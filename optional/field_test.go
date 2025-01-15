package optional

import (
	"encoding/json"
	"testing"
)

func pointerOf[T any](value T) *T {
	return &value
}

// todo: add more tests
func TestFieldUnmarshaling(t *testing.T) {
	type s struct {
		Avatar Field[string] `json:"avatar"`
	}

	tests := []struct {
		name        string
		input       string
		wantDefined bool
		wantValue   *string
	}{
		{
			name:        "with field value",
			input:       `{"avatar":"123"}`,
			wantDefined: true,
			wantValue:   pointerOf("123"),
		},
		{
			name:        "without field",
			input:       `{}`,
			wantDefined: false,
			wantValue:   nil,
		},
		{
			name:        "with null field",
			input:       `{"avatar":null}`,
			wantDefined: true,
			wantValue:   nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s s
			if err := json.Unmarshal([]byte(tt.input), &s); err != nil {
				t.Fatalf("failed to unmarshal JSON: %v", err)
			}

			if got := s.Avatar.Defined; got != tt.wantDefined {
				t.Errorf("Defined = %v, want %v", got, tt.wantDefined)
			}

			if tt.wantValue == nil {
				if s.Avatar.Value != nil {
					t.Errorf("Value = %v, want nil", *s.Avatar.Value)
				}
			} else if s.Avatar.Value == nil {
				t.Error("Value is nil, want non-nil")
			} else if *s.Avatar.Value != *tt.wantValue {
				t.Errorf("Value = %v, want %v", *s.Avatar.Value, *tt.wantValue)
			}
		})
	}
}
