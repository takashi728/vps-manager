# VPS Manager

A high-performance, minimal TUI tool for managing VPS hardening and SSH keys, written in Go.

## Features

- **TUI Navigation**: Built with Bubble Tea.
- **SSH Management**: Secure key authentication and command execution.
- **Local Key Discovery**: Automatically finds local SSH private keys.
- **Minimal Footprint**: Compiled binary.

## Setup

1. Clone the repo: `git clone https://github.com/takashi728/vps-manager`
2. Build the project:
   ```bash
   cd vps-manager
   go build -ldflags="-s -w" -o vps-manager main.go
   ```
3. Run: `./vps-manager`
