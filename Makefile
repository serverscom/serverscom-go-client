test:
	go test ./...

deps:
	go mod tidy

# Report drift between the client and the OpenAPI spec (read-only, no sync).
# The spec is fetched from the canonical URL by default; override with -spec <path|url>.
specsync:
	go run ./cmd/specsync

# Emit a Markdown brief (renames, missing endpoints, struct fields) for an AI agent.
specsync-tasks:
	go run ./cmd/specsync -instructions
