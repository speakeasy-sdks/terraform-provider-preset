// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package sdk

import (
	"Preset/internal/sdk/pkg/models/operations"
	"Preset/internal/sdk/pkg/utils"
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// supersetAPIsOpenSourceGreaterThanDatasets - APIs to manage your datasets.
type supersetAPIsOpenSourceGreaterThanDatasets struct {
	sdkConfiguration sdkConfiguration
}

func newSupersetAPIsOpenSourceGreaterThanDatasets(sdkConfig sdkConfiguration) *supersetAPIsOpenSourceGreaterThanDatasets {
	return &supersetAPIsOpenSourceGreaterThanDatasets{
		sdkConfiguration: sdkConfig,
	}
}

// DeleteAPIV1DatasetDatasetID - Delete a Dataset
// Deletes a Dataset.
//
// Replace in the URL:
//
// - `{{WorkspaceSlug}}` with the `name` retrieved through the API (using the **Get Workspaces from a Team** endpoint).
// - `{{WorkspaceRegion}}` corresponding to the `region` retrieved from the **Get Workspaces from a Team** endpoint. Refer to below table:
//
// | **`region`** | **`WorkspaceRegion`** |
// | --- | --- |
// | us-east-1 | `us2a` |
// | us-west-2 | `us1a` |
// | eu-north-1 | `eu5a` |
// | ap-northeast-1 | `ap1a` |
//
// Alternatively, access the Workspace through the UI, and get the `{{WorkspaceSlug}}` and `{{WorkspaceRegion}}` from the URL -> `https://{{WorkspaceSlug}}.{{Region}}.preset.io/superset/welcome/`.
//
// - `{{DatasetID}}` with the Dataset `id` retrieved from the **Get All Datasets from a Workspace** endpoint.
func (s *supersetAPIsOpenSourceGreaterThanDatasets) DeleteAPIV1DatasetDatasetID(ctx context.Context, request operations.DeleteAPIV1DatasetDatasetIDRequest) (*operations.DeleteAPIV1DatasetDatasetIDResponse, error) {
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url, err := utils.GenerateURL(ctx, baseURL, "/api/v1/dataset/{DatasetID}", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "DELETE", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "*/*")
	req.Header.Set("user-agent", fmt.Sprintf("speakeasy-sdk/%s %s %s %s", s.sdkConfiguration.Language, s.sdkConfiguration.SDKVersion, s.sdkConfiguration.GenVersion, s.sdkConfiguration.OpenAPIDocVersion))

	utils.PopulateHeaders(ctx, req, request)

	client := s.sdkConfiguration.SecurityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.DeleteAPIV1DatasetDatasetIDResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
	}

	return res, nil
}

// GetAPIV1Dataset - Get all Datasets from a Workspace
// Gets all Datasets available on the Worpkspace.
//
// Replace in the URL:
//
// - `{{WorkspaceSlug}}` with the `name` retrieved through the API (using the **Get Workspaces from a Team** endpoint).
// - `{{WorkspaceRegion}}` corresponding to the `region` retrieved from the **Get Workspaces from a Team** endpoint. Refer to below table:
//
// | **`region`** | **`WorkspaceRegion`** |
// | --- | --- |
// | us-east-1 | `us2a` |
// | us-west-2 | `us1a` |
// | eu-north-1 | `eu5a` |
// | ap-northeast-1 | `ap1a` |
//
// Alternatively, access the Workspace through the UI, and get the `{{WorkspaceSlug}}` and `{{WorkspaceRegion}}` from the URL -> `https://{{WorkspaceSlug}}.{{Region}}.preset.io/superset/welcome/`.
//
// Note that this endpoint returns 20 results by default. You can return up to 100 results at a time and use pagination by adding the following query parameters:
//
// ```
// ?q=(page_size:{{PageSize}},page:{{Page}})
//
// ```
//
// Replace:
//
// - `{{PageSize}}` with the desired size (min `1` max `100`).
// - `{{Page}}` with the page number (useful when the total count > `{{PageSize}}` - min `0`).
func (s *supersetAPIsOpenSourceGreaterThanDatasets) GetAPIV1Dataset(ctx context.Context, request operations.GetAPIV1DatasetRequest) (*operations.GetAPIV1DatasetResponse, error) {
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url := strings.TrimSuffix(baseURL, "/") + "/api/v1/dataset/"

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "*/*")
	req.Header.Set("user-agent", fmt.Sprintf("speakeasy-sdk/%s %s %s %s", s.sdkConfiguration.Language, s.sdkConfiguration.SDKVersion, s.sdkConfiguration.GenVersion, s.sdkConfiguration.OpenAPIDocVersion))

	utils.PopulateHeaders(ctx, req, request)

	client := s.sdkConfiguration.SecurityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.GetAPIV1DatasetResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
	}

	return res, nil
}

