package main

import (
	"flag"
	"log"

	"github.com/KrllF/metrics_for_autodocumentation/internal/app"
)

func main() {
	sourceFile := flag.String("source", "", "the path to the source file")
	mdFile := flag.String("md", "", "the path to the md file")

	flag.Parse()

	a := app.NewApp()

	if err := a.Run(*sourceFile, *mdFile); err != nil {
		log.Printf("a.Run: %v", err)

		return
	}
}
