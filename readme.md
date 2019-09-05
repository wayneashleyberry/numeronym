### numeronym

> “A **numeronym** is a number-based word. Most commonly, a numeronym is a word where a number is used to form an abbreviation (albeit not an acronym or an initialism).” — https://en.wikipedia.org/wiki/Numeronym

#### command

```sh
$ go install github.com/wayneashleyberry/numeronym/cmd/numeronym
$ numeronym "Andreessen Horowitz"
a16z
```

#### package

```go
package main

import (
	"fmt"

	"github.com/wayneashleyberry/numeronym"
)

func main() {
    s := "Andreessen Horowitz"
    fmt.Println(numeronym.FromString(s)) // a16z
}
```
