package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
	"text/template"
)

var verbose bool

var funcMap = map[string]interface{}{
	"snakecase":    SnakeCase,
	"camelcase":    CamelCase,
	"dromedarcase": DromedarCase,
	"dasherized":   Dasherized,
}

func generate(templateDir string, mappings map[string]string) {
	copyFile := func(filename string, info os.FileInfo, err error) error {
		newPath := newFilename(templateDir, filename, mappings)
		if info.IsDir() {
			log.Printf("Creating dir %v\n", newPath)
			os.MkdirAll(newPath, 0700)
		} else {
			tmpl := template.Must(template.New(path.Base(filename)).Funcs(funcMap).ParseFiles(filename))
			f, err := os.Create(newPath)
			if err != nil {
				panic(err)
			}
			writer := bufio.NewWriter(f)
			err = tmpl.Execute(writer, mappings)
			if err != nil {
				panic(err)
			}
			writer.Flush()
			f.Close()
		}
		return nil
	}
	filepath.Walk(templateDir, copyFile)
}

func newFilename(templateDir string, filename string, mappings map[string]string) string {
	newPath := strings.Replace(filename, templateDir, ".", -1)
	for k, v := range mappings {
		newPath = replace(newPath, k, v)
	}
	return newPath
}

func replace(name string, key string, value string) string {
	tmp := strings.Replace(name, key+".cc", CamelCase(value), -1)
	tmp = strings.Replace(tmp, key+".da", Dasherized(value), -1)
	tmp = strings.Replace(tmp, key+".dc", DromedarCase(value), -1)
	tmp = strings.Replace(tmp, key+".sc", SnakeCase(value), -1)
	return tmp
}

func main() {
	var templateDir string

	flag.BoolVar(&verbose, "verbose", false, "Be verbose")
	flag.StringVar(&templateDir, "templatedir", "~/.goose",
		"Directory where templates are stored")
	flag.Parse()
	log.SetPrefix("")
	if !verbose {
		log.SetOutput(ioutil.Discard)
	}

	program := path.Base(os.Args[0])
	args := flag.Args()
	log.Println(args)
	if len(args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %v [--verbose] <template> <name>\n", program)
		os.Exit(1)
	}
	template := args[0]
	name := args[1]

	log.Println("verbose:", verbose)
	log.Println("program:", program)
	log.Println("template:", template)
	log.Println("name:", name)
	log.Println("templateDir:", templateDir)
	dir := templateDir + "/" + template
	generate(dir, map[string]string{"NAME": name})
}
