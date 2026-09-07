# \WorkloadTemplatesAPI

All URIs are relative to *https://app.run.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CountTemplates**](WorkloadTemplatesAPI.md#CountTemplates) | **Get** /api/v1/workloads/templates/count | Count templates.
[**CreateDistributedTemplate**](WorkloadTemplatesAPI.md#CreateDistributedTemplate) | **Post** /api/v1/workloads/templates/distributed | Create a distributed template.
[**CreateInferenceTemplate**](WorkloadTemplatesAPI.md#CreateInferenceTemplate) | **Post** /api/v1/workloads/templates/inferences | Create a inference template.
[**CreateTrainingTemplate**](WorkloadTemplatesAPI.md#CreateTrainingTemplate) | **Post** /api/v1/workloads/templates/trainings | Create a training template.
[**CreateWorkspaceTemplate**](WorkloadTemplatesAPI.md#CreateWorkspaceTemplate) | **Post** /api/v1/workloads/templates/workspaces | Create a workspace template.
[**DeleteDistributedTemplate**](WorkloadTemplatesAPI.md#DeleteDistributedTemplate) | **Delete** /api/v1/workloads/templates/distributed/{templateId} | Delete a distributed template.
[**DeleteInferenceTemplate**](WorkloadTemplatesAPI.md#DeleteInferenceTemplate) | **Delete** /api/v1/workloads/templates/inferences/{templateId} | Delete a inference template.
[**DeleteTrainingTemplate**](WorkloadTemplatesAPI.md#DeleteTrainingTemplate) | **Delete** /api/v1/workloads/templates/trainings/{templateId} | Delete a training template.
[**DeleteWorkspaceTemplate**](WorkloadTemplatesAPI.md#DeleteWorkspaceTemplate) | **Delete** /api/v1/workloads/templates/workspaces/{templateId} | Delete a workspace template.
[**GetDistributedTemplate**](WorkloadTemplatesAPI.md#GetDistributedTemplate) | **Get** /api/v1/workloads/templates/distributed/{templateId} | Get a distributed template.
[**GetInferenceTemplate**](WorkloadTemplatesAPI.md#GetInferenceTemplate) | **Get** /api/v1/workloads/templates/inferences/{templateId} | Get a inference template.
[**GetTemplateById**](WorkloadTemplatesAPI.md#GetTemplateById) | **Get** /api/v1/workloads/templates/{templateId} | Get a template, any type.
[**GetTrainingTemplate**](WorkloadTemplatesAPI.md#GetTrainingTemplate) | **Get** /api/v1/workloads/templates/trainings/{templateId} | Get a training template.
[**GetWorkspaceTemplate**](WorkloadTemplatesAPI.md#GetWorkspaceTemplate) | **Get** /api/v1/workloads/templates/workspaces/{templateId} | Get a workspace template.
[**ListDistributedTemplates**](WorkloadTemplatesAPI.md#ListDistributedTemplates) | **Get** /api/v1/workloads/templates/distributed | List distributed templates.
[**ListInferenceTemplates**](WorkloadTemplatesAPI.md#ListInferenceTemplates) | **Get** /api/v1/workloads/templates/inferences | List inference templates.
[**ListTemplates**](WorkloadTemplatesAPI.md#ListTemplates) | **Get** /api/v1/workloads/templates | List templates.
[**ListTrainingTemplates**](WorkloadTemplatesAPI.md#ListTrainingTemplates) | **Get** /api/v1/workloads/templates/trainings | List training templates.
[**ListWorkspaceTemplates**](WorkloadTemplatesAPI.md#ListWorkspaceTemplates) | **Get** /api/v1/workloads/templates/workspaces | List workspace templates.
[**PatchDistributedTemplate**](WorkloadTemplatesAPI.md#PatchDistributedTemplate) | **Patch** /api/v1/workloads/templates/distributed/{templateId} | Patch a distributed template.
[**PatchInferenceTemplate**](WorkloadTemplatesAPI.md#PatchInferenceTemplate) | **Patch** /api/v1/workloads/templates/inferences/{templateId} | Patch a inference template.
[**PatchTrainingTemplate**](WorkloadTemplatesAPI.md#PatchTrainingTemplate) | **Patch** /api/v1/workloads/templates/trainings/{templateId} | Patch a training template.
[**PatchWorkspaceTemplate**](WorkloadTemplatesAPI.md#PatchWorkspaceTemplate) | **Patch** /api/v1/workloads/templates/workspaces/{templateId} | Patch a workspace template.
[**UpdateDistributedTemplate**](WorkloadTemplatesAPI.md#UpdateDistributedTemplate) | **Put** /api/v1/workloads/templates/distributed/{templateId} | Update a distributed template.
[**UpdateInferenceTemplate**](WorkloadTemplatesAPI.md#UpdateInferenceTemplate) | **Put** /api/v1/workloads/templates/inferences/{templateId} | Update a inference template.
[**UpdateTrainingTemplate**](WorkloadTemplatesAPI.md#UpdateTrainingTemplate) | **Put** /api/v1/workloads/templates/trainings/{templateId} | Update a training template.
[**UpdateWorkspaceTemplate**](WorkloadTemplatesAPI.md#UpdateWorkspaceTemplate) | **Put** /api/v1/workloads/templates/workspaces/{templateId} | Update a workspace template.



## CountTemplates

> CountResponse CountTemplates(ctx).FilterBy(filterBy).Search(search).IncludeAncestors(includeAncestors).Execute()

Count templates.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dsx-ai-factory/run-ai/runai-clients/workloads-client"
)

func main() {
	filterBy := []string{"Inner_example"} // []string | Filter results by a parameter. Use the format field-name operator value. Operators are `==` Equals, `!=` Not equals, `<=` Less than or equal, `>=` Greater than or equal, `=@` contains, `!@` Does not contain, `=^` Starts with and `=$` Ends with. Dates are in ISO 8601 timestamp format and available for operators `==`, `!=`, `<=` and `>=`. (optional)
	search := "test project" // string | Filter results by a free text search. (optional)
	includeAncestors := true // bool | When filtering by projectId, set includeAncestors=true to include templates from the project's parent department and cluster in the response. Any reason we are calling this includeAncestors (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkloadTemplatesAPI.CountTemplates(context.Background()).FilterBy(filterBy).Search(search).IncludeAncestors(includeAncestors).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkloadTemplatesAPI.CountTemplates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CountTemplates`: CountResponse
	fmt.Fprintf(os.Stdout, "Response from `WorkloadTemplatesAPI.CountTemplates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCountTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterBy** | **[]string** | Filter results by a parameter. Use the format field-name operator value. Operators are &#x60;&#x3D;&#x3D;&#x60; Equals, &#x60;!&#x3D;&#x60; Not equals, &#x60;&lt;&#x3D;&#x60; Less than or equal, &#x60;&gt;&#x3D;&#x60; Greater than or equal, &#x60;&#x3D;@&#x60; contains, &#x60;!@&#x60; Does not contain, &#x60;&#x3D;^&#x60; Starts with and &#x60;&#x3D;$&#x60; Ends with. Dates are in ISO 8601 timestamp format and available for operators &#x60;&#x3D;&#x3D;&#x60;, &#x60;!&#x3D;&#x60;, &#x60;&lt;&#x3D;&#x60; and &#x60;&gt;&#x3D;&#x60;. | 
 **search** | **string** | Filter results by a free text search. | 
 **includeAncestors** | **bool** | When filtering by projectId, set includeAncestors&#x3D;true to include templates from the project&#39;s parent department and cluster in the response. Any reason we are calling this includeAncestors | 

### Return type

[**CountResponse**](CountResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDistributedTemplate

> DistributedTemplate CreateDistributedTemplate(ctx).DistributedTemplateCreationRequest(distributedTemplateCreationRequest).Execute()

Create a distributed template.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dsx-ai-factory/run-ai/runai-clients/workloads-client"
)

func main() {
	distributedTemplateCreationRequest := *openapiclient.NewDistributedTemplateCreationRequest(*openapiclient.NewTemplateCreationMetaFields("my-template-name", openapiclient.ScopeType("system"))) // DistributedTemplateCreationRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkloadTemplatesAPI.CreateDistributedTemplate(context.Background()).DistributedTemplateCreationRequest(distributedTemplateCreationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkloadTemplatesAPI.CreateDistributedTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDistributedTemplate`: DistributedTemplate
	fmt.Fprintf(os.Stdout, "Response from `WorkloadTemplatesAPI.CreateDistributedTemplate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDistributedTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **distributedTemplateCreationRequest** | [**DistributedTemplateCreationRequest**](DistributedTemplateCreationRequest.md) |  | 

### Return type

[**DistributedTemplate**](DistributedTemplate.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateInferenceTemplate

> InferenceTemplate CreateInferenceTemplate(ctx).InferenceTemplateCreationRequest(inferenceTemplateCreationRequest).Execute()

Create a inference template.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dsx-ai-factory/run-ai/runai-clients/workloads-client"
)

func main() {
	inferenceTemplateCreationRequest := *openapiclient.NewInferenceTemplateCreationRequest(*openapiclient.NewTemplateCreationMetaFields("my-template-name", openapiclient.ScopeType("system"))) // InferenceTemplateCreationRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkloadTemplatesAPI.CreateInferenceTemplate(context.Background()).InferenceTemplateCreationRequest(inferenceTemplateCreationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkloadTemplatesAPI.CreateInferenceTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateInferenceTemplate`: InferenceTemplate
	fmt.Fprintf(os.Stdout, "Response from `WorkloadTemplatesAPI.CreateInferenceTemplate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateInferenceTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inferenceTemplateCreationRequest** | [**InferenceTemplateCreationRequest**](InferenceTemplateCreationRequest.md) |  | 

### Return type

[**InferenceTemplate**](InferenceTemplate.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTrainingTemplate

> TrainingTemplate CreateTrainingTemplate(ctx).TrainingTemplateCreationRequest(trainingTemplateCreationRequest).Execute()

Create a training template.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dsx-ai-factory/run-ai/runai-clients/workloads-client"
)

func main() {
	trainingTemplateCreationRequest := *openapiclient.NewTrainingTemplateCreationRequest(*openapiclient.NewTemplateCreationMetaFields("my-template-name", openapiclient.ScopeType("system"))) // TrainingTemplateCreationRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkloadTemplatesAPI.CreateTrainingTemplate(context.Background()).TrainingTemplateCreationRequest(trainingTemplateCreationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkloadTemplatesAPI.CreateTrainingTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTrainingTemplate`: TrainingTemplate
	fmt.Fprintf(os.Stdout, "Response from `WorkloadTemplatesAPI.CreateTrainingTemplate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTrainingTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **trainingTemplateCreationRequest** | [**TrainingTemplateCreationRequest**](TrainingTemplateCreationRequest.md) |  | 

### Return type

[**TrainingTemplate**](TrainingTemplate.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateWorkspaceTemplate

> WorkspaceTemplate CreateWorkspaceTemplate(ctx).WorkspaceTemplateCreationRequest(workspaceTemplateCreationRequest).Execute()

Create a workspace template.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dsx-ai-factory/run-ai/runai-clients/workloads-client"
)

func main() {
	workspaceTemplateCreationRequest := *openapiclient.NewWorkspaceTemplateCreationRequest(*openapiclient.NewTemplateCreationMetaFields("my-template-name", openapiclient.ScopeType("system"))) // WorkspaceTemplateCreationRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkloadTemplatesAPI.CreateWorkspaceTemplate(context.Background()).WorkspaceTemplateCreationRequest(workspaceTemplateCreationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkloadTemplatesAPI.CreateWorkspaceTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateWorkspaceTemplate`: WorkspaceTemplate
	fmt.Fprintf(os.Stdout, "Response from `WorkloadTemplatesAPI.CreateWorkspaceTemplate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorkspaceTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workspaceTemplateCreationRequest** | [**WorkspaceTemplateCreationRequest**](WorkspaceTemplateCreationRequest.md) |  | 

### Return type

[**WorkspaceTemplate**](WorkspaceTemplate.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDistributedTemplate

> DeleteDistributedTemplate(ctx, templateId).Execute()

Delete a distributed template.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dsx-ai-factory/run-ai/runai-clients/workloads-client"
)

func main() {
	templateId := "550e8400-e29b-41d4-a716-446655440000" // string | The  Universally Unique Identifier (UUID) of the template.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WorkloadTemplatesAPI.DeleteDistributedTemplate(context.Background(), templateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkloadTemplatesAPI.DeleteDistributedTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **string** | The  Universally Unique Identifier (UUID) of the template. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDistributedTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteInferenceTemplate

> DeleteInferenceTemplate(ctx, templateId).Execute()

Delete a inference template.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dsx-ai-factory/run-ai/runai-clients/workloads-client"
)

func main() {
	templateId := "550e8400-e29b-41d4-a716-446655440000" // string | The  Universally Unique Identifier (UUID) of the template.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WorkloadTemplatesAPI.DeleteInferenceTemplate(context.Background(), templateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkloadTemplatesAPI.DeleteInferenceTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **string** | The  Universally Unique Identifier (UUID) of the template. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteInferenceTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTrainingTemplate

> DeleteTrainingTemplate(ctx, templateId).Execute()

Delete a training template.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dsx-ai-factory/run-ai/runai-clients/workloads-client"
)

func main() {
	templateId := "550e8400-e29b-41d4-a716-446655440000" // string | The  Universally Unique Identifier (UUID) of the template.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WorkloadTemplatesAPI.DeleteTrainingTemplate(context.Background(), templateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkloadTemplatesAPI.DeleteTrainingTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **string** | The  Universally Unique Identifier (UUID) of the template. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTrainingTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteWorkspaceTemplate

> DeleteWorkspaceTemplate(ctx, templateId).Execute()

Delete a workspace template.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dsx-ai-factory/run-ai/runai-clients/workloads-client"
)

func main() {
	templateId := "550e8400-e29b-41d4-a716-446655440000" // string | The  Universally Unique Identifier (UUID) of the template.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WorkloadTemplatesAPI.DeleteWorkspaceTemplate(context.Background(), templateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkloadTemplatesAPI.DeleteWorkspaceTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **string** | The  Universally Unique Identifier (UUID) of the template. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWorkspaceTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDistributedTemplate

> DistributedTemplate GetDistributedTemplate(ctx, templateId).Execute()

Get a distributed template.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dsx-ai-factory/run-ai/runai-clients/workloads-client"
)

func main() {
	templateId := "550e8400-e29b-41d4-a716-446655440000" // string | The  Universally Unique Identifier (UUID) of the template.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkloadTemplatesAPI.GetDistributedTemplate(context.Background(), templateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkloadTemplatesAPI.GetDistributedTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDistributedTemplate`: DistributedTemplate
	fmt.Fprintf(os.Stdout, "Response from `WorkloadTemplatesAPI.GetDistributedTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **string** | The  Universally Unique Identifier (UUID) of the template. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDistributedTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DistributedTemplate**](DistributedTemplate.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInferenceTemplate

> InferenceTemplate GetInferenceTemplate(ctx, templateId).Execute()

Get a inference template.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dsx-ai-factory/run-ai/runai-clients/workloads-client"
)

func main() {
	templateId := "550e8400-e29b-41d4-a716-446655440000" // string | The  Universally Unique Identifier (UUID) of the template.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkloadTemplatesAPI.GetInferenceTemplate(context.Background(), templateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkloadTemplatesAPI.GetInferenceTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInferenceTemplate`: InferenceTemplate
	fmt.Fprintf(os.Stdout, "Response from `WorkloadTemplatesAPI.GetInferenceTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **string** | The  Universally Unique Identifier (UUID) of the template. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInferenceTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InferenceTemplate**](InferenceTemplate.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTemplateById

> SupersetTemplate GetTemplateById(ctx, templateId).Execute()

Get a template, any type.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dsx-ai-factory/run-ai/runai-clients/workloads-client"
)

func main() {
	templateId := "550e8400-e29b-41d4-a716-446655440000" // string | The  Universally Unique Identifier (UUID) of the template.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkloadTemplatesAPI.GetTemplateById(context.Background(), templateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkloadTemplatesAPI.GetTemplateById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTemplateById`: SupersetTemplate
	fmt.Fprintf(os.Stdout, "Response from `WorkloadTemplatesAPI.GetTemplateById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **string** | The  Universally Unique Identifier (UUID) of the template. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTemplateByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SupersetTemplate**](SupersetTemplate.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTrainingTemplate

> TrainingTemplate GetTrainingTemplate(ctx, templateId).Execute()

Get a training template.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dsx-ai-factory/run-ai/runai-clients/workloads-client"
)

func main() {
	templateId := "550e8400-e29b-41d4-a716-446655440000" // string | The  Universally Unique Identifier (UUID) of the template.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkloadTemplatesAPI.GetTrainingTemplate(context.Background(), templateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkloadTemplatesAPI.GetTrainingTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTrainingTemplate`: TrainingTemplate
	fmt.Fprintf(os.Stdout, "Response from `WorkloadTemplatesAPI.GetTrainingTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **string** | The  Universally Unique Identifier (UUID) of the template. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTrainingTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TrainingTemplate**](TrainingTemplate.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWorkspaceTemplate

> WorkspaceTemplate GetWorkspaceTemplate(ctx, templateId).Execute()

Get a workspace template.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dsx-ai-factory/run-ai/runai-clients/workloads-client"
)

func main() {
	templateId := "550e8400-e29b-41d4-a716-446655440000" // string | The  Universally Unique Identifier (UUID) of the template.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkloadTemplatesAPI.GetWorkspaceTemplate(context.Background(), templateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkloadTemplatesAPI.GetWorkspaceTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWorkspaceTemplate`: WorkspaceTemplate
	fmt.Fprintf(os.Stdout, "Response from `WorkloadTemplatesAPI.GetWorkspaceTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **string** | The  Universally Unique Identifier (UUID) of the template. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkspaceTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WorkspaceTemplate**](WorkspaceTemplate.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDistributedTemplates

> DistributedTemplatesListResponse ListDistributedTemplates(ctx).Offset(offset).Limit(limit).SortOrder(sortOrder).SortBy(sortBy).FilterBy(filterBy).Search(search).IncludeAncestors(includeAncestors).Execute()

List distributed templates.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dsx-ai-factory/run-ai/runai-clients/workloads-client"
)

func main() {
	offset := int32(100) // int32 | The offset of the first item returned in the collection. (optional)
	limit := int32(56) // int32 | The maximum number of entries to return. (optional) (default to 50)
	sortOrder := "sortOrder_example" // string | Sort results in descending or ascending order. (optional) (default to "asc")
	sortBy := "sortBy_example" // string | Sort results by parameter. (optional)
	filterBy := []string{"Inner_example"} // []string | Filter results by a parameter. Use the format field-name operator value. Operators are `==` Equals, `!=` Not equals, `<=` Less than or equal, `>=` Greater than or equal, `=@` contains, `!@` Does not contain, `=^` Starts with and `=$` Ends with. Dates are in ISO 8601 timestamp format and available for operators `==`, `!=`, `<=` and `>=`. (optional)
	search := "test project" // string | Filter results by a free text search. (optional)
	includeAncestors := true // bool | When filtering by projectId, set includeAncestors=true to include templates from the project's parent department and cluster in the response. Any reason we are calling this includeAncestors (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkloadTemplatesAPI.ListDistributedTemplates(context.Background()).Offset(offset).Limit(limit).SortOrder(sortOrder).SortBy(sortBy).FilterBy(filterBy).Search(search).IncludeAncestors(includeAncestors).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkloadTemplatesAPI.ListDistributedTemplates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDistributedTemplates`: DistributedTemplatesListResponse
	fmt.Fprintf(os.Stdout, "Response from `WorkloadTemplatesAPI.ListDistributedTemplates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDistributedTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **int32** | The offset of the first item returned in the collection. | 
 **limit** | **int32** | The maximum number of entries to return. | [default to 50]
 **sortOrder** | **string** | Sort results in descending or ascending order. | [default to &quot;asc&quot;]
 **sortBy** | **string** | Sort results by parameter. | 
 **filterBy** | **[]string** | Filter results by a parameter. Use the format field-name operator value. Operators are &#x60;&#x3D;&#x3D;&#x60; Equals, &#x60;!&#x3D;&#x60; Not equals, &#x60;&lt;&#x3D;&#x60; Less than or equal, &#x60;&gt;&#x3D;&#x60; Greater than or equal, &#x60;&#x3D;@&#x60; contains, &#x60;!@&#x60; Does not contain, &#x60;&#x3D;^&#x60; Starts with and &#x60;&#x3D;$&#x60; Ends with. Dates are in ISO 8601 timestamp format and available for operators &#x60;&#x3D;&#x3D;&#x60;, &#x60;!&#x3D;&#x60;, &#x60;&lt;&#x3D;&#x60; and &#x60;&gt;&#x3D;&#x60;. | 
 **search** | **string** | Filter results by a free text search. | 
 **includeAncestors** | **bool** | When filtering by projectId, set includeAncestors&#x3D;true to include templates from the project&#39;s parent department and cluster in the response. Any reason we are calling this includeAncestors | 

### Return type

[**DistributedTemplatesListResponse**](DistributedTemplatesListResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListInferenceTemplates

> InferenceTemplatesListResponse ListInferenceTemplates(ctx).Offset(offset).Limit(limit).SortOrder(sortOrder).SortBy(sortBy).FilterBy(filterBy).Search(search).IncludeAncestors(includeAncestors).Execute()

List inference templates.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dsx-ai-factory/run-ai/runai-clients/workloads-client"
)

func main() {
	offset := int32(100) // int32 | The offset of the first item returned in the collection. (optional)
	limit := int32(56) // int32 | The maximum number of entries to return. (optional) (default to 50)
	sortOrder := "sortOrder_example" // string | Sort results in descending or ascending order. (optional) (default to "asc")
	sortBy := "sortBy_example" // string | Sort results by parameter. (optional)
	filterBy := []string{"Inner_example"} // []string | Filter results by a parameter. Use the format field-name operator value. Operators are `==` Equals, `!=` Not equals, `<=` Less than or equal, `>=` Greater than or equal, `=@` contains, `!@` Does not contain, `=^` Starts with and `=$` Ends with. Dates are in ISO 8601 timestamp format and available for operators `==`, `!=`, `<=` and `>=`. (optional)
	search := "test project" // string | Filter results by a free text search. (optional)
	includeAncestors := true // bool | When filtering by projectId, set includeAncestors=true to include templates from the project's parent department and cluster in the response. Any reason we are calling this includeAncestors (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkloadTemplatesAPI.ListInferenceTemplates(context.Background()).Offset(offset).Limit(limit).SortOrder(sortOrder).SortBy(sortBy).FilterBy(filterBy).Search(search).IncludeAncestors(includeAncestors).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkloadTemplatesAPI.ListInferenceTemplates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListInferenceTemplates`: InferenceTemplatesListResponse
	fmt.Fprintf(os.Stdout, "Response from `WorkloadTemplatesAPI.ListInferenceTemplates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListInferenceTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **int32** | The offset of the first item returned in the collection. | 
 **limit** | **int32** | The maximum number of entries to return. | [default to 50]
 **sortOrder** | **string** | Sort results in descending or ascending order. | [default to &quot;asc&quot;]
 **sortBy** | **string** | Sort results by parameter. | 
 **filterBy** | **[]string** | Filter results by a parameter. Use the format field-name operator value. Operators are &#x60;&#x3D;&#x3D;&#x60; Equals, &#x60;!&#x3D;&#x60; Not equals, &#x60;&lt;&#x3D;&#x60; Less than or equal, &#x60;&gt;&#x3D;&#x60; Greater than or equal, &#x60;&#x3D;@&#x60; contains, &#x60;!@&#x60; Does not contain, &#x60;&#x3D;^&#x60; Starts with and &#x60;&#x3D;$&#x60; Ends with. Dates are in ISO 8601 timestamp format and available for operators &#x60;&#x3D;&#x3D;&#x60;, &#x60;!&#x3D;&#x60;, &#x60;&lt;&#x3D;&#x60; and &#x60;&gt;&#x3D;&#x60;. | 
 **search** | **string** | Filter results by a free text search. | 
 **includeAncestors** | **bool** | When filtering by projectId, set includeAncestors&#x3D;true to include templates from the project&#39;s parent department and cluster in the response. Any reason we are calling this includeAncestors | 

### Return type

[**InferenceTemplatesListResponse**](InferenceTemplatesListResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTemplates

> TemplatesListResponse ListTemplates(ctx).Offset(offset).Limit(limit).SortOrder(sortOrder).SortBy(sortBy).FilterBy(filterBy).Search(search).IncludeAncestors(includeAncestors).Execute()

List templates.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dsx-ai-factory/run-ai/runai-clients/workloads-client"
)

func main() {
	offset := int32(100) // int32 | The offset of the first item returned in the collection. (optional)
	limit := int32(56) // int32 | The maximum number of entries to return. (optional) (default to 50)
	sortOrder := "sortOrder_example" // string | Sort results in descending or ascending order. (optional) (default to "asc")
	sortBy := "sortBy_example" // string | Sort results by parameter. (optional)
	filterBy := []string{"Inner_example"} // []string | Filter results by a parameter. Use the format field-name operator value. Operators are `==` Equals, `!=` Not equals, `<=` Less than or equal, `>=` Greater than or equal, `=@` contains, `!@` Does not contain, `=^` Starts with and `=$` Ends with. Dates are in ISO 8601 timestamp format and available for operators `==`, `!=`, `<=` and `>=`. (optional)
	search := "test project" // string | Filter results by a free text search. (optional)
	includeAncestors := true // bool | When filtering by projectId, set includeAncestors=true to include templates from the project's parent department and cluster in the response. Any reason we are calling this includeAncestors (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkloadTemplatesAPI.ListTemplates(context.Background()).Offset(offset).Limit(limit).SortOrder(sortOrder).SortBy(sortBy).FilterBy(filterBy).Search(search).IncludeAncestors(includeAncestors).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkloadTemplatesAPI.ListTemplates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListTemplates`: TemplatesListResponse
	fmt.Fprintf(os.Stdout, "Response from `WorkloadTemplatesAPI.ListTemplates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **int32** | The offset of the first item returned in the collection. | 
 **limit** | **int32** | The maximum number of entries to return. | [default to 50]
 **sortOrder** | **string** | Sort results in descending or ascending order. | [default to &quot;asc&quot;]
 **sortBy** | **string** | Sort results by parameter. | 
 **filterBy** | **[]string** | Filter results by a parameter. Use the format field-name operator value. Operators are &#x60;&#x3D;&#x3D;&#x60; Equals, &#x60;!&#x3D;&#x60; Not equals, &#x60;&lt;&#x3D;&#x60; Less than or equal, &#x60;&gt;&#x3D;&#x60; Greater than or equal, &#x60;&#x3D;@&#x60; contains, &#x60;!@&#x60; Does not contain, &#x60;&#x3D;^&#x60; Starts with and &#x60;&#x3D;$&#x60; Ends with. Dates are in ISO 8601 timestamp format and available for operators &#x60;&#x3D;&#x3D;&#x60;, &#x60;!&#x3D;&#x60;, &#x60;&lt;&#x3D;&#x60; and &#x60;&gt;&#x3D;&#x60;. | 
 **search** | **string** | Filter results by a free text search. | 
 **includeAncestors** | **bool** | When filtering by projectId, set includeAncestors&#x3D;true to include templates from the project&#39;s parent department and cluster in the response. Any reason we are calling this includeAncestors | 

### Return type

[**TemplatesListResponse**](TemplatesListResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTrainingTemplates

> TrainingTemplatesListResponse ListTrainingTemplates(ctx).Offset(offset).Limit(limit).SortOrder(sortOrder).SortBy(sortBy).FilterBy(filterBy).Search(search).IncludeAncestors(includeAncestors).Execute()

List training templates.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dsx-ai-factory/run-ai/runai-clients/workloads-client"
)

func main() {
	offset := int32(100) // int32 | The offset of the first item returned in the collection. (optional)
	limit := int32(56) // int32 | The maximum number of entries to return. (optional) (default to 50)
	sortOrder := "sortOrder_example" // string | Sort results in descending or ascending order. (optional) (default to "asc")
	sortBy := "sortBy_example" // string | Sort results by parameter. (optional)
	filterBy := []string{"Inner_example"} // []string | Filter results by a parameter. Use the format field-name operator value. Operators are `==` Equals, `!=` Not equals, `<=` Less than or equal, `>=` Greater than or equal, `=@` contains, `!@` Does not contain, `=^` Starts with and `=$` Ends with. Dates are in ISO 8601 timestamp format and available for operators `==`, `!=`, `<=` and `>=`. (optional)
	search := "test project" // string | Filter results by a free text search. (optional)
	includeAncestors := true // bool | When filtering by projectId, set includeAncestors=true to include templates from the project's parent department and cluster in the response. Any reason we are calling this includeAncestors (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkloadTemplatesAPI.ListTrainingTemplates(context.Background()).Offset(offset).Limit(limit).SortOrder(sortOrder).SortBy(sortBy).FilterBy(filterBy).Search(search).IncludeAncestors(includeAncestors).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkloadTemplatesAPI.ListTrainingTemplates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListTrainingTemplates`: TrainingTemplatesListResponse
	fmt.Fprintf(os.Stdout, "Response from `WorkloadTemplatesAPI.ListTrainingTemplates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListTrainingTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **int32** | The offset of the first item returned in the collection. | 
 **limit** | **int32** | The maximum number of entries to return. | [default to 50]
 **sortOrder** | **string** | Sort results in descending or ascending order. | [default to &quot;asc&quot;]
 **sortBy** | **string** | Sort results by parameter. | 
 **filterBy** | **[]string** | Filter results by a parameter. Use the format field-name operator value. Operators are &#x60;&#x3D;&#x3D;&#x60; Equals, &#x60;!&#x3D;&#x60; Not equals, &#x60;&lt;&#x3D;&#x60; Less than or equal, &#x60;&gt;&#x3D;&#x60; Greater than or equal, &#x60;&#x3D;@&#x60; contains, &#x60;!@&#x60; Does not contain, &#x60;&#x3D;^&#x60; Starts with and &#x60;&#x3D;$&#x60; Ends with. Dates are in ISO 8601 timestamp format and available for operators &#x60;&#x3D;&#x3D;&#x60;, &#x60;!&#x3D;&#x60;, &#x60;&lt;&#x3D;&#x60; and &#x60;&gt;&#x3D;&#x60;. | 
 **search** | **string** | Filter results by a free text search. | 
 **includeAncestors** | **bool** | When filtering by projectId, set includeAncestors&#x3D;true to include templates from the project&#39;s parent department and cluster in the response. Any reason we are calling this includeAncestors | 

### Return type

[**TrainingTemplatesListResponse**](TrainingTemplatesListResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWorkspaceTemplates

> WorkspaceTemplatesListResponse ListWorkspaceTemplates(ctx).Offset(offset).Limit(limit).SortOrder(sortOrder).SortBy(sortBy).FilterBy(filterBy).Search(search).IncludeAncestors(includeAncestors).Execute()

List workspace templates.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dsx-ai-factory/run-ai/runai-clients/workloads-client"
)

func main() {
	offset := int32(100) // int32 | The offset of the first item returned in the collection. (optional)
	limit := int32(56) // int32 | The maximum number of entries to return. (optional) (default to 50)
	sortOrder := "sortOrder_example" // string | Sort results in descending or ascending order. (optional) (default to "asc")
	sortBy := "sortBy_example" // string | Sort results by parameter. (optional)
	filterBy := []string{"Inner_example"} // []string | Filter results by a parameter. Use the format field-name operator value. Operators are `==` Equals, `!=` Not equals, `<=` Less than or equal, `>=` Greater than or equal, `=@` contains, `!@` Does not contain, `=^` Starts with and `=$` Ends with. Dates are in ISO 8601 timestamp format and available for operators `==`, `!=`, `<=` and `>=`. (optional)
	search := "test project" // string | Filter results by a free text search. (optional)
	includeAncestors := true // bool | When filtering by projectId, set includeAncestors=true to include templates from the project's parent department and cluster in the response. Any reason we are calling this includeAncestors (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkloadTemplatesAPI.ListWorkspaceTemplates(context.Background()).Offset(offset).Limit(limit).SortOrder(sortOrder).SortBy(sortBy).FilterBy(filterBy).Search(search).IncludeAncestors(includeAncestors).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkloadTemplatesAPI.ListWorkspaceTemplates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListWorkspaceTemplates`: WorkspaceTemplatesListResponse
	fmt.Fprintf(os.Stdout, "Response from `WorkloadTemplatesAPI.ListWorkspaceTemplates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListWorkspaceTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **int32** | The offset of the first item returned in the collection. | 
 **limit** | **int32** | The maximum number of entries to return. | [default to 50]
 **sortOrder** | **string** | Sort results in descending or ascending order. | [default to &quot;asc&quot;]
 **sortBy** | **string** | Sort results by parameter. | 
 **filterBy** | **[]string** | Filter results by a parameter. Use the format field-name operator value. Operators are &#x60;&#x3D;&#x3D;&#x60; Equals, &#x60;!&#x3D;&#x60; Not equals, &#x60;&lt;&#x3D;&#x60; Less than or equal, &#x60;&gt;&#x3D;&#x60; Greater than or equal, &#x60;&#x3D;@&#x60; contains, &#x60;!@&#x60; Does not contain, &#x60;&#x3D;^&#x60; Starts with and &#x60;&#x3D;$&#x60; Ends with. Dates are in ISO 8601 timestamp format and available for operators &#x60;&#x3D;&#x3D;&#x60;, &#x60;!&#x3D;&#x60;, &#x60;&lt;&#x3D;&#x60; and &#x60;&gt;&#x3D;&#x60;. | 
 **search** | **string** | Filter results by a free text search. | 
 **includeAncestors** | **bool** | When filtering by projectId, set includeAncestors&#x3D;true to include templates from the project&#39;s parent department and cluster in the response. Any reason we are calling this includeAncestors | 

### Return type

[**WorkspaceTemplatesListResponse**](WorkspaceTemplatesListResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchDistributedTemplate

> DistributedTemplate PatchDistributedTemplate(ctx, templateId).DistributedTemplatePatchRequest(distributedTemplatePatchRequest).Execute()

Patch a distributed template.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dsx-ai-factory/run-ai/runai-clients/workloads-client"
)

func main() {
	templateId := "550e8400-e29b-41d4-a716-446655440000" // string | The  Universally Unique Identifier (UUID) of the template.
	distributedTemplatePatchRequest := *openapiclient.NewDistributedTemplatePatchRequest() // DistributedTemplatePatchRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkloadTemplatesAPI.PatchDistributedTemplate(context.Background(), templateId).DistributedTemplatePatchRequest(distributedTemplatePatchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkloadTemplatesAPI.PatchDistributedTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchDistributedTemplate`: DistributedTemplate
	fmt.Fprintf(os.Stdout, "Response from `WorkloadTemplatesAPI.PatchDistributedTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **string** | The  Universally Unique Identifier (UUID) of the template. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchDistributedTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **distributedTemplatePatchRequest** | [**DistributedTemplatePatchRequest**](DistributedTemplatePatchRequest.md) |  | 

### Return type

[**DistributedTemplate**](DistributedTemplate.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchInferenceTemplate

> InferenceTemplate PatchInferenceTemplate(ctx, templateId).InferenceTemplatePatchRequest(inferenceTemplatePatchRequest).Execute()

Patch a inference template.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dsx-ai-factory/run-ai/runai-clients/workloads-client"
)

func main() {
	templateId := "550e8400-e29b-41d4-a716-446655440000" // string | The  Universally Unique Identifier (UUID) of the template.
	inferenceTemplatePatchRequest := *openapiclient.NewInferenceTemplatePatchRequest() // InferenceTemplatePatchRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkloadTemplatesAPI.PatchInferenceTemplate(context.Background(), templateId).InferenceTemplatePatchRequest(inferenceTemplatePatchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkloadTemplatesAPI.PatchInferenceTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchInferenceTemplate`: InferenceTemplate
	fmt.Fprintf(os.Stdout, "Response from `WorkloadTemplatesAPI.PatchInferenceTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **string** | The  Universally Unique Identifier (UUID) of the template. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchInferenceTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inferenceTemplatePatchRequest** | [**InferenceTemplatePatchRequest**](InferenceTemplatePatchRequest.md) |  | 

### Return type

[**InferenceTemplate**](InferenceTemplate.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchTrainingTemplate

> TrainingTemplate PatchTrainingTemplate(ctx, templateId).TrainingTemplatePatchRequest(trainingTemplatePatchRequest).Execute()

Patch a training template.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dsx-ai-factory/run-ai/runai-clients/workloads-client"
)

func main() {
	templateId := "550e8400-e29b-41d4-a716-446655440000" // string | The  Universally Unique Identifier (UUID) of the template.
	trainingTemplatePatchRequest := *openapiclient.NewTrainingTemplatePatchRequest() // TrainingTemplatePatchRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkloadTemplatesAPI.PatchTrainingTemplate(context.Background(), templateId).TrainingTemplatePatchRequest(trainingTemplatePatchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkloadTemplatesAPI.PatchTrainingTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchTrainingTemplate`: TrainingTemplate
	fmt.Fprintf(os.Stdout, "Response from `WorkloadTemplatesAPI.PatchTrainingTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **string** | The  Universally Unique Identifier (UUID) of the template. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchTrainingTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **trainingTemplatePatchRequest** | [**TrainingTemplatePatchRequest**](TrainingTemplatePatchRequest.md) |  | 

### Return type

[**TrainingTemplate**](TrainingTemplate.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchWorkspaceTemplate

> WorkspaceTemplate PatchWorkspaceTemplate(ctx, templateId).WorkspaceTemplatePatchRequest(workspaceTemplatePatchRequest).Execute()

Patch a workspace template.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dsx-ai-factory/run-ai/runai-clients/workloads-client"
)

func main() {
	templateId := "550e8400-e29b-41d4-a716-446655440000" // string | The  Universally Unique Identifier (UUID) of the template.
	workspaceTemplatePatchRequest := *openapiclient.NewWorkspaceTemplatePatchRequest() // WorkspaceTemplatePatchRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkloadTemplatesAPI.PatchWorkspaceTemplate(context.Background(), templateId).WorkspaceTemplatePatchRequest(workspaceTemplatePatchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkloadTemplatesAPI.PatchWorkspaceTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchWorkspaceTemplate`: WorkspaceTemplate
	fmt.Fprintf(os.Stdout, "Response from `WorkloadTemplatesAPI.PatchWorkspaceTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **string** | The  Universally Unique Identifier (UUID) of the template. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchWorkspaceTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **workspaceTemplatePatchRequest** | [**WorkspaceTemplatePatchRequest**](WorkspaceTemplatePatchRequest.md) |  | 

### Return type

[**WorkspaceTemplate**](WorkspaceTemplate.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDistributedTemplate

> DistributedTemplate UpdateDistributedTemplate(ctx, templateId).DistributedTemplateUpdateRequest(distributedTemplateUpdateRequest).Execute()

Update a distributed template.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dsx-ai-factory/run-ai/runai-clients/workloads-client"
)

func main() {
	templateId := "550e8400-e29b-41d4-a716-446655440000" // string | The  Universally Unique Identifier (UUID) of the template.
	distributedTemplateUpdateRequest := *openapiclient.NewDistributedTemplateUpdateRequest(*openapiclient.NewTemplateUpdateMetaFields("my-template-name")) // DistributedTemplateUpdateRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkloadTemplatesAPI.UpdateDistributedTemplate(context.Background(), templateId).DistributedTemplateUpdateRequest(distributedTemplateUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkloadTemplatesAPI.UpdateDistributedTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDistributedTemplate`: DistributedTemplate
	fmt.Fprintf(os.Stdout, "Response from `WorkloadTemplatesAPI.UpdateDistributedTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **string** | The  Universally Unique Identifier (UUID) of the template. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDistributedTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **distributedTemplateUpdateRequest** | [**DistributedTemplateUpdateRequest**](DistributedTemplateUpdateRequest.md) |  | 

### Return type

[**DistributedTemplate**](DistributedTemplate.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateInferenceTemplate

> InferenceTemplate UpdateInferenceTemplate(ctx, templateId).InferenceTemplateUpdateRequest(inferenceTemplateUpdateRequest).Execute()

Update a inference template.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dsx-ai-factory/run-ai/runai-clients/workloads-client"
)

func main() {
	templateId := "550e8400-e29b-41d4-a716-446655440000" // string | The  Universally Unique Identifier (UUID) of the template.
	inferenceTemplateUpdateRequest := *openapiclient.NewInferenceTemplateUpdateRequest(*openapiclient.NewTemplateUpdateMetaFields("my-template-name")) // InferenceTemplateUpdateRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkloadTemplatesAPI.UpdateInferenceTemplate(context.Background(), templateId).InferenceTemplateUpdateRequest(inferenceTemplateUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkloadTemplatesAPI.UpdateInferenceTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateInferenceTemplate`: InferenceTemplate
	fmt.Fprintf(os.Stdout, "Response from `WorkloadTemplatesAPI.UpdateInferenceTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **string** | The  Universally Unique Identifier (UUID) of the template. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateInferenceTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inferenceTemplateUpdateRequest** | [**InferenceTemplateUpdateRequest**](InferenceTemplateUpdateRequest.md) |  | 

### Return type

[**InferenceTemplate**](InferenceTemplate.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTrainingTemplate

> TrainingTemplate UpdateTrainingTemplate(ctx, templateId).TrainingTemplateUpdateRequest(trainingTemplateUpdateRequest).Execute()

Update a training template.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dsx-ai-factory/run-ai/runai-clients/workloads-client"
)

func main() {
	templateId := "550e8400-e29b-41d4-a716-446655440000" // string | The  Universally Unique Identifier (UUID) of the template.
	trainingTemplateUpdateRequest := *openapiclient.NewTrainingTemplateUpdateRequest(*openapiclient.NewTemplateUpdateMetaFields("my-template-name")) // TrainingTemplateUpdateRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkloadTemplatesAPI.UpdateTrainingTemplate(context.Background(), templateId).TrainingTemplateUpdateRequest(trainingTemplateUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkloadTemplatesAPI.UpdateTrainingTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateTrainingTemplate`: TrainingTemplate
	fmt.Fprintf(os.Stdout, "Response from `WorkloadTemplatesAPI.UpdateTrainingTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **string** | The  Universally Unique Identifier (UUID) of the template. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTrainingTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **trainingTemplateUpdateRequest** | [**TrainingTemplateUpdateRequest**](TrainingTemplateUpdateRequest.md) |  | 

### Return type

[**TrainingTemplate**](TrainingTemplate.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWorkspaceTemplate

> WorkspaceTemplate UpdateWorkspaceTemplate(ctx, templateId).WorkspaceTemplateUpdateRequest(workspaceTemplateUpdateRequest).Execute()

Update a workspace template.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dsx-ai-factory/run-ai/runai-clients/workloads-client"
)

func main() {
	templateId := "550e8400-e29b-41d4-a716-446655440000" // string | The  Universally Unique Identifier (UUID) of the template.
	workspaceTemplateUpdateRequest := *openapiclient.NewWorkspaceTemplateUpdateRequest(*openapiclient.NewTemplateUpdateMetaFields("my-template-name")) // WorkspaceTemplateUpdateRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkloadTemplatesAPI.UpdateWorkspaceTemplate(context.Background(), templateId).WorkspaceTemplateUpdateRequest(workspaceTemplateUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkloadTemplatesAPI.UpdateWorkspaceTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateWorkspaceTemplate`: WorkspaceTemplate
	fmt.Fprintf(os.Stdout, "Response from `WorkloadTemplatesAPI.UpdateWorkspaceTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **string** | The  Universally Unique Identifier (UUID) of the template. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWorkspaceTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **workspaceTemplateUpdateRequest** | [**WorkspaceTemplateUpdateRequest**](WorkspaceTemplateUpdateRequest.md) |  | 

### Return type

[**WorkspaceTemplate**](WorkspaceTemplate.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

