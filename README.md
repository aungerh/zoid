## Regex Text Generator

Generate random strings from expressions such as: "###-###".

---

```
package main

import (
	"fmt"
	"log"

	"github.com/zyxan/regtgen"
)

func main() {
	allowedAlphabet := []byte("abcdefghijklmnopqrstuvwxyz")
	exp := regtgen.InitGenerator("###-###", allowedAlphabet)
	single, err := exp.Generate() // mjg-zjt
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(single)

	many, err := exp.GenerateMany(7)
	if err != nil {
		log.Panic(err)
	}
	for i := 0; i < len(many); i++ {
		fmt.Println(many[i])
	}
}

```

## Benchmarks

```
goos: darwin
goarch: amd64
pkg: github.com/zyxan/regtgen
BenchmarkGeneration-4   	   30000	     58058 ns/op	       8 B/op	       1 allocs/op
PASS
ok  	github.com/zyxan/regtgen	2.336s
Success: Benchmarks passed.
```

## State -  Work In Progress

### TODO:

- [ ] compute possible combinations with the given alphabet and pattern
- [ ] collision check
- [X] return consumable values