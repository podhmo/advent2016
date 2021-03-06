package def

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
)

// UserUserIDSubscribitions user user Id subscribitions
// swagger:model user-userId-subscribitions
type UserUserIDSubscribitions []*UserUserIDSubscribitionsItems0

// Validate validates this user user Id subscribitions
func (m UserUserIDSubscribitions) Validate(formats strfmt.Registry) error {
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

// UserUserIDSubscribitionsItems0 user user ID subscribitions items0
// swagger:model UserUserIDSubscribitionsItems0
type UserUserIDSubscribitionsItems0 struct {

	// clone url
	CloneURL string `json:"clone_url,omitempty"`

	// ISO 8601 format: YYYY-MM-DDTHH:MM:SSZ
	CreatedAt string `json:"created_at,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// fork
	Fork bool `json:"fork,omitempty"`

	// forks
	Forks int64 `json:"forks,omitempty"`

	// forks count
	ForksCount int64 `json:"forks_count,omitempty"`

	// full name
	FullName string `json:"full_name,omitempty"`

	// git url
	GitURL string `json:"git_url,omitempty"`

	// homepage
	Homepage string `json:"homepage,omitempty"`

	// html url
	HTMLURL string `json:"html_url,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// language
	Language string `json:"language,omitempty"`

	// master branch
	MasterBranch int64 `json:"master_branch,omitempty"`

	// mirror url
	MirrorURL string `json:"mirror_url,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// open issues
	OpenIssues int64 `json:"open_issues,omitempty"`

	// open issues count
	OpenIssuesCount int64 `json:"open_issues_count,omitempty"`

	// owner
	Owner *UserUserIDSubscribitionsItems0Owner `json:"owner,omitempty"`

	// private
	Private bool `json:"private,omitempty"`

	// ISO 8601 format: YYYY-MM-DDTHH:MM:SSZ
	PushedAt string `json:"pushed_at,omitempty"`

	// size
	Size int64 `json:"size,omitempty"`

	// ssh url
	SSHURL string `json:"ssh_url,omitempty"`

	// svn url
	SvnURL string `json:"svn_url,omitempty"`

	// ISO 8601 format: YYYY-MM-DDTHH:MM:SSZ
	UpdatedAt string `json:"updated_at,omitempty"`

	// url
	URL string `json:"url,omitempty"`

	// watchers
	Watchers int64 `json:"watchers,omitempty"`

	// watchers count
	WatchersCount int64 `json:"watchers_count,omitempty"`
}

// Validate validates this user user ID subscribitions items0
func (m *UserUserIDSubscribitionsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOwner(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UserUserIDSubscribitionsItems0) validateOwner(formats strfmt.Registry) error {

	if swag.IsZero(m.Owner) { // not required
		return nil
	}

	if m.Owner != nil {

		if err := m.Owner.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

// UserUserIDSubscribitionsItems0Owner user user ID subscribitions items0 owner
// swagger:model UserUserIDSubscribitionsItems0Owner
type UserUserIDSubscribitionsItems0Owner struct {

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

// Validate validates this user user ID subscribitions items0 owner
func (m *UserUserIDSubscribitionsItems0Owner) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
