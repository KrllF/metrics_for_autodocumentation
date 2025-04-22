package golang

import "github.com/KrllF/metrics_for_autodocumentation/internal/entity"

// Coverage возвращает
// 1. количество непокрытых функций
// 2. количество некорректных функций в md файле
// 3. процент покрытия (отношение количества покрытых функций ко всем функциям)
func (s *Service) Coverage(sourceFile, mdFile map[string]struct{}) entity.Stat {
	var uncovered, incorrect int64
	var coverage float64

	for funcName := range sourceFile {
		if _, exists := mdFile[funcName]; !exists {
			uncovered++
		}
	}

	for funcName := range mdFile {
		if _, exists := sourceFile[funcName]; !exists {
			incorrect++
		}
	}

	totalFunctions := len(sourceFile)
	if totalFunctions > 0 {
		covered := int64(totalFunctions) - uncovered
		coverage = float64(covered) / float64(totalFunctions) * 100
		return entity.Stat{UnCovered: uncovered, InCorrect: incorrect, Coverage: coverage}
	}
	coverage = 100

	return entity.Stat{UnCovered: uncovered, InCorrect: incorrect, Coverage: coverage}
}
