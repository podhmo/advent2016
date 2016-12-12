package def

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// Pulls pulls
// swagger:model pulls
type Pulls []*PullsItems0

// Validate validates this pulls
func (m Pulls) Validate(formats strfmt.Registry) error {
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

// PullsItems0 pulls items0
// swagger:model PullsItems0
type PullsItems0 struct {

	// links
	Links *PullsItems0Links `json:"_links,omitempty"`

	// base
	Base *PullsItems0Base `json:"base,omitempty"`

	// body
	Body string `json:"body,omitempty"`

	// ISO 8601 format: YYYY-MM-DDTHH:MM:SSZ
	ClosedAt string `json:"closed_at,omitempty"`

	// ISO 8601 format: YYYY-MM-DDTHH:MM:SSZ
	CreatedAt string `json:"created_at,omitempty"`

	// diff url
	DiffURL string `json:"diff_url,omitempty"`

	// head
	Head *PullsItems0Head `json:"head,omitempty"`

	// html url
	HTMLURL string `json:"html_url,omitempty"`

	// issue url
	IssueURL string `json:"issue_url,omitempty"`

	// ISO 8601 format: YYYY-MM-DDTHH:MM:SSZ
	MergedAt string `json:"merged_at,omitempty"`

	// number
	Number int64 `json:"number,omitempty"`

	// patch url
	PatchURL string `json:"patch_url,omitempty"`

	// state
	State interface{} `json:"state,omitempty"`

	// title
	Title string `json:"title,omitempty"`

	// ISO 8601 format: YYYY-MM-DDTHH:MM:SSZ
	UpdatedAt string `json:"updated_at,omitempty"`

	// url
	URL string `json:"url,omitempty"`

	// user
	User *PullsItems0User `json:"user,omitempty"`
}

// Validate validates this pulls items0
func (m *PullsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateBase(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateHead(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
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

func (m *PullsItems0) validateLinks(formats strfmt.Registry) error {

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

func (m *PullsItems0) validateBase(formats strfmt.Registry) error {

	if swag.IsZero(m.Base) { // not required
		return nil
	}

	if m.Base != nil {

		if err := m.Base.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *PullsItems0) validateHead(formats strfmt.Registry) error {

	if swag.IsZero(m.Head) { // not required
		return nil
	}

	if m.Head != nil {

		if err := m.Head.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

var pullsItems0TypeStatePropEnum []interface{}

func init() {
	var res []interface{}
	if err := json.Unmarshal([]byte(`["open","closed"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		pullsItems0TypeStatePropEnum = append(pullsItems0TypeStatePropEnum, v)
	}
}

// prop value enum
func (m *PullsItems0) validateStateEnum(path, location string, value interface{}) error {
	if err := validate.Enum(path, location, value, pullsItems0TypeStatePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *PullsItems0) validateState(formats strfmt.Registry) error {

	if swag.IsZero(m.State) { // not required
		return nil
	}

	return nil
}

func (m *PullsItems0) validateUser(formats strfmt.Registry) error {

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

// PullsItems0Base pulls items0 base
// swagger:model PullsItems0Base
type PullsItems0Base struct {

	// label
	Label string `json:"label,omitempty"`

	// ref
	Ref string `json:"ref,omitempty"`

	// repo
	Repo *PullsItems0BaseRepo `json:"repo,omitempty"`

	// sha
	Sha string `json:"sha,omitempty"`

	// user
	User *PullsItems0BaseUser `json:"user,omitempty"`
}

// Validate validates this pulls items0 base
func (m *PullsItems0Base) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRepo(formats); err != nil {
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

func (m *PullsItems0Base) validateRepo(formats strfmt.Registry) error {

	if swag.IsZero(m.Repo) { // not required
		return nil
	}

	if m.Repo != nil {

		if err := m.Repo.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *PullsItems0Base) validateUser(formats strfmt.Registry) error {

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

// PullsItems0BaseRepo pulls items0 base repo
// swagger:model PullsItems0BaseRepo
type PullsItems0BaseRepo struct {

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
	Owner *PullsItems0BaseRepoOwner `json:"owner,omitempty"`

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

// Validate validates this pulls items0 base repo
func (m *PullsItems0BaseRepo) Validate(formats strfmt.Registry) error {
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

func (m *PullsItems0BaseRepo) validateOwner(formats strfmt.Registry) error {

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

// PullsItems0BaseRepoOwner pulls items0 base repo owner
// swagger:model PullsItems0BaseRepoOwner
type PullsItems0BaseRepoOwner struct {

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

// Validate validates this pulls items0 base repo owner
func (m *PullsItems0BaseRepoOwner) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// PullsItems0BaseUser pulls items0 base user
// swagger:model PullsItems0BaseUser
type PullsItems0BaseUser struct {

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

// Validate validates this pulls items0 base user
func (m *PullsItems0BaseUser) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// PullsItems0Head pulls items0 head
// swagger:model PullsItems0Head
type PullsItems0Head struct {

	// label
	Label string `json:"label,omitempty"`

	// ref
	Ref string `json:"ref,omitempty"`

	// repo
	Repo *PullsItems0HeadRepo `json:"repo,omitempty"`

	// sha
	Sha string `json:"sha,omitempty"`

	// user
	User *PullsItems0HeadUser `json:"user,omitempty"`
}

// Validate validates this pulls items0 head
func (m *PullsItems0Head) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRepo(formats); err != nil {
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

func (m *PullsItems0Head) validateRepo(formats strfmt.Registry) error {

	if swag.IsZero(m.Repo) { // not required
		return nil
	}

	if m.Repo != nil {

		if err := m.Repo.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *PullsItems0Head) validateUser(formats strfmt.Registry) error {

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

// PullsItems0HeadRepo pulls items0 head repo
// swagger:model PullsItems0HeadRepo
type PullsItems0HeadRepo struct {

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
	Owner *PullsItems0HeadRepoOwner `json:"owner,omitempty"`

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

// Validate validates this pulls items0 head repo
func (m *PullsItems0HeadRepo) Validate(formats strfmt.Registry) error {
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

func (m *PullsItems0HeadRepo) validateOwner(formats strfmt.Registry) error {

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

// PullsItems0HeadRepoOwner pulls items0 head repo owner
// swagger:model PullsItems0HeadRepoOwner
type PullsItems0HeadRepoOwner struct {

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

// Validate validates this pulls items0 head repo owner
func (m *PullsItems0HeadRepoOwner) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// PullsItems0HeadUser pulls items0 head user
// swagger:model PullsItems0HeadUser
type PullsItems0HeadUser struct {

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

// Validate validates this pulls items0 head user
func (m *PullsItems0HeadUser) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// PullsItems0Links pulls items0 links
// swagger:model PullsItems0Links
type PullsItems0Links struct {

	// comments
	Comments *PullsItems0LinksComments `json:"comments,omitempty"`

	// html
	HTML *PullsItems0LinksHTML `json:"html,omitempty"`

	// review comments
	ReviewComments *PullsItems0LinksReviewComments `json:"review_comments,omitempty"`

	// self
	Self *PullsItems0LinksSelf `json:"self,omitempty"`
}

// Validate validates this pulls items0 links
func (m *PullsItems0Links) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateComments(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateHTML(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateReviewComments(formats); err != nil {
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

func (m *PullsItems0Links) validateComments(formats strfmt.Registry) error {

	if swag.IsZero(m.Comments) { // not required
		return nil
	}

	if m.Comments != nil {

		if err := m.Comments.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *PullsItems0Links) validateHTML(formats strfmt.Registry) error {

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

func (m *PullsItems0Links) validateReviewComments(formats strfmt.Registry) error {

	if swag.IsZero(m.ReviewComments) { // not required
		return nil
	}

	if m.ReviewComments != nil {

		if err := m.ReviewComments.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *PullsItems0Links) validateSelf(formats strfmt.Registry) error {

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

// PullsItems0LinksComments pulls items0 links comments
// swagger:model PullsItems0LinksComments
type PullsItems0LinksComments struct {

	// href
	Href string `json:"href,omitempty"`
}

// Validate validates this pulls items0 links comments
func (m *PullsItems0LinksComments) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// PullsItems0LinksHTML pulls items0 links HTML
// swagger:model PullsItems0LinksHTML
type PullsItems0LinksHTML struct {

	// href
	Href string `json:"href,omitempty"`
}

// Validate validates this pulls items0 links HTML
func (m *PullsItems0LinksHTML) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// PullsItems0LinksReviewComments pulls items0 links review comments
// swagger:model PullsItems0LinksReviewComments
type PullsItems0LinksReviewComments struct {

	// href
	Href string `json:"href,omitempty"`
}

// Validate validates this pulls items0 links review comments
func (m *PullsItems0LinksReviewComments) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// PullsItems0LinksSelf pulls items0 links self
// swagger:model PullsItems0LinksSelf
type PullsItems0LinksSelf struct {

	// href
	Href string `json:"href,omitempty"`
}

// Validate validates this pulls items0 links self
func (m *PullsItems0LinksSelf) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// PullsItems0User pulls items0 user
// swagger:model PullsItems0User
type PullsItems0User struct {

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

// Validate validates this pulls items0 user
func (m *PullsItems0User) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}