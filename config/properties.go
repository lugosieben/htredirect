package config

import "os"

const VERSION = "0.2.0"

var TIMEFORMAT = "2006-01-02 15:04:05"
var MAINCONFIG = "config.htredirect"

func replacePropertiesWithEnv() {
	timeFormat, found := os.LookupEnv("HTREDIRECT_TIMEFORMAT")
	if found {
		TIMEFORMAT = timeFormat
	}

	mainConfig, found := os.LookupEnv("HTREDIRECT_MAINCONFIG")
	if found {
		MAINCONFIG = mainConfig
	}
}
