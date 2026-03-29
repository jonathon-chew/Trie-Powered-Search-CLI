# Trie-Powered-Search-CLI TODO

## Goal

Build a local search CLI that indexes short text entries and returns matching results quickly for prefix queries. The primary technical goal is to demonstrate how a trie supports fast repeated prefix lookups and autocomplete-style traversal in a practical tool.

## Version 1 Definition

Version 1 of this project is a prefix-first autocomplete CLI for line-based command and snippet data.

Version 1 should:

- read a file with one entry per line,
- build a trie over those entries,
- accept a prefix query,
- return matching completions,
- track duplicate frequency counts,
- return results in deterministic order,
- support a small result limit,
- stay simple enough that the trie remains the main data structure and story.

Version 1 should not try to be:

- full-text search,
- fuzzy search,
- shell integration,
- persistent indexing across runs.

## Project Outcome

- Index useful local data such as notes, snippets, commands, bookmarks, or glossary entries.
- Search by prefix and return matching entries quickly.
- Turn a core DSA into a practical CLI you can actually use.

## Recommended Build Order

1. [x] Define the exact problem and input shape.
2. [ ] Build a brute-force baseline.
3. [x] Implement trie insertion and prefix lookup.
4. [x] Add completion collection below a matched prefix.
5. [ ] Add ranking and result limits.
6. [ ] Add one practical enhancement.

## Current Progress

- [x] Input ingestion from a file is working.
- [x] Trie insertion has reached the correct pointer-based traversal shape.
- [x] Prefix lookup has reached the correct traversal shape.
- [x] Recursive completion collection is working in the current implementation.
- [x] Duplicate entries increment a terminal-node count.
- [x] Structured output entries are collected after recursive traversal.
- [x] Results are sorted after traversal.
- [x] CLI-level result limits are supported.
- [x] A small focused test suite covers core trie behavior.

## Immediate Next Step

- [x] Finish recursive completion traversal from the node returned by `trieFind`.
- [x] Carry the matched prefix into the traversal so returned results are full entries rather than suffix fragments.
- [x] Treat `Terminal` as "emit a result here", not "stop traversing here".
- [x] Recurse once per child branch.
- [x] Decide whether `trieFind` should return a node pointer instead of a copied node value.
- [x] Add explicit tests for terminal-with-children cases and exact prefix output.
- [x] Decide whether to order output by count, alphabetically, or both.
- [x] Add explicit tests for duplicate-count behaviour.
- [ ] Replace package-level output state with function-scoped result collection.
- [ ] Restore or redesign count-aware printed output on top of the sorted results.
- [ ] Add tests for sorted and limited output behavior.

## Employer-Facing Next 5

1. [ ] Replace package-level output state with function-scoped results.
2. [ ] Restore count-aware rendering from the structured output path.
3. [ ] Add tests for sorted output and limit handling.
4. [ ] Separate trie traversal completely from CLI formatting.
5. [ ] Expose a reusable autocomplete-oriented result shape.

## Phase 1: Define The First Version

- [x] Decide what one searchable item is.
- [x] Keep the first version simple: one line equals one entry.
- [x] Choose the first input source.
- Recommended starting point: a plain text file with one entry per line.
- [ ] Decide whether matching is case-sensitive or case-insensitive.
- Recommended starting point: case-insensitive matching.
- [x] Decide whether the prefix applies to the full line only or to words within the line.
- Recommended starting point: full-line prefix matching only.

## Phase 2: Build A Brute-Force Baseline

- [x] Load all entries from the chosen input file.
- [ ] For each query, scan every entry linearly.
- [ ] Return entries whose text starts with the requested prefix.
- [ ] Use this version to confirm the CLI behaviour before introducing the trie.
- [ ] Compare all later trie results against this baseline for correctness.

## Phase 3: Learn And Model The Trie

Understand the minimum trie structure:
- [x] a root node,
- [x] children keyed by character,
- [x] a marker showing whether a full entry ends at a node,
- [ ] a link back to the stored entry or record ID.

Current project direction:
- [x] child nodes are stored as pointers,
- [x] `Value` is kept on the node for easier reasoning,
- [x] `Terminal` marks the end of a complete entry.
- [x] `Count` stores how often an exact full entry has been inserted.
- [ ] Decide what the trie stores.
- Recommended starting point: store record IDs and keep the actual entry text in a separate collection.
- [x] Decide how to represent characters.
- Recommended starting point: normalised lowercase ASCII if the data is mostly commands, snippets, and notes.

## Phase 4: Implement Core Trie Operations