// GetAPIV1DatasetExport - Export Datasets
// Exports Datasets from the Workspace.
//
// Replace in the URL:
//
// - `{{WorkspaceSlug}}` with the `name` retrieved through the API (using the **Get Workspaces from a Team** endpoint).
// - `{{WorkspaceRegion}}` corresponding to the `region` retrieved from the **Get Workspaces from a Team** endpoint. Refer to below table:
//
// | **`region`** | **`WorkspaceRegion`** |
// | --- | --- |
// | us-east-1 | `us2a` |
// | us-west-2 | `us1a` |
// | eu-north-1 | `eu5a` |
// | ap-northeast-1 | `ap1a` |
//
// Alternatively, access the Workspace through the UI, and get the `{{WorkspaceSlug}}` and `{{WorkspaceRegion}}` from the URL -> `https://{{WorkspaceSlug}}.{{Region}}.preset.io/superset/welcome/`.
//
// - `{{DatasetIDs}}` with comma separated `ids` retrieved from the **Get all Datasets from a Workspace** endpoint.
//
// _**Tip:**_ If used in Postman, select `Save Response` and `Save to a File` to get the zip export.
func (s *supersetAPIsOpenSourceGreaterThanDatasets) GetAPIV1DatasetExport(ctx context.Context, request operations.GetAPIV1DatasetExportRequest) (*operations.GetAPIV1DatasetExportResponse, error) {
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url := strings.TrimSuffix(baseURL, "/") + "/api/v1/dataset/export"

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "*/*")
	req.Header.Set("user-agent", fmt.Sprintf("speakeasy-sdk/%s %s %s %s", s.sdkConfiguration.Language, s.sdkConfiguration.SDKVersion, s.sdkConfiguration.GenVersion, s.sdkConfiguration.OpenAPIDocVersion))

	utils.PopulateHeaders(ctx, req, request)

	if err := utils.PopulateQueryParams(ctx, req, request, nil); err != nil {
		return nil, fmt.Errorf("error populating query params: %w", err)
	}

	client := s.sdkConfiguration.SecurityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.GetAPIV1DatasetExportResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
	}

	return res, nil
}

// GetAPIV1DatasetRelatedOwners - Get all possible Dataset Owners
// Gets all possible Dataset owners on the Workspace.
//
// Replace in the URL:
//
// - `{{WorkspaceSlug}}` with the `name` retrieved through the API (using the **Get Workspaces from a Team** endpoint).
// - `{{WorkspaceRegion}}` corresponding to the `region` retrieved from the **Get Workspaces from a Team** endpoint. Refer to below table:
//
// | **`region`** | **`WorkspaceRegion`** |
// | --- | --- |
// | us-east-1 | `us2a` |
// | us-west-2 | `us1a` |
// | eu-north-1 | `eu5a` |
// | ap-northeast-1 | `ap1a` |
//
// Alternatively, access the Workspace through the UI, and get the `{{WorkspaceSlug}}` and `{{WorkspaceRegion}}` from the URL -> `https://{{WorkspaceSlug}}.{{Region}}.preset.io/superset/welcome/`.
//
// Note that this endpoint returns 20 results by default. You can return up to 100 results at a time and use pagination by adding the following query parameters:
//
// ```
// ?q=(page_size:{{PageSize}},page:{{Page}})
//
// ```
//
// Replace:
//
// - `{{PageSize}}` with the desired size (min `1` max `100`).
// - `{{Page}}` with the page number (useful when the total count > `{{PageSize}}` - min `0`).
func (s *supersetAPIsOpenSourceGreaterThanDatasets) GetAPIV1DatasetRelatedOwners(ctx context.Context, request operations.GetAPIV1DatasetRelatedOwnersRequest) (*operations.GetAPIV1DatasetRelatedOwnersResponse, error) {
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url := strings.TrimSuffix(baseURL, "/") + "/api/v1/dataset/related/owners"

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "*/*")
	req.Header.Set("user-agent", fmt.Sprintf("speakeasy-sdk/%s %s %s %s", s.sdkConfiguration.Language, s.sdkConfiguration.SDKVersion, s.sdkConfiguration.GenVersion, s.sdkConfiguration.OpenAPIDocVersion))

	utils.PopulateHeaders(ctx, req, request)

	client := s.sdkConfiguration.SecurityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.GetAPIV1DatasetRelatedOwnersResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
	}

	return res, nil
}

