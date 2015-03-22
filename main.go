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

func generate(templateDir string, outputDir string, mappings map[string]string) {
	copyFile := func(filename string, info os.FileInfo, err error) error {
		newPath := newFilename(templateDir, outputDir, filename, mappings)
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
			log.Printf("Generating file %v\n", newPath)
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

func newFilename(templateDir string, outputDir string, filename string, mappings map[string]string) string {
	newPath := strings.Replace(filename, templateDir, outputDir, -1)
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
	tmp = strings.Replace(tmp, key, value, -1)
	return tmp
}

func main() {
	var templateDir string
	var outputDir string

	flag.BoolVar(&verbose, "verbose", false, "Be verbose")
	flag.StringVar(&templateDir, "templatedir", os.Getenv("HOME")+"/.goose",
		"Directory where templates are stored")
	flag.StringVar(&outputDir, "outputdir", "", "Output directory, default NAME")
	flag.Parse()
	program := path.Base(os.Args[0])
	log.SetFlags(0)
	if !verbose {
		log.SetOutput(ioutil.Discard)
	}

	args := flag.Args()
	log.Println(args)
	if len(args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %v [--templatedir <dir>] [--outputdir <dir>] [--verbose] <template> <name>\n", program)
		os.Exit(1)
	}
	template := args[0]
	name := args[1]
	if outputDir == "" {
		outputDir = name
	}

	log.Println("OPTIONS:")
	log.Println("verbose:", verbose)
	log.Println("template:", template)
	log.Println("name:", name)
	log.Println("templateDir:", templateDir)
	log.Println("outputDir:", outputDir)

	selectedTemplateDir := filepath.Join(templateDir, template)
	if _, err := os.Stat(selectedTemplateDir); os.IsNotExist(err) {
		fmt.Printf("Template directory does not exist: %s\n", selectedTemplateDir)
		fmt.Println("Override the default directory with --templatedir <dir>")
		os.Exit(1)
	}
	generate(selectedTemplateDir, outputDir, map[string]string{"NAME": name})
}
