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
	allowedAlphabet := []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	exp := regtgen.Init("###-###", allowedAlphabet)
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

Comparing [nilium/regen](https://github.com/nilium/regen) Vs. [zyxan/regtgen](https://github.com/zyxan/regtgen)

```
❯ go test -benchmem -bench=.
goos: darwin
goarch: amd64
pkg: github.com/zyxan/benchmarks
BenchmarkGoregen-4    	 1000000	      1327 ns/op	     792 B/op	      15 allocs/op
BenchmarkZyxanGen-4   	 5000000	       260 ns/op	       8 B/op	       1 allocs/op
PASS
ok  	github.com/zyxan/benchmarks	2.945s
```

## State -  Work In Progress

### TODO:

- [ ] compute possible combinations with the given alphabet and pattern
- [ ] collision check
- [X] return consumable values
