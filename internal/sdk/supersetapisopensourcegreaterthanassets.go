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

// supersetAPIsOpenSourceGreaterThanAssets - APIs to export/import an `assets` ZIP file from the Workspace, which includes all:
//
// *   databases.
// *   datasets.
// *   charts.
// *   saved queries.
type supersetAPIsOpenSourceGreaterThanAssets struct {
	sdkConfiguration sdkConfiguration
}

func newSupersetAPIsOpenSourceGreaterThanAssets(sdkConfig sdkConfiguration) *supersetAPIsOpenSourceGreaterThanAssets {
	return &supersetAPIsOpenSourceGreaterThanAssets{
		sdkConfiguration: sdkConfig,
	}
}

// GetAPIV1AssetsExport - Export Assets
// Generates and export a ZIP file from the Workspace containing all:
//
// *   Databases
// *   Datasets
// *   Charts
// *   Dashboards
// *   Saved Queries
//
// Replace in the URL:
//
// *   `{{WorkspaceSlug}}` with the `name` retrieved through the API (using the **Get Workspaces from a Team** endpoint).
// *   `{{WorkspaceRegion}}` corresponding to the `region` retrieved from the **Get Workspaces from a Team** endpoint. Refer to below table:
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
// ***Tip:*** If used in Postman, select `Save Response` and `Save to a File` to get the zip export.
func (s *supersetAPIsOpenSourceGreaterThanAssets) GetAPIV1AssetsExport(ctx context.Context, request operations.GetAPIV1AssetsExportRequest) (*operations.GetAPIV1AssetsExportResponse, error) {
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url := strings.TrimSuffix(baseURL, "/") + "/api/v1/assets/export/"

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "*/*")
	req.Header.Set("user-agent", s.sdkConfiguration.UserAgent)

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

	res := &operations.GetAPIV1AssetsExportResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
	}

	return res, nil
}

// PostAPIV1AssetsImport - Import Assets
// Imports an `assets` ZIP file.
//
// Replace in the URL:
//
// *   `{{WorkspaceSlug}}` with the `name` retrieved through the API (using the **Get Workspaces from a Team** endpoint).
// *   `{{WorkspaceRegion}}` corresponding to the `region` retrieved from the **Get Workspaces from a Team** endpoint. Refer to below table:
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
// *   `{{DatabaseYAMLFile}}` by the database YAML File you can find in your chart export under the folder `Databases`.
// *   `{{DatabasePassword}}` by your database password
// *   Chose your Chart Export Zip file as a value for the `formData`.
func (s *supersetAPIsOpenSourceGreaterThanAssets) PostAPIV1AssetsImport(ctx context.Context, request operations.PostAPIV1AssetsImportRequest) (*operations.PostAPIV1AssetsImportResponse, error) {
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url := strings.TrimSuffix(baseURL, "/") + "/api/v1/assets/import/"

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
	req.Header.Set("user-agent", s.sdkConfiguration.UserAgent)

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

	res := &operations.PostAPIV1AssetsImportResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
	}

	return res, nil
}
