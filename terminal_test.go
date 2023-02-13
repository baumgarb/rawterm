package rawterm

import (
	"testing"
)

func TestReadString_OvertypeMode_IgnoresKeysPressedAtTheEnd(t *testing.T) {
	err := runReadString(Overtype, "abcdef\n", "123456", "enter something: 123456")
	if err != nil {
		t.Error(err)
	}
}

func TestReadString_OvertypeMode_IgnoresBackspace(t *testing.T) {
	err := runReadString(Overtype, backspaceTimes(3)+"\n", "123456", "enter something: 123456")
	if err != nil {
		t.Error(err)
	}
}

func TestReadString_OvertypeMode_HandlesCursorLeftCorrectly(t *testing.T) {
	type test struct {
		input, expectedResult, expectedOutput string
	}

	tt := []test{
		{arrLeft + "abc\n", "12345a", "enter something: 123456" + arrLeft + "a"},
		{arrLeftTimes(3) + "abc\n", "123abc", "enter something: 123456" + arrLeftTimes(3) + "abc"},
		{arrLeftTimes(10) + "abc\n", "abc456", "enter something: 123456" + arrLeftTimes(6) + "abc"},
	}

	for i, td := range tt {
		err := runReadString(Overtype, td.input, td.expectedResult, td.expectedOutput)
		if err != nil {
			t.Errorf("#%v: %v", i, err)
		}
	}
}

func TestReadString_OvertypeMode_IgnoresAllCursorRightAtTheEnd(t *testing.T) {
	type test struct {
		input, expectedResult, expectedOutput string
	}

	tt := []test{
		{arrRight + "abc\n", "123456", "enter something: 123456"},
		{arrRightTimes(3) + "abc\n", "123456", "enter something: 123456"},
		{arrRightTimes(10) + "abc\n", "123456", "enter something: 123456"},
	}

	for i, td := range tt {
		err := runReadString(Overtype, td.input, td.expectedResult, td.expectedOutput)
		if err != nil {
			t.Errorf("#%v: %v", i, err)
		}
	}
}

func TestReadString_OvertypeMode_HandlesCursorLeftAndRightCorrectly(t *testing.T) {
	type test struct {
		input, expectedResult, expectedOutput string
	}

	tt := []test{
		{arrLeft + arrRight + "abc\n", "123456", "enter something: 123456" + arrLeft + arrRight},
		{arrLeftTimes(3) + arrRightTimes(1) + "abc\n", "1234ab", "enter something: 123456" + arrLeftTimes(3) + arrRightTimes(1) + "ab"},
		{arrLeftTimes(3) + arrRightTimes(1) + arrLeftTimes(2) + "abc\n", "12abc6", "enter something: 123456" + arrLeftTimes(3) + arrRightTimes(1) + arrLeftTimes(2) + "abc"},
		{
			arrLeftTimes(3) + arrRightTimes(1) + arrLeftTimes(2) + "a" + arrRightTimes(2) + "b" + arrLeftTimes(6) + "c\n",
			"c2a45b",
			"enter something: 123456" + arrLeftTimes(3) + arrRightTimes(1) + arrLeftTimes(2) + "a" + arrRightTimes(2) + "b" + arrLeftTimes(6) + "c",
		},
	}

	for i, td := range tt {
		err := runReadString(Overtype, td.input, td.expectedResult, td.expectedOutput)
		if err != nil {
			t.Errorf("#%v: %v", i, err)
		}
	}
}

func TestReadString_InsertMode_HandlesCursorLeftCorrectly(t *testing.T) {
	type test struct {
		input, expectedResult, expectedOutput string
	}

	tt := []test{
		{arrLeft + "abc\n", "12345abc6", "enter something: 123456" + arrLeft + "abc"},
		{arrLeftTimes(3) + "abc\n", "123abc456", "enter something: 123456" + arrLeftTimes(3) + "abc"},
		{arrLeftTimes(10) + "abc\n", "abc123456", "enter something: 123456" + arrLeftTimes(6) + "abc"},
	}

	for i, td := range tt {
		err := runReadString(Insert, td.input, td.expectedResult, td.expectedOutput)
		if err != nil {
			t.Errorf("#%v: %v", i, err)
		}
	}
}

