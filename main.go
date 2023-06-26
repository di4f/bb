package main

import (
	"github.com/mojosa-software/gomtool/src/mtool"
	
	"github.com/mojosa-software/goblin/src/tool/awk"
	"github.com/mojosa-software/goblin/src/tool/basename"
	"github.com/mojosa-software/goblin/src/tool/cat"
	"github.com/mojosa-software/goblin/src/tool/date"
	"github.com/mojosa-software/goblin/src/tool/ec"
	"github.com/mojosa-software/goblin/src/tool/echo"
	"github.com/mojosa-software/goblin/src/tool/ftest"
	"github.com/mojosa-software/goblin/src/tool/gfalse"
	"github.com/mojosa-software/goblin/src/tool/grange"
	"github.com/mojosa-software/goblin/src/tool/gtrue"
	"github.com/mojosa-software/goblin/src/tool/in"
	"github.com/mojosa-software/goblin/src/tool/ln"
	"github.com/mojosa-software/goblin/src/tool/ls"
	"github.com/mojosa-software/goblin/src/tool/mergelbl"
	"github.com/mojosa-software/goblin/src/tool/mk"
	"github.com/mojosa-software/goblin/src/tool/mkdir"
	"github.com/mojosa-software/goblin/src/tool/noext"
	"github.com/mojosa-software/goblin/src/tool/paths"
	"github.com/mojosa-software/goblin/src/tool/quote"
	"github.com/mojosa-software/goblin/src/tool/read"
	"github.com/mojosa-software/goblin/src/tool/sort"
	"github.com/mojosa-software/goblin/src/tool/tac"
	"github.com/mojosa-software/goblin/src/tool/uniq"
	"github.com/mojosa-software/goblin/src/tool/urlprs"
	"github.com/mojosa-software/goblin/src/tool/useprog"
	"github.com/mojosa-software/goblin/src/tool/wc"
	"github.com/mojosa-software/goblin/src/tool/whoami"
	"github.com/mojosa-software/goblin/src/tool/yes"
	"github.com/mojosa-software/goblin/src/tool/script"
)

func main() {
	tools := mtool.Tools{
		"basename": mtool.Tool{basename.Run, "get base name of file path", ""},
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
		"mk":       mtool.Tool{mk.Run, "file dependency system, simpler make", ""},
		"awk":      mtool.Tool{awk.Run, "simple scripting language for working with string templates", ""},
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
		"script": mtool.Tool{
			script.Run,
			"run embedded anko",
			"",
		},
	}

	mtool.Main("goblin", tools)
}
