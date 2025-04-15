package cli

import (
	"fmt"

	"github.com/KrllF/metrics_for_autodocumentation/internal/entity"
)

func (c *Handler) Run(sourceFile, mdFile string) (entity.Stat, error) {
	ret, err := c.serv.GetMetrics(sourceFile, mdFile)
	if err != nil {
		return entity.Stat{}, fmt.Errorf("c.serv.GetMetrics: %w", err)
	}

	return ret, nil
}
