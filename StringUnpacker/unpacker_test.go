package main

import (
	"testing"
)

func TestUnpack(t *testing.T) {
	type Tests struct {
		incoming string
		expected string
	}
	tests := []Tests{
		{"", ""},
		{"a4bc2d5e", "aaaabccddddde"},
		{"a15b11", "aaaaaaaaaaaaaaabbbbbbbbbbb"},
		{"abcd", "abcd"},
		{"a10b20", "aaaaaaaaaabbbbbbbbbbbbbbbbbbbb"},
		{"45", ""},
		{"012", ""},

		{`qwe\415a2`, `qwe444444444444444aa`},
		{`qwe\415\310\\\\\\3`, `qwe4444444444444443333333333\\\\\`},
		{`qwe\4\5`, `qwe45`},
		{`qwe\45`, `qwe44444`},
		{`qwe\\5`, `qwe\\\\\`},
		{`qwe\\\5`, `qwe\5`},
		{`qwe\\2\3\\2`, `qwe\\3\\`},
		{`\\`, `\`},
		{`\\\3\4\\`, `\34\`},
		{`\45q2w3e10`, `44444qqwwweeeeeeeeee`},
		{`\417\310`, `444444444444444443333333333`},
	}
	for _, test := range tests {
		unpackedString := Unpack(test.incoming)
		if test.expected != unpackedString {
			t.Errorf("Unpack(%v) | Actual result: %v Expected result: %v", test.incoming, unpackedString, test.expected)
		}
	}
}
