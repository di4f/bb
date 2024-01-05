package sort

import (
	"fmt"
	"os"
	"sort"
	"vultras.su/core/bb/input"
	"vultras.su/core/cli/mtool"
)

func Run(flagSet *mtool.Flags) {
	lines, _ := input.ReadAllLines(os.Stdin)
	sort.Strings(lines)
	for _, l := range lines {
		fmt.Print(l)
	}

}