- [x] Insert a searchable entry into the trie.
- [x] Walk the trie for a given prefix.
- [x] Detect when a prefix does not exist.
- [x] Handle the case where a prefix is itself a complete entry.
- [x] Handle the case where multiple entries share the same prefix.

Implementation reminder:

- insertion should mutate real nodes, not copied values,
- search should walk from the current node to the next child node,
- `Terminal` answers "does a full entry end here?" rather than controlling traversal.

## Phase 5: Collect Matching Results

- [x] Once a prefix node is found, traverse below it to gather matches.
- [x] Choose a traversal style.
- Recommended starting point: depth-first traversal is simple and fine for the first version.
- [x] Return matching entries in a predictable order.
- [x] Decide how many matches to return at most.
- Recommended starting point: cap results to a small default such as 10 or 20.
- [ ] Include current frequency count in the emitted output.

Implementation reminder:

- subtree collection is not a single linear walk,
- it needs recursion,
- the recursive shape is:
  - if current node is terminal, emit the current built text,
  - for each child, extend the text and recurse into that child,
- do not return from inside the child loop if you want all matches.

## Phase 6: Add Ranking

- [x] Start with a simple ordering rule.
- Recommended starting point: alphabetical order.
- [x] Evaluate count-based ordering for more frequently used commands.

- Consider future ranking signals:
- shorter matches first,
- most frequently selected entries,
- most recently used entries,
- custom score fields.
- Keep ranking separate from trie traversal if possible.

## Phase 7: Make It Useful

- [ ] Choose one practical enhancement after the trie search works.
- Good options:
- case-normalisation improvements,
- persisted index snapshots,
- indexing multiple files in a directory,
- highlighting matched prefixes in output,
- frequency-based suggestion ranking.

## Reusable Autocomplete Direction

The longer-term useful version of this project is not just a standalone search tool. It is also a reusable autocomplete engine for other CLI tools.

Target use cases:

- command snippet autocomplete,
- subcommand autocomplete,
- flag autocomplete,
- argument autocomplete for controlled value sets.

The key architectural shift for that future is:

- trie logic should gather and return structured results,
- CLI code should decide how to format and print them.

Recommended result shape for that direction:

- matched text,
- frequency count,
- whether the match is an exact terminal match,
- optional score or rank value,
- optional source metadata later if the project grows.

The practical implication is:

- avoid baking too much printing logic into trie traversal,
- move toward "collect results, then sort/limit/print".

## Design Decisions To Make Early

- What is being indexed:
- [x] raw lines,
- titles,
- commands,
- snippets.
- What matching means:
- [x] full-line prefix only,
- word-level prefix matching,
- case-sensitive or case-insensitive.
- What is stored in the trie:
- full strings,
- record IDs pointing to stored entries.
- What success looks like for version 1:
- load input reliably,
- answer prefix queries correctly,
- return stable output,
- feel clearly more structured than brute-force scanning.

## DSA Concepts To Focus On

- tree-based data structures,
- prefix traversal,
- recursive or iterative subtree collection,
- indexing versus scanning,
- time and space tradeoffs,
- separating storage from lookup structures.

## Edge Cases To Think About

- empty queries,
- empty lines in the input file,
- [x] duplicate entries,
- a word being a prefix of another word,
- nonexistent prefixes,
- very large result sets,
- case-normalisation rules,
- character handling beyond simple ASCII.

## Testing Checklist

- [x] inserted entries can be found by valid prefixes,
- [x] nonexistent prefixes return no matches,
- [x] shared prefixes behave correctly,
- [x] an entry that is itself a prefix of another still appears correctly,
- [x] a terminal node can still have children and both results are surfaced,
- [x] duplicate input is handled intentionally,
- [ ] duplicate counts are printed consistently in the current output format,
- [ ] result limits are respected,
- case rules behave exactly as designed.

## Good Example Test Data

- `go`
- `goal`
- `gone`
- `git`
- `github`
- `gist`

These examples are useful because they expose shared-prefix behaviour and terminal-with-children cases quickly.

## Definition Of Done For Version 1

- Read a text file containing one entry per line.
- Build an in-memory trie index from those entries.
- Accept a prefix query.
- Return matching entries in a stable order.
- Include or intentionally suppress duplicate-count metadata.
- Keep the implementation simple enough that the trie design is easy to explain and reason about.

## Nice Next Steps After Version 1

- Add indexing for entire directories.
- Persist the built index between runs.
- Add better ranking based on usage.
- Support matching on individual words inside a line.
- Compare performance against the brute-force baseline.
