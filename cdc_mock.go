package rawterm

import (
	cdc "github.com/containerd/console"
)

type consoleMock struct {
	File cdc.File
}

func newConsoleMock() cdc.Console {
	return &consoleMock{}
}

func (c *consoleMock) Resize(cdc.WinSize) error     { return nil }
func (c *consoleMock) ResizeFrom(cdc.Console) error { return nil }
func (c *consoleMock) SetRaw() error                { return nil }
func (c *consoleMock) DisableEcho() error           { return nil }
func (c *consoleMock) Reset() error                 { return nil }
func (c *consoleMock) Close() error                 { return nil }
func (c *consoleMock) Fd() uintptr                  { return 0 }
func (c *consoleMock) Name() string                 { return "" }
func (c *consoleMock) Read(b []byte) (int, error)   { return 0, nil }
func (c *consoleMock) Write(b []byte) (int, error)  { return 0, nil }
func (c *consoleMock) Size() (cdc.WinSize, error)   { return cdc.WinSize{}, nil }
