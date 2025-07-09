package main

import "testing"


func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "  My name iS Hasan",
			expected: []string{"my", "name", "is", "hasan"},
		},
		{
			input:    " thIs is teSting    ",
			expected: []string{"this", "is", "testing"},
		},
		{
			input:    "     ",
			expected: []string{},
		},
	}

	for _, c := range cases {
		actual := CleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("expected %d elements after cleaning input, got %d", len(c.expected), len(actual))
			continue 
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("CleanInput failed: expected %s, got %s", word, expectedWord)
				
			}
		}
	}
}