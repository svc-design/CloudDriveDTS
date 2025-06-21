# CloudDriveDTS

CloudDriveDTS aims to provide a private, multi-cloud synchronization engine.
It starts as a minimal command-line tool and will evolve into a flexible
solution supporting bidirectional sync across popular providers. Credentials
are encrypted locally using AES-256 so you remain in control of your data.

"CloudDriveDTS — Own your cloud. Sync your way." 🪧

## Features

- **Cloud authentication management** via OAuth for multiple drives
- **Bidirectional sync engine** for local ↔ cloud and cloud ↔ cloud
- **Conflict resolution** using timestamps, hashes or manual override
- **Watchers & schedulers** for real‑time sync and periodic scans
- **Encrypted local storage** of configuration and tokens
- **Plugin architecture** to add services like WebDAV, S3 or iCloud
- **CLI and optional Web UI** for managing sync tasks

## Project Layout

```text
clouddrivedts/
├── cmd/                # CLI entry point
├── adapter/            # Cloud drive adapters
├── engine/             # Sync engine and change detection
├── watcher/            # Filesystem watcher helpers
├── store/              # State storage (e.g. SQLite)
├── config/             # Config loading and validation
├── crypto/             # AES encryption utilities
├── utils/              # Misc utilities
├── webui/              # Optional web control panel
└── README.md
```

## Example Commands

```bash
clouddrivedts init
clouddrivedts auth gdrive
clouddrivedts auth onedrive
clouddrivedts sync --src ~/Projects --dst gdrive:/Backups --bi-sync
clouddrivedts run --daemon
clouddrivedts status
```

## Example Configuration

`~/.clouddrivedts.yaml`

```yaml
profiles:
  - name: sync_docs
    source: ~/Documents
    target:
      provider: gdrive
      path: /DocumentsBackup
    bidirectional: true
    schedule: "0 */2 * * *"
    ignore:
      - "*.tmp"
      - ".DS_Store"
```

Credentials entered during `clouddrivedts auth` are encrypted with AES-256 and
stored under `~/.clouddrivedts/<provider>/credentials.json`.

To build the CLI from source run:

```bash
go build ./cmd/cloudvault
```

For architectural details, see [docs/design.md](docs/design.md).
