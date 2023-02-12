package path

import (
	"path"
	"strings"

	"github.com/ethanmidgley/quick-folders/pkg/open"
)

type Config struct {
	RootDirectory string       `yaml:"rootDirectory"`
	Paths         map[string]P `yaml:"paths"`
}

type P struct {
	Path     string       `yaml:"path"`
	Absolute bool         `yaml:"absolute"`
	Children map[string]P `yaml:"children"`
}

func (c *Config) OpenPath(key string) {
	// get key split by -
	keys := strings.Split(key, "-")
	if p, has := c.Paths[keys[0]]; has {
		absPath := p.Path
		if !p.Absolute {
			absPath = path.Join(c.RootDirectory, p.Path)
		}
		p.OpenPath(keys[1:], absPath)
	}

}

func (p *P) OpenPath(key []string, filepath string) {

	if len(key) == 0 {
		// we are fully recursed
		fullPath := filepath
		if p.Absolute {
			fullPath = p.Path
		}
		open.OpenFolder(fullPath)
	} else {
		if p, has := p.Children[key[0]]; has {
			absPath := p.Path
			if !p.Absolute {
				absPath = path.Join(filepath, p.Path)
			}
			p.OpenPath(key[1:], absPath)
		}

	}

}
