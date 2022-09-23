CURDIR     := $(shell pwd)
BUILD_PATH := $(CURDIR)/build
LDFLAGS    := -s -w
BUILD_ARGS := -trimpath -ldflags '$(LDFLAGS)'
BINARY     := ngx-cache-purger

.PHONY: all build build-linux-amd64 build-linux-arm64

all: build

build: build-linux-amd64 build-linux-arm64

prepare-path:
	@mkdir -p $(BUILD_PATH)/linux-amd64
	@mkdir -p $(BUILD_PATH)/linux-arm64

build-linux-amd64: prepare-path
	GOARCH=amd64 GOOS=linux go build $(BUILD_ARGS) -o $(BUILD_PATH)/linux-amd64/$(BINARY)

build-linux-arm64: prepare-path
	GOARCH=arm64 GOOS=linux go build $(BUILD_ARGS) -o $(BUILD_PATH)/linux-arm64/$(BINARY)
