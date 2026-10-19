package main

import (
	"bytes"
	"fmt"
	"log"
	"text/template"
)

type InspectOpenGauss struct {
	Name  string `mapstructure:"name"`
	Value string `mapstructure:"value"`
}

var TemplateFuncMap = template.FuncMap{
	"test": func(str string) string {
		return fmt.Sprintf("testFunc_%s", str)
	},
}

func TemplateProcess(t *template.Template, data any) (string, error) {
	var tmplBytes bytes.Buffer
	if err := t.Execute(&tmplBytes, data); err != nil {
		return "", err
	}
	return tmplBytes.String(), nil
}

func main() {
	var contentTemplate = `{{.Value}}{{test .Name}}`
	t, err := template.New("opengauss").Funcs(TemplateFuncMap).Parse(contentTemplate)
	if err != nil {
		log.Fatal(err)
	}
	data := InspectOpenGauss{
		Name:  "name",
		Value: "test_value",
	}
	result, err := TemplateProcess(t, data)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(result)
}
