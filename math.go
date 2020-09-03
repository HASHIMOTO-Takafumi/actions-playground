package main

func gcd(x int, y int) int {
	a, b := x, y
	if b > a {
		a, b = b, a
	}
	for {
		r := a % b
		if r == 0 {
			return b
		}
		a = b
		b = r
	}
}
