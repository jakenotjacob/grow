package main

import (
	"fmt"
	"net"
	"os"
	"text/template"
)

type Worker struct {
	Id      int
	Conns   chan net.Conn
	ConnLen int
}

const workerview = `
Worker[{{.Id}}] {{ range .Conns }} {{ printf "#" }} {{ end }}
`

// Returns worker id dispatched to or error
func Dispatch(workers []*Worker, conn net.Conn, signals chan int) (interface{}, error) {
	// Dispatch connection to first worker with capacity available
	for _, worker := range workers {
		if len(worker.Conns) < cap(worker.Conns) {
			worker.Conns <- conn
			worker.ConnLen = len(worker.Conns)
			signals <- worker.Id
		}
		//TODO closing causes infinite loop
		//close(signals)
		return worker.Id, fmt.Errorf("All workers are at full capacity! (sending to waiting room?)")
	}
	return nil, fmt.Errorf("Failed to dispatch %v to worker pool", conn)
}

func main() {

	signals := make(chan int)

	// Gather the workforce
	workers := []*Worker{
		&Worker{Id: 2, Conns: make(chan net.Conn, 1)},
		&Worker{Id: 3, Conns: make(chan net.Conn, 1)},
	}

	go func(chan int) {
		for {
			select {
			case <-signals:
				tmpl := template.Must(template.New("workerview").Parse(workerview))
				for worker := range workers {
					tmpl.Execute(os.Stdout, worker)
				}
			}
		}
	}(signals)
	// Listen for signals to print changes

	// Listen up
	server, err := net.Listen("tcp", ":31337")
	if err != nil {
		fmt.Printf("Couldn't start server!: %v", err)
	}

	// Dispatch to workers
	for {
		conn, err := server.Accept()
		if err != nil {
			fmt.Printf("Failed to accept client: %v", err)
		}
		go Dispatch(workers, conn, signals)
	}

	// We don't error error here since template.Must() will panic on failure
}
