package shared

type RepositoryDtoBuilder struct {
	instance RepositoryDto
}

func (builder RepositoryDtoBuilder) FromRepository(repository RepositoryDto) RepositoryDtoBuilder {
	builder.instance = RepositoryDto{
		Id:                           repository.Id,
		URL:                          repository.URL,
		Name:                         repository.Name,
		Alias:                        repository.Alias,
		StarsTotalCount:              repository.StarsTotalCount,
		UseAgile:                     repository.UseAgile,
		UseDevOps:                    repository.UseDevOps,
		UseGithubPipelines:           repository.UseGithubPipelines,
		UseCircleCI:                  repository.UseCircleCI,
		UseJenkins:                   repository.UseJenkins,
		UseGitLabPipelines:           repository.UseGitLabPipelines,
		UseAzureDevops:               repository.UseAzureDevops,
		UseTravisCI:                  repository.UseTravisCI,
		UseHarness:                   repository.UseHarness,
		UseBitBucketPipelines:        repository.UseBitBucketPipelines,
		WasCloned:                    repository.WasCloned,
		Meta:                         repository.Meta,
		ProjectContributors:          repository.ProjectContributors,
		ProjectCommits:               repository.ProjectCommits,
		CommitsIntervalInDays:        repository.CommitsIntervalInDays,
		ContributorsInfo:             repository.ContributorsInfo,
		HasCommitsInInterval:         repository.HasCommitsInInterval,
		LastCommitDateInterval:       repository.LastCommitDateInterval,
		ProjectType:                  repository.ProjectType,
		ProjectTypeVersion:           repository.ProjectTypeVersion,
		ProjectIssuesEffortTotal:     repository.ProjectIssuesEffortTotal,
		ProjectIssuesCount:           repository.ProjectIssuesCount,
		ProjectCodeSmellsEffortTotal: repository.ProjectCodeSmellsEffortTotal,
		ProjectCodeSmellsCount:       repository.ProjectCodeSmellsCount,
		ProjectSonarComponentsCount:  repository.ProjectSonarComponentsCount,
		ProjectSonarInfo:             repository.ProjectSonarInfo,
		Bugs:                         repository.Bugs,
		SqaleRating:                  repository.SqaleRating,
		ReliabilityRating:            repository.ReliabilityRating,
		Complexity:                   repository.Complexity,
		CognitiveComplexity:          repository.CognitiveComplexity,
		DuplicatedBlocks:             repository.DuplicatedBlocks,
		DuplicatedFiles:              repository.DuplicatedFiles,
		DuplicatedLines:              repository.DuplicatedLines,
		CodeSmells:                   repository.CodeSmells,
		LinesOfCodesFromSonar:        repository.LinesOfCodesFromSonar,
		SqaleIndex:                   repository.SqaleIndex,
		SqaleDebtRatio:               repository.SqaleDebtRatio,
		QualityGateDetails:           repository.QualityGateDetails,
		Vulnerabilities:              repository.Vulnerabilities,
		SecurityRating:               repository.SecurityRating,
		Classes:                      repository.Classes,
		CommentLines:                 repository.CommentLines,
		Coverage:                     repository.Coverage,
		Tests:                        repository.Tests,

		LinesOfCodesFromCk:             repository.LinesOfCodesFromCk,
		CouplingBetweenObjects:         repository.CouplingBetweenObjects,
		CouplingBetweenObjectsModified: repository.CouplingBetweenObjectsModified,
	}
	return builder
}

func (builder RepositoryDtoBuilder) Create(repository RepositoryDto) RepositoryDtoBuilder {
	meta := "{}"
	contributorsInfo := "{}"
	projectSonarInfo := "{}"

	builder.instance = RepositoryDto{
		Id:                           0,
		URL:                          "",
		Name:                         "",
		Alias:                        "",
		StarsTotalCount:              0,
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
		ContributorsInfo:             contributorsInfo,
		HasCommitsInInterval:         false,
		LastCommitDateInterval:       nil,
		ProjectType:                  "",
		ProjectTypeVersion:           "",
		ProjectIssuesEffortTotal:     0,
		ProjectIssuesCount:           0,
		ProjectCodeSmellsEffortTotal: 0,
		ProjectCodeSmellsCount:       0,
		ProjectSonarComponentsCount:  0,
		ProjectSonarInfo:             projectSonarInfo,

		Bugs:                  "",
		SqaleRating:           "",
		ReliabilityRating:     "",
		Complexity:            "",
		CognitiveComplexity:   "",
		DuplicatedBlocks:      "",
		DuplicatedFiles:       "",
		DuplicatedLines:       "",
		CodeSmells:            "",
		LinesOfCodesFromSonar: "",
		SqaleIndex:            "",
		SqaleDebtRatio:        "",
		QualityGateDetails:    "",
		Vulnerabilities:       "",
		SecurityRating:        "",
		Classes:               "",
		CommentLines:          "",
		Coverage:              "",
		Tests:                 "",

		LinesOfCodesFromCk:             "",
		CouplingBetweenObjects:         "",
		CouplingBetweenObjectsModified: "",
	}
	return builder
}

