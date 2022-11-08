package pr

import (
	"time"
)

type PullRequest struct {
	Title     string     `json:"title"`
	URL       string     `json:"url"`
	UpdatedAt *time.Time `json:"updatedAt"`
}

func (pr PullRequest) IsRecent() bool {
	now := time.Now()
	return now.Sub(*pr.UpdatedAt) < 24*time.Hour
}
