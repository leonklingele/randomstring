package randomstring_test

import (
	"errors"
	"fmt"
	"testing"
	"unicode/utf8"

	"github.com/leonklingele/randomstring/v2"
)

func TestRandomStringLength(t *testing.T) {
	t.Parallel()

	ls := []int{1, 5, 10, 50, 100}
	dict := randomstring.CharsASCII

	for _, l := range ls {
		t.Run(fmt.Sprintf("len-%d", l), func(t *testing.T) {
			t.Parallel()

			s, err := randomstring.Generate(l, dict)
			if err != nil {
				t.Error(err)
			}

			if len(s) != l {
				t.Errorf("unexpected length of generated string: want %d, got %d", l, len(s))
			}
		})
	}
}

func TestRandomStringInvalidLength(t *testing.T) {
	t.Parallel()

	ls := []int{-1, 0}
	dict := randomstring.CharsAlpha

	for _, l := range ls {
		t.Run(fmt.Sprintf("len-%d", l), func(t *testing.T) {
			t.Parallel()

			if _, err := randomstring.Generate(l, dict); !errors.Is(err, randomstring.ErrInvalidLengthSpecified) {
				t.Errorf("unexpected error for length %d: %v", l, err)
			}
		})
	}
}

func TestRandomStringInvalidChars(t *testing.T) {
	t.Parallel()

	if _, err := randomstring.Generate(1, ""); !errors.Is(err, randomstring.ErrInvalidDictSpecified) {
		t.Errorf("unexpected error using empty dictionary: %v", err)
	}
}

func TestRandomStringWithNonASCII(t *testing.T) {
	t.Parallel()

	const (
		dict = "世界" //nolint:gosmopolitan // We explicitly want to test this
		l    = 5
	)

	s, err := randomstring.Generate(l, dict)
	if err != nil {
		t.Fatal(err)
	}

	if got := utf8.RuneCountInString(s); got != l {
		t.Fatalf("invalid length of string, got %d, want %d", got, l)
	}
}

func BenchmarkGenerate(b *testing.B) {
	const (
		l    = 32
		dict = randomstring.CharsAlphaNum
	)

	for b.Loop() {
		s, err := randomstring.Generate(l, randomstring.CharsAlphaNum)
		if err != nil {
			b.Fatal(err)
		}

		if len(s) != l {
			b.Fatalf("invalid length of %q: %d != %d", s, len(s), l)
		}
	}
}

func BenchmarkGenerator(b *testing.B) {
	const (
		l    = 32
		dict = randomstring.CharsAlphaNum //nolint:goconst // We don't care to duplicate function-scoped consts
	)

	genl32, err := randomstring.NewGenerator(l, dict)
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()

	for b.Loop() {
		s, err := genl32.Generate()
		if err != nil {
			b.Fatal(err)
		}

		if len(s) != l {
			b.Fatalf("invalid length of %q: %d != %d", s, len(s), l)
		}
	}
}
