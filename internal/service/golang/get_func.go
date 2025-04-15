package helper

import (
	"fmt"
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
)

// GetFileFunc получить названия функций из исходного файла
func (s *Service) GetFileFunc(filePath string) (map[string]struct{}, error) {
	fset := token.NewFileSet()

	node, err := parser.ParseFile(fset, filePath, nil, parser.ParseComments)
	if err != nil {
		return nil, fmt.Errorf("parser.ParseFile: %w", err)
	}

	conf := types.Config{Importer: importer.Default()}

	info := &types.Info{
		Defs: make(map[*ast.Ident]types.Object),
	}

	_, err = conf.Check("mypackage", fset, []*ast.File{node}, info)
	if err != nil {
		return nil, fmt.Errorf("conf.Check: %w", err)
	}

	ret := make(map[string]struct{})
	for id, obj := range info.Defs {
		if obj != nil {
			ret[id.Name] = struct{}{}
		}
	}

	return ret, nil
}
