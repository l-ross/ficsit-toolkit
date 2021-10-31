package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strings"

	"github.com/l-ross/ficsit-toolkit/save"
)

func main() {
	err := realMain()
	if err != nil {
		log.Fatal(err)
	}
}

func realMain() error {
	f, err := os.Open("./Simple.sav")
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

	paths := make([]string, 0)

	for _, e := range s.Entities {
		if strings.HasPrefix(e.TypePath, "/Game/FactoryGame/Buildable") {
			paths = append(paths, e.TypePath)
		}
	}

	sort.Strings(paths)

	for _, p := range paths {
		fmt.Println(p)
	}

	return nil
}
