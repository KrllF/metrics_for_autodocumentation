package checkStruct

import (
	"fmt"
	"os"
	"path/filepath"
)

func (s *Service) ListDirByWalk(path string, isSource bool) map[string][]string {
	mp := make(map[string][]string)

	path = filepath.Clean(path)
	rootDirName := filepath.Base(path)

	_ = filepath.Walk(path, func(wPath string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("Ошибка при обходе %q: %v\n", wPath, err)
			return err
		}

		if path == wPath {
			return nil
		}

		if info.IsDir() && shouldSkipDir(info.Name()) {
			return filepath.SkipDir
		}

		if !info.IsDir() && isSource && !isSourceFile(info.Name()) {
			return nil
		}

		relPath, err := filepath.Rel(path, wPath)
		if err != nil {
			return err
		}

		key := filepath.Join(rootDirName, filepath.Dir(relPath))

		if info.IsDir() {
			mp[filepath.Join(rootDirName, relPath)] = []string{}
		} else {
			mp[key] = append(mp[key], info.Name())
		}

		return nil
	})

	return mp
}

func isSourceFile(name string) bool {
	ext := filepath.Ext(name)
	switch ext {
	case ".go", ".py":
		return true
	default:
		return false
	}
}

func shouldSkipDir(name string) bool {
	skip := []string{
		".git", ".github", ".idea", ".vscode", "docs", "build", "testdata", "__pycache__", ".devcontainer",
	}
	for _, s := range skip {
		if name == s {
			return true
		}
	}

	return false
}
