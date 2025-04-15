package app

import (
	"fmt"

	"github.com/KrllF/metrics_for_autodocumentation/internal/entity"
)

func (a *App) Run(sourceFile, mdFile string) (entity.Stat, error) {
	ret, err := a.hand.Run(sourceFile, mdFile)
	if err != nil {
		return entity.Stat{}, fmt.Errorf("a.hand.Run: %w", err)
	}

	return ret, nil
}
