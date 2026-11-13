# Side-B Dev TODO

## Context

`dev` is 3 commits ahead of `main`. The last work was a half-finished SQLite
schema migration:

- New schema/table design is done (`internal/db/initSchema.sql`).
- New types are done (`pkg/types/types.go`).
- API now returns `types.ReleaseData` (`internal/api/requests.go`).
- DB CRUD and the Bubble Tea UI were never updated.

The project currently does not compile. Goal of this pass: get back to a
working list/add/delete app using the new schema. Do **not** touch
`oldSide-B.db` and do **not** wire up `collection_items` yet.

Current build error:

```text
internal/ui/commands.go:26:24: cannot use record (variable of type *types.ReleaseData)
as *types.Record value in argument to store.AddRecord
```

---

## 1. Fix the AddRecord data contract

- [ ] In `internal/db/records.go`, change:
  ```go
  func (s *RecordStore) AddRecord(rec *types.Record) error
  ```
  to:
  ```go
  func (s *RecordStore) AddRecord(data *types.ReleaseData) error
  ```
- [ ] Delete the current no-op stub body (returns `nil` without writing).
- [ ] `internal/ui/commands.go:26` should then compile unchanged
  (`store.AddRecord(record)` where `record` is `*types.ReleaseData`).

### Implement `AddRecord` as one transaction

Write `releases`, `artists`, `release_artists`, and `tracks` atomically:

```go
tx, err := s.db.Begin()
if err != nil { return err }
defer tx.Rollback()
// ... all inserts/updates ...
return tx.Commit()
```

Upsert the release (12 columns in the new schema):

```sql
INSERT INTO releases (
  id, master_id, title, artists_sort, year, release_date,
  country, barcode, notes, thumb_url, raw_json, fetched_at
) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
ON CONFLICT(id) DO UPDATE SET
  master_id    = excluded.master_id,
  title        = excluded.title,
  artists_sort = excluded.artists_sort,
  year         = excluded.year,
  release_date = excluded.release_date,
  country      = excluded.country,
  barcode      = excluded.barcode,
  notes        = excluded.notes,
  thumb_url    = excluded.thumb_url,
  raw_json     = excluded.raw_json,
  fetched_at   = excluded.fetched_at;
```

Upsert each artist:

```sql
INSERT INTO artists (id, name, thumb_url) VALUES (?, ?, ?)
ON CONFLICT(id) DO UPDATE SET
  name      = excluded.name,
  thumb_url = excluded.thumb_url;
```

Replace links and tracks for that release (re-adding a release must not
duplicate them):

```sql
DELETE FROM release_artists WHERE release_id = ?;
DELETE FROM tracks WHERE release_id = ?;
```

Then insert every `data.Links` row:

```sql
INSERT INTO release_artists
  (release_id, artist_id, anv, join_str, position)
VALUES (?, ?, ?, ?, ?);
```

And every `data.Tracks` row:

```sql
INSERT INTO tracks
  (release_id, position, title, duration, track_index, type_)
VALUES (?, ?, ?, ?, ?, ?);
```

Notes:
- The schema uses `type_`, not `type`.
- `artists_sort` is `NOT NULL DEFAULT ''`.
- `fetched_at` defaults to `datetime('now')` only if omitted; providing a
  value is fine as long as it is consistent.

---

## 2. Fix the read path

- [ ] In `pkg/types/types.go`, add `Notes string` to `types.Release`
  (the schema already has `releases.notes`).
- [ ] Add a display row type (or redefine `Record`):
  ```go
  type ReleaseSummary struct {
      ID      int
      Artists string
      Title   string
      Year    int
      Country string
      Barcode string
  }
  ```
- [ ] Replace `GetAllRecords` in `internal/db/records.go`. It currently does
  `SELECT * FROM releases` and scans 6 fields, but the new table has 12
  columns, so it fails at runtime.
