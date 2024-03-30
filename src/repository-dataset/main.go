package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/luiz04nl/devops-ic-collector/src/repository-dataset/shared"

	_ "github.com/mattn/go-sqlite3"
)

//after utiliza
// "pageInfo": {
//   "endCursor": "Y3Vyc29yOjIwMA==",
//   "startCursor": "Y3Vyc29yOjEwMQ=="
// },

func main() {
	query := `
  {
    search(query: "is:public stars:>=3224 language:Java",
      type: REPOSITORY,
      first: 100, , after: null) {
      repositoryCount
      pageInfo {
        endCursor
        startCursor
      }
      edges {
        node {
          ... on Repository {
            name
            url
            stargazers {
              totalCount
            }
          }
        }
      }
    }
  }
  `

	gitHubSearchResponse, err := shared.ExecuteGraphQLQuery(query)
	if err != nil {
		fmt.Println(err)
		return
	}

	// log.Println("gitHubSearchResponse: ", gitHubSearchResponse)

	gitHubSearchResponseJsonData, err := json.Marshal(&gitHubSearchResponse)
	if err != nil {
		log.Fatalf("Erro ao converter para JSON: %v", err)
	}

	log.Println("gitHubSearchResponseJsonData: ", string(gitHubSearchResponseJsonData))

	// var repositories = shared.GitHubSearchResponseToRepositories(gitHubSearchResponse)

	// log.Println("repositories: ", repositories)
}
