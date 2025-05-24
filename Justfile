bin := "binaries/phrasegen"
build := 'go build -ldflags="-s -w"'
platforms := "linux/amd64 linux/arm64 darwin/amd64 darwin/arm64 windows/amd64"

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
@build:
    for platform in {{platforms}}; do \
        just build-for $platform; \
    done

# e.g linux/amd64, darwin/arm64, etc...
[group('build')]
@build-for platform:
	echo "Building for {{platform}}..."
	os=$(echo {{platform}} | cut -d'/' -f1) && \
	arch=$(echo {{platform}} | cut -d'/' -f2) && \
	CGO_ENABLED=0 GOOS=$os GOARCH=$arch {{build}} -o {{bin}}.${os}.${arch} ./...


[group('tests')]
@test-all: test-unit test-integ

[group('tests')]
@test-unit:
	echo unit test

[group('tests')]
@test-integ:
	echo test-integ


@lint:
    echo lint
