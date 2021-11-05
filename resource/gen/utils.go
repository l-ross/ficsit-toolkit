package main

import (
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
	u "unicode"
)

var startDigit = regexp.MustCompile(`^\d`)

func createNameFromClassName(className string) string {
	// Create Name based on the ClassName
	// We want to:
	// - Remove the _C suffix
	// - Ideally remove everything before the first underscore. However, we need to keep it
	//   if this would result in the Name starting with a number as we would then have an invalid
	//   variable name.
	// - Replace hyphens with underscores.

	s := strings.Split(className, "_")
	l := len(s)
	name := strings.Join(s[1:l-1], "")
	if startDigit.MatchString(name) {
		// Name starts with a digit so add the first element back.
		name = strings.Join(s[0:l-1], "")
	}
	name = strings.ReplaceAll(name, "-", "_")

	return name
}

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
