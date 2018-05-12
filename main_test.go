package test

import "testing"

// Data provider for Concat test
func concatDataProvider() [][]string {
	arr := [][]string{}
	arr = append(arr, []string{"test-", "test", "test-test"})
	arr = append(arr, []string{"", "", ""})

	return arr
}

func TestConcat(t *testing.T) {
	for _, el := range concatDataProvider() {
		if out := Concat(el[0], el[1]); out != el[2] {
			t.Errorf("Concat(%s, %s) = %s, want %s", el[0], el[1], out, el[2])
		}
	}
}
