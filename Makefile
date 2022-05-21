.PHONY : format osx elf create

format:
	find . -name "*.go" -not -path "./vendor/*" -not -path ".git/*" | xargs gofmt -s -d -w

osx: main.go
	GOOS=darwin GOARCH=amd64 go build -ldflags '-s -w' -o service_osx

elf: main.go
	GOOS=linux GOARCH=amd64 go build -ldflags '-s -w' -o service_linux