func (builder RepositoryDtoBuilder) WithWasNotCloned() RepositoryDtoBuilder {
	builder.instance.WasCloned = 0
	return builder
}

func (builder RepositoryDtoBuilder) WithWasCloned() RepositoryDtoBuilder {
	builder.instance.WasCloned = 1
	return builder
}

func (builder RepositoryDtoBuilder) WithUseNotDevops() RepositoryDtoBuilder {
	builder.instance.UseDevOps = 0
	return builder
}

func (builder RepositoryDtoBuilder) WithUseDevops() RepositoryDtoBuilder {
	builder.instance.UseDevOps = 1
	return builder
}

func (builder RepositoryDtoBuilder) WithUseGithubPipelines() RepositoryDtoBuilder {
	builder.instance.UseGithubPipelines = 1
	return builder
}

func (builder RepositoryDtoBuilder) WithUseCircleCI() RepositoryDtoBuilder {
	builder.instance.UseCircleCI = 1
	return builder
}

func (builder RepositoryDtoBuilder) WithJenkins() RepositoryDtoBuilder {
	builder.instance.UseJenkins = 1
	return builder
}

func (builder RepositoryDtoBuilder) WithGitLabPipelines() RepositoryDtoBuilder {
	builder.instance.UseGitLabPipelines = 1
	return builder
}

func (builder RepositoryDtoBuilder) WithAzureDevops() RepositoryDtoBuilder {
	builder.instance.UseAzureDevops = 1
	return builder
}

func (builder RepositoryDtoBuilder) WithTravisCI() RepositoryDtoBuilder {
	builder.instance.UseTravisCI = 1
	return builder
}

func (builder RepositoryDtoBuilder) WithHarness() RepositoryDtoBuilder {
	builder.instance.UseHarness = 1
	return builder
}

func (builder RepositoryDtoBuilder) WithBitbucketPipelines() RepositoryDtoBuilder {
	builder.instance.UseBitBucketPipelines = 1
	return builder
}

func (builder RepositoryDtoBuilder) WithUseNotAgile() RepositoryDtoBuilder {
	builder.instance.UseAgile = 0
	return builder
}

func (builder RepositoryDtoBuilder) WithUseAgile() RepositoryDtoBuilder {
	builder.instance.UseAgile = 1
	return builder
}

func (builder RepositoryDtoBuilder) WithProjectContributors(projectContributors int) RepositoryDtoBuilder {
	builder.instance.ProjectContributors = projectContributors
	return builder
}

func (builder RepositoryDtoBuilder) WithProjectCommits(projectCommits int) RepositoryDtoBuilder {
	builder.instance.ProjectCommits = projectCommits
	return builder
}

func (builder RepositoryDtoBuilder) WithCommitsIntervalInDays(commitsIntervalInDays float32) RepositoryDtoBuilder {
	builder.instance.CommitsIntervalInDays = commitsIntervalInDays
	return builder
}

func (builder RepositoryDtoBuilder) WithContributorsInfo(contributorsInfo string) RepositoryDtoBuilder {
	builder.instance.ContributorsInfo = contributorsInfo
	return builder
}

func (builder RepositoryDtoBuilder) WithHasCommitsInInterval(hasCommitsInInterval bool) RepositoryDtoBuilder {
	builder.instance.HasCommitsInInterval = hasCommitsInInterval
	return builder
}

func (builder RepositoryDtoBuilder) WithLastCommitDateInterval(lastCommitDateInterval *string) RepositoryDtoBuilder {
	builder.instance.LastCommitDateInterval = lastCommitDateInterval
	return builder
}

func (builder RepositoryDtoBuilder) WithProjectType(projectType string) RepositoryDtoBuilder {
	builder.instance.ProjectType = projectType
	return builder
}

func (builder RepositoryDtoBuilder) WithProjectTypeVersion(projectTypeVersion string) RepositoryDtoBuilder {
	builder.instance.ProjectTypeVersion = projectTypeVersion
	return builder
}

func (builder RepositoryDtoBuilder) WithProjectIssuesEffortTotal(projectIssuesEffortTotal int) RepositoryDtoBuilder {
	builder.instance.ProjectIssuesEffortTotal = projectIssuesEffortTotal
	return builder
}

func (builder RepositoryDtoBuilder) WithProjectIssuesCount(projectIssuesCount int) RepositoryDtoBuilder {
	builder.instance.ProjectIssuesCount = projectIssuesCount
	return builder
}

func (builder RepositoryDtoBuilder) WithProjectCodeSmellsEffortTotal(projectCodeSmellsEffortTotal int) RepositoryDtoBuilder {
	builder.instance.ProjectCodeSmellsEffortTotal = projectCodeSmellsEffortTotal
	return builder
}

func (builder RepositoryDtoBuilder) WithProjectCodeSmellsCount(projectCodeSmellsCount int) RepositoryDtoBuilder {
	builder.instance.ProjectCodeSmellsCount = projectCodeSmellsCount
	return builder
}

