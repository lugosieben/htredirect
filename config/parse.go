package config

import (
	"fmt"
	"regexp"

	"gopkg.in/yaml.v3"
)

type rawRule struct {
	Field      string `yaml:"field"`
	Comparator string `yaml:"comparator"`
	Value      string `yaml:"value"`
}

type rawEntry struct {
	Target string    `yaml:"target"`
	Method string    `yaml:"method"`
	Rules  []rawRule `yaml:"rules"`
}

type rawConfig struct {
	Port    int        `yaml:"port"`
	Entries []rawEntry `yaml:"entries"`
}

type ParsedConfig struct {
	Port    int
	Entries []*Entry
}

func ParseRuleField(s string) (RuleField, error) {
	switch s {
	case "host":
		return RuleFieldHost, nil
	case "path":
		return RuleFieldPath, nil
	default:
		return 0, fmt.Errorf("unknown rule field: %s", s)
	}
}

func ParseRuleComparator(s string) (RuleComparator, error) {
	switch s {
	case "equal":
		return RuleComparatorEqual, nil
	case "equal-insensitive":
		return RuleComparatorEqualInsensitive, nil
	case "notequal":
		return RuleComparatorNotEqual, nil
	case "regex":
		return RuleComparatorRegEx, nil
	case "notregex":
		return RuleComparatorNotRegEx, nil
	case "prefix":
		return RuleComparatorPrefix, nil
	case "suffix":
		return RuleComparatorSuffix, nil
	default:
		return 0, fmt.Errorf("unknown rule comparator: %s", s)
	}
}

func ParseMethod(s string) (Method, error) {
	switch s {
	case "permanent":
		return MethodPermanent, nil
	case "temporary":
		return MethodTemporary, nil
	default:
		return 0, fmt.Errorf("unknown method: %s", s)
	}
}

func ParseYAML(data []byte) (*ParsedConfig, error) {
	var rawCfg rawConfig
	if err := yaml.Unmarshal(data, &rawCfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal YAML: %w", err)
	}

	var entries []*Entry
	for i, rawEnt := range rawCfg.Entries {
		method, err := ParseMethod(rawEnt.Method)
		if err != nil {
			return nil, fmt.Errorf("entry %d: %w", i, err)
		}

		var rules []*Rule
		for j, rawRule := range rawEnt.Rules {
			field, err := ParseRuleField(rawRule.Field)
			if err != nil {
				return nil, fmt.Errorf("entry %d, rule %d: %w", i, j, err)
			}

			comparator, err := ParseRuleComparator(rawRule.Comparator)
			if err != nil {
				return nil, fmt.Errorf("entry %d, rule %d: %w", i, j, err)
			}

			if comparator == RuleComparatorRegEx {
				if _, err := regexp.Compile(rawRule.Value); err != nil {
					return nil, fmt.Errorf("entry %d, rule %d: invalid regex: %w", i, j, err)
				}
			}

			rule := &Rule{
				Field:      field,
				Comparator: comparator,
				Value:      rawRule.Value,
			}
			rules = append(rules, rule)
		}

		entry := &Entry{
			Target: rawEnt.Target,
			Method: method,
			Rules:  rules,
		}
		entries = append(entries, entry)
	}

	return &ParsedConfig{
		Port:    rawCfg.Port,
		Entries: entries,
	}, nil
}
