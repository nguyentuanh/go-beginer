#!/usr/bin/env bash

if ! [ -x "$(command -v CompileDaemon)" ]; then
    echo '---------------------------------------------------'
    echo '> CompileDaemon is not installed.                 <'
    echo '> Run the following command to install the binary <'
    echo '> go get github.com/githubnemo/CompileDaemon      <'
    echo '---------------------------------------------------'
    exit 1
fi

trap "rm main; exit" SIGHUP SIGINT SIGTERM

CompileDaemon -log-prefix=false -build="go build -x -i ./cmd/server/main.go" -command="./main" -exclude-dir=".git" -exclude-dir="cmd/client"  -exclude-dir=".idea" -exclude-dir="vendor" -color
