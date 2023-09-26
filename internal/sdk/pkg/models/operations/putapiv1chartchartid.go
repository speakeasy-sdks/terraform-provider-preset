// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type PutAPIV1ChartChartIDRequestBody struct {
}

type PutAPIV1ChartChartIDRequest struct {
	ChartID     string                           `pathParam:"style=simple,explode=false,name=ChartID"`
	Referer     *string                          `header:"style=simple,explode=false,name=Referer"`
	RequestBody *PutAPIV1ChartChartIDRequestBody `request:"mediaType=application/json"`
}

type PutAPIV1ChartChartIDResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}
