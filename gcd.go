package gomath

// GCDint64 returns Greatest Common Divisor of 2 int64
func GCDint64(a, b int64) int64 {
	c := a % b

	if c == 0 {
		return b
	}

	return GCDint64(b, c)
}
