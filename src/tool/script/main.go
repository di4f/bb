// +build !appengine

package script

import (
	"bufio"
	//"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"

	"github.com/surdeus/goscript/core"
	"github.com/surdeus/goscript/env"
	_ "github.com/surdeus/goscript/packages"
	"github.com/surdeus/goscript/parser"
	"github.com/surdeus/goscript/vm"
	"github.com/surdeus/gomtool/src/mtool"
	"os/exec"
)

type Command struct {
	input io.Reader
	*exec.Cmd
	output io.Writer
}

const version = "0.1.8"

var (
	flagExecute string
	file        string
	args        []string
	e           *env.Env
	//flag *mtool.Flags
)

func Run(flagSet *mtool.Flags) {
	
	printVersion := flagSet.Bool("v", false, "prints out the version and then exits")
	flagSet.StringVar(&flagExecute, "e", "", "execute the Anko code")
	flagSet.Parse()
	args := flagSet.Args()
	
	if *printVersion {
		fmt.Println(version)
		os.Exit(0)
	}
	
	setupEnv()
	if flagExecute != "" {
		exitCode := runNonInteractive()
		os.Exit(exitCode)
	}
	
	if len(args) < 1 {
		file = "#stdin"
	} else {
		file = args[0]
	}
	
	
	exitCode := runNonInteractive()
	os.Exit(exitCode)
}

func Cmd(args ...string) *Command {
		if len(args) < 1 {
			panic("too few arguments")
		}
		
		cmd := exec.Command(args[0], args[1:]...)
		
		return &Command{ Cmd: cmd }
	}

func (cmd *Command)IO(input io.Reader, output io.Writer) *Command {
	cmd.input = input
	cmd.output = output
	return cmd
}

func (cmd *Command) Run() error {
		input := cmd.input
		output := cmd.output
		
		stdin, err := cmd.StdinPipe()
		if err != nil {
			return err
		}
		
		stdout, err := cmd.StdoutPipe()
		if err != nil {
			return err
		}
		
		err = cmd.Start()
		if err != nil {
			return err
		}
		go func(){
			_, err = io.Copy(stdin, input)
			if err != nil {
				panic(err)
			}
		}()
		
		go func() {
			_, err = io.Copy(output, stdout)
			if err != nil {
				panic(err)
			}
		}()
	
		err = cmd.Wait()
		if err != nil {
			fmt.Println(err)
			return err
		}
		
		return nil
}

func setupEnv() {
	e = env.NewEnv()
	e.Define("Args", args)
	e.Define("Cmd", Cmd)
	e.Define("Rcmd", func(args ...string) bool {
		cmd := Cmd(args...).IO(os.Stdin, os.Stdout)
		return cmd.Run() == nil
	})
	core.Import(e)
}

func runNonInteractive() int {
	var source string
	if flagExecute != "" {
		source = flagExecute
	} else {
		var (
			sourceBytes []byte
			err error
		)
		if file == "#stdin" {
			sourceBytes, err = io.ReadAll(os.Stdin)
		} else {
			sourceBytes, err = ioutil.ReadFile(file)
		}
		if err != nil {
			fmt.Println("ReadFile error:", err)
			return 2
		}
		source = string(sourceBytes)
	}

	_, err := vm.Execute(e, nil, source)
	if err != nil {
		fmt.Println("Execute error:", err)
		return 4
	}

	return 0
}

func runInteractive() int {
	var following bool
	var source string
	scanner := bufio.NewScanner(os.Stdin)

	parser.EnableErrorVerbose()

	for {
		if following {
			source += "\n"
			fmt.Print("  ")
		} else {
			fmt.Print("> ")
		}

		if !scanner.Scan() {
			break
		}
		source += scanner.Text()
		if source == "" {
			continue
		}
		if source == "quit()" {
			break
		}

		stmts, err := parser.ParseSrc(source)

		if e, ok := err.(*parser.Error); ok {
			es := e.Error()
			if strings.HasPrefix(es, "syntax error: unexpected") {
				if strings.HasPrefix(es, "syntax error: unexpected $end,") {
					following = true
					continue
				}
			} else {
				if e.Pos.Column == len(source) && !e.Fatal {
					fmt.Fprintln(os.Stderr, e)
					following = true
					continue
				}
				if e.Error() == "unexpected EOF" {
					following = true
					continue
				}
			}
		}

		following = false
		source = ""
		var v interface{}

		if err == nil {
			v, err = vm.Run(e, nil, stmts)
		}
		if err != nil {
			if e, ok := err.(*vm.Error); ok {
				fmt.Fprintf(os.Stderr, "%d:%d %s\n", e.Pos.Line, e.Pos.Column, err)
			} else if e, ok := err.(*parser.Error); ok {
				fmt.Fprintf(os.Stderr, "%d:%d %s\n", e.Pos.Line, e.Pos.Column, err)
			} else {
				fmt.Fprintln(os.Stderr, err)
			}
			continue
		}

		fmt.Printf("%#v\n", v)
	}

	if err := scanner.Err(); err != nil {
		if err != io.EOF {
			fmt.Fprintln(os.Stderr, "ReadString error:", err)
			return 12
		}
	}

	return 0
}
