package def

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
)

// PullsComment pulls comment
// swagger:model pullsComment
type PullsComment struct {

	// links
	Links *PullsCommentLinks `json:"_links,omitempty"`

	// body
	Body string `json:"body,omitempty"`

	// commit id
	CommitID string `json:"commit_id,omitempty"`

	// ISO 8601 format: YYYY-MM-DDTHH:MM:SSZ
	CreatedAt string `json:"created_at,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// path
	Path string `json:"path,omitempty"`

	// position
	Position int64 `json:"position,omitempty"`

	// ISO 8601 format: YYYY-MM-DDTHH:MM:SSZ
	UpdatedAt string `json:"updated_at,omitempty"`

	// url
	URL string `json:"url,omitempty"`

	// user
	User *PullsCommentUser `json:"user,omitempty"`
}

// Validate validates this pulls comment
func (m *PullsComment) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateUser(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PullsComment) validateLinks(formats strfmt.Registry) error {

	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {

		if err := m.Links.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *PullsComment) validateUser(formats strfmt.Registry) error {

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

// PullsCommentLinks pulls comment links
// swagger:model PullsCommentLinks
type PullsCommentLinks struct {

	// html
	HTML *PullsCommentLinksHTML `json:"html,omitempty"`

	// pull request
	PullRequest *PullsCommentLinksPullRequest `json:"pull_request,omitempty"`

	// self
	Self *PullsCommentLinksSelf `json:"self,omitempty"`
}

// Validate validates this pulls comment links
func (m *PullsCommentLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHTML(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePullRequest(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSelf(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PullsCommentLinks) validateHTML(formats strfmt.Registry) error {

	if swag.IsZero(m.HTML) { // not required
		return nil
	}

	if m.HTML != nil {

		if err := m.HTML.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *PullsCommentLinks) validatePullRequest(formats strfmt.Registry) error {

	if swag.IsZero(m.PullRequest) { // not required
		return nil
	}

	if m.PullRequest != nil {

		if err := m.PullRequest.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *PullsCommentLinks) validateSelf(formats strfmt.Registry) error {

	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {

		if err := m.Self.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

// PullsCommentLinksHTML pulls comment links HTML
// swagger:model PullsCommentLinksHTML
type PullsCommentLinksHTML struct {

	// href
	Href string `json:"href,omitempty"`
}

// Validate validates this pulls comment links HTML
func (m *PullsCommentLinksHTML) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// PullsCommentLinksPullRequest pulls comment links pull request
// swagger:model PullsCommentLinksPullRequest
type PullsCommentLinksPullRequest struct {

	// href
	Href string `json:"href,omitempty"`
}

// Validate validates this pulls comment links pull request
func (m *PullsCommentLinksPullRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// PullsCommentLinksSelf pulls comment links self
// swagger:model PullsCommentLinksSelf
type PullsCommentLinksSelf struct {

	// href
	Href string `json:"href,omitempty"`
}

// Validate validates this pulls comment links self
func (m *PullsCommentLinksSelf) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// PullsCommentUser pulls comment user
// swagger:model PullsCommentUser
type PullsCommentUser struct {

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

// Validate validates this pulls comment user
func (m *PullsCommentUser) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}