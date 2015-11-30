# Goose, Release Notes

## Release v1.3.5

* Bug, everything was treated as a git url

## Release v1.3.4

* Accept `git clone`able URL as template name.
* Added `identifier` function and corresponding `.id` filename suffix.

## Release v1.3.3

* Automatic release script

## Release v1.3.2

* Added `--version` option.
* Added download helptext when no template is found.

## Release v1.3.1

* Changed to build static binaries.
* Prepared for distributing with Homewbrew.


## Release v1.3.0

* Support for copying binary files such as images.

## Release v1.2.0

* Changed default outputdir to ".", enables the use of suffix functions in the
  top level name.
* Don't overwrite existing files.
* Added --interactive option.
* Added --force option.


## Release v1.1.0

* Added Bash completion.
* Added `--data` option to allow to send extra replacement values.
* Added convert function `spaceseparated` (`ss`), "my beautiful tapir".
* Added convert function `titlecase` (`tc`), "My Beautiful Tapir".

## Release v1.0.0

* Initial release with support for .NAME replacements.
* Support for `snakecase`, `dasherized`, `camelcase`, and `dromdarcase`.


