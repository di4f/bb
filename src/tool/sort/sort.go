package sort

import (
	"fmt"
	"os"
	"sort"
	"github.com/mojosa-software/goblin/src/input"
	"github.com/mojosa-software/gomtool/src/mtool"
)

func Run(flagSet *mtool.Flags) {
	lines, _ := input.ReadAllLines(os.Stdin)
	sort.Strings(lines)
	for _, l := range lines {
		fmt.Print(l)
	}

}
