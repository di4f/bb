package packages

import (
	"encoding/json"
	"reflect"

	"github.com/surdeus/goblin/src/tool/anko/env"
)

func init() {
	env.Packages["encoding/json"] = map[string]reflect.Value{
		"Marshal":   reflect.ValueOf(json.Marshal),
		"Unmarshal": reflect.ValueOf(json.Unmarshal),
	}
}
