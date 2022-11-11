package pr

import (
	"time"

	"github.com/AlecAivazis/survey/v2"
)

type PullRequests []PullRequest

func (prs PullRequests) FilterByHours(hours int) PullRequests {
	now := time.Now()
	recents := PullRequests{}
	for _, pr := range prs {
		if now.Sub(*pr.UpdatedAt) < time.Duration(hours)*time.Hour {
			recents = append(recents, pr)
		}
	}
	return recents
}

func (prs PullRequests) GetMyPRs(u string) PullRequests {
	myPRs := PullRequests{}
	for _, pr := range prs {
		if pr.Author.Login == u {
			myPRs = append(myPRs, pr)
		}
	}
	return myPRs
}

func (prs PullRequests) GetMyReviews(u string) PullRequests {
	reviews := PullRequests{}
	for _, pr := range prs {
		for _, review := range pr.Reviews {
			if review.Author.Login == u {
				reviews = append(reviews, pr)
			}
		}
	}
	return reviews
}

func (prs PullRequests) Select(label string) PullRequests {
	res := []string{}
	opts := []string{}
	for _, pr := range prs {
		opts = append(opts, pr.Title)
	}

	prompt := &survey.MultiSelect{
		Message: label,
		Options: opts,
	}
	survey.AskOne(prompt, &res)

	results := PullRequests{}
	for _, title := range res {
		for _, pr := range prs {
			if title == pr.Title {
				results = append(results, pr)
			}
		}
	}

	return results
}
