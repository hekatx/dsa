package utils

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomString() string {
	var sb strings.Builder
	k := len(alphabet)

	n := int(RandomInt(3, 10))

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func GetMockValues() []any {
	v := []any{}

	vlen := int(RandomInt(4, 10))

	for i := 0; i < vlen; i++ {
		s := RandomString()
		v = append(v, s)
	}

	return v
}
