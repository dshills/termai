package aitool

import "encoding/json"

type structProp struct {
	name     string
	desc     string
	props    map[string]Property
	required bool
}

func (sd structProp) Name() string {
	return sd.name
}

func (sd structProp) Type() string {
	return "object"
}

func (sd structProp) MarshalJSON() ([]byte, error) {
	type custom struct {
		Type        string              `json:"type"`
		Description string              `json:"description,omitempty"`
		Properties  map[string]Property `json:"properties"`
	}
	cus := custom{Type: "object", Description: sd.desc, Properties: sd.props}
	return json.MarshalIndent(&cus, "", "\t")
}

func (sd structProp) UnmarshalJSON([]byte) error {
	return nil
}

func (sd structProp) Required() bool {
	return sd.required
}

func NewStruct(name, desc string, required bool, props ...Property) Property {
	sd := structProp{name: name, desc: desc, required: required, props: make(map[string]Property)}
	for _, p := range props {
		sd.props[p.Name()] = p
	}
	return &sd
}

/*
{
  "type": "object",
  "properties": {
    "number": { "type": "number" },
    "street_name": { "type": "string" },
    "street_type": { "enum": ["Street", "Avenue", "Boulevard"] }
  }
}
*/
