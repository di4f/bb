package main

import(
	"fmt"
	"os"
	"path"
	"github.com/k1574/goblin/cat"
	"github.com/k1574/goblin/echo"
	"github.com/k1574/goblin/mkdir"
	"github.com/k1574/goblin/gtrue"
	"github.com/k1574/goblin/gfalse"
	"github.com/k1574/goblin/sort"
	"github.com/k1574/goblin/tac"
	"github.com/k1574/goblin/ls"
	"github.com/k1574/goblin/yes"
	"github.com/k1574/goblin/date"
	"github.com/k1574/goblin/uniq"
)

func main() {
	var(
		utilName string
		args []string
	)
	

	utilsMap := map[string]  interface{}  {
		"cat": cat.Run,
		"mkdir" : mkdir.Run,
		"echo" : echo.Run,
		"true" : gtrue.Run,
		"false" : gfalse.Run,
		"sort" : sort.Run,
		"tac" : tac.Run,
		"ls" : ls.Run,
		"yes" : yes.Run,
		"date" : date.Run,
		"uniq" : uniq.Run,
	}

	if binBase := path.Base(os.Args[0]) ; binBase != "goblin" {
		utilName = binBase
		args = os.Args[:]
	} else {
		if len(os.Args)<2  {
			for k, _ := range utilsMap {
				fmt.Printf("%s\n", k)
			}
			os.Exit(0)
		}
		utilName = os.Args[1]
		args = os.Args[1:]
	}

	if _, ok := utilsMap[utilName] ; !ok {
		fmt.Printf("%s: No such uitl as '%s'.\n", os.Args[0], utilName )
		os.Exit(1)
	}
	status := utilsMap[utilName].(func([]string) int )(args)
	os.Exit(status)
}
