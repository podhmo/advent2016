package def

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
)

// Readme readme
// swagger:model readme
type Readme struct {

	// links
	Links *ReadmeLinks `json:"_links,omitempty"`

	// content
	Content string `json:"content,omitempty"`

	// encoding
	Encoding string `json:"encoding,omitempty"`

	// git url
	GitURL string `json:"git_url,omitempty"`

	// html url
	HTMLURL string `json:"html_url,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// path
	Path string `json:"path,omitempty"`

	// sha
	Sha string `json:"sha,omitempty"`

	// size
	Size int64 `json:"size,omitempty"`

	// type
	Type string `json:"type,omitempty"`

	// url
	URL string `json:"url,omitempty"`
}

// Validate validates this readme
func (m *Readme) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Readme) validateLinks(formats strfmt.Registry) error {

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

// ReadmeLinks readme links
// swagger:model ReadmeLinks
type ReadmeLinks struct {

	// git
	Git string `json:"git,omitempty"`

	// html
	HTML string `json:"html,omitempty"`

	// self
	Self string `json:"self,omitempty"`
}

// Validate validates this readme links
func (m *ReadmeLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
