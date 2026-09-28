/*
 * SPDX-FileCopyrightText: Copyright (c) 2026 NVIDIA CORPORATION & AFFILIATES. All rights reserved.
 * SPDX-License-Identifier: Apache-2.0
 */

package dao

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/dsx-ai-factory/run-ai/control-plane/assets-service/generated/pkg/api"
	"github.com/dsx-ai-factory/run-ai/control-plane/assets-service/pkg/database"
	"github.com/dsx-ai-factory/run-ai/control-plane/assets-service/pkg/database/dbmodel"
	"github.com/dsx-ai-factory/run-ai/control-plane/assets-service/pkg/database/dbtables"
	"github.com/dsx-ai-factory/run-ai/control-plane/workloads-and-policies/auth"
	"github.com/dsx-ai-factory/run-ai/runai-common-packages/runaidb"
	"github.com/dsx-ai-factory/run-ai/runai-common-packages/runaidb/orm"
	"github.com/google/uuid"
)

// CreateAsset atomically writes a k8s-object-backed asset: it inserts the
// assets_metadata row and the assets_spec row in a single transaction.
// The spec is any JSON-marshalable value (typically the typed swagger struct);
// the DAO converts it to the map-shaped storage type so callers do not have
// to marshal the spec themselves.
func CreateAsset(
	ctx context.Context,
	kind api.AssetKind,
	meta api.AssetCreationRequest,
	spec any,
	nameValidationQuery *database.ParametrizedQuery,
) (*api.AssetMeta, *api.Error) {
	tx := runaidb.DB().WithContext(ctx).Begin()
	committedTx := false
	defer orm.FinalizeTransaction(tx, &committedTx)

	result, err := createAssetInTransaction(ctx, tx, kind, meta, spec, nameValidationQuery)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, database.Error(database.COMMIT_CREATIONS, kind).WithDetails(err)
	}
	committedTx = true

	return result, nil
}

func createAssetInTransaction(
	ctx context.Context,
	tx orm.DB,
	kind api.AssetKind,
	meta api.AssetCreationRequest,
	spec any,
	nameValidationQuery *database.ParametrizedQuery,
) (*api.AssetMeta, *api.Error) {
	id := uuid.New().String()

	assetMeta, err := createAssetMeta(ctx, tx, id, kind, meta, nameValidationQuery)
	if err != nil {
		if strings.Contains(err.Error(), alreadyExistsError) {
			return nil, database.ErrorWithCode(database.CREATE, kind, http.StatusConflict).WithDetails(err)
		}
		return nil, database.Error(database.CREATE, kind).WithDetails(err)
	}

	specMap, marshalErr := toJsonBMap(spec)
	if marshalErr != nil {
		return nil, database.ErrorWithCode(database.CREATE, kind, http.StatusBadRequest).WithDetails(marshalErr)
	}

	specRec := dbmodel.AssetSpecRec{
		AssetId:  id,
		TenantId: auth.TenantIdOf(ctx),
		Spec:     specMap,
	}

	if _, err := tx.Create(specRec); err != nil {
		return nil, database.Error(database.CREATE, kind).WithDetails(err)
	}

	return assetMeta, nil
}

// toJsonBMap converts an arbitrary JSON-marshalable value into the map-shaped
// storage representation used by the assets_spec.spec column. Round-tripping
// through JSON is intentional: it normalizes typed structs, pointers, and any
// existing map into the exact `map[string]any` shape the storage layer stores.
func toJsonBMap(v any) (dbmodel.JsonBMapFromString, error) {
	bytes, err := json.Marshal(v)
	if err != nil {
		return nil, fmt.Errorf("marshal spec to json: %w", err)
	}
	var m dbmodel.JsonBMapFromString
	if err := json.Unmarshal(bytes, &m); err != nil {
		return nil, fmt.Errorf("unmarshal spec json into map: %w", err)
	}
	return m, nil
}

// fromJsonBMap is the inverse of toJsonBMap: it reads the map-shaped stored
// spec into the caller's typed swagger struct via a JSON round-trip. Used by
// per-kind list/get DAOs (e.g. ListRcts) to materialize a typed response.
func fromJsonBMap(m dbmodel.JsonBMapFromString, out any) error {
	bytes, err := json.Marshal(m)
	if err != nil {
		return fmt.Errorf("marshal stored spec map: %w", err)
	}
	if err := json.Unmarshal(bytes, out); err != nil {
		return fmt.Errorf("unmarshal stored spec into %T: %w", out, err)
	}
	return nil
}

