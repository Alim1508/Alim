package main

import (
	"bytes"
	"testing"
)

type testCase struct {
	input    string
	expected string
}

func TestMaskLinks(t *testing.T) {
	testCases := []testCase{
		{"Here's my spammy page http://hehefouls.netHAHAHA see you.", "Here's my spammy page [LINK REMOVED] see you."},
		{"http://example.com is a good site.", "[LINK REMOVED] is a good site."},
		{"Visit our site at http://example.com and enjoy!", "Visit our site at [LINK REMOVED] and enjoy!"},
		{"This text has no links.", "This text has no links."},
		{"MixedCase HTTP://example.com should also work.", "MixedCase [LINK REMOVED] should also work."},
	}

	for _, tc := range testCases {
		result := maskLinks(tc.input)
		if !bytes.Equal(result, []byte(tc.expected)) {
			t.Errorf("maskLinks(%q) = %q; want %q", tc.input, result, tc.expected)
		}
	}
}
