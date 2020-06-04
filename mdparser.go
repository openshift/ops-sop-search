package sopsearch

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

type MDFile struct {
	Path    string
	Name    string
	Content string
}

//Scans for files ending in ".md", then creates an MDFile obj for each
// file and then adds it to the slice. the slice is returned at the end.
func ScanForMD(path string) ([]MDFile, error) {

	//variable for slice of MDFile obj
	var files []MDFile

	//method which finds all MDFiles in given path
	matches, err := FindFiles(path, "*.md")
	if err != nil {
		return []MDFile{}, err
	}

	//creates MDFile objectts and adds them to slice
	for _, f := range matches {
		path := f
		name := filepath.Base(f)
		content, err := ioutil.ReadFile(f)
		if err != nil {
			return []MDFile{}, err
		}
		cont := string(content)
		files = append(files, MDFile{path, name, cont})
	}

	return files, nil
}

//searches the given root and using the pattern, it finds all the files that end
// in that particular pattern.
func FindFiles(root, pattern string) ([]string, error) {
	var matches []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if matched, err := filepath.Match(pattern, filepath.Base(path)); err != nil {
			return err
		} else if matched {
			matches = append(matches, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return matches, nil
}

//func thatll convert an md file to an sop struct with info known so far
// aka name, path, content
func (mf MDFile) ToSOP() (Sop, error) {
	var sop Sop
	sop.Path = mf.Path
	sop.Name = mf.Name
	sop.Content = mf.Content
	return sop, nil
}
