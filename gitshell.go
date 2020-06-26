package sopsearch

import (
	"os/exec"
	"strings"
	"time"
)

type Dates struct {
	Oldest time.Time
	Newest time.Time
}

func GitClone(repo string) ([]byte, error) {

	cmd := exec.Command("git", "clone", repo)
	return cmd.Output()
}

func GitPull() ([]byte, error) {
	//in order to do this, need to be in ops-sop no matter what
	cmd := exec.Command("git", "pull", "origin")
	return cmd.Output()
}

func GitLog(filepath string) ([]byte, error) {

	cmd := exec.Command("git", "log", "--pretty=%an%x09%ai", filepath)
	return cmd.Output()
}

func GetAuthorsAndDates(data []byte) ([]string, Dates, error) {

	var auths []string
	var date Dates
	format := "2006-01-02 15:04:05 -0700"
	list := strings.Split(string(data), "\n")
	list = list[:len(list)-1]

	for _, line := range list {
		section := strings.Split(line, "\t")
		if isUnique(auths, section[0]) {
			auths = append(auths, section[0])
		}
	}

	//last updated date will be the first item in list
	str := strings.Split(list[0], "\t")
	updated, err := time.Parse(format, str[1])
	if err != nil {
		return []string{}, Dates{}, nil
	}

	//creation date will be the last item in list - split string grab section [1] parse into date!!
	str = strings.Split(list[len(list)-1], "\t")
	creation, err := time.Parse(format, str[1])
	if err != nil {
		return []string{}, Dates{}, nil
	}

	date.Oldest = creation
	date.Newest = updated
	return auths, date, nil
}

func isUnique(list []string, pot string) bool {
	for _, str := range list {
		if str == pot {
			return false
		}
	}
	return true
}
