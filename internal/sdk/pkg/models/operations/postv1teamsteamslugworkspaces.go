// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type PostV1TeamsTeamSlugWorkspacesRequestBody struct {
}

type PostV1TeamsTeamSlugWorkspacesRequest struct {
	RequestBody *PostV1TeamsTeamSlugWorkspacesRequestBody `request:"mediaType=application/json"`
	TeamSlug    string                                    `pathParam:"style=simple,explode=false,name=TeamSlug"`
}

type PostV1TeamsTeamSlugWorkspacesResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}
