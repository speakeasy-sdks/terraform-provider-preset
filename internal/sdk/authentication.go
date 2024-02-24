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

// Authentication - API to authenticate and get a JWT token to interact with the Preset/Superset APIs.
type Authentication struct {
	sdkConfiguration sdkConfiguration
}

func newAuthentication(sdkConfig sdkConfiguration) *Authentication {
	return &Authentication{
		sdkConfiguration: sdkConfig,
	}
}

// PostV1Auth - Get a JWT Token
// To interact with the Preset API, it's required to generate an API Key, that's used to generate a JWT token.
//
// 1. Generate an API Key on the Preset Manager UI. Refer to [our documentation](https://docs.preset.io/docs/the-preset-api).
// 2. Copy the API `token` and `secret`.
//
// Replace in the body:
//
// - `{{APIToken}}` with the `token` from the UI.
// - `{{APISecret}}` with the `secret` from the UI.
func (s *Authentication) PostV1Auth(ctx context.Context, request operations.PostV1AuthRequest) (*operations.PostV1AuthResponse, error) {
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url := strings.TrimSuffix(baseURL, "/") + "/v1/auth/"

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

	client := s.sdkConfiguration.DefaultClient

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

	res := &operations.PostV1AuthResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
	}

	return res, nil
}
