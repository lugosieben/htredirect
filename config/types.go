package config

import (
	"net/http"
	"regexp"
	"strings"
)

type RuleField int

const (
	RuleFieldHost RuleField = iota
	RuleFieldPath
)

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

type Rule struct {
	Field      RuleField
	Comparator RuleComparator
	Value      string
}

func (r Rule) Match(host string, path string) (bool, error) {
	var usedField string
	switch r.Field {
	case RuleFieldHost:
		usedField = host
	case RuleFieldPath:
		usedField = path
	}

	switch r.Comparator {
	case RuleComparatorEqual:
		return usedField == r.Value, nil
	case RuleComparatorEqualInsensitive:
		return strings.EqualFold(usedField, r.Value), nil
	case RuleComparatorNotEqual:
		return usedField != r.Value, nil
	case RuleComparatorRegEx:
		return regexp.MatchString(r.Value, usedField)
	case RuleComparatorNotRegEx:
		matched, err := regexp.MatchString(r.Value, usedField)
		return !matched, err
	case RuleComparatorPrefix:
		return strings.HasPrefix(r.Value, usedField), nil
	case RuleComparatorSuffix:
		return strings.HasSuffix(r.Value, usedField), nil
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

func (m *Method) String() string {
	switch *m {
	case MethodPermanent:
		return "MethodPermanent"
	case MethodTemporary:
		return "MethodTemporary"
	}
	return "Unknown Method"
}

type Entry struct {
	Target string
	Method Method
	Rules  []*Rule
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
