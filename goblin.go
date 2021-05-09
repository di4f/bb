package main

import(
	"fmt"
	"os"
	"path"
	"goblin/cat"
	"goblin/echo"
	"goblin/mkdir"
	"goblin/gtrue"
	"goblin/gfalse"
	"goblin/sort"
	"goblin/tac"
	"goblin/ls"
	"goblin/yes"
	"goblin/date"
	"goblin/uniq"
	"goblin/quote"
	"goblin/urlprs"
	"goblin/noext"
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
		"quote" : quote.Run,
		"urlprs" : urlprs.Run,
		"noext" : noext.Run,
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