// ListRcts returns a paginated list of resource claim template assets. The
// offset/limit contract matches the swagger: fetch limit+1 rows so we can set
// RctListResponse.Next when the next page exists (i.e. when the returned page
// is full). Mirrors ListNFSAssets structurally, plus pagination + spec
// unmarshaling from the shared assets_spec table.
func ListRcts(ctx context.Context, offset, limit int32) (*api.RctListResponse, *api.Error) {
	combinedList := make([]dbmodel.AssetSpecCombinedRec, 0)
	query := dbmodel.AssetMetaRec{
		AssetMeta: dbmodel.AssetMeta{
			Kind: api.ASSETKIND_RESOURCE_CLAIM_TEMPLATE,
		},
	}

	clonePresetAssetsToTenant(ctx, auth.TenantIdOf(ctx))

	result, err := GetAssetMeta(ctx, runaidb.DB().WithContext(ctx), query, nil,
		fmt.Sprintf("%s.%s", dbtables.TableNameAssetsSpec, "spec"))
	if err != nil {
		return nil, err
	}

	// Fetch limit+1 so we can detect "is there a next page" without a count query.
	fetchLimit := int(limit) + 1
	_, scanErr := result.
		Order(&query.AssetMeta.Id).
		Offset(int(offset)).
		Limit(fetchLimit).
		Scan(&combinedList)
	if scanErr != nil {
		return nil, api.NewError(http.StatusServiceUnavailable, "error occurred while attempting to list rct assets", scanErr)
	}

	response := api.RctListResponse{
		Entries: make([]api.Rct, 0, len(combinedList)),
	}

	// If we got more than limit rows, there is a next page: trim the extra row
	// and set the next offset. Otherwise leave Next as nil (no more pages).
	if int32(len(combinedList)) > limit {
		combinedList = combinedList[:limit]
		nextOffset := offset + limit
		response.Next = nextOffset
	}

	for _, row := range combinedList {
		var rctSpec api.RctSpec
		if err := fromJsonBMap(row.Spec, &rctSpec); err != nil {
			return nil, database.Error(database.LIST, api.ASSETKIND_RESOURCE_CLAIM_TEMPLATE).WithDetails(err)
		}
		response.Entries = append(response.Entries, api.Rct{
			Meta: row.AssetMeta,
			Spec: rctSpec,
		})
	}

	return &response, nil
}

// GetRctById fetches a single resource claim template joined from
// assets_metadata + assets_spec, unmarshals the JSON spec into the typed
// RctSpec, and returns the typed asset. Mirrors GetNFSAssetById.
func GetRctById(ctx context.Context, id string) (*api.Rct, *api.Error) {
	var combinedRec dbmodel.AssetSpecCombinedRec
	query := dbmodel.AssetMetaRec{
		AssetMeta: dbmodel.AssetMeta{
			Kind: api.ASSETKIND_RESOURCE_CLAIM_TEMPLATE,
			Id:   id,
		},
	}

	result, err := GetAssetMeta(ctx, runaidb.DB().WithContext(ctx), query, nil,
		fmt.Sprintf("%s.%s", dbtables.TableNameAssetsSpec, "spec"))
	if err != nil {
		return nil, err
	}

	scanRows, scanErr := result.Scan(&combinedRec)
	if scanErr != nil || scanRows == 0 {
		return nil, database.Error(database.GET, api.ASSETKIND_RESOURCE_CLAIM_TEMPLATE).WithDetailsFromDbResult(scanRows, scanErr)
	}

	var rctSpec api.RctSpec
	if err := fromJsonBMap(combinedRec.Spec, &rctSpec); err != nil {
		return nil, database.Error(database.GET, api.ASSETKIND_RESOURCE_CLAIM_TEMPLATE).WithDetails(err)
	}

	return &api.Rct{
		Meta: combinedRec.AssetMeta,
		Spec: rctSpec,
	}, nil
}

// UpdateRct updates the assets_metadata row and overwrites the JSON spec in
// assets_spec for a resource claim template, atomically. On success it clears
// any cached policy compliance for the asset and notifies watchers (agent
// sync) via NotifyAssetModified. Mirrors UpdateNFSAsset.
func UpdateRct(ctx context.Context, id string, request api.RctUpdateRequest, nameValidationQuery *database.ParametrizedQuery) (*api.Rct, *api.Error) {
	tx := runaidb.DB().WithContext(ctx).Begin()
	committedTx := false
	defer orm.FinalizeTransaction(tx, &committedTx)

	result, assetMeta, err := updateRctInTransaction(ctx, tx, id, request, nameValidationQuery)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, database.Error(database.COMMIT_UPDATES, api.ASSETKIND_RESOURCE_CLAIM_TEMPLATE).WithDetails(err)
	}
	committedTx = true

	NotifyAssetModified(ctx, id, assetMeta.Kind, ScopeInfoOfDbMeta(assetMeta))

	return result, nil
}

