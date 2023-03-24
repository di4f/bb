package main

import (
	"github.com/surdeus/goblin/src/tool/awk"
	"github.com/surdeus/goblin/src/tool/basename"
	"github.com/surdeus/goblin/src/tool/cat"
	"github.com/surdeus/goblin/src/tool/date"
	"github.com/surdeus/goblin/src/tool/ec"
	"github.com/surdeus/goblin/src/tool/echo"
	"github.com/surdeus/goblin/src/tool/ftest"
	"github.com/surdeus/goblin/src/tool/gfalse"
	"github.com/surdeus/goblin/src/tool/grange"
	"github.com/surdeus/goblin/src/tool/gtrue"
	"github.com/surdeus/goblin/src/tool/in"
	"github.com/surdeus/goblin/src/tool/ln"
	"github.com/surdeus/goblin/src/tool/ls"
	"github.com/surdeus/goblin/src/tool/mergelbl"
	"github.com/surdeus/goblin/src/tool/mk"
	"github.com/surdeus/goblin/src/tool/mkdir"
	"github.com/surdeus/goblin/src/tool/noext"
	"github.com/surdeus/goblin/src/tool/path"
	"github.com/surdeus/goblin/src/tool/paths"
	"github.com/surdeus/goblin/src/tool/quote"
	"github.com/surdeus/goblin/src/tool/read"
	"github.com/surdeus/goblin/src/tool/sort"
	"github.com/surdeus/goblin/src/tool/tac"
	"github.com/surdeus/goblin/src/tool/uniq"
	"github.com/surdeus/goblin/src/tool/urlprs"
	"github.com/surdeus/goblin/src/tool/useprog"
	"github.com/surdeus/goblin/src/tool/wc"
	"github.com/surdeus/goblin/src/tool/whoami"
	"github.com/surdeus/goblin/src/tool/yes"
	"github.com/surdeus/gomtool/src/mtool"
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
		"useprog":  mtool.Tool{useprog.Run, "print the name of the first existing program in arg list", ""},
		"path":     mtool.Tool{path.Run, "print cross platform path based on cmd arguments", ""},
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
	}

	mtool.Main("goblin", tools)
}
