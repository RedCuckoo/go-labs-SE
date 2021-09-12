package requests

import (
	"github.com/deliveroo/jsonrpc-go"
	"github.com/redcuckoo/go-labs-SE/lab4/resources"
)

// GetUpdateCity returns country, name, population, capital, error
func GetUpdateCity(params resources.Params) (uint64, string, uint64, bool, error) {
	if len(params) != 4 {
		return 0, "", 0, false, jsonrpc.InvalidParams("method expects four param")
	}

	country, ok := params[0].(uint64)
	if !ok {
		return 0, "", 0, false, jsonrpc.InvalidParams("country should be number")
	}

	name, ok := params[1].(string)
	if !ok {
		return 0, "", 0, false, jsonrpc.InvalidParams("name should be string")
	}

	population, ok := params[2].(uint64)
	if !ok {
		return 0, "", 0, false, jsonrpc.InvalidParams("population should be number")
	}

	capital, ok := params[3].(bool)
	if !ok {
		return 0, "", 0, false, jsonrpc.InvalidParams("capital should be boolean")
	}

	return country, name, population, capital, nil
}
