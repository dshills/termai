package aitool

import "encoding/json"

type NumericType string

const (
	NumericFloat NumericType = "number"
	NumericInt   NumericType = "integer"
	NumericBool  NumericType = "boolean"
	NumericNil   NumericType = "null"
)

type numericProp struct {
	name     string
	numType  NumericType
	desc     string
	required bool
}

func (nd numericProp) Name() string {
	return nd.name
}

func (nd numericProp) Type() string {
	return string(nd.numType)
}

func (nd numericProp) MarshalJSON() ([]byte, error) {
	type custom struct {
		Type        string `json:"type,omitempty"`
		Description string `json:"description,omitempty"`
	}
	cus := custom{Type: string(nd.numType), Description: nd.desc}
	return json.MarshalIndent(&cus, "", "\t")
}

func (nd numericProp) UnmarshalJSON([]byte) error {
	return nil
}

func (nd numericProp) Required() bool {
	return nd.required
}

func NewNumeric(name, desc string, required bool, numType NumericType) Property {
	return &numericProp{name: name, desc: desc, required: required, numType: numType}
}

/*
{ "type": "boolean" }
{ "type": "integer" }
{ "type": "numeric" }
{ "type": "null" }
*/
