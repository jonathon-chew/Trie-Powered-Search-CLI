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
- Track duplicate entries by frequency at terminal nodes.
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

## Current Status

This project is now partially implemented.

Current progress includes:

- file ingestion for newline-delimited input,
- a trie node model with child pointers,
- insertion logic that walks or creates child nodes byte by byte,
- prefix lookup that walks to the node matching the requested prefix,
- recursive completion traversal below the matched prefix node,
- terminal-node frequency counting for duplicate entries,
- output that includes the current frequency count for each match.

Still to do:

- add explicit tests around exact-prefix and terminal-with-children behaviour,
- handle edge cases such as empty lines,
- add stable output ordering and tests,
- decide whether to sort matches by count,
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

Planned commands once implementation begins:

- `go run ./cmd/Trie-Powered-Search-CLI`
- `go build ./cmd/Trie-Powered-Search-CLI`
- `go test ./...`

Current learning checkpoint:

- insertion is in the right structural shape,
- prefix search is in the right structural shape,
- recursive completion traversal is working,
- duplicate-count tracking is working,
- the next main step is tightening correctness and test coverage around exact-prefix behaviour, count-aware ordering, and output ordering.

## Project Structure

```text
cmd/Trie-Powered-Search-CLI/    future search CLI entrypoint
internal/                       indexing, ranking, and CLI internals
pkg/                            optional reusable trie and search packages
doc/                            design notes and benchmark ideas
scripts/                        helper scripts
```
