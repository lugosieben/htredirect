package config

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"
)

type RuleField int

const (
	RuleFieldHost RuleField = iota
	RuleFieldPath
)

func (rf RuleField) String() string {
	switch rf {
	case RuleFieldHost:
		return "HOST"
	case RuleFieldPath:
		return "PATH"
	}
	return "UNKNOWN"
}

type RuleComparator int

const (
	RuleComparatorEqual RuleComparator = iota
	RuleComparatorEqualInsensitive
	RuleComparatorNotEqual
	RuleComparatorRegEx
	RuleComparatorNotRegEx
	RuleComparatorPrefix
	RuleComparatorSuffix
)

func (rc RuleComparator) String() string {
	switch rc {
	case RuleComparatorEqual:
		return "EQUALS"
	case RuleComparatorEqualInsensitive:
		return "EQUALS CASE-INSENSITIVE"
	case RuleComparatorNotEqual:
		return "NOT EQUALS"
	case RuleComparatorRegEx:
		return "REGEX"
	case RuleComparatorNotRegEx:
		return "NOT REGEX"
	case RuleComparatorPrefix:
		return "PREFIX"
	case RuleComparatorSuffix:
		return "SUFFIX"
	}
	return "UNKNOWN"
}

type RuleMod int

const (
	RuleModNot RuleMod = iota
	RuleModLower
)

func (rcm RuleMod) String() string {
	switch rcm {
	case RuleModNot:
		return "NOT"
	case RuleModLower:
		return "LOWER"
	}
	return "UNKNOWN"
}

type Rule struct {
	Field      RuleField
	Comparator RuleComparator
	Mods       []RuleMod
	Value      string
}

func (r Rule) String() string {
	field := r.Field.String()
	comparator := r.Comparator.String()
	mods := r.Mods
	modStrings := make([]string, len(mods))
	for idx, mod := range mods {
		modStrings[idx] = mod.String()
	}
	modsString := strings.Join(modStrings, " ")

	return fmt.Sprintf("%s %s %s \"%s\"", field, modsString, comparator, r.Value)
}

func (r Rule) Match(host string, path string) (bool, error) {
	var field string
	switch r.Field {
	case RuleFieldHost:
		field = host
	case RuleFieldPath:
		field = path
	}

	invert := false
	for _, mod := range r.Mods {
		switch mod {
		case RuleModNot:
			invert = !invert
		case RuleModLower:
			field = strings.ToLower(field)
		}
	}

	match, err := r.Compare(field)
	if err != nil {
		return false, err
	}

	if invert {
		return !match, nil
	}
	return match, nil
}

func (r Rule) Compare(field string) (bool, error) {
	switch r.Comparator {
	case RuleComparatorEqual:
		return field == r.Value, nil
	case RuleComparatorEqualInsensitive:
		return strings.EqualFold(field, r.Value), nil
	case RuleComparatorNotEqual:
		return field != r.Value, nil
	case RuleComparatorRegEx:
		return regexp.MatchString(r.Value, field)
	case RuleComparatorNotRegEx:
		matched, err := regexp.MatchString(r.Value, field)
		return !matched, err
	case RuleComparatorPrefix:
		return strings.HasPrefix(r.Value, field), nil
	case RuleComparatorSuffix:
		return strings.HasSuffix(r.Value, field), nil
	}

	return false, nil
}

func (r Rule) MatchRequest(req *http.Request) (bool, error) {
	return r.Match(req.Host, req.URL.Path)
}

type Method int

const (
	MethodPermanent Method = iota
	MethodTemporary
)

func (m Method) String() string {
	switch m {
	case MethodPermanent:
		return "Permanent"
	case MethodTemporary:
		return "Temporary"
	}
	return "Unknown Method"
}

type Entry struct {
	Target string
	Method Method
	Rules  []Rule
}

func (e Entry) Match(host string, path string) (bool, error) {
	for _, rule := range e.Rules {
		match, err := rule.Match(host, path)
		if err != nil {
			return false, err
		}
		if !match {
			return false, nil
		}
	}
	return true, nil
}

func (e Entry) MatchRequest(req *http.Request) (bool, error) {
	return e.Match(req.Host, req.URL.Path)
}
