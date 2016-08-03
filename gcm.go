package gomath

func GCD(a, b int64) int64 {
	c := a % b

	if c == 0 {
		return b
	}

	return GCD(b, c)
}
