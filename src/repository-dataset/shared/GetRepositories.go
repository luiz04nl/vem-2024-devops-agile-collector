package shared

import (
	"fmt"
)

func GetRepositories() ([]RepositoryDto, error) {
	var dtos []RepositoryDto

	// var minStars = 1612
	var minStars = 3224
	// Me Parece que mesmo paginando nao consigo buscar mais de 1000 registros
	var after *string = nil

	var hasNextPage = true

	for hasNextPage {

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

		if endCursor == "" {
			hasNextPage = false
		}
		// hasNextPage = currentGitHubGraphQLRepositoriesResponseDto.Data.Search.PageInfo.hasNextPage

		after = &endCursor

		// hasNextPage = false
		// fmt.Println("endCursor:", endCursor)
		// fmt.Println("after:", after)
	}

	return dtos, nil
}
