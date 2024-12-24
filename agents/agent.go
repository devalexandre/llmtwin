package agents

import (
	"fmt"
	"github.com/devalexandre/llmtwin/state"
)

type Agent struct {
	tools map[string]func(state.State) (string, error)
}

func NewAgent() *Agent {
	return &Agent{tools: make(map[string]func(state.State) (string, error))}
}

func (a *Agent) RegisterTool(name string, tool func(state.State) (string, error)) {
	a.tools[name] = tool
}

func (a *Agent) Execute(state state.State, toolName string) (string, error) {
	tool, exists := a.tools[toolName]
	if !exists {
		return "", fmt.Errorf("ferramenta não encontrada: %s", toolName)
	}
	return tool(state)
}
