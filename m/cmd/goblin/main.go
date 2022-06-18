package main

import(
	"github.com/k1574/gomtool/m/multitool"
	"github.com/k1574/goblin/m/tool/cat"
	"github.com/k1574/goblin/m/tool/echo"
	"github.com/k1574/goblin/m/tool/mkdir"
	"github.com/k1574/goblin/m/tool/gtrue"
	"github.com/k1574/goblin/m/tool/gfalse"
	"github.com/k1574/goblin/m/tool/sort"
	"github.com/k1574/goblin/m/tool/tac"
	"github.com/k1574/goblin/m/tool/ls"
	"github.com/k1574/goblin/m/tool/yes"
	"github.com/k1574/goblin/m/tool/date"
	"github.com/k1574/goblin/m/tool/uniq"
	"github.com/k1574/goblin/m/tool/quote"
	"github.com/k1574/goblin/m/tool/urlprs"
	"github.com/k1574/goblin/m/tool/noext"
	"github.com/k1574/goblin/m/tool/mergelbl"
	"github.com/k1574/goblin/m/tool/basename"
	"github.com/k1574/goblin/m/tool/ec"
	"github.com/k1574/goblin/m/tool/read"
	"github.com/k1574/goblin/m/tool/wc"
	"github.com/k1574/goblin/m/tool/ftest"
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
		"wc" : wc.Run,
		"ftest" : ftest.Run,
	}

	multitool.Main("goblin", tools)
}
