package reorderlogs

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestReorderLogFiles(t *testing.T) {
	// arrange
	logs := []string{
		"dig1 8 1 5 1",
		"let1 art can",
		"dig2 3 6",
		"let2 own kit dig",
		"let3 art zero"}

	expectedResult := []string{
		"let1 art can",
		"let3 art zero",
		"let2 own kit dig",
		"dig1 8 1 5 1",
		"dig2 3 6"}

	// act
	result := ReorderLogFiles(logs)

	// asset
	if !cmp.Equal(expectedResult, result) {
		t.Errorf("Expected %v. Actual %v", expectedResult, result)
	}
}
