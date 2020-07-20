package sopsearch

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/pkg/errors"
)

type Indexer interface {
	CreateOrUpdateIndex(index, documentID, body string) error
}

func IndexSOP(indexer Indexer, sops map[string]string) error {
	for key, content := range sops {
		err := indexer.CreateOrUpdateIndex("sop", key, content)
		if err != nil {
			return err
		}
	}
	return nil
}

func RunIndex(index bool, config Config) error {

	if index {

		log.Printf("Indexing %s now...\n", config.RepoName)

		start := time.Now()

		ec, err := NewElasticClient(
			[]string{config.ElasticURL}, "", "")
		if err != nil {
			return err
		}

		path, err := os.Getwd()
		if err != nil {
			return err
		}

		md, ad, err := ScanForFiles(path)
		if err != nil {
			return err
		}

		sop, err := ToBulkSOP(md, ad)
		if err != nil {
			return err
		}

		jmap, err := ToBulkJSON(sop)
		if err != nil {
			return err
		}

		err = IndexSOP(&ec, jmap)
		if err != nil {
			return err
		}

		elapsed := time.Since(start)
		log.Printf("Indexing complete! Time elapsed: %v\n", elapsed)
	}
	return nil
}

func NeedReIndex(config Config) error {

	ret, err := GitPull(config.GitScript)
	if err != nil {
		msg := fmt.Sprintf("Could not pull repo %s\n error: %s\n", config.RepoName, err)
		return errors.Wrap(err, msg)
	}

	if string(ret) != string("Already up to date.\n") {
		return RunIndex(true, config)
	}

	log.Printf("Repo %s is up to date, no need to re-index.\n", config.RepoName)
	return RunIndex(false, config)
}
