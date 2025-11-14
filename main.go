package main

import (
	"fmt"
	"net/http"
	"os"
	"shoot-github-client/github"
	"time"

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

	githubC := shoot.NewRest[github.Client](
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
	orgs, ex, err := githubC.ListOrgsForUser(nil, nil)
	if err != nil {
		panic(err)
	}

	if ex != nil {
		fmt.Printf("exception: %+v\n", ex)
		return
	}

	for _, org := range orgs {
		fmt.Println(org.Id())
		fmt.Println(org.Login())
		fmt.Println(org.Url())
	}
}
