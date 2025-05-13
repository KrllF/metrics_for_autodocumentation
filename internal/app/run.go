package app

import (
	"fmt"
)

func (a *App) Run(sourceFile, mdFile string) error {
	structStat, stat, err := a.hand.Run(sourceFile, mdFile)
	if err != nil {
		return fmt.Errorf("a.hand.Run: %w", err)
	}
	fmt.Printf("\nудалось ли покрыть %v, процент покрытия сгенерированными файлами %v\n\n", structStat.OkCoverageStruct, structStat.CoverageStruct)
	fmt.Printf("количество непокрытых функций: %d,\nколичество некорректных функций в md: %d,\nпокрытие: %v",
		stat.UnCovered, stat.InCorrect, stat.Coverage)

	return nil
}
