package main

import(
	"github.com/k1574/gomtool/m/multitool"
	"github.com/k1574/goblin/m/sep/cat"
	"github.com/k1574/goblin/m/sep/echo"
	"github.com/k1574/goblin/m/sep/mkdir"
	"github.com/k1574/goblin/m/sep/gtrue"
	"github.com/k1574/goblin/m/sep/gfalse"
	"github.com/k1574/goblin/m/sep/sort"
	"github.com/k1574/goblin/m/sep/tac"
	"github.com/k1574/goblin/m/sep/ls"
	"github.com/k1574/goblin/m/sep/yes"
	"github.com/k1574/goblin/m/sep/date"
	"github.com/k1574/goblin/m/sep/uniq"
	"github.com/k1574/goblin/m/sep/quote"
	"github.com/k1574/goblin/m/sep/urlprs"
	"github.com/k1574/goblin/m/sep/noext"
	"github.com/k1574/goblin/m/sep/mergelbl"
	"github.com/k1574/goblin/m/sep/basename"
	"github.com/k1574/goblin/m/sep/ec"
	"github.com/k1574/goblin/m/sep/read"
)

func main() {
	tools :=  multitool.Tools {
		"basename" : basename.Run,
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
		"mergelbl" : mergelbl.Run,
		"ec" : ec.Run,
		"read" : read.Run,
	}

	multitool.Main("goblin", tools)
}
