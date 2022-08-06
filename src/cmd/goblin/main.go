package main

import(
	"github.com/surdeus/gomtool/src/multitool"
	"github.com/surdeus/goblin/src/tool/cat"
	"github.com/surdeus/goblin/src/tool/echo"
	"github.com/surdeus/goblin/src/tool/mkdir"
	"github.com/surdeus/goblin/src/tool/gtrue"
	"github.com/surdeus/goblin/src/tool/gfalse"
	"github.com/surdeus/goblin/src/tool/sort"
	"github.com/surdeus/goblin/src/tool/tac"
	"github.com/surdeus/goblin/src/tool/ls"
	"github.com/surdeus/goblin/src/tool/yes"
	"github.com/surdeus/goblin/src/tool/date"
	"github.com/surdeus/goblin/src/tool/uniq"
	"github.com/surdeus/goblin/src/tool/quote"
	"github.com/surdeus/goblin/src/tool/urlprs"
	"github.com/surdeus/goblin/src/tool/noext"
	"github.com/surdeus/goblin/src/tool/mergelbl"
	"github.com/surdeus/goblin/src/tool/basename"
	"github.com/surdeus/goblin/src/tool/ec"
	"github.com/surdeus/goblin/src/tool/read"
	"github.com/surdeus/goblin/src/tool/wc"
	"github.com/surdeus/goblin/src/tool/ftest"
	"github.com/surdeus/goblin/src/tool/grange"
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
		"range" : grange.Run,
	}

	multitool.Main("goblin", tools)
}
