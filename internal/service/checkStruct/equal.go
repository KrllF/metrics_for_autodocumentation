package checkStruct

import (
	"log"
	"path/filepath"
	"slices"
	"sort"
	"strings"

	"github.com/KrllF/metrics_for_autodocumentation/internal/entity"
)

func (s *Service) EqualStruct(sourceFile, mdFile string) (entity.StructStat, error) {
	sourceRaw := s.ListDirByWalk(sourceFile, true)
	mdRaw := s.ListDirByWalk(mdFile, false)

	sourceM := stripRoot(sourceRaw)
	mdM := stripRoot(mdRaw)

	equal := true
	var totalFiles, coveredFiles int

	log.Println("исходный:")
	log.Println(sourceFile)
	log.Println("сгенерированный:")
	log.Println(mdFile)

	for path, sourceFiles := range sourceM {
		sourceNames := stripExtensions(sourceFiles)
		totalFiles += len(sourceNames)

		mdFiles, ok := mdM[path]
		mdNames := stripExtensions(mdFiles)

		if !ok {
			log.Printf("Отсутствует директория в документации: %s\n", path)
			equal = false
			continue
		}

		mdSet := make(map[string]struct{}, len(mdNames))
		for _, name := range mdNames {
			mdSet[name] = struct{}{}
		}

		for _, name := range sourceNames {
			if _, found := mdSet[name]; found {
				coveredFiles++
			}
		}

		sort.Strings(sourceNames)
		sort.Strings(mdNames)

		if !slices.Equal(sourceNames, mdNames) {
			log.Printf("Несовпадение файлов в директории %s:\n", path)
			log.Printf("Ожидается: %v\n", sourceNames)
			log.Printf("Найдено:   %v\n", mdNames)
			equal = false
		}
	}

	for path := range mdM {
		if _, ok := sourceM[path]; !ok {
			log.Printf("Лишняя директория в документации: %s\n", path)
			equal = false
		}
	}

	var coverage float64
	if totalFiles > 0 {
		coverage = float64(coveredFiles) / float64(totalFiles)
	} else {
		coverage = 1.0
	}

	return entity.StructStat{CoverageStruct: coverage, OkCoverageStruct: equal}, nil
}

func stripRoot(m map[string][]string) map[string][]string {
	newMap := make(map[string][]string)
	for path, files := range m {
		parts := strings.SplitN(path, string(filepath.Separator), 2)
		relPath := ""
		if len(parts) == 2 {
			relPath = parts[1]
		}
		newMap[relPath] = files
	}

	return newMap
}

func stripExtensions(files []string) []string {
	res := make([]string, 0, len(files))
	for _, f := range files {
		name := strings.TrimSuffix(f, filepath.Ext(f))
		res = append(res, name)
	}

	return res
}
