package github

import (
	"time"
)

/* structure
AuthenticatedUser
	Plan
*/
// AuthenticatedUser : auto generated JSON container
type AuthenticatedUser struct {
	AvatarURL         string    `json:"avatar_url" example:"https://github.com/images/error/octocat_happy.gif"`
	Bio               string    `json:"bio" example:"There once was..."`
	Blog              string    `json:"blog" example:"https://github.com/blog"`
	Collaborators     int       `json:"collaborators" example:"8"`
	Company           string    `json:"company" example:"GitHub"`
	CreatedAt         time.Time `json:"created_at" example:"2008-01-14T04:33:35Z"`
	DiskUsage         int       `json:"disk_usage" example:"10000"`
	Email             string    `json:"email" example:"octocat@github.com"`
	EventsURL         string    `json:"events_url" example:"https://api.github.com/users/octocat/events{/privacy}"`
	Followers         int       `json:"followers" example:"20"`
	FollowersURL      string    `json:"followers_url" example:"https://api.github.com/users/octocat/followers"`
	Following         int       `json:"following" example:"0"`
	FollowingURL      string    `json:"following_url" example:"https://api.github.com/users/octocat/following{/other_user}"`
	GistsURL          string    `json:"gists_url" example:"https://api.github.com/users/octocat/gists{/gist_id}"`
	GravatarID        string    `json:"gravatar_id" example:""`
	HTMLURL           string    `json:"html_url" example:"https://github.com/octocat"`
	Hireable          bool      `json:"hireable" example:"False"`
	ID                int       `json:"id" example:"1"`
	Location          string    `json:"location" example:"San Francisco"`
	Login             string    `json:"login" example:"octocat"`
	Name              string    `json:"name" example:"monalisa octocat"`
	OrganizationsURL  string    `json:"organizations_url" example:"https://api.github.com/users/octocat/orgs"`
	OwnedPrivateRepos int       `json:"owned_private_repos" example:"100"`
	Plan              Plan      `json:"plan"`
	PrivateGists      int       `json:"private_gists" example:"81"`
	PublicGists       int       `json:"public_gists" example:"1"`
	PublicRepos       int       `json:"public_repos" example:"2"`
	ReceivedEventsURL string    `json:"received_events_url" example:"https://api.github.com/users/octocat/received_events"`
	ReposURL          string    `json:"repos_url" example:"https://api.github.com/users/octocat/repos"`
	SiteAdmin         bool      `json:"site_admin" example:"False"`
	StarredURL        string    `json:"starred_url" example:"https://api.github.com/users/octocat/starred{/owner}{/repo}"`
	SubscriptionsURL  string    `json:"subscriptions_url" example:"https://api.github.com/users/octocat/subscriptions"`
	TotalPrivateRepos int       `json:"total_private_repos" example:"100"`
	Type              string    `json:"type" example:"User"`
	URL               string    `json:"url" example:"https://api.github.com/users/octocat"`
	UpdatedAt         time.Time `json:"updated_at" example:"2008-01-14T04:33:35Z"`
}

// Plan : auto generated JSON container
type Plan struct {
	Collaborators int    `json:"collaborators" example:"0"`
	Name          string `json:"name" example:"Medium"`
	PrivateRepos  int    `json:"private_repos" example:"20"`
	Space         int    `json:"space" example:"400"`
}