func updateRctInTransaction(ctx context.Context, tx orm.DB, id string, request api.RctUpdateRequest,
	nameValidationQuery *database.ParametrizedQuery,
) (*api.Rct, *dbmodel.AssetMeta, *api.Error) {
	assetMeta, rowsAffected, apiErr, dbErr := updateAssetMeta(ctx, tx, id, request.Meta, nameValidationQuery, nil)
	if apiErr != nil {
		return nil, nil, apiErr
	}
	if dbErr != nil || rowsAffected == 0 {
		return nil, nil, database.Error(database.UPDATE, api.ASSETKIND_RESOURCE_CLAIM_TEMPLATE).WithDetailsFromDbResult(rowsAffected, dbErr)
	}

	specMap, marshalErr := toJsonBMap(request.Spec)
	if marshalErr != nil {
		return nil, assetMeta, database.ErrorWithCode(database.UPDATE, api.ASSETKIND_RESOURCE_CLAIM_TEMPLATE, http.StatusBadRequest).WithDetails(marshalErr)
	}

	recId := dbmodel.AssetSpecRec{AssetId: id}
	update := dbmodel.AssetSpecUpdate{Spec: specMap}

	saveRows, saveErr := tx.Where(recId).Save(&update)
	if saveErr != nil {
		return nil, assetMeta, database.Error(database.UPDATE, api.ASSETKIND_RESOURCE_CLAIM_TEMPLATE).WithDetailsFromDbResult(saveRows, saveErr)
	}

	if policyErr := deletePolicyComplianceOfAsset(ctx, tx, assetMeta.Id); policyErr != nil {
		return nil, assetMeta, policyErr
	}

	return &api.Rct{
		Meta: assetMeta.ToApiAssetMeta(),
		Spec: request.Spec,
	}, assetMeta, nil
}

// DeleteRct removes a resource claim template. Deletion type follows the
// existing asset pattern (see DeleteNFSAsset / DeleteComputeAsset): SoftDelete
// marks assets_metadata but keeps the spec (so agent sync can still see the
// asset while the delete propagates to the cluster); HardDelete removes the
// spec row plus policy compliance and workload asset references.
func DeleteRct(ctx context.Context, assetMeta *dbmodel.AssetMeta, deletionType DeletionType) *api.Error {
	tx := runaidb.DB().WithContext(ctx).Begin()
	committedTx := false
	defer orm.FinalizeTransaction(tx, &committedTx)

	err := deleteRctInTransaction(ctx, tx, assetMeta.Id, deletionType)
	if err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return database.Error(database.COMMIT_DELETIONS, api.ASSETKIND_RESOURCE_CLAIM_TEMPLATE).WithDetails(err)
	}
	committedTx = true

	NotifyAssetDeleted(ctx, assetMeta.Id, assetMeta.Kind, ScopeInfoOfDbMeta(assetMeta))

	return nil
}

func deleteRctInTransaction(ctx context.Context, tx orm.DB, id string, deletionType DeletionType) *api.Error {
	rowsAffected, apiErr, dbErr := deleteAssetMeta(ctx, tx, id, deletionType)
	if apiErr != nil {
		return apiErr
	}
	if dbErr != nil || rowsAffected == 0 {
		return database.Error(database.DELETE, api.ASSETKIND_RESOURCE_CLAIM_TEMPLATE).WithDetailsFromDbResult(rowsAffected, dbErr)
	}

	if deletionType == SoftDelete {
		return nil
	}

	recId := dbmodel.AssetSpecRec{AssetId: id}
	delRows, delErr := tx.Where(&recId).Delete(&recId)
	if delErr != nil || delRows == 0 {
		return database.Error(database.DELETE, api.ASSETKIND_RESOURCE_CLAIM_TEMPLATE).WithDetailsFromDbResult(delRows, delErr)
	}

	if policyErr := deletePolicyComplianceOfAsset(ctx, tx, id); policyErr != nil {
		return policyErr
	}

	if err := deleteWorkloadAssetsByAssetId(ctx, tx, id); err != nil {
		return err
	}

	return nil
}
