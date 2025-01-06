package main

import (
	"context"
	"fmt"
	"log"

	"github.com/devalexandre/llmtwin/agents"
	"github.com/devalexandre/llmtwin/state"
	"github.com/devalexandre/llmtwin/tools"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"
)

func main() {
	// Inicializa o agente e a ferramenta.
	agent := agents.NewAgent()
	weatherTool := &tools.WeatherTool{APIKey: "API_KEY"}

	// Registra a ferramenta no agente.
	agent.RegisterTool("weather", weatherTool.Execute)

	// Configura o estado inicial com o nome da cidade.
	s := state.State{
		Data: map[string]interface{}{
			"city": "Pocos de Caldas, BR",
		},
	}

	// Executa a ferramenta através do agente.
	result, err := agent.Execute(s, "weather")
	if err != nil {
		fmt.Printf("Erro ao executar a ferramenta: %v\n", err)
		return
	}

	// Exibe o resultado.
	fmt.Printf("Resultado da ferramenta: %s\n", result)

	llm, err := ollama.New(ollama.WithModel("splitpierre/bode-alpaca-pt-br"))
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()

	content := []llms.MessageContent{
		llms.TextParts(llms.ChatMessageTypeSystem, result),
		llms.TextParts(llms.ChatMessageTypeHuman, "Dada a temperatura atual de poços de caldas, oque posso vestir hoje?"),
	}

	completion, err := llm.GenerateContent(ctx, content)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(completion.Choices[0].Content)
}
