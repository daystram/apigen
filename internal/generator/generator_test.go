package generator_test

import (
	"path/filepath"
	"testing"

	"github.com/daystram/apigen/internal/definition"
	"github.com/daystram/apigen/internal/generator"
)

var cases = map[string]struct {
	svc         definition.Service
	pkg         string
	expectFiles []string
	expectFail  bool
}{
	"ok1": {
		svc: definition.Service{
			Host:      "localhost",
			Port:      8080,
			Endpoints: []definition.Endpoint{},
		},
		pkg:         "github.com/daystram/apigen",
		expectFiles: []string{"main.go", "controllers/init.go"},
		expectFail:  false,
	},
	"ok2": {
		svc: definition.Service{
			Host: "localhost",
			Port: 8080,
			Endpoints: []definition.Endpoint{{
				Name:    "User",
				Path:    "/user",
				Actions: []definition.Action{},
			}},
		},
		pkg:        "github.com/daystram/apigen",
		expectFail: false,
	},
	"ok3": {
		svc: definition.Service{
			Host: "localhost",
			Port: 8080,
			Endpoints: []definition.Endpoint{{
				Name: "User",
				Path: "/user",
				Actions: []definition.Action{{
					Method:      "GET",
					Description: "Get all users",
				}, {
					Method:      "POST",
					Description: "Add user",
				}},
			}, {
				Name: "Item",
				Path: "/item/:id",
				Actions: []definition.Action{{
					Method:      "GET",
					Description: "Get item",
				}, {
					Method:      "POST",
					Description: "Add item",
				}, {
					Method:      "PUT",
					Description: "Edit item",
				}},
			}},
		},
		pkg:        "github.com/daystram/apigen",
		expectFail: false,
	},
	"bad-port": {
		svc: definition.Service{
			Host:      "localhost",
			Endpoints: []definition.Endpoint{},
		},
		pkg:        "github.com/daystram/apigen",
		expectFail: true,
	},
	"bad-endpoint-name": {
		svc: definition.Service{
			Host: "localhost",
			Port: 8080,
			Endpoints: []definition.Endpoint{{
				Name: "",
				Path: "/user",
			}},
		},
		pkg:        "github.com/daystram/apigen",
		expectFail: true,
	},
	"bad-endpoint-action-method": {
		svc: definition.Service{
			Host: "localhost",
			Port: 8080,
			Endpoints: []definition.Endpoint{{
				Name: "User",
				Path: "/user",
				Actions: []definition.Action{{
					Method:      "PATCH",
					Description: "Get all users",
				}},
			}},
		},
		pkg:        "github.com/daystram/apigen",
		expectFail: true,
	},
}

func TestGenerator(t *testing.T) {
	t.Parallel()
	for name, tt := range cases {
		name, tt := name, tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			fg, err := generator.Generate(tt.svc, tt.pkg)

			if !tt.expectFail && err != nil {
				t.Errorf("unexpected error: %s", err)
			}
			if tt.expectFail && err == nil {
				t.Errorf("expected error")
			}

			if len(tt.expectFiles) != 0 {
				if len(fg) != len(tt.expectFiles) {
					t.Errorf("incorrect number of generated files")
				}
				fc := make(map[string]bool)
				for _, f := range tt.expectFiles {
					fc[f] = false
				}
				for _, f := range fg {
					name := filepath.Join(f.Dir, f.Name)
					if found, exist := fc[name]; found {
						t.Errorf("duplicate file: %s", name)
					} else if !exist {
						t.Errorf("unexpected file: %s", name)
					}
					fc[name] = true
				}
			}
		})
	}
}
