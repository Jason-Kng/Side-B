# Side-B

## Project

Go TUI for browsing/managing a vinyl record collection via the Discogs API. Very early stage. Bubble Tea + Lip Gloss frontend, SQLite storage, Discogs REST API.

## Build & Run

```sh
go build -o Side-B ./cmd/Side-B/main.go
./Side-B
```

No test suite, no linter config, no CI pipeline. `go vet ./...` and `go build` are the only verification steps.

## Architecture

```
cmd/Side-B/main.go          Entry point: loads .env, opens DB, starts Bubble Tea program
internal/api/requests.go    Discogs API calls (RequestRelease, ReturnRecord)
internal/db/                SQLite layer: connection.go, store.go (schema), records.go (CRUD)
internal/ui/                Bubble Tea TUI: model/view/update + popup sub-models (add, delete)
pkg/types/types.go          Shared data types (Record, Item, Folder, User)
pkg/logger/logger.go        File-based logger writes to debug.log
exp/                        Experimental scratch files (not part of main build)
```

## Key Conventions

- **`.env` required**: App loads `DISCOGS_TOKEN` via `godotenv`. Without it, the app crashes on startup.
- **DB path is hardcoded**: `./Side-B.db` relative to working directory in `cmd/Side-B/main.go`.
- **Bubble Tea v2**: Uses `charm.land/bubbletea/v2` (not v1). API differs from v1 — `View()` returns `tea.View`, not `string`.
- **Elm architecture**: Model → Update → View. All UI state changes go through `Update` returning `tea.Cmd`.
- **`internal/` is private**: Standard Go layout. Nothing in `internal/` is importable outside this module.
- **`tea.Model` interface**: Sub-models (AddModel, DeleteModel) must implement `Init()`, `Update()`, and `View()` even if `Init`/`View` are unused stubs — required because `Update` returns `tea.Model`.
- **Message types**: Use `PopupCancelledMsg` for dismiss/escape, `ErrorMsg` for errors, `RecordAddedMsg`/`RecordDeletedMsg` for successful operations. Never reuse success messages for cancellation.

## Known Issues

- No `.env` existence check — app crashes with `log.Fatal` if `.env` is missing.
