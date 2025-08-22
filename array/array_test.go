package array

import (
	"reflect"
	"sort"
	"testing"
)

func TestArrayShuffle(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	shuffled := ArrayShuffle(arr)

	if len(arr) != len(shuffled) {
		t.Errorf("Expected length %d, but got %d", len(arr), len(shuffled))
	}

	sort.Ints(shuffled)
	if !reflect.DeepEqual(arr, shuffled) {
		t.Errorf("Expected %v, but got %v", arr, shuffled)
	}
}

func TestArrayChunk(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		size int
		want [][]int
	}{
		{
			name: "Test with empty array",
			arr:  []int{},
			size: 2,
			want: nil,
		},
		{
			name: "Test with size > len(arr)",
			arr:  []int{1, 2, 3},
			size: 5,
			want: [][]int{{1, 2, 3}},
		},
		{
			name: "Test with size < len(arr)",
			arr:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			size: 3,
			want: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}, {10}},
		},
		{
			name: "Test with size = len(arr)",
			arr:  []int{1, 2, 3, 4, 5},
			size: 5,
			want: [][]int{{1, 2, 3, 4, 5}},
		},
		{
			name: "Test with zero size",
			arr:  []int{1, 2, 3, 4, 5},
			size: 0,
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ArrayChunk(tt.arr, tt.size); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ArrayChunk() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArrayUnique(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want []int
	}{
		{
			name: "Test with duplicates",
			arr:  []int{1, 2, 2, 3, 4, 4, 5},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "Test with no duplicates",
			arr:  []int{1, 2, 3, 4, 5},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "Test with empty array",
			arr:  []int{},
			want: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ArrayUnique(tt.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ArrayUnique() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsArrayIntersect(t *testing.T) {
	tests := []struct {
		name string
		arr1 []int
		arr2 []int
		want bool
	}{
		{
			name: "Test with intersection",
			arr1: []int{1, 2, 3},
			arr2: []int{3, 4, 5},
			want: true,
		},
		{
			name: "Test with no intersection",
			arr1: []int{1, 2, 3},
			arr2: []int{4, 5, 6},
			want: false,
		},
		{
			name: "Test with empty arrays",
			arr1: []int{},
			arr2: []int{},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsArrayIntersect(tt.arr1, tt.arr2); got != tt.want {
				t.Errorf("IsArrayIntersect() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInArray(t *testing.T) {
	tests := []struct {
		name string
		val  int
		arr  []int
		want bool
	}{
		{
			name: "Test with value in array",
			val:  3,
			arr:  []int{1, 2, 3, 4, 5},
			want: true,
		},
		{
			name: "Test with value not in array",
			val:  6,
			arr:  []int{1, 2, 3, 4, 5},
			want: false,
		},
		{
			name: "Test with empty array",
			val:  1,
			arr:  []int{},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InArray(tt.val, tt.arr); got != tt.want {
				t.Errorf("InArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArrayIntersection(t *testing.T) {
	tests := []struct {
		name string
		arr1 []int
		arr2 []int
		want []int
	}{
		{
			name: "Test with intersection",
			arr1: []int{1, 2, 3},
			arr2: []int{3, 4, 5},
			want: []int{3},
		},
		{
			name: "Test with no intersection",
			arr1: []int{1, 2, 3},
			arr2: []int{4, 5, 6},
			want: []int{},
		},
		{
			name: "Test with empty arrays",
			arr1: []int{},
			arr2: []int{},
			want: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ArrayIntersection(tt.arr1, tt.arr2)
			sort.Ints(got)
			sort.Ints(tt.want)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ArrayIntersection() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArrayUnion(t *testing.T) {
	tests := []struct {
		name string
		arr1 []int
		arr2 []int
		want []int
	}{
		{
			name: "Test with union",
			arr1: []int{1, 2, 3},
			arr2: []int{3, 4, 5},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "Test with no common elements",
			arr1: []int{1, 2, 3},
			arr2: []int{4, 5, 6},
			want: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name: "Test with empty arrays",
			arr1: []int{},
			arr2: []int{},
			want: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ArrayUnion(tt.arr1, tt.arr2)
			sort.Ints(got)
			sort.Ints(tt.want)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ArrayUnion() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArrayDifference(t *testing.T) {
	tests := []struct {
		name string
		arr1 []int
		arr2 []int
		want []int
	}{
		{
			name: "Test with difference",
			arr1: []int{1, 2, 3},
			arr2: []int{3, 4, 5},
			want: []int{1, 2},
		},
		{
			name: "Test with no common elements",
			arr1: []int{1, 2, 3},
			arr2: []int{4, 5, 6},
			want: []int{1, 2, 3},
		},
		{
			name: "Test with empty arrays",
			arr1: []int{},
			arr2: []int{},
			want: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ArrayDifference(tt.arr1, tt.arr2)
			sort.Ints(got)
			sort.Ints(tt.want)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ArrayDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}
