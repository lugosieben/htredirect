package config

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

var Port int
var WebPort int
var Entries []*Entry

func cleanString(s string) string {
	if s == "" {
		return s
	}

	s = strings.ReplaceAll(s, "\r", "")
	parts := strings.Fields(s)
	return strings.Join(parts, " ")
}

func cleanUpperString(s string) string {
	return strings.ToUpper(cleanString(s))
}

func removeCommentLines(s string) string {
	var out []string
	for _, line := range strings.Split(s, "\n") {
		var outLine []rune
		dashBefore := false
		for _, r := range line {
			if r == '-' {
				if dashBefore {
					outLine = outLine[:len(outLine)-1]
					break
				}
				dashBefore = true
			}
			outLine = append(outLine, r)
		}
		out = append(out, string(outLine))
	}
	return strings.Join(out, "\n")
}

func expandHTRFString(htrfString string) ([]string, error) {
	return expandHTRFStringRecursive(htrfString, make(map[string]bool))
}

func expandHTRFStringRecursive(htrfString string, readFiles map[string]bool) ([]string, error) {
	var outStatements []string
	inStatements := strings.Split(cleanString(removeCommentLines(htrfString)), ";")

	for _, statement := range inStatements {
		if strings.HasPrefix(statement, "READ") {
			filePath := cleanString(statement[len("READ"):])
			if filePath == "" {
				return nil, fmt.Errorf("missing file path in READ statement: %s", statement)
			}
			if readFiles[filePath] {
				return nil, fmt.Errorf("circular reference detected in READ statements: %s", filePath)
			}
			readFiles[filePath] = true
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

func Load() {
	fmt.Println("Loading configuration")

	rulesDat, err := os.ReadFile(MAINCONFIG)
	if err != nil {
		panic(err)
	}

	htrfStatements, err := expandHTRFString(string(rulesDat))
	if err != nil {
		panic(err)
	}

	entries, err := ParseEntryStrings(htrfStatements)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Entries found: %d\n", len(*entries))

	Port = 80
	WebPort = 8080
	Entries = *entries
}
