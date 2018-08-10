package regtgen

// TODO: Allow get total possible number of permutations

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"regexp/syntax"
	"strings"
	"time"
)

type parser struct {
	pattern    string
	flags      syntax.Flags
	validChars []byte
	rx         *syntax.Regexp
}

func InitGenerator(pattern string, validCharacters []byte) *parser {
	p := &parser{pattern: pattern}
	reg, err := syntax.Parse(p.pattern, p.flags)
	if err != nil {
		log.Fatal(err)
	}

	p.validChars = validCharacters
	p.rx = reg.Simplify()
	return p
}

// Generate creates a random string parting from a regular expression
func (p *parser) generate() (string, error) {
	// Current generator supports only `OpCharClass'
	switch p.rx.Op {
	case 3:
		b := strings.Builder{}

		for i := 0; i < len(p.rx.Rune); i++ {
			if p.rx.Rune[i] == 35 {
				b.WriteByte(p.validChars[randomint(len(p.validChars))])
			} else {
				b.WriteRune(p.rx.Rune[i])
			}
		}
		return b.String(), nil
	default:
		return "", errors.New("unsupported case")
	}
}

func (p *parser) Generate(n int) (string, error) {
	for i := 0; i < n; i++ {
		value, err := p.generate()
		if err != nil {
			log.Panic(err)
		}
		fmt.Println(value)
	}
	return "", nil
}

// --- UTILS ---

func randomint(max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max)
}
