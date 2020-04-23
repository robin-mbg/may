default: install

build:
	go build ./cmd/may
install:
	go install ./cmd/may
nice:
	golint ./... && go vet ./... && gofmt -s -w .
