## Regex Text Generator

Generate random strings from expressions such as: "###-###". The subset of regular expressions this package accepts is reduced but curated in order to be blazing fast.

---

```
package main

import (
	"fmt"

	"github.com/zyxan/regtgen"
)

func main() {
	allowedAlphabet := []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	gen := regtgen.Init("[###]{3}-fixed-[###]{3}", allowedAlphabet)
	single, _ := gen.Generate() // tpF-fixed-bWV
	fmt.Println(single)
}
```

## Benchmarks

Comparing [nilium/regen](https://github.com/nilium/regen) vs [lucasjones/reggen](https://github.com/lucasjones/reggen) vs [zyxan/regtgen](https://github.com/zyxan/regtgen)

```
❯ go test -benchmem -bench=.
goos: darwin
goarch: amd64
pkg: github.com/zyxan/benchmarks
BenchmarkGoregen-4    	 1000000	      1266 ns/op	     792 B/op	      15 allocs/op
BenchmarkReggen-4     	 3000000	       409 ns/op	      64 B/op	      13 allocs/op
BenchmarkZyxanGen-4   	 5000000	       258 ns/op	       8 B/op	       1 allocs/op
PASS
ok  	github.com/zyxan/benchmarks	4.522s
```

## State -  Work In Progress

### TODO:

- [ ] compute possible combinations with the given alphabet and pattern
- [ ] collision check
- [X] return consumable values
