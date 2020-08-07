package sopsearch

import (
	"encoding/json"
	"time"
)

// Sop struct contains all the information that will be in the index for the sop document
type Sop struct {
	Name         string    `json:"name"`
	Path         string    `json:"path"`
	Author       []string  `json:"author"`
	CreationDate time.Time `json:"creationDate"`
	LastUpdated  time.Time `json:"lastUpdated"`
	Commit       string    `json:"commit"`
	Content      string    `json:"content"`
	Tags         []string  `json:"tags"`
	Link         string    `json:"links"`
}

// ToBulkJSON takes all the Sop objects and puts them into the map that will be used
//when calling IndexSOP after encoding the Sop into JSON. The name of the SOP will be
//the key and the json encoded sop as the value for the key-value pair.
func ToBulkJSON(s []Sop) (map[string]string, error) {
	jmap := make(map[string]string)
	for _, f := range s {
		obj, err := json.Marshal(f)
		if err != nil {
			return map[string]string{}, nil
		}
		jmap[f.Name] = string(obj)
	}
	return jmap, nil
}
