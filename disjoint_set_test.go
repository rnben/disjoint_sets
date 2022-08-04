package graph

import (
	"testing"
)

func TestNewDisjointSet(t *testing.T) {
	values := []int{1, 2, 3, 4, 5, 6, 8, 9, 10}
	set := NewDisjointSet(values)

	for _, v := range values {
		if got := set.FindHead(v); v != got {
			t.Fatalf("Expected head %d, got %d", v, got)
		}
	}
}

func TestUnion(t *testing.T) {
	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	set := NewDisjointSet(values)

	set.Union(1, 2)
	set.Union(2, 2)
	set.Union(2, 3)
	set.Union(3, 4)
	set.Union(4, 5)
	set.Union(5, 6)
	set.Union(6, 7)
	set.Union(7, 8)
	set.Union(9, 8)

	// test not same set
	if set.IsSameSet(1, 10) {
		t.Fatal("Expected not same set")
	}

	// swap small, big
	set.Union(9, 10)

	// first elem  exist
	if set.IsSameSet(-1, 10) {
		t.Fatal("Expected not same set")
	}

	// second elem  not exist
	if set.IsSameSet(1, -10) {
		t.Fatal("Expected not same set")
	}

	// test same set
	if !set.IsSameSet(1, 10) {
		t.Fatal("Expected same set")
	}

	// check all union result
	if len(set.sizeMap) != 1 {
		t.Fatalf("Expected len(sizeMap) = 1, got %d", len(set.sizeMap))
	}

	var (
		head  element[int]
		count int
	)
	for head, count = range set.sizeMap {
	}

	if count != len(values) {
		t.Fatalf("Expected count = %d values, got %d", len(values), count)
	}

	for _, v := range values {
		if got := set.FindHead(v); head.value != got {
			t.Fatalf("Expected %d, got %d", v, got)
		}
	}
}
