package shared

type PydrillerInfosDto struct {
	RepositoryName         string  `json:"repositoryName"`
	Contributors           int     `json:"contributors"`
	Commits                int     `json:"commits"`
	CommitsIntervalInDays  float32 `json:"commitsIntervalInDays"`
	ContributorsInfo       any     `json:"contributorsInfo"`
	HasCommitsInInterval   bool    `json:"hasCommitsInInterval"`
	LastCommitDateInterval *string `json:"lastCommitDateInterval"`
}
