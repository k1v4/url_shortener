package randomGen

import (
	"math/rand"
	"time"
)

func NewRandomString(size int) string {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

	// алфавит
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"abcdefghijklmnopqrstuvwxyz" +
		"0123456789" + "_")

	b := make([]rune, size)

	for i := range b {
		b[i] = chars[rnd.Intn(len(chars))]
	}

	return string(b)
}
