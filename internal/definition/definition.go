package definition

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Service struct {
	Host      string     `json:"host"`
	Port      int        `json:"port"`
	Endpoints []Endpoint `json:"endpoints"`
}

type Endpoint struct {
	Name    string   `json:"name"`
	Path    string   `json:"path"`
	Actions []Action `json:"actions"`
}

type Action struct {
	Method      string `json:"method"`
	Description string `json:"description"`
}

var _ json.Unmarshaler = (*Service)(nil)
var _ json.Unmarshaler = (*Endpoint)(nil)
var _ json.Unmarshaler = (*Action)(nil)

func (s *Service) UnmarshalJSON(data []byte) error {
	raw := make(map[string]json.RawMessage)
	err := json.Unmarshal(data, &raw)
	if err != nil {
		return err
	}

	err = json.Unmarshal(raw["host"], &s.Host)
	if err != nil {
		return fmt.Errorf("invalid host %s", raw["host"])
	}
	err = json.Unmarshal(raw["port"], &s.Port)
	if err != nil {
		return fmt.Errorf("invalid port %s", raw["port"])
	}
	err = json.Unmarshal(raw["endpoints"], &s.Endpoints)
	if err != nil {
		return err
	}

	if s.Port == 0 {
		return fmt.Errorf("missing port")
	}
	return nil
}

func (e *Endpoint) UnmarshalJSON(data []byte) error {
	raw := make(map[string]json.RawMessage)
	err := json.Unmarshal(data, &raw)
	if err != nil {
		return err
	}

	err = json.Unmarshal(raw["name"], &e.Name)
	if err != nil {
		return fmt.Errorf("invalid name %s", raw["name"])
	}
	err = json.Unmarshal(raw["path"], &e.Path)
	if err != nil {
		return fmt.Errorf("invalid path %s", raw["path"])
	}
	err = json.Unmarshal(raw["actions"], &e.Actions)
	if err != nil {
		return err
	}

	if e.Name == "" {
		return fmt.Errorf("missing name")
	}
	if e.Path == "" {
		return fmt.Errorf("endpoint %s missing path", e.Name)
	}
	return nil
}

func (a *Action) UnmarshalJSON(data []byte) error {
	raw := make(map[string]json.RawMessage)
	err := json.Unmarshal(data, &raw)
	if err != nil {
		return err
	}

	err = json.Unmarshal(raw["method"], &a.Method)
	if err != nil {
		return fmt.Errorf("invalid method %s", raw["method"])
	}
	err = json.Unmarshal(raw["description"], &a.Description)
	if err != nil {
		return err
	}

	if a.Method == "" {
		return fmt.Errorf("missing method")
	}
	if a.Method != http.MethodGet && a.Method != http.MethodPost &&
		a.Method != http.MethodPut && a.Method != http.MethodDelete {
		return fmt.Errorf("unsupported method %s", a.Method)
	}
	return nil
}
