package main

import (
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
	u "unicode"
)

var fgRegexp = regexp.MustCompile(`^FG`)

func nativeClassToTypeNameAndPkgName(s string) (string, string) {
	split := strings.Split(s, ".")
	if len(split) == 1 {
		panic(fmt.Sprintf("unexpected NativeClass: %s", s))
	}

	typeName := split[1]
	typeName = strings.ReplaceAll(typeName, "'", "")
	pkgName := fgRegexp.ReplaceAllString(typeName, "")

	return typeName, pkgName
}

func firstCharUpper(s string) string {
	tmp := []rune(s)
	tmp[0] = u.ToUpper(tmp[0])
	return string(tmp)
}

// Source: https://gist.github.com/stoewer/fbe273b711e6a06315d19552dd4d33e6
var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

func toSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

func sortFields(sf []*StructField) {
	sort.Slice(sf, func(i, j int) bool {
		if sf[i].Name < sf[j].Name {
			return true
		}
		return false
	})
}

func recreateDir(dirName string) error {
	err := os.RemoveAll(dirName)
	if err != nil {
		return err
	}

	err = os.Mkdir(dirName, 0700)
	if err != nil {
		return err
	}

	return nil
}
