// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type PutV1TeamsTeamSlugRequestBody struct {
}

type PutV1TeamsTeamSlugRequest struct {
	RequestBody *PutV1TeamsTeamSlugRequestBody `request:"mediaType=application/json"`
	TeamSlug    string                         `pathParam:"style=simple,explode=false,name=TeamSlug"`
}

func (o *PutV1TeamsTeamSlugRequest) GetRequestBody() *PutV1TeamsTeamSlugRequestBody {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

func (o *PutV1TeamsTeamSlugRequest) GetTeamSlug() string {
	if o == nil {
		return ""
	}
	return o.TeamSlug
}

type PutV1TeamsTeamSlugResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PutV1TeamsTeamSlugResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PutV1TeamsTeamSlugResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PutV1TeamsTeamSlugResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
