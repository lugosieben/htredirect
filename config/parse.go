package config

import (
	"fmt"
	"strings"
)

func cleanString(s string) string {
	return strings.TrimSpace(s)
}

func cleanUpperString(s string) string {
	return strings.ToUpper(cleanString(s))
}

func ParseEntryStrings(entries []string) (*[]*Entry, error) {
	var parsedEntries []*Entry
	for _, entryString := range entries {
		if cleanString(entryString) == "" {
			continue
		}
		entry, err := ParseEntry(entryString)
		if err != nil {
			return nil, fmt.Errorf("error parsing entry: %s", err)
		}
		parsedEntries = append(parsedEntries, entry)
	}

	return &parsedEntries, nil
}

func ParseEntriesString(entriesString string) (*[]*Entry, error) {
	entries := strings.Split(entriesString, ";")
	return ParseEntryStrings(entries)
}

func ParseEntry(entryString string) (*Entry, error) {
	prefix := "REDIRECT WHERE"

	if !strings.HasPrefix(cleanUpperString(entryString), prefix) {
		return nil, fmt.Errorf("entry does not start with '%s': %s", prefix, entryString)
	}
	mainEntry := strings.Trim(entryString[len(prefix):], " \n")
	parts := strings.Split(mainEntry, "TO")
	rulesString := strings.Trim(parts[0], " \n")
	ruleStrings := strings.Split(rulesString, ",")
	rules := make([]*Rule, len(ruleStrings))

	for i, ruleString := range ruleStrings {
		ruleParts := strings.Split(cleanString(ruleString), " ")
		field, err := ParseRuleField(ruleParts[0])
		if err != nil {
			return nil, err
		}
		comparator, err := ParseRuleComparator(ruleParts[len(ruleParts)-2])
		if err != nil {
			return nil, err
		}
		comparatorMods, err := ParseRuleComparatorMods(ruleParts[1 : len(ruleParts)-2])
		if err != nil {
			return nil, err
		}

		rules[i] = &Rule{
			Field:      field,
			Comparator: comparator,
			Mods:       comparatorMods,
			Value:      ruleParts[len(ruleParts)-1],
		}
	}

	redirectString := strings.Trim(parts[1], " \n")
	redirectParts := strings.Split(redirectString, " ")
	target := redirectParts[0]
	method, err := ParseMethod(redirectParts[1])
	if err != nil {
		return nil, err
	}

	return &Entry{
		Rules:  rules,
		Target: target,
		Method: method,
	}, nil
}

func ParseRuleField(s string) (RuleField, error) {
	switch cleanUpperString(s) {
	case "HOST":
		return RuleFieldHost, nil
	case "PATH":
		return RuleFieldPath, nil
	default:
		return 0, fmt.Errorf("unknown rule field: %s", s)
	}
}

func ParseRuleComparator(s string) (RuleComparator, error) {
	switch cleanUpperString(s) {
	case "EQUALS":
		return RuleComparatorEqual, nil
	case "MATCHES":
		return RuleComparatorRegEx, nil
	case "PREFIX":
		return RuleComparatorPrefix, nil
	case "SUFFIX":
		return RuleComparatorSuffix, nil
	default:
		return 0, fmt.Errorf("unknown rule comparator: %s", s)
	}
}

func ParseRuleComparatorMod(s string) (RuleMod, error) {
	switch cleanUpperString(s) {
	case "NOT":
		return RuleModNot, nil
	case "LOWER":
		return RuleModLower, nil
	default:
		return 0, fmt.Errorf("unknown rule comparator mod: %s", s)
	}
}

func ParseRuleComparatorMods(strings []string) ([]RuleMod, error) {
	var mods []RuleMod
	for _, s := range strings {
		mod, err := ParseRuleComparatorMod(s)
		if err != nil {
			return nil, err
		}
		mods = append(mods, mod)
	}
	return mods, nil
}

func ParseMethod(s string) (Method, error) {
	switch cleanUpperString(s) {
	case "PERMANENT":
		return MethodPermanent, nil
	case "TEMPORARY":
		return MethodTemporary, nil
	default:
		return 0, fmt.Errorf("unknown method: %s", s)
	}
}
