package rawterm

import (
	"fmt"
	"log"

	sio "io" // sio = standard io change

	"github.com/baumgarb/rawterm/internal/io"
	cdc "github.com/containerd/console" // cdc = container d console
)

// InputMode represents the mode in which users can enter an input. Default is FixedLength.
type InputMode int

const (
	// Mode in which the final input has to have the same length as the provided default value.
	//
	// This mode ignores Backspace and Delete keys. You can only use arrow keys to navigate within the boundaries of the
	// provided default value and overwrite individual characters. The entered value will therefore have the same
	// length as the provided default value.
	FixedLength = iota

	// Mode in which the user are free to enter whatever they want.
	//
	// This mode accepts Backspace and Delete key presses and treats them as such. You can insert or remove characters at any
	// position. The entered value can therefore have an arbitrary length which differs from the length of the provided
	// default value.
	Flexible
)

type terminal struct {
	in, out io.File
	// Only set cdc if you want to mock it (e.g. for tests).
	cdc  cdc.Console
	Mode InputMode
}

// Instantiates a new raw terminal with InputMode set to FixedLength by default. You can provide specific descriptors via the
// in and out parameters. Note that they have to be valid terminals / consoles, though. If you pass a regular file descriptor
// we'll panic since there's no step further.
//
// Most of the time you'll invoke this function as follows:
//
//  term := rawterm.New(os.Stdin, os.Stdout)
//  input, err := term.ReadString("Enter some six digit number", "123456")
func New(in, out io.File) *terminal {
	return &terminal{
		in:   in,
		out:  out,
		Mode: FixedLength,
	}
}

// ReadString reads a string from the terminal. It prints the msg provided followed by a colon and blank (: ) and the defaultValue,
// see also the following format
//
// <msg>: <defaultValue>
//
// Using arrow keys you can only navigate the cursor left and right within the boundaries of the defaultValue. See InputMode for more
// information on various modes that are supported.
func (c *terminal) ReadString(msg, defaultValue string) (string, error) {
	fmt.Fprint(c.out, io.FormatMsg(msg))
	fmt.Fprint(c.out, defaultValue)

	if c.cdc == nil {
		// Tests will be the only ones setting c.cdc explicitly to mock this thing away
		c.cdc = cdc.Current()
	}
	defer c.cdc.Reset()
	if err := c.cdc.SetRaw(); err != nil {
		return "", fmt.Errorf("enhanced console: error entering raw mode: %w", err)
	}

	cursorX := len(defaultValue)
	maxCursorX := len(defaultValue)
	keepGoing := true
	result := []rune(defaultValue)
	for keepGoing {
		b := readOneByteFrom(c.in)
		switch b {
		case '\r':
			fallthrough
		case '\n':
			c.out.WriteString("\r\n")
			keepGoing = false
		case 127:
			if c.Mode == FixedLength {
				break
			}
			if cursorX <= 0 {
				break
			}

			result = removeCharAt(result, cursorX-1)
			cursorX--
			maxCursorX--
			c.out.WriteString(fmt.Sprintf("\033[D\033[K%v", string(result[cursorX:])))

			moveCursorBack := maxCursorX-cursorX > 0
			if moveCursorBack {
				c.out.WriteString(fmt.Sprintf("\033[%vD", maxCursorX-cursorX))
			}

		case 27:
			b2 := readOneByteFrom(c.in)
			switch b2 {
			case '[':
				b3 := readOneByteFrom(c.in)
				switch b3 {
				case 'C':
					if cursorX >= maxCursorX {
						break
					}
					c.out.WriteString("\033[C")
					cursorX++
				case 'D':
					if cursorX <= 0 {
						break
					}
					c.out.WriteString("\033[D")
					cursorX--
				case '3':
					b4 := readOneByteFrom(c.in)
					switch b4 {
					case '~':
						if c.Mode == FixedLength {
							break
						}
						if cursorX < 0 || cursorX == maxCursorX {
							break
						}

						result = removeCharAt(result, cursorX)
						maxCursorX--
						c.out.WriteString(fmt.Sprintf("\033[K%v", string(result[cursorX:])))

						moveCursorBack := maxCursorX-cursorX > 0
						if moveCursorBack {
							c.out.WriteString(fmt.Sprintf("\033[%vD", maxCursorX-cursorX))
						}
					}
				}
			}
		default:
			if c.Mode == FixedLength && cursorX >= maxCursorX {
				break
			}
			c.out.Write([]byte{b})
			switch c.Mode {
			case FixedLength:
				result[cursorX] = rune(b)
			case Flexible:
				result = insertCharAt(result, cursorX, rune(b))
			}
			cursorX++
		}
	}

	return string(result), nil
}

func readOneByteFrom(r sio.Reader) byte {
	xb := make([]byte, 1)
	_, err := r.Read(xb)
	if err != nil {
		log.Fatalf("enhanced console: error reading input byte: %v", err)
	}
	return xb[0]
}
