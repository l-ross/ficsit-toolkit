package main

import (
	"log"
	"os"

	"github.com/l-ross/satisfactorysave"
)

func main() {
	err := realMain()
	if err != nil {
		log.Fatal(err)
	}
}

func realMain() error {
	f, err := os.Open("./Luke.sav")
	if err != nil {
		return err
	}
	defer f.Close()

	p, err := satisfactorysave.NewParser(f)
	if err != nil {
		return err
	}

	_, err = p.Parse()
	if err != nil {
		return err
	}

	return nil
}
