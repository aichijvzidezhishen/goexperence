package generic

import (
	"fmt"
	"slices"
)

func SliceNewCha() {
	sl := []string{"hello", "world", "golang"}

	for i, s := range slices.All(sl) {
		fmt.Printf("%d : %s\n", i, s)
	}

	for i, s := range slices.Backward(sl) {
		fmt.Printf("%d : %s\n", i, s)
	}
}
