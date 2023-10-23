package main

import (
	"github.com/reklesio/mtool"
	
	"github.com/reklesio/tk/tool/cat"
	"github.com/reklesio/tk/tool/date"
	"github.com/reklesio/tk/tool/ec"
	"github.com/reklesio/tk/tool/echo"
	"github.com/reklesio/tk/tool/ftest"
	"github.com/reklesio/tk/tool/gfalse"
	"github.com/reklesio/tk/tool/grange"
	"github.com/reklesio/tk/tool/gtrue"
	"github.com/reklesio/tk/tool/in"
	"github.com/reklesio/tk/tool/ln"
	"github.com/reklesio/tk/tool/ls"
	"github.com/reklesio/tk/tool/mergelbl"
	"github.com/reklesio/tk/tool/mkdir"
	"github.com/reklesio/tk/tool/noext"
	"github.com/reklesio/tk/tool/paths"
	"github.com/reklesio/tk/tool/quote"
	"github.com/reklesio/tk/tool/read"
	"github.com/reklesio/tk/tool/sort"
	"github.com/reklesio/tk/tool/tac"
	"github.com/reklesio/tk/tool/uniq"
	"github.com/reklesio/tk/tool/urlprs"
	"github.com/reklesio/tk/tool/useprog"
	"github.com/reklesio/tk/tool/wc"
	"github.com/reklesio/tk/tool/whoami"
	"github.com/reklesio/tk/tool/yes"
)

func main() {
	tools := mtool.Tools{
		"cat":      mtool.Tool{cat.Run, "print file data to the standard output", ""},
		"mkdir":    mtool.Tool{mkdir.Run, "make new directory", ""},
		"echo":     mtool.Tool{echo.Run, "print strings to the standard output", ""},
		"true":     mtool.Tool{gtrue.Run, "exit with true status", ""},
		"false":    mtool.Tool{gfalse.Run, "exit with false status", ""},
		"sort":     mtool.Tool{sort.Run, "sort strings inputed from standard input", ""},
		"tac":      mtool.Tool{tac.Run, "print strings from standard input in reversed order", ""},
		"ls":       mtool.Tool{ls.Run, "list directory content", ""},
		"yes":      mtool.Tool{yes.Run, "print string infinite/exact amount times", ""},
		"date":     mtool.Tool{date.Run, "print current date", ""},
		"uniq":     mtool.Tool{uniq.Run, "filter repeated strings", ""},
		"quote":    mtool.Tool{quote.Run, "quote words containing space characters", ""},
		"urlprs":   mtool.Tool{urlprs.Run, "parse URLs", ""},
		"noext":    mtool.Tool{noext.Run, "print file path without extension", ""},
		"mergelbl": mtool.Tool{mergelbl.Run, "merge line by line", ""},
		"ec":       mtool.Tool{ec.Run, "render escape sequences", ""},
		"read":     mtool.Tool{read.Run, "read lines and exit", ""},
		"wc":       mtool.Tool{wc.Run, "count words, bytes, runes etc", ""},
		"ftest":    mtool.Tool{ftest.Run, "filter files by specified features", ""},
		"range":    mtool.Tool{grange.Run, "too lazy", ""},
		"in":       mtool.Tool{in.Run, "filter strings from stdin that aren not in arguments", ""},
		"which":  mtool.Tool{useprog.Run, "print the name or the path of the first existing program in arg list", ""},
		"paths": mtool.Tool{
			paths.Run,
			"convert UNIX slash separated paths into the OS compatible ones",
			"",
		},
		"whoami": mtool.Tool{
			whoami.Run,
			"print current user name",
			"",
		},
		"ln": mtool.Tool{
			ln.Run,
			"link files",
			"",
		},
	}

	mtool.Main("tk", tools)
}
