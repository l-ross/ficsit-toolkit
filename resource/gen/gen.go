package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go/format"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"sort"
	"strings"
	"text/template"

	"github.com/spf13/pflag"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

var file string

func init() {
	pflag.StringVarP(&file, "file", "f", "./Docs.json", "Path to the Docs.json file")
}

func main() {
	pflag.Parse()

	err := realMain()
	if err != nil {
		log.Fatal(err)
	}
}

func realMain() error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	r := transform.NewReader(f, unicode.UTF16(unicode.LittleEndian, unicode.ExpectBOM).NewDecoder().Transformer)

	resources := make([]Resource, 0)

	err = json.NewDecoder(r).Decode(&resources)
	if err != nil {
		return err
	}

	for _, r := range resources {
		if r.NativeClass != "Class'/Script/FactoryGame.FGItemDescriptor'" {
			continue
		}

		err = r.generate()
		if err != nil {
			return err
		}
	}

	return nil
}

type Resource struct {
	NativeClass string   `json:"NativeClass"`
	Classes     []*Class `json:"Classes"`

	Name         string        `json:"-"`
	StructFields []StructField `json:"-'"`
	structFields map[string]string
}

type StructField struct {
	Name  string
	Type  string
	Value string
}

type Class struct {
	Name   string
	Type   string
	Fields []StructField

	className string
	rawFields map[string]interface{}
}

var lowerM = regexp.MustCompile(`^m`)

func (c *Class) UnmarshalJSON(b []byte) error {
	c.rawFields = make(map[string]interface{})

	err := json.Unmarshal(b, &c.rawFields)
	if err != nil {
		return err
	}

	c.className = c.rawFields["ClassName"].(string)
	delete(c.rawFields, "ClassName")

	s := strings.Split(c.className, "_")
	l := len(s)
	c.Name = strings.Join(s[1:l-1], "")

	for n, v := range c.rawFields {
		n2 := lowerM.ReplaceAllString(n, "M")

		if n2 != n {
			c.rawFields[n2] = v
			delete(c.rawFields, n)
		}

	}

	return nil
}

func (r *Resource) generate() error {
	r.Name = strings.Split(r.NativeClass, ".FG")[1]
	r.Name = strings.ReplaceAll(r.Name, "'", "")

	r.structFields = make(map[string]string)
	r.StructFields = make([]StructField, 0)

	for _, c := range r.Classes {
		err := r.determineFields(c)
		if err != nil {
			return err
		}

		c.Type = r.Name
	}

	sort.Slice(r.Classes, func(i, j int) bool {
		if r.Classes[i].Name < r.Classes[j].Name {
			return true
		}
		return false
	})

	for n, t := range r.structFields {
		r.StructFields = append(r.StructFields, StructField{
			Name: n,
			Type: string(t),
		})
	}

	sort.Slice(r.StructFields, func(i, j int) bool {
		if r.StructFields[i].Name < r.StructFields[j].Name {
			return true
		}
		return false
	})

	fileName := toSnakeCase(r.Name)
	fileName = fmt.Sprintf("../%s.go", fileName)

	var b bytes.Buffer

	err := tpl.Execute(&b, r)
	if err != nil {
		return err
	}

	formatted, err := format.Source(b.Bytes())
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(fileName, formatted, 0644)
	if err != nil {
		return err
	}

	return nil
}

var floatRegexp = regexp.MustCompile(`^\d+\.\d+$`)

var intRegexp = regexp.MustCompile(`^\d+$`)

func (r *Resource) determineFields(c *Class) error {
	c.Fields = make([]StructField, 0)

	for n, f := range c.rawFields {
		switch f.(type) {
		case string:
			fStr := f.(string)

			sf := StructField{
				Name: n,
			}

			switch {
			case strings.ToUpper(fStr) == "TRUE":
				sf.Type = "bool"
				sf.Value = "true"
			case strings.ToUpper(fStr) == "FALSE":
				sf.Type = "bool"
				sf.Value = "false"
			case strings.HasPrefix(fStr, "SS_"):
				sf.Type = "StackSize"

				v, ok := stackSize[fStr]
				if !ok {
					return fmt.Errorf("failed to parse stack size value %s", fStr)
				}
				sf.Value = v
			case strings.HasPrefix(fStr, "RF_"):
				sf.Type = Form

				v, ok := resourceForm[fStr]
				if !ok {
					return fmt.Errorf("failed to parse resource form value %s", fStr)
				}
				sf.Value = v
			case strings.HasPrefix(fStr, "(("):
			case strings.HasPrefix(fStr, "("):

				ty, va, err := genStruct(fStr, false)
				if err != nil {
					return err
				}
				sf.Type = ty
				sf.Value = va
			case floatRegexp.MatchString(fStr):
				sf.Type = "float64"
				sf.Value = fStr
			case intRegexp.MatchString(fStr):
				sf.Type = "int"
				sf.Value = fStr
			default:
				sf.Type = "string"
				sf.Value = fmt.Sprintf("`%s`", fStr)
			}

			if t2, ok := r.structFields[n]; ok && t2 != sf.Type {
				return fmt.Errorf("type mismatch in field %s with types %s and %s", n, sf.Type, t2)
			}

			if sf.Type != "" {
				c.Fields = append(c.Fields, sf)
			}

			r.structFields[n] = sf.Type
		default:
			return fmt.Errorf("unknown type of field %s", n)
		}
	}

	sort.Slice(c.Fields, func(i, j int) bool {
		if c.Fields[i].Name < c.Fields[j].Name {
			return true
		}
		return false
	})

	return nil
}

type FieldType string

const (
	Boolean        FieldType = "bool"
	StackSize                = "StackSize"
	Form                     = "Form"
	ArrayOfStructs           = "[]struct{}"
	Struct                   = "struct{}"
	Float                    = "float64"
	Int                      = "int"
	String                   = "string"
)

var tpl = template.Must(template.New("resource").Parse(`
package resource

type {{.Name}} struct{
	Name string
	{{range .StructFields}}		{{.Name}} {{.Type}}
{{end}}
}

var (
	{{range .Classes}}	{{.Name}} = {{.Type}}{
		{{range .Fields}}		{{.Name}}: {{.Value}},
		{{end}}
	}

	{{end}}
)


`))

var stackSize = map[string]string{
	"SS_ONE":    "One",
	"SS_SMALL":  "Small",
	"SS_MEDIUM": "Medium",
	"SS_BIG":    "Big",
	"SS_HUGE":   "Huge",
	"SS_FLUID":  "Fluid",
}

var resourceForm = map[string]string{
	"RF_INVALID": "Invalid",
	"RF_SOLID":   "Solid",
	"RF_LIQUID":  "Liquid",
	"RF_GAS":     "Gas",
	"RF_HEAT":    "Heat",
}

// Source: https://gist.github.com/stoewer/fbe273b711e6a06315d19552dd4d33e6
var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

func toSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}
