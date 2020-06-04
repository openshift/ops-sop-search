package main

import (
	"fmt"

	sopsearch "github.com/lareplog/ops-sop-search"
)

func main() {

	files, err := sopsearch.ScanForMD("/home/laurenreplogle/ops-sop/v4/troubleshoot")
	if err != nil {
		fmt.Println(err)
	}
	n := 1
	for _, f := range files {
		fmt.Println(n)
		fmt.Println(f)
		n++
	}

	test := sopsearch.MDFile{"/test/file.md", "file.md", "hi"}
	test.ToSOP()
	fmt.Println(test)
}
