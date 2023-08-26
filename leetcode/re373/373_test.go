package re373

import (
	"fmt"
	"testing"
)

func TestHeap(t *testing.T) {
	h2 := &intHeap{}
	h2.pushInt([]int{1, 2})
	h2.pushInt([]int{1, 4})
	h2.pushInt([]int{2, 2})
	fmt.Printf("minimum: %d\n", h2.front())
	for !h2.isEmpty() {
		fmt.Printf("%d ", h2.popInt())
	}
}
