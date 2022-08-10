package main

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"rename-feature-tests/flags"
	"strings"
	"sync"
)

//const absPath = "/Users/ekrivenko/work-repos/acceptance-test/web-js-acceptance-test/test/features/selfcare"
//
////const file = "/Users/ekrivenko/work-repos/acceptance-test/api-acceptance-test/src/test/resources/features/selfcare/pz/els/add/AddEls.feature"
//const stream = "selfcare/"
//const prefixTag = "@autoTestExternalId-web"

// -path - absolute path for your stream
// -tag - prefix tag

var stream string
var prefixTag string

var wg sync.WaitGroup

func getDottedPath(path, stream string) string {
	index := strings.Index(path, stream)

	path = strings.TrimSuffix(path, ".feature")

	l := len([]rune(stream))

	path = path[index+l:]

	pathArray := strings.Split(path, string(os.PathSeparator))
	return stream + strings.Join(pathArray, ".")
}

// Prepare line for concat result tag
func cutLine(line, prefix string) string {
	index := strings.Index(line, prefix)
	if index == -1 {
		// remove \n after tags
		return strings.Replace(line, "\n", "", 1)
	}
	// remove prefix tag and space after last tag before prefix
	return line[:index-1]
}

func workWithFile(path, stream string) {
	input, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	dottedPath := getDottedPath(path, stream)
	lines := strings.Split(string(input), "\n")
	var aTestCounter int
	for i, l := range lines {

		isLineContainsATestTag := strings.Contains(l, "@atest")

		if isLineContainsATestTag {
			// counter for test cases for last symbol in result tag
			aTestCounter++

			resultTag := fmt.Sprintf(" %s.%s.%d", prefixTag, dottedPath, aTestCounter)
			lines[i] = cutLine(l, prefixTag) + resultTag
		}
	}

	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile(path, []byte(output), 0677)
	if err != nil {
		panic(err)
	}
	wg.Done()
}

func main() {
	f := flags.GetFlags()
	absPath := *f.AbsPath
	prefixTag = *f.TagPrefix
	ex := *f.ExcludedFolderName
	targetFolder := *f.TargetFolderName

	if absPath == "" || prefixTag == "" {
		log.Fatalln("Передайте --path и --tag")
	}

	strs := strings.Split(absPath, string(os.PathSeparator))

	if targetFolder != "" {
		absPath = path.Join(absPath, targetFolder)
		if _, err := os.Stat(absPath); err != nil {
			if os.IsNotExist(err) {
				log.Fatal("Папки не существует")
			}
			log.Fatalln(err)
		}
	}

	stream = strs[len(strs)-1]
	filepath.Walk(absPath, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		if ex != "" && strings.Contains(path, ex) {
			return nil
		}

		if strings.HasSuffix(info.Name(), ".feature") {
			wg.Add(1)
			go workWithFile(path, stream)
		}
		return nil
	})
	wg.Wait()
}