func (builder RepositoryDtoBuilder) WithProjectSonarComponentsCount(projectSonarComponentsCount int) RepositoryDtoBuilder {
	builder.instance.ProjectSonarComponentsCount = projectSonarComponentsCount
	return builder
}

func (builder RepositoryDtoBuilder) WithProjectSonarInfo(projectSonarInfo string) RepositoryDtoBuilder {
	builder.instance.ProjectSonarInfo = projectSonarInfo
	return builder
}

func (builder RepositoryDtoBuilder) WithBugs(bugs string) RepositoryDtoBuilder {
	builder.instance.Bugs = bugs
	return builder
}

func (builder RepositoryDtoBuilder) WithSqaleRating(sqaleRating string) RepositoryDtoBuilder {
	builder.instance.SqaleRating = sqaleRating
	return builder
}

func (builder RepositoryDtoBuilder) WithReliabilityRating(reliabilityRating string) RepositoryDtoBuilder {
	builder.instance.ReliabilityRating = reliabilityRating
	return builder
}

func (builder RepositoryDtoBuilder) WithComplexity(complexity string) RepositoryDtoBuilder {
	builder.instance.Complexity = complexity
	return builder
}

func (builder RepositoryDtoBuilder) WithCognitiveComplexity(cognitiveComplexity string) RepositoryDtoBuilder {
	builder.instance.CognitiveComplexity = cognitiveComplexity
	return builder
}

func (builder RepositoryDtoBuilder) WithDuplicatedBlocks(duplicatedBlocks string) RepositoryDtoBuilder {
	builder.instance.DuplicatedBlocks = duplicatedBlocks
	return builder
}

func (builder RepositoryDtoBuilder) WithDuplicatedFiles(duplicatedFiles string) RepositoryDtoBuilder {
	builder.instance.DuplicatedFiles = duplicatedFiles
	return builder
}

func (builder RepositoryDtoBuilder) WithDuplicatedLines(duplicatedLines string) RepositoryDtoBuilder {
	builder.instance.DuplicatedLines = duplicatedLines
	return builder
}

func (builder RepositoryDtoBuilder) WithCodeSmells(codeSmells string) RepositoryDtoBuilder {
	builder.instance.CodeSmells = codeSmells
	return builder
}

func (builder RepositoryDtoBuilder) WithLinesOfCodesFromSonar(linesOfCodesFromSonar string) RepositoryDtoBuilder {
	builder.instance.LinesOfCodesFromSonar = linesOfCodesFromSonar
	return builder
}

func (builder RepositoryDtoBuilder) WithSqaleIndex(sqaleIndex string) RepositoryDtoBuilder {
	builder.instance.SqaleIndex = sqaleIndex
	return builder
}

func (builder RepositoryDtoBuilder) WithSqaleDebtRatio(sqaleDebtRatio string) RepositoryDtoBuilder {
	builder.instance.SqaleDebtRatio = sqaleDebtRatio
	return builder
}

func (builder RepositoryDtoBuilder) WithQualityGateDetails(qualityGateDetails string) RepositoryDtoBuilder {
	builder.instance.QualityGateDetails = qualityGateDetails
	return builder
}

func (builder RepositoryDtoBuilder) WithVulnerabilities(vulnerabilities string) RepositoryDtoBuilder {
	builder.instance.Vulnerabilities = vulnerabilities
	return builder
}

func (builder RepositoryDtoBuilder) WithSecurityRating(securityRating string) RepositoryDtoBuilder {
	builder.instance.SecurityRating = securityRating
	return builder
}

func (builder RepositoryDtoBuilder) WithClasses(classes string) RepositoryDtoBuilder {
	builder.instance.Classes = classes
	return builder
}

func (builder RepositoryDtoBuilder) WithCommentLines(commentLines string) RepositoryDtoBuilder {
	builder.instance.CommentLines = commentLines
	return builder
}

func (builder RepositoryDtoBuilder) WithCoverage(coverage string) RepositoryDtoBuilder {
	builder.instance.Coverage = coverage
	return builder
}

func (builder RepositoryDtoBuilder) WithTests(tests string) RepositoryDtoBuilder {
	builder.instance.Tests = tests
	return builder
}

func (builder RepositoryDtoBuilder) WithLinesOfCodesFromCk(linesOfCodesFromCk string) RepositoryDtoBuilder {
	builder.instance.LinesOfCodesFromCk = linesOfCodesFromCk
	return builder
}

func (builder RepositoryDtoBuilder) WithCouplingBetweenObjects(couplingBetweenObjects string) RepositoryDtoBuilder {
	builder.instance.CouplingBetweenObjects = couplingBetweenObjects
	return builder
}

func (builder RepositoryDtoBuilder) WithCouplingBetweenObjectsModified(couplingBetweenObjectsModified string) RepositoryDtoBuilder {
	builder.instance.CouplingBetweenObjectsModified = couplingBetweenObjectsModified
	return builder
}

func (builder RepositoryDtoBuilder) Build() RepositoryDto {
	return builder.instance
}
