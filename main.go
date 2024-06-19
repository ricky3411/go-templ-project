package main 

import (
	"html/template"
	"os"
)

type Page struct {
	Title string
	Body string
}

func main() {
	t := template.Must(template.New("page").Parse(`
	<html>
		<head>
			<title>{{ .Title }}</title>
		</head>
		<body>
			<h1>{{ .Title }}</h1>
			<p>{{ .Body }}</p>
		</body>
	</html>
	`))

	p := Page{Title: "My Go Page", Body: "This is a simple Go page using templates."}
	err := t.Execute(os.Stdout, p)
    if err != nil {
        panic(err)
    }
}