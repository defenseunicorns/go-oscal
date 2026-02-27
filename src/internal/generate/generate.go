package generate

import (
	"fmt"
	"go/format"
	"reflect"
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

	// Initialize the build process.
	err = config.initBuild(&schema)
	if err != nil {
		return typeBytes, err
	}

	err = config.buildStructs()
	if err != nil {
		return typeBytes, err
	}

	// Add header comment
	typeString := fmt.Sprintf("%s\n", headerComment)

	// Add the package name
	typeString += fmt.Sprintf("package %s\n\n", pkgName)

	// Add additional imports
	typeString += buildImportString()

	// Parse the schema ID for the version reference
	schemaId := schema.ID
	if schemaId == nil {
		return typeBytes, fmt.Errorf("unable to find $id in schema")
	}

	version, err := extractVersion(*schemaId)
	if err != nil {
		return typeBytes, err
	}
	typeString += fmt.Sprintf("const Version = %q\n", version)

	// Add the struct definitions in order of creation.
	for _, ref := range config.refQueue.History() {
		typeString += config.structMap[ref] + "\n\n"
	}

	typeBytes = []byte(typeString)

	// Format the bytes
	typeBytes, err = format.Source(typeBytes)
	if err != nil {
		return typeBytes, err
	}

	return typeBytes, nil
}

// initBuild initializes the build process by adding the root schema to the ref queue and building the definitions map.
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

	// No properties, so we need to add the properties from the oneOf or anyOf schemas.
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
		if schema.AnyOf != nil {
			for _, anyOf := range schema.AnyOf {
				if anyOf.TypeObject.Properties != nil {
					for key, prop := range anyOf.TypeObject.Properties {
						// Add all properties from anyOf variants, not just those with direct refs
						schema.Properties[key] = prop
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
	// While the ref queue is not empty, pop the ref and build the struct string.
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

	// Merge properties from anyOf or allOf schemas if the definition has no properties
	if len(def.Properties) == 0 {
		if def.Properties == nil {
			def.Properties = map[string]jsonschema.SchemaOrBool{}
		}
		if def.AnyOf != nil {
			for _, anyOf := range def.AnyOf {
				if anyOf.TypeObject.Properties != nil {
					for key, prop := range anyOf.TypeObject.Properties {
						def.Properties[key] = prop
					}
				}
				// Note: We don't merge required fields from anyOf variants
				// because only one variant is active at a time, so all fields
				// should be optional at the merged struct level
			}
		}
		if def.AllOf != nil {
			for _, allOf := range def.AllOf {
				if allOf.TypeObject.Properties != nil {
					for key, prop := range allOf.TypeObject.Properties {
						def.Properties[key] = prop
					}
				}
				// For allOf, we should merge required fields since all variants must be satisfied
				if allOf.TypeObject.Required != nil {
					def.Required = append(def.Required, allOf.TypeObject.Required...)
				}
			}
		}
	}

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

	// create a slice of the keys in the properties map
	var keys []string
	for key := range def.Properties {
		if !KeysToIgnore[key] {
			keys = append(keys, key)
		}
	}
	// Sort the keys alphabetically
	slices.Sort(keys)

	// If there are no properties, return a map[string]interface{} type
	if len(keys) == 0 {
		structString = fmt.Sprintf("type %s = map[string]interface{}", name)
		return structString, err
	}

	// Generate aliases for the struct if they exist
	if aliases, ok := Aliases[name]; ok {
		for _, alias := range aliases {
			structString += fmt.Sprintf("type %s = %s\n", alias, name)
		}
	}

	// Add top level struct definition
	structString += fmt.Sprintf("type %s struct {\n", name)

	// Add the properties to the struct string
	for _, key := range keys {
		// Set the parent of the property schema to the definition
		// Allows for the parent to be used in case of duplicate types
		def.Properties[key].TypeObject.Parent = &def

		// Get the property schema
		prop := def.Properties[key]
		propSchema := prop.TypeObject

		// Build the property name, type, and tags
		propName := FmtFieldName(key)
		propType, err := c.buildTypeString(*propSchema)
		if err != nil {
			return structString, err
		}

		propType = addPointerIfOptionalNonPrimitive(required[key], propType)
		propTags := buildTagString(c.tags, key, required[key])
		structString += fmt.Sprintf("\t%s %s %s\n", propName, propType, propTags)
	}
	// Close the struct
	structString += "}"
	if err != nil {
		return structString, err
	}

	return structString, err
}

// buildTypeString builds the type string for a given property.
func (c *GeneratorConfig) buildTypeString(property jsonschema.Schema) (propType string, err error) {
	var possibleRefs []string

	if property.Type != nil && property.Type.SimpleTypes != nil {
		jsonType := getJsonOrCustomType(property)
		// convert json type to go type
		propType = getGoType(jsonType)
		// if the type is not primitive, we need to add the name of the type
		if !isPrimitiveOrCustomJsonType(jsonType) {
			name, err := c.findSubType(property)
			if err != nil {
				return "", err
			}
			propType += name
		}
		return propType, err
	} else if property.Ref != nil {
		// If the property is a ref, we need to find the schema for the ref and build the type string for that schema.
		def, ok := c.definitions[*property.Ref]
		if !ok {
			return "", fmt.Errorf("unable to find schema for %s", *property.Ref)
		}
		// Set the parent of the definition to the property
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
			// Set the parent of the definition to the property
			def.Parent = &property
			return c.buildTypeString(def)
		} else {
			// Could not determine the type of the prop so return an error.
			return "", fmt.Errorf("could not determine type for property %v", property)
		}
	}
}

// findSubType finds the name of the type for the given schema.
func (c *GeneratorConfig) findSubType(schema jsonschema.Schema) (name string, err error) {
	simpleType := getJsonOrCustomType(schema)
	switch simpleType {
	case "object":
		// If the schema has a ref, we need to find the name of the ref.
		ref, err := getRef(schema)
		if err != nil {
			return name, err
		}
		// Create a name from the ref
		name = getNameFromRef(ref)
		// Check if the name is already in use and if so, append the parent name to the name.
		ref, name = c.handleDuplicates(ref, name, schema)
		c.nameMap[name] = ref
		// Add the ref to the ref queue if it is not already in the queue.
		c.refQueue.Add(ref)
		// Add the schema to the definitions map if it is not already in the map.
		if _, ok := c.definitions[ref]; !ok {
			c.definitions[ref] = schema
		}
	case "array":
		// If the schema has items, we need to find the name of the items.
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
		name = simpleType
	}

	return name, err
}

// handleDuplicates checks if the name is already in use and if so, appends the parent name to the name.
func (c *GeneratorConfig) handleDuplicates(ref string, name string, schema jsonschema.Schema) (string, string) {
	if currentRef, ok := c.nameMap[name]; ok {
		// Points to a different definition
		if currentRef != ref {
			// If the definitions are the same, return the current ref and name
			existing := c.definitions[currentRef]
			if reflect.DeepEqual(existing, schema) {
				return currentRef, name
			}
			// If the definitions are different, try the title
			if schema.Title != nil && (existing.Title == nil || *existing.Title != *schema.Title) {
				newName := getNameFromTitle(*schema.Title)
				return c.handleDuplicates(getRefWithName(newName), newName, schema)
			}
			// If the title is the same, try the parent
			parent := schema.Parent
			parentRef, _ := getRef(*parent)
			if parentRef != "" {
				prefix := getNameFromRef(parentRef)
				if prefix != name {
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
