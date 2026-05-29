# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Commands

```bash
# Build
go build ./...

# Run
go run main.go

# Run with vendored dependencies
go run -mod=vendor main.go

# Update vendor directory after changing go.mod
go mod vendor
```

## Architecture

This is a Go project (module `github.com/Khnzh/HassApi`) that bridges a **Telegram bot** with the **Home Assistant** (HASS) REST API.

- **`main.go`** — Entry point. Calls `bot.ListenAndServe()`. Also contains commented-out standalone CLI code for directly updating a Home Assistant automation's trigger time via the HA REST API (`/api/config/automation/config/{id}`).

- **`bot/bot.go`** — Package `bot`. Contains `ListenAndServe()`, which runs a long-polling Telegram bot using `go-telegram-bot-api/v5`. Currently echoes messages back; intended to be extended to dispatch Home Assistant API calls.

- **Home Assistant integration** — The HA API is at `http://192.168.1.18:8123`. Auth uses a long-lived token from the `HASS_TOKEN` environment variable. The helper functions `hassGet` / `hassPost` in `main.go` wrap HTTP calls with the Bearer token header.

- **Dependencies** are vendored under `vendor/`. The only external dependency is `github.com/go-telegram-bot-api/telegram-bot-api/v5`.

## Environment Variables

| Variable | Purpose |
|---|---|
| `HASS_TOKEN` | Home Assistant long-lived access token |

The Telegram bot token is currently hardcoded in `bot/bot.go` — it should be moved to an environment variable.
