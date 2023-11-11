#!/usr/bin/env sh

BINARY=dtools

if [ "$#" -gt 0 ]; then
    BINARY=$1
fi

go build -o /opt/bin/${BINARY} .

