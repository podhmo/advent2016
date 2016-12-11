package def

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
)

// Asset asset
// swagger:model asset
type Asset struct {

	// content type
	ContentType string `json:"content_type,omitempty"`

	// created at
	CreatedAt string `json:"created_at,omitempty"`

	// download count
	DownloadCount float64 `json:"download_count,omitempty"`

	// id
	ID float64 `json:"id,omitempty"`

	// label
	Label string `json:"label,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// size
	Size float64 `json:"size,omitempty"`

	// state
	State string `json:"state,omitempty"`

	// updated at
	UpdatedAt string `json:"updated_at,omitempty"`

	// uploader
	Uploader *AssetUploader `json:"uploader,omitempty"`

	// url
	URL string `json:"url,omitempty"`
}

// Validate validates this asset
func (m *Asset) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateUploader(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Asset) validateUploader(formats strfmt.Registry) error {

	if swag.IsZero(m.Uploader) { // not required
		return nil
	}

	if m.Uploader != nil {

		if err := m.Uploader.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

// AssetUploader asset uploader
// swagger:model AssetUploader
type AssetUploader struct {

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
	ID float64 `json:"id,omitempty"`

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

// Validate validates this asset uploader
func (m *AssetUploader) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
