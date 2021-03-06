package def

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
)

// RateLimit rate limit
// swagger:model rate_limit
type RateLimit struct {

	// rate
	Rate *RateLimitRate `json:"rate,omitempty"`
}

// Validate validates this rate limit
func (m *RateLimit) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRate(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RateLimit) validateRate(formats strfmt.Registry) error {

	if swag.IsZero(m.Rate) { // not required
		return nil
	}

	if m.Rate != nil {

		if err := m.Rate.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

// RateLimitRate rate limit rate
// swagger:model RateLimitRate
type RateLimitRate struct {

	// limit
	Limit int64 `json:"limit,omitempty"`

	// remaining
	Remaining int64 `json:"remaining,omitempty"`

	// reset
	Reset int64 `json:"reset,omitempty"`
}

// Validate validates this rate limit rate
func (m *RateLimitRate) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
