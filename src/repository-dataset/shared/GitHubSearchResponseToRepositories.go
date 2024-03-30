package shared

func GitHubSearchResponseToRepositories(response *GitHubSearchResponse) []Repository {
	var Repositorys []Repository

	for _, edge := range response.Data.Search.Edges {
		Repositorys = append(Repositorys, Repository{
			Name:            edge.Node.Name,
			URL:             edge.Node.URL,
			StarsTotalCount: edge.Node.Stargazers.TotalCount,
		})
	}

	return Repositorys
}
