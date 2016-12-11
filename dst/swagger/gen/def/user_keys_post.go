package def

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

// UserKeysPost user keys post
// swagger:model user-keys-post
type UserKeysPost struct {

	// key
	Key string `json:"key,omitempty"`

	// title
	Title string `json:"title,omitempty"`
}

// Validate validates this user keys post
func (m *UserKeysPost) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
