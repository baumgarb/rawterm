package rawterm

import (
	"fmt"
	"strings"

	"github.com/baumgarb/rawterm/internal/io"
)

const (
	arrLeft      = "\033[D"
	arrRight     = "\033[C"
	backspaceKey = "\x7F"
	delKey       = "\033[3~"
)

func arrLeftTimes(n int) string {
	return strings.Repeat(arrLeft, n)
}

func arrRightTimes(n int) string {
	return strings.Repeat(arrRight, n)
}

func backspaceTimes(n int) string {
	return strings.Repeat(backspaceKey, n)
}

func delTimes(n int) string {
	return strings.Repeat(delKey, n)
}

func runReadString(mode InputMode, keysPressed, expectedReturnValue, expectedOutput string) error {
	in := io.NewFileMock([]byte(keysPressed))
	out := io.NewFileMock([]byte{})
	sut := New(in, out)
	sut.Mode = mode
	sut.cdc = newConsoleMock()

	actualReturnValue, err := sut.ReadString("enter something", "123456")
	actualOutput := strings.Trim(string(out.Buffer), "\r\n")
	if err != nil {
		return fmt.Errorf("expected no error, got %v", err)
	}
	if expectedReturnValue != actualReturnValue {
		return fmt.Errorf("expected result '%v', got '%v'", expectedReturnValue, actualReturnValue)
	}
	if expectedOutput != "*" && expectedOutput != actualOutput {
		return fmt.Errorf("expected output '%v' (len %v), got '%v' (len %v)", expectedOutput, len(expectedOutput), actualOutput, len(actualOutput))
	}
	return nil
}
