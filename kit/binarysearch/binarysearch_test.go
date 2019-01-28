package binarysearch

import (
	"log"
	"testing"
)

func TestBinarySearchLarge(t *testing.T) {
	log.Println("Case 1:", BinarySearchLarge([]int{0, 3}, 0, 1, 2))
	log.Println("Case 2:", BinarySearchLarge([]int{2, 3}, 0, 1, 2))
	log.Println("Case 3:", BinarySearchLarge([]int{0, 2}, 0, 1, 2))
	log.Println("Case 4:", BinarySearchLarge([]int{0, 1, 2}, 0, 2, 2))
	log.Println("Case 5:", BinarySearchLarge([]int{0, 1, 2}, 0, 2, 1))
	log.Println("Case 6:", BinarySearchLarge([]int{0, 1, 2}, 0, 2, -1))
	log.Println("Case 7:", BinarySearchLarge([]int{0, 1, 2}, 0, 2, 3))
}
