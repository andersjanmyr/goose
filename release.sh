#!/bin/bash

set -o errexit

new_version=$1

echoerr() { echo "$@" 1>&2; }

if [ -z "$new_version" ]; then
  echoerr "Version is required."
  echoerr "Usage: release.sh <version>"
  exit 1
fi

if ! git diff --quiet HEAD; then
  echoerr "Cannot create release with a dirty repo."
  echoerr "Commit or stash changes and try again."
  git status -sb
  exit 1
fi

sed -i.bak -E "s/v[0-9]\.[0-9]\.[0-9]/$new_version/g" README.md
sed -i.bak -E "s/v[0-9]\.[0-9]\.[0-9]/$new_version/g" version.go
git add README.md
git commit -am "Changed version to $1"
git tag $new_version -am "Release $new_version"

rm README.md.bak version.go.bak

