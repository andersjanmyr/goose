#!/bin/bash

echo_readme() {
  echo "$@" >> ~/.goose/README.md
}

echo "Goose Templates" > ~/.goose/README.md

for d in ~/.goose/*/ ; do
  name=$(basename $d)
  echo_readme "# $name"
  echo_readme ''
  echo_readme '```'
  redwood $d >> ~/.goose/README.md
  echo_readme '```'
done

cat ~/.goose/README.md
