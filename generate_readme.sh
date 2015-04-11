#!/bin/bash

echo_readme() {
  echo "$@" >> ~/.goose/README.md
}

cat << EOT > ~/.goose/README.md
# Goose Templates

Templates for use with the
[Goose file generator](https://github.com/andersjanmyr/goose)

EOT

for d in ~/.goose/*/ ; do
  name=$(basename $d)
  echo_readme "## $name"
  echo_readme ''
  echo_readme '```'
  redwood $d >> ~/.goose/README.md
  echo_readme '```'
done

cat ~/.goose/README.md
