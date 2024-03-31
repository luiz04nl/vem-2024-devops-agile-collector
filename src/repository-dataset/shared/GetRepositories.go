package shared

import (
	"fmt"
)

func GetRepositories() ([]RepositoryDto, error) {
	var dtos []RepositoryDto

	//after utiliza
	// "pageInfo": {
	//   "endCursor": "Y3Vyc29yOjIwMA==",
	//   "startCursor": "Y3Vyc29yOjEwMQ=="
	// },

	var minStars = 1612
	var after = "null"

	var thereIsTheNextPage = true

	for thereIsTheNextPage {

		fmt.Println("thereIsTheNextPage:", thereIsTheNextPage)
		fmt.Println("after:", after)

		query := fmt.Sprintf(`
    {
      search(query: "is:public stars:>=%d language:Java",
        type: REPOSITORY,
        first: 100, after: "%s") {
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
    `, minStars, after)

		currentGitHubGraphQLRepositoriesResponseDto, err := ExecuteGraphQLQuery(query)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}

		var currentRepositoriesDto = GitHubGraphQLRepositoriesResponseDtoToRepositoriesDto(currentGitHubGraphQLRepositoriesResponseDto)

		dtos = append(dtos, currentRepositoriesDto.Repositories...)

		var endCursor = currentGitHubGraphQLRepositoriesResponseDto.Data.Search.PageInfo.EndCursor

		if endCursor == "null" {
			thereIsTheNextPage = false
		}
		after = endCursor

		// thereIsTheNextPage = false
		fmt.Println("endCursor:", endCursor)
		fmt.Println("after:", after)
	}

	return dtos, nil
}
