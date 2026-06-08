package config

import (
	"fmt"
	"os"
)

var Port int
var WebPort int
var Entries []Entry

func Load() {
	fmt.Println("Loading configuration")

	fmt.Printf("Reading config: %s\n", MAINCONFIG)
	rulesDat, err := os.ReadFile(MAINCONFIG)
	if err != nil {
		panic(err)
	}

	htrfStatements, err := expandHTRFString(string(rulesDat))
	if err != nil {
		panic(err)
	}
	fmt.Printf("Statements found: %d\n", len(htrfStatements))

	entries, err := ParseEntryStrings(htrfStatements)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Parsed %d entries\n", len(entries))

	Port = 80
	WebPort = 8080
	Entries = entries
}
