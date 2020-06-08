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
}

//the map will have SOP name as key and JSON object as value in string
func (s Sop) toJSON() (map[string]string, error) {
	var jmap map[string]string
	obj, err := json.Marshal(s)
	if err != nil {
		return map[string]string{}, err
	}
	jmap[s.Name] = string(obj)
	//shouldn't i be doing this for all SOPs?
	return jmap, nil
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
