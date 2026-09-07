package internal

import (
	"context"

	openapi "github.com/dsx-ai-factory/run-ai/runai-clients/workloads-client/generated"
)

type WorkloadsExternalClient struct {
	WorkloadsClient
}

type WorkloadsClient struct {
	ApiClient *openapi.APIClient
}

func (wc *WorkloadsClient) GetInferenceAncestors(ctx context.Context, resourceId string) (*openapi.Ancestors, error) {
	request := wc.ApiClient.InferencesAPI.GetInferenceAncestors(ctx, resourceId)

	ancestors, _, err := wc.ApiClient.InferencesAPI.GetInferenceAncestorsExecute(request)
	if err != nil {
		return nil, err
	}

	return ancestors, nil
}

func (wc *WorkloadsClient) GetAncestors(ctx context.Context, resourceId string) (*openapi.Ancestors, error) {
	request := wc.ApiClient.WorkloadsAPI.GetWorkloadAncestors(ctx, resourceId)

	ancestors, _, err := wc.ApiClient.WorkloadsAPI.GetWorkloadAncestorsExecute(request)
	if err != nil {
		return nil, err
	}

	return ancestors, nil
}

func (wc *WorkloadsClient) GetWorkloads(ctx context.Context,
	limit, offset int32, filter []string, includeDeleted bool, sortBy string, sortOrder string) (*openapi.GetWorkloads200Response, error) {
	request := wc.ApiClient.WorkloadsAPI.GetWorkloads(ctx)
	request = request.Limit(limit)
	request = request.Offset(offset)
	request = request.Deleted(includeDeleted)

	if len(filter) != 0 {
		request = request.FilterBy(filter)
	}

	if len(sortBy) != 0 {
		request = request.SortBy(sortBy)
	}

	if len(sortOrder) != 0 {
		request = request.SortOrder(sortOrder)
	}

	workloadList, _, err := request.Execute()
	if err != nil {
		return nil, err
	}

	return workloadList, nil
}
