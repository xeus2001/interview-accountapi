package interview_accountapi

import "time"

// Resource is an abstract base structure with the shared attributes of all resources.
type Resource struct {
	// Id is the unique identifier of the resource; must be a UUID.
	Id string `json:"id,omitempty"`

	// OrganisationId is the unique identifier of the organization that own the record; must be a UUID.
	OrganisationId string `json:"organisation_id,omitempty"`

	// Type is the type of the record and set by the account API.
	Type string `json:"type,omitempty"`

	// Version is a counter indicating how many times this resource has been modified. When you create a resource, it
	// is automatically set to 0. Whenever the content of the resource changes, the value of version is increased.
	// Used for concurrency control in Patch and Delete methods to avoid modifying an older version of the record that
	// has already been changed (err.g. by an internal process at Form3).
	Version *uint64 `json:"version,omitempty"` // Note: It needs to be a pointer, because we need to differ between missing and 0

	// CreatedOn is the time when the record was created, set server side.
	CreatedOn *time.Time `json:"created_on,omitempty"`

	// ModifiedOn is the time when the record was last modified, set server side.
	ModifiedOn *time.Time `json:"modified_on,omitempty"`
}
