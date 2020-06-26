package main

import (
	"log"
	"os"

	sopsearch "github.com/lareplog/ops-sop-search"
)

func init() {

	os.Mkdir("/home/laurenreplogle/test", os.FileMode(0777))
	os.Chdir("/home/laurenreplogle/test")

	_, err := sopsearch.GitClone("git@github.com:openshift/ops-sop.git")
	if err != nil {
		log.Panicln(err)
	}

	os.Chdir("ops-sop") //need to be in the directory first!
	sopsearch.Run(true) //automatically true so it will do initial index!
}

func main() {

	sopsearch.Routine()
	//time.Sleep(11 * time.Minute)
}
