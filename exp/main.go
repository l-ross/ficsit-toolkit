package main

import (
	"log"
	"os"
	"regexp"

	"github.com/l-ross/ficsit-toolkit/factory"
)

func main() {
	err := realMain()
	if err != nil {
		log.Fatal(err)
	}
}

var idRegexp = regexp.MustCompile(`.*_C_(\d+)`)

func realMain() error {
	f, err := os.Open("./Simple.sav")
	if err != nil {
		return err
	}
	defer f.Close()

	//s, err := save.Parse(f)
	//if err != nil {
	//	return err
	//}
	//
	//uniqueIDs := make(map[string]string)
	//
	//for _, c := range s.Components {
	//	matches := idRegexp.FindStringSubmatch(c.ParentEntityName)
	//	if matches != nil {
	//		entityName := matches[0]
	//		id := matches[1]
	//
	//		if v, ok := uniqueIDs[id]; ok && v != entityName {
	//			fmt.Printf("Duplicates %q %q\n", v, entityName)
	//			os.Exit(0)
	//		}
	//	}
	//}

	//
	//d, err := json.MarshalIndent(s, "", "  ")
	//if err != nil {
	//	return err
	//}

	//d, err := save.DumpBody(f)
	//if err != nil {
	//	return err
	//}

	//err = ioutil.WriteFile("Simple.json", d, 0644)
	//if err != nil {
	//	return err
	//}

	_, err = factory.Load(f)
	if err != nil {
		return err
	}

	//for _, e := range s.Entities {
	//	if e.InstanceName == "Persistent_Level:PersistentLevel.Build_ConstructorMk1_C_2147415632" {
	//		spew.Dump(e.References)
	//	}
	//}

	//paths := make([]string, 0)
	//
	//for _, e := range s.Entities {
	//	if strings.HasPrefix(e.TypePath, "/Game/FactoryGame/Buildable") {
	//		paths = append(paths, e.TypePath)
	//	}
	//}
	//
	//sort.Strings(paths)
	//
	//for _, p := range paths {
	//	fmt.Println(p)
	//}

	return nil
}
