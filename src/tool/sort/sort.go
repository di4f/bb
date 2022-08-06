package sort

import (
	"fmt"
	"os"
	"sort"
	"flag"
	"github.com/surdeus/goblin/src/input"
)

func Run(args []string) {
	flagSet := flag.NewFlagSet(args[0], flag.ExitOnError)	
	flagSet.Parse(args[1:])

	lines, _ := input.ReadAllLines(os.Stdin)
	sort.Strings(lines)
	for _, l := range lines {
		fmt.Print(l)
	}

}
