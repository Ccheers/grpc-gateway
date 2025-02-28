/*
 * A Bit of Everything
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Contact: none@example.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package abe

import (
	"time"
)

// Intentionally complicated message type to cover many features of Protobuf.
type ExamplepbABitOfEverythingServiceUpdateBody struct {
	SingleNested *ABitOfEverythingNested  `json:"singleNested"`
	Nested       []ABitOfEverythingNested `json:"nested"`
	// Float value field
	FloatValue          float32                         `json:"floatValue"`
	DoubleValue         float64                         `json:"doubleValue"`
	Int64Value          string                          `json:"int64Value"`
	Uint64Value         string                          `json:"uint64Value"`
	Int32Value          int32                           `json:"int32Value"`
	Fixed64Value        string                          `json:"fixed64Value"`
	Fixed32Value        int64                           `json:"fixed32Value"`
	BoolValue           bool                            `json:"boolValue"`
	StringValue         string                          `json:"stringValue"`
	BytesValue          string                          `json:"bytesValue"`
	Uint32Value         int64                           `json:"uint32Value"`
	EnumValue           *ExamplepbNumericEnum           `json:"enumValue"`
	PathEnumValue       *PathenumPathEnum               `json:"pathEnumValue"`
	NestedPathEnumValue *MessagePathEnumNestedPathEnum  `json:"nestedPathEnumValue"`
	Sfixed32Value       int32                           `json:"sfixed32Value"`
	Sfixed64Value       string                          `json:"sfixed64Value"`
	Sint32Value         int32                           `json:"sint32Value"`
	Sint64Value         string                          `json:"sint64Value"`
	RepeatedStringValue []string                        `json:"repeatedStringValue"`
	OneofEmpty          *interface{}                    `json:"oneofEmpty"`
	OneofString         string                          `json:"oneofString"`
	MapValue            map[string]ExamplepbNumericEnum `json:"mapValue"`
	// Map of string description.
	MappedStringValue        map[string]string                 `json:"mappedStringValue"`
	MappedNestedValue        map[string]ABitOfEverythingNested `json:"mappedNestedValue"`
	NonConventionalNameValue string                            `json:"nonConventionalNameValue"`
	TimestampValue           time.Time                         `json:"timestampValue"`
	RepeatedEnumValue        []ExamplepbNumericEnum            `json:"repeatedEnumValue"`
	// Repeated numeric enum description.
	RepeatedEnumAnnotation []ExamplepbNumericEnum `json:"repeatedEnumAnnotation"`
	// Numeric enum description.
	EnumValueAnnotation *ExamplepbNumericEnum `json:"enumValueAnnotation"`
	// Repeated string description.
	RepeatedStringAnnotation []string `json:"repeatedStringAnnotation"`
	// Repeated nested object description.
	RepeatedNestedAnnotation []ABitOfEverythingNested `json:"repeatedNestedAnnotation"`
	// Nested object description.
	NestedAnnotation                           *ABitOfEverythingNested `json:"nestedAnnotation"`
	Int64OverrideType                          int64                   `json:"int64OverrideType"`
	RequiredStringViaFieldBehaviorAnnotation   string                  `json:"requiredStringViaFieldBehaviorAnnotation"`
	OutputOnlyStringViaFieldBehaviorAnnotation string                  `json:"outputOnlyStringViaFieldBehaviorAnnotation"`
	OptionalStringValue                        string                  `json:"optionalStringValue,omitempty"`
	// Only digits are allowed.
	ProductId                           []string `json:"productId"`
	OptionalStringField                 string   `json:"optionalStringField"`
	RequiredStringField1                string   `json:"requiredStringField1"`
	RequiredStringField2                string   `json:"requiredStringField2"`
	RequiredFieldBehaviorJsonNameCustom string   `json:"required_field_behavior_json_name_custom"`
	RequiredFieldSchemaJsonNameCustom   string   `json:"required_field_schema_json_name_custom"`
	TrailingOnly                        string   `json:"trailingOnly"`
	// Trailing only dot.
	TrailingOnlyDot string `json:"trailingOnlyDot"`
	// Trailing both.
	TrailingBoth string `json:"trailingBoth"`
	// This is an example of a multi-line comment.  Trailing multiline.
	TrailingMultiline string   `json:"trailingMultiline"`
	Uuids             []string `json:"uuids"`
}
