package main

type GitHubSearchResponse struct {
	Data struct {
		Search struct {
			RepositoryCount int `json:"repositoryCount"`
			PageInfo        struct {
				EndCursor   string `json:"endCursor"`
				StartCursor string `json:"startCursor"`
			} `json:"pageInfo"`
			Edges []struct {
				Node struct {
					Name       string `json:"name"`
					URL        string `json:"url"`
					Stargazers struct {
						TotalCount int `json:"totalCount"`
					} `json:"stargazers"`
					// Collaborators is omitted due to permission issues
				} `json:"node"`
			} `json:"edges"`
		} `json:"search"`
	} `json:"data"`
	Errors []struct {
		Type      string        `json:"type"`
		Path      []interface{} `json:"path"`
		Locations []struct {
			Line   int `json:"line"`
			Column int `json:"column"`
		} `json:"locations"`
		Message string `json:"message"`
	} `json:"errors"`
}

type Repositorie struct {
	Name            string
	URL             string
	StarsTotalCount int
}
