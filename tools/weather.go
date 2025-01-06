package tools

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/devalexandre/llmtwin/state"
)

// WeatherTool busca informações meteorológicas pelo nome da cidade.
type WeatherTool struct {
	APIKey string
}

func (w *WeatherTool) FetchWeather(city string) (string, error) {
	// Codifica o nome da cidade para evitar problemas com caracteres especiais.
	encodedCity := url.QueryEscape(city)
	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&APPID=%s&units=metric", encodedCity, w.APIKey)

	// Faz a requisição HTTP.
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("falha ao buscar dados da API: %s", resp.Status)
	}

	// Lê a resposta e converte para JSON.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return "", err
	}

	// Extrai as informações desejadas.
	main, ok := result["main"].(map[string]interface{})
	if !ok {
		return "", errors.New("dados inesperados no formato da resposta")
	}

	temp := main["temp"].(float64)
	weather := result["weather"].([]interface{})[0].(map[string]interface{})
	description := weather["description"].(string)

	return fmt.Sprintf("A temperatura atual em %s é %.1f°C com %s.", city, temp, description), nil
}

// Execute é o método da Tool para integração com o agente.
func (w *WeatherTool) Execute(s state.State) (string, error) {
	city, ok := s.Data["city"].(string)
	if !ok {
		return "", errors.New("cidade não fornecida no estado")
	}

	return w.FetchWeather(city)
}
