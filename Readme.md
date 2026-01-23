# JSON to Go Struct CLI

A command-line tool to convert JSON objects into Go struct definitions.

## Features

- Converts JSON to properly formatted Go structs
- Handles nested objects and arrays
- Converts `snake_case` keys to `PascalCase` field names
- Automatically copies generated struct to clipboard
- Properly aligned and formatted output using `go/format`

## Install

**Requires Go 1.25+**

```bash
go install github.com/samit22/json-go@latest
```

## Usage

```bash
json-go jtg '<json-string>'
```

### Examples

**Simple JSON:**

```bash
json-go jtg '{"name": "John", "age": 30, "email": "john@example.com", "active": true}'
```

**Output:**

```go
type AutoStruct struct {
	Name   string  `json:"name"`
	Age    float64 `json:"age"`
	Email  string  `json:"email"`
	Active bool    `json:"active"`
}
```

**Nested JSON with snake_case:**

```bash
json-go jtg '{"user": {"first_name": "John", "last_name": "Doe"}, "is_premium": true}'
```

**Output:**

```go
type AutoStruct struct {
	User struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
	} `json:"user"`
	IsPremium bool `json:"is_premium"`
}
```

**Array of objects:**

```bash
json-go jtg '[{"id": 1, "name": "Item"}]'
```

## License

Apache License 2.0
