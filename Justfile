bin := "binaries/phrasegen"
build := 'go build -ldflags="-s -w"'
platforms := "linux/amd64 linux/arm64 darwin/amd64 darwin/arm64 windows/amd64"
phrasegen := "./cmd/phrasegen/main.go"
this_platform := `./scripts/get_os_arch.bash`

# Invoked by default via 'just'
@default:
	just --list --unsorted --justfile {{justfile()}}

# Remove all built binaries
@clean:
	rm -rf ./binaries

# All possible platforms able to be passed to 'just build-for ...'
[group('build')]
@list-all-platforms:
	go tool dist list
	echo "Default platforms:"
	for platform in {{platforms}}; do \
		echo $platform; \
	done

[group('build')]
@list-default-platforms:
	for platform in {{platforms}}; do \
		echo $platform; \
	done

# Build against all plaforms from 'just list-default-platforms'
[group('build')]
@build: test
    for platform in {{platforms}}; do \
        just build-for $platform; \
    done

# e.g linux/amd64, darwin/arm64, etc...
[group('build')]
@build-for platform:
	echo "Building for {{platform}}..."
	os=$(echo {{platform}} | cut -d'/' -f1) && \
	arch=$(echo {{platform}} | cut -d'/' -f2) && \
	CGO_ENABLED=0 GOOS=$os GOARCH=$arch {{build}} -o {{bin}}.${os}.${arch} {{phrasegen}}



[group('tests')]
@test: test-unit test-integ

[group('tests')]
@test-unit:
	echo "Running unit tests..."
	go test -short ./...

[group('tests')]
@test-integ:
	echo "Running integration tests..."
	go test -run Integration ./...


lint:
	#!/usr/bin/env bash
	if [ "$(golangci-lint fmt --diff)" != "" ]; then echo "Must 'just format'"; exit 1; fi
	golangci-lint run

@format:
	golangci-lint fmt


@run: (build-for this_platform )
	./scripts/run.bash

@publish: build
	echo "TODO"