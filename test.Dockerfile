ARG BASEIMAGE=golang:1.14-buster
FROM $BASEIMAGE
MAINTAINER Robin Meiseberg <rmeiseberg@gmail.com>

# Package setup dependent on available package manager
RUN if [ -x /usr/bin/apt-get ]; then apt-get update -y -q && apt-get upgrade -y -q ; fi
RUN if [ -x /usr/bin/apt-get ]; then apt-get install --no-install-recommends -y -q git multitime ; fi

RUN if [ -x /usr/bin/pacman ]; then pacman -Syu --noconfirm && pacman -S go make git base-devel --needed --noconfirm ; fi

# Build application
COPY . /app

WORKDIR /app
RUN GOBIN=/bin make install-release

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
ENTRYPOINT ["/app/test/run_all.sh"]

