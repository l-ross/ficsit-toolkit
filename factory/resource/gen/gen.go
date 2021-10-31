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
		err = r.generate()
		if err != nil {
			return err
		}
	}

	return nil
}

type Resource struct {
	// Fields that are populated by Docs.json
	NativeClass string   `json:"NativeClass"`
	Classes     []*Class `json:"Classes"`

	// TypeName is derived from NativeClass.
	// e.g `Class'/Script/FactoryGame.FGItemDescriptor'` becomes `FGItemDescriptor`
	TypeName string `json:"-"`

	// PkgName is the name of the package this resource will be in.
	// Derived from the TypeName.
	// e.g. `FGItemDescriptor` becomes `item_descriptor`
	PkgName string `json:"-"`

	// StructDefs defines the names and types that will appear in the definition
	// of the struct for this Resource type
	StructDefs []*StructField `json:"-'"`

	// Map defining the names of the fields that will appear in the struct and their types
	// This is populated as we parse the classes within the resource and ensures that
	// we don't end up with the same field in multiple classes but different types.
	// Is used to populate StructDefs
	// Key = name of field
	// Value = type of field
	structDef map[string]string

	// If true then one of the Classes is using a type defined in the base resource package, so we
	// need to import it.
	ImportResourcePkg bool `json:"-"`

	HasFullName bool `json:"-"`

	// As the classes are being parsed any fields where the value is an array are populated here.
	// Once we have finished parsing all the classes we can then figure out the format of these fields.
	// We need to wait until we have parsed all classes as the fields within the array may differ
	// by class.
	arrayFields map[string][]interface{}
}

type StructField struct {
	Name  string
	Type  string
	Value string
}

// Class is a class as defined with Docs.json
type Class struct {
	// The ClassName as taken from Docs.json
	ClassName string

	// The FullName as taken from Docs.json
	// Not all classes have this field.
	FullName string

	// Name is derived from the ClassName and will be used as the variable name.
	// e.g. ClassName `Desc_NuclearWaste_C` becomes `NuclearWaste`
	Name string

	// The name of the resource type that the Class variable will be a type of.
	ResourceTypeName string

	// The fields for this Class variable
	Fields []*StructField

	HasFullName bool `json:"-"`

	// Set by Class.UnmarshalJSON.
	// We initially read all fields of a class in as interface types.
	rawFields map[string]interface{}
}

var invalidChars = regexp.MustCompile(`[^a-zA-Z0-9_]`)
var startDigit = regexp.MustCompile(`^\d`)

func (c *Class) UnmarshalJSON(b []byte) error {
	// Unmarshal all the fields
	c.rawFields = make(map[string]interface{})
	err := json.Unmarshal(b, &c.rawFields)
	if err != nil {
		return err
	}

	// Set ClassName
	c.ClassName = c.rawFields["ClassName"].(string)
	delete(c.rawFields, "ClassName")

	if i, ok := c.rawFields["FullName"]; ok {
		if s, ok := i.(string); ok {
			c.FullName = s
			delete(c.rawFields, "FullName")
			c.HasFullName = true
		}
	}

	// Set Name based on the ClassName
	// We want to:
	// - Remove the _C suffix
	// - Ideally remove everything before the first underscore. However, we need to keep it
	//   if this would result in the Name starting with a number as we would then have an invalid
	//   variable name.
	// - Replace hyphens with underscores.
	s := strings.Split(c.ClassName, "_")
	l := len(s)
	c.Name = strings.Join(s[1:l-1], "")
	if startDigit.MatchString(c.Name) {
		// Name starts with a digit so add the first element back.
		c.Name = strings.Join(s[0:l-1], "")
	}
	c.Name = strings.ReplaceAll(c.Name, "-", "_")

	for name, value := range c.rawFields {
		// Make sure the first character is upper case so that
		// the field is public in Go.
		// Lots of the fields start 'm'
		newName := firstCharUpper(name)
		// Remove any characters that would not be valid in the field name in Go.
		// Some fields contain '?'
		newName = invalidChars.ReplaceAllString(newName, "")

		// If the above has resulted in newName differing from name then update
		// c.rawFields
		if newName != name {
			c.rawFields[newName] = value
			delete(c.rawFields, name)
		}
	}

	return nil
}

