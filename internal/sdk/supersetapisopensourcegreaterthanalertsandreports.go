// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package sdk

import (
	"bytes"
	"context"
	"fmt"
	"github.com/Preset/terraform-provider-Preset/internal/sdk/pkg/models/operations"
	"github.com/Preset/terraform-provider-Preset/internal/sdk/pkg/utils"
	"io"
	"net/http"
	"strings"
)

// SupersetAPIsOpenSourceGreaterThanAlertsAndReports - APIs to manage your Alerts & Reports.
type SupersetAPIsOpenSourceGreaterThanAlertsAndReports struct {
	sdkConfiguration sdkConfiguration
}

func newSupersetAPIsOpenSourceGreaterThanAlertsAndReports(sdkConfig sdkConfiguration) *SupersetAPIsOpenSourceGreaterThanAlertsAndReports {
	return &SupersetAPIsOpenSourceGreaterThanAlertsAndReports{
		sdkConfiguration: sdkConfig,
	}
}

// GetAPIV1Report - Get all Reports from a Workspace
// Gets all Reports created on the Workspace.
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
// Note that the response includes a `count` value, indicating the total count of Alerts. 100 alerts would be included on the response - if `count > 100`, you can access the remaining items by increasing the `page` value on the `q` parameter:
//
// ```
// ?q=(filters:!((col:type,opr:eq,value:Report)),page_size:{{PageSize}},page:{{Page}})
//
// ```
func (s *SupersetAPIsOpenSourceGreaterThanAlertsAndReports) GetAPIV1Report(ctx context.Context, request operations.GetAPIV1ReportRequest) (*operations.GetAPIV1ReportResponse, error) {
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url := strings.TrimSuffix(baseURL, "/") + "/api/v1/report/"

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "*/*")
	req.Header.Set("user-agent", s.sdkConfiguration.UserAgent)

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

	res := &operations.GetAPIV1ReportResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
	}

	return res, nil
}

// GetAPIV1ReportInfo - Get Alerts & Reports API metadata Info
// Gets metadata information about the Alerts & Reports API endpoints.
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
func (s *SupersetAPIsOpenSourceGreaterThanAlertsAndReports) GetAPIV1ReportInfo(ctx context.Context, request operations.GetAPIV1ReportInfoRequest) (*operations.GetAPIV1ReportInfoResponse, error) {
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url := strings.TrimSuffix(baseURL, "/") + "/api/v1/report/_info"

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

	res := &operations.GetAPIV1ReportInfoResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
	}

	return res, nil
}

// GetAPIV1ReportAlertIDORReportID - Get an Alert/Report
// Gets a specific Alert/Report from the Workspace.
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
// - `{{AlertID OR ReportID}}` with the `id` retrieved from one of the endpoints below:
//   - **Get all Alerts and Reports from a Workspace**
//   - **Get all Alerts from a Workspace**
//   - **Get all Reports from a Workspace**
func (s *SupersetAPIsOpenSourceGreaterThanAlertsAndReports) GetAPIV1ReportAlertIDORReportID(ctx context.Context, request operations.GetAPIV1ReportAlertIDORReportIDRequest) (*operations.GetAPIV1ReportAlertIDORReportIDResponse, error) {
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url := strings.TrimSuffix(baseURL, "/") + "/api/v1/report/{AlertID OR ReportID}"

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

	res := &operations.GetAPIV1ReportAlertIDORReportIDResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
	}

	return res, nil
}

