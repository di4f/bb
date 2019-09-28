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
	backslashSeqFlagPtr := flag.Bool("e", false, "Interpret backslash special terminal characters.")
	flag.Parse()

	if *newLineFlagPtr {
		*newLineStrPtr = ""	
	}
	if *joinStrsFlagPtr {
		*joinStrPtr = ""
	}

	printStr := strings.Join(flag.Args(), *joinStrPtr)
	if *backslashSeqFlagPtr {
		seqs := map[string] string {
			"\\\\" : "\\",
			"\\a" : "\a",
			"\\b" : "\b",
			/*"\\c" : "\c",
			"\\e" : "\e",*/
			"\\f" : "\f",
			"\\n" : "\n",
			"\\r" : "\r",
			"\\t" : "\t",
			"\\v" : "\v",
		}
		for k, v := range seqs {
			printStr = strings.ReplaceAll(printStr, k, v)
		}
	}

	fmt.Printf("%s%s",
	            printStr,
	            *newLineStrPtr,)
}
