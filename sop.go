package sopsearch

import (
	"encoding/json"
	"time"
)

type Sop struct {
	Name         string    `json:"name"`
	Path         string    `json:"path"`
	Author       []string  `json:"author"`
	CreationDate time.Time `json:"creationDate"`
	LastUpdated  time.Time `json:"lastUpdated"`
	Commit       string    `json:"commit"`
	Content      string    `json:"content"`
	Tags         []string  `json:"tags"`
	Links        []string  `json:"links"`
}

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
