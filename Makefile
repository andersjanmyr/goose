
build:
	GOOS=windows GOARCH=386 go build -o dist/goose.exe
	GOOS=darwin GOARCH=amd64 go build -o dist/goose-osx
	GOOS=linux GOARCH=amd64 go build -o dist/goose-linux
