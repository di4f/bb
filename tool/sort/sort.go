package sort

import (
	"fmt"
	"os"
	"sort"
	"github.com/di4f/bb/input"
	"github.com/di4f/cli/mtool"
)

func Run(flagSet *mtool.Flags) {
	lines, _ := input.ReadAllLines(os.Stdin)
	sort.Strings(lines)
	for _, l := range lines {
		fmt.Print(l)
	}

}
