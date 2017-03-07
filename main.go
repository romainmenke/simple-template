package main

import (
	"bytes"
	"flag"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func main() {

	var err error

	source := flag.String("source", "./", "source directory")
	out := flag.String("out", "./", "output directory")
	flag.Parse()

	sourceDir := strings.TrimSuffix(*source, "/")
	outDir := strings.TrimSuffix(*out, "/") + "/"
	createIfMissing(outDir)

	t := template.New(sourceDir)

	templ := getTemplates(sourceDir)
	if len(templ) > 0 {
		t, err = t.ParseFiles(templ...)
		if err != nil {
			panic(err)
		}
	}

	exclude := flag.Args()

	files, err := ioutil.ReadDir(sourceDir)
	if err != nil {
		panic(err)
	}

FILE_ITERATOR:
	for _, f := range files {
		if !isFile(sourceDir + "/" + f.Name()) {
			continue FILE_ITERATOR
		}

		for _, exc := range exclude {
			if strings.Contains(f.Name(), exc) {
				continue FILE_ITERATOR
			}
		}

		if !strings.HasSuffix(f.Name(), ".html") {
			continue FILE_ITERATOR
		}

		var buf bytes.Buffer
		err = t.ExecuteTemplate(&buf, f.Name(), nil)
		if err != nil {
			panic(err)
		}

		writeFile(buf.Bytes(), f.Name(), *out)
	}
}

func readFile(name string) []byte {
	buf := bytes.NewBuffer(nil)
	file, err := os.Open(name)
	if err != nil {
		panic(err)
	}
	_, err = io.Copy(buf, file)
	if err != nil {
		panic(err)
	}
	file.Close()
	return buf.Bytes()
}

func writeFile(content []byte, fileName string, out string) {

	err := ioutil.WriteFile(out+"/"+fileName, content, 0644)
	if err != nil {
		panic(err)
	}

}

func isFile(path string) bool {
	fileInfo, err := os.Stat(path)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return !fileInfo.IsDir()
}

func createIfMissing(path string) {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		os.Mkdir(path, os.ModePerm)
	}
}
