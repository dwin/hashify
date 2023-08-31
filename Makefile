include Makefile.defs

.PHONY: openapi
openapi:
	go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@v1.12.4
	oapi-codegen -config resources/openapi/model.config.yaml openapi.yaml
	oapi-codegen -config resources/openapi/server.config.yaml openapi.yaml

.PHONE: mock
generate: mock
	go generate ./...

.PHONY: mock
mock:
	go install github.com/vektra/mockery/v2@v2.20.0

run:
	go run cmd/hashify/main.go