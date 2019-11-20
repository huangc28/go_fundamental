package statrannum

import "math/rand"

const MaxRand = 16

func StatRandomNumbers(n int) (int, int) {
	var a, b int

	for i := 0; i < n; i++ {
		if rand.Intn(MaxRand) < MaxRand/2 {
			a = a + 1
		} else {
			b++
		}
	}

	return a, b
}
