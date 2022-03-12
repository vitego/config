package config

import (
	"github.com/valyala/fastjson"
	"strings"
)

var (
	// store is the main config storage, we use a second map for cache store's
	// result to get much faster the called value.
	// First call (with store) : "Vitego" 28.146µs
	// Second call (with cache) : "Vitego" 8.66µs
	store map[string]*fastjson.Value
	cache map[string]string
)

func Get(path string) string {
	if cache[path] == "" {
		cache[path] = walkInKeys(path)
	}

	return cache[path]
}

func walkInKeys(path string) string {
	k := strings.Split(path, ".")
	return string(store[k[0]].GetStringBytes(k[1:]...))
}
