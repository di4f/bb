package read
/* Plan9 read or something to read into variables. */
import(
	"os"
	"bufio"
	"fmt"
	"github.com/omnipunk/cli/mtool"
)

var(
	nLines int
)

func Run(flagSet *mtool.Flags) {	
	flagSet.IntVar(&nLines, "n", 1, "amount of lines")
	flagSet.Parse()
	//args := flagSet.Args()

	if nLines <= 0 {
		flagSet.Usage()
	}

	rd := bufio.NewReader(os.Stdin)
	for nLines != 0 {
		line, err := rd.ReadString('\n')
		if err != nil {
			os.Exit(0)
		}
		fmt.Print(line)
		nLines -= 1
	}
}