func TestReadString_InsertMode_HandlesCursorLeftAndRightCorrectly(t *testing.T) {
	type test struct {
		input, expectedResult, expectedOutput string
	}

	tt := []test{
		{arrLeft + arrRight + "abc\n", "123456abc", "enter something: 123456" + arrLeft + arrRight + "abc"},
		{arrLeftTimes(3) + arrRightTimes(1) + "abc\n", "1234abc56", "enter something: 123456" + arrLeftTimes(3) + arrRightTimes(1) + "abc"},
		{arrLeftTimes(3) + arrRightTimes(1) + arrLeftTimes(2) + "abc\n", "12abc3456", "enter something: 123456" + arrLeftTimes(3) + arrRightTimes(1) + arrLeftTimes(2) + "abc"},
		{
			arrLeftTimes(3) + arrRightTimes(1) + arrLeftTimes(2) + "a" + arrRightTimes(2) + "b" + arrLeftTimes(6) + "c\n",
			"c12a34b56",
			"enter something: 123456" + arrLeftTimes(3) + arrRightTimes(1) + arrLeftTimes(2) + "a" + arrRightTimes(2) + "b" + arrLeftTimes(6) + "c",
		},
	}

	for i, td := range tt {
		err := runReadString(Insert, td.input, td.expectedResult, td.expectedOutput)
		if err != nil {
			t.Errorf("#%v: %v", i, err)
		}
	}
}

func TestReadString_InsertMode_IgnoresAllCursorRightAtTheEnd(t *testing.T) {
	type test struct {
		input, expectedResult, expectedOutput string
	}

	tt := []test{
		{arrRight + "abc\n", "123456abc", "enter something: 123456abc"},
		{arrRightTimes(3) + "abc\n", "123456abc", "enter something: 123456abc"},
		{arrRightTimes(10) + "abc\n", "123456abc", "enter something: 123456abc"},
	}

	for i, td := range tt {
		err := runReadString(Insert, td.input, td.expectedResult, td.expectedOutput)
		if err != nil {
			t.Errorf("#%v: %v", i, err)
		}
	}
}

func TestReadString_InsertMode_HandlesBackspaceCorrectly(t *testing.T) {
	type test struct {
		input, expectedResult, expectedOutput string
	}

	tt := []test{
		{backspaceKey + "\n", "12345", "enter something: 123456\033[D\033[K"},
		{backspaceTimes(3) + "\n", "123", "enter something: 123456\033[D\033[K\033[D\033[K\033[D\033[K"},
		{backspaceTimes(3) + "abc\n", "123abc", "enter something: 123456\033[D\033[K\033[D\033[K\033[D\033[Kabc"},
		{backspaceTimes(10) + "\n", "", "*"},
		{backspaceKey + arrLeftTimes(2) + backspaceKey + "abc\n", "12abc45", "*"},
		{backspaceKey + arrLeftTimes(2) + backspaceKey + "abc\n", "12abc45", "*"},
		{backspaceKey + arrLeftTimes(10) + backspaceKey + "abc\n", "abc12345", "*"},
		{backspaceKey + arrLeftTimes(10) + arrRightTimes(3) + backspaceKey + "abc\n", "12abc45", "*"},
		{backspaceKey + arrLeftTimes(10) + arrRightTimes(10) + backspaceKey + "abc\n", "1234abc", "*"},
	}

	for i, td := range tt {
		err := runReadString(Insert, td.input, td.expectedResult, td.expectedOutput)
		if err != nil {
			t.Errorf("#%v: %v", i, err)
		}
	}
}

func TestReadString_InsertMode_HandlesDeleteCorrectly(t *testing.T) {
	type test struct {
		input, expectedResult, expectedOutput string
	}

	tt := []test{
		{delKey + "\n", "123456", "*"},
		{arrLeft + delKey + "\n", "12345", "*"},
		{arrLeftTimes(3) + delKey + "\n", "12356", "*"},
		{arrLeftTimes(10) + delTimes(10) + "\n", "", "*"},
		{arrLeftTimes(10) + delTimes(10) + "\n", "", "*"},
		{arrLeftTimes(10) + delTimes(10) + "abc\n", "abc", "*"},
		{arrLeftTimes(10) + delKey + arrRight + delKey + "abc\n", "2abc456", "*"},
	}

	for i, td := range tt {
		err := runReadString(Insert, td.input, td.expectedResult, td.expectedOutput)
		if err != nil {
			t.Errorf("#%v: %v", i, err)
		}
	}
}
