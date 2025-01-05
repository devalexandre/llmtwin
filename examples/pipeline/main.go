package main

import (
	"fmt"
	"log"

	"github.com/devalexandre/llmtwin/pipeline"
)

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

func main() {

	p := pipeline.NewPipeline()

	p.AddStage(&TemperatureCheckStage{Threshold: 75.0})

	inputData := map[string]interface{}{
		"temperature": 80.5,
		"pressure":    101.3,
	}

	outputData, err := p.Execute(inputData)
	if err != nil {
		log.Fatal("Pipeline execution failed:", err)
	}

	fmt.Println("Pipeline Output:", outputData)
}
