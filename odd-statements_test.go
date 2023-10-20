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
		if strings.Contains(strings.ToLower(s), strings.ToLower(substring)) {
			return true
		}
	}
	return false
}

func TestRandomStatement(t *testing.T) {
	want := "This is a random statement"
	if got := RandomStatement(); got != want {
		t.Errorf("RandomStatement() = %q, want = %q", got, want)
	}

}

func TestHelloInGerman(t *testing.T) {
	want := "Guten Tag!"
	if got := HelloInGerman(); got != want {
		t.Errorf("HelloInGerman() = %q, want = %q", got, want)
	}

}
