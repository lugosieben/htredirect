package config

import (
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/lugosieben/htredirect/internal/util"
)

func expandHTRFString(htrfString string) ([]string, error) {
	return expandHTRFStringRecursive(htrfString, map[string]bool{MAINCONFIG: true})
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
				return nil, fmt.Errorf("tried to READ already read file: %s", filePath)
			}
			readFiles[filePath] = true
			fmt.Printf("Reading inner config: %s\n", filePath)
			innerDat, err := os.ReadFile(filePath)
			if err != nil {
				return nil, err
			}
			innerStatements, err := expandHTRFStringRecursive(string(innerDat), readFiles)
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

func sortExpandedHTRFStrings(htrfStrings []string) ([]string, []string, error) {
	var setStatements []string
	var redirectStatements []string
	for _, htrfString := range htrfStrings {
		if strings.HasPrefix(util.CleanUpperString(htrfString), "SET") {
			setStatements = append(setStatements, htrfString)
		} else if strings.HasPrefix(util.CleanUpperString(htrfString), "REDIRECT") {
			redirectStatements = append(redirectStatements, htrfString)
		} else {
			return nil, nil, fmt.Errorf("invalid statement prefix for statement: %s", htrfString)
		}
	}
	return setStatements, redirectStatements, nil
}
