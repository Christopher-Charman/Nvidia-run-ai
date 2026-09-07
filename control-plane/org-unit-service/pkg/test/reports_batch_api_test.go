/*
 * SPDX-FileCopyrightText: Copyright (c) 2026 NVIDIA CORPORATION & AFFILIATES. All rights reserved.
 * SPDX-License-Identifier: Apache-2.0
 */

package testing

import (
	"context"
	"fmt"
	"net/http"

	authorization_api "github.com/dsx-ai-factory/run-ai/control-plane/authorization/pkg/client/generated"
	authorization_utils "github.com/dsx-ai-factory/run-ai/control-plane/authorization/pkg/client/helpers"
	metrics_swagger "github.com/dsx-ai-factory/run-ai/control-plane/metrics-service/pkg/client/generated"
	metrics_mocks "github.com/dsx-ai-factory/run-ai/control-plane/metrics-service/pkg/client/mocks"
	"github.com/dsx-ai-factory/run-ai/control-plane/org-unit-service/generated/pkg/api"
	"github.com/dsx-ai-factory/run-ai/control-plane/org-unit-service/pkg/services/reports"
	"github.com/google/uuid"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"
)

var _ = Describe("BatchDeleteReports API tests", func() {
	ctx := SetContext(TestTenant, TestUser)

	makeSvc := func(mc *metrics_mocks.MetricsClientContract) reports.ReportsApiService {
		return reports.ReportsApiService{MetricsClient: mc}
	}

	makeReport := func(id, createdBy string) metrics_swagger.Report {
		return metrics_swagger.Report{Id: id, CreatedBy: createdBy}
	}

	makeBatchGetResp := func(succeeded []metrics_swagger.Report, failed []metrics_swagger.BatchResponseFailureItem) *metrics_swagger.ReportsBatchGetResponse {
		return &metrics_swagger.ReportsBatchGetResponse{
			Id:        uuid.NewString(),
			Succeeded: succeeded,
			Failed:    failed,
		}
	}

	makeBatchDeleteResp := func(batchId string, accepted []string, failed []metrics_swagger.BatchResponseFailureItem) *metrics_swagger.BatchResponseAsync {
		return &metrics_swagger.BatchResponseAsync{
			Id:       batchId,
			Accepted: accepted,
			Failed:   failed,
		}
	}

	Context("request validation", func() {
		It("rejects empty ids with 400", func() {
			mc := &metrics_mocks.MetricsClientContract{}
			svc := makeSvc(mc)

			resp, err := svc.BatchDeleteReports(ctx, api.BatchDeleteReportsRequest{Ids: []string{}})

			Expect(err).To(BeNil())
			Expect(resp.Code).To(Equal(http.StatusBadRequest))
			mc.AssertNumberOfCalls(GinkgoT(), "BatchGetReportsByIds", 0)
		})

		It("rejects duplicate ids with 400", func() {
			id := uuid.NewString()
			mc := &metrics_mocks.MetricsClientContract{}
			svc := makeSvc(mc)

			resp, err := svc.BatchDeleteReports(ctx, api.BatchDeleteReportsRequest{Ids: []string{id, id}})

			Expect(err).To(BeNil())
			Expect(resp.Code).To(Equal(http.StatusBadRequest))
			mc.AssertNumberOfCalls(GinkgoT(), "BatchGetReportsByIds", 0)
		})
	})

	Context("all reports owned by the caller", func() {
		It("returns 202 with all reports accepted", func() {
			id1, id2 := uuid.NewString(), uuid.NewString()
			batchOpId := uuid.NewString()

			mc := &metrics_mocks.MetricsClientContract{}
			mc.On("BatchGetReportsByIds", mock.Anything, []string{id1, id2}).Return(
				makeBatchGetResp(
					[]metrics_swagger.Report{makeReport(id1, TestUser), makeReport(id2, TestUser)},
					nil,
				),
				http.StatusOK, nil,
			)
			mc.On("BatchDeleteReports", mock.Anything, []string{id1, id2}).Return(
				makeBatchDeleteResp(batchOpId, []string{id1, id2}, nil),
				http.StatusAccepted, nil,
			)

			svc := makeSvc(mc)
			resp, err := svc.BatchDeleteReports(ctx, api.BatchDeleteReportsRequest{Ids: []string{id1, id2}})

			Expect(err).To(BeNil())
			Expect(resp.Code).To(Equal(http.StatusAccepted))
			body := resp.Body.(api.BatchDeleteReportsResponse)
			Expect(body.Id).To(Equal(batchOpId))
			Expect(body.Accepted).To(ConsistOf(id1, id2))
			Expect(body.Failed).To(BeEmpty())
			Expect(body.TotalRequested).To(Equal(int32(2)))
			Expect(body.TotalAccepted).To(Equal(int32(2)))
			Expect(body.TotalFailed).To(Equal(int32(0)))
			mc.AssertExpectations(GinkgoT())
		})
	})

	Context("some reports not found in metrics-service", func() {
		It("not-found reports appear in failed, found ones proceed to delete", func() {
			foundId := uuid.NewString()
			missingId := uuid.NewString()
			batchOpId := uuid.NewString()

			mc := &metrics_mocks.MetricsClientContract{}
			mc.On("BatchGetReportsByIds", mock.Anything, []string{foundId, missingId}).Return(
				makeBatchGetResp(
					[]metrics_swagger.Report{makeReport(foundId, TestUser)},
					[]metrics_swagger.BatchResponseFailureItem{{Id: missingId, Code: http.StatusNotFound, Message: "not found"}},
				),
				http.StatusOK, nil,
			)
			mc.On("BatchDeleteReports", mock.Anything, []string{foundId}).Return(
				makeBatchDeleteResp(batchOpId, []string{foundId}, nil),
				http.StatusAccepted, nil,
			)

			svc := makeSvc(mc)
			resp, err := svc.BatchDeleteReports(ctx, api.BatchDeleteReportsRequest{Ids: []string{foundId, missingId}})

			Expect(err).To(BeNil())
			Expect(resp.Code).To(Equal(http.StatusAccepted))
			body := resp.Body.(api.BatchDeleteReportsResponse)
			Expect(body.Id).To(Equal(batchOpId))
			Expect(body.Accepted).To(ConsistOf(foundId))
			Expect(body.Failed).To(HaveLen(1))
			Expect(body.Failed[0].Id).To(Equal(missingId))
			Expect(body.TotalRequested).To(Equal(int32(2)))
			Expect(body.TotalAccepted).To(Equal(int32(1)))
			Expect(body.TotalFailed).To(Equal(int32(1)))
			mc.AssertExpectations(GinkgoT())
		})
	})

	Context("permission checks", func() {
		It("rejects report owned by another user when caller has no tenant scope", func() {
			othersId := uuid.NewString()

			mc := &metrics_mocks.MetricsClientContract{}
			mc.On("BatchGetReportsByIds", mock.Anything, []string{othersId}).Return(
				makeBatchGetResp(
					[]metrics_swagger.Report{makeReport(othersId, "other-user@run.ai")},
					nil,
				),
				http.StatusOK, nil,
			)

			svc := makeSvc(mc)
			resp, err := svc.BatchDeleteReports(ctx, api.BatchDeleteReportsRequest{Ids: []string{othersId}})

			Expect(err).To(BeNil())
			Expect(resp.Code).To(Equal(http.StatusAccepted))
			body := resp.Body.(api.BatchDeleteReportsResponse)
			Expect(body.Id).NotTo(BeEmpty())
			Expect(body.Accepted).To(BeEmpty())
			Expect(body.Failed).To(HaveLen(1))
			Expect(body.Failed[0].Id).To(Equal(othersId))
			Expect(int(body.Failed[0].Code)).To(Equal(http.StatusForbidden))
			mc.AssertNumberOfCalls(GinkgoT(), "BatchDeleteReports", 0)
		})

		It("allows tenant admin to delete another user's report", func() {
			othersId := uuid.NewString()
			batchOpId := uuid.NewString()
			tenantAdminCtx := context.WithValue(ctx, authorization_utils.PermittedScopesKey,
				&authorization_api.PermittedScopes{Tenants: []string{TestTenant}})

			mc := &metrics_mocks.MetricsClientContract{}
			mc.On("BatchGetReportsByIds", mock.Anything, []string{othersId}).Return(
				makeBatchGetResp(
					[]metrics_swagger.Report{makeReport(othersId, "other-user@run.ai")},
					nil,
				),
				http.StatusOK, nil,
			)
			mc.On("BatchDeleteReports", mock.Anything, []string{othersId}).Return(
				makeBatchDeleteResp(batchOpId, []string{othersId}, nil),
				http.StatusAccepted, nil,
			)

			svc := makeSvc(mc)
			resp, err := svc.BatchDeleteReports(tenantAdminCtx, api.BatchDeleteReportsRequest{Ids: []string{othersId}})

			Expect(err).To(BeNil())
			Expect(resp.Code).To(Equal(http.StatusAccepted))
			body := resp.Body.(api.BatchDeleteReportsResponse)
			Expect(body.Id).To(Equal(batchOpId))
			Expect(body.Accepted).To(ConsistOf(othersId))
			Expect(body.Failed).To(BeEmpty())
			mc.AssertExpectations(GinkgoT())
		})
	})

	Context("all reports forbidden", func() {
		It("response Id is a generated uuid — no batch op was submitted to metrics-service", func() {
			othersId := uuid.NewString()

			mc := &metrics_mocks.MetricsClientContract{}
			mc.On("BatchGetReportsByIds", mock.Anything, []string{othersId}).Return(
				makeBatchGetResp(
					[]metrics_swagger.Report{makeReport(othersId, "other-user@run.ai")},
					nil,
				),
				http.StatusOK, nil,
			)

			svc := makeSvc(mc)
			resp, err := svc.BatchDeleteReports(ctx, api.BatchDeleteReportsRequest{Ids: []string{othersId}})

			Expect(err).To(BeNil())
			Expect(resp.Code).To(Equal(http.StatusAccepted))
			body := resp.Body.(api.BatchDeleteReportsResponse)
			Expect(body.Id).NotTo(BeEmpty())
			Expect(body.Accepted).To(BeEmpty())
			Expect(body.Failed).To(HaveLen(1))
			mc.AssertNumberOfCalls(GinkgoT(), "BatchDeleteReports", 0)
		})
	})

	Context("metrics-service returns 5xx during delete", func() {
		It("folds permitted ids into failed and still returns 202", func() {
			id1, id2 := uuid.NewString(), uuid.NewString()

			mc := &metrics_mocks.MetricsClientContract{}
			mc.On("BatchGetReportsByIds", mock.Anything, []string{id1, id2}).Return(
				makeBatchGetResp(
					[]metrics_swagger.Report{makeReport(id1, TestUser), makeReport(id2, TestUser)},
					nil,
				),
				http.StatusOK, nil,
			)
			mc.On("BatchDeleteReports", mock.Anything, []string{id1, id2}).Return(
				(*metrics_swagger.BatchResponseAsync)(nil),
				http.StatusInternalServerError,
				fmt.Errorf("metrics-service unavailable"),
			)

			svc := makeSvc(mc)
			resp, err := svc.BatchDeleteReports(ctx, api.BatchDeleteReportsRequest{Ids: []string{id1, id2}})

			Expect(err).To(BeNil())
			Expect(resp.Code).To(Equal(http.StatusAccepted))
			body := resp.Body.(api.BatchDeleteReportsResponse)
			Expect(body.Id).NotTo(BeEmpty())
			Expect(body.Accepted).To(BeEmpty())
			Expect(body.Failed).To(HaveLen(2))
			Expect(body.TotalRequested).To(Equal(int32(2)))
			Expect(body.TotalAccepted).To(Equal(int32(0)))
			Expect(body.TotalFailed).To(Equal(int32(2)))
			mc.AssertExpectations(GinkgoT())
		})
	})

	Context("mixed: some not found, some forbidden, some accepted", func() {
		It("partitions results correctly across all outcome categories", func() {
			ownId := uuid.NewString()
			othersId := uuid.NewString()
			missingId := uuid.NewString()
			batchOpId := uuid.NewString()

			mc := &metrics_mocks.MetricsClientContract{}
			mc.On("BatchGetReportsByIds", mock.Anything, []string{ownId, othersId, missingId}).Return(
				makeBatchGetResp(
					[]metrics_swagger.Report{
						makeReport(ownId, TestUser),
						makeReport(othersId, "other-user@run.ai"),
					},
					[]metrics_swagger.BatchResponseFailureItem{{Id: missingId, Code: http.StatusNotFound, Message: "not found"}},
				),
				http.StatusOK, nil,
			)
			mc.On("BatchDeleteReports", mock.Anything, []string{ownId}).Return(
				makeBatchDeleteResp(batchOpId, []string{ownId}, nil),
				http.StatusAccepted, nil,
			)

			svc := makeSvc(mc)
			resp, err := svc.BatchDeleteReports(ctx, api.BatchDeleteReportsRequest{Ids: []string{ownId, othersId, missingId}})

			Expect(err).To(BeNil())
			Expect(resp.Code).To(Equal(http.StatusAccepted))
			body := resp.Body.(api.BatchDeleteReportsResponse)
			Expect(body.Id).To(Equal(batchOpId))
			Expect(body.Accepted).To(ConsistOf(ownId))
			Expect(body.Failed).To(HaveLen(2))
			Expect(body.TotalRequested).To(Equal(int32(3)))
			Expect(body.TotalAccepted).To(Equal(int32(1)))
			Expect(body.TotalFailed).To(Equal(int32(2)))
			mc.AssertExpectations(GinkgoT())
		})
	})
})
