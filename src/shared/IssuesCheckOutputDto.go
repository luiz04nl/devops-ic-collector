package shared

type IssuesCheckOutputDto struct {
	Total int `json:"total"`

	Page      int `json:"p"`
	PageLimit int `json:"ps"`

	EffortTotal int   `json:"effortTotal"`
	Issues      []any `json:"issues"`
	Components  []any `json:"components"`
}
