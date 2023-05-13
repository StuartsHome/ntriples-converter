package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/StuartsHome/cmd/config"
	"github.com/StuartsHome/cmd/ntriples"
)

func main() {
	var path string
	var namespace string
	flag.StringVar(&path, "filepath", "", "")
	flag.StringVar(&namespace, "namespace", "", "")
	flag.Parse()

	// Parse args.
	if path == "" {
		log.Fatalln("filepath required")
	}
	if namespace == "" {
		log.Fatalln("namespace required")
	}

	// Create config.
	config := config.New(
		"namespace",
		fmt.Sprintf("output-%s.ttl", time.Now().Format("20060102")),
	)

	// Create ntriples object.
	n := ntriples.New(
		*config,
	)
	err := n.Construct(path)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("complete")
}
