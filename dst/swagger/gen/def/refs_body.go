package def

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

// RefsBody refs body
// swagger:model refsBody
type RefsBody struct {

	// ref
	Ref string `json:"ref,omitempty"`

	// sha
	Sha string `json:"sha,omitempty"`
}

// Validate validates this refs body
func (m *RefsBody) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
