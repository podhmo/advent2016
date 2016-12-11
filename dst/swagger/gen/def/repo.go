package def

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
)

// Repo repo
// swagger:model repo
type Repo struct {

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

	// has downloads
	HasDownloads bool `json:"has_downloads,omitempty"`

	// has issues
	HasIssues bool `json:"has_issues,omitempty"`

	// has wiki
	HasWiki bool `json:"has_wiki,omitempty"`

	// homepage
	Homepage string `json:"homepage,omitempty"`

	// html url
	HTMLURL string `json:"html_url,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// language
	Language string `json:"language,omitempty"`

	// master branch
	MasterBranch string `json:"master_branch,omitempty"`

	// mirror url
	MirrorURL string `json:"mirror_url,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// open issues
	OpenIssues int64 `json:"open_issues,omitempty"`

	// open issues count
	OpenIssuesCount int64 `json:"open_issues_count,omitempty"`

	// organization
	Organization *RepoOrganization `json:"organization,omitempty"`

	// owner
	Owner *RepoOwner `json:"owner,omitempty"`

	// parent
	Parent *RepoParent `json:"parent,omitempty"`

	// private
	Private bool `json:"private,omitempty"`

	// ISO 8601 format: YYYY-MM-DDTHH:MM:SSZ
	PushedAt string `json:"pushed_at,omitempty"`

	// size
	Size int64 `json:"size,omitempty"`

	// source
	Source *RepoSource `json:"source,omitempty"`

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

// Validate validates this repo
func (m *Repo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOrganization(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateOwner(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateParent(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSource(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Repo) validateOrganization(formats strfmt.Registry) error {

	if swag.IsZero(m.Organization) { // not required
		return nil
	}

	if m.Organization != nil {

		if err := m.Organization.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *Repo) validateOwner(formats strfmt.Registry) error {

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

func (m *Repo) validateParent(formats strfmt.Registry) error {

	if swag.IsZero(m.Parent) { // not required
		return nil
	}

	if m.Parent != nil {

		if err := m.Parent.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *Repo) validateSource(formats strfmt.Registry) error {

	if swag.IsZero(m.Source) { // not required
		return nil
	}

	if m.Source != nil {

		if err := m.Source.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

// RepoOrganization repo organization
// swagger:model RepoOrganization
type RepoOrganization struct {

	// avatar url
	AvatarURL string `json:"avatar_url,omitempty"`

	// gravatar id
	GravatarID string `json:"gravatar_id,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// login
	Login string `json:"login,omitempty"`

	// type
	Type string `json:"type,omitempty"`

	// url
	URL string `json:"url,omitempty"`
}

// Validate validates this repo organization
func (m *RepoOrganization) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// RepoOwner repo owner
// swagger:model RepoOwner
type RepoOwner struct {

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

// Validate validates this repo owner
func (m *RepoOwner) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// RepoParent Is present when the repo is a fork. Parent is the repo this repo was forked from.
// swagger:model RepoParent
type RepoParent struct {

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
	MasterBranch string `json:"master_branch,omitempty"`

	// mirror url
	MirrorURL string `json:"mirror_url,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// open issues
	OpenIssues int64 `json:"open_issues,omitempty"`

	// open issues count
	OpenIssuesCount int64 `json:"open_issues_count,omitempty"`

	// owner
	Owner *RepoParentOwner `json:"owner,omitempty"`

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

// Validate validates this repo parent
func (m *RepoParent) Validate(formats strfmt.Registry) error {
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

func (m *RepoParent) validateOwner(formats strfmt.Registry) error {

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

// RepoParentOwner repo parent owner
// swagger:model RepoParentOwner
type RepoParentOwner struct {

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

// Validate validates this repo parent owner
func (m *RepoParentOwner) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// RepoSource Is present when the repo is a fork. Source is the ultimate source for the network.
// swagger:model RepoSource
type RepoSource struct {

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
	MasterBranch string `json:"master_branch,omitempty"`

	// mirror url
	MirrorURL string `json:"mirror_url,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// open issues
	OpenIssues int64 `json:"open_issues,omitempty"`

	// open issues count
	OpenIssuesCount int64 `json:"open_issues_count,omitempty"`

	// owner
	Owner *RepoSourceOwner `json:"owner,omitempty"`

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

// Validate validates this repo source
func (m *RepoSource) Validate(formats strfmt.Registry) error {
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

func (m *RepoSource) validateOwner(formats strfmt.Registry) error {

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

// RepoSourceOwner repo source owner
// swagger:model RepoSourceOwner
type RepoSourceOwner struct {

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

// Validate validates this repo source owner
func (m *RepoSourceOwner) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
