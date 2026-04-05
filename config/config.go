package config

import (
	"fmt"
	"os"
)

var Port int
var Entries []*Entry

func Load() {
	fmt.Println("Loading configuration")

	path := "htredirect.yml"
	dat, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	parsed, err := ParseYAML(dat)
	if err != nil {
		panic(err)
	}

	Port = parsed.Port
	Entries = parsed.Entries
}
