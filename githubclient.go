package sopsearch

import (
	"context"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/google/go-github/v32/github"
	"golang.org/x/oauth2"
)

type GithubClient struct {
	gclient *github.Client
}

func NewGithubClient(token string) GithubClient {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)
	return GithubClient{client}
}

func (gc GithubClient) ListOrganizations(username string) error {
	orgs, _, err := gc.gclient.Organizations.List(context.Background(), username, nil)
	if err != nil {
		return err
	}
	for _, org := range orgs {
		fmt.Println(org)
	}
	return nil
}

func (gc GithubClient) ListRepos(username string) error {
	//orgs, _, err := gc.gclient.Organizations.List(context.Background(), username, nil)
	ctx := context.Background()
	repos, _, err := gc.gclient.Repositories.List(ctx, username, nil)
	if err != nil {
		return err
	}
	//repo, res, err := gc.gclient.Repositories.Get(ctx, username, "https://github.com/lareplog/ops-sop")
	if err != nil {
		return err
	}
	for _, org := range repos {
		fmt.Println(org)
	}
	return nil
}

func (gc GithubClient) getRepo(username string) (*github.Repository, error) {

	ctx := context.Background()
	repo, _, err := gc.gclient.Repositories.Get(ctx, username, "ops-sop")
	if err != nil {
		return nil, err
	}

	return repo, nil
}
func (gc GithubClient) NeedsUpdating(username string, org string) (bool, error) {
	base, err := gc.getRepo(username)
	if err != nil {
		return false, err
	}
	ops, err := gc.getRepo(org)
	if err != nil {
		return false, err
	}

	if base.GetUpdatedAt() == ops.GetUpdatedAt() {
		return false, nil
	}
	return true, nil

	//since one repo will never change (openshift/ops-sop), can just get it from here
	//compare timestamp of openshift/sop-ops to timestamp of given repo, if not same
	//  then you need to call in git pull for given repo and get it up to date?
	//    is that the only way to do this??? is there an easier way???!!!
}

func GetToken(path string) (string, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	token := strings.TrimSpace(string(content))
	token = strings.TrimSuffix(token, "\n")
	return token, nil
}
