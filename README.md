# FestMap Backend

Go backend for FestMap. It serves JSON content for the frontend and can optionally host static frontend files.

## Status

- Language: Go
- Main HTTP port: 8080
- Current API endpoint: `/api/text`
- Primary data source: `homePage.json`

## SDD Workflow

This backend follows Spec-Driven Development.

Before implementing or changing behavior:

1. Define product intent in `mission.md`.
2. Add or update feature description in `features.md`.
3. Place work in `roadMap.md` milestone.
4. Confirm technical constraints in `techStack.md`.

Implementation starts only after the spec entry exists and is approved.

## Documentation Index

- Vision and requirements: `mission.md`
- Feature tracking: `features.md`
- Milestone planning: `roadMap.md`
- Technology constraints: `techStack.md`

## Project Structure

- `main.go` - server bootstrap, routing setup, optional tracer initialization
- `mainPage.go` - handler for `/api/text`, reads `homePage.json`
- `homePage.json` - local data payload for API response
- `jaeger.go` - OpenTelemetry/Jaeger setup (currently not enabled in `main.go`)
- `dataBase.go` - database-related code (currently not enabled in `main.go`)
- `docker-compose.yml` - backend and Jaeger stack
- `Dockerfile` - container build definition

## Prerequisites

- Go 1.23+
- Optional: Docker + Docker Compose

## Run Locally

From backend directory:

```bash
go run .
```

Server starts on:

- `http://localhost:8080`

## API

### GET /api/text

Returns the contents of `homePage.json` as `application/json`.

Example:

```bash
curl http://localhost:8080/api/text
```

## Frontend Integration

Frontend proxy is expected to forward:

- `/api` -> `http://localhost:8080`

This allows frontend calls to `/api/text` without CORS setup in local development.

## Static Frontend Hosting (Optional)

`main.go` currently registers:

- `/` -> static file server at `./frontend1/dist`

If this directory is missing, API routes can still work, but root file serving may return 404.

## Docker

Build image:

```bash
docker build -t backend .
```

Run stack with Jaeger:

```bash
docker compose up --build
```

Useful endpoints:

- Backend: `http://localhost:8080`
- Jaeger UI: `http://localhost:16686`

## Observability Notes

Jaeger initialization exists but is currently commented out in `main.go`:

- `initJaegerMain()` is disabled

Enable it only after confirming exporter settings and runtime expectations.

## Common Troubleshooting

### `listen tcp :8080: bind: address already in use`

Another process is using port 8080.

```bash
lsof -nP -iTCP:8080 -sTCP:LISTEN
kill <PID>
```

### `Could not open JSON file`

`homePage.json` was not found from current working directory.

- Run `go run .` from backend root.
- Ensure `homePage.json` exists in the same directory.

### Root route returns 404

Expected if `./frontend1/dist` does not exist. This does not block `/api/text`.

## Next SDD Step

Define the first backend milestone in `roadMap.md` with clear acceptance criteria for:

- API contract for `/api/text`
- error handling behavior
- optional static file serving expectations
