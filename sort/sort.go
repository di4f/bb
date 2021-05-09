package sort

import (
	"fmt"
	"os"
	"sort"
	"flag"
	"goblin/input"
)

func Run(args []string) int {
	flagSet := flag.NewFlagSet(args[0], flag.ExitOnError)	
	flagSet.Parse(args[1:])
	status := 0

	lines, _ := input.ReadAllLines(os.Stdin)
	sort.Strings(lines)
	for _, l := range lines {
		fmt.Print(l)
	}

	return status
}
