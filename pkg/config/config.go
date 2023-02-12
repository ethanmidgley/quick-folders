package config

import (
	"os"

	p "github.com/ethanmidgley/quick-folders/pkg/path"
	"gopkg.in/yaml.v3"
)

func Load(path string) (*p.Config, error) {

	c := p.Config{}

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	d := yaml.NewDecoder(file)

	if err = d.Decode(&c); err != nil {
		return nil, err
	}

	return &c, err

}
