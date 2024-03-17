package main

func gitHubSearchResponseToRepositories(response GitHubSearchResponse) []Repositorie {
	var repositories []Repositorie

	for _, edge := range response.Data.Search.Edges {
		repositories = append(repositories, Repositorie{
			Name:            edge.Node.Name,
			URL:             edge.Node.URL,
			StarsTotalCount: edge.Node.Stargazers.TotalCount,
		})
	}

	return repositories
}
