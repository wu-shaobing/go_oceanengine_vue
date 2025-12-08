# WARP.md

This file provides guidance to WARP (warp.dev) when working with code in this repository.

## Project overview

OceanEngine is an ad management platform built on the 巨量引擎 (OceanEngine) Marketing API. It consists of a Go backend (`backend/`), a Vue 3 frontend (`frontend/`), an additional SDK-based API server (`server/`), and vendored SDKs (`sdk/`, `juliang/`).

Read `README.md` and `docs/README.md` first for a human-oriented introduction, architecture diagram, and links to detailed backend/frontend design docs.

## Common commands

All commands assume the repo root is `/Users/wushaobing911/Desktop/oceanengine`.

### Backend (main Go service in `backend/`)

From repo root:

- Install Go deps (if not using Docker):
  - `cd backend && go mod download`
- Local development (no Docker):
  - Run DB migrations: `cd backend && go run cmd/migrate/main.go`
  - Start HTTP server: `cd backend && go run cmd/server/main.go`
- Using `Makefile` helpers (in `backend/`):
  - Build binary: `cd backend && make build`
  - Run server: `cd backend && make run`
  - Dev mode with hot reload (requires `air`): `cd backend && make dev`
  - Run tests: `cd backend && make test` (equivalent to `go test -v ./...`)
  - Lint (requires `golangci-lint`): `cd backend && make lint`
  - Format: `cd backend && make fmt`
  - Tidy deps: `cd backend && make mod`
- Go testing examples:
  - All tests with coverage: `cd backend && go test -coverprofile=coverage.out ./... && go tool cover -html=coverage.out`
  - Single package: `cd backend && go test ./internal/app/advertiser`
  - Single test by name: `cd backend && go test ./internal/app/advertiser -run "TestAdvertiser.*"`
- Docker / Compose:
  - One-click deploy (recommended for local full stack): `./deploy.sh` (runs backend `docker compose`, migrations, and seeds data)
  - Directly from `backend/`:
    - Start stack: `cd backend && docker compose up -d`
    - Stop stack: `cd backend && docker compose down`
    - Tail app logs: `cd backend && docker compose logs -f app`

### Frontend (Vue 3 app in `frontend/`)

From repo root:

- Install deps: `cd frontend && npm install`
- Dev server: `cd frontend && npm run dev` (Vite dev server on port 3000, proxies `/api` to `http://localhost:8080`)
- Production build: `cd frontend && npm run build` (runs `vue-tsc` type-check + Vite build)
- Preview built app: `cd frontend && npm run preview`
- Lint: `cd frontend && npm run lint`
- Tests (Vitest; scripts are referenced in docs and `vitest.config.ts`):
  - Run unit tests (if `"test"` script exists): `cd frontend && npm run test`
  - Run Vitest UI (if `"test:ui"` script exists): `cd frontend && npm run test:ui`
  - Example: run a single test file via Vitest directly: `cd frontend && npx vitest src/views/DashboardView.test.ts --runTestsByPath`

### SDK-based API server in `server/`

This is a separate HTTP API service built directly on the vendored SDK in `sdk/` (used mainly for SDK-style deployments like Tencent Cloud).

Key commands (from `server/README.md`):

- Local run:
  - `cd server && cp .env.example .env`
  - Edit `.env` to set `OCEANENGINE_APP_ID` / `OCEANENGINE_APP_SECRET`.
  - `cd server && go mod tidy`
  - `cd server && go run ./cmd/api`
- Example Docker build and run:
  - `cd server && docker build -f deployments/docker/Dockerfile -t oceanengine-api:latest .`
  - `docker run -d -p 8080:8080 -e OCEANENGINE_APP_ID=... -e OCEANENGINE_APP_SECRET=... oceanengine-api:latest`

### Production deployment

#### HTTPS/SSL setup (Let's Encrypt)

Nginx config template and SSL setup script are in `backend/deployments/nginx/`:

```bash
# On production server (requires root)
cd backend/deployments/nginx
sudo ./setup-ssl.sh your-domain.com admin@example.com
```

This script:
1. Installs nginx and certbot
2. Configures Nginx with HTTPS + HTTP→HTTPS redirect
3. Requests Let's Encrypt certificate
4. Sets up automatic certificate renewal

After SSL is configured:
- Copy frontend build: `cp -r frontend/dist/* /var/www/oceanengine/`
- Update OceanEngine OAuth callback URL to `https://your-domain.com/auth/callback`

#### Change admin password

Default admin credentials are `admin / admin123`. **Change immediately after deployment:**

```bash
# Interactive mode (prompts for password)
cd backend && ./scripts/change-admin-password.sh

# Or non-interactive
cd backend && ./scripts/change-admin-password.sh "YourSecurePass123"
```

Password requirements: 8-32 chars, must contain uppercase, lowercase, and digit.

## High-level architecture

### Overall system

High-level architecture (see `docs/README.md` for diagram):

- Browser-based frontend (Vue 3 + Vite) served behind Nginx or similar reverse proxy.
- Backend REST API (Go + Gin) providing business logic, persistence, auth, and integrations with OceanEngine.
- MySQL as primary data store; Redis for caching and some queue-like features.
- OceanEngine Marketing API as the external ad platform.

The main business entrypoints are in `backend/` (for this project’s application) and `frontend/` (admin UI). The `server/` and `sdk/` trees are vendored SDK and related services.

