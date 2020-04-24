package gcd

func gcd(x uint, y uint) uint {

	var shift uint = 0

	if x == y || y == 0 {
		return x
	}

	if x == 0 {
		return y
	}

	for shift := 0; (x|y)&1 == 0; shift++ {
		x = x >> 1
		y = y >> 1
	}

	for (x & 1) == 0 {
		x = x >> 1
	}

	for y == 0 {

		for (y & 1) == 0 {
			y = y >> 1
		}

		if x > y {
			t := x
			x = y
			y = t
		}

		y = y - x

	}

	y = y << shift

	return y
}