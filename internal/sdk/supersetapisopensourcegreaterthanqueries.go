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

type SupersetAPIsOpenSourceGreaterThanQueries struct {
	sdkConfiguration sdkConfiguration
}

func newSupersetAPIsOpenSourceGreaterThanQueries(sdkConfig sdkConfiguration) *SupersetAPIsOpenSourceGreaterThanQueries {
	return &SupersetAPIsOpenSourceGreaterThanQueries{
		sdkConfiguration: sdkConfig,
	}
}

// GetAPIV1Query - Get All Workspace Queries
func (s *SupersetAPIsOpenSourceGreaterThanQueries) GetAPIV1Query(ctx context.Context) (*operations.GetAPIV1QueryResponse, error) {
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url := strings.TrimSuffix(baseURL, "/") + "/api/v1/query/"

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "*/*")
	req.Header.Set("user-agent", s.sdkConfiguration.UserAgent)

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

	res := &operations.GetAPIV1QueryResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
	}

	return res, nil
}
