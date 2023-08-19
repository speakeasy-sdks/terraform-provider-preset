// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type PutAPIV1AnnotationLayerAnnotationLayerIDAnnotationAnnotationIDRequestBody struct {
}

type PutAPIV1AnnotationLayerAnnotationLayerIDAnnotationAnnotationIDRequest struct {
	AnnotationID      string                                                                     `pathParam:"style=simple,explode=false,name=AnnotationID"`
	AnnotationLayerID string                                                                     `pathParam:"style=simple,explode=false,name=AnnotationLayerID"`
	Referer           *string                                                                    `header:"style=simple,explode=false,name=Referer"`
	RequestBody       *PutAPIV1AnnotationLayerAnnotationLayerIDAnnotationAnnotationIDRequestBody `request:"mediaType=application/json"`
}

type PutAPIV1AnnotationLayerAnnotationLayerIDAnnotationAnnotationIDResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
