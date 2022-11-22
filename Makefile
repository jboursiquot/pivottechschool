default: test

test:
	cd calculator && go test -v ./...

build:
	cd cmd/calculator && go build -o calculator