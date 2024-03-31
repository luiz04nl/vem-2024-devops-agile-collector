package shared

import (
	"fmt"
)

// func GetRepositories() ([]RepositoryDto, error) {
func GetRepositories() (*GitHubGraphQLRepositoriesResponseDto, error) {

	//after utiliza
	// "pageInfo": {
	//   "endCursor": "Y3Vyc29yOjIwMA==",
	//   "startCursor": "Y3Vyc29yOjEwMQ=="
	// },

	query := `
    {
      search(query: "is:public stars:>=1612 language:Java",
        type: REPOSITORY,
        first: 100, after: null) {
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

	gitHubGraphQLRepositoriesResponseDto, err := ExecuteGraphQLQuery(query)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return gitHubGraphQLRepositoriesResponseDto, nil

	// log.Println("GitHubGraphQLRepositoriesResponseDto: ", GitHubGraphQLRepositoriesResponseDto)

	// gitHubGraphQLRepositoriesResponseDtoJsonData, err := json.Marshal(&gitHubGraphQLRepositoriesResponseDto)
	// if err != nil {
	// 	log.Fatalf("Erro ao converter para JSON: %v", err)
	// }

	// log.Println("gitHubGraphQLRepositoriesResponseDtoJsonData: ", string(gitHubGraphQLRepositoriesResponseDtoJsonData))

	// var repositories = shared.GitHubGraphQLRepositoriesResponseDtoToRepositoriesDto(GitHubGraphQLRepositoriesResponseDto)

	// log.Println("repositories: ", repositories)
}
