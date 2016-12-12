package def

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

// TeamMembership team membership
// swagger:model teamMembership
type TeamMembership struct {

	// state
	State string `json:"state,omitempty"`

	// url
	URL string `json:"url,omitempty"`
}

// Validate validates this team membership
func (m *TeamMembership) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}