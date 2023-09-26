// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type PatchV1TeamsTeamSlugMembershipsUserIDRequestBody struct {
}

type PatchV1TeamsTeamSlugMembershipsUserIDRequest struct {
	RequestBody *PatchV1TeamsTeamSlugMembershipsUserIDRequestBody `request:"mediaType=application/json"`
	TeamSlug    string                                            `pathParam:"style=simple,explode=false,name=TeamSlug"`
	UserID      string                                            `pathParam:"style=simple,explode=false,name=UserID"`
}

type PatchV1TeamsTeamSlugMembershipsUserIDResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}
