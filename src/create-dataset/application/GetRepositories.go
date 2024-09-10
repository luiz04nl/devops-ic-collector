package application

import (
	"fmt"

	"github.com/luiz04nl/devops-ic-collector/src/shared"
)

func GetRepositories() ([]shared.RepositoryDto, error) {
	var dtos []shared.RepositoryDto
	var minStars = 3224
	var after *string = nil
	var hasNextPage = true

	for hasNextPage {
		var filterAfter string = ""
		if after != nil {
			filterAfter = fmt.Sprintf(`, after: "%s"`, *after)
		}

		// languages := "js,ts,java,python"
		languages := "java"

		query := fmt.Sprintf(`
    {
      search(query: "is:public language:%v NOT android NOT jvm NOT spring NOT platform_frameworks_base NOT hbase stars:>=%d",
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
    `, languages, minStars, filterAfter)

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

		after = &endCursor
	}

	return dtos, nil
}
