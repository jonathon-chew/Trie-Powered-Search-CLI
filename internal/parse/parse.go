package parse

import (
	"Trie-Powered-Search-CLI/pkg/utils"
	"bufio"
	"fmt"
	"os"
)

type Node struct {
	Children map[byte]*Node
	Value    byte
	Terminal bool
}

func (n *Node) trieFind(keys []byte) (Node, error) {

	current := n

	for _, i := range keys {
		if current.Children[i] == nil {
			return Node{}, fmt.Errorf("No Node Found")
		}
		current = current.Children[i]
	}

	return *current, nil
}

func (n *Node) trieReturn(currentText string) {

	if n.Terminal {
		fmt.Println(currentText)
	}

	for _, each_child := range n.Children {
		nextText := currentText + string(each_child.Value)
		each_child.trieReturn(nextText)
	}
}

func (n *Node) trieInsert(line []byte) {
	current := n

	for i := 0; i < len(line); i++ { // Loop through the line

		//Ensure Children maps exist
		if current.Children == nil {
			current.Children = make(map[byte]*Node)
		}

		if _, ok := current.Children[line[i]]; ok == false {
			current.Children[line[i]] = &Node{Value: line[i]}
			current = current.Children[line[i]]
		} else {
			current = current.Children[line[i]]
		}
	}

	current.Terminal = true
}

func File(fileName string, pattern string) {
	file_pointer, err := os.Open(fileName)
	if err != nil {
		utils.Broke(err.Error())
	}

	defer file_pointer.Close()

	scanner := bufio.NewScanner(file_pointer)
	// Use a larger buffer for better I/O performance (default is 64KB)
	buf := make([]byte, 0, 1024*1024) // 1MB buffer
	scanner.Buffer(buf, 10*1024*1024) // Max 10MB line size

	var root = new(Node)
	// var child map[byte]Node
	for scanner.Scan() { // Loop through the file
		line := scanner.Bytes()
		root.trieInsert(line)
	}

	if err := scanner.Err(); err != nil {
		utils.Broke("Error reading file: " + err.Error())
	}

	foundNode, err := root.trieFind([]byte(pattern))
	if err != nil {
		utils.Broke("No matches for the pattern")
	}

	foundNode.trieReturn(pattern)
}
