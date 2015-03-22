package main

import (
	"bufio"
	"flag"
	"fmt"
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

func newFilename(templateDir string, path string, mappings map[string]string) string {
	newPath := strings.Replace(path, templateDir, ".", -1)
	return strings.Replace(newPath, "NAME", mappings["NAME"], -1)
}

func main() {
	var templateDir string

	flag.BoolVar(&verbose, "verbose", false, "Be verbose")
	flag.StringVar(&templateDir, "templatedir", "~/.goose",
		"Directory where templates are stored")
	flag.Parse()

	program := path.Base(os.Args[0])
	args := flag.Args()
	log.Println(args)
	if len(args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %v [--verbose] <template> <name>\n", program)
		os.Exit(1)
	}
	template := args[0]
	name := args[1]

	log.Println("program:", program)
	log.Println("template:", template)
	log.Println("name:", name)
	log.Println("verbose:", verbose)
	log.Println("templateDir:", templateDir)
	dir := templateDir + "/" + template
	generate(dir, map[string]string{"NAME": name})
}
