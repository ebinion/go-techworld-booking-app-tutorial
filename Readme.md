# Using Go to build a really basic command line tool aka (Go for JS/TS folks)

Following along with ["TechWorld with Nana" tutorial](https://www.youtube.com/watch?v=yyUHQIec83I&t=241sg)

## How to run

1. Check Go Version in go.mod
1. Download & install Go
1. Download or git clone this code
1. `cd` into directory
1. run `go run .`
1. Have fun

## Notes

## Tabs not spaces

Go formatter uses tabs by defaults. Let's no fight the formatter.

To run it from the command line, use:

```
gofmt -w .
```

### Variables in Go

- `const` for immutable variables
- `var` for mutable variables
  I’m still not sure sure why JS decided to use 3 types when `var` and let mean the same thing

### Data Structures in Go

- `struct` is for flexible key/value pairs
- `map` all values must be the same type
- `arrays` are still arrays

To create `map` or `struct` it’s advisable to use the built-in `make` function to create your initial object.

```
var bookings = make([]UserData, 0)
type UserData struct {
	firstName string
	lastName string
	email string
	ticketsCount uint
}
```

### Type declarations and assignment

Declarations are not really necessary to predefine types, unless you need a struct.

Most assignment happens like this

```
var thing = 50
// or
var anotherThing int
```

### Think “I’m making a module with packages”

All code is organized into packages.

```
package main
```

Import sub packages by referencing the module name plus package name

```
import (
	"go-booking-app/helper"
)
```
