package def

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
)

// ContributorsStats contributors stats
// swagger:model contributorsStats
type ContributorsStats []*ContributorsStatsItems0

// Validate validates this contributors stats
func (m ContributorsStats) Validate(formats strfmt.Registry) error {
	var res []error

	for i := 0; i < len(m); i++ {

		if swag.IsZero(m[i]) { // not required
			continue
		}

		if m[i] != nil {

			if err := m[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContributorsStatsItems0 contributors stats items0
// swagger:model ContributorsStatsItems0
type ContributorsStatsItems0 struct {

	// author
	Author *ContributorsStatsItems0Author `json:"author,omitempty"`

	// The Total number of commits authored by the contributor.
	Total int64 `json:"total,omitempty"`

	// weeks
	Weeks []*ContributorsStatsItems0WeeksItems0 `json:"weeks"`
}

// Validate validates this contributors stats items0
func (m *ContributorsStatsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuthor(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateWeeks(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ContributorsStatsItems0) validateAuthor(formats strfmt.Registry) error {

	if swag.IsZero(m.Author) { // not required
		return nil
	}

	if m.Author != nil {

		if err := m.Author.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *ContributorsStatsItems0) validateWeeks(formats strfmt.Registry) error {

	if swag.IsZero(m.Weeks) { // not required
		return nil
	}

	for i := 0; i < len(m.Weeks); i++ {

		if swag.IsZero(m.Weeks[i]) { // not required
			continue
		}

		if m.Weeks[i] != nil {

			if err := m.Weeks[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

// ContributorsStatsItems0Author contributors stats items0 author
// swagger:model ContributorsStatsItems0Author
type ContributorsStatsItems0Author struct {

	// avatar url
	AvatarURL string `json:"avatar_url,omitempty"`

	// gravatar id
	GravatarID string `json:"gravatar_id,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// login
	Login string `json:"login,omitempty"`

	// url
	URL string `json:"url,omitempty"`
}

// Validate validates this contributors stats items0 author
func (m *ContributorsStatsItems0Author) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContributorsStatsItems0WeeksItems0 contributors stats items0 weeks items0
// swagger:model ContributorsStatsItems0WeeksItems0
type ContributorsStatsItems0WeeksItems0 struct {

	// Number of additions.
	A int64 `json:"a,omitempty"`

	// Number of commits.
	C int64 `json:"c,omitempty"`

	// Number of deletions.
	D int64 `json:"d,omitempty"`

	// Start of the week.
	W string `json:"w,omitempty"`
}

// Validate validates this contributors stats items0 weeks items0
func (m *ContributorsStatsItems0WeeksItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}