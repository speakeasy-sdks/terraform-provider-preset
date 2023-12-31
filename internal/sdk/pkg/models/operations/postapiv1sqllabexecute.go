// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type PostAPIV1SqllabExecuteRequestBody struct {
}

type PostAPIV1SqllabExecuteRequest struct {
	Referer     *string                            `header:"style=simple,explode=false,name=Referer"`
	RequestBody *PostAPIV1SqllabExecuteRequestBody `request:"mediaType=application/json"`
}

func (o *PostAPIV1SqllabExecuteRequest) GetReferer() *string {
	if o == nil {
		return nil
	}
	return o.Referer
}

func (o *PostAPIV1SqllabExecuteRequest) GetRequestBody() *PostAPIV1SqllabExecuteRequestBody {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

type PostAPIV1SqllabExecuteResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PostAPIV1SqllabExecuteResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PostAPIV1SqllabExecuteResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PostAPIV1SqllabExecuteResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
