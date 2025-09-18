package commands

import (
	"redis/resp"
	"sync"
)

var Handlers = map[string]func([]resp.Val) resp.Val{
	"PING": ping,
	"SET":  set,
	"GET":  get,
	"HSET": hset,
	"HGET": hget,
}

func ping(args []resp.Val) resp.Val {
	if len(args) == 0 {
		return resp.Val{Typ: "String", Str: "PONG"}
	}

	return resp.Val{Typ: "String", Str: args[0].Bulk}
}

var SETs = map[string]string{}
var SETsMu = sync.RWMutex{}

func set(args []resp.Val) resp.Val {
	if len(args) != 2 {
		return resp.Val{Typ: "Error", Str: "ERR wrong number of arguments for 'set' command"}
	}

	key := args[0].Bulk
	value := args[1].Bulk

	SETsMu.Lock()
	SETs[key] = value
	SETsMu.Unlock()

	return resp.Val{Typ: "String", Str: "OK"}
}

func get(args []resp.Val) resp.Val {
	if len(args) != 1 {
		return resp.Val{Typ: "Error", Str: "ERR wrong number of arguments for 'get' command"}
	}

	key := args[0].Bulk

	SETsMu.RLock()
	value, ok := SETs[key]
	SETsMu.RUnlock()

	if !ok {
		return resp.Val{Typ: "Null"}
	}

	return resp.Val{Typ: "Bulk", Bulk: value}
}

var HSETs = map[string]map[string]string{}
var HSETsMu = sync.RWMutex{}

func hset(args []resp.Val) resp.Val {
	if len(args) != 3 {
		return resp.Val{Typ: "Error", Str: "ERR wrong number of arguments for 'hset' command"}
	}

	hash := args[0].Bulk
	key := args[1].Bulk
	value := args[2].Bulk

	HSETsMu.Lock()
	if _, ok := HSETs[hash]; !ok {
		HSETs[hash] = map[string]string{}
	}
	HSETs[hash][key] = value
	HSETsMu.Unlock()

	return resp.Val{Typ: "String", Str: "OK"}
}

func hget(args []resp.Val) resp.Val {
	if len(args) != 2 {
		return resp.Val{Typ: "Error", Str: "ERR wrong number of arguments for 'hget' command"}
	}

	hash := args[0].Bulk
	key := args[1].Bulk

	HSETsMu.RLock()
	value, ok := HSETs[hash][key]
	HSETsMu.RUnlock()

	if !ok {
		return resp.Val{Typ: "Null"}
	}

	return resp.Val{Typ: "Bulk", Bulk: value}
}
