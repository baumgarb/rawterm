# Rawterm

A super tiny Go module which allows to prompt users for command line input while the terminal is in raw mode. This provides some more flexibility where you can provide default values on the commandl ine which can be overwritten depending on the `InputMode` you're in.

## Use cases

TBD

## Usage

Here's a fully working example. Error handling's omitted for brevity.

```go
package main

import (
	"fmt"
	"os"

	"github.com/baumgarb/rawterm"
)

func main() {
	term := rawterm.New(os.Stdin, os.Stdout)

	// Terminal is in 'Overtype' input mode by default. That's useful if you're asking for input with a specific length and in a specific format
	// so that users only need to overwrite very few parts of it.
	//
	// Users can only use left and right arrow keys and overwrite the character at the cursor. Nothing else.
	startTS, _ := term.ReadString("Enter start timestamp", "2023-02-12T15:00:00.000Z")
	endTS, _ := term.ReadString("Enter end timestamp", "2023-12-31T23:59:59.999Z")

	// Change to 'Insert' mode so users can delete characters and enter an answer of variable length
	term.Mode = rawterm.Insert
	name, _ := term.ReadString("Enter your full name", "Jane Doe")

	fmt.Printf("You entered the following information: \n\tStart: %v\n\tEnd: %v\n\tFull name: %v\n", startTS, endTS, name)
}
```
