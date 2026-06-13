# gtv

**Gollem Trace Viewer** — a local web UI for browsing [gollem](https://github.com/gollem-dev/gollem) trace files.

`gtv` starts an HTTP server that serves a single-page web app (embedded in the binary) and exposes a small JSON API over a trace source. Traces can be read from a local directory or from a Google Cloud Storage bucket.

## Install

```sh
go install github.com/gollem-dev/gtv@latest
```

## Usage

Browse traces stored in a local directory:

```sh
gtv --dir ./traces
```

Browse traces stored in Google Cloud Storage (uses Application Default Credentials):

```sh
gtv --gs gs://my-bucket/path/to/traces/
```

By default the server listens on `:18900` and opens your browser automatically.

### Flags

| Flag | Env | Default | Description |
| --- | --- | --- | --- |
| `--addr` | `GTV_ADDR` | `:18900` | Server listen address |
| `--dir` | `GTV_DIR` | | Local directory containing trace JSON files |
| `--gs` | `GTV_GS` | | Google Cloud Storage URI (e.g. `gs://bucket/prefix/`) |
| `--no-browser` | `GTV_NO_BROWSER` | `false` | Do not open the browser automatically |

`--dir` and `--gs` are mutually exclusive; exactly one must be specified.

## API

| Method | Path | Description |
| --- | --- | --- |
| `GET` | `/api/health` | Health check |
| `GET` | `/api/traces?path=&page_size=&page_token=` | List entries (files/directories) under a path |
| `GET` | `/api/traces/{path...}` | Get a single trace by path (without `.json` suffix) |

## Development

```sh
# Backend
go test ./...
go vet ./...

# Frontend (rebuilds the embedded assets under frontend/dist)
cd frontend
corepack enable
pnpm install --frozen-lockfile
pnpm build
```

The built frontend assets in `frontend/dist` are embedded into the Go binary via `frontend/static.go`.
