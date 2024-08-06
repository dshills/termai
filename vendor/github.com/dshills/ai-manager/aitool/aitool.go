package aitool

import "encoding/json"

type Property interface {
	json.Marshaler
	json.Unmarshaler
	Name() string
	Type() string
	Required() bool
}

type Tool struct {
	fxName string
	fxDesc string
	params []Property
}

func (t *Tool) Name() string {
	return t.fxName
}

func (t *Tool) MarshalJSON() ([]byte, error) {
	parmd := paramDef{Type: "object", Properties: make(map[string]Property)}
	for _, p := range t.params {
		parmd.Properties[p.Name()] = p
		if p.Required() {
			parmd.Required = append(parmd.Required, p.Name())
		}
	}
	fund := funcDef{Name: t.fxName, Description: t.fxDesc, Parameters: parmd}
	tool := toolDef{Type: "function", Function: fund}
	return json.MarshalIndent(&tool, "", "\t")
}

func NewTool(fxName, fxDesc string, params ...Property) *Tool {
	return &Tool{fxName: fxName, fxDesc: fxDesc, params: params}
}

type toolDef struct {
	Type     string  `json:"type"`
	Function funcDef `json:"function"`
}

type funcDef struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Parameters  paramDef `json:"parameters"`
}

type paramDef struct {
	Type       string              `json:"type"`
	Properties map[string]Property `json:"properties"`
	Required   []string            `json:"required"`
}

/*
	ToolChoice string `json:"tool_choice"`
{
  "model": "gpt-4-turbo",
  "messages": [
    {
      "role": "user",
      "content": "Whats the weather like in Boston today?"
    }
  ],
  "tools": [
    {
      "type": "function",
      "function": {
        "name": "get_current_weather",
        "description": "Get the current weather in a given location",
        "parameters": {
          "type": "object",
          "properties": {
            "location": {
              "type": "string",
              "description": "The city and state, e.g. San Francisco, CA"
            },
            "unit": {
              "type": "string",
              "enum": [
                "celsius",
                "fahrenheit"
              ]
            }
          },
          "required": [
            "location"
          ]
        }
      }
    }
  ],
  "tool_choice": "auto"
}
*/
