package shared

type SonarMeasuresDto struct {
	Metric    string `json:"metric"`
	Value     string `json:"value"`
	BestValue bool   `json:"bestValue"`
}

type SonarMeasuresComponentOutputDto struct {
	Key       string             `json:"key"`
	Name      string             `json:"name"`
	Qualifier string             `json:"qualifier"`
	Measures  []SonarMeasuresDto `json:"measures"`
}

type SonarMeasuresOutputDto struct {
	Component SonarMeasuresComponentOutputDto `json:"component"`
}
