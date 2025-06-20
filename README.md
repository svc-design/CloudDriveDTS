# CloudVault

CloudVault aims to provide a unified tool for synchronizing files across
multiple cloud storage providers. The core CLI is written in Go and currently
supports simple login, sync and daemon commands. It starts as a minimal tool
for one-way synchronization from a local directory to Google Drive and is
planned to evolve into a multi-cloud, bidirectional solution with a web-based
management interface.

Credentials entered during `cloudvault login` are encrypted using AES-256 with
an encryption key supplied interactively. The resulting encrypted blob is stored
under `~/.cloudvault/<provider>/<region>/credentials.json`.

To build the CLI from source run:

```bash
go build ./cmd/cloudvault
```

For details about the architecture and planned phases, see
[docs/design.md](docs/design.md).
