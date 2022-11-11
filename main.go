package main

import (
	"fmt"

	"github.com/flexphere/gh-nipo/pr"
)

func main() {
	prs := pr.GetPRs()
	myPRs := prs.GetMyPRs("flexphere").FilterByHours(24).Select("Select PRs")
	reviewPRs := prs.GetMyReviews("flexphere").FilterByHours(24).Select("Select Reviews")

	if len(myPRs)+len(reviewPRs) == 0 {
		return
	}

	fmt.Println(":クラッカー:")
	for _, pr := range append(myPRs, reviewPRs...) {
		fmt.Printf("[%s](%s)\n", pr.Title, pr.URL)
	}
}
