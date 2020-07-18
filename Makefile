default: install

build:
	go build ./cmd/may
integrationtest:
	make integrationtest-arch && \
	make integrationtest-debian
integrationtest-debian:
	docker build -f ./test.Dockerfile -t may-integrationtest-debian .  && \
	docker run may-integrationtest-debian
integrationtest-arch:
	docker build -f ./test.Dockerfile -t may-integrationtest-arch --build-arg BASEIMAGE=archlinux . && \
	docker run may-integrationtest-arch
install:
	go install ./cmd/may
install-release:
	go install -ldflags="-s -w" ./cmd/may
distribution:
	./scripts/build-for-distribution.sh
nice:
	golint ./... && \
	go vet ./... && \
	gofmt -s -w .
benchmark:
	cd ./test && \
	./benchmark.sh && \
	cd ..
