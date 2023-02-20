package utils

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRandomCurrency(t *testing.T) {
	curr := RandomCurrency()
	require.Contains(t, currencies, curr)
}

func TestRandomName(t *testing.T) {
	for i := 5; i < 11; i++ {
		name := RandomName(i)
		nameBytes := []byte(name)
		for _, char := range nameBytes {
			require.Contains(t, []byte(alphabets), char)
		}
		require.Len(t, nameBytes, i)
	}
}

func TestRandomNumber(t *testing.T) {
	for min := 0; min < 100; min++ {
		for max := 100; max < 1000; max += 10 {
			number := RandomNumber(int64(min), int64(max))
			require.LessOrEqual(t, int64(min), number)
			require.GreaterOrEqual(t, int64(max), number)
		}
	}
}
