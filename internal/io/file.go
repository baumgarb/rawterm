package io

import sio "io"

type File interface {
	Write(b []byte) (n int, err error)
	WriteString(s string) (n int, err error)
	sio.Reader
}
