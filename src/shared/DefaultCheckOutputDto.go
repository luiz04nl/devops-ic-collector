package shared

type DefaultCheckOutputDto struct {
	ProjectType        string `json:"projectType"`
	ProjectTypeVersion string `json:"projectTypeVersion"`
	Repository         string `json:"repository"`
	AnalysisSuccess    string `json:"analysisSuccess"`
  // FilesAtRootDir    string `json:"filesAtRootDir"`
}
