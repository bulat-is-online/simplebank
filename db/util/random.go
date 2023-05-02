package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdeghiklmnoprtsuwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomInt generates a random integer between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomString generates random string with lenght n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// Returns randomly generated Owner name by running RandomString
func RandomOwner() string {
	return RandomString(6)
}

// Randomly generates balance
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

// Randomly defines currency code

func RandomCurrency() string {
	currencies := []string{USD, EUR, CAD}
	n := len(currencies)
	return currencies[rand.Intn(n-1)]
}

// Returns ID of a random account. Wont work if amount of records ar less than max range
func RandomID() int64 {
	return RandomInt(3, 11)
}
