package github

import (
	"context"
	"net/http"

	"github.com/lopolopen/shoot"
)

//go:generate go tool shoot rest -type=Client

type Client interface {
	//shoot: headers=
	// {Accept:application/vnd.github+json}
	// {X-GitHub-Api-Version:2022-11-28}
	shoot.RestClient[Client]

	// ListOrgsForUser lists organizations for the authenticated user.
	// GitHub API docs: https://docs.github.com/en/rest/orgs/orgs#list-organizations-for-the-authenticated-user
	//shoot: Get("/user/orgs")
	ListOrgsForUser(ctx context.Context, per_page *int, page *int) ([]*Org, *http.Response, error)

	// ListProjectsForUser lists projects for a user.
	// GitHub API docs: https://docs.github.com/en/rest/projects/projects#list-projects-for-user
	//shoot: Get("/users/{username}/projectsV2")
	ListProjectsForUser(ctx context.Context, username string, beforce *string, after *string, per_page *int) ([]*ProjectV2, *http.Response, error)

	// ListReposForOrg lists repositories for an organization.
	// GitHub API docs: https://docs.github.com/en/rest/repos/repos#list-organization-repositories
	//shoot: Get("/orgs/{org}/repos")
	//shoot: alias={typ:type}
	ListReposForOrg(ctx context.Context, org string, typ *string, per_page *int, page *int) ([]*Repo, *http.Response, error)
}
