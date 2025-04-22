package golang

import (
	"fmt"

	"github.com/KrllF/metrics_for_autodocumentation/internal/entity"
)

func (s *Service) GetMetrics(sourceFile, mdFile string) (entity.Stat, error) {
	sour, err := s.GetFileFunc(sourceFile)
	if err != nil {
		return entity.Stat{}, fmt.Errorf("s.GetFileFunc: %w", err)
	}
	md, err := s.GetMarkDownFunc(mdFile)
	if err != nil {
		return entity.Stat{}, fmt.Errorf("s.GetMarkDownFunc: %w", err)
	}

	return s.Coverage(sour, md), nil
}
