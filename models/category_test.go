package models

import "testing"

func TestReturnsNonEmptySlice(t *testing.T) {
	categories := GetCategories()

	if len(categories) == 0 {
		t.Errorf("Expected %d, got %d", 0, len(categories))
	}
}
