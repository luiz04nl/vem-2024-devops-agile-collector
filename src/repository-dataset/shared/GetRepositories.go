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
	// var after = "Y3Vyc29yOjIwMA=="
	var after *string = nil

	var thereIsTheNextPage = true

	for thereIsTheNextPage {

		// fmt.Println("thereIsTheNextPage:", thereIsTheNextPage)
		// fmt.Println("after:", after)

		var filterAfter string = ""
		if after != nil {
			filterAfter = fmt.Sprintf(`, after: "%s"`, *after)
		}

		query := fmt.Sprintf(`
    {
      search(query: "is:public stars:>=%d language:Java",
        type: REPOSITORY,
        first: 100 %v) {
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
    `, minStars, filterAfter)

		currentGitHubGraphQLRepositoriesResponseDto, err := ExecuteGraphQLQuery(query)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}

		var endCursor = currentGitHubGraphQLRepositoriesResponseDto.Data.Search.PageInfo.EndCursor

		var currentRepositoriesDto = GitHubGraphQLRepositoriesResponseDtoToRepositoriesDto(currentGitHubGraphQLRepositoriesResponseDto)

		dtos = append(dtos, currentRepositoriesDto.Repositories...)

		if endCursor == "null" || endCursor == "" {
			thereIsTheNextPage = false
		}
		after = &endCursor

		// thereIsTheNextPage = false
		// fmt.Println("endCursor:", endCursor)
		// fmt.Println("after:", after)
	}

	return dtos, nil
}
