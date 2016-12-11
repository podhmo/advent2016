package def

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

// NotificationMarkRead notification mark read
// swagger:model notificationMarkRead
type NotificationMarkRead struct {

	// last read at
	LastReadAt string `json:"last_read_at,omitempty"`
}

// Validate validates this notification mark read
func (m *NotificationMarkRead) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
