package sopsearch

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jasonlvhit/gocron"
)

type Indexer interface {
	CreateOrUpdateIndex(index, documentID, body string) error
}

//call func from mdparser that grabs mdfile, convert it into sop struct
//  convert sop struct into json slice of bytes ->string
//    for each string, create index
func IndexSOP(indexer Indexer, sops map[string]string) error {
	for key, content := range sops {
		err := indexer.CreateOrUpdateIndex("sop", key, content)
		if err != nil {
			return err
		}
	}
	return nil
}

func Run(index bool) {

	//if it does need to be reindexed, reindex!
	if index {
		start := time.Now()
		ec, err := NewElasticClient(
			[]string{"http://localhost:9200"},
			"user",
			"pass",
		)
		if err != nil {
			fmt.Println(err)
		}
		//os.Chdir("/home/laurenreplogle/ops-sop")
		path, er := os.Getwd()
		if er != nil {
			log.Println(er)
		}
		md, ad, err := ScanForFiles(path)
		if err != nil {
			fmt.Println(err)
		}
		sop, err := ToBulkSOP(md, ad)
		if err != nil {
			fmt.Println(err)
		}
		jmap, err := ToBulkJSON(sop)
		if err != nil {
			fmt.Println(err)
		}
		err = IndexSOP(&ec, jmap)
		if err != nil {
			fmt.Println(err)
		}
		elapsed := time.Since(start)
		fmt.Println("success!")
		fmt.Println(elapsed)
	}
}

func needReIndex() {

	ret, err := GitPull()
	if err != nil {
		//if error occurs, stop! log.Panic?
		log.Panicln(err)
	}

	if string(ret) == string("Already up to date.\n") {
		Run(false)
	} else {
		Run(true)
	}
}

func Routine() {

	gocron.Every(2).Minute().Do(needReIndex)
	<-gocron.Start()
}
