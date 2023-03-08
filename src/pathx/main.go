package pathx

// The package implements crossplatform structured
// paths.

import (
	"path"
	fp "path/filepath"
	"strings"
)

type Value string
type Path struct {
	Values []Value
	IsAbs  bool
}

func (p Path) Append(values ...string) Path {
	for _, v := range values {
		p.Values = append(p.Values, Value(v))
	}
	return p
}

func From(p string) Path {
	ret := Path{}
	if len(p) == 0 {
		return ret
	}

	/*	p = path.Clean(p)
		if p[0] == '/' {
			ret.IsAbs = true
		} */
	p, _ = strings.CutSuffix(p, "/")
	svalues := strings.Split(p, "/")

	values := make([]Value, len(svalues))
	for i, s := range svalues {
		values[i] = Value(s)
	}
	ret.Values = values

	return ret
}

func FromReal(p string) Path {
	return From(fp.ToSlash(p))
}

func (v Value) IsValid() bool {
	return v.Err() == nil
}

func (v Value) Err() error {
	return nil
}

func (p Path) StringValues() []string {
	ret := []string{}
	if p.IsAbs {
		ret = append(ret, "/")
	}

	for _, v := range p.Values {
		ret = append(ret, string(v))
	}

	return ret
}

func (p Path) Real() string {
	return strings.Join(p.StringValues(), string(fp.Separator))
}

func (p Path) String() string {
	ret := ""
	if p.IsAbs {
		ret = "/"
	}

	ret += path.Join(p.StringValues()...)
	return ret
}

func (p Path) IsValid() bool {
	for _, v := range p.Values {
		if !v.IsValid() {
			return false
		}
	}

	return true
}

func (p Path) Err() error {
	for _, v := range p.Values {
		if !v.IsValid() {
			return v.Err()
		}
	}

	return nil
}
