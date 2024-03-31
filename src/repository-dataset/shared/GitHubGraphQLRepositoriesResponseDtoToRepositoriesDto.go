package shared

func GitHubGraphQLRepositoriesResponseDtoToRepositoriesDto(response *GitHubGraphQLRepositoriesResponseDto) RepositoriesDto {
	var Repositories []RepositoryDto

	for _, edge := range response.Data.Search.Edges {
		Repositories = append(Repositories, RepositoryDto{
			Name:            edge.Node.Name,
			URL:             edge.Node.URL,
			StarsTotalCount: edge.Node.Stargazers.TotalCount,
		})
	}

	return RepositoriesDto{
		Repositories: Repositories,
		Count:        response.Data.Search.RepositoryCount,
	}
}
