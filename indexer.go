package sopsearch

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
