package shared

type ContributorInfo struct {
	Id             string
	Commits        int
	IntervalInDays float32
}

type ContributorsInfo = []ContributorInfo
