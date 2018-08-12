package regtgen

import (
	"errors"
	"log"
	"regexp/syntax"
	"strconv"
	"strings"

	"github.com/valyala/fastrand"
)

var (
	reg *syntax.Regexp
	err error
)

type block struct {
	validChars []byte
	flags      syntax.Flags
	rx         *syntax.Regexp
}

func Init(pattern string, validCharacters []byte) *block {
	if reg, err = syntax.Parse(pattern, 0); err != nil {
		log.Fatal(err)
	}

	b := &block{
		validChars: validCharacters,
		rx:         reg.Simplify(),
	}

	return b
}

func (b *block) GenerateMany(n int) ([]string, error) {
	res := []string{}
	for i := 0; i < n; i++ {
		val, err := b.generate()
		if err != nil {
			log.Panic(err)
		}
		res = append(res, val)
	}
	return res, nil
}

func (b *block) Generate() (string, error) {
	value, err := b.generate()
	if err != nil {
		log.Panic(err)
	}
	return value, nil
}

// Generate creates a random string from a regular expression
func (b *block) generate() (string, error) {
	switch b.rx.Op {
	case 3:
		sb := strings.Builder{}
		for i := 0; i < len(b.rx.Rune); i++ {
			if b.rx.Rune[i] == '#' {
				sb.WriteByte(b.validChars[fastrand.Uint32n(uint32(len(b.validChars)))])
			} else {
				sb.WriteRune(b.rx.Rune[i])
			}
		}
		return sb.String(), nil
	default:
		return "", errors.New("unsupported case: " + strconv.Itoa(int(b.rx.Op)))
	}
}
