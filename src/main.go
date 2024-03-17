package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	accessToken := os.Getenv("GITHUB_ACCESS_TOKEN")
	if accessToken == "" {
		fmt.Println("Token de acesso pessoal não encontrado. Defina a variável de ambiente GITHUB_ACCESS_TOKEN.")
		return
	}

	// jsonMapInstance := map[string]string{
	// 	"query": `
	//       {
	//         user(login: "luiz04nl") {
	//           id
	//           name
	//         }
	//       }
	//   `,
	// }

	// jsonMapInstance := map[string]string{
	// 	"query": `
	//       {
	//             viewer {
	//                 login
	//               }
	//       }
	//   `,
	// }

	jsonMapInstance := map[string]string{
		"query": `
    {
      search(query: "is:public", type: REPOSITORY, first: 3) {
        repositoryCount
        pageInfo {
          endCursor
          startCursor
        }
        edges {
          node {
            ... on Repository {
              name,
              id
	              nameWithOwner
	              description
	              url
	              collaborators(first: 3, after: null) {
	                  edges {
	                      permission
	                      node {
	                          id
	                          name
	                          email
	                      }
	                  }
	                  pageInfo {
	                      hasNextPage
	                      endCursor
	                  }
	              }

            }
          }
        }
      }
}
    `,
	}

	jsonResult, err := json.Marshal(jsonMapInstance)

	if err != nil {
		fmt.Printf("There was an error marshaling the JSON instance %v", err)
	}

	req, err := http.NewRequest("POST", "https://api.github.com/graphql", bytes.NewBuffer(jsonResult))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+accessToken)

	if err != nil {
		fmt.Println("Erro ao criar requisição:", err)
		return
	}

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

	// Exibe os dados do usuário obtidos da Graph API do GitHub
	fmt.Println("Resposta da Graph API do GitHub:")
	fmt.Println(string(body))
}
