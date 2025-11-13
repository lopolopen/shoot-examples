package github

import "github.com/lopolopen/shoot"

//go:generate go tool shoot rest -type=Client

type Client interface {
	//shoot: headers=
	// {Accept:application/vnd.github+json}
	// {X-GitHub-Api-Version:2022-11-28}
	shoot.RestClient[Client]

	//shoot: Get("/user/orgs")
	ListOrgsForUser(per_page int, page int) ([]Org, map[string]string, error)
}