// GetAPIV1DatasetDatasetID - Get a Dataset
// Get a specific Dataset from a Workspace.
//
// Replace in the URL:
//
// - `{{WorkspaceSlug}}` with the `name` retrieved through the API (using the **Get Workspaces from a Team** endpoint).
// - `{{WorkspaceRegion}}` corresponding to the `region` retrieved from the **Get Workspaces from a Team** endpoint. Refer to below table:
//
// | **`region`** | **`WorkspaceRegion`** |
// | --- | --- |
// | us-east-1 | `us2a` |
// | us-west-2 | `us1a` |
// | eu-north-1 | `eu5a` |
// | ap-northeast-1 | `ap1a` |
//
// Alternatively, access the Workspace through the UI, and get the `{{WorkspaceSlug}}` and `{{WorkspaceRegion}}` from the URL -> `https://{{WorkspaceSlug}}.{{Region}}.preset.io/superset/welcome/`.
//
// - `{{DatasetID}}` with the Dataset `id` retrieved from the **Get All Datasets from a Workspace** endpoint.
func (s *supersetAPIsOpenSourceGreaterThanDatasets) GetAPIV1DatasetDatasetID(ctx context.Context, request operations.GetAPIV1DatasetDatasetIDRequest) (*operations.GetAPIV1DatasetDatasetIDResponse, error) {
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url, err := utils.GenerateURL(ctx, baseURL, "/api/v1/dataset/{DatasetID}", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "*/*")
	req.Header.Set("user-agent", fmt.Sprintf("speakeasy-sdk/%s %s %s %s", s.sdkConfiguration.Language, s.sdkConfiguration.SDKVersion, s.sdkConfiguration.GenVersion, s.sdkConfiguration.OpenAPIDocVersion))

	utils.PopulateHeaders(ctx, req, request)

	client := s.sdkConfiguration.SecurityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.GetAPIV1DatasetDatasetIDResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
	}

	return res, nil
}

// PostAPIV1Dataset - Create a Virtual Dataset
// Creates a new Virtual Dataset.
//
// Replace in the URL:
//
// - `{{WorkspaceSlug}}` with the `name` retrieved through the API (using the **Get Workspaces from a Team** endpoint).
// - `{{WorkspaceRegion}}` corresponding to the `region` retrieved from the **Get Workspaces from a Team** endpoint. Refer to below table:
//
// | **`region`** | **`WorkspaceRegion`** |
// | --- | --- |
// | us-east-1 | `us2a` |
// | us-west-2 | `us1a` |
// | eu-north-1 | `eu5a` |
// | ap-northeast-1 | `ap1a` |
//
// Alternatively, access the Workspace through the UI, and get the `{{WorkspaceSlug}}` and `{{WorkspaceRegion}}` from the URL -> `https://{{WorkspaceSlug}}.{{Region}}.preset.io/superset/welcome/`.
//
// Replace in the body:
//
// - `{{DBID}}` with the database `id` retrieved using the **Get Databases** endoint.
// - `{{SchemaName}}` with the desired schema.
// - `{{TableName}}` with the desired table.
// - `{{DatasetSQL}}` with the SQL query to power your dataset.
// - `{{OwnerID}}` with the `id` for the desired owner account(s) retrieved from the **Get all possible Dataset Owners** endpoint. Use a comma to separate multiple IDs (for example, `[2,5]`).
func (s *supersetAPIsOpenSourceGreaterThanDatasets) PostAPIV1Dataset(ctx context.Context, request operations.PostAPIV1DatasetRequest) (*operations.PostAPIV1DatasetResponse, error) {
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url := strings.TrimSuffix(baseURL, "/") + "/api/v1/dataset"

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request, "RequestBody", "json")
	if err != nil {
		return nil, fmt.Errorf("error serializing request body: %w", err)
	}

	debugBody := bytes.NewBuffer([]byte{})
	debugReader := io.TeeReader(bodyReader, debugBody)

	req, err := http.NewRequestWithContext(ctx, "POST", url, debugReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "*/*")
	req.Header.Set("user-agent", fmt.Sprintf("speakeasy-sdk/%s %s %s %s", s.sdkConfiguration.Language, s.sdkConfiguration.SDKVersion, s.sdkConfiguration.GenVersion, s.sdkConfiguration.OpenAPIDocVersion))

	req.Header.Set("Content-Type", reqContentType)

	utils.PopulateHeaders(ctx, req, request)

	client := s.sdkConfiguration.SecurityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Request.Body = io.NopCloser(debugBody)
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.PostAPIV1DatasetResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
	}

	return res, nil
}

