package main

import (
	"github.com/omnipunk/cli/mtool"
	
	"github.com/omnipunk/tk/tool/cat"
	"github.com/omnipunk/tk/tool/date"
	"github.com/omnipunk/tk/tool/ec"
	"github.com/omnipunk/tk/tool/echo"
	"github.com/omnipunk/tk/tool/ftest"
	"github.com/omnipunk/tk/tool/grange"
	"github.com/omnipunk/tk/tool/in"
	"github.com/omnipunk/tk/tool/ln"
	"github.com/omnipunk/tk/tool/ls"
	"github.com/omnipunk/tk/tool/mergelbl"
	"github.com/omnipunk/tk/tool/mkdir"
	"github.com/omnipunk/tk/tool/paths"
	"github.com/omnipunk/tk/tool/quote"
	"github.com/omnipunk/tk/tool/read"
	"github.com/omnipunk/tk/tool/sort"
	"github.com/omnipunk/tk/tool/tac"
	"github.com/omnipunk/tk/tool/uniq"
	"github.com/omnipunk/tk/tool/urlprs"
	"github.com/omnipunk/tk/tool/useprog"
	"github.com/omnipunk/tk/tool/wc"
	"github.com/omnipunk/tk/tool/whoami"
	"github.com/omnipunk/tk/tool/yes"
	"os"
)

var root = mtool.T("tk").Subs(
	mtool.T("cat").Func(cat.Run).Desc(
		"concatenate files",
	),
	mtool.T("mkdir").Func(mkdir.Run).Desc(
		"make new directories",
	),
	mtool.T("echo").Func(echo.Run).Desc(
		"print strings",
	),
	mtool.T("true").Func(func(flags *mtool.Flags){
		os.Exit(0)
	}).Desc("exit successfuly"),
	mtool.T("false").Func(func(flags *mtool.Flags){
		os.Exit(1)
	}).Desc("exit with failure"),
	mtool.T("sort").Func(sort.Run).Desc(
		"sort strings",
	),
	mtool.T("tac").Func(tac.Run).Desc(
		"reversed cat",
	),
	mtool.T("ls").Func(ls.Run).Desc(
		"list files",
	),
	mtool.T("yes").Func(yes.Run).Desc(
		"repeat string",
	),
	mtool.T("date").Func(date.Run).Desc(
		"print date",
	),
	mtool.T("uniq").Func(uniq.Run).Desc(
		"filter repeated strings",
	),
	mtool.T("quote").Func(quote.Run).Desc(
		"quote strings with spaces",
	),
	mtool.T("urlprs").Func(urlprs.Run).Desc(
		"parse URL",
	),
	mtool.T("read").Func(read.Run).Desc(
		"read lines",
	),
	mtool.T("ec").Func(ec.Run).Desc(
		"render escape characters",
	),
	mtool.T("lbl").Func(mergelbl.Run).Desc(
		"merge files line by line",
	),
	mtool.T("ftest").Func(ftest.Run).Desc(
		"filter files",
	),
	mtool.T("wc").Func(wc.Run).Desc(
		"word, rune, byte counts",
	),
	mtool.T("range").Func(grange.Run).Desc(
		"print num range",
	),
	mtool.T("in").Func(in.Run).Desc(
		"print only strings that are in arguments",
	),
	mtool.T("which").Func(useprog.Run).Desc(
		"print path to executable",
	),
	mtool.T("whoami").Func(whoami.Run).Desc(
		"print your username",
	),
	mtool.T("ln").Func(ln.Run).Desc(
		"link files",
	),
	mtool.T("paths").Func(paths.Run).Desc(
		"print different parts of paths",
	),
).Desc(
	"ToolKit, BusyBox-like not POSIX-compatible utilities",
)

func main() {
	root.Run(os.Args[1:])
}
