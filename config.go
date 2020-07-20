package sopsearch

import (
	"context"
	"strconv"

	v1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type Config struct {
	Time       int
	ElasticURL string
	RepoURL    string
	RepoName   string
	GitScript  string
}

func GetConfig(kubecli client.Client) (Config, error) {
	var confmap v1.ConfigMap

	confmap, err := getConfigMap(context.TODO(), kubecli)
	if err != nil {
		return Config{}, err
	}

	config := Config{
		ElasticURL: confmap.Data["elastic"],
		RepoURL:    confmap.Data["repourl"],
		RepoName:   confmap.Data["reponame"],
		GitScript:  confmap.Data["gitscript"],
	}
	config.Time, err = strconv.Atoi(confmap.Data["time"])
	if err != nil {
		return Config{}, err
	}
	return config, nil
}
