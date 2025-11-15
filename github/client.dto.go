package github

//go:generate go tool shoot new -getset -json -file=$GOFILE

type Org struct {
	id                 int64
	node_id            string
	login              string
	url                string
	repos_url          string
	events_url         string
	hooks_url          string
	issues_url         string
	members_url        string
	public_members_url string
	avatar_url         string
	description        string
}
