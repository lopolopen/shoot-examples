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
}
