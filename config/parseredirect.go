package config

import (
	"fmt"
	"strings"

	"github.com/lugosieben/htredirect/internal/util"
)

func ParseRedirectStrings(redirectStrings []string) ([]Entry, error) {
	var parsedEntries []Entry
	for _, redirectString := range redirectStrings {
		if util.CleanString(redirectString) == "" {
			continue
		}
		entry, err := ParseRedirectString(redirectString)
		if err != nil {
			return nil, fmt.Errorf("error parsing entry: %s", err)
		}
		parsedEntries = append(parsedEntries, *entry)
	}

	return parsedEntries, nil
}

func ParseRedirectString(redirectString string) (*Entry, error) {
	prefix := "REDIRECT WHERE"
	cleaned := util.CleanString(redirectString)
	cleanedUpper := strings.ToUpper(cleaned)

	if !strings.HasPrefix(cleanedUpper, prefix) {
		return nil, fmt.Errorf("entry does not start with '%s': %s", prefix, redirectString)
	}

	main := strings.TrimSpace(cleaned[len(prefix):])
	parts := strings.SplitN(main, "TO", 2)
	if len(parts) < 2 {
		return nil, fmt.Errorf("missing TO in entry: %s", redirectString)
	}
	rulesString := strings.TrimSpace(parts[0])
	ruleStrings := util.CleanSplit(rulesString, ",")
	rules := make([]Rule, len(ruleStrings))

	for i, ruleString := range ruleStrings {
		ruleParts := util.ShellSplit(util.CleanString(ruleString))
		if len(ruleParts) < 3 {
			return nil, fmt.Errorf("invalid rule format: %s", ruleString)
		}
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

		rules[i] = Rule{
			Field:      field,
			Comparator: comparator,
			Mods:       comparatorMods,
			Value:      ruleParts[len(ruleParts)-1],
		}
	}

	redirectionString := strings.TrimSpace(parts[1])
	redirectionParts := strings.Fields(redirectionString)
	if len(redirectionParts) < 2 {
		return nil, fmt.Errorf("invalid redirect target/method: %s", redirectionString)
	}
	target := redirectionParts[0]
	method, err := ParseMethod(redirectionParts[1])
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
	switch util.CleanUpperString(s) {
	case "HOST":
		return RuleFieldHost, nil
	case "PATH":
		return RuleFieldPath, nil
	default:
		return 0, fmt.Errorf("unknown rule field: %s", s)
	}
}

func ParseRuleComparator(s string) (RuleComparator, error) {
	switch util.CleanUpperString(s) {
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
	switch util.CleanUpperString(s) {
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
	switch util.CleanUpperString(s) {
	case "PERMANENT":
		return MethodPermanent, nil
	case "TEMPORARY":
		return MethodTemporary, nil
	default:
		return 0, fmt.Errorf("unknown method: %s", s)
	}
}
