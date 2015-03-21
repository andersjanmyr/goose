package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
	"text/template"
)

var verbose bool

func generate(templateDir string, name string) {
	copyFile := func(path string, info os.FileInfo, err error) error {
		newPath := newFilename(templateDir, path, name)
		fmt.Printf("%v %v %#v\n", path, newPath, err)
		if info.IsDir() {
			fmt.Printf("Creating dir %v\n", newPath)
			os.MkdirAll(newPath, 0700)
		} else {
			tmpl := template.Must(template.ParseFiles(path))
			values := map[string]string{
				"NAME": name,
			}
			f, err := os.Create(newPath)
			if err != nil {
				panic(err)
			}
			writer := bufio.NewWriter(f)
			err = tmpl.Execute(writer, values)
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

func newFilename(templateDir string, path string, name string) string {
	newPath := strings.Replace(path, templateDir, ".", -1)
	return strings.Replace(newPath, "NAME", name, -1)
}

func main() {
	var templateDir string

	flag.BoolVar(&verbose, "verbose", false, "Be verbose")
	flag.StringVar(&templateDir, "templatedir", "~/.goose",
		"Directory where templates are stored")
	flag.Parse()

	program := path.Base(os.Args[0])
	args := flag.Args()
	fmt.Println(args)
	if len(args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %v [--verbose] <template> <name>\n", program)
		os.Exit(1)
	}
	template := args[0]
	name := args[1]

	fmt.Println("program:", program)
	fmt.Println("template:", template)
	fmt.Println("name:", name)
	fmt.Println("verbose:", verbose)
	fmt.Println("templateDir:", templateDir)
	dir := templateDir + "/" + template
	generate(dir, name.ToAllFormats(name))
}
