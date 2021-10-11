package main

import (
	"fmt"
	"os"
	"text/template"
)

func main() {

	type Book struct {
		Author string
		Title  string
	}

	t := template.New("book template")

	t, _ = t.Parse(`{{range .Book}}
     Book title: {{ .Title }}
{{end}}`)

	p := make([]Book, 0, 2)
	p = []Book{

		{Author: "J.Tolkien", Title: "LOTR"},
		{Author: "Orwel", Title: "1984"},
	}

	if err := t.Execute(os.Stdout, p); err != nil {
		fmt.Println(err)
	}

}
