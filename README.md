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
	res, err := exp.Generate(3)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(res)
}
```

## State -  Work In Progress

### TODO:

- compute possible combinations with the given alphabet and pattern
- collision check
- return consumable values