package main

import (
	"os"

	"github.com/singhdurgesh/rednote/cmd/rednote"
	"github.com/singhdurgesh/rednote/cmd/worker"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "worker" {
		worker.Init()
	} else {
		rednote.Init()
	}
}
