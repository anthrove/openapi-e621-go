package main

//go:generate go tool oapi-codegen -config cfg.yml openapi.yaml
//go:generate go mod tidy
//go:generate go fmt ./...
