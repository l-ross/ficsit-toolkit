package main

import (
	"log"
	"os"

	"github.com/davecgh/go-spew/spew"

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

	//for _, e := range s.Entities {
	//	if e.InstanceName == "Persistent_Level:PersistentLevel.Build_ConstructorMk1_C_2147415632" {
	//		spew.Dump(e.References)
	//	}
	//}

	spew.Dump(s.Header)

	return nil
}
