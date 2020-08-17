package gentest

import (
	"log"
	"math/rand"
	"time"
)

func getRandom() *rand.Rand {
	seed := time.Now().UTC().UnixNano()
	log.Printf("Random seed: %d\n", seed)
	rng := rand.New(rand.NewSource(seed))
	return rng
}

func RandomString(length int, rng *rand.Rand) string {
	runes := make([]rune, length)
	for i := 0; i < length; i++ {
		runes[i] = rune(rng.Intn(0x1000))
	}
	return string(runes)
}

func RandomStrings(n int, length int) []string {
	randStrings := make([]string, n)
	rng := getRandom()
	for i := 0; i < n; i++ {
		randStrings[i] = RandomString(length, rng)
	}
	return randStrings
}
