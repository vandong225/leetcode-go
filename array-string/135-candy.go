package arraystring

// func candy(ratings []int) int {
// 	n := len(ratings)
// 	lr := make([]int, n)
// 	rl := make([]int, n)

// 	for i := 0; i < n; i++ {
// 		lr[i] = 1
// 		rl[i] = 1
// 	}

// 	for i := 1; i < n; i++ {
// 		if ratings[i] > ratings[i-1] {
// 			lr[i] = lr[i-1] + 1
// 		}
// 	}

// 	for i := n - 2; i >= 0; i-- {
// 		if ratings[i] > ratings[i+1] {
// 			rl[i] = rl[i+1] + 1
// 		}
// 	}

// 	candy := 0
// 	for i := 0; i < n; i++ {
// 		candy += int(math.Max(float64(lr[i]), float64(rl[i])))
// 	}
// 	return candy
// }

func candy(ratings []int) int {
	n := len(ratings)
	inc, desc, peak, candy := 1, 0, 1, 1

	for i := 1; i < n; i++ {
		if ratings[i] > ratings[i-1] {
			inc++
			desc = 0
			candy += inc
			peak = inc
		} else if ratings[i] < ratings[i-1] {
			desc++
			inc = 1
			candy += desc
			if desc >= peak {
				candy++
			}
		} else {
			inc = 1
			desc = 0
			candy += inc
			peak = inc
		}

	}

	return candy
}
