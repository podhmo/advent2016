package def

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
)

// SearchUserByEmail search user by email
// swagger:model search-user-by-email
type SearchUserByEmail struct {

	// user
	User *SearchUserByEmailUser `json:"user,omitempty"`
}

// Validate validates this search user by email
func (m *SearchUserByEmail) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateUser(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SearchUserByEmail) validateUser(formats strfmt.Registry) error {

	if swag.IsZero(m.User) { // not required
		return nil
	}

	if m.User != nil {

		if err := m.User.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

// SearchUserByEmailUser search user by email user
// swagger:model SearchUserByEmailUser
type SearchUserByEmailUser struct {

	// blog
	Blog string `json:"blog,omitempty"`

	// company
	Company string `json:"company,omitempty"`

	// created
	Created string `json:"created,omitempty"`

	// created at
	CreatedAt string `json:"created_at,omitempty"`

	// email
	Email string `json:"email,omitempty"`

	// followers count
	FollowersCount int64 `json:"followers_count,omitempty"`

	// following count
	FollowingCount int64 `json:"following_count,omitempty"`

	// gravatar id
	GravatarID string `json:"gravatar_id,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// location
	Location string `json:"location,omitempty"`

	// login
	Login string `json:"login,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// public gist count
	PublicGistCount int64 `json:"public_gist_count,omitempty"`

	// public repo count
	PublicRepoCount int64 `json:"public_repo_count,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this search user by email user
func (m *SearchUserByEmailUser) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}