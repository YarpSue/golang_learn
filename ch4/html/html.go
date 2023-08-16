package main

import (
	"html/template"
	"log"
	"os"
	"testmode/ch4/github"
	"time"
)

const templ = `{{.TotalCount}} issues:
{{range .Items}}----------------------------------------
Number: {{.Number}}
User: {{.User.Login}}
Title: {{.Title | printf "%.64s"}}
Age: {{.CreatedAt | daysAgo}} days
{{end}}`

const temp2 = `
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
`



func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

func main() {
	//report, err := template.New("report").
	//	Funcs(template.FuncMap{"daysAgo": daysAgo}).
	//	Parse(templ)
	//if err != nil {
	//	log.Fatal(err)
	//}
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	//if err := report.Execute(os.Stdout, result); err != nil {
	//	log.Fatal(err)
	//}

	var issueList = template.Must(template.New("issuelist").Parse(temp2))
	if err := issueList.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}