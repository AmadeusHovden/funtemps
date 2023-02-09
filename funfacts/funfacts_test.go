package funfacts

import (
	"reflect"
	"testing"
)

func TestGetFunFacts(t *testing.T) {
	funFacts := FunFacts{

		Sun: []string{
			"Visste du at temperaturen i solens kjerne er 15 000 000℃",
			"Visste du at temperaturen på ytre lag av solen er 5778K",
		},
		Terra: []string{
			"Visste du at høyeste målt temperatur på jordens overflate er 56℃",
			"Visste du at laveste temperatur på jordens overflate er -89.4℃",
			"Visste du at temperaturen i jordens indre kjerne er 9392K ",
		},
		Luna: []string{
			"Visste du at temperaturen på månens overflate om natten -183℃",
			"Visste du at temperaturen på månens overflate om dagen 106℃",
		},
	}

	type test struct {
		input string
		want  []string
	}

	tests := []test{
		{"sun", funFacts.Sun},
		{"luna", funFacts.Luna},
		{"terra", funFacts.Terra},
		{"unknown", nil},
	}

	for _, tc := range tests {
		got := GetFunFacts(tc.input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}
