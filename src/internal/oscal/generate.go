package oscal

import (
	"fmt"
	"go/format"
	"slices"

	"github.com/swaggest/jsonschema-go"
)

// BaseFlags represents command-line flags for the base go-oscal command.
type BaseFlags struct {
	InputFile  string // -f / --input-file
	OutputFile string // -o / --output-file
	Pkg        string // -p / --pkg
	Tags       string // -t / --tags
}

type GeneratorConfig struct {
	tags        []string
	pkgName     string
	refQueue    RefQueue
	definitions map[string]jsonschema.Schema
	structMap   map[string]string
	nameMap     map[string]string
}

func NewGeneratorConfig(tags []string, pkgName string) GeneratorConfig {
	return GeneratorConfig{
		tags:        tags,
		pkgName:     pkgName,
		refQueue:    NewRefQueue(),
		definitions: map[string]jsonschema.Schema{},
		structMap:   map[string]string{},
		nameMap:     map[string]string{},
	}
}

// Generate a struct definition given a JSON string representation of an object.
func Generate(oscalSchema []byte, pkgName string, tags []string) (typeBytes []byte, err error) {
	schema, err := populateSchema(oscalSchema)
	if err != nil {
		return typeBytes, err
	}

	config := NewGeneratorConfig(tags, pkgName)

	err = config.initBuild(&schema)
	if err != nil {
		return typeBytes, err
	}

	err = config.buildStructs()
	if err != nil {
		return typeBytes, err
	}

	typeString := fmt.Sprintf("%s\n", headerComment)
	typeString += fmt.Sprintf("package %s\n\n", pkgName)
	for _, ref := range config.refQueue.History() {
		typeString += config.structMap[ref] + "\n\n"
	}

	typeBytes = []byte(typeString)

	typeBytes, err = format.Source(typeBytes)
	if err != nil {
		return typeBytes, err
	}

	return typeBytes, nil
}

func (c *GeneratorConfig) initBuild(schema *jsonschema.Schema) (err error) {
	if schema.ID == nil {
		return fmt.Errorf("unable to find $id in schema")
	}

	if schema.Definitions != nil {
		c.definitions, err = getDefinitionMap(*schema)
		if err != nil {
			return err
		}
	}

	if schema.Ref != nil {
		c.refQueue.Add(*schema.Ref)
		return nil
	}

	if schema.Properties == nil {
		schema.Properties = map[string]jsonschema.SchemaOrBool{}
		if schema.OneOf != nil {
			for _, oneOf := range schema.OneOf {
				if oneOf.TypeObject.Properties != nil {
					for key, prop := range oneOf.TypeObject.Properties {
						if prop.TypeObject.Ref != nil && !RefsToIgnore[*prop.TypeObject.Ref] {
							schema.Properties[key] = prop
						}
					}
				}
			}
		}
	}

	// add the schema to the definitions map
	c.definitions[*schema.ID] = *schema
	// Add the root schema to the ref queue.
	c.refQueue.Add(*schema.ID)
	return err
}

func (c *GeneratorConfig) buildStructs() (err error) {
	for c.refQueue.Len() > 0 {
		ref := c.refQueue.Pop()

		// Look up the schema for the ref, return an error if not found.
		def, ok := c.definitions[ref]
		if !ok {
			return fmt.Errorf("unable to find schema for %s", ref)
		}

		// Build the struct string for the ref.
		structString, err := c.buildStructString(def)
		if err != nil {
			return err
		}

		// Add the struct string to the struct map.
		c.structMap[ref] = structString
	}

	return err
}

func (c *GeneratorConfig) buildStructString(def jsonschema.Schema) (structString string, err error) {

	// Create a map of required fields
	required := map[string]bool{}
	for _, req := range def.Required {
		required[req] = true
	}

	// Get the name of the struct
	name, err := c.findSubType(def)
	if err != nil {
		return structString, err
	}

	structStr := fmt.Sprintf("type %s struct {\n", name)
	var keys []string
	for key := range def.Properties {
		keys = append(keys, key)
	}
	slices.Sort(keys)

	for _, key := range keys {
		def.Properties[key].TypeObject.Parent = &def
		prop := def.Properties[key]
		propSchema := prop.TypeObject
		propName := FmtFieldName(key)
		propType, err := c.buildTypeString(*propSchema)
		if err != nil {
			return structStr, err
		}
		propTags := buildTagString(c.tags, key, required[key])
		structStr += fmt.Sprintf("\t%s %s %s\n", propName, propType, propTags)
	}
	structStr += "}"
	formattedStruct, err := format.Source([]byte(structStr))
	if err != nil {
		return structStr, err
	}

	return string(formattedStruct), err
}

