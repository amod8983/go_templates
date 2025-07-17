package render

import (
	"log"
	"os"
	"text/template"
)

var templ *template.Template

func init() {
	templ = template.Must(template.ParseFiles("./templates/range-template.gohtml"))
}

func RenderRangeTemplate(isObj bool) {
	fruits := []string{"apple", "orange", "banana", "kiwi", "grapes"}
	appleProp := map[string]string{
		"name": "Apple",
		"weight": "100gm",
		"color": "red",
	}
	var err  error
	if isObj {
		err = templ.Execute(os.Stdout, appleProp)
	} else {
		err = templ.Execute(os.Stdout, fruits)
	}
	if err != nil {
		log.Fatalln(err)
	}
}