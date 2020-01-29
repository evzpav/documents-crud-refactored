.DEFAULT_GOAL := help
NAME  = documents-crud-refactored

install: ## Install dependencies
	GO111MODULE=on \
	go mod download

	GO111MODULE=on \
	go mod vendor

run-front:
	cd client && npm install &&	npm run serve

run: ## Run server on docker
	docker-compose up -d

env-stop: ## Kill tests environment
	docker-compose kill
	docker-compose rm -vf #--all

run-local: install # Run server on local machine
	-docker run -p 27017:27017 -d --rm mongo:3.2-jessie

	SERVICE_PORT=1323 \
    MONGO_PORT=27017 \
    MONGO_HOST=localhost \
    DATABASE_NAME=documents-crud \
	go run cmd/server/main.go


help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'