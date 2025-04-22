package checkStruct

import (
	"slices"
	"sort"
)

func (s *Service) EqualStruct(sourceFile, mdFile string) (bool, error) {
	sourceM := s.ListDirByWalk(sourceFile)
	mdM := s.ListDirByWalk(mdFile)

	// TODO: на данный момент только
	// базовая проверка на равенство структур
	// + не оптимально
	// сгенерированного файла и исходного проекта
	// просто по ключам в получившихся мапах
	// при обходе ListDirByWalk
	keySource := make([]string, len(sourceM))
	keyM := make([]string, len(mdM))

	for key := range sourceM {
		keySource = append(keySource, key)
	}

	for key := range mdM {
		keyM = append(keyM, key)
	}

	if len(keySource) != len(keyM) {
		return false, nil
	}

	sort.Strings(keySource)
	sort.Strings(keyM)

	return slices.Equal(keyM, keySource), nil
}
