package sopsearch

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type MDFile struct {
	Path         string
	Name         string
	Content      string
	Author       []string
	CreationDate time.Time
	LastUpdated  time.Time
	Tags         []string
}

type ADFile struct {
	Path         string
	Name         string
	Content      string
	Author       []string
	CreationDate time.Time
	LastUpdated  time.Time
	Tags         []string
}

// ScanForFiles scans for files ending in ".md", then creates an MDFile obj for each
// file and then adds it to the slice. the slice is returned at the end.
func ScanForFiles(path string) ([]MDFile, []ADFile, error) {

	var mfiles []MDFile

	var afiles []ADFile

	mdmatches, err := FindFiles(path, "*.md")
	if err != nil {
		return []MDFile{}, []ADFile{}, err
	}

	admatches, err := FindFiles(path, "*.asciidoc")
	if err != nil {
		return []MDFile{}, []ADFile{}, err
	}

	for _, f := range mdmatches {
		path := f
		name := strings.TrimSuffix(filepath.Base(f), filepath.Ext(filepath.Base(f)))
		content, err := ioutil.ReadFile(f)
		if err != nil {
			return []MDFile{}, []ADFile{}, err
		}
		cont := string(content)
		git, err := GitLog(f)
		if err != nil {
			return []MDFile{}, []ADFile{}, err
		}
		auth, dat, err := GetAuthorsAndDates(git)
		if err != nil {
			return []MDFile{}, []ADFile{}, err
		}
		tags := strings.Split(filepath.Dir(f), "/")
		tags = append(tags, "markdown")
		mfiles = append(mfiles, MDFile{path, name, cont, auth, dat.Oldest, dat.Newest, tags})
	}

	for _, f := range admatches {
		path := f
		name := strings.TrimSuffix(filepath.Base(f), filepath.Ext(filepath.Base(f)))
		content, err := ioutil.ReadFile(f)
		if err != nil {
			return []MDFile{}, []ADFile{}, err
		}
		cont := string(content)
		git, err := GitLog(f)
		if err != nil {
			return []MDFile{}, []ADFile{}, err
		}
		auth, dat, err := GetAuthorsAndDates(git)
		if err != nil {
			return []MDFile{}, []ADFile{}, err
		}
		tags := strings.Split(filepath.Dir(f), "/")
		tags = append(tags, "asciidoc")
		afiles = append(afiles, ADFile{path, name, cont, auth, dat.Oldest, dat.Newest, tags})
	}

	return mfiles, afiles, nil
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

func ToBulkSOP(mf []MDFile, af []ADFile) ([]Sop, error) {

	var sop Sop
	var slice []Sop

	for _, x := range mf {
		sop.Path = x.Path
		sop.Name = x.Name
		sop.Content = x.Content
		sop.Author = x.Author
		sop.CreationDate = x.CreationDate
		sop.LastUpdated = x.LastUpdated
		sop.Tags = x.Tags
		slice = append(slice, sop)
	}

	for _, y := range af {
		sop.Path = y.Path
		sop.Name = y.Name
		sop.Content = y.Content
		sop.Author = y.Author
		sop.CreationDate = y.CreationDate
		sop.LastUpdated = y.LastUpdated
		sop.Tags = y.Tags
		slice = append(slice, sop)
	}

	return slice, nil
}
