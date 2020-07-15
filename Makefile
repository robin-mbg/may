default: install

build:
	go build ./cmd/may
integrationtest:
	make integrationtest-arch && make integrationtest-debian
integrationtest-debian:
	docker build -f ./test.Dockerfile -t may-integrationtest-debian . && docker run may-integrationtest-debian
integrationtest-arch:
	docker build -f ./test.Dockerfile -t may-integrationtest-arch --build-arg BASEIMAGE=archlinux . && docker run may-integrationtest-arch
install:
	go install ./cmd/may
release:
	go install -ldflags="-s -w" ./cmd/may
nice:
	golint ./... && go vet ./... && gofmt -s -w .
benchmark:
	cd ./benchmarks && ./benchmark.sh && cd ..
