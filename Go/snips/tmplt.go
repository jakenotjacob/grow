package main

import (
	"fmt"
	"os"
	"text/template"
)

type Person struct {
	Name string
}

func main() {
	const letter = `
Hello, {{ .Name }}.
`

	t := template.Must(template.New("letter").Parse(letter))
	p := &Person{"Joe"}

	err := t.Execute(os.Stdout, p)
	if err != nil {
		fmt.Println("Ohnodidntwork")
	}

}
