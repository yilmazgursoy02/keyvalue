package main

import (
	"fmt"
	"regexp"
)

type KeyValueMatcher struct {
	pattern *regexp.Regexp
}

func NewKeyValueMatcher(key string) (*KeyValueMatcher, error) {
	pattern := fmt.Sprintf(`#\s?%s:\s?(.*)`, key)
	matcher, err := regexp.Compile(pattern)
	if err != nil {
		return nil, err
	}
	return &KeyValueMatcher{pattern: matcher}, nil
}

func (m *KeyValueMatcher) Match(text string) string {
	results := m.pattern.FindStringSubmatch(text)
	if len(results) == 2 {
		return results[1]
	}
	return ""
}
