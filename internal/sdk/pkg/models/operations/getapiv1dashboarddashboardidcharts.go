// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type GetAPIV1DashboardDashboardIDChartsRequest struct {
	DashboardID string `pathParam:"style=simple,explode=false,name=DashboardID"`
}

type GetAPIV1DashboardDashboardIDChartsResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}