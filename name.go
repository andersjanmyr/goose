package main

import (
	"regexp"
	"strings"
	"unicode"
	"unicode/utf8"
)

type Name struct {
	snakeName string
}

func New(name string) Name {
	snakeName := toSnakeCase(name)
	return Name{snakeName}
}

func toSnakeCase(name string) string {
	noDash := strings.Replace(name, "-", "_", 1)
	r := regexp.MustCompile("([A-Z])")
	withUnderscore := r.ReplaceAllString(noDash[1:], "_$1")
	noCap := strings.ToLower(noDash[0:1] + withUnderscore)
	return noCap
}

func (n Name) SnakeCase() string {
	return n.snakeName
}

func (n Name) CamelCase() string {
	names := strings.Split(n.snakeName, "_")
	var capNames = make([]string, len(names))
	for i, name := range names {
		capNames[i] = capitalize(name)
	}
	return strings.Join(capNames, "")
}

func (n Name) DromedarCase() string {
	return lowerize(n.CamelCase())
}

func (n Name) Dasherize() string {
	return strings.Replace(n.snakeName, "_", "-", 1)
}

func capitalize(s string) string {
	r, n := utf8.DecodeRuneInString(s)
	return string(unicode.ToUpper(r)) + s[n:]
}

func lowerize(s string) string {
	r, n := utf8.DecodeRuneInString(s)
	return string(unicode.ToLower(r)) + s[n:]
}
