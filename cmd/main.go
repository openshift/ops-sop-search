package main

import (
	"fmt"
	"time"

	sopsearch "github.com/lareplog/ops-sop-search"
)

func main() {
	start := time.Now()
	ec, err := sopsearch.NewElasticClient(
		[]string{"http://localhost:9200"},
		"user",
		"pass",
	)
	if err != nil {
		fmt.Println(err)
	}
	//set up everything so IndexSOP goes successfully lol
	md, ad, err := sopsearch.ScanForFiles("/home/laurenreplogle/ops-sop")
	if err != nil {
		fmt.Println(err)
	}
	sop, err := sopsearch.ToBulkSOP(md, ad)
	if err != nil {
		fmt.Println(err)
	}
	jmap, err := sopsearch.ToBulkJSON(sop)
	if err != nil {
		fmt.Println(err)
	}
	err = sopsearch.IndexSOP(&ec, jmap)
	if err != nil {
		fmt.Println(err)
	}
	elapsed := time.Since(start)
	fmt.Println("success!")
	fmt.Println(elapsed)
}
