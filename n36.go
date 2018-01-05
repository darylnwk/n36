package n36

import (
	"errors"
	"math"
	"math/rand"
	"strings"
	"time"
)

// N36 represents a numeric map
type N36 struct {
	charset  string
	seedRand *rand.Rand
}

const (
	CharRange62 = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	CharRange36 = "0123456789abcdefghijklmnopqrstuvwxyz"
)

// New creates a new n36 numeric map
func New(charset string) *N36 {
	n := &N36{
		charset:  charset,
		seedRand: rand.New(rand.NewSource(time.Now().UnixNano())),
	}

	return n
}

// Iton converts a uint64 value to string
func (n *N36) Iton(i uint64) string {

	r := ""
	l := uint64(len(n.charset))

	for ; i > 0; i = i / l {
		j := i % l
		r = n.charset[j:j+1] + r
	}

	return r
}

// Ntoi converts a string to uint64
func (n *N36) Ntoi(s string) (uint64, error) {
	var r uint64
	base := len(n.charset)
	fbase := float64(base)
	l := len(s)

	r = 0

	for i := 0; i < l; i++ {
		n := strings.Index(n.charset, s[i:i+1])

		if n < 0 {
			return 0, errors.New("n36.Ntoi: character not part of charset")
		}

		//n * base^(l-i-1) + r
		r = uint64(n)*uint64(math.Pow(fbase, float64(l-i)-1)) + r
	}

	return uint64(r), nil
}

// Random creates an (l)-long random string based on the character set
func (n *N36) Random(l int) string {
	max := len(n.charset)
	runes := []rune(n.charset)

	b := make([]rune, l)
	for i := range b {
		b[i] = runes[n.seedRand.Intn(max)]
	}

	return string(b)
}
