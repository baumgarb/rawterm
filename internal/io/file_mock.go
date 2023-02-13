package io

import "io"

type fileMock struct {
	Buffer []byte
	i      int
}

func NewFileMock(b []byte) *fileMock {
	return &fileMock{b, 0}
}

func (f *fileMock) Write(b []byte) (int, error) {
	f.Buffer = append(f.Buffer, b...)
	return len(b), nil
}

func (f *fileMock) WriteString(s string) (int, error) {
	b := []byte(s)
	return f.Write(b)
}

func (f *fileMock) Read(b []byte) (int, error) {
	if f.i >= len(f.Buffer) {
		return 0, io.EOF
	}

	i := 0
	for i < len(b) && f.i < len(f.Buffer) {
		b[i] = f.Buffer[f.i]
		f.i = f.i + 1
		i = i + 1
	}
	return i, nil
}