- [ ] Query joins `releases`, `release_artists`, and `artists`:
  ```sql
  SELECT
    r.id,
    COALESCE(GROUP_CONCAT(a.name, ', '), '') AS artist_names,
    r.title,
    COALESCE(r.year, 0),
    COALESCE(r.country, ''),
    COALESCE(r.barcode, '')
  FROM releases r
  LEFT JOIN release_artists ra ON ra.release_id = r.id
  LEFT JOIN artists a ON a.id = ra.artist_id
  GROUP BY r.id
  ORDER BY r.artists_sort COLLATE NOCASE, r.title;
  ```
  If artist ordering matters, query ordered rows and aggregate in Go instead
  of relying on `GROUP_CONCAT`.
- [ ] Return `[]types.ReleaseSummary`.
- [ ] Remove or retire the old flat `types.Record` (it has `ArtistID`, which
  no longer exists in the schema).

---

## 3. Update the table

- [ ] In `internal/ui/table.go`, change the columns to something like:

  | Column      | Width |
  |-------------|-------|
  | Release ID  | 10    |
  | Artists     | 25    |
  | Album Title | 25    |
  | Year        | 6     |
  | Country     | 10    |
  | Barcode     | 15    |

- [ ] Build rows from `types.ReleaseSummary`; remove all `ArtistID` usage.
- [ ] Handle `Year == 0` as an empty string instead of `"0"` if desired.
- [ ] Optional: return an error from `BuildTable` and surface it in the UI
  instead of only logging it.

---

## 4. Fix the connection path

- [ ] In `internal/db/connection.go`, `Conn(path string)` currently ignores
  `path` and hardcodes `Side-B.db`.
- [ ] Build the DSN from the argument, e.g.:

  ```go
  dsn := path + "?_foreign_keys=on&_journal_mode=WAL"
  ```

  (`main.go` already passes `./Side-B.db`.)

Note: `InitSchema` uses `CREATE TABLE IF NOT EXISTS`, so it will not migrate
an old-format database. The app will create a fresh `Side-B.db` with the new
schema, which is expected for this pass.

---

## 5. Complete the API mapping

- [ ] In `internal/api/requests.go`, read the body once:

  ```go
  body, err := io.ReadAll(resp.Body)
  // json.Unmarshal(body, &itemData)
  ```

- [ ] Populate the currently commented-out fields:
  - `RawJSON: string(body)`
  - `FetchedAt: time.Now().UTC().Format(time.RFC3339)`
- [ ] Set `Release.Notes = itemData.Notes`.
- [ ] Leave the barcode selection logic as-is for now (first "Scanned"
  barcode wins, otherwise the first barcode).

---

## 6. Small correctness fixes

- [ ] In `internal/ui/update.go`, global keys are checked before forwarding
  to the active popup. While a popup is open, `q`/`a`/`d` can still trigger
  global actions. Forward `tea.KeyMsg` to the popup first when
  `m.popupMode != PopupNone`; only handle global bindings when no popup is
  open.
- [ ] When opening a popup, return `popup.Init()` (cursor blink) instead of
  `nil` if the cursor should blink.
- [ ] In `DeleteRecord`, use a transaction and clean up orphaned artists:

  ```sql
  DELETE FROM releases WHERE id = ?;
  DELETE FROM artists
  WHERE id NOT IN (SELECT DISTINCT artist_id FROM release_artists);
  ```

  Foreign-key cascades are already enabled by the DSN for `release_artists`,
  `tracks`, and `collection_items`.

---

## 7. Verification

- [ ] `go build ./...`
- [ ] `go vet ./...`
- [ ] Run the app with the existing `.env`:
  ```sh
  go build -o Side-B ./cmd/Side-B/main.go
  ./Side-B
  ```
- [ ] Add a release by Discogs ID. Confirm:
  - It appears in the table with artist name(s).
  - `releases`, `artists`, `release_artists`, and `tracks` are populated
    (`sqlite3 Side-B.db ...`).
  - Re-adding the same ID updates instead of duplicating links/tracks.
- [ ] Delete a release. Confirm the row disappears and child rows cascade.
- [ ] Confirm `oldSide-B.db` was not modified.

---

## Deferred (do not do in this pass)

- [ ] `collection_items` integration (ratings, notes, sync status).
- [ ] Migration of `oldSide-B.db` (it has 8 records in the old flat
  `releases` table; would need Discogs re-fetch for artists/tracks).
- [ ] Cover image rendering from `exp/imgTest.go` / `go-termimg`.
- [ ] Tests, lint config, CI.
