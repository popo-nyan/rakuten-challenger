package rc

import (
	"crypto/rand"
	"strings"
)

type Mdata struct {
	Mask string `json:"mask"`
	Key  string `json:"key"`
	Seed uint32 `json:"seed"`
}

func validateAgainstMask(e, t string) bool {
	if t == "" {
		return true
	}
	return strings.HasPrefix(e, t)
}

func randomStringGenerator(eLength int, tLength int) string {
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	n := tLength - eLength
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return ""
	}
	for i := range b {
		b[i] = charset[int(b[i])%len(charset)]
	}
	return string(b)
}

func SolvePow(e string, t uint32, n string) string {
	var a string

	for {
		randomStr := randomStringGenerator(len(e), 16)
		a = e + randomStr
		c := x64hash128(a, t)
		if validateAgainstMask(c, n) {
			break
		}
	}
	return a
}
