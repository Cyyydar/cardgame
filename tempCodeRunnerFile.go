package main

import (
    "html/template"
    "os"
)

type a struct {
    Title   []string
    Article [][]string
}

var data = &a{
    Title: []string{"One", "Two", "Three"},
    Article: [][]string{
        []string{"a", "b", "c"},
        []string{"d", "e"},
        []string{"f", "g", "h", "i"}},
}

var tmplSrc = `
{{range $i, $a := .Title}}
  Title: {{$i}} {{$a}}
  {{range $article := index $.Article $i}}
    Article: {{$article}}.
  {{end}}
{{end}}`

func main() {
    tmpl := template.Must(template.New("test").Parse(tmplSrc))
    tmpl.Execute(os.Stdout, data)
}
