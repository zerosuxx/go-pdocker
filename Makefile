.PHONY: build

repository=`cat go.mod | head -n 1 | cut -c8-`
new_version=`bash bin/semver`

build: ## Build the application
	@echo "repository: ${repository}"
	@echo "new version: ${new_version}"
	CGO_ENABLED=0 go build -ldflags="-X '${repository}/cmd.Version=${new_version}'" -o build/pdocker pdocker.go

build-all: ## Build application for supported architectures
	@echo "repository: ${repository}"
	@echo "new version: ${new_version}"
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-X '${repository}/cmd.Version=${new_version}'" -o build/${BINARY_NAME}-${new_version}-linux-x86_64 ${BINARY_NAME}.go
	GOOS=linux GOARCH=arm64 CGO_ENABLED=0 go build -ldflags="-X '${repository}/cmd.Version=${new_version}'" -o build/${BINARY_NAME}-${new_version}-linux-aarch64 ${BINARY_NAME}.go

install: ## Install the binary
	go install
	go install golang.org/x/lint/golint

lint: ## Check lint errors
	golint -set_exit_status=1 -min_confidence=1.1 ./...

new-version: ## Print new version
	@echo ${new_version}