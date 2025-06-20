# CloudVault Design Document

CloudVault is a command-line tool and optional web UI for synchronizing files
across multiple cloud storage providers. The project is structured to make it
simple to extend support for new providers while keeping the core sync engine
agnostic of any specific API implementation.

## Project Layout

```
cmd/         - entry point for the CLI (implemented in Go)
adapters/    - provider adapters (e.g. gdrive.go, dropbox.go)
engine/      - core sync logic and conflict resolution
watcher/     - cross-platform filesystem watching
config/      - configuration loading and saving
store/       - persistent metadata (SQLite)
utils/       - helpers for encryption, logging and I/O
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
- Use a private `.cloudvault/` directory for storing credentials and sync state.
- Support a `.syncignore` file to exclude paths from synchronization.
- Run `cloudvault login icloud --region <region>` to authenticate via an
  interactive prompt that stores encrypted credentials locally.

## Example Commands

```
cloudvault login icloud --region cn
cloudvault login icloud --region global
cloudvault sync --src icloud/cn --dst icloud/global
cloudvault daemon --config ./sync.yaml
```

