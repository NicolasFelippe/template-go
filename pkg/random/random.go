package random

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Int generates a random integer between min and max
func Int(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// String generates a random string of length n
func String(n int) string {
	var sb strings.Builder
	k := len(alphabet)
	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

// Owner generates a random owner name
func Owner() string {
	return String(6)
}

// Money generates a random amount of money
func Money() int64 {
	return Int(0, 10000)
}

// Currency generates a random currency code
func Currency() string {
	currencies := []string{"EUR", "USD", "CAD", "BRL"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}

// Currency generates a random email
func Email() string {
	return fmt.Sprintf("%s@email.com", String(6))
}
