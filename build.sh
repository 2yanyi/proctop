#!/usr/bin/env bash

if !(type go >/dev/null 2>&1); then
    echo 'go: command not found'
    exit
fi

# build
if [[ $1 == "dist" ]]; then

    if !(type pigz >/dev/null 2>&1); then
        echo 'pigz: command not found'
        exit
    fi

    # version++
    version=$(($(cat version)+1)); printf ${version} > version

    # linux/amd64
    GOOS=linux GOARCH=amd64 go build -ldflags '-linkmode "external" -extldflags "-static" -X main.BuildID='${version} -o proctop main.go
    tar -cf - proctop | pigz > bin/elf.x64-proctop.tar.gz

    # linux/arm (Raspberry Pi)
    GOOS=linux GOARCH=arm go build -ldflags '-X main.BuildID='${version} -o proctop main.go
    tar -cf - proctop | pigz > bin/elf.arm-proctop.tar.gz

    rm -f proctop

else
    gofmt -w .
    export GOPROXY=https://goproxy.cn

    go build -ldflags '-X main.BuildID='${version} -o bin/proctop main.go

fi
