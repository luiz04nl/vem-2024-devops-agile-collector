package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func executeGraphQLQuery(query string) (*GitHubSearchResponse, error) {
	accessToken := os.Getenv("GITHUB_ACCESS_TOKEN")
	if accessToken == "" {
		return nil, fmt.Errorf("token de acesso pessoal não encontrado. Defina a variável de ambiente GITHUB_ACCESS_TOKEN")
	}

	jsonMapInstance := map[string]string{
		"query": query,
	}

	jsonResult, err := json.Marshal(jsonMapInstance)
	if err != nil {
		return nil, fmt.Errorf("there was an error marshaling the JSON instance %v", err)
	}

	req, err := http.NewRequest("POST", "https://api.github.com/graphql", bytes.NewBuffer(jsonResult))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+accessToken)

	if err != nil {
		return nil, fmt.Errorf("erro ao criar requisição: %v", err)
	}

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("erro ao enviar requisição: %v", err)
	}
	defer resp.Body.Close()

	jsonBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("erro ao ler resposta: %v", err)
	}

	var gitHubSearchResponse GitHubSearchResponse
	gitHubSearchResponseErr := json.Unmarshal(jsonBody, &gitHubSearchResponse)
	if err != nil {
		log.Fatalf("json.Unmarshal falhou: %v", gitHubSearchResponseErr)
	}

	return &gitHubSearchResponse, nil
}
