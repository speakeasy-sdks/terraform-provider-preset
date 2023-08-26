// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type PostAPIV1AnnotationLayerRequestBody struct {
}

type PostAPIV1AnnotationLayerRequest struct {
	Referer     *string                              `header:"style=simple,explode=false,name=Referer"`
	RequestBody *PostAPIV1AnnotationLayerRequestBody `request:"mediaType=application/json"`
}

type PostAPIV1AnnotationLayerResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}