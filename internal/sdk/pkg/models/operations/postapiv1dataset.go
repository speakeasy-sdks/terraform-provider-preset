// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type PostAPIV1DatasetRequestBody struct {
}

type PostAPIV1DatasetRequest struct {
	RequestBody *PostAPIV1DatasetRequestBody `request:"mediaType=application/json"`
}

type PostAPIV1DatasetResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}
