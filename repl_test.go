package main

import (
	"testing"
)

type cases []struct {
	input    string
	expected []string
}

func TestCleanInput(t *testing.T) {
	tests := cases{{input: "  Hello  World  ",
		expected: []string{"hello", "world"}},
		{input: "I'm very happy! ! !",
			expected: []string{"i'm", "very", "happy!", "!", "!"}},
		{input: "   Multiple   spaces   between words   ",
			expected: []string{"multiple", "spaces", "between", "words"}},
	}

	for _, c := range tests {
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Errorf("Error: %v produced %v", c.input, actual)
			continue
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Error: got %v but actual %v", word, expectedWord)
				t.Fail()
			}
		}
	}

}
