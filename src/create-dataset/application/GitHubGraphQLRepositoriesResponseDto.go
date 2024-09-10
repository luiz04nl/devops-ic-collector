package application

type GitHubGraphQLRepositoriesResponseDto struct {
	Data struct {
		Search struct {
			RepositoryCount int `json:"repositoryCount"`
			PageInfo        struct {
				EndCursor   string `json:"endCursor"`
				StartCursor string `json:"startCursor"`
				hasNextPage bool   `json:"hasNextPage"`
			} `json:"pageInfo"`
			Edges []struct {
				Node struct {
					Name       string `json:"name"`
					URL        string `json:"url"`
					Stargazers struct {
						TotalCount int `json:"totalCount"`
					} `json:"stargazers"`
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
