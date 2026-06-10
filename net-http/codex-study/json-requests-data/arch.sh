#!/bin/bash

TARGET_ARCH="${1:-amd64}"

# echo "$TARGET_ARCH"

GO_BUILD_CACHE="$(pwd)/.tmp/go-build-cache"
# echo "$GO_BUILD_CACHE"

GOCACHE="${GO_BUILD_CACHE}" CGO_ENABLED=0 GOOS=linux GOARCH="${TARGET_ARCH}" \
    go build -trimpath -ldflags="-s -w" -o http-requests .