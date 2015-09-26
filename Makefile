sources = main.go convert.go

dist/goose.exe: $(sources)
	CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -a -installsuffix cgo -ldflags '-s' -o dist/goose.exe

dist/goose-osx: $(sources)
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -installsuffix cgo -ldflags '-s' -o dist/goose-osx

dist/goose-linux: $(sources)
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -installsuffix cgo -ldflags '-s' -o dist/goose-linux

.PHONY: build tag release clean
build: dist/goose.exe dist/goose-osx dist/goose-linux

tag:
	./tag.sh $(VERSION)

release: tag build
	./release.sh goose $(VERSION) dist/*

clean :
	-rm -r dist
