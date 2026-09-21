# teltonika-playground

Codec parsing and connection manager.

## Development container

This repository includes a devcontainer configuration with the latest stable Go image and common tooling.

- Open the folder in VS Code.
- Run **Dev Containers: Reopen in Container**.
- The container runs `go mod tidy` after creation.

## Project layout

- `cmd/teltonika-playground` — application entrypoint
- `internal/app` — internal package code

## Local commands

```bash
go test ./...
go run ./cmd/teltonika-playground
```
