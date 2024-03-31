package model

// A trick for data encapsulation with unexported types and JSON Handling 
type issueUnexported struct {
	Key    string         `json:"key"`
	Fields map[string]any `json:"fields"`
}

type Issue struct {
	issueUnexported
}

func (iss Issue) GetKey() string {
	return iss.Key
}

func (iss Issue) GetType() string {
	return iss.Fields["type"].(string)
}

func (iss Issue) GetParam(key string) any {
	return iss.Fields[key]
}

func (iss Issue) GetAllParams() *map[string]any {
	return &iss.Fields
}

type Issues struct {
	Issues []Issue `json:"issues"`
}
