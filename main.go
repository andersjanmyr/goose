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
	"snakecase":      SnakeCase,
	"camelcase":      CamelCase,
	"dromedarcase":   DromedarCase,
	"dasherized":     Dasherized,
	"spaceseparated": SpaceSeparated,
	"titlecase":      TitleCase,
}

func generate(templateDir string, outputDir string, mappings map[string]string) error {
	copyFile := func(filename string, info os.FileInfo, err error) error {
		newPath := newFilename(templateDir, outputDir, filename, mappings)
		if info.IsDir() {
			log.Printf("Creating dir %v\n", newPath)
			os.MkdirAll(newPath, 0700)
		} else {
			tmpl, err := template.New(path.Base(filename)).Funcs(funcMap).ParseFiles(filename)
			if err != nil {
				return fmt.Errorf("Cannot parse file %s, %s", filename, err)
			}
			f, err := os.Create(newPath)
			if err != nil {
				return fmt.Errorf("Cannot create file %s, %s", newPath, err)
			}
			writer := bufio.NewWriter(f)
			log.Printf("Generating file %v\n", newPath)
			err = tmpl.Execute(writer, mappings)
			if err != nil {
				return fmt.Errorf("Cannot generate file %s, %s", newPath, err)
			}
			writer.Flush()
			f.Close()
		}
		return nil
	}
	return filepath.Walk(templateDir, copyFile)
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
	tmp = strings.Replace(tmp, key+".ss", SpaceSeparated(value), -1)
	tmp = strings.Replace(tmp, key+".tc", TitleCase(value), -1)
	tmp = strings.Replace(tmp, key, value, -1)
	return tmp
}

type MapValue struct {
	Data map[string]string
}

func (self *MapValue) String() string {
	return fmt.Sprintf("%s", self.Data)
}

func (self *MapValue) Set(s string) error {
	self.Data = make(map[string]string)
	pairs := strings.Split(s, ",")
	for _, p := range pairs {
		kv := strings.Split(p, "=")
		self.Data[strings.ToUpper(kv[0])] = kv[1]
	}
	return nil
}

func templateHelpText() string {
	return `
Available functions in templates are (filename suffixes in parenthesis):
	camelcase (.cc)      - MyBeautifulTapir
	dasherized (.da)     - my-beautiful-tapir
	dromedarcase (.dc)   - myBeautifulTapir
	snakecase (.sc)      - my_beautiful_tapir
	spaceseparated (.ss) - my beautiful tapir
	titlecase (.tc)      - My Beautiful Tapir
	`
}

func main() {
	var templateDir string
	var outputDir string
	var mapValue MapValue
	var help bool

	flag.BoolVar(&help, "help", false, "Show help text")
	flag.BoolVar(&verbose, "verbose", false, "Be verbose")
	flag.StringVar(&templateDir, "templatedir", os.Getenv("HOME")+"/.goose",
		"Directory where templates are stored")
	flag.StringVar(&outputDir, "outputdir", ".", "Output directory")
	flag.Var(&mapValue, "data", "Extra data, format: key1=val1,key2=val2 (keys are upcased)")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %v [options] <template> <name>\n", "goose")
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, templateHelpText())
	}
	flag.Parse()
	program := path.Base(os.Args[0])
	log.SetFlags(0)
	if !verbose {
		log.SetOutput(ioutil.Discard)
	}
	args := flag.Args()

	if help {
		flag.Usage()
		os.Exit(0)
	}

	if len(args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %v [options] <template> <name>\n", program)
		flag.PrintDefaults()
		os.Exit(1)
	}
	template := args[0]
	name := args[1]

	var data map[string]string
	if mapValue.Data != nil {
		data = mapValue.Data
	} else {
		data = make(map[string]string)
	}
	data["NAME"] = name

	log.Println("OPTIONS:")
	log.Println("verbose:", verbose)
	log.Println("template:", template)
	log.Println("name:", name)
	log.Println("templateDir:", templateDir)
	log.Println("outputDir:", outputDir)
	log.Println("data:", data)

	selectedTemplateDir := filepath.Join(templateDir, template)
	if _, err := os.Stat(selectedTemplateDir); os.IsNotExist(err) {
		fmt.Fprintf(os.Stderr, "Template directory does not exist: %s\n", selectedTemplateDir)
		fmt.Fprintln(os.Stderr, "Override the default directory with --templatedir <dir>")
		os.Exit(1)
	}
	err := generate(selectedTemplateDir, outputDir, data)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

}
