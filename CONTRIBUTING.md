# Contributing

Thanks for helping improve this taskForge. This project aims to stay **clear, copyable, and production-leaning** — prefer small, well-documented changes over large unrelated refactors.

## Before you start

1. Read the [README](README.md) and [docs/ARCHITECTURE.md](docs/ARCHITECTURE.md).
2. Open an issue for large features or breaking changes so direction can be aligned early.
3. Use a feature branch off the default branch.

## Development setup

Follow [docs/DEVELOPMENT.md](docs/DEVELOPMENT.md).

Minimum smoke check before a PR:

```bash
bun install
cd apps/backend
cp -n .env.sample .env   # if needed
task deps:up
task run                 # ensure /status responds, then stop
```

If you touch TypeScript packages:

```bash
bun run build
```

If you touch Go modules:

```bash
cd apps/backend
task tidy
go test ./...
```

## Pull request guidelines

- Keep PRs focused (one concern per PR when possible).
- Update docs when behavior, env vars, scripts, or architecture change.
- Do not commit secrets (`.env`, API keys, license keys).
- Match existing style: Go in `internal/` layers; TS packages under `packages/*`.
- Prefer Bun commands for JS (`bun install`, `bun run …`), not npm/yarn.

### Commit messages

Use short, imperative subjects:

```text
add health contract for redis optional check
fix cors env key to match koanf tag
docs: document redis host port mapping
```

## Code map (where to put things)

| Change type | Location |
| --- | --- |
| HTTP route | `apps/backend/internal/router` + `handler` |
| Business logic | `apps/backend/internal/service` |
| SQL / persistence | `apps/backend/internal/repository` + Tern migrations |
| Shared request/response schema | `packages/zod` |
| Public API contract / OpenAPI | `packages/openapi` |
| Email template | `packages/emails` then `bun run export` |
| Docs | `README.md`, `docs/*` |

## Reporting bugs

Include:

- OS and versions (Go, Bun, Docker)
- Steps to reproduce
- Expected vs actual behavior
- Relevant logs (redact secrets)

## Security

Do not file public issues for sensitive vulnerabilities. Contact the maintainers privately if you discover a security problem.

## License

By contributing, you agree that your contributions will be licensed under the [MIT License](LICENSE).
