package main

//go:generate go tool oapi-codegen -config cfg.yml ../E621OpenAPI/openapi.yaml
//go:generate go mod tidy
//go:generate go fmt ./...
