package config

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/lugosieben/htredirect/internal/util"
)

func ParseSetStrings(setStrings []string) error {
	for _, setString := range setStrings {
		return ParseSetString(setString)
	}
	return nil
}

func ParseSetString(s string) error {
	prefix := "SET "
	cleaned := util.CleanString(s)
	cleanedUpper := strings.ToUpper(cleaned)

	if !strings.HasPrefix(cleanedUpper, prefix) {
		return fmt.Errorf("set string does not start with '%s': %s", prefix, s)
	}

	cleaned = strings.TrimSpace(cleaned[len(prefix):])
	parts := util.ShellSplit(cleaned)
	if parts[1] != "=" {
		return fmt.Errorf("missing or invalid '=' in set string: %s", cleaned)
	}
	key := parts[0]
	value := parts[2]
	switch util.CleanUpperString(key) {
	case "PORT":
		in, err := strconv.Atoi(value)
		if err != nil {
			return err
		}
		Port = in
	default:
		return fmt.Errorf("invalid configuration key in set string: %s", key)
	}
	return nil
}
