package app

import (
	"fmt"
)

func (a *App) Run(sourceFile, mdFile string) error {
	stat, err := a.hand.Run(sourceFile, mdFile)
	if err != nil {
		return fmt.Errorf("a.hand.Run: %w", err)
	}
	fmt.Printf("количество непокрытых функций: %d,\nколичество некорректных функций в md: %d,\nпокрытие: %v",
		stat.UnCovered, stat.InCorrect, stat.Coverage)

	return nil
}
