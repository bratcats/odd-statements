package odd_statements

import (
	"strings"
	"testing"
)

func TestCatRandomStatement(t *testing.T) {
	// creat a slice of strings
	want := []string{"cheetah", "cat", "kitten"}
	if got := CatRandomStatement(); !containsSubstring(got, want) {
		t.Errorf("CatRandomStatement() = %q does not contain any %q occurence", got, want)
	}
}

func containsSubstring(s string, substrings []string) bool {
	for _, substring := range substrings {
		if strings.Contains(s, substring) {
			return true
		}
	}
	return false
}
