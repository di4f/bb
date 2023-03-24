package noext
/* Print file path without extension. */

import(
	"fmt"
	"github.com/surdeus/gomtool/src/mtool"
)

var(
	arg0 string
	args []string
	dot = '.'
	slash =  '/'
)

func NoExt(p string) string {
	l := len(p)-1
	i := l
	for ; rune(p[i])!=dot ; i -= 1 {
		if rune(p[i]) == slash || i==0 {
			return p
		}
	}
	return p[:i]
}

func Run(flagSet *mtool.Flags) {
	flagSet.Parse()
	args := flagSet.Args()
	if len(args) < 1 { flagSet.Usage() }
	fmt.Printf("%s", NoExt(args[0]))
}
