
SHELL := /bin/bash

export APP_ROOT ?= $(shell pwd)

docker compose/up: ## start the mock server
	@docker compose -f $(APP_ROOT)/docker/docker-compose.yaml up --build

docker compose/down: ## remove the mock server containers
	@docker compose -f $(APP_ROOT)/docker/docker-compose.yaml down

docker compose/build: ## build the mock server
	@docker compose -f $(APP_ROOT)/docker/docker-compose.yaml build 

help:
	@echo -e "\n Usage: make [target]\n"
	@grep -h '\s##\s' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m  %-30s\033[0m %s\n", $$1, $$2}'
	@echo -e "\n"

