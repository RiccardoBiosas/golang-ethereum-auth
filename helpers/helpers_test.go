package helpers

import (
	"testing"
)

func TestRandomStringGeneratorDuplicates(t *testing.T) {
	//naive test to check whether the "random" string generator produces duplicates
	var output []string
	for n := 0; n < 10000; n++ {
		str := GenerateRandomString(20)
		output = append(output, str)
	}
	for i, entry := range output {
		for _, check := range output[i+1:] {
			if entry == check {
				t.Errorf("Found that %v is equal to %v; expected no duplicates;", entry, check)
			}
		}
	}
}
