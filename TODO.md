# Trie-Powered-Search-CLI TODO

## Goal

Build a local search CLI that indexes short text entries and returns matching results quickly for prefix queries. The main learning goal is to understand how a trie enables fast repeated prefix lookups compared to brute-force scanning.

## Project Outcome

- Index useful local data such as notes, snippets, commands, bookmarks, or glossary entries.
- Search by prefix and return matching entries quickly.
- Turn a core DSA into a practical CLI you can actually use.

## Recommended Build Order

1. Define the exact problem and input shape.
2. Build a brute-force baseline.
3. Implement trie insertion and prefix lookup.
4. Add completion collection below a matched prefix.
5. Add ranking and result limits.
6. Add one practical enhancement.

## Phase 1: Define The First Version

- Decide what one searchable item is.
- Keep the first version simple: one line equals one entry.
- Choose the first input source.
- Recommended starting point: a plain text file with one entry per line.
- Decide whether matching is case-sensitive or case-insensitive.
- Recommended starting point: case-insensitive matching.
- Decide whether the prefix applies to the full line only or to words within the line.
- Recommended starting point: full-line prefix matching only.

## Phase 2: Build A Brute-Force Baseline

- Load all entries from the chosen input file.
- For each query, scan every entry linearly.
- Return entries whose text starts with the requested prefix.
- Use this version to confirm the CLI behaviour before introducing the trie.
- Compare all later trie results against this baseline for correctness.

## Phase 3: Learn And Model The Trie

- Understand the minimum trie structure:
- a root node,
- children keyed by character,
- a marker showing whether a full entry ends at a node,
- a link back to the stored entry or record ID.
- Decide what the trie stores.
- Recommended starting point: store record IDs and keep the actual entry text in a separate collection.
- Decide how to represent characters.
- Recommended starting point: normalised lowercase ASCII if the data is mostly commands, snippets, and notes.

## Phase 4: Implement Core Trie Operations

- Insert a searchable entry into the trie.
- Walk the trie for a given prefix.
- Detect when a prefix does not exist.
- Handle the case where a prefix is itself a complete entry.
- Handle the case where multiple entries share the same prefix.

## Phase 5: Collect Matching Results

- Once a prefix node is found, traverse below it to gather matches.
- Choose a traversal style.
- Recommended starting point: depth-first traversal is simple and fine for the first version.
- Return matching entries in a predictable order.
- Decide how many matches to return at most.
- Recommended starting point: cap results to a small default such as 10 or 20.

## Phase 6: Add Ranking

- Start with a simple ordering rule.
- Recommended starting point: alphabetical order.
- Consider future ranking signals:
- shorter matches first,
- most frequently selected entries,
- most recently used entries,
- custom score fields.
- Keep ranking separate from trie traversal if possible.

## Phase 7: Make It Useful

- Choose one practical enhancement after the trie search works.
- Good options:
- case-normalisation improvements,
- persisted index snapshots,
- indexing multiple files in a directory,
- highlighting matched prefixes in output,
- frequency-based suggestion ranking.

## Design Decisions To Make Early

- What is being indexed:
- raw lines,
- titles,
- commands,
- snippets.
- What matching means:
- full-line prefix only,
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
- duplicate entries,
- a word being a prefix of another word,
- nonexistent prefixes,
- very large result sets,
- case-normalisation rules,
- character handling beyond simple ASCII.

## Testing Checklist

- inserted entries can be found by valid prefixes,
- nonexistent prefixes return no matches,
- shared prefixes behave correctly,
- an entry that is itself a prefix of another still appears correctly,
- duplicate input is handled intentionally,
- result limits are respected,
- case rules behave exactly as designed.

## Good Example Test Data

- `go`
- `goal`
- `gone`
- `git`

These examples are useful because they expose shared-prefix behaviour quickly.

## Definition Of Done For Version 1

- Read a text file containing one entry per line.
- Build an in-memory trie index from those entries.
- Accept a prefix query.
- Return matching entries in a stable order.
- Keep the implementation simple enough that the trie design is easy to explain and reason about.

## Nice Next Steps After Version 1

- Add indexing for entire directories.
- Persist the built index between runs.
- Add better ranking based on usage.
- Support matching on individual words inside a line.
- Compare performance against the brute-force baseline.
