
# LLM Twin

`llmtwin` is a powerful framework that integrates Large Language Models (LLMs) with real-time state management and tool execution. It is designed to help developers create intelligent systems that monitor, analyze, and react to changes in data or states using the capabilities of LLMs.

---

## Key Features

- **State Management:** Manage dynamic states with tools to update and retrieve values efficiently.
- **Agent and Tool System:** Create agents that register tools for specific functionalities, enabling modular and extensible workflows.
- **LLM Integration:** Direct integration with [LangChainGo](https://github.com/tmc/langchaingo) to leverage OpenAI's models 
- **Event-Driven Design:** Ideal for monitoring and reacting to real-time changes, such as database updates or IoT sensor inputs.
- **Modularity:** A clean separation of concerns between state management, tool registration, and LLM execution.

---

## Installation

To get started, clone the repository and install the required dependencies:

```bash
go get github.com/devalexandre/llmtwin
```

---

## Getting Started

Here’s an example of how to use `llmtwin` to manage a dynamic state, register tools, and interact with an LLM.

### 1. Initialize State Management

Create a new state and update it with dynamic values:

```go
import "github.com/devalexandre/llmtwin/state"

s := state.NewState()
s.Update("temperature", 75.3)
s.Update("pressure", 101.5)
```

### 2. Register Tools with an Agent

Tools are functions that perform specific tasks. Register tools to an agent and execute them based on the state:

```go
import "github.com/devalexandre/llmtwin/agents"

agent := agents.NewAgent()
agent.RegisterTool("analyze-state", func(s state.State) (string, error) {
    temp, _ := s.Get("temperature")
    return fmt.Sprintf("Current temperature is: %.1f", temp), nil
})
```

Execute the tool:

```go
toolOutput, err := agent.Execute(*s, "analyze-state")
if err != nil {
    log.Fatal("Error executing tool:", err)
}
fmt.Println("Tool Output:", toolOutput)
```

### 3. Integrate with an LLM

Use LangChainGo to integrate an LLM and generate recommendations or insights:

```go
import (
    "context"
    "github.com/tmc/langchaingo/llms/openai"
)

ctx := context.Background()
llm, err := openai.New(
    openai.WithModel("gpt-4"),
    openai.WithToken("your-openai-api-key"),
)
if err != nil {
    log.Fatal("Error initializing LLM:", err)
}

prompt := fmt.Sprintf("The current state is: %s. Please recommend an action.", toolOutput)
response, err := llm.Call(ctx, prompt)
if err != nil {
    log.Fatal("Error calling LLM:", err)
}

fmt.Println("LLM Response:", response)
```

---

## Example Use Case

### Monitoring an E-Commerce System

**Scenario:** Monitor new sales in a database and use an LLM to generate recommendations for marketing campaigns.

1. Use `state` to track the latest sales data.
2. Register tools for database queries.
3. Send sales data to the LLM for analysis.

```go
agent.RegisterTool("query-sales", func(s state.State) (string, error) {
    salesData := "[{"product":"Shirt","amount":29.99}]"
    return salesData, nil
})

salesOutput, _ := agent.Execute(*s, "query-sales")
prompt := fmt.Sprintf("Recent sales data: %s. Recommend a marketing strategy.", salesOutput)
response, _ := llm.Call(ctx, prompt)
fmt.Println("Marketing Strategy:", response)
```

---

## Advanced Features

- **Event-Driven Execution:** Use systems like NATS or RabbitMQ to trigger tools and LLM interactions based on real-time events.
- **Database Integration:** Connect to databases and dynamically query data with registered tools.
- **Extensibility:** Easily add new tools and integrate other LLM providers.

---

## Contributing

Contributions are welcome! Please submit a pull request or open an issue for feedback and suggestions.

---

## License

This project is licensed under the MIT License. See the LICENSE file for details.

---

## Support

For support or inquiries, please contact [devalexandre](https://github.com/devalexandre) or open an issue in the repository.
