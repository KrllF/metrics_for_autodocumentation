package cli

import (
	"fmt"
	"sync"

	"github.com/KrllF/metrics_for_autodocumentation/internal/entity"
)

func (c *Handler) Run(sourceFile, mdFile string) (entity.StructStat, entity.Stat, error) {
	var (
		structStat entity.StructStat
		// stat       entity.Stat
		err1 error
		// err2       error
	)

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		structStat, err1 = c.servStruct.EqualStruct(sourceFile, mdFile)
	}()

	// go func() {
	// 	defer wg.Done()
	// 	stat, err2 = c.servGo.GetMetrics(sourceFile, mdFile)
	// }()

	wg.Wait()

	if err1 != nil {
		return entity.StructStat{}, entity.Stat{}, fmt.Errorf("c.servStruct.EqualStruct: %w", err1)
	}
	// if err2 != nil {
	// 	return entity.StructStat{}, entity.Stat{}, fmt.Errorf("c.servGo.GetMetrics: %w", err2)
	// }

	return structStat, entity.Stat{}, nil
}
