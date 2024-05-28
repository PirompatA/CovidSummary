package entity

// SummaryResponse represents the summary response structure
type SummaryResponse struct {
	Province map[string]int `json:"Province"`
	AgeGroup map[string]int `json:"AgeGroup"`
}
