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
	//possibly add []string for tags?
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

//Scans for files ending in ".md", then creates an MDFile obj for each
// file and then adds it to the slice. the slice is returned at the end.
func ScanForFiles(path string) ([]MDFile, []ADFile, error) {

	//variable for slice of MDFile obj
	var mfiles []MDFile

	//variable for slice of ADFile obj
	var afiles []ADFile

	//method which finds all MDFiles in given path
	mdmatches, err := FindFiles(path, "*.md")
	if err != nil {
		return []MDFile{}, []ADFile{}, err
	}

	//method which finds all ADFiles in given path
	admatches, err := FindFiles(path, "*.asciidoc")
	if err != nil {
		return []MDFile{}, []ADFile{}, err
	}

	//creates MDFile objects and adds them to slice
	for _, f := range mdmatches {
		path := f
		name := strings.TrimSuffix(filepath.Base(f), filepath.Ext(filepath.Base(f)))
		content, err := ioutil.ReadFile(f) //before or after this: where divinding file into sections would be?
		if err != nil {
			return []MDFile{}, []ADFile{}, err
		}
		cont := string(content)
		//getting author names and dates!
		git, err := GitLog(f) //CHANGED
		if err != nil {
			return []MDFile{}, []ADFile{}, err
		}
		//old, new, err := AddData(git)
		auth, dat, err := GetAuthorsAndDates(git)
		if err != nil {
			return []MDFile{}, []ADFile{}, err
		}
		// var authors []string
		// authors = append(authors, old.Name, new.Name)
		// first := old.Date
		// last := new.Date
		//getting tags
		tags := strings.Split(filepath.Dir(f), "/")
		tags = append(tags, "markdown")
		mfiles = append(mfiles, MDFile{path, name, cont, auth, dat.Oldest, dat.Newest, tags})
	}

	//creates ADFile objects and adds them to slice
	for _, f := range admatches {
		path := f
		name := strings.TrimSuffix(filepath.Base(f), filepath.Ext(filepath.Base(f)))
		content, err := ioutil.ReadFile(f)
		if err != nil {
			return []MDFile{}, []ADFile{}, err
		}
		cont := string(content)
		git, err := GitLog(f) //CHANGED
		if err != nil {
			return []MDFile{}, []ADFile{}, err
		}
		//old, new, err := AddData(git)
		auth, dat, err := GetAuthorsAndDates(git)
		if err != nil {
			return []MDFile{}, []ADFile{}, err
		}
		// var authors []string
		// authors = append(authors, old.Name, new.Name)
		// first := old.Date
		// last := new.Date
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

//func thatll convert an md file to an sop struct with info known so far
// aka name, path, content
func (mf MDFile) ToSOP() (Sop, error) {
	var sop Sop
	sop.Path = mf.Path
	sop.Name = mf.Name
	sop.Content = mf.Content
	sop.Tags = mf.Tags
	return sop, nil
}

func ToBulkSOP(mf []MDFile, af []ADFile) ([]Sop, error) {

	//does same as above but for a large amount of mdFiles
	var sop Sop
	var slice []Sop

	//creating SOP objs from MDFile obj
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

	//creating SOP objs from ADFile obj
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
