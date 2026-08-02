# Development guide

Day-to-day workflows for this monorepo.

## First-time setup

```bash
# repo root
bun install

cd apps/backend
cp .env.sample .env
# set at least:
# - BOILERPLATE_DATABASE.PASSWORD=postgres (matches docker-compose)
# - BOILERPLATE_REDIS.ADDRESS=localhost:6380
# - placeholders OK for Clerk/Resend/New Relic on pure local smoke tests

task deps:up
task run
```

Confirm:

```bash
curl -i http://localhost:8080/status
```

## Backend Taskfile

Run all commands from `apps/backend` (or `task -d apps/backend <task>` from the repo root).

| Task | Purpose |
| --- | --- |
| `task help` | List tasks |
| `task deps:up` | `docker compose up -d` (Postgres + Redis) |
| `task deps:down` | Stop compose stack |
| `task deps:logs` | Follow container logs |
| `task run` | `go run ./cmd/go-taskForge` |
| `task migrations:new name=...` | Create a new Tern SQL migration |
| `task migrations:up` | Apply migrations (interactive confirm) |
| `task tidy` | Format + `go mod tidy` + verify |

Default migration DSN:

```text
postgres://postgres:postgres@localhost:5432/taskForge?sslmode=disable
```

Override with `BOILERPLATE_DB_DSN` when invoking Task if needed.

## Docker services

Defined in `apps/backend/docker-compose.yml`:

| Service | Container | Host port | Credentials / notes |
| --- | --- | --- | --- |
| Postgres 15 | `taskForge-postgres` | `5432` | user/password/db: `postgres` / `postgres` / `taskForge` |
| Redis 7 | `taskForge-redis` | `6380` → `6379` | no password |

If port `5432` or `6380` is busy, change the **left** side of the port mapping and update `.env` accordingly.

## TypeScript packages

From the **repository root**:

```bash
bun install
bun run build          # turbo build
bun run dev            # turbo persistent watchers
bun run typecheck
bun run clean
```

Workspace packages are declared in root `package.json` as `packages/*`.

### `@taskForge/zod`

```bash
cd packages/zod
bun run build   # tsc + tsc-alias → dist/
bun run dev     # watch
```

Add new schemas under `src/`, export them from `src/index.ts`.

### `@taskForge/openapi`

Depends on built `@taskForge/zod`.

```bash
cd packages/zod && bun run build
cd ../openapi
bun run build
bun run gen     # writes openapi.json (package + apps/backend/static)
```

Add contracts under `src/contracts/`, register them in `src/contracts/index.ts`, then regenerate.

### `@taskForge/emails`

```bash
cd packages/emails
bun run dev      # React Email preview :3001
bun run export   # HTML → apps/backend/templates/emails
```

After exporting, restart or redeploy the Go API so it picks up template file changes (templates are read from disk / embeds depending on the email helper implementation).

## Adding a new API endpoint (recommended path)

1. **Schema** — add Zod types in `packages/zod`.
2. **Contract** — add a ts-rest route in `packages/openapi/src/contracts`.
3. **Generate** — `bun run gen` in `packages/openapi`.
4. **Repository** — implement persistence in `apps/backend/internal/repository` (and migrations).
5. **Service** — business logic in `internal/service`.
6. **Handler** — HTTP adapter in `internal/handler`.
7. **Router** — register on `/api/v1` in `internal/router`.
8. **Tests** — use `internal/testing` helpers + Testcontainers where useful.

## Migrations

Create:

```bash
cd apps/backend
task migrations:new name=create_users
```

Edit the generated SQL between Tern’s up/down separators:

```sql
-- migrate up statements here

---- create above / drop below ----

-- migrate down statements here
```

Apply locally:

```bash
task migrations:up
```

Empty up-steps will fail with Tern’s `no sql in forward migration step` — always put real SQL in new migrations.

## Linting / formatting

- Go: `task tidy` (fmt) and `apps/backend/.golangci.yml` when you run golangci-lint in CI or locally.
- TS: add package-level lint scripts as the frontend/packages mature; Turbo already has `lint` / `format` task slots.

## Common failures

| Symptom | Likely cause | Fix |
| --- | --- | --- |
| `bind: address already in use` on `:8080` | Another process on 8080 | Free the port or change `BOILERPLATE_SERVER.PORT` |
| Redis container fails on `6379` | Host Redis already bound | Use mapped host port `6380` (default in this repo) |
| Config validation failed (`CORS` / `Redis`) | Env key mismatch | Match `.env.sample` keys exactly (`CORS_ALLOWED_ORIGINS`, top-level `BOILERPLATE_REDIS.ADDRESS`) |
| Tern: no sql in forward step | Empty migration file | Fill up SQL or remove empty migration |
| OpenAPI gen can’t resolve `@taskForge/zod` | Zod not built / install stale | `bun install` then build zod first |

## Suggested editor setup

- Go: official Go extension / gopls
- TypeScript: workspace TypeScript version from packages
- Bun: prefer `bun` over npm/yarn/pnpm for this repo (see `.cursor/rules`)
