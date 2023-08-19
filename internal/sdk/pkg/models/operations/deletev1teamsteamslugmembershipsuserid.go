// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type DeleteV1TeamsTeamSlugMembershipsUserIDRequest struct {
	TeamSlug string `pathParam:"style=simple,explode=false,name=TeamSlug"`
	UserID   string `pathParam:"style=simple,explode=false,name=UserID"`
}

type DeleteV1TeamsTeamSlugMembershipsUserIDResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