### Backend (`backend/`)

Backend is a modular Go service with the following major layers (see `backend/` section in root `README.md` and `docs/backend/*.md` for detailed docs):

- Entrypoints (`backend/cmd/`):
  - `cmd/server/`: main HTTP server binary.
  - `cmd/migrate/`: database migrations/initialization.
  - `cmd/task/`: scheduled/cron-like tasks.
- Configuration (`backend/config/`):
  - YAML + environment-variable-driven config, loaded via Viper.
  - Sensitive values are expected via environment variables or `.env` (see `backend/.env.example` and `config/settings.example.yml`).
- Application layer (`backend/internal/app/`):
  - Business modules such as advertisers, campaigns, creatives, reports, 千川电商、星图达人、本地推、RBAC auth, etc.
  - Each module typically exposes service structs and handlers used by routers.
- HTTP layer:
  - `backend/internal/router/`: route registration, grouping, versioning (e.g. `/api/v1/...`).
  - `backend/internal/middleware/`: cross-cutting middleware (JWT auth, CORS, rate limiting, logging, recovery, etc.).
- Infrastructure / shared packages (`backend/pkg/`):
  - `auth/`: JWT auth helpers and RBAC integration.
  - `cache/`: Redis wrappers, common caching utilities.
  - `database/`: GORM setup, connection management, migrations.
  - `oceanengine/`: **project-oriented wrapper around the full OceanEngine SDK**, exposing simplified methods for core ad operations (advertisers, campaigns, creatives, reports, 千川、本地推、星图等). This is the recommended API for most backend business code.
- Scripts & deployments:
  - `backend/scripts/`: operational scripts.
  - `backend/deployments/`: Docker/Kubernetes manifests for deploying the backend stack.

**Routing & auth flow (backend):**

- HTTP requests enter via Gin router in `internal/router/`, passing through middleware for logging, CORS, JWT validation, and rate limiting.
- Handlers in `internal/app/*` modules use `pkg/database` for persistence, `pkg/cache` for Redis, and `pkg/oceanengine` for interacting with the external Marketing API.
- RBAC and user/session management are implemented in system management modules (users, roles, menus, operation logs) as described in `docs/backend/*.md`.

### Frontend (`frontend/`)

Vue 3 SPA built with Vite and TypeScript, following a conventional feature-based layout (see `docs/frontend/*.md` for more on component/state design):

- `src/api/`: Axios-based API wrappers for backend endpoints.
- `src/components/`: Reusable UI components.
- `src/composables/`: Vue Composition API utilities (shared logic hooks).
- `src/router/`: Vue Router configuration for major views (e.g., advertisers, campaigns, creatives, reports, system settings).
- `src/stores/`: Pinia stores for global and domain-specific state (auth, layout, current advertiser, filters, etc.).
- `src/views/`: Page-level components corresponding to main business features (广告主管理、广告系列、广告组、创意、素材库、数据报表、系统管理等).
- `src/utils/`: Utility helpers (date/time formatting via `dayjs`, number formatting, chart helpers, etc.).

Bundling & performance (see `vite.config.ts` and `docs/PROJECT_RUNNABLE_ANALYSIS.md`):

- Vite with aliases (`@` → `frontend/src`).
- Build split into `vue-vendor` and `chart-vendor` chunks.
- Terser used to strip `console` and `debugger` in production.

### Vendored SDK and API server (`sdk/`, `juliang/`, `server/`)

- `sdk/` and `sdk/marketing-api/`:
  - Vendored upstream Go SDK for the OceanEngine Marketing API (including 千川、星图、本地推、企业号、服务市场等模块).
  - `sdk/marketing-api/README.md` and linked docs describe full API surface, request/response models, and examples.
  - `sdk/go.mod` and `server/go.mod` show this is treated as a Go module, with telemetry helpers.
- `backend/pkg/oceanengine/`:
  - Thin, project-specific wrapper around the full SDK, simplifying common flows (OAuth, advertisers, campaigns, creatives, reports, etc.) and adding retry/error handling tuned for this app.
  - Prefer using this from backend business code instead of importing `sdk/marketing-api` directly, unless the wrapper is missing a needed API.
- `server/`:
  - A standalone API server built directly on the vendored SDK (module `github.com/bububa/oceanengine/server` with `replace` pointing at `../sdk`).
  - Provides a smaller HTTP surface for OAuth and a subset of advertiser/report operations; used for alternative deployment targets (Docker, Kubernetes, Tencent Cloud SCF/TKE/CVM).
- `juliang/`:
  - Additional upstream documentation and example index pages for the SDK; useful as reference for long-tail Marketing API methods but not part of the main app runtime.

### Documentation (`docs/`)

`docs/` contains authoritative design docs and should be consulted before making large changes:

- `docs/README.md`: overall architecture, tech stack, and module list.
- `docs/backend/*.md`: backend directory structure, database and API design, auth, middleware, caching, logging, deployment.
- `docs/frontend/*.md`: frontend project structure, component patterns, state management, API integration, routing, and style guide.
- `docs/PROJECT_RUNNABLE_ANALYSIS.md`: up-to-date end-to-end runbook and validation report (startup modes, env vars, security hardening, test commands, and performance notes).

Future Warp instances should rely on these docs for deeper architectural questions rather than inferring structure solely from the code layout.