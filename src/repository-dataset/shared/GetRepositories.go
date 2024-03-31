package shared

import (
	"fmt"
)

func GetRepositories() ([]RepositoryDto, error) {
	// func GetRepositories() (*GitHubGraphQLRepositoriesResponseDto, error) {

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

	// return gitHubGraphQLRepositoriesResponseDto, nil

	var repositoriesDto = GitHubGraphQLRepositoriesResponseDtoToRepositoriesDto(gitHubGraphQLRepositoriesResponseDto)

	return repositoriesDto.Repositories, nil
}
