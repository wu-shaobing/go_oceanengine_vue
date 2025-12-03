---
name: dev-stop
description: Stop all running development servers (backend and frontend)
---

# Stop Development Environment

Gracefully stop all development servers and clean up processes.

## Usage

```
/dev-stop
```

## What It Does

1. **Find Running Processes**
   - Searches for `go run` processes (backend)
   - Finds `vite` dev server processes (frontend)
   - Lists all found processes

2. **Graceful Shutdown**
   - Sends SIGTERM to allow cleanup
   - Waits for graceful shutdown (5s timeout)
   - Force kills if necessary (SIGKILL)

3. **Cleanup**
   - Removes temporary files
   - Clears lock files if present
   - Reports final status

## Manual Stop

If automatic stop fails:

```bash
# Find processes
ps aux | grep "go run"
ps aux | grep "vite"

# Kill by PID
kill <PID>

# Or kill by port
lsof -ti:8080 | xargs kill
lsof -ti:3000 | xargs kill
```

## Related Commands

- `/dev-start` - Start development environment
- `/dev-restart` - Restart all servers
