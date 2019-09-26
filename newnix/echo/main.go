/* Simple 'echo' implementation. */
package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	newLineFlagPtr := flag.Bool("n", false,
	                            "Don't add new line character('-N' is lower priority).")
	newLineStrPtr := flag.String("N", "\n", "Use this instead new line character.") 
	joinStrsFlagPtr := flag.Bool("j", false, "Join strings('-J' is lower priority).")
	joinStrPtr := flag.String("J", " ", "Use instead of space as separator.") 
	flag.Parse()

	if *newLineFlagPtr {
		*newLineStrPtr = ""	
	}
	if *joinStrsFlagPtr {
		*joinStrPtr = ""
	}

	fmt.Printf("%s%s",
	            strings.Join(flag.Args(), *joinStrPtr),
	            *newLineStrPtr,)	
}
