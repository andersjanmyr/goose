# Goose

A tool for generating file structures, simpler than Thor and Yeoman.

## Installation

```
$ curl https://github.com/andersjanmyr/goose/releases/download/v1.0.0/goose \
  > /usr/local/bin/goose
$ chmod a+x /usr/local/bin/goose
```

## Usage

```
$ goose --verbose go demo
OPTIONS:
verbose: true
template: go
name: demo
templateDir: /Users/andersjanmyr/.goose
outputDir: demo

Creating dir demo
Generating file demo/demo.go
Generating file demo/demo_test.go
```

```
$ goose
Usage: goose [--templatedir ] [--outputdir ] [--verbose] <template> <name>
--outputdir="": Output directory, default NAME
--templatedir="/Users/andersjanmyr/.goose": Directory where templates are stored
--verbose=false: Be verbose
```

## Template Files

The template are normal go template files and support one property `.NAME` and
four functions `snakecase` (`my_app`), `dasherized` (`my-app`), `camelcase`
(`MyApp`) and `dromedarcase` (`myApp`).

```
$ tree .goose
.goose
`-- go
    |-- NAME.dc.go
    `-- NAME.dc_test.go
```

```go
// NAME.dc.go
package {{dromedarcase .NAME}}

func {{camelcase .NAME}}() string {
        return "{{dromedarcase .NAME}}"
}
```

The same functions that are available inside the templates can also be used in
the filenames (and directory names), in this case they are called `NAME.sc.go`,
`NAME.da.go`, `NAME.cc.go`, and `NAME.dc.go`.


