package config

import (
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/lugosieben/htredirect/internal/util"
)

func expandHTRFString(htrfString string) ([]string, error) {
	return expandHTRFStringRecursive(htrfString, make(map[string]bool))
}

func expandHTRFStringRecursive(htrfString string, readFiles map[string]bool) ([]string, error) {
	var outStatements []string
	inStatements := util.CleanSplit(util.CleanString(util.RemoveCommentLines(htrfString)), ";")

	for _, statement := range inStatements {
		if strings.HasPrefix(util.CleanUpperString(statement), "READ") {
			filePath := util.ShellSplit(statement)[1]
			if filePath == "" {
				return nil, fmt.Errorf("missing file path in READ statement: %s", statement)
			}
			if readFiles[filePath] {
				return nil, fmt.Errorf("circular reference detected in READ statements: %s", filePath)
			}
			readFiles[filePath] = true
			fmt.Printf("Reading inner config: %s\n", filePath)
			innerDat, err := os.ReadFile(filePath)
			if err != nil {
				return nil, err
			}
			innerStatements, err := expandHTRFString(string(innerDat))
			if err != nil {
				return nil, err
			}
			outStatements = slices.Concat(outStatements, innerStatements)
			continue
		}
		outStatements = append(outStatements, statement)
	}

	return outStatements, nil
}
