package utils

import (
	"math/rand"
	"strings"
	"time"
)

var alphabets string = "abcdefghijklmnopqrstuvwxyz"
var currencies = []string{"USD", "RUB", "EUR"}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomNumber(min, max int64) (number int64) {
	return rand.Int63n(max-min+1) + min
}

func RandomName(n int) string {
	builder := strings.Builder{}
	k := len(alphabets)
	for i := 0; i < n; i++ {
		builder.WriteByte(alphabets[rand.Intn(k)])
	}
	return builder.String()
}

func RandomCurrency() string {
	return currencies[rand.Intn(2)]
}
