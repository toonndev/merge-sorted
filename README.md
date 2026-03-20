# merge-sorted

Merges three integer slices into one ascending-sorted slice using the two-pointer technique — no built-in sort functions used.

## Input

| Parameter     | Order      |
|---------------|------------|
| `collection1` | Descending |
| `collection2` | Ascending  |
| `collection3` | Ascending  |

## How it works

1. Reverse `collection1` → ascending
2. Merge result with `collection2` using two-pointer
3. Merge result with `collection3` using two-pointer

## Install dependencies

```bash
go mod download
```

## Run example

```bash
go run cmd/main.go
```

## Run unit tests

```bash
go test ./...
```

## Run tests with coverage report

```bash
go test ./... -cover
```

For a detailed HTML coverage report:

```bash
go test ./... -coverprofile=coverage.out && go tool cover -html=coverage.out
```
