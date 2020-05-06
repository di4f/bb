package ls
import(
	"os"
	"fmt"
	"flag"
	"strings"
)

func IsDir(p string) (bool, error) {
	finfo, e := os.Stat(p)
	if e != nil {
		return false, e
	}
	return finfo.IsDir(), e
}

func ReadDir(p string) ([]os.FileInfo, error) {
	f, e := os.Open(p)
	if e != nil {
		return nil, e
	}
	l, e := f.Readdir(-1)
	if e != nil {
			return nil, e
	}
	f.Close()

	return l, nil
}

func Stat(p string) (os.FileInfo, error) {
	f, e := os.Open(p)
	if e != nil {
		return nil, e
	}
	s, e := f.Stat()
	f.Close()
	return s, nil
}

func ls(p string, fold int) error {
	isDir, e := IsDir(p)
	if e != nil {
		return e
	}

	pp := strings.TrimRight(p, "/")

	if !isDir {
		fmt.Println(pp)
		return nil
	}
	
	if fold>0 { 
		/* It's a directory and it can be ls-ed. */
		l, e := ReadDir(pp)
		if e!=nil {
			return e
		}
		for _, f := range l {
			s := pp+"/"+f.Name()
			if b, _ := IsDir(s) ; b && fold!=1 {
				fmt.Println(s)
			}
			ls(s, fold-1)
		}
	} else {
		/* It's finish directory. Fold==0 or fold<0. */
		fmt.Println(pp)
	}
	
	return nil
}

func Run(args []string) int {
	status := 0
	arg0 := args[0]
	args = args[1:]
	flagSet := flag.NewFlagSet(arg0, flag.ExitOnError)
	var foldLvl int
	flagSet.IntVar(&foldLvl, "r", 1, "List recursively with choosing deepness, can't be negative or zero.")
	flagSet.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s: %s [files]\n", arg0, arg0)
		flagSet.PrintDefaults()
	}
	flagSet.Parse(args)
	args = flagSet.Args()

	if foldLvl<0 {
		flagSet.Usage()
		return 1
	}

	if foldLvl==0 && len(args)==0 {
		flagSet.Usage()
		return 1
	}
	
	if len(args) == 0 {
		foldLvl -= 1
		if l, e := ReadDir(".") ; e != nil {
			status = 1
			fmt.Fprintf(os.Stderr, "%s: %s.\n", arg0, e)
		} else {
			for _, f := range l {
				isDir, _ := IsDir(f.Name())

				fmt.Println(f.Name())
				if isDir && foldLvl>0 {
					e := ls(f.Name(), foldLvl)
					if e!=nil {
						status = 1
						fmt.Fprintf(os.Stderr, "%s: %s.\n", arg0, e)
					}
				}
			}
		}
	} else {
		for _, p := range args {
			e := ls(p, foldLvl)
			if e != nil {
				status = 1
				fmt.Fprintf(os.Stderr, "%s: %s.\n", arg0, e)
			}
			
		}
	}
	return status
}
