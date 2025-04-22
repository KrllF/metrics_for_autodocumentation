package golang

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

// GetMarkDownFunc получить функции из markdown
func (s *Service) GetMarkDownFunc(filePath string) (map[string]struct{}, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("os.Open: %v", err)
	}
	defer file.Close()

	pattern := `^## FunctionDef\s+(\w+)$`
	re := regexp.MustCompile(pattern)

	functionNames := make(map[string]struct{})

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		matches := re.FindStringSubmatch(line)
		if len(matches) > 1 {
			functionNames[matches[1]] = struct{}{}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("scanner.Err: %v", err)
	}

	return functionNames, nil
}
