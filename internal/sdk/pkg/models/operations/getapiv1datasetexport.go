// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type GetAPIV1DatasetExportRequest struct {
	// Comma separated list of Dataset IDs.
	Q *string `queryParam:"style=form,explode=true,name=q"`
}

type GetAPIV1DatasetExportResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}