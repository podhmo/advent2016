package autogen

/* structure
AuthenticatedUser
	Plan
*/
// AuthenticatedUser : auto generated JSON container
type AuthenticatedUser struct {
	AvatarURL         string `json:"avatar_url" example:"string"`
	Bio               string `json:"bio" example:"string"`
	Blog              string `json:"blog" example:"string"`
	Collaborators     int    `json:"collaborators" example:"0"`
	Company           string `json:"company" example:"string"`
	CreatedAt         string `json:"created_at" example:"string"`
	DiskUsage         int    `json:"disk_usage" example:"0"`
	Email             string `json:"email" example:"string"`
	Followers         int    `json:"followers" example:"0"`
	Following         int    `json:"following" example:"0"`
	GravatarID        string `json:"gravatar_id" example:"string"`
	HTMLURL           string `json:"html_url" example:"string"`
	Hireable          int    `json:"hireable" example:"True"`
	ID                int    `json:"id" example:"0"`
	Location          string `json:"location" example:"string"`
	Login             string `json:"login" example:"string"`
	Name              string `json:"name" example:"string"`
	OwnedPrivateRepos int    `json:"owned_private_repos" example:"0"`
	Plan              Plan   `json:"plan"`
	PrivateGists      int    `json:"private_gists" example:"0"`
	PublicGists       int    `json:"public_gists" example:"0"`
	PublicRepos       int    `json:"public_repos" example:"0"`
	TotalPrivateRepos int    `json:"total_private_repos" example:"0"`
	Type              string `json:"type" example:"string"`
	URL               string `json:"url" example:"string"`
}

// Plan : auto generated JSON container
type Plan struct {
	Collaborators int    `json:"collaborators" example:"0"`
	Name          string `json:"name" example:"string"`
	PrivateRepos  int    `json:"private_repos" example:"0"`
	Space         int    `json:"space" example:"0"`
}
