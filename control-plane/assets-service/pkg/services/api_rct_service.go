/*
 * SPDX-FileCopyrightText: Copyright (c) 2026 NVIDIA CORPORATION & AFFILIATES. All rights reserved.
 * SPDX-License-Identifier: Apache-2.0
 */

package services

import (
	"context"
	"errors"
	"net/http"

	"github.com/dsx-ai-factory/run-ai/control-plane/assets-service/generated/pkg/api"
	"github.com/dsx-ai-factory/run-ai/control-plane/assets-service/pkg/database/dao"
	"github.com/dsx-ai-factory/run-ai/control-plane/assets-service/pkg/projects"
	"github.com/dsx-ai-factory/run-ai/control-plane/assets-service/pkg/query_types"
	"github.com/dsx-ai-factory/run-ai/control-plane/assets-service/pkg/services/response_patching"
	"github.com/dsx-ai-factory/run-ai/control-plane/assets-service/pkg/services/scope"
	audit "github.com/dsx-ai-factory/run-ai/control-plane/audit-service/pkg/client"
	backend "github.com/dsx-ai-factory/run-ai/control-plane/authorization/pkg/client/generated"
	"github.com/dsx-ai-factory/run-ai/control-plane/workloads-and-policies/auth"
	"github.com/dsx-ai-factory/run-ai/runai-application-packages/convutils"
	"github.com/dsx-ai-factory/run-ai/runai-common-packages/logging"
	"github.com/dsx-ai-factory/run-ai/runai-common-packages/runaictx"
)

// RctApiService is a service that implements the logic for the ResourceClaimTemplates servicer.
// This service should implement the business logic for every endpoint for the ResourceClaimTemplates API.
type RctApiService struct{}

// NewRctApiService creates a default api service.
func NewRctApiService() *RctApiService {
	return &RctApiService{}
}

// CreateRct - Create resource claim template.
func (s *RctApiService) CreateRct(ctx context.Context, request api.RctCreationRequest) (api.ImplResponse, error) {
	assetKind := api.ASSETKIND_RESOURCE_CLAIM_TEMPLATE
	audit.AuditFromCtx(ctx).WithEntityType(string(assetKind)).WithEntityName(request.Meta.Name).WithClusterId(safeStrPtr(request.Meta.ClusterId))

	validationError := ValidateAssetCreateRequestMeta(ctx, &request.Meta, assetKind)
	if validationError != nil {
		return NewErrorResponse(validationError), nil
	}

	rbacErr := auth.ValidateAccessByActionAndAssetKind(ctx, ScopeInfoOfCreateRequest(&request.Meta), string(assetKind), backend.ACTION_CREATE, convutils.Deref(request.Meta.Private, false))
	if rbacErr != nil {
		return NewErrorResponse(api.NewError(int(rbacErr.Code), rbacErr.Message, errors.New(rbacErr.Details))), nil
	}

	nameValidationQuery, err := scope.GetNameValidationQuery(
		ctx,
		request.Meta.Name,
		request.Meta.Scope,
		request.Meta.ProjectId,
		request.Meta.DepartmentId,
		request.Meta.ClusterId,
		request.Meta.Private,
		assetKind,
		nil,
	)
	if err != nil {
		return NewErrorResponse(err), nil
	}

	resultMeta, actionError := dao.CreateAsset(ctx, assetKind, request.Meta, request.Spec, nameValidationQuery)
	if actionError != nil {
		return NewErrorResponse(actionError), nil
	}

	audit.AuditFromCtx(ctx).WithEntityId(resultMeta.Id)

	result := api.Rct{
		Meta: *resultMeta,
		Spec: request.Spec,
	}

	return api.Response(http.StatusAccepted, result), nil
}

