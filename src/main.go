package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	accessToken := os.Getenv("GITHUB_ACCESS_TOKEN")

	username := "luiz04nl"

	// URL da API do GitHub para obter informações do usuário
	url := fmt.Sprintf("https://api.github.com/users/%s", username)

	// Cria uma requisição HTTP
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Erro ao criar requisição:", err)
		return
	}

	// Adiciona o cabeçalho de autorização com o token de acesso
	req.Header.Set("Authorization", "token "+accessToken)

	// Realiza a requisição HTTP
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Erro ao enviar requisição:", err)
		return
	}
	defer resp.Body.Close()

	// Lê a resposta da requisição
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Erro ao ler resposta:", err)
		return
	}

	// Exibe os dados do usuário obtidos da API
	fmt.Println("Resposta da API do GitHub:")
	fmt.Println(string(body))
}
