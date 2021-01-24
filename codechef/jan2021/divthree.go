package jan2021

import (
	"fmt"
	"math"

	"github.com/amiedeep/go-lang-competitive/utils"
)

func main() {
	var t, n, k, d, s, tmp int

	fmt.Scan(&t)

	o := make([]int, t)

	for i := 0; i < t; i++ {
		fmt.Scan(&n, &k, &d)

		s = 0

		for i := 0; i < n; i++ {
			fmt.Scan(&tmp)
			s += tmp
		}

		o[i] = int(math.Min(float64(s/k), float64(d)))

	}

	utils.PrintSlice(o)

}
