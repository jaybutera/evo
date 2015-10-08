package perm

import (
	"math/rand"
)

// RandSlice returns a random slice of the argument along with the boundaries.
// That is to say:
//     sub == slice[left:right]
func RandSlice(slice []int) (sub []int, left, right int) {
	left = rand.Intn(len(slice))
	right = left
	for right == left {
		right = rand.Intn(len(slice))
	}
	if right < left {
		left, right = right, left
	}
	return slice[left:right], left, right
}

// Search searches an int slice for a particular value and returns the index.
// If the value is not found, Search returns -1.
func Search(slice []int, val int) (idx int) {
	for idx = range slice {
		if slice[idx] == val {
			return idx
		}
	}
	return -1
}

// Reverse reverses an int slice.
func Reverse(slice []int) {
	i := 0
	j := len(slice) - 1
	for i < j {
		slice[i], slice[j] = slice[j], slice[i]
		i++
		j--
	}
}

// Validate panics if the argument is not a permutation.
// This can be useful when testing custom operators.
func Validate(slice []int) {
	n := len(slice)
	for i := 0; i < n; i++ {
		if Search(slice, i) == -1 {
			panic("invalid permutation")
		}
	}
}
