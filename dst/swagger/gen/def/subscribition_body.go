package def

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

// SubscribitionBody subscribition body
// swagger:model subscribitionBody
type SubscribitionBody struct {

	// ignored
	Ignored bool `json:"ignored,omitempty"`

	// subscribed
	Subscribed bool `json:"subscribed,omitempty"`
}

// Validate validates this subscribition body
func (m *SubscribitionBody) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
