// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type PostAPIV1ChartDataRequestBody struct {
}

type PostAPIV1ChartDataRequest struct {
	RequestBody *PostAPIV1ChartDataRequestBody `request:"mediaType=application/json"`
	// Flag to force refresh data.
	Force *string `queryParam:"style=form,explode=true,name=force"`
	// ID of the Chart to be refreshed.
	SliceID *string `queryParam:"style=form,explode=true,name=slice_id"`
}

type PostAPIV1ChartDataResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}