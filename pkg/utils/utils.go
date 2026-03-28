package utils

import (
	"fmt"
	"os"
)

func Broke(message string) {
	fmt.Println(message)
	os.Exit(1)
}
