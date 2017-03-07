package main

import (
	"io/ioutil"
	"strings"
)

func getTemplates(pathElements ...string) []string {

	var templateFiles []string

	if len(pathElements) == 0 {
		pathElements = []string{".", "templates"}
	}

	path := strings.Join(pathElements[:], "/")

	files, _ := ioutil.ReadDir(path)
	for _, f := range files {
		if f.IsDir() {
			newPathElements := append(pathElements, f.Name())
			templateFiles = append(templateFiles, getTemplates(newPathElements...)...)
		} else if len(strings.Split(f.Name(), ".")) > 1 && strings.Split(f.Name(), ".")[1] == "html" {
			templateFiles = append(templateFiles, path+"/"+f.Name())
		}
	}
	return templateFiles
}
