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

	// Historico de commits e pull requets mais recentos em repositorio
	jsonMapInstance := map[string]string{
		"query": `
	  {
	    repository(owner: "luiz04nl", name: "devops-ic-collector") {
	      defaultBranchRef {
	        target {
	          ... on Commit {
	            history(first: 10) {
	              edges {
	                node {
                    committedDate
	                  oid
	                  messageHeadline
	                  author {
	                    name
	                    date
	                  }
	                }
	              }
	            }
	          }
	        }
	      }
	      pullRequests(last: 10, orderBy: {field: CREATED_AT, direction: DESC}) {
	        edges {
	          node {
	            title
	            state
	            author {
	              login
	            }
	            createdAt
	          }
	        }
	      }
	    }
	  }
	`,
	}

	// // Obter repositorios com filtros
	// jsonMapInstance := map[string]string{
	// 	"query": `
	//   {
	//   search(query: "is:public stars:>=10", type: REPOSITORY, first: 3) {
	//     repositoryCount
	//     pageInfo {
	//       endCursor
	//       startCursor
	//     }
	//     edges {
	//       node {
	//         ... on Repository {
	//           name
	//           id
	//           nameWithOwner
	//           description
	//           url
	//           stargazers {
	//             totalCount
	//           }
	//           collaborators(first: 3, after: null) {
	//             totalCount
	//             edges {
	//               permission
	//               node {
	//                 id
	//                 name
	//                 email
	//               }
	//             }
	//             pageInfo {
	//               hasNextPage
	//               endCursor
	//             }
	//           }
	//         }
	//       }
	//     }
	//   }
	// }
	// `,
	// }

	//

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
