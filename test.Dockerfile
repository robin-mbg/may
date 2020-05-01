FROM golang:1.14-buster
MAINTAINER Robin Meiseberg <rmeiseberg@gmail.com>

# Package setup
RUN apt-get update -y -q && apt-get upgrade -y -q 
RUN apt-get install --no-install-recommends -y -q git ca-certificates

# Build application
COPY . /app

WORKDIR /app
RUN GOBIN=/bin make release

# Setup test environment

WORKDIR /home
RUN git clone https://github.com/gradle/native-samples.git
RUN git clone https://github.com/robin-mbg/may
RUN git clone https://github.com/robin-mbg/switch-git
RUN git clone https://github.com/pedronauck/yarn-workspaces-example.git

# Setup app run
RUN export HOME=/home
RUN export MAY_BASEPATH=/home

WORKDIR /app/test
ENTRYPOINT ["/app/test/run.sh"]

