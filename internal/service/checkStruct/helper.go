package checkStruct

import (
	"context"
	"errors"
	"os"
	"path/filepath"
	"runtime"
	"sync"
)

type dirEntry struct {
	path  string
	files []string
}

func (s *Service) ListDirByWalk(path string) (map[string][]string, error) {
	return s.ListDirByWalkContext(context.Background(), path)
}

func (s *Service) ListDirByWalkContext(ctx context.Context, path string) (map[string][]string, error) {
	path = filepath.Clean(path)
	rootDirName := filepath.Base(path)

	var (
		wg         sync.WaitGroup
		mu         sync.Mutex
		result     = make(map[string][]string)
		workers    = runtime.NumCPU()
		dirQueue   = make(chan string, workers*2)
		errorChan  = make(chan error, 1)
		ctxCancel  context.CancelFunc
		hasError   bool
	)

	ctx, ctxCancel = context.WithCancel(ctx)
	defer ctxCancel()

	// Start worker pool
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for dir := range dirQueue {
				select {
				case <-ctx.Done():
					return
				default:
					if err := s.processDir(dir, path, rootDirName, &mu, result); err != nil {
						if !hasError {
							hasError = true
							errorChan <- err
							ctxCancel()
						}
						return
					}
				}
			}
		}()
	}

	// Initial directory to process
	dirQueue <- path

	// Wait for completion in separate goroutine
	go func() {
		wg.Wait()
		close(errorChan)
	}()

	select {
	case err := <-errorChan:
		return nil, err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

func (s *Service) processDir(dir, basePath, rootDirName string, mu *sync.Mutex, result map[string][]string) error {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return err
	}

	relPath, err := filepath.Rel(basePath, dir)
	if err != nil {
		return err
	}
	key := filepath.Join(rootDirName, relPath)

	var files []string
	for _, entry := range entries {
		if entry.IsDir() {
			if entry.Name() == ".git" {
				continue
			}
			files = append(files, entry.Name()+string(filepath.Separator))
		} else {
			files = append(files, entry.Name())
		}
	}

	mu.Lock()
	result[key] = files
	mu.Unlock()

	return nil
}
