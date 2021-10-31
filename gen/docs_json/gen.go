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
		//if r.NativeClass != "Class'/Script/FactoryGame.FGSchematic'" {
		//	continue
		//}

		fmt.Println(r.NativeClass)

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

	TypeName     string `json:"-"`
	PkgName      string
	StructFields []StructField `json:"-'"`
	structFields map[string]string

	ImportResourcePkg bool

	arrayFields map[string][]interface{}
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

	ClassName string
	rawFields map[string]interface{}
}

var lowerM = regexp.MustCompile(`^m`)

var invalidChars = regexp.MustCompile(`[^a-zA-Z0-9_]`)

var startDigit = regexp.MustCompile(`^\d`)

func (c *Class) UnmarshalJSON(b []byte) error {
	c.rawFields = make(map[string]interface{})

	err := json.Unmarshal(b, &c.rawFields)
	if err != nil {
		return err
	}

	c.ClassName = c.rawFields["ClassName"].(string)
	delete(c.rawFields, "ClassName")

	s := strings.Split(c.ClassName, "_")
	l := len(s)
	c.Name = strings.Join(s[1:l-1], "")
	if startDigit.MatchString(c.Name) {
		c.Name = strings.Join(s[0:l-1], "")
	}

	c.Name = strings.ReplaceAll(c.Name, "-", "_")

	for n, v := range c.rawFields {
		n2 := lowerM.ReplaceAllString(n, "M")
		n2 = invalidChars.ReplaceAllString(n2, "")

		if n2 != n {
			c.rawFields[n2] = v
			delete(c.rawFields, n)
		}
	}

	return nil
}

var fgRegexp = regexp.MustCompile(`^FG`)

func (r *Resource) generate() error {
	r.TypeName = strings.Split(r.NativeClass, ".")[1]
	r.TypeName = strings.ReplaceAll(r.TypeName, "'", "")
	r.PkgName = fgRegexp.ReplaceAllString(r.TypeName, "")

	r.structFields = make(map[string]string)
	r.StructFields = make([]StructField, 0)
	r.arrayFields = make(map[string][]interface{})

	for _, c := range r.Classes {
		fields, err := r.determineFields(c.rawFields)
		if err != nil {
			return err
		}

		c.Fields = fields
		c.Type = r.TypeName
	}

	err := r.determineArrayFields()
	if err != nil {
		return err
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
			Type: t,
		})
	}

	sort.Slice(r.StructFields, func(i, j int) bool {
		if r.StructFields[i].Name < r.StructFields[j].Name {
			return true
		}
		return false
	})

	fileName := toSnakeCase(r.PkgName)
	dirName := fmt.Sprintf("../../resource/%s", fileName)
	fileName = fmt.Sprintf("%s/%s.go", dirName, fileName)

	os.RemoveAll(dirName)

	err = os.Mkdir(dirName, 0700)
	if err != nil {
		return err
	}

	var b bytes.Buffer

	err = tpl.Execute(&b, r)
	if err != nil {
		return err
	}

	formatted, err := format.Source(b.Bytes())
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(fileName, formatted, 0700)
	if err != nil {
		return err
	}

	return nil
}

func (r *Resource) determineArrayFields() error {
	for name, field := range r.arrayFields {
		structTypes := make(map[string]string)
		values := make([][][]*StructField, len(r.Classes))

		for i, f := range field {
			if f == nil {
				continue
			}
			f2 := f.([]interface{})

			arrayValues := make([][]*StructField, 0)
			for _, f3 := range f2 {

				f4 := f3.(map[string]interface{})

				arrayValue := make([]*StructField, 0)

				for k, v := range f4 {
					sf, err := r.createStructField(k, v)
					if err != nil {
						return err
					}

					if t, ok := structTypes[sf.Name]; ok && t != sf.Type {
						return fmt.Errorf("type mismatch in field %s with types %s and %s", sf.Name, sf.Type, t)
					}

					structTypes[sf.Name] = sf.Type

					arrayValue = append(arrayValue, sf)
				}

				sort.Slice(arrayValue, func(i, j int) bool {
					if arrayValue[i].Name < arrayValue[j].Name {
						return true
					}
					return false
				})

				arrayValues = append(arrayValues, arrayValue)

			}

			values[i] = arrayValues

		}

		sortedTypes := make([]*StructField, 0)

		for n, t := range structTypes {
			sortedTypes = append(sortedTypes, &StructField{
				Name: n,
				Type: t,
			})
		}

		sort.Slice(sortedTypes, func(i, j int) bool {
			if sortedTypes[i].Name < sortedTypes[j].Name {
				return true
			}
			return false
		})

		var b bytes.Buffer

		err := structTpl.Execute(&b, sortedTypes)
		if err != nil {
			return err
		}

		r.structFields[name] = b.String()

		a := array{
			Fields: sortedTypes,
		}

		for i, v := range values {
			if v == nil {
				r.Classes[i].Fields = append(r.Classes[i].Fields, StructField{
					Name:  name,
					Value: "nil",
				})
				continue
			}

			a.Values = v

			var b bytes.Buffer

			err := arrayStructValTpl.Execute(&b, a)
			if err != nil {
				return err
			}

			r.Classes[i].Fields = append(r.Classes[i].Fields, StructField{
				Name:  name,
				Value: b.String(),
			})

		}
	}

	return nil
}

type array struct {
	Fields []*StructField
	Values [][]*StructField
}

