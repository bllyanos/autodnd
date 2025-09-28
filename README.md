# DND - Do Not Disturb Mode Toggle

A simple Go utility for toggling mako notification "Do Not Disturb" mode on Linux systems.

## What it does

This tool manages an `auto_dnd` mode for the mako notification daemon:

- **Toggles** the `auto_dnd` mode using `makoctl`
- **Checks** if DND is currently active or inactive
- **Auto-installs** the required configuration to your mako config if missing
- **Sends welcome notification** when DND is deactivated

When DND mode is active, notifications become invisible but are still received in the background.

## Prerequisites

- [mako](https://github.com/emersion/mako) notification daemon
- `makoctl` command available in PATH
- `notify-send` for welcome notifications
- Go 1.24.5+ for building from source

## Installation

### Build from source

```bash
git clone <repository-url>
cd dnd
go build -o dnd .
```

### Usage

```bash
./dnd
```

The tool will:
1. Check if the `auto_dnd` mode is configured in your mako config
2. Add the configuration automatically if missing
3. Toggle the DND mode
4. Report the current status

## How it works

The tool adds this configuration to `~/.config/mako/config`:

```ini
# --generated auto_dnd
[mode=auto_dnd]
invisible=1
# --end
```

This creates an `auto_dnd` mode where notifications are invisible but still logged.

## Development

```bash
# Run directly
go run .

# Format code
go fmt ./...

# Run tests
go test ./...

# Clean dependencies
go mod tidy
```

## License

[Add your license here]