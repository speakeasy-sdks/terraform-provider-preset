// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type GetAPIV1DatabaseDatabaseIDRequest struct {
	DatabaseID string `pathParam:"style=simple,explode=false,name=DatabaseID"`
}

type GetAPIV1DatabaseDatabaseIDResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
