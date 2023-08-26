// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type GetV1TeamsTeamSlugPermissionsRequest struct {
	TeamSlug string `pathParam:"style=simple,explode=false,name=TeamSlug"`
}

type GetV1TeamsTeamSlugPermissionsResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}