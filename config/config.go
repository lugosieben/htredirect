package config

import (
	"fmt"
	"os"
)

var Port int
var WebPort int
var Entries []*Entry

func Load() {
	fmt.Println("Loading configuration")

	rulesPath := "rules.htredirect"
	rulesDat, err := os.ReadFile(rulesPath)
	if err != nil {
		panic(err)
	}

	entries, err := ParseEntriesString(string(rulesDat))
	if err != nil {
		panic(err)
	}
	fmt.Printf("Entries found: %d\n", len(*entries))

	Port = 80
	WebPort = 8080
	Entries = *entries
}
