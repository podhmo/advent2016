package def

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// PostRepo post repo
// swagger:model postRepo
type PostRepo struct {

	// True to create an initial commit with empty README. Default is false.
	AutoInit bool `json:"auto_init,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// Desired language or platform .gitignore template to apply. Use the name of the template without the extension. For example, "Haskell" Ignored if auto_init parameter is not provided.
	GitignoreTemplate string `json:"gitignore_template,omitempty"`

	// True to enable downloads for this repository, false to disable them. Default is true.
	HasDownloads bool `json:"has_downloads,omitempty"`

	// True to enable issues for this repository, false to disable them. Default is true.
	HasIssues bool `json:"has_issues,omitempty"`

	// True to enable the wiki for this repository, false to disable it. Default is true.
	HasWiki bool `json:"has_wiki,omitempty"`

	// homepage
	Homepage string `json:"homepage,omitempty"`

	// name
	// Required: true
	Name *string `json:"name"`

	// True to create a private repository, false to create a public one. Creating private repositories requires a paid GitHub account.
	Private bool `json:"private,omitempty"`

	// The id of the team that will be granted access to this repository. This is only valid when creating a repo in an organization.
	TeamID int64 `json:"team_id,omitempty"`
}

// Validate validates this post repo
func (m *PostRepo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostRepo) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}
