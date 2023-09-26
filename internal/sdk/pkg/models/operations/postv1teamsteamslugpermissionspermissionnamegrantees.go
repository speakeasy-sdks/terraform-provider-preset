// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type PostV1TeamsTeamSlugPermissionsPermissionNameGranteesRequestBody struct {
}

type PostV1TeamsTeamSlugPermissionsPermissionNameGranteesRequest struct {
	PermissionName string                                                           `pathParam:"style=simple,explode=false,name=PermissionName"`
	RequestBody    *PostV1TeamsTeamSlugPermissionsPermissionNameGranteesRequestBody `request:"mediaType=application/json"`
	TeamSlug       string                                                           `pathParam:"style=simple,explode=false,name=TeamSlug"`
}

type PostV1TeamsTeamSlugPermissionsPermissionNameGranteesResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}
