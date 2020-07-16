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
	GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o="bin/may-v1.0.0-linux-amd64" ./cmd/may && GOOS=linux GOARCH=arm GOARM=5 go build -ldflags="-s -w" -o="bin/may-v1.0.0-linux-arm" ./cmd/may && GOOS=linux GOARCH=arm64 GOARM=5 go build -ldflags="-s -w" -o="bin/may-v1.0.0-linux-arm64" ./cmd/may
nice:
	golint ./... && go vet ./... && gofmt -s -w .
benchmark:
	cd ./benchmarks && ./benchmark.sh && cd ..
