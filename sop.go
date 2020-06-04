package sopsearch

import "time"

//
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
