package parse

import (
	"io"
	"os"
	"slices"
	"strings"
	"testing"
)

func buildTrie(lines ...string) *Node {
	root := new(Node)
	for _, line := range lines {
		root.trieInsert([]byte(line))
	}
	return root
}

func captureStdout(t *testing.T, fn func()) string {
	t.Helper()

	oldStdout := os.Stdout
	reader, writer, err := os.Pipe()
	if err != nil {
		t.Fatalf("create stdout pipe: %v", err)
	}

	os.Stdout = writer
	defer func() {
		os.Stdout = oldStdout
	}()

	fn()

	if err := writer.Close(); err != nil {
		t.Fatalf("close stdout writer: %v", err)
	}

	output, err := io.ReadAll(reader)
	if err != nil {
		t.Fatalf("read stdout: %v", err)
	}

	return string(output)
}

func normalizeOutputLines(output string) []string {
	lines := strings.Split(strings.TrimSpace(output), "\n")
	normalized := make([]string, 0, len(lines))

	for _, line := range lines {
		line = strings.Join(strings.Fields(line), " ")
		if line != "" {
			normalized = append(normalized, line)
		}
	}

	slices.Sort(normalized)

	return normalized
}

func TestTrieInsertTracksDuplicateCounts(t *testing.T) {
	root := buildTrie("git log", "git log", "git log --oneline")

	foundNode, err := root.trieFind([]byte("git log"))
	if err != nil {
		t.Fatalf("find prefix: %v", err)
	}

	if !foundNode.Terminal {
		t.Fatalf("expected prefix node to be terminal")
	}

	if foundNode.Count != 2 {
		t.Fatalf("expected duplicate count 2, got %d", foundNode.Count)
	}
}

func TestTrieFindReturnsErrorForMissingPrefix(t *testing.T) {
	root := buildTrie("git status", "git stash")

	if _, err := root.trieFind([]byte("git log")); err == nil {
		t.Fatalf("expected missing prefix to return error")
	}
}

func TestTrieReturnIncludesExactPrefixAndChildren(t *testing.T) {
	root := buildTrie("git log", "git log", "git log --oneline", "git status")

	foundNode, err := root.trieFind([]byte("git log"))
	if err != nil {
		t.Fatalf("find prefix: %v", err)
	}

	output := captureStdout(t, func() {
		foundNode.trieReturn("git log")
	})

	got := normalizeOutputLines(output)
	want := []string{
		"git log ( 2 )",
		"git log --oneline ( 1 )",
	}

	if !slices.Equal(got, want) {
		t.Fatalf("unexpected trie output\nwant: %v\ngot:  %v", want, got)
	}
}
