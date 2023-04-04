package randomstring_test

import (
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
		l := l

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
		l := l

		t.Run(fmt.Sprintf("len-%d", l), func(t *testing.T) {
			t.Parallel()

			if _, err := randomstring.Generate(l, dict); err != randomstring.ErrInvalidLengthSpecified {
				t.Errorf("unexpected error for length %d: %v", l, err)
			}
		})
	}
}

func TestRandomStringInvalidChars(t *testing.T) {
	t.Parallel()

	if _, err := randomstring.Generate(1, ""); err != randomstring.ErrInvalidDictSpecified {
		t.Errorf("unexpected error using empty dictionary: %v", err)
	}
}

func TestRandomStringWithNonASCII(t *testing.T) {
	t.Parallel()

	const (
		dict = "世界"
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
