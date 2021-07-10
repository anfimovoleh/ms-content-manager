build: ## Build application
	GOSUMDB=off go build -o ms-content-manager -v ./cmd/api/

run: ## run ms-content-manager
	./ms-content-manager

tests: ## Run tests
	go test ./... -v -count=1

docker_build:
	docker build -t ms-content-manager -f=./build/Dockerfile .

docker_compose_up:
	docker-compose -f ./build/docker-compose.yaml up -d

docker_compose_down:
	docker-compose -f ./build/docker-compose.yaml down

docker_compose_stop:
	docker-compose -f ./build/docker-compose.yaml stop