package main

import (
	"log"
	"os"

	"github.com/jasonlvhit/gocron"
	sopsearch "github.com/lareplog/ops-sop-search"
)

func init() {

	log.Printf("Configuring Application...\n")

	cli, err := sopsearch.GetKubeClient()
	if err != nil {
		log.Printf("Failed to get kubernetes client.\n error: %s\n", err)
		os.Exit(1)
	}

	conf, err := sopsearch.GetConfig(cli)
	if err != nil {
		log.Printf("Failed to get config object.\n error: %s\n", err)
		os.Exit(1)
	}

	err = sopsearch.GitClone(conf.GitScript, conf.RepoURL)
	if err != nil {
		log.Printf("Failed to clone repo %s.\n error: %s\n", conf.RepoURL, err)
		os.Exit(1)
	}

	log.Printf("Cloned %s Repo!\n", conf.RepoName)

	err = os.Chdir(conf.RepoName)
	if err != nil {
		log.Printf("Couldn't find %s directory.\n error: %s\n", conf.RepoName, err)
		os.Exit(1)
	}

	log.Printf("Start Initial Indexing...\n")

	sopsearch.RunIndex(true, conf)
}

func main() {

	log.Printf("Starting Routine...\n")
	cli, err := sopsearch.GetKubeClient()
	if err != nil {
		log.Printf("Failed to get kubernetes client.\n error: %s\n", err)
		os.Exit(1)
	}

	config, err := sopsearch.GetConfig(cli)
	if err != nil {
		log.Printf("Failed to get config object.\n error: %s\n", err)
		os.Exit(1)
	}

	err = gocron.Every(uint64(config.Time)).Minute().Do(sopsearch.NeedReIndex, config)
	if err != nil {
		log.Printf("Error when indexing: %s\n", err)
	}
	<-gocron.Start()
}
