package checkStruct

import (
	"context"
	"slices"
	"sort"
)

func (s *Service) EqualStruct(sourceFile, mdFile string) (bool, error) {
	return s.EqualStructContext(context.Background(), sourceFile, mdFile)
}

func (s *Service) EqualStructContext(ctx context.Context, sourceFile, mdFile string) (bool, error) {
	sourceM, err := s.ListDirByWalkContext(ctx, sourceFile)
	if err != nil {
		return false, err
	}

	mdM, err := s.ListDirByWalkContext(ctx, mdFile)
	if err != nil {
		return false, err
	}

	// Compare structure
	keySource := make([]string, 0, len(sourceM))
	keyMd := make([]string, 0, len(mdM))

	for key := range sourceM {
		keySource = append(keySource, key)
	}

	for key := range mdM {
		keyMd = append(keyMd, key)
	}

	if len(keySource) != len(keyMd) {
		return false, nil
	}

	sort.Strings(keySource)
	sort.Strings(keyMd)

	if !slices.Equal(keyMd, keySource) {
		return false, nil
	}

	// Compare contents (simplified)
	for key, sourceFiles := range sourceM {
		mdFiles := mdM[key]
		if len(sourceFiles) != len(mdFiles) {
			return false, nil
		}

		sort.Strings(sourceFiles)
		sort.Strings(mdFiles)
		if !slices.Equal(mdFiles, sourceFiles) {
			return false, nil
		}
	}

	return true, nil
}
