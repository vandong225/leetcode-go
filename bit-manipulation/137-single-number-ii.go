package bitmanipulation

func singleNumberII(nums []int) int {
	res := 0

	for i := 0; i < 32; i++ {
		cnt := 0

		for _, v := range nums {
			if (v>>i)&1 == 1 {
				cnt++
			}
		}

		if cnt%3 != 0 {
			res |= 1 << i
		}
	}

	return int(int32(res))
}
