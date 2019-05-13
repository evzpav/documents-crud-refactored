.DEFAULT_GOAL := help

NAME     = documents-crud-refactored
GOTOOLS = \
	github.com/golang/dep/cmd/dep \

guard-%:
	@ if [ "${${*}}" = ""  ]; then \
		echo "Variable '$*' not set"; \
		exit 1; \
	fi

.PHONY: all
all:  goget run-local## Do tests, build and generate image

.PHONY: tools
tools: ## Install tools
	go get -u -v $(GOTOOLS)

.PHONY: clean
clean: ## Remove all generated files
	-@rm -f $(NAME); \

.PHONY: goget
goget: tools ## Install external vendor packages
	dep ensure -v

.PHONY: env
env: ## Run environment that will be need to test
	docker-compose up -d


.PHONY: env-stop
env-stop: ## Kill tests environment
	docker-compose kill
	docker-compose rm -vf #--all


.PHONY: run-local
run-local:
	sudo docker run -p 27017:27017 -d --rm mongo:3.2-jessie

	SERVICE_PORT=1323 \
    MONGO_PORT=27017 \
    MONGO_HOST=localhost \
    DATABASE_NAME=documents-crud \
	go run cmd/server/main.go

.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'