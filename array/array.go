package array

import (
	"math/rand"
	"time"
)

// ArrayShuffle shuffles an array.
// It takes an array of any type and returns a new shuffled array.
func ArrayShuffle[T any](arr []T) []T {
	newArr := make([]T, len(arr))
	copy(newArr, arr)
	rand.New(rand.NewSource(time.Now().UnixNano())).Shuffle(len(newArr), func(i, j int) {
		newArr[i], newArr[j] = newArr[j], newArr[i]
	})
	return newArr
}

// ArrayChunk chunks an array into smaller arrays of a specified size.
// It takes an array of any type and a chunk size, and returns a new 2D array.
func ArrayChunk[T any](arr []T, size int) [][]T {
	if size <= 0 {
		return nil
	}
	var chunks [][]T
	for i := 0; i < len(arr); i += size {
		end := i + size
		if end > len(arr) {
			end = len(arr)
		}
		chunks = append(chunks, arr[i:end])
	}
	return chunks
}

// ArrayUnique removes duplicate values from an array.
// It takes an array of a comparable type and returns a new array with unique values.
func ArrayUnique[T comparable](arr []T) []T {
	inResult := make(map[T]bool)
	result := make([]T, 0)
	for _, str := range arr {
		if _, ok := inResult[str]; !ok {
			inResult[str] = true
			result = append(result, str)
		}
	}
	return result
}

// InArray checks if a value exists in an array.
// It takes a value of a comparable type and an array of the same type,
// and returns true if the value is found in the array, false otherwise.
func InArray[T comparable](val T, arr []T) bool {
	for _, v := range arr {
		if v == val {
			return true
		}
	}
	return false
}

// IsArrayIntersect checks if two arrays have common elements.
// It takes two arrays of a comparable type and returns true if they have at least one common element, false otherwise.
func IsArrayIntersect[T comparable](arr1, arr2 []T) bool {
	s := make(map[T]struct{}, len(arr1))
	for _, v := range arr1 {
		s[v] = struct{}{}
	}

	for _, v := range arr2 {
		if _, ok := s[v]; ok {
			return true
		}
	}
	return false
}

// ArrayIntersection returns the intersection of two arrays.
// It takes two arrays of a comparable type and returns a new array containing elements that are present in both arrays.
func ArrayIntersection[T comparable](arr1, arr2 []T) []T {
	set := make(map[T]struct{})
	var intersection []T
	for _, v := range arr1 {
		set[v] = struct{}{}
	}
	for _, v := range arr2 {
		if _, ok := set[v]; ok {
			intersection = append(intersection, v)
		}
	}
	return ArrayUnique(intersection)
}

// ArrayUnion returns the union of two arrays.
// It takes two arrays of a comparable type and returns a new array containing all unique elements from both arrays.
func ArrayUnion[T comparable](arr1, arr2 []T) []T {
	return ArrayUnique(append(arr1, arr2...))
}

// ArrayDifference returns the difference of two arrays.
// It takes two arrays of a comparable type and returns a new array containing elements that are in the first array but not in the second.
func ArrayDifference[T comparable](arr1, arr2 []T) []T {
	set := make(map[T]struct{})
	difference := make([]T, 0)
	for _, v := range arr2 {
		set[v] = struct{}{}
	}
	for _, v := range arr1 {
		if _, ok := set[v]; !ok {
			difference = append(difference, v)
		}
	}
	return difference
}
