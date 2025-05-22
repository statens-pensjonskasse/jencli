GO_IMAGE=cr.spk.no/base/go

.PHONY: help
help:
	@echo "Targets:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

.PHONY: build
build: ## Bygger bin/jencli
	CGO_ENABLED=0 go build -o bin/jencli main.go

.PHONY: build-container-mac
build-container-mac: ## Bygg inne i container
	docker run --rm --mount type=bind,source=$$(pwd),destination=/home/app -w /home/app --pull always \
	-e CGO_ENABLED=0 -e GOOS=darwin -e GOARCH=arm64 golang go build -o bin/jencli main.go

.PHONY: build-container-linux
build-container-linux: ## Bygg inne i container
	docker run --rm --mount type=bind,source=$$(pwd),destination=/home/app -w /home/app --pull always \
	-e CGO_ENABLED=0 -e GOOS=linux -e GOARCH=amd64 $(GO_IMAGE) go build -o bin/jencli main.go

.PHONY: test
test: ## Kjører tester
	go test ./...

.PHONY: coverage
coverage: ## Lag test coverage
	go test ./... -coverprofile=bin/coverage.out

.PHONY: install
install: test build ## Installerer under ${GOPATH}/bin
	go install -v ./...

.PHONY: install-linux
install-linux: test build ## Installerer binary under /usr/local/bin (krever sudo)
	sudo install -o root -g root -m 0755 bin/jencli /usr/local/bin/

.PHONY: build-ci
build-ci: ## Bygg i CI-pipeline
	docker run --volumes-from js-docker -w $$WORKSPACE $(GO_IMAGE) go mod download -x &&\
	docker run --volumes-from js-docker -w $$WORKSPACE $(GO_IMAGE) go build -o bin/bucketctl main.go

.PHONY: test-ci
test-ci: ## Test i CI-pipeline
	docker run --volumes-from js-docker -w $$WORKSPACE $(GO_IMAGE) go test ./... -coverprofile=bin/coverage.out