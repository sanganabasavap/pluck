package main

import (
	"os"
	"testing"
	"text/template"
)

func TestTemplateGeneration(t *testing.T) {
	data := struct {
		Name map[string]string
	}{Name: map[string]string{"out.text": "some sample"}}

	textTmpl := `{{range $key,$value:=.Name}}
value for {{$key}} is {{$value}}
{{end}}
`
	templ := template.New("yaml")
	parse, err := templ.Parse(textTmpl)
	if err != nil {
		return
	}
	err = parse.Execute(os.Stdout, data)
}

func TestLoadConfig(t *testing.T) {

}
