package reorderlogs

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestReorderLogFiles(t *testing.T) {
	// arrange
	testCases := []struct {
		logs, expectedResults []string
	}{
		{
			[]string{"dig1 8 1 5 1", "let1 art can", "dig2 3 6", "let2 own kit dig", "let3 art zero"},
			[]string{"let1 art can", "let3 art zero", "let2 own kit dig", "dig1 8 1 5 1", "dig2 3 6"},
		},
		{
			[]string{"8 fj dzz k", "5r 446 6 3", "63 gu psub", "5 ba iedrz", "6s 87979 5", "3r 587 01", "jc 3480612", "bb wsrd kp", "b aq cojj", "r5 6316 71"},
			[]string{"b aq cojj", "5 ba iedrz", "8 fj dzz k", "63 gu psub", "bb wsrd kp", "5r 446 6 3", "6s 87979 5", "3r 587 01", "jc 3480612", "r5 6316 71"},
		},
	}

	for _, testCase := range testCases {

		// act
		actualResult := ReorderLogFiles(testCase.logs)

		// asset
		if !cmp.Equal(testCase.expectedResults, actualResult) {
			t.Errorf("Expected %v. Actual %v", testCase.expectedResults, actualResult)
		}
	}
}
