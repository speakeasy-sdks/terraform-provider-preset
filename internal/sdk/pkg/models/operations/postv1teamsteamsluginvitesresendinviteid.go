// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type PostV1TeamsTeamSlugInvitesResendInviteIDRequest struct {
	InviteID string `pathParam:"style=simple,explode=false,name=InviteID"`
	TeamSlug string `pathParam:"style=simple,explode=false,name=TeamSlug"`
}

type PostV1TeamsTeamSlugInvitesResendInviteIDResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
