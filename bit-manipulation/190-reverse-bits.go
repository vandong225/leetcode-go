package bitmanipulation

func reverseBits(n int) int {
	res := 0

	for i := 0; i < 32; i++ {
		bit := n & 1
		res = (res << 1) + bit
		n = n >> 1
	}

	return res
}
