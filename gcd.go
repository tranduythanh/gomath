package gomath

// GCDint64 returns Greatest Common Divisor of 2 int64
func GCD(a, b int64) int64 {
	c := a % b

	if c == 0 {
		return b
	}

	return GCD(b, c)
}
