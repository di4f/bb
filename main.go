package main

import (
	"fmt"
	"github.com/omnipunk/cli/mtool"

	"github.com/omnipunk/bb/tool/cat"
	"github.com/omnipunk/bb/tool/date"
	"github.com/omnipunk/bb/tool/ec"
	"github.com/omnipunk/bb/tool/echo"
	"github.com/omnipunk/bb/tool/ftest"
	"github.com/omnipunk/bb/tool/grange"
	"github.com/omnipunk/bb/tool/in"
	"github.com/omnipunk/bb/tool/ln"
	"github.com/omnipunk/bb/tool/ls"
	"github.com/omnipunk/bb/tool/mergelbl"
	"github.com/omnipunk/bb/tool/mkdir"
	"github.com/omnipunk/bb/tool/paths"
	"github.com/omnipunk/bb/tool/quote"
	"github.com/omnipunk/bb/tool/read"
	"github.com/omnipunk/bb/tool/sort"
	"github.com/omnipunk/bb/tool/tac"
	"github.com/omnipunk/bb/tool/uniq"
	"github.com/omnipunk/bb/tool/urlprs"
	"github.com/omnipunk/bb/tool/useprog"
	"github.com/omnipunk/bb/tool/wc"
	"github.com/omnipunk/bb/tool/whoami"
	"github.com/omnipunk/bb/tool/yes"
	"os"
)

var root = mtool.T("bb").Subs(
	mtool.T("cat").Func(cat.Run).Desc(
		"concatenate files",
	).Usage(
		"[file1 file2 ...fileN]",
	),
	mtool.T("mkdir").Func(mkdir.Run).Desc(
		"make new directories",
	).Usage(
		"<dir1 [dir2 dir3 ...dirN]>",
	),
	mtool.T("echo").Func(echo.Run).Desc(
		"print strings",
	).Usage(
		"[str1 str2 ...strN]",
	),
	mtool.T("true").Func(func(flags *mtool.Flags) {
		os.Exit(0)
	}).Desc("exit successfuly"),
	mtool.T("false").Func(func(flags *mtool.Flags) {
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
	).Usage(
		"[fileDir1 fileDir2 ...fileDirN]",
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
	mtool.T("env").Func(func(flags *mtool.Flags) {
		flags.Parse()
		envs := os.Environ()
		for _, env := range envs {
			fmt.Println(env)
		}
	}).Desc(
		"print all the environment variables",
	),
).Desc(
	"not POSIX compatible BusyBox utilities",
)

func main() {
	root.Run(os.Args[1:])
}
