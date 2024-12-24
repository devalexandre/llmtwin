package main

import (
	"context"
	"fmt"
	"github.com/devalexandre/llmtwin/agents"
	"github.com/devalexandre/llmtwin/state"
	"github.com/tmc/langchaingo/llms/openai"
	"log"
)

func main() {

	s := state.NewState()
	s.Update("temperature", 75.3)
	s.Update("pressure", 101.5)

	ctx := context.Background()
	llm, err := openai.New(
		openai.WithModel("gpt-4o"),
		openai.WithToken(""),
	)
	if err != nil {
		log.Fatal(err)
	}

	agent := agents.NewAgent()
	agent.RegisterTool("analyze-state", func(s state.State) (string, error) {
		temp, _ := s.Get("temperature")
		return fmt.Sprintf("Temperatura atual é: %v", temp), nil
	})

	toolOutput, err := agent.Execute(*s, "analyze-state")
	if err != nil {
		log.Fatal(err)
	}
	response, err := llm.Call(ctx, fmt.Sprintf("O estado atual é: %s. Recomende uma ação.", toolOutput))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Resposta do LLM:", response)

}
