package io

import "fmt"

func FormatMsg(msg string) string {
	return fmt.Sprintf("%v: ", msg)
}

func FormatMsgWithDefault(msg, defaultValue string) string {
	return fmt.Sprintf("%v (default = %v): ", msg, defaultValue)
}
