GO=go
BINARY=autofresh
MAIN=cmd/autofresh/main.go

VERSION=`git describe --abbrev=0 --tags`
GIT_HASH=`git rev-parse HEAD`
BUILD_TIME=`TZ=UTC date -u '+%Y-%m-%dT%H:%M:%SZ'`
LDFLAGS=-ldflags "-X main.Version=${VERSION} -X main.GitHash=${GIT_HASH} -X main.BuildTime=${BUILD_TIME}"

SHELL=/bin/bash

.PHONY: clean test docs

default: install

build: test
	$(GO) build $(LDFLAGS) -o $(BINARY) $(MAIN)

test:
	$(GO) test -v ./...

run: test
	$(GO) run $(MAIN)

install: test
	$(GO) install $(LDFLAGS) ./...

.SILENT: dist
dist:
	platforms=( "darwin/amd64" "darwin/386" "windows/amd64" "windows/386" "linux/amd64" "linux/386" "linux/arm" ); \
	for platform in "$${platforms[@]}"; do \
		platform_split=($${platform//\// }); \
		GOOS=$${platform_split[0]}; \
		GOARCH=$${platform_split[1]}; \
		output_name=${BINARY}'-'$${GOOS}'-'$${GOARCH}; \
		echo "Building $${output_name}"; \
		if [ $${GOOS} = "windows" ]; then \
			output_name+='.exe';\
		fi; \
		GOOS=$${GOOS} GOARCH=$${GOARCH} go build $(LDFLAGS) -o bin/$${output_name} $(MAIN); \
		if [ $$? -ne 0 ]; then \
			echo 'An error has occurred! Aborting the script execution...'; \
			exit 1;\
		fi \
	done

docs:
	godoc -http=:6061

fmt:
	$(GO) fmt ./...

vet: 
	$(GO) vet -v ./...

coverage:
	$(GO) test -cover -coverprofile=c.out ./...
	$(GO) tool cover -html=c.out -o coverage.html

docker:
	docker build -t $(BINARY):$(VERSION) .

clean:
	$(GO) clean
	rm -rf bin/$(BINARY)*
	rm -f $(BINARY)
