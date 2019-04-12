package main

import (
	"os"
	"text/template"
)

type ToDo struct {
	Name        string
	Description string
}

func main() {
	td := ToDo{Name: "Go", Description: "Get basic knowledge in Go"}
	t, err := template.New("Todos").Parse("You have a task named \"{{ .Name}}\" with description: \"{{ .Description}}\"\n")
	if err != nil {
		panic(err)
	}
	err = t.Execute(os.Stdout, td)
	if err != nil {
		panic(err)
	}
}