func (r *Resource) generate() error {
	r.TypeName, r.PkgName = nativeClassToTypeNameAndPkgName(r.NativeClass)

	r.structDef = make(map[string]string)
	r.StructDefs = make([]*StructField, 0)
	r.arrayFields = make(map[string][]interface{})

	// For each Class parse all of its fields.
	for _, c := range r.Classes {
		if c.HasFullName {
			r.HasFullName = true
		}

		fields, err := r.parseFields(c.rawFields)
		if err != nil {
			return err
		}

		c.Fields = fields
		c.ResourceTypeName = r.TypeName
	}

	// If any of the classes contained fields with an array value
	// then parse them and update the classes with the new fields.
	//
	// We have to parse the array fields for all classes at once
	// to ensure the struct definition is correct as not all instances of
	// an array may contain all fields.
	//
	// Example below shows the same field in 3 classes with different fields in each array item.
	//
	// "mUnlocks": [],
	//
	// "mUnlocks": [
	// 	{
	//		"Class": "BP_UnlockRecipe_C",
	//		"mRecipes": "(BlueprintGeneratedClass'\"/Game/FactoryGame/Recipes/Buildings/Recipe_Workshop.Recipe_Workshop_C\"',BlueprintGeneratedClass'\"/Game/FactoryGame/Recipes/Equipment/Recipe_PortableMiner.Recipe_PortableMiner_C\"')"
	//	},
	// ],
	//
	// "mUnlocks": [
	//	{
	//		"Class": "BP_UnlockScannableResource_C",
	//		"mResourcesToAddToScanner": "(BlueprintGeneratedClass'\"/Game/FactoryGame/Resource/RawResources/OreCopper/Desc_OreCopper.Desc_OreCopper_C\"')",
	//		"mResourcePairsToAddToScanner": "((ResourceDescriptor=BlueprintGeneratedClass'\"/Game/FactoryGame/Resource/RawResources/OreCopper/Desc_OreCopper.Desc_OreCopper_C\"'))"
	//	}
	// ],
	err := r.parseArrayFields()
	if err != nil {
		return err
	}

	// Sort classes and their fields so that code gen is consistent.
	sort.Slice(r.Classes, func(i, j int) bool {
		if r.Classes[i].Name < r.Classes[j].Name {
			return true
		}
		return false
	})

	for _, c := range r.Classes {
		sortFields(c.Fields)
	}

	// Create the struct definition for the resource type and sort the fields.
	for n, t := range r.structDef {
		r.StructDefs = append(r.StructDefs, &StructField{
			Name: n,
			Type: t,
		})
	}
	sortFields(r.StructDefs)

	fileName := toSnakeCase(r.PkgName)
	dirName := fmt.Sprintf("../%s", fileName)
	fileName = fmt.Sprintf("%s/%s.go", dirName, fileName)

	err = recreateDir(dirName)
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

// Parse each field in v and return a StructField for each.
func (r *Resource) parseFields(v map[string]interface{}) ([]*StructField, error) {
	fields := make([]*StructField, 0)

	for n, f := range v {
		sf, err := r.createStructField(n, f)
		if err != nil {
			return nil, err
		}

		// sf will be nil if the field was an array.
		// We have to skip these here and deal with them later.
		if sf != nil {
			// Check that we are being consistent with the type of a field.
			// e.g. If the field "TEST" was a float in one class but a string in another
			// then something has gone wrong.
			if storedType, ok := r.structDef[n]; ok && storedType != sf.Type {
				return nil, fmt.Errorf("type mismatch in field %s with types %s and %s", n, sf.Type, storedType)
			}

			r.structDef[n] = sf.Type
			fields = append(fields, sf)
		}
	}

	return fields, nil
}

var floatRegexp = regexp.MustCompile(`^-?\d+\.\d+$`)
var intRegexp = regexp.MustCompile(`^-?\d+$`)

func (r *Resource) createStructField(k string, v interface{}) (*StructField, error) {
	sf := &StructField{
		Name: firstCharUpper(k),
	}

	switch v.(type) {
	case string:
		// Field is a string, based on its value figure out the correct type.
		vString := v.(string)

		switch {
		case strings.ToUpper(vString) == "TRUE":
			sf.Type = "bool"
			sf.Value = "true"
		case strings.ToUpper(vString) == "FALSE":
			sf.Type = "bool"
			sf.Value = "false"
		case isEnum(vString):
			r.ImportResourcePkg = true
			value, t, err := getEnum(vString)
			if err != nil {
				return nil, err
			}

			sf.Value = value
			sf.Type = t

		//case strings.HasPrefix(vString, "SS_"):
		//	r.ImportResourcePkg = true
		//	sf.Type = "resource.StackSize"
		//
		//	v, ok := stackSize[vString]
		//	if !ok {
		//		return nil, fmt.Errorf("failed to parse stack size value %s", vString)
		//	}
		//	sf.Value = v
		//case strings.HasPrefix(vString, "RF_"):
		//	r.ImportResourcePkg = true
		//	sf.Type = "resource.Form"
		//
		//	v, ok := resourceForm[vString]
		//	if !ok {
		//		return nil, fmt.Errorf("failed to parse resource form value %s", vString)
		//	}
		//	sf.Value = v
		//case strings.HasPrefix(vString, "EST_"):
		//	r.ImportResourcePkg = true
		//	sf.Type = "resource.SchematicType"
		//
		//	v, ok := schematicType[vString]
		//	if !ok {
		//		return nil, fmt.Errorf("failed to parse schematic type value %s", vString)
		//	}
		//	sf.Value = v
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
	case []interface{}:
		// Field is an array.
		// Add to arrayFields so we can deal with it latter
		arr := v.([]interface{})

		if len(arr) == 0 {
			r.arrayFields[k] = append(r.arrayFields[k], nil)
		} else {
			r.arrayFields[k] = append(r.arrayFields[k], arr)
		}

		return nil, nil
	default:
		return nil, fmt.Errorf("unsupported type in field %s", k)
	}

	return sf, nil
}

func (r *Resource) parseArrayFields() error {
	// Here be dragons 🐉 and crufty code.

	for name, field := range r.arrayFields {
		// Keep track of the type of each field in the struct definition.
		structTypes := make(map[string]string)
		values := make([][][]*StructField, len(r.Classes))

		for i, f := range field {
			if f == nil {
				continue
			}

			// Example classArray:
			//  [
			//		{
			//			"Class": "BP_UnlockRecipe_C",
			//			"mRecipes": "(BlueprintGeneratedClass'\"/Game/FactoryGame/Recipes/Buildings/Recipe_SmelterBasicMk1.Recipe_SmelterBasicMk1_C\"',BlueprintGeneratedClass'\"/Game/FactoryGame/Recipes/Buildings/Recipe_PowerLine.Recipe_PowerLine_C\"',BlueprintGeneratedClass'\"/Game/FactoryGame/Recipes/Smelter/Recipe_IngotCopper.Recipe_IngotCopper_C\"',BlueprintGeneratedClass'\"/Game/FactoryGame/Recipes/Constructor/Recipe_Wire.Recipe_Wire_C\"',BlueprintGeneratedClass'\"/Game/FactoryGame/Recipes/Constructor/Recipe_Cable.Recipe_Cable_C\"')"
			//		},
			//		{
			//			"Class": "BP_UnlockScannableResource_C",
			//			"mResourcesToAddToScanner": "(BlueprintGeneratedClass'\"/Game/FactoryGame/Resource/RawResources/OreCopper/Desc_OreCopper.Desc_OreCopper_C\"')",
			//			"mResourcePairsToAddToScanner": "((ResourceDescriptor=BlueprintGeneratedClass'\"/Game/FactoryGame/Resource/RawResources/OreCopper/Desc_OreCopper.Desc_OreCopper_C\"'))"
			//		}
			//	],
			classArray := f.([]interface{})

			arrayValues := make([][]*StructField, 0)
			for _, classArrayElement := range classArray {
				// Example classArrayElement:
				//		{
				//			"Class": "BP_UnlockRecipe_C",
				//			"mRecipes": "(BlueprintGeneratedClass'\"/Game/FactoryGame/Recipes/Buildings/Recipe_SmelterBasicMk1.Recipe_SmelterBasicMk1_C\"',BlueprintGeneratedClass'\"/Game/FactoryGame/Recipes/Buildings/Recipe_PowerLine.Recipe_PowerLine_C\"',BlueprintGeneratedClass'\"/Game/FactoryGame/Recipes/Smelter/Recipe_IngotCopper.Recipe_IngotCopper_C\"',BlueprintGeneratedClass'\"/Game/FactoryGame/Recipes/Constructor/Recipe_Wire.Recipe_Wire_C\"',BlueprintGeneratedClass'\"/Game/FactoryGame/Recipes/Constructor/Recipe_Cable.Recipe_Cable_C\"')"
				//		},

				// Example classArrayElementField:
				// "Class": "BP_UnlockRecipe_C",
				classArrayElementField := classArrayElement.(map[string]interface{})

				arrayValue := make([]*StructField, 0)

				for k, v := range classArrayElementField {
					sf, err := r.createStructField(k, v)
					if err != nil {
						return err
					}

					// Check that all fields with the same name have the same type.
					if storedType, ok := structTypes[sf.Name]; ok && storedType != sf.Type {
						return fmt.Errorf("type mismatch in field %s with types %s and %s", sf.Name, sf.Type, storedType)
					}

					structTypes[sf.Name] = sf.Type
					arrayValue = append(arrayValue, sf)
				}

				sortFields(arrayValue)
				arrayValues = append(arrayValues, arrayValue)
			}

			values[i] = arrayValues
		}

		// Create type definition for array of structs.

		arrayStructDefs := make([]*StructField, 0)
		for n, t := range structTypes {
			arrayStructDefs = append(arrayStructDefs, &StructField{
				Name: n,
				Type: t,
			})
		}
		sortFields(arrayStructDefs)

		var b bytes.Buffer

		err := arrayStructDef.Execute(&b, arrayStructDefs)
		if err != nil {
			return err
		}

		r.structDef[name] = b.String()

		// Update each class to add their array fields.
		a := array{
			Fields: arrayStructDefs,
		}

		for i, v := range values {
			if v == nil {
				r.Classes[i].Fields = append(r.Classes[i].Fields, &StructField{
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

			r.Classes[i].Fields = append(r.Classes[i].Fields, &StructField{
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

var tpl = template.Must(template.New("resource").Parse(`
// Code generated by ../../gen/docs_json. DO NOT EDIT.

package {{.PkgName}}

import (
	"fmt"

	{{if .ImportResourcePkg}}"github.com/l-ross/ficsit-toolkit/resource" {{end}}
)

type {{.TypeName}} struct{
	ClassName string	{{if .HasFullName}} 
	FullName string {{end}}
	{{range .StructDefs}}		{{.Name}} {{.Type}}
{{end}}
}

var (
	{{range .Classes}}	{{.Name}} = {{.ResourceTypeName}}{
		ClassName: "{{.ClassName}}", {{if .HasFullName}}
		FullName: "{{.FullName}}", {{end}}
		{{range .Fields}}		{{.Name}}: {{.Value}},
		{{end}}
	}

	{{end}}
)

func GetByClassName(className string) (*{{.TypeName}}, error) {
	if v, ok := classNameToVar[className]; ok {
		return v, nil
	}

	return nil, fmt.Errorf("failed to find {{.TypeName}} with class name %s", className)
}

var classNameToVar = map[string]*{{.TypeName}} {
	{{range .Classes}}	"{{.ClassName}}": &{{.Name}},
	{{end}}
}

{{if .HasFullName}}
func GetByFullName(fullName string) (*{{.TypeName}}, error) {
	if v, ok := fullNameToVar[fullName]; ok {
		return v, nil
	}

	return nil, fmt.Errorf("failed to find {{.TypeName}} with full name %s", fullName)
}

var fullNameToVar = map[string]*{{.TypeName}} {
	{{range .Classes}} "{{.FullName}}": &{{.Name}},
	{{end}}
}
{{end}}

`))

var arrayStructDef = template.Must(template.New("ArrayStructDef").Parse(` []struct {
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
