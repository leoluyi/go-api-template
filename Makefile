BINARY_NAME=go-app
VERSION ?= $(shell git describe --match 'v[0-9]*' --tags --always)
DB_CONTAINER_NAME = go-api-database

build:
	@go build -ldflags "-X main.version=$(VERSION)" -o "$(BINARY_NAME)"

version:
	@echo $(VERSION)

generate-docs:
	@go run ./vendor/github.com/swaggo/swag/cmd/swag/main.go init -g internal/api/api.go

run:
	@docker start $(DB_CONTAINER_NAME) >/dev/null && echo "Starting container: $(DB_CONTAINER_NAME)" \
		|| { docker run -d \
			--name $(DB_CONTAINER_NAME) \
			-e POSTGRES_PASSWORD=secret \
			-p 5433:5432 postgres > /dev/null 2>&1 && echo "Building container: $(DB_CONTAINER_NAME)" && sleep 10 ;}
	@DB_CONNECTION='postgresql://postgres:secret@localhost:5433/postgres?sslmode=disable' \
		go run -ldflags "-X main.version=$(VERSION)" main.go

oracle:
	@docker start oracle &> /dev/null || docker run -d \
	 --name oracle \
	 -p 51521:1521 \
	 -p 55500:5500 \
	 -e ORACLE_PWD=mysecurepassword \
	 -e ORACLE_CHARACTERSET=AL32UTF8 \
	 oracle/database:18.4.0-xe > /dev/null

docker:
	@docker build -t go-api:$(VERSION) .

test:
	@go test -v ./...

generate:
	@go generate ./...

tools:
	@cat tools.go | grep _ | awk -F'"' '{print $$2}' | xargs -tI % go install %

.PHONY: clean
clean:
	@go clean
	@rm -f "$(BINARY_NAME)"
