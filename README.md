# Trie-Powered-Search-CLI

## Summary

`Trie-Powered-Search-CLI` is a local search tool for building fast prefix-based lookup over personal notes, command snippets, bookmarks, or glossary files. It turns the classic trie into something practical you can use day to day while learning how indexed string search differs from brute-force scanning.

## Why This Project Exists

This project is meant to teach:

- trie construction and prefix traversal,
- the tradeoffs between lookup speed and memory usage,
- ranking and suggestion strategies for autocomplete,
- how to turn a core data structure into a useful CLI workflow.

## Planned Capabilities

- Index a directory of plain-text files or line-based records.
- Search by prefix and return matching entries quickly.
- Show autocomplete suggestions ranked by frequency, recency, or score.
- Rebuild or persist the index for repeated local usage.

## Architecture Sketch

- A parser ingests local files into searchable records.
- A trie stores prefixes and links them to matching record IDs.
- A ranking layer decides which suggestions should surface first.
- The CLI coordinates indexing, querying, and output formatting.

## Milestones

1. Index newline-delimited text and support basic prefix lookup.
2. Add ranked autocomplete suggestions and case-normalisation rules.
3. Add incremental reindexing or persisted snapshots.
4. Add highlighting, filtering, and benchmark comparisons against linear scans.

## Current Status

This project is currently scaffolded but not implemented. The module, placeholder CLI, and folder layout exist, but the indexing, ranking, and persistence logic still need to be built.

## Development Notes

Planned commands once implementation begins:

- `go run ./cmd/Trie-Powered-Search-CLI`
- `go build ./cmd/Trie-Powered-Search-CLI`
- `go test ./...`

## Project Structure

```text
cmd/Trie-Powered-Search-CLI/    future search CLI entrypoint
internal/                       indexing, ranking, and CLI internals
pkg/                            optional reusable trie and search packages
doc/                            design notes and benchmark ideas
scripts/                        helper scripts
```