func (r *Resource) createStructField(k string, v interface{}) (*StructField, error) {
	sf := &StructField{
		Name: lowerM.ReplaceAllString(k, "M"),
	}

	switch v.(type) {
	case string:
		vString := v.(string)

		switch {
		case strings.ToUpper(vString) == "TRUE":
			sf.Type = "bool"
			sf.Value = "true"
		case strings.ToUpper(vString) == "FALSE":
			sf.Type = "bool"
			sf.Value = "false"
		case strings.HasPrefix(vString, "SS_"):
			r.ImportResourcePkg = true
			sf.Type = "resource.StackSize"

			v, ok := stackSize[vString]
			if !ok {
				return nil, fmt.Errorf("failed to parse stack size value %s", vString)
			}
			sf.Value = v
		case strings.HasPrefix(vString, "RF_"):
			r.ImportResourcePkg = true
			sf.Type = "resource.Form"

			v, ok := resourceForm[vString]
			if !ok {
				return nil, fmt.Errorf("failed to parse resource form value %s", vString)
			}
			sf.Value = v
		case floatRegexp.MatchString(vString):
			sf.Type = "float64"
			sf.Value = vString
		case intRegexp.MatchString(vString):
			sf.Type = "int"
			sf.Value = vString
		default:
			sf.Type = "string"
			sf.Value = fmt.Sprintf("`%s`", vString)
		}
	default:
		return nil, fmt.Errorf("unsupported type in field %s", k)
	}

	return sf, nil
}

var floatRegexp = regexp.MustCompile(`^-?\d+\.\d+$`)

var intRegexp = regexp.MustCompile(`^-?\d+$`)

func (r *Resource) determineFields(v map[string]interface{}) ([]StructField, error) {
	fields := make([]StructField, 0)

	i := 0
	for n, f := range v {
		sf := StructField{
			Name: n,
		}

		switch f.(type) {
		case string:
			fStr := f.(string)

			switch {
			case strings.ToUpper(fStr) == "TRUE":
				sf.Type = "bool"
				sf.Value = "true"
			case strings.ToUpper(fStr) == "FALSE":
				sf.Type = "bool"
				sf.Value = "false"
			case strings.HasPrefix(fStr, "SS_"):
				r.ImportResourcePkg = true
				sf.Type = "resource.StackSize"

				v, ok := stackSize[fStr]
				if !ok {
					return nil, fmt.Errorf("failed to parse stack size value %s", fStr)
				}
				sf.Value = v
			case strings.HasPrefix(fStr, "RF_"):
				r.ImportResourcePkg = true
				sf.Type = "resource.Form"

				v, ok := resourceForm[fStr]
				if !ok {
					return nil, fmt.Errorf("failed to parse resource form value %s", fStr)
				}
				sf.Value = v
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
				return nil, fmt.Errorf("type mismatch in field %s with types %s and %s", n, sf.Type, t2)
			}

			r.structFields[n] = sf.Type

		case []interface{}:
			f2 := f.([]interface{})

			if len(f2) == 0 {
				r.arrayFields[n] = append(r.arrayFields[n], nil)
			} else {
				r.arrayFields[n] = append(r.arrayFields[n], f.([]interface{}))
			}

		default:
			return nil, fmt.Errorf("unknown type of field %s", n)
		}

		if sf.Type != "" {
			fields = append(fields, sf)
		}

		i++
	}

	sort.Slice(fields, func(i, j int) bool {
		if fields[i].Name < fields[j].Name {
			return true
		}
		return false
	})

	return fields, nil
}

var tpl = template.Must(template.New("resource").Parse(`
package {{.PkgName}}

import (
	"fmt"

	{{if .ImportResourcePkg}}"github.com/l-ross/ficsit-toolkit/resource" {{end}}
)

type {{.TypeName}} struct{
	ClassName string
	{{range .StructFields}}		{{.Name}} {{.Type}}
{{end}}
}

var (
	{{range .Classes}}	{{.Name}} = {{.Type}}{
		ClassName: "{{.ClassName}}",
		{{range .Fields}}		{{.Name}}: {{.Value}},
		{{end}}
	}

	{{end}}
)

func GetByClassName(className string) (*{{.TypeName}}, error) {
	if v, ok := classNameToVar[className]; ok {
		return v, nil
	}

	return nil, fmt.Errorf("failed to find {{.TypeName}} with name %s", className)
}

var classNameToVar = map[string]*{{.TypeName}} {
	{{range .Classes}}	"{{.ClassName}}": &{{.Name}},
	{{end}}
}

`))

var structTpl = template.Must(template.New("ArrayStructDev").Parse(` []struct {
	{{range .}} {{.Name}} {{.Type}}
	{{end}}
}

`))

var arrayStructValTpl = template.Must(template.New("ArrayStructVal").Parse(`[]struct {
	{{range .Fields}} {{.Name}} {{.Type}}
	{{end}}
}{ 
	{{range .Values}} { 
	{{range .}}{{.Name}}: {{.Value}},
		{{end}}
	},
	{{end}}
}`))

//{{range .Values}} {{.Name}}: {{.Value}},
//	{{end}}

var stackSize = map[string]string{
	"SS_ONE":    "resource.One",
	"SS_SMALL":  "resource.Small",
	"SS_MEDIUM": "resource.Medium",
	"SS_BIG":    "resource.Big",
	"SS_HUGE":   "resource.Huge",
	"SS_FLUID":  "resource.Fluid",
}

var resourceForm = map[string]string{
	"RF_INVALID": "resource.Invalid",
	"RF_SOLID":   "resource.Solid",
	"RF_LIQUID":  "resource.Liquid",
	"RF_GAS":     "resource.Gas",
	"RF_HEAT":    "resource.Heat",
}

// Source: https://gist.github.com/stoewer/fbe273b711e6a06315d19552dd4d33e6
var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

func toSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}
