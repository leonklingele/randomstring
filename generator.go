package randomstring

import (
	"crypto/rand"
	"errors"
	"math"
	"math/big"
	"unicode/utf8"
)

const (
	// CharsNum contains numbers from 0-9
	CharsNum = "0123456789"
	// CharsAlpha contains the full English alphabet: letters a-z and A-Z
	CharsAlpha = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	// CharsAlphaNum is a combination of CharsNum and CharsAlpha
	CharsAlphaNum = CharsNum + CharsAlpha
	// CharsASCII contains all printable ASCII characters in code range [32, 126]
	CharsASCII = CharsAlphaNum + " !\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~"
)

var (
	// ErrInvalidLengthSpecified is returned when the length specified is invalid
	ErrInvalidLengthSpecified = errors.New("invalid password length specified")
	// ErrInvalidDictSpecified is returned when the dictionary specified is invalid
	ErrInvalidDictSpecified = errors.New("invalid password dictionary specified")
)

type Generator struct {
	l    int
	dict []rune
	max  *big.Int
}

func (g *Generator) Generate() (string, error) {
	buf := make([]rune, g.l)
	for i := 0; i < g.l; i++ {
		index, err := randomInt(g.max)
		if err != nil {
			return "", err
		}

		buf[i] = g.dict[index]
	}

	return string(buf), nil
}

func NewGenerator(l int, dict string) (*Generator, error) {
	// Length needs to be in range [1, 1<<31-1]
	if l <= 0 || l > math.MaxInt32 {
		return nil, ErrInvalidLengthSpecified
	}

	dlen := utf8.RuneCountInString(dict)
	if dlen == 0 {
		return nil, ErrInvalidDictSpecified
	}

	max := big.NewInt(int64(dlen))

	return &Generator{
		l:    l,
		dict: []rune(dict),
		max:  max,
	}, nil
}

// Generate generates a cryptographically secure and unbiased random string of length 'l' using alphabet 'dict'
func Generate(l int, dict string) (string, error) {
	gen, err := NewGenerator(l, dict)
	if err != nil {
		return "", err
	}

	return gen.Generate()
}

func randomInt(max *big.Int) (int, error) {
	i, err := rand.Int(rand.Reader, max)
	if err != nil {
		return 0, err
	}

	return int(i.Int64()), nil
}
