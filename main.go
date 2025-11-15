package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"shoot-github-client/github"
	"time"

	github78 "github.com/google/go-github/v78/github"
	"github.com/lopolopen/shoot"
	"github.com/lopolopen/shoot/middleware"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Token string `yaml:"token"`
}

func main() {
	data, err := os.ReadFile("etc/app.yaml")
	// data, err := os.ReadFile("etc/private.yaml")
	if err != nil {
		panic(err)
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		panic(err)
	}

	useShootExaple(cfg)

	// useGoogleClientExample(cfg)
}

func useGoogleClientExample(cfg Config) {
	c := github78.NewClient(nil).WithAuthToken(cfg.Token)
	orgs, resp, err := c.Organizations.List(context.Background(), "", nil)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(orgs)
	fmt.Println(*resp)
}

func useShootExaple(cfg Config) {
	c := shoot.NewRest[github.Client](
		shoot.BaseURL("https://api.github.com"),
		shoot.Timeout(3000*time.Millisecond),
		shoot.EnableLogging(true),
		shoot.Use(func(next http.RoundTripper) http.RoundTripper {
			return middleware.RoundTripper(func(req *http.Request) (*http.Response, error) {
				req.Header.Add("Authorization", "token "+cfg.Token)
				return next.RoundTrip(req)
			})
		}),
	)
	orgs, resp, err := c.ListOrgsForUser(context.Background(), nil, nil)
	if err != nil {
		fmt.Println(err)
	}

	if resp != nil {
		fmt.Println(resp.StatusCode)
		fmt.Println(resp.Header.Get("etag"))
	}

	for _, org := range orgs {
		fmt.Println(org.Id())
		fmt.Println(org.Login())
		fmt.Println(org.Url())
	}

	repos, resp, err := c.ListReposForOrg(context.Background(), "lopolopen", nil, nil, nil)
	if err != nil {
		fmt.Println(err)
	}

	if resp != nil {
		fmt.Println(resp.StatusCode)
	}

	for _, repo := range repos {
		fmt.Println(repo.Id())
		fmt.Println(repo.Name())
		fmt.Println(repo.FullName())
		fmt.Println(repo.Private())
		if repo.Owner() != nil {
			fmt.Println("owner: ", repo.Owner().Login())
			fmt.Println("type: ", repo.Owner().Typ())
		}
		fmt.Println(repo.HtmlUrl())
		fmt.Println(repo.Description())
	}

	projs, resp, err := c.ListProjectsForUser(context.Background(), "redochenzhen", nil, nil, nil)
	if err != nil {
		fmt.Println(err)
	}

	if resp != nil {
		fmt.Println(resp.StatusCode)
	}

	for _, proj := range projs {
		fmt.Println(proj.Id())
		fmt.Println(proj.Title())
		fmt.Println(proj.Description())
		if proj.Owner() != nil {
			fmt.Println(proj.Owner().Login())
		}
		if proj.Creator() != nil {
			fmt.Println(proj.Creator().Login())
		}
	}
}
