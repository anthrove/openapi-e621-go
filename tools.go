package e621

//go:generate go tool oapi-codegen -config cfg.yml https://raw.githubusercontent.com/anthrove/E621OpenAPI/refs/heads/master/openapi.yaml
//go:generate go mod tidy
//go:generate go fmt ./...
