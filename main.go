package main

import (
	"log"

	"github.com/samhep0803/togo/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
