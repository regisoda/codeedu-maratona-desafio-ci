package main

import "testing"

func TestSumNumbers1(t *testing.T) {
	resultExpected := 10
  	var result = sumNumbers(5,5)
  	if result != 10{
	    t.Errorf("Expected: %d, returns: %d", resultExpected, result)
	}
}

func TestSumNumbers2(t *testing.T){
	testResults := []struct {
		x int
		y int
		n int
	}{
		{1, 2, 3},
		{2, 3, 5},
		{0, 1, 1},
	}

	for _, test := range testResults {
		result  := sumNumbers(test.x, test.y)
		if result != test.n {
			t.Errorf("Sum (%d+%d) incorrect, returns: %d, expected: %d.", test.x, test.y, result, test.n)
		}
	}
}
