package input

import (
	"math/rand"
	"time"
)

func combineStringArray(random bool, arrays_in ...[]string) (out []string) {
	if len(arrays_in) > 1 {

		var A []string
		var B []string
		for n, array := range arrays_in {

			if n == 0 {
				A = array
			} else {
				// segunda vuelta
				if n == 1 {
					B = array
				} else { //tercera vuelta o mas
					A = out
					B = array
				}
			}

			out = permutation(A, B)

		}

		if random {
			rand.New(rand.NewSource(time.Now().UnixNano())) // now
			// rand.Seed(time.Now().UnixNano())// before
			rand.Shuffle(len(out), func(i, j int) { out[i], out[j] = out[j], out[i] })
		}
		return
	} else if len(arrays_in) == 1 {
		return arrays_in[0]
	}
	return
}

func permutation(A, B []string) (out []string) {

	for _, a := range A {
		for _, b := range B {
			out = append(out, a+" "+b)
		}
	}

	return
}
