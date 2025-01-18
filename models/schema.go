package models

const (
	// FieldTypeByte is the byte type
	FieldTypeByte FieldType = "byte"
	// FieldTypeString is the string type
	FieldTypeString = "string"
	// FieldTypeDateStr is the int type
	FieldTypeDateStr = "date"
	// FieldTypeInt is the int type
	FieldTypeInt = "int"
	//FeildTypeInt64 is the float type
	FieldTypeInt64 = "long"
	// FieldTypeFloat is the float type
	FieldTypeFloat32 = "float"
	// FieldTypeDouble is the double type
	FieldTypeFloat64 = "double"
	// FieldTypeBool is the bool type
	FieldTypeBool = "bool"
	// FieldTypeObject is the object type
	FieldTypeObject = "object"
	// FieldTypeArray is the array type
	FieldTypeArray = "array"
)

// FieldType is the type of the field
type FieldType string

// Schema represents the structure of an input/output schema with various attributes.
//
// Fields:
//   - Name: The name of the input/output.
//   - Description: The description of the input/output.
//   - Type: The type of the field.
//   - Items: Used when the type is an array, represents the schema of the array items.
//   - Properties: Used when the type is an object, represents the schema of the object properties.
//   - Enum: The enum of the field used.
//   - Default: The default value of the field.
//   - Required: The required flag of the field.
type Schema struct {
	// Name is the name of the input/output
	Name string `json:"name" yaml:"name"`
	// Description is the description of the input/output
	Description string `json:"description" yaml:"description"`
	// Type is the type of the field
	Type FieldType `json:"type" yaml:"type"`
	// Items used when the type is an array
	Items *Schema `json:"items" yaml:"items"`
	// Properties used when the type is an object
	Properties []*Schema `json:"properties" yaml:"properties"`
	// Enum is the enum of the field used
	Enum any `json:"enum" yaml:"enum"`
	// Default is the default value of the field
	Default any `json:"default" yaml:"default"`
	// Required is the required flag of the field
	Required bool `json:"required" yaml:"required"`
}
