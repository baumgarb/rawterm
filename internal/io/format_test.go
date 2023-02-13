package io

import "testing"

func TestFormatMsg(t *testing.T) {
	type test struct {
		input, expect string
	}

	tt := []test{
		{"enter something", "enter something: "},
		{"Enter something", "Enter something: "},
		{"Enter : something", "Enter : something: "},
		{"  Enter : something  ", "  Enter : something  : "},
	}

	for i, td := range tt {
		result := FormatMsg(td.input)
		if td.expect != result {
			t.Errorf("%v: expected '%v', got '%v'", i, td.expect, result)
		}
	}
}

func TestFormatMsgWithDefault(t *testing.T) {
	type test struct {
		msg, defaultValue, expect string
	}

	tt := []test{
		{"enter something", "abc", "enter something (default = abc): "},
		{"Enter something", "abc", "Enter something (default = abc): "},
		{"Enter : something", "abc", "Enter : something (default = abc): "},
		{"  Enter : something  ", "abc", "  Enter : something   (default = abc): "},
	}

	for i, td := range tt {
		result := FormatMsgWithDefault(td.msg, td.defaultValue)
		if td.expect != result {
			t.Errorf("%v: expected '%v', got '%v'", i, td.expect, result)
		}
	}
}
