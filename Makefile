default: install

build:
	go build .
install:
	go install .
nice:
	golint ./... && go vet ./... && gofmt -s -w .
