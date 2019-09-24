/* Simple 'echo' implementation. */
package main

import (
	"fmt"
	"strings"
	"os"
)

func main() {
	strsId := 1
	firstOpt := os.Args[strsId]
	nextLineStr := "\n"
	joinStr := " "
	if firstOpt == "-n" {
		nextLineStr = ""
		strsId += 1
	}
	fmt.Printf("%s%s",
	            strings.Join(os.Args[strsId:], joinStr),
	            nextLineStr,)	
}
