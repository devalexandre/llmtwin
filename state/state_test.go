package state_test

import (
	"testing"

	"github.com/devalexandre/llmtwin/state"
)

func TestState_UpdateAndGet(t *testing.T) {
	tests := []struct {
		name      string
		key       string
		value     interface{}
		expectKey string
		expectVal interface{}
		expectOk  bool
		update    bool // Flag to determine if `Update` should be called
	}{
		{
			name:      "Add new key-value pair",
			key:       "username",
			value:     "alex",
			expectKey: "username",
			expectVal: "alex",
			expectOk:  true,
			update:    true,
		},
		{
			name:      "Update existing key",
			key:       "age",
			value:     25,
			expectKey: "age",
			expectVal: 25,
			expectOk:  true,
			update:    true,
		},
		{
			name:      "Key not found",
			key:       "nonexistent",
			value:     nil,
			expectKey: "nonexistent",
			expectVal: nil,
			expectOk:  false,
			update:    false, // Do not update this key
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			state := state.NewState()

			// Update state with key-value pair if needed
			if tt.update {
				state.Update(tt.key, tt.value)
			}

			// Get the value from the state
			result, ok := state.Get(tt.expectKey)

			if ok != tt.expectOk {
				t.Errorf("expected exists: %v, got: %v", tt.expectOk, ok)
			}

			if result != tt.expectVal {
				t.Errorf("expected value: %v, got: %v", tt.expectVal, result)
			}
		})
	}
}

func TestState_EmptyState(t *testing.T) {
	state := state.NewState()

	// Attempt to get a key from an empty state
	_, exists := state.Get("missingKey")

	if exists {
		t.Errorf("expected key to not exist, but it does")
	}
}
