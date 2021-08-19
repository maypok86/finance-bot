package random

import (
	"math"
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// IntByRange generates a random integer between min and max
func IntByRange(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// Int generates a random integer between 1 and math.MaxInt32
func Int() int {
	return int(IntByRange(1, math.MaxInt32))
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

// Timestamp generates a random timestamp
func Timestamp() time.Time {
	min := time.Date(1970, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	max := time.Date(2070, 1, 0, 0, 0, 0, 0, time.UTC).Unix()

	return time.Unix(IntByRange(min, max), 0)
}
