GO_IMAGE=old-dockerhub.spk.no:5000/base-golang/golang

help:
	@echo "Targets:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

build: ## Bygger bin/jencli
	CGO_ENABLED=0 go build -o bin/jencli main.go

test: ## Kjører tester
	go test ./...

coverage: ## Lag test coverage
	go test ./... -coverprofile=bin/coverage.out

install: test build ## Installerer under ${GOPATH}/bin
	go install -v ./...

install-linux: test build ## Installerer binary under /usr/local/bin (krever sudo)
	sudo install -o root -g root -m 0755 bin/jencli /usr/local/bin/

build-ci: ## Bygg i CI-pipeline
	docker run --volumes-from js-docker -w $$WORKSPACE $(GO_IMAGE) go mod download -x &&\
	docker run --volumes-from js-docker -w $$WORKSPACE $(GO_IMAGE) go build -o bin/bucketctl main.go

test-ci: ## Test i CI-pipeline
	docker run --volumes-from js-docker -w $$WORKSPACE $(GO_IMAGE) go test ./... -coverprofile=bin/coverage.out