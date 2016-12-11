package def

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
)

// RepoDeployments repo deployments
// swagger:model repo-deployments
type RepoDeployments []*RepoDeploymentsItems0

// Validate validates this repo deployments
func (m RepoDeployments) Validate(formats strfmt.Registry) error {
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

// RepoDeploymentsItems0 repo deployments items0
// swagger:model RepoDeploymentsItems0
type RepoDeploymentsItems0 struct {

	// created at
	CreatedAt string `json:"created_at,omitempty"`

	// creator
	Creator *RepoDeploymentsItems0Creator `json:"creator,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// payload
	Payload string `json:"payload,omitempty"`

	// sha
	Sha string `json:"sha,omitempty"`

	// statuses url
	StatusesURL string `json:"statuses_url,omitempty"`

	// updated at
	UpdatedAt string `json:"updated_at,omitempty"`

	// url
	URL string `json:"url,omitempty"`
}

// Validate validates this repo deployments items0
func (m *RepoDeploymentsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreator(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RepoDeploymentsItems0) validateCreator(formats strfmt.Registry) error {

	if swag.IsZero(m.Creator) { // not required
		return nil
	}

	if m.Creator != nil {

		if err := m.Creator.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

// RepoDeploymentsItems0Creator repo deployments items0 creator
// swagger:model RepoDeploymentsItems0Creator
type RepoDeploymentsItems0Creator struct {

	// avatar url
	AvatarURL string `json:"avatar_url,omitempty"`

	// events url
	EventsURL string `json:"events_url,omitempty"`

	// followers url
	FollowersURL string `json:"followers_url,omitempty"`

	// following url
	FollowingURL string `json:"following_url,omitempty"`

	// gists url
	GistsURL string `json:"gists_url,omitempty"`

	// gravatar id
	GravatarID string `json:"gravatar_id,omitempty"`

	// html url
	HTMLURL string `json:"html_url,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// login
	Login string `json:"login,omitempty"`

	// organizations url
	OrganizationsURL string `json:"organizations_url,omitempty"`

	// received events url
	ReceivedEventsURL string `json:"received_events_url,omitempty"`

	// repos url
	ReposURL string `json:"repos_url,omitempty"`

	// site admin
	SiteAdmin bool `json:"site_admin,omitempty"`

	// starred url
	StarredURL string `json:"starred_url,omitempty"`

	// subscriptions url
	SubscriptionsURL string `json:"subscriptions_url,omitempty"`

	// type
	Type string `json:"type,omitempty"`

	// url
	URL string `json:"url,omitempty"`
}

// Validate validates this repo deployments items0 creator
func (m *RepoDeploymentsItems0Creator) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
