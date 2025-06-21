# CloudDriveDTS Design Document

CloudDriveDTS is a command-line tool with an optional web UI for synchronizing
files across multiple cloud storage providers. The architecture is designed to
be provider agnostic so new adapters can be added with minimal effort.

**Slogan:** "CloudDriveDTS — Own your cloud. Sync your way."

## Features

- OAuth-based authentication for various cloud drives
- Bidirectional sync between local directories and remote clouds
- Conflict resolution based on timestamps, hashes or manual selection
- Real-time watchers with optional cron-like schedulers
- AES-encrypted storage of configuration files and tokens
- Plugin system for extending to WebDAV, S3, iCloud and more
- Command-line interface with an optional web dashboard

## Project Layout

```
cmd/         - entry point for the CLI (implemented in Go)
adapter/     - provider adapters (e.g. gdrive.go, dropbox.go)
engine/      - core sync logic and conflict resolution
watcher/     - cross-platform filesystem watching
config/      - configuration loading and saving
store/       - persistent metadata (SQLite)
crypto/      - AES encryption utilities
utils/       - helper functions
webui/       - optional web management interface
```

## Development Phases

1. **MVP**
   - Basic CLI for one-way sync from a local directory to Google Drive.
   - OAuth authentication.
   - Detect additions, deletions and modifications.
   - Support configuration files and logging.

2. **Multi-cloud & Bidirectional Sync**
   - Add adapters for services like OneDrive and Dropbox.
   - Enable bidirectional sync with conflict detection and incremental uploads.
   - Persist sync state in SQLite for reliability.

3. **Visualization & Task Management**
   - Provide a web UI and API server for managing tasks and viewing history.
   - Support multiple sync tasks and audit logs.

4. **Advanced Features & Private Deployment**
   - Encrypted sync using tools like rclone or encfs.
   - Dockerized deployment and support for private cloud backends (WebDAV, S3).
   - Multi-user permissions through the web interface.

## Security Guidelines

- Encrypt all cloud provider access tokens locally (e.g. AES‑256 with
a password-derived key).
- Use a private `.clouddrivedts/` directory for storing credentials and sync state.
- Support a `.syncignore` file to exclude paths from synchronization.
- Run `clouddrivedts auth icloud --region <region>` to authenticate via an
  interactive prompt that stores encrypted credentials locally.

## Example Commands

```
clouddrivedts auth gdrive
clouddrivedts auth onedrive
clouddrivedts sync --src ~/Projects --dst gdrive:/Backups --bi-sync
clouddrivedts run --daemon
clouddrivedts status
```

