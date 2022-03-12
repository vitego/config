package config

import (
	"embed"
	"github.com/valyala/fastjson"
	"path/filepath"
	"strings"
)

func Parse(dir string, files embed.FS) (err error) {
	d, err := files.ReadDir(dir)
	if err != nil {
		return
	}

	store = make(map[string]*fastjson.Value)
	cache = make(map[string]string)

	for _, f := range d {
		var (
			content []byte
			p       fastjson.Parser
		)

		ext := filepath.Ext(f.Name())
		name := strings.TrimRight(f.Name(), ext)

		if f.IsDir() || ext != ".json" {
			continue
		}

		content, err = files.ReadFile(filepath.Join(dir, f.Name()))
		if err != nil {
			continue
		}

		store[name], err = p.Parse(string(content))
		if err != nil {
			return
		}
	}

	return
}
