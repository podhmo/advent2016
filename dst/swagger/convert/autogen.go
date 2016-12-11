package convert

import (
	def "github.com/podhmo/advent2016/dst/swagger/gen/def"
	autogen "github.com/podhmo/advent2016/src/github"
)

// AuthenticatedUserToUser :
func AuthenticatedUserToUser(from *autogen.AuthenticatedUser) *def.User {
	to := &def.User{}
	to.AvatarURL = from.AvatarURL
	to.Bio = from.Bio
	to.Blog = from.Blog
	tmp1 := (int64)(from.Collaborators)
	to.Collaborators = tmp1
	to.Company = from.Company
	to.CreatedAt = from.CreatedAt
	tmp2 := (int64)(from.DiskUsage)
	to.DiskUsage = tmp2
	to.Email = from.Email
	tmp3 := (int64)(from.Followers)
	to.Followers = tmp3
	tmp4 := (int64)(from.Following)
	to.Following = tmp4
	to.GravatarID = from.GravatarID
	to.Hireable = from.Hireable
	to.HTMLURL = from.HTMLURL
	tmp5 := (int64)(from.ID)
	to.ID = tmp5
	to.Location = from.Location
	to.Login = from.Login
	to.Name = from.Name
	tmp6 := (int64)(from.OwnedPrivateRepos)
	to.OwnedPrivateRepos = tmp6
	to.Plan = PlanToUserPlan(&(from.Plan))
	tmp7 := (int64)(from.PrivateGists)
	to.PrivateGists = tmp7
	tmp8 := (int64)(from.PublicGists)
	to.PublicGists = tmp8
	tmp9 := (int64)(from.PublicRepos)
	to.PublicRepos = tmp9
	tmp10 := (int64)(from.TotalPrivateRepos)
	to.TotalPrivateRepos = tmp10
	to.Type = from.Type
	to.URL = from.URL
	return to
}

// PlanToUserPlan :
func PlanToUserPlan(from *autogen.Plan) *def.UserPlan {
	to := &def.UserPlan{}
	tmp11 := (int64)(from.Collaborators)
	to.Collaborators = tmp11
	to.Name = from.Name
	tmp12 := (int64)(from.PrivateRepos)
	to.PrivateRepos = tmp12
	tmp13 := (int64)(from.Space)
	to.Space = tmp13
	return to
}
