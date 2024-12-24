package main

import (
	"context"
	"fmt"
	"log"

	"github.com/devalexandre/llmtwin/pipeline"
	"github.com/devalexandre/llmtwin/state"
	"github.com/tmc/langchaingo/llms/openai"
)

// Define a stage that updates the state based on a threshold
type TemperatureCheckStage struct {
	Threshold float64
}

func (t *TemperatureCheckStage) ShouldExecute(data map[string]interface{}) bool {
	temp, exists := data["temperature"].(float64)
	return exists && temp > t.Threshold
}

func (t *TemperatureCheckStage) Process(data map[string]interface{}) (map[string]interface{}, error) {
	data["temperature_warning"] = "High temperature detected!"
	return data, nil
}

// Define a stage that prepares a prompt and updates the state
type PreparePromptStage struct{}

func (p *PreparePromptStage) ShouldExecute(data map[string]interface{}) bool {
	return true
}

func (p *PreparePromptStage) Process(data map[string]interface{}) (map[string]interface{}, error) {
	// Create a prompt using the data and update the state
	prompt := fmt.Sprintf("Temperature: %.1f, Pressure: %.1f. Analyze the data and recommend actions.",
		data["temperature"].(float64), data["pressure"].(float64))
	data["prompt"] = prompt
	return data, nil
}

func main() {
	// Initialize the state
	s := state.NewState()
	s.Update("temperature", 80.5)
	s.Update("pressure", 101.5)

	// Create a pipeline
	p := pipeline.NewPipeline()
	p.AddStage(&TemperatureCheckStage{Threshold: 75.0})
	p.AddStage(&PreparePromptStage{})

	// Convert state to input data for the pipeline
	inputData := make(map[string]interface{})
	inputData["temperature"], _ = s.Get("temperature")
	inputData["pressure"], _ = s.Get("pressure")

	// Execute the pipeline
	outputData, err := p.Execute(inputData)
	if err != nil {
		log.Fatal("Pipeline execution failed:", err)
	}

	// Update the state with pipeline results
	if warning, exists := outputData["temperature_warning"]; exists {
		s.Update("warning", warning)
	}
	if prompt, exists := outputData["prompt"]; exists {
		s.Update("prompt", prompt)
	}

	// Retrieve the final prompt from the state
	finalPrompt, _ := s.Get("prompt")

	// Initialize the LLM
	ctx := context.Background()
	llm, err := openai.New(
		openai.WithModel("gpt-4"),
		openai.WithToken("your-openai-api-key"),
	)
	if err != nil {
		log.Fatal("Error initializing LLM:", err)
	}

	// Call the LLM with the prompt
	response, err := llm.Call(ctx, finalPrompt.(string))
	if err != nil {
		log.Fatal("Error calling LLM:", err)
	}

	fmt.Println("LLM Response:", response)
}
