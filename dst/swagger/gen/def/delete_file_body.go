package def

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
)

// DeleteFileBody delete file body
// swagger:model deleteFileBody
type DeleteFileBody struct {

	// committer
	Committer *DeleteFileBodyCommitter `json:"committer,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// sha
	Sha string `json:"sha,omitempty"`
}

// Validate validates this delete file body
func (m *DeleteFileBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCommitter(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeleteFileBody) validateCommitter(formats strfmt.Registry) error {

	if swag.IsZero(m.Committer) { // not required
		return nil
	}

	if m.Committer != nil {

		if err := m.Committer.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

// DeleteFileBodyCommitter delete file body committer
// swagger:model DeleteFileBodyCommitter
type DeleteFileBodyCommitter struct {

	// email
	Email string `json:"email,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this delete file body committer
func (m *DeleteFileBodyCommitter) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
