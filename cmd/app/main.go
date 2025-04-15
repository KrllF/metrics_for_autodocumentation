package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/KrllF/metrics_for_autodocumentation/internal/app"
)

func main() {
	sourceFile := flag.String("source", "", "the path to the source file")
	mdFile := flag.String("md", "", "the path to the md file")

	flag.Parse()

	a := app.NewApp()

	stat, err := a.Run(*sourceFile, *mdFile)
	if err != nil {
		log.Printf("a.Run: %v", err)

		return
	}
	fmt.Printf("количество непокрытых функций: %d,\nколичество некорректных функций в md: %d,\nпокрытие: %v",
		stat.UnCovered, stat.InCorrect, stat.Coverage)
}
