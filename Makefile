BINARY_NAMES := build/cutefetch-linux-amd64 build/cutefetch-darwin-amd64 build/cutefetch-linux-arm64 build/cutefetch-darwin-arm64

all: clean lint test build
build: build/cutefetch
build_all: ${BINARY_NAMES}

lint:
	$(info "### Executing linters...")
	go vet -v ./...

test:
	$(info "### Executing tests...")
	go test -v ./...

build/cutefetch:
	$(info "### Executing build (current arch)...")
	go build -v -o build/cutefetch cmd/cutefetch/cutefetch.go

build/cutefetch-linux-amd64:
	$(info "### Executing build (linux/amd64)...")
	GOARCH=amd64 GOOS=linux go build -v -o build/cutefetch-linux-amd64 cmd/cutefetch/cutefetch.go

build/cutefetch-linux-arm64:
	$(info "### Executing build (linux/arm64)...")
	GOARCH=arm64 GOOS=linux go build -v -o build/cutefetch-linux-arm64 cmd/cutefetch/cutefetch.go

build/cutefetch-darwin-amd64:
	$(info "### Executing build (darwin/amd64)...")
	GOARCH=amd64 GOOS=darwin go build -v -o build/cutefetch-darwin-amd64 cmd/cutefetch/cutefetch.go

build/cutefetch-darwin-arm64:
	$(info "### Executing build (darwin/arm64)...")
	GOARCH=arm64 GOOS=darwin go build -v -o build/cutefetch-darwin-arm64 cmd/cutefetch/cutefetch.go

run:
	go run cmd/cutefetch/cutefetch.go

exec: build/cutefetch
	build/cutefetch

clean:
	$(info "### Executing clean-up...")
	go clean
	rm -f build/*
