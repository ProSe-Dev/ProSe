package main

import (
	"log"

	"github.com/ProSe-Dev/prose/prose/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
