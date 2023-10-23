package quote
/* Quote quotes string if it contains white space character. */

import(
	"os"
	"io"
	"fmt"
	"unicode"
	"bufio"
	"github.com/reklesio/mtool"
)

func HasWhiteSpace(s string) bool {
	for _, r := range s {
		if(unicode.IsSpace(r)){
			return true
		}
	}
	return false
}

func Run(flagSet *mtool.Flags) {	
	flagSet.Parse()

	r := bufio.NewReader(os.Stdin)
	for{
		l, e := r.ReadString('\n')	
		if e==io.EOF {
			break
		}
		last := len(l) - 1
		if l[last] == '\n' {
			l = l[:last]	
		}
		if HasWhiteSpace(l) {
			fmt.Printf("'%s'\n", l)
		}else {
			fmt.Println(l)
		}
	}
}
