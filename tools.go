package e621

//go:generate go tool oapi-codegen -config cfg.yml https://e621.wiki/openapi.yaml
//go:generate go mod tidy
//go:generate go fmt ./...
