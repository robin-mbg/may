default: nice

nice:
	golint ./... && go vet ./... && gofmt -s -w .
