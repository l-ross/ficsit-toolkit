package main

import (
	"fmt"
	"log"
	"os"

	"github.com/l-ross/ficsit-toolkit/save"
)

func main() {
	err := realMain()
	if err != nil {
		log.Fatal(err)
	}
}

func realMain() error {
	f, err := os.Open("./Example.sav")
	if err != nil {
		return err
	}
	defer f.Close()

	s, err := save.Parse(f)
	if err != nil {
		return err
	}

	fmt.Println(len(s.Components))
	fmt.Println(s.Components[0].TypePath)
	fmt.Println(len(s.Entities))
	fmt.Println(s.Entities[0].TypePath)

	fOut, err := os.OpenFile("out.sav", os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer fOut.Close()

	err = save.Serialize(s, fOut)
	if err != nil {
		return err
	}

	f2, err := os.Open("out.sav")
	if err != nil {
		return err
	}

	s2, err := save.Parse(f2)
	if err != nil {
		return err
	}

	fmt.Println(len(s2.Components))
	fmt.Println(s2.Components[0].TypePath)
	fmt.Println(len(s2.Entities))
	fmt.Println(s2.Entities[0].TypePath)

	return nil
}
