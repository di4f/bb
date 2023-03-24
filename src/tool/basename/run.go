package basename
/* Print file path without extension. */

import(
	"fmt"
	"path"
	"github.com/surdeus/gomtool/src/mtool"
)

var(
	arg0 string
	args []string
	slash =  '/'
)

func Base(p string) string {
	return path.Base(p)
}

func Run(flags *mtool.Flags) {
	flags.Parse()
	args := flags.Args()

	lasti := len(args) - 1
	for i, v := range args {
		fmt.Printf("%s", Base(v))
		if i != lasti {
			fmt.Println("")
		}
	}

}
