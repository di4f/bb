/* Simple 'echo' implementation. */
package echo

import (
	"os"
	"flag"
	"strings"
)

type Echoer struct {
	out *os.File
	backslashMap map[string] string
}
func (e Echoer)BackslashInterpret(str string) string {
	for k, v := range e.backslashMap {
		str  = strings.ReplaceAll(str, k, v)
	}
	return str
}

func (e Echoer)Echo(str string) {
	e.out.Write( []byte(str[:len(str)]) )
}

func New(out *os.File, backslashMap map[string] string) *Echoer {
	e := new(Echoer)
	e.out = out
	e.backslashMap = backslashMap
	return e
}
func Run(args []string) int {
	var(
		newLineStr string
		newLineFlag bool
		joinStr string
		joinStrsFlag bool
		backslashSeqFlag bool
	)
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
	e := New(os.Stdout, seqs)
	flagSet := flag.NewFlagSet(args[0], flag.ExitOnError)
	flagSet.BoolVar(&newLineFlag, "n", false,
	                            "Don't add new line character('-N' is lower priority).")
	flagSet.StringVar(&newLineStr, "N", "\n", "Use this instead new line character.") 
	flagSet.BoolVar(&joinStrsFlag, "j", false, "Join strings('-J' is lower priority).")
	flagSet.StringVar(&joinStr, "J", " ", "Use instead of space as separator.") 
	flagSet.BoolVar(&backslashSeqFlag, "e", false, "Interpret backslash special terminal characters(Characters from join options will be interpreted too).")
	flagSet.Parse(args[1:])

	if newLineFlag {
		newLineStr = ""	
	}
	if joinStrsFlag {
		joinStr = ""
	}

	printStr := strings.Join(flagSet.Args(), joinStr) +newLineStr
	if backslashSeqFlag {
		printStr = e.BackslashInterpret(printStr)
	}

	e.Echo(printStr)
	return 0
}
