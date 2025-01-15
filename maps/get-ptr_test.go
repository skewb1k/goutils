package maps

import (
	"testing"
)

func TestPtrFromMap(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string]any
		key      string
		expected *string
	}{
		{
			name: "valid string value",
			input: map[string]any{
				"name": "John",
			},
			key:      "name",
			expected: strPtr("John"),
		},
		{
			name: "valid string value",
			input: map[string]any{
				"name": "John",
			},
			key:      "name",
			expected: strPtr("John"),
		},
		{
			name: "key not found",
			input: map[string]any{
				"age": 25,
			},
			key:      "name",
			expected: nil,
		},
		{
			name: "wrong type",
			input: map[string]any{
				"name": 42,
			},
			key:      "name",
			expected: nil,
		},
		{
			name:     "empty map",
			input:    map[string]any{},
			key:      "name",
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := PtrFromMap[string](tt.input, tt.key)

			if (result == nil && tt.expected != nil) ||
				(result != nil && tt.expected == nil) ||
				(result != nil && tt.expected != nil && *result != *tt.expected) {
				t.Errorf("PtrFromMap() = %v, expected %v", result, tt.expected)
			}
		})
	}

	// Test with different types
	t.Run("integer type", func(t *testing.T) {
		m := map[string]any{"age": 25}
		result := PtrFromMap[int](m, "age")
		if result == nil || *result != 25 {
			t.Errorf("PtrFromMap() = %v, expected 25", result)
		}
	})

	t.Run("boolean type", func(t *testing.T) {
		m := map[string]any{"active": true}
		result := PtrFromMap[bool](m, "active")
		if result == nil || *result != true {
			t.Errorf("PtrFromMap() = %v, expected true", result)
		}
	})
}

// Helper function to create string pointer
func strPtr(s string) *string {
	return &s
}
