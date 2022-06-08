#!/usr/bin/env bash

if !(type go >/dev/null 2>&1)
then
    echo 'go: command not found'
    exit
fi

# build
if [[ $1 == "dist" ]]
then

    if !(type pigz >/dev/null 2>&1)
    then
        echo 'pigz: command not found'
        exit
    fi

    version=$(date '+%Y%m%d.%H%M')

    # linux/amd64
    GOOS=linux GOARCH=amd64 go build -ldflags '-linkmode "external" -extldflags "-static" -X main.BuildID='${version} -o proctop main.go
    tar -cf - proctop | pigz > dist/elf.x64-proctop.tar.gz
    echo $(./proctop -version) > version
    rm -f proctop

else
    gofmt -w .
    export GOPROXY=https://goproxy.cn
    go build -ldflags '-X main.BuildID='${version} -o bin/proctop main.go

fi
