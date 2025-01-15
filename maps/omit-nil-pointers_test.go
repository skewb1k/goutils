package maps

import (
	"reflect"
	"testing"
)

func ptr[T any](v T) *T {
	return &v
}

func TestOmitNilPointers(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string]any
		expected map[string]any
	}{
		{
			name: "map with nil values",
			input: map[string]any{
				"a": nil,
				"b": "hello",
				"c": ptr("world"),
			},
			expected: map[string]any{
				"b": "hello",
				"c": "world",
			},
		},
		{
			name: "map with nil pointers",
			input: map[string]any{
				"a": (*string)(nil),
				"b": ptr(42),
				"c": 100,
			},
			expected: map[string]any{
				"b": 42,
				"c": 100,
			},
		},
		{
			name: "map with mixed types",
			input: map[string]any{
				"str":    ptr("string"),
				"int":    ptr(123),
				"bool":   ptr(true),
				"direct": 456,
				"nil":    nil,
			},
			expected: map[string]any{
				"str":    "string",
				"int":    123,
				"bool":   true,
				"direct": 456,
			},
		},
		{
			name:     "empty map",
			input:    map[string]any{},
			expected: map[string]any{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := OmitNilPointers(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("OmitNilPointers() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestOmitNilPointersWithIntKeys(t *testing.T) {
	input := map[int]any{
		1: ptr("one"),
		2: nil,
		3: (*string)(nil),
		4: 42,
	}
	expected := map[int]any{
		1: "one",
		4: 42,
	}

	result := OmitNilPointers(input)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("OmitNilPointers() = %v, want %v", result, expected)
	}
}
