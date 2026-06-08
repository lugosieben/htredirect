package config

import (
	"fmt"
	"os"
)

var Port = 80
var WebPort = 8080
var Redirects []Entry

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
	setStatements, redirectStatements, err := sortExpandedHTRFStrings(htrfStatements)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Statements found: %d (%d config, %d redirects)\n", len(htrfStatements), len(setStatements), len(redirectStatements))

	fmt.Printf("Parsing general configuration\n")
	err = ParseSetStrings(setStatements)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Parsing redirects\n")
	redirects, err := ParseRedirectStrings(redirectStatements)
	if err != nil {
		panic(err)
	}
	Redirects = redirects
}
