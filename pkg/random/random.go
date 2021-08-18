package random

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Int generates a random integer between min and max
func Int(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// String generates a random string of length n
func String(size int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < size; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// Name generates a random owner name
func Name() string {
	return String(255)
}
