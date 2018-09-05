package leetcode

import (
	"fmt"
	"testing"
)

func TestTopKFrequent(t *testing.T) {
	topK := topKFrequent([]int{1, 2, 2, 1, 1, 1, 1, 2, 2, 3, 4, 4, 4, 5}, 3)
	if fmt.Sprint(topK) != "[1 2 4]" {
		t.Error("Not the top K", topK)
	}
}
