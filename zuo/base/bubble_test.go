package base

import (
	"fmt"
	"testing"
)

func TestBubble(t *testing.T) {
	fmt.Println(bubble([]int{7, 2, 3, 8, 4, 5, 6, 1}))
	fmt.Println(bubble([]int{7, 1, 8, 8, 4, 9, 6, 1}))
}
