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

type User struct {
	login      string
	id         int64
	node_id    string
	avatar_url string
	typ        string `json:"type"`
}

type ProjectV2 struct {
	id          string
	node_id     string
	title       string
	description string
	owner       *User
	creator     *User
}

type Repo struct {
	id          int64
	node_id     string
	name        string
	full_name   string
	private     bool
	owner       *User
	html_url    string
	description string
}
