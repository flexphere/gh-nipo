package pr

import (
	"time"
)

type Author struct {
	Login string `json:"login"`
}

type PullRequest struct {
	Author    Author     `json:"author"`
	Title     string     `json:"title"`
	URL       string     `json:"url"`
	UpdatedAt *time.Time `json:"updatedAt"`
	Reviews   []struct {
		ID          string    `json:"id"`
		Author      Author    `json:"author"`
		SubmittedAt time.Time `json:"submittedAt"`
	} `json:"reviews"`
}
