# Servers.com go client

![Test](https://github.com/serverscom/serverscom-go-client/workflows/Test/badge.svg)
[![GoDoc](https://godoc.org/github.com/serverscom/serverscom-go-client/pkg?status.svg)](https://godoc.org/github.com/serverscom/serverscom-go-client/pkg)

serverscom-go-client is a Go client library for accessing the [Servers.com API](https://developers.servers.com/api-documentation/v1/)

## Development status

Unstable, in development

## Example

```go
client := serverscom.NewClient("my-jwt-token")

hosts, err := client.Hosts.Collection().Collect()
if err != nil {
  panic(err.Error())
}

for h, i := range hosts {
  log.Println(fmt.Sprintf("Host: %s, with title: %s, private ipv4: %s, public ipv4: %s", h.ID, h.Title, h.PrivateIPv4Address, h.PublicIPv4Address))
}
```

## Checking spec drift

`cmd/specsync` is a read-only tool that compares this client against the OpenAPI spec and
reports what needs syncing. It never edits the client or the spec — it only prints
findings. The spec is fetched from the canonical URL
(`https://developers.servers.com/api-documentation/v1/index.json`) by default; pass
`-spec <path>` to use a local file offline.

```sh
make specsync        # print the drift report
make specsync-tasks  # emit a Markdown brief for an AI agent to act on
```

The report has these sections:

- **Missing** — operations in the spec that the client does not implement.
- **Renamed / moved** — operations the client already implements but at an old path; the
  spec changed the path. These are paired by the `operation/<id>` doc-comment annotation
  (with a path-similarity fallback), so a rename shows as one line (`old → new`) instead
  of an unrelated missing+stale pair.
- **Stale** — client operations whose path is not found in the spec and that are not a
  rename (e.g. a removed endpoint).
- **Field diffs** (advisory) — for matched operations, response/request fields the spec
  has but the Go struct lacks (and vice versa), each with its mapped Go type. Heuristic;
  never affects the exit code.

`make specsync-tasks` (or `-instructions`) emits a Markdown brief grouped into renames to
fix, missing endpoints to add (with typed request/response fields), struct fields to add
(ready-to-paste `Field Type` + json tag lines), and stale ops to investigate.

Useful flags: `-tag "<Tag>"` to focus on one API group, `-json` for machine-readable
output, `-fail-on-missing` to gate CI, and `-debug` to list unresolved methods and
per-operation coverage detail. 

## Copyright

The library is available as open source under the terms of the [MIT License](http://opensource.org/licenses/MIT).
