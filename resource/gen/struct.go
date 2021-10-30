package main

import (
	"bytes"
	"fmt"
	"text/template"
)

type state int

const (
	Start state = iota
	StartElem
	Name
	Value
	EndElem
	End
)

func genStruct(str string, isArray bool) (string, string, error) {
	s := Start

	fields := make([]fieldDef, 0)
	fd := fieldDef{}

	buffer := ""

	for _, v := range str {
		v2 := string(v)
		switch s {
		case Start:
			if v2 != "(" {
				return "", "", fmt.Errorf("expected '(' but got %q", v2)
			}

			s = Name
			if isArray {
				s = StartElem
			}
		case Name:
			if v2 == "=" {
				if len(buffer) == 0 {
					return "", "", fmt.Errorf("empty name")
				}
				fd.Name = buffer
				buffer = ""
				s = Value
				continue
			}

			if v2 == ")" {
				break
			}

			buffer += v2
		case Value:
			if v2 == "," || v2 == ")" {
				if len(buffer) == 0 {
					return "", "", fmt.Errorf("empty value")
				}
				fd.Value = buffer
				fd.Type = determineType(buffer)
				fields = append(fields, fd)
				fd = fieldDef{}
				buffer = ""

				switch v2 {
				case ",":
					s = Name
				case ")":
					s = End
				}
				s = Name
				continue
			}

			buffer += v2
		case End:
		default:
			return "", "", fmt.Errorf("unknown state %v", s)
		}
	}

	var b1 bytes.Buffer

	err := fieldDefTpl.Execute(&b1, fields)
	if err != nil {
		return "", "", err
	}

	var b2 bytes.Buffer

	err = valueDefTpl.Execute(&b2, fields)
	if err != nil {
		return "", "", err
	}

	return b1.String(), b2.String(), nil
}

type fieldDef struct {
	Name  string
	Type  string
	Value string
}

type example struct {
	x struct {
		y int
		z int
	}
}

var e = example{
	x: struct {
		y int
		z int
	}{
		y: 0,
		z: 0,
	},
}

func determineType(s string) string {
	switch {
	case intRegexp.MatchString(s):
		return "int"
	default:
		return "string"
	}
}

var fieldDefTpl = template.Must(template.New("FieldDef").Parse(`	struct {
	{{range .}} 		{{.Name}} {{.Type}}
	{{end}}}`))

var valueDefTpl = template.Must(template.New("ValueDef").Parse(`	struct {
	{{range .}} 		{{.Name}} {{.Type}}
	{{end}}}{
		{{range .}} {{.Name}}: {{.Value}},
		{{end}}}`))
