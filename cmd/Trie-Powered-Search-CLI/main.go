package main

import (
	"os"

	"Trie-Powered-Search-CLI/internal/cli"
	"Trie-Powered-Search-CLI/internal/parse"
	"Trie-Powered-Search-CLI/pkg/utils"
)

func main() {
	Flags := cli.CLI(os.Args[1:])

	if Flags.Err {
		utils.Broke(Flags.ErrMessage)
	} else if Flags.Pattern == "" {
		utils.Broke("No pattern detected")
	} else if Flags.FileName == "" {
		utils.Broke("No file path detected")
	}

	parse.File(Flags.FileName, Flags.Pattern, Flags.Limit)

	os.Exit(0)
}
