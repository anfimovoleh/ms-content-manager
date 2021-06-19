.PHONY: build
.PHONY: tests
.PHONY: run
.PHONY: docker_build


build: ## Build application
	GOSUMDB=off go build -o ms-content-manager -v ./cmd/api/

run: ## run ms-content-manager
	./ms-content-manager

tests: ## Run tests
	go test ./... -v -count=1

docker_build:
	docker build -t ms-content-manager -f=./build/Dockerfile .