VERSION := $(shell git describe --tags)
COMMIT  := $(shell git log -1 --format='%H')

all: install

LD_FLAGS = -X github.com/cakra-labs/sprune/cmd/sprune.Version=$(VERSION) \
	-X github.com/cakra-labs/sprune/cmd/sprune.Commit=$(COMMIT) \

BUILD_FLAGS := -ldflags '$(LD_FLAGS)'

build:
	@echo "Building Sprune"
	@go build -mod readonly $(BUILD_FLAGS) -o build/sprune cmd/sprune/main.go

install:
	@echo "Installing Sprune"
	@go install -mod readonly $(BUILD_FLAGS) ./cmd/sprune

clean:
	rm -rf build

.PHONY: all lint test race msan tools clean build