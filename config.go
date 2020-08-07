package sopsearch

import (
	"context"
	"strconv"

	v1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// Config struct holds all the important values used across the program.
type Config struct {
	Time       int
	ElasticURL string
	GitURL     string
	RepoName   string
	GitScript  string
	RepoURL    string
}

// GetConfig uses a Kubernetes client and then gets the configmap object. Then it puts
//the data from that configmap object into the Config object which can then be used and
//accessed across the whole program.
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
		GitURL:     confmap.Data["giturl"],
	}
	config.Time, err = strconv.Atoi(confmap.Data["time"])
	if err != nil {
		return Config{}, err
	}
	return config, nil
}
