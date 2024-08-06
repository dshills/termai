package aitool

import "encoding/json"

type stringProp struct {
	name     string
	desc     string
	enums    []string
	required bool
}

func (sd stringProp) Name() string {
	return sd.name
}

func (sd stringProp) Type() string {
	return "string"
}

func (sd stringProp) MarshalJSON() ([]byte, error) {
	type custom struct {
		Type        string   `json:"type"`
		Description string   `json:"description,omitempty"`
		Enum        []string `json:"enum,omitempty"`
	}
	cus := custom{Type: "string", Description: sd.desc, Enum: sd.enums}
	return json.MarshalIndent(&cus, "", "\t")
}

func (sd stringProp) UnmarshalJSON([]byte) error {
	return nil
}

func (sd stringProp) Required() bool {
	return sd.required
}

func NewString(name, desc string, required bool, enums ...string) Property {
	return &stringProp{name: name, desc: desc, required: required, enums: enums}
}

/*
{ "type": "string" }
*/
