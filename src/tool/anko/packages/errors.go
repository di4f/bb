package packages

import (
	"errors"
	"reflect"

	"github.com/surdeus/goblin/src/tool/anko/env"
)

func init() {
	env.Packages["errors"] = map[string]reflect.Value{
		"New": reflect.ValueOf(errors.New),
	}
}
