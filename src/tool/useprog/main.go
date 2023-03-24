package useprog

import (
	"os/exec"
	"fmt"
	"github.com/surdeus/gomtool/src/mtool"
)

func Run(flagSet *mtool.Flags) {
	flagSet.Parse()
	args := flagSet.Args()

	if len(args) < 1 {
		flagSet.Usage()
	}

	for _, v := range args {
		_, err := exec.LookPath(v)
		if err != nil {
			continue
		}
		fmt.Printf("%s", v)
		break
	}
}
