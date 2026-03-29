# Trie-Powered-Search-CLI

## Summary

`Trie-Powered-Search-CLI` is a local search tool for building fast prefix-based lookup over personal notes, command snippets, bookmarks, or glossary files. It turns the classic trie into something practical you can use day to day while learning how indexed string search differs from brute-force scanning.

Today it supports:

- indexing a newline-delimited text file,
- prefix lookup over full lines,
- recursive completion traversal below the matched prefix,
- duplicate counting on exact full entries.

## Why This Project Exists

This project is meant to teach:

- trie construction and prefix traversal,
- the tradeoffs between lookup speed and memory usage,
- ranking and suggestion strategies for autocomplete,
- how to turn a core data structure into a useful CLI workflow.

## Current Capabilities

- Index a plain-text file with one entry per line.
- Search by full-line prefix and print matching completions.
- Track duplicate entries by frequency at terminal nodes.
- Print counts beside each match.

## Planned Capabilities

- Index a directory of plain-text files or line-based records.
- Show autocomplete suggestions ranked by frequency, recency, or score.
- Rebuild or persist the index for repeated local usage.

## Architecture Sketch

- A parser ingests local files into searchable records.
- A trie stores prefixes in a tree of nodes.
- Each node tracks:
  - child nodes keyed by byte,
  - whether a full entry ends at that node,
  - the node byte for easier reasoning while learning,
  - a count for how many times the full entry has been inserted.
- A ranking layer decides which suggestions should surface first.
- The CLI coordinates indexing, querying, and output formatting.

## Milestones

1. Index newline-delimited text and support basic prefix lookup.
2. Add ranked autocomplete suggestions and case-normalisation rules.
3. Add incremental reindexing or persisted snapshots.
4. Add highlighting, filtering, and benchmark comparisons against linear scans.

## Usage

Current CLI flags:

- `--file`, `-f` for the input file
- `--pattern`, `-p` for the prefix to search

Example:

```bash
go run ./cmd/Trie-Powered-Search-CLI --file ./testdata/example.txt --pattern git
```

Example output:

```text
git status ( 1 )
git add . ( 1 )
git commit -m "" ( 1 )
git push ( 1 )
git pull ( 1 )
```

The output format is currently:

- `<full match> ( <count> )`

Counts represent how many times the exact full entry was inserted into the trie.

## Current Status

This project is usable today for newline-delimited files and full-line prefix queries.

Implemented:

- file ingestion for newline-delimited input,
- a trie node model with child pointers,
- insertion logic that walks or creates child nodes byte by byte,
- prefix lookup that walks to the node matching the requested prefix,
- recursive completion traversal below the matched prefix node,
- terminal-node frequency counting for duplicate entries,
- focused automated tests for core trie behavior.

Still to do before calling it polished:

- make output ordering deterministic,
- add more edge-case coverage around empty lines and formatting,
- decide whether to sort results alphabetically, by count, or both,
- add practical ranking or persistence later.

## Current Implementation Notes

The trie design currently being explored uses a node shape with:

- `Children` keyed by `byte` and pointing to child nodes,
- `Value` to make the path easier to visualise while learning,
- `Terminal` to mark where a complete entry ends,
- `Count` to track how often a full entry has been inserted.

The intended flow is:

1. Insert each line by starting at the root.
2. For each byte in the line:
   - ensure the current node can hold children,
   - create the child if it does not already exist,
   - move to that child.
3. Mark the final node as terminal and increment its count.
4. Search for a prefix by walking child-to-child without creating nodes.
5. Collect completions by recursively traversing every child below the matched prefix node.
6. Print each completion together with its current count.

One important distinction in this project is:

- `trieFind` identifies whether a prefix path exists and returns the node where that prefix ends.
- completion gathering is a separate recursive subtree walk.
- duplicate lines do not create separate terminal entries; instead they increment the matched terminal node count.

## Development Notes

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
