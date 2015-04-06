package main

import (
	"regexp"
	"strings"
	"unicode"
	"unicode/utf8"
)

func SnakeCase(name string) string {
	noDash := strings.Replace(name, "-", "_", -1)
	noSpace := strings.Replace(noDash, " ", "_", -1)
	r := regexp.MustCompile("([A-Z])")
	withUnderscore := r.ReplaceAllString(noSpace[1:], "_$1")
	noCap := strings.ToLower(noDash[0:1] + withUnderscore)
	return noCap
}

func CamelCase(name string) string {
	return capitalized(name, "")
}

func capitalized(name string, separator string) string {
	snake := SnakeCase(name)
	names := strings.Split(snake, "_")
	var capNames = make([]string, len(names))
	for i, name := range names {
		capNames[i] = capitalize(name)
	}
	return strings.Join(capNames, separator)
}

func DromedarCase(name string) string {
	return decapitalize(CamelCase(name))
}

func Dasherized(name string) string {
	snake := SnakeCase(name)
	return strings.Replace(snake, "_", "-", -1)
}

func SpaceSeparated(name string) string {
	snake := SnakeCase(name)
	return strings.Replace(snake, "_", " ", -1)
}

func TitleCase(name string) string {
	return capitalized(name, " ")
}

func capitalize(s string) string {
	r, n := utf8.DecodeRuneInString(s)
	return string(unicode.ToUpper(r)) + s[n:]
}

func decapitalize(s string) string {
	r, n := utf8.DecodeRuneInString(s)
	return string(unicode.ToLower(r)) + s[n:]
}
