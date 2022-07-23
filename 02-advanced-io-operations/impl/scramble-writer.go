package impl

import (
	"io"
	"math/rand"
	"unicode"
	"unicode/utf8"
)

type ScrambleWriter struct {
	w io.Writer
	r *rand.Rand
	c float64
}

func NewScrambleWriter(w io.Writer, r *rand.Rand, chance float64) *ScrambleWriter {
	return &ScrambleWriter{w: w, r: r, c: chance}
}

func (s *ScrambleWriter) shambleWrite(runes []rune, sep rune) (n int, err error) {
	// scramble after first letter
	for i := 1; i < len(runes)-1; i++ {
		if s.r.Float64() > s.c {
			continue
		}

		j := s.r.Intn(len(runes)-i) + 1
		runes[i], runes[j] = runes[j], runes[i]
	}
	if sep != 0 {
		runes = append(runes, sep)
	}

	var b = make([]byte, 10)

	for _, r := range runes {
		v, err := s.w.Write(b[:utf8.EncodeRune(b, r)])
		if err != nil {
			return n, err
		}
		n += v
	}
	return
}

func (s *ScrambleWriter) Write(b []byte) (n int, err error) {
	var runes = make([]rune, 0, 10)

	for r, i, w := rune(0), 0, 0; i < len(b); i += w {
		r, w = utf8.DecodeRune(b[i:])
		if unicode.IsLetter(r) {
			runes = append(runes, r)
			continue
		}

		v, err := s.shambleWrite(runes, r)
		if err != nil {
			return n, err
		}
		n += v
		runes = runes[:0]
	}

	if len(runes) != 0 {
		v, err := s.shambleWrite(runes, 0)
		if err != nil {
			return n, err
		}
		n += v
	}
	return
}
