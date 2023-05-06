package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/StuartsHome/cmd/ntriples"
)

func main() {
	var path string
	flag.StringVar(&path, "filepath", "", "")
	flag.Parse()

	//
	if path == "" {
		log.Fatalln("filepath required")
	}

	currTime := time.Now().Format("YYYY-MM-DD")

	// Create ntriples object.
	n := ntriples.New(fmt.Sprintf("output%s.ttl", currTime))
	err := n.Construct(path)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("complete")
}
