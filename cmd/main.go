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
	var inputPath string
	var namespace string
	flag.StringVar(&inputPath, "inputpath", "", "")
	flag.StringVar(&namespace, "namespace", "", "")
	flag.Parse()

	// Parse args.
	if err := validateRequiredArgs(inputPath, namespace); err != nil {
		log.Fatalln(err)
	}

	// Create config.
	config := config.New(
		config.Namespace(namespace),
		inputPath,
		fmt.Sprintf("output-%s.ttl", time.Now().Format("20060102")),
	)

	// Create ntriples object.
	n := ntriples.New(
		config,
	)
	err := n.Construct()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("complete")
}

func validateRequiredArgs(args ...interface{}) error {
	for _, arg := range args {
		if arg.(string) == "" {
			return fmt.Errorf("inputpath, namespace required")
		}
	}
	return nil
}