// ListRcts - List resource claim templates.
func (s *RctApiService) ListRcts(ctx context.Context, offset int32, limit int32) (api.ImplResponse, error) {
	assetKind := api.ASSETKIND_RESOURCE_CLAIM_TEMPLATE

	// this API does not work with the "classic" queryParams so just initial it to empty
	// when we work with the "standard" filter this needs to be refactored
	queryParams, queryError := query_types.NewQueryParams("", "", "", false, 0, "", "", nil)
	if queryError != nil {
		return NewErrorResponse(queryError), nil
	}

	_, authError := query_types.ApplyRbacRules(ctx, queryParams)
	if authError != nil {
		return NewErrorResponse(authError), nil
	}

	result, actionError := dao.ListRcts(ctx, offset, limit)
	if actionError != nil {
		return NewErrorResponse(actionError), nil
	}

	patchingPointers := response_patching.LoadResponsePatchingPtrs(result)

	patchErr := projects.PatchProjectNames(ctx, patchingPointers)
	if patchErr != nil {
		logging.WithCtx(runaictx.LoggerHook(ctx)).Warnf("list %s: Failed to patch project names: %v", assetKind, patchErr)
	}

	return api.Response(http.StatusOK, result), nil
}

// GetRctById - Retrieve a resource claim template.
func (s *RctApiService) GetRctById(ctx context.Context, id string) (api.ImplResponse, error) {
	asset, actionError := dao.GetRctById(ctx, id)
	if actionError != nil {
		return NewErrorResponse(actionError), nil
	}

	if _, _, authError := query_types.ApplyRbacRulesGetById(ctx, asset.Meta); authError != nil {
		return NewErrorResponse(authError), nil
	}

	return api.Response(http.StatusOK, asset), nil
}

// UpdateRctById - Update a resource claim template.
func (s *RctApiService) UpdateRctById(ctx context.Context, id string, request api.RctUpdateRequest) (api.ImplResponse, error) {
	assetKind := api.ASSETKIND_RESOURCE_CLAIM_TEMPLATE
	audit.AuditFromCtx(ctx).WithEntityType(string(assetKind)).WithEntityId(id).WithEntityName(request.Meta.Name)

	assetMeta, getError := dao.GetAssetMetaById(ctx, id)
	if getError != nil {
		return NewErrorResponse(getError), nil
	}
	audit.AuditFromCtx(ctx).WithClusterId(safeStrPtr(assetMeta.ClusterId))

	rbacErr := auth.ValidateAccessByActionAndAssetKind(ctx, ScopeInfoOfDbMeta(assetMeta), string(assetKind), backend.ACTION_UPDATE, convutils.Deref(assetMeta.Private, false))
	if rbacErr != nil {
		return NewErrorResponse(api.NewError(int(rbacErr.Code), rbacErr.Message, errors.New(rbacErr.Details))), nil
	}

	nameValidationQuery, err := scope.GetNameValidationQuery(
		ctx,
		request.Meta.Name,
		assetMeta.Scope,
		assetMeta.ProjectId,
		assetMeta.DepartmentId,
		assetMeta.ClusterId,
		assetMeta.Private,
		assetKind,
		&id,
	)
	if err != nil {
		return NewErrorResponse(err), nil
	}

	result, actionError := dao.UpdateRct(ctx, id, request, nameValidationQuery)
	if actionError != nil {
		return NewErrorResponse(actionError), nil
	}

	return api.Response(http.StatusAccepted, result), nil
}

// DeleteRctById - Delete a resource claim template.
func (s *RctApiService) DeleteRctById(ctx context.Context, id string) (api.ImplResponse, error) {
	assetKind := api.ASSETKIND_RESOURCE_CLAIM_TEMPLATE
	audit.AuditFromCtx(ctx).WithEntityType(string(assetKind)).WithEntityId(id)

	assetMeta, getError := dao.GetAssetMetaById(ctx, id)
	if getError != nil {
		return NewErrorResponse(getError), nil
	}
	audit.AuditFromCtx(ctx).WithEntityName(assetMeta.Name).WithClusterId(safeStrPtr(assetMeta.ClusterId))

	rbacErr := auth.ValidateAccessByActionAndAssetKind(ctx, ScopeInfoOfDbMeta(assetMeta), string(assetKind), backend.ACTION_DELETE, convutils.Deref(assetMeta.Private, false))
	if rbacErr != nil {
		return NewErrorResponse(api.NewError(int(rbacErr.Code), rbacErr.Message, errors.New(rbacErr.Details))), nil
	}

	if actionError := dao.DeleteRct(ctx, assetMeta, dao.HardDelete); actionError != nil {
		return NewErrorResponse(actionError), nil
	}

	return NewDeletionResponse(), nil
}
