package pipeline_test

import (
	"errors"
	"testing"

	"github.com/devalexandre/llmtwin/pipeline"
)

type mockStage struct {
	shouldExecute bool
	processFunc   func(data map[string]interface{}) (map[string]interface{}, error)
}

func (m *mockStage) ShouldExecute(data map[string]interface{}) bool {
	return m.shouldExecute
}

func (m *mockStage) Process(data map[string]interface{}) (map[string]interface{}, error) {
	return m.processFunc(data)
}

func TestPipeline_Execute(t *testing.T) {
	tests := []struct {
		name        string
		stages      []pipeline.Stage
		inputData   map[string]interface{}
		expectData  map[string]interface{}
		expectError bool
	}{
		{
			name: "Single stage executes successfully",
			stages: []pipeline.Stage{
				&mockStage{
					shouldExecute: true,
					processFunc: func(data map[string]interface{}) (map[string]interface{}, error) {
						data["processed"] = true
						return data, nil
					},
				},
			},
			inputData:   map[string]interface{}{"initial": "value"},
			expectData:  map[string]interface{}{"initial": "value", "processed": true},
			expectError: false,
		},
		{
			name: "Stage does not execute when ShouldExecute is false",
			stages: []pipeline.Stage{
				&mockStage{
					shouldExecute: false,
					processFunc: func(data map[string]interface{}) (map[string]interface{}, error) {
						data["processed"] = true
						return data, nil
					},
				},
			},
			inputData:   map[string]interface{}{"initial": "value"},
			expectData:  map[string]interface{}{"initial": "value"},
			expectError: false,
		},
		{
			name: "Stage returns error during processing",
			stages: []pipeline.Stage{
				&mockStage{
					shouldExecute: true,
					processFunc: func(data map[string]interface{}) (map[string]interface{}, error) {
						return nil, errors.New("processing failed")
					},
				},
			},
			inputData:   map[string]interface{}{"initial": "value"},
			expectData:  nil,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := pipeline.NewPipeline()
			for _, stage := range tt.stages {
				p.AddStage(stage)
			}

			result, err := p.Execute(tt.inputData)

			if (err != nil) != tt.expectError {
				t.Errorf("expected error: %v, got: %v", tt.expectError, err)
			}

			if !tt.expectError && !equalMaps(result, tt.expectData) {
				t.Errorf("expected result: %v, got: %v", tt.expectData, result)
			}
		})
	}
}

func equalMaps(a, b map[string]interface{}) bool {
	if len(a) != len(b) {
		return false
	}
	for k, v := range a {
		if b[k] != v {
			return false
		}
	}
	return true
}
