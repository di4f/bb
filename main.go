package main

import(
	"fmt"
	"os"
	"path"
	"github.com/jienfak/goblin/cat"
	"github.com/jienfak/goblin/echo"
	"github.com/jienfak/goblin/mkdir"
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