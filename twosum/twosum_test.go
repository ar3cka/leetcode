package twosum

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestTwoSum(t *testing.T) {

	// arrange
	nums := []int{2, 7, 11, 15}
	target := 9
	expectedResult := []int{0, 1}

	// act
	result := TwoSum(nums, target)

	// asset
	if !cmp.Equal(expectedResult, result) {
		t.Errorf("Expected %v. Actual %v", expectedResult, result)
	}
}
