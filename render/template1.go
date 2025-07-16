package render

import (
	"log"
	"os"
	"text/template"
	"time"
)

func RenderTemplate1() {
	// All the file reference is from root, no matter from which package it is called
	tpl1, err := template.ParseFiles("./templates/template1.gohtml")

	if err != nil {
		log.Fatalln(err)
	}

// | Value  | Means...           |
// | ------ | ------------------ |
// | `02`   | Day (of month)     |
// | `01`   | Month (numeric)    |
// | `Jan`  | Month (short name) |
// | `2006` | Year               |
// | `15`   | Hour (24h)         |
// | `03`   | Hour (12h)         |
// | `04`   | Minute             |
// | `05`   | Second             |
// | `PM`   | AM/PM              |
// | `MST`  | Time zone          |

	errExc := tpl1.Execute(os.Stdout, map[string]string{
		"CurrentTime": time.Now().Format("02-Jan-2006"),
		"Message":     "Hello world",
	})

	if errExc != nil {
		log.Fatalln(errExc)
	}
}
