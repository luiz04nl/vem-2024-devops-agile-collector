package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// GitHubSearchResponse
func executeGraphQLQuery(query string) ([]byte, error) {
	accessToken := os.Getenv("GITHUB_ACCESS_TOKEN")
	if accessToken == "" {
		return nil, fmt.Errorf("Token de acesso pessoal não encontrado. Defina a variável de ambiente GITHUB_ACCESS_TOKEN.")
	}

	jsonMapInstance := map[string]string{
		"query": query,
	}

	jsonResult, err := json.Marshal(jsonMapInstance)
	if err != nil {
		return nil, fmt.Errorf("There was an error marshaling the JSON instance %v", err)
	}

	req, err := http.NewRequest("POST", "https://api.github.com/graphql", bytes.NewBuffer(jsonResult))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+accessToken)

	if err != nil {
		return nil, fmt.Errorf("Erro ao criar requisição: %v", err)
	}

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Erro ao enviar requisição: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Erro ao ler resposta: %v", err)
	}

	return body, nil
}
