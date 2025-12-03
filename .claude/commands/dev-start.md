---
name: dev-start
description: Quick start development environment with backend and frontend running in parallel
---

# Start Development Environment

Quickly start both backend and frontend development servers in parallel with proper environment setup.

## Usage

```
/dev-start
```

## What It Does

1. **Environment Check**
   - Verifies Go and Node.js are installed
   - Checks for required environment variables
   - Validates `.env` files exist

2. **Dependency Check**
   - Checks if `go.mod` dependencies are current
   - Verifies `node_modules` are installed
   - Offers to install if missing

3. **Start Services**
   - Backend on port 8080 (Go + Gin)
   - Frontend on port 3000 (Vite dev server)
   - Both running concurrently

4. **Health Check**
   - Waits for servers to be ready
   - Validates API connectivity
   - Opens browser to frontend

## Prerequisites

### Backend `.env` file
```bash
APP_ID=your_qianchuan_app_id
APP_SECRET=your_qianchuan_app_secret
COOKIE_SECRET=random_32_byte_secret
PORT=8080
GIN_MODE=debug
CORS_ORIGIN=http://localhost:3000
COOKIE_DOMAIN=localhost
COOKIE_SECURE=false
SESSION_NAME=qianchuan_session
```

### Frontend `.env` file
```bash
VITE_API_BASE_URL=http://localhost:8080/api
VITE_OAUTH_APP_ID=your_app_id
VITE_OAUTH_REDIRECT_URI=http://localhost:3000/auth/callback
VITE_APP_TITLE=千川SDK管理平台
```

## Troubleshooting

### Port Already in Use
```
Error: Port 8080 already in use
Solution: /dev-stop or manually kill the process
```

### Missing Dependencies
```
Error: go.mod dependencies not installed
Solution: Run `make install-backend`
```

### Environment Variables Missing
```
Error: Required env vars not found
Solution: Copy .env.example to .env and configure
```

## Related Commands

- `/dev-stop` - Stop all development servers
- `/dev-backend` - Start only backend
- `/dev-frontend` - Start only frontend
- `/dev-logs` - View server logs
