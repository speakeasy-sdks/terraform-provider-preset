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

func (o *PostAPIV1AnnotationLayerRequest) GetReferer() *string {
	if o == nil {
		return nil
	}
	return o.Referer
}

func (o *PostAPIV1AnnotationLayerRequest) GetRequestBody() *PostAPIV1AnnotationLayerRequestBody {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

type PostAPIV1AnnotationLayerResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PostAPIV1AnnotationLayerResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PostAPIV1AnnotationLayerResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PostAPIV1AnnotationLayerResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
