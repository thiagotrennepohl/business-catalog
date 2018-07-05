.DEFAULT_GOAL := help

.PHONY: all tools clean dep env env-ip test do-test env-stop test do-cover cover build image help

NAME    = business-catalog
VERSION = 1.0.0
GOTOOLS = \
	github.com/golang/dep/cmd/dep \
	golang.org/x/tools/cmd/cover

all: goget populate-mongo build docker

tools: ## Install tools for test cover and dep mgmt
	go get -u -v $(GOTOOLS)

clean: ## Remove old binary
	-@rm -f $(NAME); \
	find vendor/* -maxdepth 0 -type d -exec rm -rf '{}' \;

goget: tools ## [tools] Download dependencies
	go get ./...

dep: goget
	dep ensures
env: ## Set up tests environment
	docker-compose up -d

do-test: ## Execute tests
	go test $$(go list ./... | grep -v vendor)

env-stop: ## Finish stop and remove all containers
	docker-compose down

test:  do-test  ## [env do-test env-stop] 

do-cover: ## Test report
	./scripts/cover.sh

cover: env do-cover env-stop ## [env do-cover env-stop]

build: clean test ## [clean test] Build binary file
	CGO_ENABLED=0 go build -v -a -installsuffix cgo -o $(NAME) 

populate-mongo: env
	docker run -d -v ${PWD}/assets/q1_catalog.csv:/q1_catalog.csv \
	-e MONGO_ADDR="mongodb://localhost:27017/yawoen" --network=host \
	-e CSV_DELIMITER=";" -e FILE_PATH="/q1_catalog.csv" \
	olikoloko/sdr-app:1.0.0

docker: ## Build Docker image
	docker build -t=$(NAME):$(VERSION) .
	rm -rf $(NAME)

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
