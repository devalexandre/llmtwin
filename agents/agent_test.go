package agents_test

import (
	"errors"
	"testing"

	"github.com/devalexandre/llmtwin/agents"
	"github.com/devalexandre/llmtwin/state"
)

func TestAgent_RegisterAndExecuteTool(t *testing.T) {
	tests := []struct {
		name       string
		toolName   string
		toolFunc   func(state.State) (string, error)
		inputState state.State
		expectErr  bool
		expectRes  string
	}{
		{
			name:     "Tool executes successfully",
			toolName: "greet",
			toolFunc: func(s state.State) (string, error) {
				return "Hello, " + s.Data["name"].(string), nil
			},
			inputState: state.State{Data: map[string]interface{}{"name": "Alex"}},
			expectErr:  false,
			expectRes:  "Hello, Alex",
		},
		{
			name:     "Tool execution fails",
			toolName: "errorTool",
			toolFunc: func(s state.State) (string, error) {
				return "", errors.New("execution failed")
			},
			inputState: state.State{},
			expectErr:  true,
			expectRes:  "",
		},
		{
			name:       "Tool not found",
			toolName:   "nonexistent",
			toolFunc:   nil,
			inputState: state.State{},
			expectErr:  true,
			expectRes:  "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			agent := agents.NewAgent()

			// Register tool if toolFunc is provided
			if tt.toolFunc != nil {
				agent.RegisterTool(tt.toolName, tt.toolFunc)
			}

			// Execute the tool
			result, err := agent.Execute(tt.inputState, tt.toolName)

			if (err != nil) != tt.expectErr {
				t.Errorf("expected error: %v, got: %v", tt.expectErr, err)
			}

			if result != tt.expectRes {
				t.Errorf("expected result: %v, got: %v", tt.expectRes, result)
			}
		})
	}
}