// PostAPIV1DatasetImport - Import a Dataset
// Imports a Dataset.
//
// Replace in the URL and on the `Referer` header:
//
// - `{{WorkspaceSlug}}` with the `name` retrieved through the API (using the **Get Workspaces from a Team** endpoint).
// - `{{WorkspaceRegion}}` corresponding to the `region` retrieved from the **Get Workspaces from a Team** endpoint. Refer to below table:
//
// | **`region`** | **`WorkspaceRegion`** |
// | --- | --- |
// | us-east-1 | `us2a` |
// | us-west-2 | `us1a` |
// | eu-north-1 | `eu5a` |
// | ap-northeast-1 | `ap1a` |
//
// Alternatively, access the Workspace through the UI, and get the `{{WorkspaceSlug}}` and `{{WorkspaceRegion}}` from the URL -> `https://{{WorkspaceSlug}}.{{Region}}.preset.io/superset/welcome/`.
//
// In the body:
//
// - Select your Dataset ZIP file as a value for the `formData`.
// - For the `passwords` field:
//   - Replace `{{DatabaseYAMLFile}}` by the database YAML file name. You can find it in your export file, under the `databases` folder.
//   - Replace `{{DatabasePassword}}` by the DB password.
//
// - For the `overwrite` field:
//   - If the DB Connection already exists on the destination Workspace, set it as `true` to overwrite it.
func (s *supersetAPIsOpenSourceGreaterThanDatasets) PostAPIV1DatasetImport(ctx context.Context, request operations.PostAPIV1DatasetImportRequest) (*operations.PostAPIV1DatasetImportResponse, error) {
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url := strings.TrimSuffix(baseURL, "/") + "/api/v1/dataset/import/"

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request, "RequestBody", "multipart")
	if err != nil {
		return nil, fmt.Errorf("error serializing request body: %w", err)
	}

	debugBody := bytes.NewBuffer([]byte{})
	debugReader := io.TeeReader(bodyReader, debugBody)

	req, err := http.NewRequestWithContext(ctx, "POST", url, debugReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "*/*")
	req.Header.Set("user-agent", fmt.Sprintf("speakeasy-sdk/%s %s %s %s", s.sdkConfiguration.Language, s.sdkConfiguration.SDKVersion, s.sdkConfiguration.GenVersion, s.sdkConfiguration.OpenAPIDocVersion))

	req.Header.Set("Content-Type", reqContentType)

	utils.PopulateHeaders(ctx, req, request)

	client := s.sdkConfiguration.SecurityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Request.Body = io.NopCloser(debugBody)
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.PostAPIV1DatasetImportResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
	}

	return res, nil
}

