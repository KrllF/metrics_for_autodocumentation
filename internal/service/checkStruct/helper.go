package checkStruct

import (
	"fmt"
	"os"
	"path/filepath"
)

func (s *Service) ListDirByWalk(path string) map[string][]string {
	mp := make(map[string][]string)

	path = filepath.Clean(path)

	rootDirName := filepath.Base(path)

	filepath.Walk(path, func(wPath string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("Ошибка при обходе %q: %v\n", wPath, err)
			return err
		}

		if path == wPath {
			return nil
		}

		if info.IsDir() && info.Name() == ".git" {
			return filepath.SkipDir
		}

		relPath, err := filepath.Rel(path, wPath)
		if err != nil {
			fmt.Printf("filepath.Rel %q: %v\n", wPath, err)
			return err
		}

		key := filepath.Join(rootDirName, relPath)

		if info.IsDir() {
			mp[key] = []string{}
			parentDir := filepath.Dir(wPath)
			relParentDir, err := filepath.Rel(path, parentDir)
			if err != nil {
				fmt.Printf("filepath.Rel %q: %v\n", parentDir, err)
				return err
			}
			parentDirKey := filepath.Join(rootDirName, relParentDir)
			mp[parentDirKey] = append(mp[parentDirKey], info.Name())

			return nil
		}

		dir := filepath.Dir(wPath)
		relDir, err := filepath.Rel(path, dir)
		if err != nil {
			fmt.Printf("filepath.Rel %q: %v\n", dir, err)
			return err
		}
		dirKey := filepath.Join(rootDirName, relDir)
		mp[dirKey] = append(mp[dirKey], info.Name())

		return nil
	})

	return mp
}
