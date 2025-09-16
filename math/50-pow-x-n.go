package math

// x^n = (x^2)^n/2 n%2== 0
// x^n = x.(x^2)^(n-1)/2
// x^n = (1/x)^-n n <0

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}

	if n == 1 {
		return x
	}

	if n < 0 {
		return myPow(1/x, -n)
	}

	if n%2 != 0 {
		return x * myPow(x*x, (n-1)/2)
	}

	return myPow(x*x, n/2)
}
