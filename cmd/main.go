package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/ethanmidgley/quick-folders/pkg/config"
)

func main() {
	exPath, _ := os.Executable()
	conf, err := config.Load(filepath.Join(exPath, "./config.yaml"))
	if err != nil {
		log.Panic(err)
	}

	conf.OpenPath(os.Args[1])

}
