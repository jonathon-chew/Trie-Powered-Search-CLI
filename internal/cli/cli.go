package cli

import (
	"Trie-Powered-Search-CLI/pkg/utils"
	"fmt"
	"os"
)

type Flags struct {
	FileName   string
	Pattern    string
	ErrMessage string
	Err        bool
}

func CLI(args []string) Flags {

	var flags Flags

	for i := 0; i < len(args); i++ {
		arg := args[i]
		switch arg {
		default:
			fmt.Println("Unrecongised command", arg)
		case "--file", "-f":
			if len(args) > i+1 {
				if _, err := os.Stat(args[i+1]); err != nil {
					flags.Err = true
					flags.ErrMessage = err.Error()
					return flags
				}
				flags.FileName = args[i+1]
				i++
			} else {
				utils.Broke("No argument found")
			}
		case "--pattern", "-p":
			if len(args) > i+1 {
				flags.Pattern = args[i+1]
				i++
			} else {
				utils.Broke("No argument found")
			}
		}
	}

	return flags
}
