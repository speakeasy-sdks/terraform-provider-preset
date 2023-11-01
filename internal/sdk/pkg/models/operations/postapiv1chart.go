// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type PostAPIV1ChartRequestBody struct {
}

type PostAPIV1ChartRequest struct {
	RequestBody *PostAPIV1ChartRequestBody `request:"mediaType=application/json"`
}

func (o *PostAPIV1ChartRequest) GetRequestBody() *PostAPIV1ChartRequestBody {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

type PostAPIV1ChartResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PostAPIV1ChartResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PostAPIV1ChartResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PostAPIV1ChartResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
