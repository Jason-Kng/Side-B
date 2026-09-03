PRAGMA foreign_keys = ON;

CREATE TABLE IF NOT EXISTS releases (
  id INTEGER PRIMARY KEY NOT NULL,
  master_id INTEGER,
  title TEXT NOT NULL,
  artists_sort TEXT NOT NULL DEFAULT '',
  year INTEGER,
  release_date TEXT,
  country TEXT,
  barcode TEXT,
  notes TEXT,
  thumb_url TEXT,
  raw_json TEXT,
  fetched_at TEXT NOT NULL DEFAULT (datetime('now'))
);

CREATE TABLE IF NOT EXISTS artists (
  id INTEGER PRIMARY KEY NOT NULL,
  name TEXT NOT NULL,
  thumb_url TEXT
);

CREATE TABLE IF NOT EXISTS release_artists (
  release_id INTEGER NOT NULL REFERENCES releases(id) ON DELETE CASCADE,
  artist_id INTEGER NOT NULL REFERENCES artists(id),
  anv TEXT,
  join_str TEXT,
  position INTEGER NOT NULL,
  PRIMARY KEY (release_id, artist_id, position)
);

CREATE TABLE IF NOT EXISTS tracks (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  release_id INTEGER NOT NULL REFERENCES releases(id) ON DELETE CASCADE,
  position TEXT,
  title TEXT NOT NULL,
  duration TEXT,
  track_index INTEGER NOT NULL,
  type_ TEXT NOT NULL DEFAULT 'track'
);

CREATE INDEX IF NOT EXISTS idx_tracks_release
  ON tracks(release_id, track_index);

CREATE TABLE IF NOT EXISTS collection_items (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  release_id INTEGER NOT NULL REFERENCES releases(id) ON DELETE CASCADE,
  discogs_instance_id INTEGER UNIQUE,
  date_acquired TEXT,
  personal_rating INTEGER CHECK (personal_rating BETWEEN 1 AND 5),
  personal_notes TEXT,
  sync_status TEXT,
    CHECK (sync_status IN ('synced', 'pending_add', 'pending_delete', 'local_only')),
  created_at TEXT NOT NULL DEFAULT (datetime('now'))
);
