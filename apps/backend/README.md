# Backend (`apps/backend`)

Go HTTP API for the monorepo (Echo + Postgres + Redis).

## Quick commands

```bash
cp .env.sample .env   # once
task deps:up          # Postgres :5432, Redis :6380
task run              # http://localhost:8080
```

Useful endpoints:

- `GET /status` — health
- `GET /docs` — OpenAPI UI

## Docs

Full project documentation lives at the repository root:

- [README](../../README.md)
- [Architecture](../../docs/ARCHITECTURE.md)
- [Development](../../docs/DEVELOPMENT.md)
- [Configuration](../../docs/CONFIGURATION.md)
- [Contributing](../../CONTRIBUTING.md)

## Layout

```text
cmd/go-taskForge/   # main
internal/             # application code
static/               # OpenAPI assets
templates/emails/     # HTML from @taskForge/emails
docker-compose.yml
Taskfile.yml
```
