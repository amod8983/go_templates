package render

import (
	"log"
	"os"
	"text/template"
	"time"
)

func RenderTemplate1() {
	tpl1, err := template.ParseFiles("./templates/template1.gohtml")

	if err != nil {
		log.Fatalln(err)
	}

	errExc := tpl1.Execute(os.Stdout, map[string]string{
		"CurrentTime": time.Now().Format("02-Jan-2006"),
		"Message":     "Hello world",
	})

	if errExc != nil {
		log.Fatalln(errExc)
	}
}
