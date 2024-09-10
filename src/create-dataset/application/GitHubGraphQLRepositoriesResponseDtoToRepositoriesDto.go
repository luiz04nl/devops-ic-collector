package application

import (
	"crypto/sha256"
	"encoding/hex"

	"github.com/luiz04nl/devops-ic-collector/src/shared"
)

func generateRepoHash(repoURL string) string {
	repoURLBytes := []byte(repoURL)
	hash := sha256.New()
	hash.Write(repoURLBytes)
	hashBytes := hash.Sum(nil)
	repoHash := hex.EncodeToString(hashBytes)
	return repoHash
}

func GitHubGraphQLRepositoriesResponseDtoToRepositoriesDto(response *GitHubGraphQLRepositoriesResponseDto) shared.RepositoriesDto {
	var Repositories []shared.RepositoryDto
	meta := "{}"

	for _, edge := range response.Data.Search.Edges {
		uniqueHash := generateRepoHash(edge.Node.URL)

		Repositories = append(Repositories, shared.RepositoryDto{
			Name:                         edge.Node.Name,
			URL:                          edge.Node.URL,
			Alias:                        edge.Node.Name + "-" + uniqueHash,
			StarsTotalCount:              edge.Node.Stargazers.TotalCount,
			UseAgile:                     0,
			UseDevOps:                    0,
			UseGithubPipelines:           0,
			UseCircleCI:                  0,
			UseJenkins:                   0,
			UseGitLabPipelines:           0,
			UseAzureDevops:               0,
			UseTravisCI:                  0,
			UseHarness:                   0,
			UseBitBucketPipelines:        0,
			WasCloned:                    0,
			Meta:                         &meta,
			ProjectContributors:          0,
			ProjectCommits:               0,
			CommitsIntervalInDays:        0,
			ContributorsInfo:             "{}",
			HasCommitsInInterval:         false,
			LastCommitDateInterval:       nil,
			ProjectType:                  "",
			ProjectTypeVersion:           "",
			ProjectIssuesEffortTotal:     0,
			ProjectIssuesCount:           0,
			ProjectCodeSmellsEffortTotal: 0,
			ProjectCodeSmellsCount:       0,
			ProjectSonarComponentsCount:  0,
			ProjectSonarInfo:             "{}",
			Bugs:                         "",
			SqaleRating:                  "",
			ReliabilityRating:            "",
			Complexity:                   "",
			CognitiveComplexity:          "",
			DuplicatedBlocks:             "",
			DuplicatedFiles:              "",
			DuplicatedLines:              "",
			CodeSmells:                   "",
			LinesOfCodesFromSonar:        "",
			SqaleIndex:                   "",
			SqaleDebtRatio:               "",
			QualityGateDetails:           "",
			Vulnerabilities:              "",
			SecurityRating:               "",
			Classes:                      "",
			CommentLines:                 "",
			Coverage:                     "",
			Tests:                        "",

			LinesOfCodesFromCk:             "",
			CouplingBetweenObjects:         "",
			CouplingBetweenObjectsModified: "",
		})
	}

	return shared.RepositoriesDto{
		Repositories: Repositories,
		Count:        response.Data.Search.RepositoryCount,
	}
}
