package p20

import (
	"reflect"
	"testing"
)

func TestIsValid(t *testing.T) {
	var tests = []struct {
		s    string
		want bool
	}{
		{"]", false},
		{"()", true},
		{"()[]{}", true},
		{"(]", false},
		{"([)]", false},
		{"{[]}", true},
	}
	for _, test := range tests {
		if got := isValid(test.s); !reflect.DeepEqual(got, test.want) {
			t.Errorf("isValid(%q) = %v", test.s, got)
		}
	}
}
