package definition

import (
	"encoding/json"
	"io"
	"os"
)

type Parser struct{}

func NewParser() *Parser {
	return &Parser{}
}

func (p *Parser) Parse(r io.Reader) (Service, error) {
	dec := json.NewDecoder(r)
	var s Service
	if err := dec.Decode(&s); err != nil {
		return Service{}, err
	}
	return s, nil
}

func (p *Parser) ParseFile(path string) (Service, error) {
	f, err := os.Open(path)
	if err != nil {
		return Service{}, err
	}
	defer f.Close()
	return p.Parse(f)
}
