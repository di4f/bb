/* Yes program implementation. */
package uniq

import(
	"os"
	"fmt"
	"bufio"
	"io"
	"github.com/di4f/cli/mtool"
)

func Run(flagSet *mtool.Flags) {
	var(
		Uflag bool
	)
	flagSet.BoolVar(&Uflag, "U", false, "Print every line just one time.")
	flagSet.Parse()
	//args := flagSet.Args()

	r := bufio.NewReader(os.Stdin)
	u := make(map[string]int)
	if Uflag {
		for{
			l, e := r.ReadString('\n')	
			if e==io.EOF {
				break
			}
			_, haskey := u[l]
			if !haskey {
				u[l] = 1
				fmt.Print(l)
			}
		}
	}

}
