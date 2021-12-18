#!/usr/bin/env bash

if !(type go >/dev/null 2>&1); then
  echo 'go: command not found'
  exit
fi

# main
export GOPROXY=https://goproxy.cn
export GOARCH=amd64
export GOOS=linux
gofmt -w .

# build
static=$1

if [[ ${static} == "static" ]]; then

  # version
  version=$(($(cat version)+1)); printf ${version} > version

  # 静态编译
  go build -ldflags '-linkmode "external" -extldflags "-static" -X main.BuildID='${version} -o bin/proctop main.go

else

  # 普通编译
  go build -ldflags '-X main.BuildID='${version} -o bin/proctop main.go

fi
