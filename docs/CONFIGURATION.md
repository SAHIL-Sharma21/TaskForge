# Configuration reference

Backend configuration is loaded in `apps/backend/internal/config` via:

1. `godotenv/autoload` (loads `.env` if present)
2. Koanf env provider with prefix `BOILERPLATE_`
3. Struct validation (`go-playground/validator`)

Keys use `.` nesting in env names. Example:

```bash
BOILERPLATE_DATABASE.HOST=localhost
```

becomes `config.Database.Host`.

Copy the sample file:

```bash
cp apps/backend/.env.sample apps/backend/.env
```

---

## Primary

| Variable | Example | Required | Description |
| --- | --- | --- | --- |
| `BOILERPLATE_PRIMARY.ENV` | `local` | yes | Runtime environment. When **not** `local`, Tern migrations run on startup. |

Suggested values: `local`, `development`, `staging`, `production`.

---

## Server

| Variable | Example | Required | Description |
| --- | --- | --- | --- |
| `BOILERPLATE_SERVER.PORT` | `8080` | yes | HTTP listen port |
| `BOILERPLATE_SERVER.READ_TIMEOUT` | `30` | yes | Read timeout (seconds) |
| `BOILERPLATE_SERVER.WRITE_TIMEOUT` | `30` | yes | Write timeout (seconds) |
| `BOILERPLATE_SERVER.IDLE_TIMEOUT` | `60` | yes | Idle timeout (seconds) |
| `BOILERPLATE_SERVER.CORS_ALLOWED_ORIGINS` | `http://localhost:3000` | yes | Allowed CORS origins (comma-separated if multiple, depending on unmarshal support) |

---

## Database

| Variable | Example | Required | Description |
| --- | --- | --- | --- |
| `BOILERPLATE_DATABASE.HOST` | `localhost` | yes | Postgres host |
| `BOILERPLATE_DATABASE.PORT` | `5432` | yes | Postgres port |
| `BOILERPLATE_DATABASE.USER` | `postgres` | yes | DB user |
| `BOILERPLATE_DATABASE.PASSWORD` | `postgres` | no* | DB password (`*` empty allowed by struct; Compose uses `postgres`) |
| `BOILERPLATE_DATABASE.NAME` | `taskForge` | yes | Database name |
| `BOILERPLATE_DATABASE.SSL_MODE` | `disable` | yes | libpq SSL mode |
| `BOILERPLATE_DATABASE.MAX_OPEN_CONNS` | `25` | yes | Pool max open |
| `BOILERPLATE_DATABASE.MAX_IDLE_CONNS` | `25` | yes | Pool max idle |
| `BOILERPLATE_DATABASE.CONN_MAX_LIFETIME` | `300` | yes | Conn max lifetime (seconds) |
| `BOILERPLATE_DATABASE.CONN_MAX_IDLE_TIME` | `300` | yes | Conn max idle time (seconds) |

Docker Compose defaults (local):

```text
user=postgres password=postgres db=taskForge port=5432
```

---

## Redis

| Variable | Example | Required | Description |
| --- | --- | --- | --- |
| `BOILERPLATE_REDIS.ADDRESS` | `localhost:6380` | yes | `host:port` for go-redis / Asynq |

> Use `host:port` (not a `redis://` URL) for the current client options.

Compose maps host **6380** → container **6379**.

---

## Auth (Clerk)

| Variable | Example | Required | Description |
| --- | --- | --- | --- |
| `BOILERPLATE_AUTH.SECRET_KEY` | `sk_test_...` | yes | Clerk secret key |

Replace the sample `"secret"` before enabling protected routes in real environments.

---

## Integrations

| Variable | Example | Required | Description |
| --- | --- | --- | --- |
| `BOILERPLATE_INTEGRATION.RESEND_API_KEY` | `re_...` | yes | Resend API key for outbound email |

---

## Observability

### Service identity

| Variable | Example | Required | Description |
| --- | --- | --- | --- |
| `BOILERPLATE_OBSERVABILITY.SERVICE_NAME` | `taskForge` | yes* | Service name in telemetry |
| `BOILERPLATE_OBSERVABILITY.ENVIRONMENT` | `development` | yes* | Telemetry environment label |

### Logging

| Variable | Example | Required | Description |
| --- | --- | --- | --- |
| `BOILERPLATE_OBSERVABILITY.LOGGING.LEVEL` | `debug` | yes* | Log level |
| `BOILERPLATE_OBSERVABILITY.LOGGING.FORMAT` | `console` | yes* | `console` or JSON-oriented formats supported by logger setup |
| `BOILERPLATE_OBSERVABILITY.LOGGING.SLOW_QUERY_THRESHOLD` | `100ms` | no | Slow query warning threshold |

### New Relic

| Variable | Example | Required | Description |
| --- | --- | --- | --- |
| `BOILERPLATE_OBSERVABILITY.NEW_RELIC.LICENSE_KEY` | `xxxxxxxx` | yes* | NR ingest license |
| `BOILERPLATE_OBSERVABILITY.NEW_RELIC.APP_LOG_FORWARDING_ENABLED` | `true` | no | Forward app logs |
| `BOILERPLATE_OBSERVABILITY.NEW_RELIC.DISTRIBUTED_TRACING_ENABLED` | `true` | no | Distributed tracing |
| `BOILERPLATE_OBSERVABILITY.NEW_RELIC.DEBUG_LOGGING` | `false` | no | Agent debug logs |

### Health checks

| Variable | Example | Required | Description |
| --- | --- | --- | --- |
| `BOILERPLATE_OBSERVABILITY.HEALTH_CHECKS.ENABLED` | `true` | no | Enable health subsystem |
| `BOILERPLATE_OBSERVABILITY.HEALTH_CHECKS.INTERVAL` | `30s` | yes* | Interval |
| `BOILERPLATE_OBSERVABILITY.HEALTH_CHECKS.TIMEOUT` | `5s` | yes* | Per-check timeout |
| `BOILERPLATE_OBSERVABILITY.HEALTH_CHECKS.CHECKS` | `database,redis` | no | Checks to run |

\*Required by observability struct validation when that subtree is active; see `internal/config/observability.go`. Note that `LoadConfig` currently applies defaults for observability after load — prefer keeping `.env` aligned with defaults until that path is cleaned up.

---

## Task / CLI overrides

| Variable | Used by | Description |
| --- | --- | --- |
| `BOILERPLATE_DB_DSN` | `task migrations:up` | Full Postgres URL for Tern CLI |

Example:

```bash
BOILERPLATE_DB_DSN='postgres://postgres:postgres@localhost:5432/taskForge?sslmode=disable' task migrations:up
```

---

## Production checklist

- [ ] `PRIMARY.ENV` is not `local`
- [ ] Strong DB password + `SSL_MODE` appropriate for your host
- [ ] Redis reachable and secured (ACL/TLS as needed; update client options if you add auth)
- [ ] Real Clerk secret and Resend API key
- [ ] Real New Relic license (or disable NR integrations deliberately)
- [ ] CORS origins locked to real frontends
- [ ] Secrets injected by your platform (not baked into images)
- [ ] Migrations reviewed and applied (startup migrate and/or CI/CD release job)
