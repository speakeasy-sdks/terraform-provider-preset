// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type PutAPIV1DatasetDatasetIDRequestBody struct {
}

type PutAPIV1DatasetDatasetIDRequest struct {
	DatasetID   string                               `pathParam:"style=simple,explode=false,name=DatasetID"`
	Referer     *string                              `header:"style=simple,explode=false,name=Referer"`
	RequestBody *PutAPIV1DatasetDatasetIDRequestBody `request:"mediaType=application/json"`
}

type PutAPIV1DatasetDatasetIDResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}