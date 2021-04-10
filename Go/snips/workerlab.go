package main

import (
	"log"
	"os"
	"text/template"
)

type Worker struct {
	Id    int8
	Conns []int8
}

func main() {
	const workerview = `
Worker[{{.Id}}]{{range .Conns}}{{ printf "#"}}{{end}}
`
	tmpl := template.Must(template.New("workerview").Parse(workerview))
	if tmpl == nil {
		log.Fatal("Could not parse worker template!")
	}

	w1 := &Worker{Id: 9, Conns: []int8{1, 2}}
	w2 := &Worker{Id: 9, Conns: []int8{9, 6, 4, 5, 2}}
	w3 := &Worker{Id: 9, Conns: []int8{3, 9, 8}}

	workers := []*Worker{w1, w2, w3}
	//tmpl.Execute(os.Stdout, w1)

	for _, worker := range workers {
		tmpl.Execute(os.Stdout, worker)
	}
}
