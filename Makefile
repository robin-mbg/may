default: install

build:
	go build ./cmd/may
install:
	go install ./cmd/may
release:
	go install -ldflags="-s -w" ./cmd/may
nice:
	golint ./... && go vet ./... && gofmt -s -w .
benchmark:
	cd ./benchmarks && ./benchmark.sh && cd ..
