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
)

// PresetManagerAPIsGreaterThanEmbedded - APIs associated with the Embedded functionality.
type PresetManagerAPIsGreaterThanEmbedded struct {
	sdkConfiguration sdkConfiguration
}

func newPresetManagerAPIsGreaterThanEmbedded(sdkConfig sdkConfiguration) *PresetManagerAPIsGreaterThanEmbedded {
	return &PresetManagerAPIsGreaterThanEmbedded{
		sdkConfiguration: sdkConfig,
	}
}

// PostAPIV1TeamsTeamSlugWorkspacesWorkspaceSlugGuestToken - Create a new Guest Token
// Creates a new Guest Token to be used with Embedded.
//
// Replace in the URL:
//
// - `{{WorkspaceSlug}}` with the `name` retrieved through the API (using the **Get Workspaces from a Team** endpoint).
// - `{{TeamSlug}}` with the `name` retrieved through the API (using the **Get Preset Teams** endpoint). Alternatively, access the team administration through the UI, and get the `{{TeamSlug}}` from the URL -> `https://manage.app.preset.io/app/teams/{{TeamSlug}}/members`.
//
// For instructions on how to populate the body, refer to [our documentation](https://preset-io.github.io/embedded-beta-docs/v1).
func (s *PresetManagerAPIsGreaterThanEmbedded) PostAPIV1TeamsTeamSlugWorkspacesWorkspaceSlugGuestToken(ctx context.Context, request operations.PostAPIV1TeamsTeamSlugWorkspacesWorkspaceSlugGuestTokenRequest) (*operations.PostAPIV1TeamsTeamSlugWorkspacesWorkspaceSlugGuestTokenResponse, error) {
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url, err := utils.GenerateURL(ctx, baseURL, "/api/v1/teams/{TeamSlug}/workspaces/{WorkspaceSlug}/guest-token/", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

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

	res := &operations.PostAPIV1TeamsTeamSlugWorkspacesWorkspaceSlugGuestTokenResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
	}

	return res, nil
}
