package sopsearch

import "fmt"

type Indexer interface {
	CreateOrUpdateIndex(index, documentID, body string) error
}
type FakeElasticClient struct {
}

func (fec *FakeElasticClient) CreateOrUpdateIndex(index, documentID, body string) error {
	// This can do what ever i want its fake
	fmt.Println("fake!")
	return nil
}

//call func from mdparser that grabs mdfile, convert it into sop struct
//  convert sop struct into json slice of bytes ->string
//    for each string, create index
func indexMD(indexer Indexer) error {
	err := indexer.CreateOrUpdateIndex("hey", "12", "body")
	if err != nil {
		return err
	}
	return nil
}
func Whatever() error {
	fec := FakeElasticClient{}
	ec, err := NewElasticClient(
		[]string{"http://localhost:9200"},
		"user",
		"pass",
	)
	if err != nil {
		return err
	}
	err = indexMD(&fec)
	if err != nil {
		return err
	}
	err = indexMD(&ec)
	if err != nil {
		return err
	}
	return nil
}
