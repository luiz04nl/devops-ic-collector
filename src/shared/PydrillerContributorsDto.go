package shared

type PydrillerContributorsDto struct {
	RepositoryAlias string `json:"repositoryAlias"`
	RepositoryName  string `json:"repositoryName"`
	Contributors    int    `json:"contributors"`
}
