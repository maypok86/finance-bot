package random

import (
	"crypto/rand"
	"math"
	"math/big"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

// IntByRange generates a random integer between min and max.
func IntByRange(min, max int64) int64 {
	r, _ := rand.Int(rand.Reader, big.NewInt(max-min+1))
	return min + r.Int64()
}

// Int generates a random integer between 1 and math.MaxInt32.
func Int() int {
	return int(IntByRange(1, math.MaxInt32))
}

// String generates a random string of length n.
func String(size int) string {
	var sb strings.Builder

	for i := 0; i < size; i++ {
		r, _ := rand.Int(rand.Reader, big.NewInt(int64(len(alphabet))))
		c := alphabet[r.Int64()]
		sb.WriteByte(c)
	}

	return sb.String()
}

// Name generates a random owner name.
func Name() string {
	return String(255)
}

// Timestamp generates a random timestamp.
func Timestamp() time.Time {
	min := time.Date(1970, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	max := time.Date(2070, 1, 0, 0, 0, 0, 0, time.UTC).Unix()

	return time.Unix(IntByRange(min, max), 0)
}
