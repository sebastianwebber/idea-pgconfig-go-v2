package config

import (
	"fmt"
	"math/bits"
)

func format(t OutputType, value interface{}) string {

	if t == Bytes {
		return humanBytes(value.(uint64))
	}

	if t == Number {
		return fmt.Sprintf("%d", value)
	}

	return fmt.Sprintf("%s", value)
}

func humanBytes(bytes uint64) string {
	if bytes < 1024 {
		return fmt.Sprintf("%d bytes", bytes)
	}

	base := uint(bits.Len64(bytes) / 10)
	val := float64(bytes) / float64(uint64(1<<(base*10)))

	return fmt.Sprintf("%.0f%cB", val, " KMGTPE"[base])
}
