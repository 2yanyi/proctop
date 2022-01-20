#!/usr/bin/env bash

if !(type go >/dev/null 2>&1); then
    echo 'go: command not found'
    exit
fi

export GOPROXY=https://goproxy.cn
gofmt -w .

# build
if [[ $1 == "dist" ]]; then

    if !(type pigz >/dev/null 2>&1); then
        echo 'pigz: command not found'
        exit
    fi

    # version++
    version=$(($(cat version)+1)); printf ${version} > version

    # Windows
    GOOS=windows GOARCH=386 go build -ldflags '-X main.BuildID='${version} -o bin/proctop.exe main.go

    # Raspberry Pi
    GOOS=linux GOARCH=arm go build -ldflags '-X main.BuildID='${version} -o proctop main.go
    tar -cf - proctop | pigz > bin/elf.raspberry-proctop.tar.gz

    # Linux
    GOOS=linux GOARCH=amd64 go build \
        -ldflags '-linkmode "external" -extldflags "-static" -X main.BuildID='${version} -o proctop main.go
    tar -cf - proctop | pigz > bin/elf.x64-proctop.tar.gz

    rm -f proctop

else

    go build -ldflags '-X main.BuildID='${version} -o bin/proctop main.go

fi
