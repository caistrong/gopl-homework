package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/caistrong/gopl-homework/src/chapter4/work4_10/github"
)

// run command
// go run ./src/chapter4/work4_10/work4_10.go 知识图谱 简单智能问答
func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("total count %d issues: \n", result.TotalCount)
	now := time.Now()
	aMonthAgo := now.AddDate(0, -1, 0)
	aYearAgo := now.AddDate(-1, 0, 0)

	lessMonthIssues := []*github.Issue{}
	lessYearMoreThanMonthIssues := []*github.Issue{}
	moreThanYearIssues := []*github.Issue{}
	for _, item := range result.Items {
		if item.CreatedAt.After(aMonthAgo) {
			lessMonthIssues = append(lessMonthIssues, item)
		} else if item.CreatedAt.After(aYearAgo) {
			lessYearMoreThanMonthIssues = append(lessYearMoreThanMonthIssues, item)
		} else {
			moreThanYearIssues = append(moreThanYearIssues, item)
		}
	}

	for _, issue := range lessMonthIssues {
		fmt.Printf("最近一月:\t%s\t%s\t%s\t%s\n", issue.Title, issue.HTMLURL, issue.CreatedAt, issue.User.Login)
	}

	for _, issue := range lessYearMoreThanMonthIssues {
		fmt.Printf("最近一年:\t%s\t%s\t%s\t%s\n", issue.Title, issue.HTMLURL, issue.CreatedAt, issue.User.Login)
	}

	for _, issue := range moreThanYearIssues {
		fmt.Printf("一年以前:\t%s\t%s\t%s\t%s\n", issue.Title, issue.HTMLURL, issue.CreatedAt, issue.User.Login)
	}
}
