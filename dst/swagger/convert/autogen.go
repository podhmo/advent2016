package convert

import (
	def "github.com/podhmo/advent2016/dst/swagger/gen/def"
	github "github.com/podhmo/advent2016/src/github"
)

// AuthenticatedUserToUser :
func AuthenticatedUserToUser(from *github.AuthenticatedUser) *def.User {
	to := &def.User{}
	tmp1 := (string)(from.AvatarURL)
	to.AvatarURL = tmp1
	to.Bio = from.Bio
	tmp2 := (string)(from.Blog)
	to.Blog = tmp2
	tmp3 := (int64)(from.Collaborators)
	to.Collaborators = tmp3
	to.Company = from.Company
	to.CreatedAt = TimeToString(from.CreatedAt)
	tmp4 := (int64)(from.DiskUsage)
	to.DiskUsage = tmp4
	to.Email = from.Email
	tmp5 := (int64)(from.Followers)
	to.Followers = tmp5
	tmp6 := (int64)(from.Following)
	to.Following = tmp6
	to.GravatarID = from.GravatarID
	to.Hireable = from.Hireable
	tmp7 := (string)(from.HTMLURL)
	to.HTMLURL = tmp7
	tmp8 := (int64)(from.ID)
	to.ID = tmp8
	to.Location = from.Location
	to.Login = from.Login
	to.Name = from.Name
	tmp9 := (int64)(from.OwnedPrivateRepos)
	to.OwnedPrivateRepos = tmp9
	to.Plan = PlanToUserPlan(&(from.Plan))
	tmp10 := (int64)(from.PrivateGists)
	to.PrivateGists = tmp10
	tmp11 := (int64)(from.PublicGists)
	to.PublicGists = tmp11
	tmp12 := (int64)(from.PublicRepos)
	to.PublicRepos = tmp12
	tmp13 := (int64)(from.TotalPrivateRepos)
	to.TotalPrivateRepos = tmp13
	to.Type = from.Type
	tmp14 := (string)(from.URL)
	to.URL = tmp14
	return to
}

// PlanToUserPlan :
func PlanToUserPlan(from *github.Plan) *def.UserPlan {
	to := &def.UserPlan{}
	tmp15 := (int64)(from.Collaborators)
	to.Collaborators = tmp15
	to.Name = from.Name
	tmp16 := (int64)(from.PrivateRepos)
	to.PrivateRepos = tmp16
	tmp17 := (int64)(from.Space)
	to.Space = tmp17
	return to
}
