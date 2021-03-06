package def

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
)

// Commit commit
// swagger:model commit
type Commit struct {

	// author
	Author *CommitAuthor `json:"author,omitempty"`

	// commit
	Commit *CommitCommit `json:"commit,omitempty"`

	// committer
	Committer *CommitCommitter `json:"committer,omitempty"`

	// files
	Files []*CommitFilesItems0 `json:"files"`

	// parents
	Parents []*CommitParentsItems0 `json:"parents"`

	// sha
	Sha string `json:"sha,omitempty"`

	// stats
	Stats *CommitStats `json:"stats,omitempty"`

	// url
	URL string `json:"url,omitempty"`
}

// Validate validates this commit
func (m *Commit) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuthor(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateCommit(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateCommitter(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateFiles(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateParents(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStats(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Commit) validateAuthor(formats strfmt.Registry) error {

	if swag.IsZero(m.Author) { // not required
		return nil
	}

	if m.Author != nil {

		if err := m.Author.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *Commit) validateCommit(formats strfmt.Registry) error {

	if swag.IsZero(m.Commit) { // not required
		return nil
	}

	if m.Commit != nil {

		if err := m.Commit.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *Commit) validateCommitter(formats strfmt.Registry) error {

	if swag.IsZero(m.Committer) { // not required
		return nil
	}

	if m.Committer != nil {

		if err := m.Committer.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *Commit) validateFiles(formats strfmt.Registry) error {

	if swag.IsZero(m.Files) { // not required
		return nil
	}

	for i := 0; i < len(m.Files); i++ {

		if swag.IsZero(m.Files[i]) { // not required
			continue
		}

		if m.Files[i] != nil {

			if err := m.Files[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *Commit) validateParents(formats strfmt.Registry) error {

	if swag.IsZero(m.Parents) { // not required
		return nil
	}

	for i := 0; i < len(m.Parents); i++ {

		if swag.IsZero(m.Parents[i]) { // not required
			continue
		}

		if m.Parents[i] != nil {

			if err := m.Parents[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *Commit) validateStats(formats strfmt.Registry) error {

	if swag.IsZero(m.Stats) { // not required
		return nil
	}

	if m.Stats != nil {

		if err := m.Stats.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

// CommitAuthor commit author
// swagger:model CommitAuthor
type CommitAuthor struct {

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

// Validate validates this commit author
func (m *CommitAuthor) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// CommitCommit commit commit
// swagger:model CommitCommit
type CommitCommit struct {

	// author
	Author *CommitCommitAuthor `json:"author,omitempty"`

	// committer
	Committer *CommitCommitCommitter `json:"committer,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// tree
	Tree *CommitCommitTree `json:"tree,omitempty"`

	// url
	URL string `json:"url,omitempty"`
}

// Validate validates this commit commit
func (m *CommitCommit) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuthor(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateCommitter(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTree(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CommitCommit) validateAuthor(formats strfmt.Registry) error {

	if swag.IsZero(m.Author) { // not required
		return nil
	}

	if m.Author != nil {

		if err := m.Author.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *CommitCommit) validateCommitter(formats strfmt.Registry) error {

	if swag.IsZero(m.Committer) { // not required
		return nil
	}

	if m.Committer != nil {

		if err := m.Committer.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *CommitCommit) validateTree(formats strfmt.Registry) error {

	if swag.IsZero(m.Tree) { // not required
		return nil
	}

	if m.Tree != nil {

		if err := m.Tree.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

// CommitCommitAuthor commit commit author
// swagger:model CommitCommitAuthor
type CommitCommitAuthor struct {

	// ISO 8601 format: YYYY-MM-DDTHH:MM:SSZ
	Date string `json:"date,omitempty"`

	// email
	Email string `json:"email,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this commit commit author
func (m *CommitCommitAuthor) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// CommitCommitCommitter commit commit committer
// swagger:model CommitCommitCommitter
type CommitCommitCommitter struct {

	// ISO 8601 format: YYYY-MM-DDTHH:MM:SSZ
	Date string `json:"date,omitempty"`

	// email
	Email string `json:"email,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this commit commit committer
func (m *CommitCommitCommitter) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// CommitCommitTree commit commit tree
// swagger:model CommitCommitTree
type CommitCommitTree struct {

	// sha
	Sha string `json:"sha,omitempty"`

	// url
	URL string `json:"url,omitempty"`
}

// Validate validates this commit commit tree
func (m *CommitCommitTree) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// CommitCommitter commit committer
// swagger:model CommitCommitter
type CommitCommitter struct {

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

// Validate validates this commit committer
func (m *CommitCommitter) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// CommitFilesItems0 commit files items0
// swagger:model CommitFilesItems0
type CommitFilesItems0 struct {

	// additions
	Additions int64 `json:"additions,omitempty"`

	// blob url
	BlobURL string `json:"blob_url,omitempty"`

	// changes
	Changes int64 `json:"changes,omitempty"`

	// deletions
	Deletions int64 `json:"deletions,omitempty"`

	// filename
	Filename string `json:"filename,omitempty"`

	// patch
	Patch string `json:"patch,omitempty"`

	// raw url
	RawURL string `json:"raw_url,omitempty"`

	// status
	Status string `json:"status,omitempty"`
}

// Validate validates this commit files items0
func (m *CommitFilesItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// CommitParentsItems0 commit parents items0
// swagger:model CommitParentsItems0
type CommitParentsItems0 struct {

	// sha
	Sha string `json:"sha,omitempty"`

	// url
	URL string `json:"url,omitempty"`
}

// Validate validates this commit parents items0
func (m *CommitParentsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// CommitStats commit stats
// swagger:model CommitStats
type CommitStats struct {

	// additions
	Additions int64 `json:"additions,omitempty"`

	// deletions
	Deletions int64 `json:"deletions,omitempty"`

	// total
	Total int64 `json:"total,omitempty"`
}

// Validate validates this commit stats
func (m *CommitStats) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
