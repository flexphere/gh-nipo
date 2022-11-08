package pr

import "github.com/AlecAivazis/survey/v2"

type PullRequests []PullRequest

func (prs PullRequests) GetRecents() PullRequests {
	recents := PullRequests{}
	for _, pr := range prs {
		if pr.IsRecent() {
			recents = append(recents, pr)
		}
	}
	return recents
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
