package backend

type Question struct {
	ID       string                 `json:"id"`
	Question string                 `json:"question"`
	X        map[string]interface{} `json:"-"`
}

type QuestionResultFieldResultItem struct {
	ID    int    `json:"id"`
	Label string `json:"label"`
	Score []int  `json:"score"`
}

type QuestionResult struct {
	Results []QuestionResultFieldResultItem `json:"results"`
	X       map[string]interface{}          `json:"-"`
}

type QuestionDesc struct {
	Question string                          `json:"question"`
	Results  []QuestionResultFieldResultItem `json:"results"`
}
