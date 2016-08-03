package gomath

// LCMint64 returns Least Common Multiple of 2 int64
func LCMint64(a, b int64) int64 {
	return a * b / GCDint64(a, b)
}
