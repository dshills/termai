package aitool

import "encoding/json"

type arrayProp struct {
	name     string
	itemType Property
	desc     string
	required bool
}

func (ad arrayProp) Name() string {
	return ad.name
}

func (ad arrayProp) Type() string {
	return "array"
}

func (ad arrayProp) MarshalJSON() ([]byte, error) {
	type cusItem struct {
		Type string `json:"type"`
	}
	type custom struct {
		Type        string  `json:"type,omitempty"`
		Description string  `json:"description,omitempty"`
		Items       cusItem `json:"items"`
	}
	ct := custom{Type: "array", Description: ad.desc, Items: cusItem{Type: ad.itemType.Type()}}
	return json.MarshalIndent(&ct, "", "\t")
}

func (ad arrayProp) UnmarshalJSON([]byte) error {
	return nil
}

func (ad arrayProp) Required() bool {
	return ad.required
}

func NewArray(name, desc string, required bool, itemType Property) Property {
	return &arrayProp{name: name, desc: desc, itemType: itemType, required: required}
}

/*
{
  "type": "array",
  "items": {
    "type": "number"
  }
}
*/
