package def

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
)

// SearchIssues search issues
// swagger:model search-issues
type SearchIssues struct {

	// items
	Items []*SearchIssuesItemsItems0 `json:"items"`

	// total count
	TotalCount int64 `json:"total_count,omitempty"`
}

// Validate validates this search issues
func (m *SearchIssues) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateItems(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SearchIssues) validateItems(formats strfmt.Registry) error {

	if swag.IsZero(m.Items) { // not required
		return nil
	}

	for i := 0; i < len(m.Items); i++ {

		if swag.IsZero(m.Items[i]) { // not required
			continue
		}

		if m.Items[i] != nil {

			if err := m.Items[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

// SearchIssuesItemsItems0 search issues items items0
// swagger:model SearchIssuesItemsItems0
type SearchIssuesItemsItems0 struct {

	// assignee
	Assignee interface{} `json:"assignee,omitempty"`

	// body
	Body string `json:"body,omitempty"`

	// closed at
	ClosedAt interface{} `json:"closed_at,omitempty"`

	// comments
	Comments int64 `json:"comments,omitempty"`

	// comments url
	CommentsURL string `json:"comments_url,omitempty"`

	// created at
	CreatedAt string `json:"created_at,omitempty"`

	// events url
	EventsURL string `json:"events_url,omitempty"`

	// html url
	HTMLURL string `json:"html_url,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// labels
	Labels []*SearchIssuesItemsItems0LabelsItems0 `json:"labels"`

	// labels url
	LabelsURL string `json:"labels_url,omitempty"`

	// milestone
	Milestone interface{} `json:"milestone,omitempty"`

	// number
	Number int64 `json:"number,omitempty"`

	// pull request
	PullRequest *SearchIssuesItemsItems0PullRequest `json:"pull_request,omitempty"`

	// score
	Score float64 `json:"score,omitempty"`

	// state
	State string `json:"state,omitempty"`

	// title
	Title string `json:"title,omitempty"`

	// updated at
	UpdatedAt string `json:"updated_at,omitempty"`

	// url
	URL string `json:"url,omitempty"`

	// user
	User *SearchIssuesItemsItems0User `json:"user,omitempty"`
}

// Validate validates this search issues items items0
func (m *SearchIssuesItemsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLabels(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePullRequest(formats); err != nil {
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

func (m *SearchIssuesItemsItems0) validateLabels(formats strfmt.Registry) error {

	if swag.IsZero(m.Labels) { // not required
		return nil
	}

	for i := 0; i < len(m.Labels); i++ {

		if swag.IsZero(m.Labels[i]) { // not required
			continue
		}

		if m.Labels[i] != nil {

			if err := m.Labels[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *SearchIssuesItemsItems0) validatePullRequest(formats strfmt.Registry) error {

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

func (m *SearchIssuesItemsItems0) validateUser(formats strfmt.Registry) error {

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

// SearchIssuesItemsItems0LabelsItems0 search issues items items0 labels items0
// swagger:model SearchIssuesItemsItems0LabelsItems0
type SearchIssuesItemsItems0LabelsItems0 struct {

	// color
	Color string `json:"color,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// url
	URL string `json:"url,omitempty"`
}

// Validate validates this search issues items items0 labels items0
func (m *SearchIssuesItemsItems0LabelsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// SearchIssuesItemsItems0PullRequest search issues items items0 pull request
// swagger:model SearchIssuesItemsItems0PullRequest
type SearchIssuesItemsItems0PullRequest struct {

	// diff url
	DiffURL interface{} `json:"diff_url,omitempty"`

	// html url
	HTMLURL interface{} `json:"html_url,omitempty"`

	// patch url
	PatchURL interface{} `json:"patch_url,omitempty"`
}

// Validate validates this search issues items items0 pull request
func (m *SearchIssuesItemsItems0PullRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// SearchIssuesItemsItems0User search issues items items0 user
// swagger:model SearchIssuesItemsItems0User
type SearchIssuesItemsItems0User struct {

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

	// starred url
	StarredURL string `json:"starred_url,omitempty"`

	// subscriptions url
	SubscriptionsURL string `json:"subscriptions_url,omitempty"`

	// type
	Type string `json:"type,omitempty"`

	// url
	URL string `json:"url,omitempty"`
}

// Validate validates this search issues items items0 user
func (m *SearchIssuesItemsItems0User) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
