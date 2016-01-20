package scope

import (
	"reflect"
	"testing"
)

func TestAdd(t *testing.T) {
	// Empty
	checkAdd(t, []string{}, []string{}, []string{})

	// No new scopes
	checkAdd(t, []string{"A"}, []string{}, []string{"A"})

	// Duplicates
	checkAdd(t, []string{"A"}, []string{"A"}, []string{"A"})

	// Unsorted
	checkAdd(t, []string{"B", "A"}, []string{"A", "B"}, []string{"A", "B"})

	// Additional new scopes
	checkAdd(t, []string{"B", "A"}, []string{"C", "A", "B"}, []string{"A", "B", "C"})

	// No existing scopes
	checkAdd(t, []string{}, []string{"C", "A", "B"}, []string{"A", "B", "C"})
}

func checkAdd(t *testing.T, s1, s2, expected []string) {
	actual := Add(s1, s2)
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v + %v to be %v, but got %v", s1, s2, expected, actual)
	}
}

func TestCovers(t *testing.T) {
	// Empty request
	checkCovers(t, []string{}, []string{}, true)
	checkCovers(t, []string{"A"}, []string{}, true)
	checkCovers(t, []string{"B", "A"}, []string{}, true)

	// Equal request
	checkCovers(t, []string{"A"}, []string{"A"}, true)
	// Covered request
	checkCovers(t, []string{"B", "A"}, []string{"A"}, true)
	// Sorting difference
	checkCovers(t, []string{"B", "A"}, []string{"A", "B"}, true)
	// Superset
	checkCovers(t, []string{"B", "A", "C"}, []string{"A", "B"}, true)

	// Empty has
	checkCovers(t, []string{}, []string{"A"}, false)
	// Different has
	checkCovers(t, []string{"B"}, []string{"A"}, false)
	// Partially overlapping has
	checkCovers(t, []string{"A", "B"}, []string{"A", "C"}, false)
}

func checkCovers(t *testing.T, has, requested []string, expected bool) {
	actual := Covers(has, requested)
	if actual != expected {
		if expected {
			t.Errorf("Expected %v to cover %v, but it did not", has, requested)
		} else {
			t.Errorf("Expected %v to not cover %v, but it did", has, requested)
		}
	}
}