func (c *GeneratorConfig) buildTypeString(property jsonschema.Schema) (propType string, err error) {
	var possibleRefs []string

	if property.Type != nil && property.Type.SimpleTypes != nil {
		jsonType := string(*property.Type.SimpleTypes)
		// convert json type to go type
		propType = getGoType(jsonType)
		// if the type is not primitive, we need to add the name of the type
		if !isPrimitiveType(jsonType) {
			name, err := c.findSubType(property)
			if err != nil {
				return "", err
			}
			propType += name
		}
		return propType, err
	} else if property.Ref != nil {
		def, ok := c.definitions[*property.Ref]
		if !ok {
			return "", fmt.Errorf("unable to find schema for %s", *property.Ref)
		}
		def.Parent = &property
		return c.buildTypeString(def)

	} else {
		// TODO: Handle anyOf, allOf, oneOf
		// We should be creating new structs for these types.
		// Right now assumes that the first is a ref to a primitive per the current patterns. This may not be true.
		// Should probably create a new struct with each of the possible types. (logic for creating golang enum?, just use primitives?, "or" types?)
		for _, schema := range property.AllOf {
			if schema.TypeObject.Ref != nil {
				possibleRefs = append(possibleRefs, *schema.TypeObject.Ref)
			}
		}
		for _, schema := range property.AnyOf {
			if schema.TypeObject.Ref != nil {
				possibleRefs = append(possibleRefs, *schema.TypeObject.Ref)
			}
		}
		// Find the first possible ref and recurse.
		if len(possibleRefs) > 0 {
			def, ok := c.definitions[possibleRefs[0]]
			if !ok {
				return "", fmt.Errorf("unable to find schema for %s", possibleRefs[0])
			}
			def.Parent = &property
			return c.buildTypeString(def)
		} else {
			// Could not determine the type of the prop so return an error.
			return "", fmt.Errorf("could not determine type for property %v", property)
		}
	}
}

func (c *GeneratorConfig) findSubType(schema jsonschema.Schema) (name string, err error) {
	simpleType := getSimpleType(schema)
	switch simpleType {
	case "object":
		ref := getRef(schema)
		name = getNameFromRef(ref)
		ref, name = c.handleDuplicates(ref, name, schema)
		c.nameMap[name] = ref
		c.refQueue.Add(ref)
		if _, ok := c.definitions[ref]; !ok {
			c.definitions[ref] = schema
		}
	case "array":
		if schema.Items.SchemaOrBoolEns() != nil {
			def := *schema.Items.SchemaOrBool.TypeObject
			def.Parent = &schema
			name, err = c.buildTypeString(def)
		} else if schema.Items.SchemaArray != nil {
			def := *schema.Items.SchemaArray[0].TypeObject
			def.Parent = &schema
			name, err = c.buildTypeString(def)
		} else {
			err = fmt.Errorf("could not determine name for %v", schema)
		}
	case "":
		if schema.Ref != nil {
			def := c.definitions[*schema.Ref]
			def.Parent = &schema
			name, err = c.buildTypeString(def)
		} else if schema.ID != nil {
			def := c.definitions[*schema.ID]
			def.Parent = &schema
			name, err = c.buildTypeString(def)
		} else {
			err = fmt.Errorf("could not determine name for %v", schema)
		}
	default:
		name = getSimpleType(schema)
	}

	return name, err
}

func (c *GeneratorConfig) handleDuplicates(ref string, name string, schema jsonschema.Schema) (string, string) {
	if currentRef, ok := c.nameMap[name]; ok {
		if currentRef != ref {
			parent := schema.Parent
			parentRef := getRef(*parent)
			if parentRef != "" {
				prefix := getNameFromRef(parentRef)
				if prefix != name {
					// pattern := regexp.MustCompile(`[a-z]+`)
					// split := pattern.Split(prefix, -1)
					// if len(split) > 2 {
					// 	prefix = strings.Join(split, "")
					// }
					newName := prefix + name
					newRef := "#/definitions/" + newName
					return c.handleDuplicates(newRef, newName, *parent)
				}
			}
			return c.handleDuplicates(ref, name, *parent)
		}
	}
	return ref, name
}
