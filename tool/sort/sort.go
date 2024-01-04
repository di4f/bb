package sort

import (
	"fmt"
	"os"
	"sort"
	"github.com/vultras/bb/input"
	"github.com/vultras/cli/mtool"
)

func Run(flagSet *mtool.Flags) {
	lines, _ := input.ReadAllLines(os.Stdin)
	sort.Strings(lines)
	for _, l := range lines {
		fmt.Print(l)
	}

}
