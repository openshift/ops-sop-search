package sopsearch

import (
	"os/exec"
	"strings"
	"time"
)

// Dates contains the creation date of the SOP document and the date it was last updated
type Dates struct {
	Oldest time.Time
	Newest time.Time
}

// GitClone performs a git clone of repo contained in the repo string. The script string
//containts the location of the shell script that performs the git clone.
func GitClone(script string, repo string) error {

	cmd := exec.Command("sh", script, "-c", repo)
	return cmd.Run()
}

// GitPull performs a git pull of the repo. The script string contains the location
//of the shell script that performs the git pull. It will then return the output of
//the git pull to determine if the repo was actually updated or not.
func GitPull(script string) ([]byte, error) {

	cmd := exec.Command("sh", script, "-p")
	return cmd.Output()
}

// GitLog performs a git log of a file specified by the filepath string. It then returns
//the output of the command.
func GitLog(filepath string) ([]byte, error) {

	cmd := exec.Command("git", "log", "--pretty=%an%x09%ai", filepath)
	return cmd.Output()
}

// GetAuthorsAndDates takes the slice of bytes (data) which contains the git log of a file
//and then goes through and finds all the unique authors of the file as well as the creation
//date and date the file was last updated and then returns that information.
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

// isUnique checks the list of authors against a potential author to see if they
//are already in the list. If they aren't it will return true. If it is in the list
//it will return false.
func isUnique(list []string, pot string) bool {
	for _, str := range list {
		if str == pot {
			return false
		}
	}
	return true
}