// PutAPIV1DatasetDatasetID - Update a Virtual Dataset
// Updates a Virtual Dataset.
//
// Replace in the URL and on the `Referer` header:
//
// - `{{WorkspaceSlug}}` with the `name` retrieved through the API (using the **Get Workspaces from a Team** endpoint).
// - `{{WorkspaceRegion}}` corresponding to the `region` retrieved from the **Get Workspaces from a Team** endpoint. Refer to below table:
//
// | **`region`** | **`WorkspaceRegion`** |
// | --- | --- |
// | us-east-1 | `us2a` |
// | us-west-2 | `us1a` |
// | eu-north-1 | `eu5a` |
// | ap-northeast-1 | `ap1a` |
//
// Alternatively, access the Workspace through the UI, and get the `{{WorkspaceSlug}}` and `{{WorkspaceRegion}}` from the URL -> `https://{{WorkspaceSlug}}.{{Region}}.preset.io/superset/welcome/`.
//
// - `{{DatasetID}}` with the Dataset `id` retrieved from the **Get All Datasets from a Workspace** endpoint.
//
// Replace in the body:
//
// - `{{DBID}}` with the database `id` retrieved using the **Get Databases** endoint.
// - `{{SchemaName}}` with the desired schema.
// - `{{TableName}}` with the desired table.
// - `{{DatasetSQL}}` with the SQL query to power your dataset.
// - `{{OwnerID}}` with the `id` for the desired owner account(s) retrieved from the **Get all possible Dataset Owners** endpoint. Use a comma to separate multiple IDs (for example, `[2,5]`).
func (s *supersetAPIsOpenSourceGreaterThanDatasets) PutAPIV1DatasetDatasetID(ctx context.Context, request operations.PutAPIV1DatasetDatasetIDRequest) (*operations.PutAPIV1DatasetDatasetIDResponse, error) {
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url, err := utils.GenerateURL(ctx, baseURL, "/api/v1/dataset/{DatasetID}", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request, "RequestBody", "json")
	if err != nil {
		return nil, fmt.Errorf("error serializing request body: %w", err)
	}

	debugBody := bytes.NewBuffer([]byte{})
	debugReader := io.TeeReader(bodyReader, debugBody)

	req, err := http.NewRequestWithContext(ctx, "PUT", url, debugReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "*/*")
	req.Header.Set("user-agent", fmt.Sprintf("speakeasy-sdk/%s %s %s %s", s.sdkConfiguration.Language, s.sdkConfiguration.SDKVersion, s.sdkConfiguration.GenVersion, s.sdkConfiguration.OpenAPIDocVersion))

	req.Header.Set("Content-Type", reqContentType)

	utils.PopulateHeaders(ctx, req, request)

	client := s.sdkConfiguration.SecurityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Request.Body = io.NopCloser(debugBody)
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.PutAPIV1DatasetDatasetIDResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
	}

	return res, nil
}

// PutAPIV1DatasetDatasetIDRefresh - Refresh a Dataset
// Refreshes and updates columns of a Dataset.
//
// Replace in the URL and on the `Referer` header:
//
// - `{{WorkspaceSlug}}` with the `name` retrieved through the API (using the **Get Workspaces from a Team** endpoint).
// - `{{WorkspaceRegion}}` corresponding to the `region` retrieved from the **Get Workspaces from a Team** endpoint. Refer to below table:
//
// | **`region`** | **`WorkspaceRegion`** |
// | --- | --- |
// | us-east-1 | `us2a` |
// | us-west-2 | `us1a` |
// | eu-north-1 | `eu5a` |
// | ap-northeast-1 | `ap1a` |
//
// Alternatively, access the Workspace through the UI, and get the `{{WorkspaceSlug}}` and `{{WorkspaceRegion}}` from the URL -> `https://{{WorkspaceSlug}}.{{Region}}.preset.io/superset/welcome/`.
//
// - `{{DatasetID}}` with the Dataset `id` retrieved from the **Get All Datasets from a Workspace** endpoint.
func (s *supersetAPIsOpenSourceGreaterThanDatasets) PutAPIV1DatasetDatasetIDRefresh(ctx context.Context, request operations.PutAPIV1DatasetDatasetIDRefreshRequest) (*operations.PutAPIV1DatasetDatasetIDRefreshResponse, error) {
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url, err := utils.GenerateURL(ctx, baseURL, "/api/v1/dataset/{DatasetID}/refresh", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request, "RequestBody", "json")
	if err != nil {
		return nil, fmt.Errorf("error serializing request body: %w", err)
	}

	debugBody := bytes.NewBuffer([]byte{})
	debugReader := io.TeeReader(bodyReader, debugBody)

	req, err := http.NewRequestWithContext(ctx, "PUT", url, debugReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "*/*")
	req.Header.Set("user-agent", fmt.Sprintf("speakeasy-sdk/%s %s %s %s", s.sdkConfiguration.Language, s.sdkConfiguration.SDKVersion, s.sdkConfiguration.GenVersion, s.sdkConfiguration.OpenAPIDocVersion))

	req.Header.Set("Content-Type", reqContentType)

	utils.PopulateHeaders(ctx, req, request)

	client := s.sdkConfiguration.SecurityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Request.Body = io.NopCloser(debugBody)
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.PutAPIV1DatasetDatasetIDRefreshResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
	}

	return res, nil
}