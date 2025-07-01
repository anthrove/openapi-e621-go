# Current Problems

- I had to change a few types of the original swagger
- For some reason the generated code is missing a few types, i had to add them manually
```go
type AccessDeniedReason string
type AccessDeniedSuccess string
type MessageErrorSuccess string
type NotFoundReason string
type NotFoundSuccess string
type WarningRecordType string
```


# TODOS: Improvements

- OpenAPI Docs has a  [Nullable type](https://github.com/oapi-codegen/oapi-codegen?tab=readme-ov-file#generating-nullable-types). We should implement this for everything that is nullable.
- We could also generate a basic Server implementation for mocking purposes, if we need to