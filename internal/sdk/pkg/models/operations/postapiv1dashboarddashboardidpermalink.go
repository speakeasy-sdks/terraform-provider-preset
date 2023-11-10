// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type PostAPIV1DashboardDashboardIDPermalinkRequestBody struct {
}

type PostAPIV1DashboardDashboardIDPermalinkRequest struct {
	DashboardID string                                             `pathParam:"style=simple,explode=false,name=DashboardID"`
	Referer     *string                                            `header:"style=simple,explode=false,name=Referer"`
	RequestBody *PostAPIV1DashboardDashboardIDPermalinkRequestBody `request:"mediaType=application/json"`
}

func (o *PostAPIV1DashboardDashboardIDPermalinkRequest) GetDashboardID() string {
	if o == nil {
		return ""
	}
	return o.DashboardID
}

func (o *PostAPIV1DashboardDashboardIDPermalinkRequest) GetReferer() *string {
	if o == nil {
		return nil
	}
	return o.Referer
}

func (o *PostAPIV1DashboardDashboardIDPermalinkRequest) GetRequestBody() *PostAPIV1DashboardDashboardIDPermalinkRequestBody {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

type PostAPIV1DashboardDashboardIDPermalinkResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PostAPIV1DashboardDashboardIDPermalinkResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PostAPIV1DashboardDashboardIDPermalinkResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PostAPIV1DashboardDashboardIDPermalinkResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
