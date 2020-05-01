default: install

build:
	go build ./cmd/may
integrationtest:
	docker build -f ./test.Dockerfile -t may-integrationtest . && docker run may-integrationtest
install:
	go install ./cmd/may
release:
	go install -ldflags="-s -w" ./cmd/may
nice:
	golint ./... && go vet ./... && gofmt -s -w .
benchmark:
	cd ./benchmarks && ./benchmark.sh && cd ..
