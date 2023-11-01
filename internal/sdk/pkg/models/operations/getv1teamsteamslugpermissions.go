// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type GetV1TeamsTeamSlugPermissionsRequest struct {
	TeamSlug string `pathParam:"style=simple,explode=false,name=TeamSlug"`
}

func (o *GetV1TeamsTeamSlugPermissionsRequest) GetTeamSlug() string {
	if o == nil {
		return ""
	}
	return o.TeamSlug
}

type GetV1TeamsTeamSlugPermissionsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetV1TeamsTeamSlugPermissionsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetV1TeamsTeamSlugPermissionsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetV1TeamsTeamSlugPermissionsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
