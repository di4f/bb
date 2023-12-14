package ls
import(
	"os"
	"fmt"
	"strings"
	"regexp"
	"path"
	"github.com/di4f/cli/mtool"
)

var(
	listHidden bool
	args []string
)

var slashRegexp = regexp.MustCompile("/+")

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

func
Stat(p string) (os.FileInfo, error) {
	f, e := os.Open(p)
	if e != nil {
		return nil, e
	}
	s, e := f.Stat()
	f.Close()
	return s, nil
}

func findString(slice []string, val string) (int, bool) {
    for i, v := range slice {
        if v == val {
            return i, true
        }
    }
    return -1, false
}

func
isHidden(p string) bool {
	if path.Base(p)[0] == '.' {
		return true
	}
	return false
}

func shouldList(p string) bool {
	if _, found := findString(args, p) ; found {
		return true
	}
	if !listHidden && isHidden(p) { 
		return false
	}
	return true
}

func
deleteExceedSlashes(p string) string {
	p = slashRegexp.ReplaceAllString(p, "/")
	if p != "/" { // Do not trim if it is root dir.
		p = strings.TrimRight(p, "/")
	}
	return p
}

func
ls(p string, fold int) error {
	if !shouldList(p) {
		return nil
	}
	
	p = deleteExceedSlashes(p)
	
	isDir, e := IsDir(p)
	if e != nil {
		return e
	}


	if !isDir {
		fmt.Println(p)
		return nil
	}
	
	if fold>0 { 
		/* It's a directory and it can be ls-ed. */
		l, e := ReadDir(p)
		if e!=nil {
			return e
		}
		for _, f := range l {
			s := p+"/"+f.Name()
			if b, _ := IsDir(s) ; b && fold!=1 {
				fmt.Println(s)
			}
			ls(s, fold-1)
		}
	} else {
		/* It's finish directory. Fold==0 or fold<0. */
		fmt.Println(p)
	}
	
	return nil
}

func Run(flagSet *mtool.Flags) {
	var foldLvl int
	flagSet.IntVar(&foldLvl, "r", 1, "List recursively with choosing deepness, can't be negative or zero.")
	flagSet.BoolVar(&listHidden, "a", false, "List hidden files.")
	flagSet.Parse()
	args = flagSet.Args()

	if foldLvl<0 {
		flagSet.Usage()
	}

	if foldLvl==0 && len(args)==0 {
		flagSet.Usage()
	}
	
	if len(args) == 0 {
		foldLvl -= 1
		if l, e := ReadDir(".") ; e != nil {
			fmt.Fprintf(os.Stderr, "%s.\n", e)
		} else {
			for _, f := range l {
				if !shouldList(f.Name()) {
					continue
				}
				isDir, _ := IsDir(f.Name())

				fmt.Println(f.Name())
				if isDir && foldLvl>0 {
					e := ls(f.Name(), foldLvl)
					if e!=nil {
						fmt.Fprintf(os.Stderr,
							"%s\n", e)
					}
				}
			}
		}
	} else {
		for _, p := range args {
			e := ls(p, foldLvl)
			if e != nil {
				fmt.Fprintf(os.Stderr, "%s\n", e)
			}
			
		}
	}
}
