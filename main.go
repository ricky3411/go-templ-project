package main

import (
	"html/template"
	"os"

	"github.com/gin-gonic/gin"
)

type Page struct {
	Title string
	Body  string
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

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Testing Gin",
		})
	})
	r.Run(":8080")
}
