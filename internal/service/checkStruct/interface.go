package checkStruct

import "context"

type DirectoryLister interface {
	ListDirByWalk(path string) (map[string][]string, error)
	ListDirByWalkContext(ctx context.Context, path string) (map[string][]string, error)
}

type StructureComparer interface {
	EqualStruct(sourceFile, mdFile string) (bool, error)
	EqualStructContext(ctx context.Context, sourceFile, mdFile string) (bool, error)
}

type Service struct{}
