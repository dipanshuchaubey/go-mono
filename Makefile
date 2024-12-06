FILES = $(shell find protos -name "*proto")
SERVICE = $(if $(service),$(service),all)

# Replace _ with - in service name
BUILD_NAME = $(echo $$SERVICE | tr '_' '-')

.PHONY: protos
protos:
	protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
	$(FILES)

build:
	@if [ "$(SERVICE)" = "all" ]; then \
		for dir in cmd/*; do \
			echo "Building $$dir"; \
			build_name=$$(basename $$dir | tr '_' '-'); \
			go build -o bin/$$build_name $$dir/main.go; \
			echo "=-=-=-=-=-=-=-=-=-=-=-=-=-=-="; \
		done; \
	else \
		go build -o bin/$(BUILD_NAME) cmd/$(SERVICE)/main.go; \
	fi

run:
	@if [ "$(SERVICE)" = "all" ]; then \
		for dir in cmd/*; do \
			echo "Running $$dir"; \
			go run $$dir/main.go; \
			echo "=-=-=-=-=-=-=-=-=-=-=-=-=-=-="; \
		done; \
	else \
		go run cmd/$(SERVICE)/main.go; \
	fi