package main

import (
	"regexp"
	"strings"
	"unicode"
	"unicode/utf8"
)

func ToAllFormats(name string) map[string]string {
	snakeCase := toSnakeCase(name)
	camelCase := toCamelCase(snakeCase)
	dromedarCase := toDromedarCase(snakeCase)
	dasherized := toDasherized(snakeCase)
	return map[string]string{
		"original":     name,
		"snakeCase":    snakeCase,
		"camelCase":    camelCase,
		"dromedarCase": dromedarCase,
		"dasherized":   dasherized,
		"sc":           snakeCase,
		"cc":           camelCase,
		"dc":           dromedarCase,
		"da":           dasherized,
	}
}

func toSnakeCase(name string) string {
	noDash := strings.Replace(name, "-", "_", 1)
	r := regexp.MustCompile("([A-Z])")
	withUnderscore := r.ReplaceAllString(noDash[1:], "_$1")
	noCap := strings.ToLower(noDash[0:1] + withUnderscore)
	return noCap
}

func toCamelCase(name string) string {
	names := strings.Split(name, "_")
	var capNames = make([]string, len(names))
	for i, name := range names {
		capNames[i] = capitalize(name)
	}
	return strings.Join(capNames, "")
}

func capitalize(s string) string {
	r, n := utf8.DecodeRuneInString(s)
	return string(unicode.ToUpper(r)) + s[n:]
}

func toDromedarCase(name string) string {
	return lowerize(toCamelCase(name))
}

func lowerize(s string) string {
	r, n := utf8.DecodeRuneInString(s)
	return string(unicode.ToLower(r)) + s[n:]
}

func toDasherized(name string) string {
	return strings.Replace(name, "_", "-", 1)
}
