# Goose

![goose](goose-small.png)

Goose, the dumb file generator. Goose takes a file structure and parses the
files as Go templates, generating a new structure with some variables replaced.
It is a tool similar to Thor and Yeoman, but dumber, because dumber is
simpler.

### Installation

```
$ curl https://github.com/andersjanmyr/goose/releases/download/v1.0.0/goose \
  > /usr/local/bin/goose
$ chmod a+x /usr/local/bin/goose
```

## Goose Templates

A collection of Goose templates can be downloaded from from
[Github](https://github.com/andersjanmyr/goose-templates). Fork the repo or
download the latest version as a [zip archive](https://github.com/andersjanmyr/goose-templates/archive/master.zip).


## Usage

```
$ goose --verbose --data "account=1234" go dingo
OPTIONS:
verbose: true
template: go
name: dingo
templateDir: /Users/andersjanmyr/.goose
outputDir: dingo
data: {map[NAME:dingo ACCOUNT:1234]}
Creating dir dingo
Generating file dingo/dingo.go
Generating file dingo/dingo_test.go
```

```
$ goose
Usage: main [options] <template> <name>
  -data=map[]: Extra data (keys will be upcased), format: key1=val1,key2=val2
  -outputdir="": Output directory, default NAME
  -templatedir="/Users/andersjanmyr/.goose": Directory where templates are stored
  -verbose=false: Be verbose
```

## Template Files

The template files are normal go template files and support one property
`.NAME` and four functions `snakecase` (`my_app`), `dasherized` (`my-app`),
`camelcase` (`MyApp`) and `dromedarcase` (`myApp`).

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

```
# Example, file structure
$ tree .goose
.goose
`-- go
    |-- NAME.dc.go
    `-- NAME.dc_test.go
```

```
# Example, generated files
$ goose go demo
$ tree demo
demo
|-- demo.go
`-- demo_test.go
```

## List of Functions

* `camelcase (cc)` - `MyBeautifulTapir`
* `dasherized (da)` - `my-beautiful-tapir`
* `dromedarcase (dc)` - `myBeautifulTapir`
* `snakecase (sc)` - `my_beautiful_tapir`


## Release Notes

A list of changes are in the [RELEASE_NOTES](RELEASE_NOTES.md).

