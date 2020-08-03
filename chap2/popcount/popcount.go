package popcount

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i >> 1] + byte(i & 1)
	}
}

func PopCount(x uint64) int {
	ans := 0
	for shift := 0; shift < 8 * 8; shift += 8 {
		ans += int(pc[x >> shift])
	}
	return ans
}

func Count2(x uint64) int {
	ans := 0
	for x != 0 {
		ans += int(x & 1)
		x >>= 1
	}
	return ans
}

func Count3(x uint64) int {
	ans := 0
	for x != 0 {
		ans += 1
		x = x & (x - 1)
	}
	return ans
}