// PostAPIV1Report - Create a Dashboard Alert
// Creates a Dashboard Alert on the Workspace through the API.
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
// Replace in the body:
//
// - `{{ActiveBooleanFlag}}` with:
//   - **true** to create it enabled.
//   - **false** to create it disabled.
//
// - `{{AlertName}}` with a name for your alert.
// - `{{DashboardID}}` with the `id` of the desired Dashboard, retrieved using the **Get all Dashboards from a Workspace** endpoint.
// - `{{AlertDescription}}` with a description for it.
// - `{{CRONSchedule}}` with the desired frequency (in [cron](https://en.wikipedia.org/wiki/Cron)).
// - `{{Timezone}}` with the desired timezone. The list of valid options can be retrieved using the **Get Alerts & Reports API metadata Info** endpoint.
// - `{{ForceBooleanFlag}}` with:
//   - **true** to ignore cache.
//   - **false** to use cache if available.
//
// - `{{DatabaseID}}` with the `id` of the database that should be used to execute the SQL validation. You can retrieve this `id` using the **Get all Database Connections from a Workspace** endpoint.
// - `{{SQLCondition}}` with the SQL query that should be validated by the alert. **Note that the SQL query should return only one column**, for example `select count(\\\\\\*) from {{MyTable}}`.
// - `{{ValidatorType}}` with:
//   - `operator` when performing number comparisson.
//   - `not null` to check if the SQL result is not null. When using `not null`, the `validator_config_json` should be empty:
//
// ``` json
// "validator_type": "not null",
// "validator_config_json": {}
//
// ```
//
// - `{{Operator}}` with the operation that should be used to analyze the SQL result. Available options:
//   - `==` to check if SQL result is equal the threshold value.
//   - `<` to check if the SQL result is smaller than the t
//   - the threshold value.
//   - `>` to check if the SQL result is larger than the the threshold value.
//   - `<=` to check if the SQL result is not larger than the the threshold value.
//   - `>=` to check if the SQL result is not smaller than the the threshold value.
//   - `!=` to check if the SQL result is different than the threshold value.
//
// - `{{Threshold}}` with the condition value.
// - For the `owners` field:
//   - Replace `{{OwnerID}}` with the owner's account ID on the Workspace level (you can retrieve this ID using the **Get all possible Chart Owners** endpoint).
//   - This field is an array, so multiple owners can be added, separated by comma.
//
// - For the `recipients` field:
//   - Replace `{{ReportType}}` with `Email` or `Slack`.
//   - `{{TargetInfo}}` with the email address/Slack user handler.
//   - This field is an array, so multiple recipient configuration can be added (comma separated).
//
// - `{{LogRetention}}` with the retention period (in days). Default and max value is `90`_._
// - `{{WorkingTimeout}}` with time out settings (in seconds). Default value is `3600`.
// - `{{GracePeriod}}` with a grace period (in seconds). Default value is `14400`.
func (s *SupersetAPIsOpenSourceGreaterThanAlertsAndReports) PostAPIV1Report(ctx context.Context, request operations.PostAPIV1ReportRequest) (*operations.PostAPIV1ReportResponse, error) {
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url := strings.TrimSuffix(baseURL, "/") + "/api/v1/report/"

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request, false, true, "RequestBody", "json", `request:"mediaType=application/json"`)
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

	res := &operations.PostAPIV1ReportResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
	}

	return res, nil
}

// PutAPIV1ReportAlertIDORReportID - Disable an Alert/Report
// Disables a specific Alert/Report from the Workspace.
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
// - `{{AlertID OR ReportID}}` with the `id` retrieved from one of the endpoints below:
//   - **Get all Alerts and Reports from a Workspace**
//   - **Get all Alerts from a Workspace**
//   - **Get all Reports from a Workspace**
func (s *SupersetAPIsOpenSourceGreaterThanAlertsAndReports) PutAPIV1ReportAlertIDORReportID(ctx context.Context, request operations.PutAPIV1ReportAlertIDORReportIDRequest) (*operations.PutAPIV1ReportAlertIDORReportIDResponse, error) {
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url := strings.TrimSuffix(baseURL, "/") + "/api/v1/report/{AlertID OR ReportID}"

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request, false, true, "RequestBody", "json", `request:"mediaType=application/json"`)
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

	res := &operations.PutAPIV1ReportAlertIDORReportIDResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
	}

	return res, nil
}
