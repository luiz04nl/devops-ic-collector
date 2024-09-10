package shared

type RepositoryDto struct {
	Id                           int     `db:"id"`
	URL                          string  `db:"url"`
	Name                         string  `db:"name"`
	Alias                        string  `db:"alias"`
	StarsTotalCount              int     `db:"starsTotalCount"`
	UseAgile                     int     `db:"useAgile"`
	UseDevOps                    int     `db:"useDevOps"`
	UseGithubPipelines           int     `db:"useGithubPipelines"`
	UseCircleCI                  int     `db:"useCircleCI"`
	UseJenkins                   int     `db:"useJenkins"`
	UseGitLabPipelines           int     `db:"useGitLabPipelines"`
	UseAzureDevops               int     `db:"useAzureDevops"`
	UseTravisCI                  int     `db:"useTravisCI"`
	UseHarness                   int     `db:"useHarness"`
	UseBitBucketPipelines        int     `db:"useBitBucketPipelines"`
	WasCloned                    int     `db:"wasCloned"`
	Meta                         *string `db:"meta"`
	ProjectContributors          int     `db:"projectContributors"`
	ProjectCommits               int     `db:"projectCommits"`
	CommitsIntervalInDays        float32 `db:"commitsIntervalInDays"`
	ContributorsInfo             string  `db:"contributorsInfo"`
	HasCommitsInInterval         bool    `db:"hasCommitsInInterval"`
	LastCommitDateInterval       *string `db:"lastCommitDateInterval"`
	ProjectType                  string  `db:"projectType"`
	ProjectTypeVersion           string  `db:"projectTypeVersion"`
	ProjectIssuesEffortTotal     int     `db:"projectIssuesEffortTotal"`
	ProjectIssuesCount           int     `db:"projectIssuesCount"`
	ProjectCodeSmellsEffortTotal int     `db:"projectCodeSmellsEffortTotal"`
	ProjectCodeSmellsCount       int     `db:"projectCodeSmellsCount"`
	ProjectSonarComponentsCount  int     `db:"projectSonarComponentsCount"`
	ProjectSonarInfo             string  `db:"projectSonarInfo"`
	Bugs                         string  `db:"bugs"`
	SqaleRating                  string  `db:"sqaleRating"`
	ReliabilityRating            string  `db:"reliabilityRating"`
	Complexity                   string  `db:"complexity"`
	CognitiveComplexity          string  `db:"cognitiveComplexity"`
	DuplicatedBlocks             string  `db:"duplicatedBlocks"`
	DuplicatedFiles              string  `db:"duplicatedFiles"`
	DuplicatedLines              string  `db:"duplicatedLines"`
	CodeSmells                   string  `db:"codeSmells"`
	LinesOfCodesFromSonar        string  `db:"linesOfCodesFromSonar"`
	SqaleIndex                   string  `db:"sqaleIndex"`
	SqaleDebtRatio               string  `db:"sqaleDebtRatio"`
	QualityGateDetails           string  `db:"qualityGateDetails"`
	Vulnerabilities              string  `db:"vulnerabilities"`
	SecurityRating               string  `db:"securityRating"`
	Classes                      string  `db:"classes"`
	CommentLines                 string  `db:"commentLines"`
	Coverage                     string  `db:"coverage"`
	Tests                        string  `db:"tests"`

	LinesOfCodesFromCk             string `db:"linesOfCodesFromCk"`
	CouplingBetweenObjects         string `db:"couplingBetweenObjects"`
	CouplingBetweenObjectsModified string `db:"couplingBetweenObjectsModified"`
}
