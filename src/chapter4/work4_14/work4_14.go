package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"

	"github.com/caistrong/gopl-homework/src/chapter4/work4_10/github"
)

// Note: 原题描述不是很清晰，不知道所谓的bug报告、里程碑和对应的用户信息是想展示什么。但是我觉得基础的技术和我下面的实现大差不差，因此复用之前Issues的API

// run command
// go run ./src/chapter4/work4_14/work4_14.go
// then
// visit http://localhost:8888/?q=%E7%9F%A5%E8%AF%86%E5%9B%BE%E8%B0%B1%20%E6%99%BA%E8%83%BD%E9%97%AE%E7%AD%94%E7%B3%BB%E7%BB%9F
// or http://localhost:8888/?q=the%%20go%%20program%%20language
func main() {
	//启动一个web服务器
	http.HandleFunc("/", handle)
	fmt.Printf("visit http://localhost:8888/?q=the%%20go%%20program%%20language\n")
	log.Fatal(http.ListenAndServe("localhost:8888", nil))
}

func handle(w http.ResponseWriter, r *http.Request) {
	query := r.FormValue("q")
	if query == "" {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprintf(w, "<html><body>You need a param 'q', you can visit like<a href=\"http://localhost:8888/?q=the%%20go%%20program%%20language\">http://localhost:8888/?q=the%%20go%%20program%%20language</a></body></html>")
		return
	}
	var result *github.IssuesSearchResult
	var keywords = strings.Split(query, " ")
	result, _ = github.SearchIssues(keywords)

	var issueList = template.Must(template.New("issuelist").Parse(`
<h1>{{.TotalCount}} issues</h1>
<table>
<tr style='text-align: left'>
  <th>#</th>
  <th>State</th>
  <th>User</th>
  <th>Title</th>
</tr>
{{range .Items}}
<tr>
  <td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
  <td>{{.State}}</td>
  <td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
  <td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
</tr>
{{end}}
</table>
`))
	issueList.Execute(w, result)
}
