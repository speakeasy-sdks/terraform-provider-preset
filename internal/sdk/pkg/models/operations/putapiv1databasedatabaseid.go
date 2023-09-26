// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type PutAPIV1DatabaseDatabaseIDRequest struct {
	DatabaseID  string  `pathParam:"style=simple,explode=false,name=DatabaseID"`
	RequestBody *string `request:"mediaType=*/*"`
}

type PutAPIV1DatabaseDatabaseIDResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}
