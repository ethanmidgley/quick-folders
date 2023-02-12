package main

import (
	"log"
	"os"

	"github.com/ethanmidgley/quick-folders/pkg/config"
)

func main() {

	conf, err := config.Load("./config.yaml")
	if err != nil {
		log.Panic(err)
	}

	conf.OpenPath(os.Args[1])

}
