package input

import (
	"os"
	"io"
	"bufio"
)

func ReadAllRaw(f *os.File) ([]byte, error) {
	ret := make([]byte, 0)
	b := make([]byte, 512)
	for {
		n, e := f.Read(b)
		if n>0 {
			ret = append(ret, b[:n]...)
		}
		if e == io.EOF {
			break
		}else if e != nil {
			return nil, e
		}
	}
	return ret, nil
}

func ReadAllLines(f *os.File)( []string, error ){
	r := bufio.NewReader(f)
	a := make([]string, 0)
	for {
		l, e := r.ReadString('\n')
		if e==io.EOF {
			break
		}else if e != nil {
			return nil, e
		}
		a = append(a, l)
	}
	return a, nil
}
