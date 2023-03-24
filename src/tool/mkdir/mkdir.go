package mkdir

import (
	"fmt"
	"os"

	"github.com/surdeus/goblin/src/pathx"
	"github.com/surdeus/gomtool/src/mtool"
)

func Run(flagSet *mtool.Flags) {
	var (
		parentFlag bool
		modeArg    int
	)
	flagSet.BoolVar(&parentFlag, "p", false, "No error if existing, make parent as needed.")
	flagSet.IntVar(&modeArg, "m", 0766, "Set file `mode`.")
	flagSet.Parse()
	args := flagSet.Args()
	if len(args) == 0 {
		flagSet.Usage()
		os.Exit(1)
	}
	mode := os.FileMode(modeArg)
	for _, path := range args {
		var e error
		pth := pathx.From(path).Real()
		if parentFlag {
			e = os.MkdirAll(pth, mode)
		} else {
			e = os.Mkdir(pth, mode)
		}
		if e != nil {
			fmt.Fprintf(os.Stderr, "%s\n", e)
		}
	}
}
