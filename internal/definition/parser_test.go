package definition_test

import (
	"testing"

	"github.com/daystram/apigen/internal/definition"
)

var cases = map[string]struct {
	file       string
	expectFail bool
}{
	"ok1": {
		file:       "testdata/ok1.json",
		expectFail: false,
	},
	"ok2": {
		file:       "testdata/ok2.json",
		expectFail: false,
	},
	"empty1": {
		file:       "testdata/empty1.json",
		expectFail: true,
	},
	"empty2": {
		file:       "testdata/empty2.json",
		expectFail: true,
	},
	"no-file": {
		file:       "",
		expectFail: true,
	},
	"bad-json": {
		file:       "testdata/bad-json.json",
		expectFail: true,
	},
	"bad-port1": {
		file:       "testdata/bad-port1.json",
		expectFail: true,
	},
	"bad-port2": {
		file:       "testdata/bad-port2.json",
		expectFail: true,
	},
	"bad-host": {
		file:       "testdata/bad-host.json",
		expectFail: true,
	},
	"bad-endpoint-path1": {
		file:       "testdata/bad-endpoint-path1.json",
		expectFail: true,
	},
	"bad-endpoint-path2": {
		file:       "testdata/bad-endpoint-path2.json",
		expectFail: true,
	},
	"bad-endpoint-name1": {
		file:       "testdata/bad-endpoint-name1.json",
		expectFail: true,
	},
	"bad-endpoint-name2": {
		file:       "testdata/bad-endpoint-name2.json",
		expectFail: true,
	},
	"bad-endpoint-action-method1": {
		file:       "testdata/bad-endpoint-action-method1.json",
		expectFail: true,
	},
	"bad-endpoint-action-method2": {
		file:       "testdata/bad-endpoint-action-method2.json",
		expectFail: true,
	},
	"bad-endpoint-action-method3": {
		file:       "testdata/bad-endpoint-action-method3.json",
		expectFail: true,
	},
	"bad-endpoint-action-description": {
		file:       "testdata/bad-endpoint-action-description.json",
		expectFail: true,
	},
}

func TestParser(t *testing.T) {
	t.Parallel()
	for name, tt := range cases {
		name, tt := name, tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			p := definition.NewParser()
			_, err := p.ParseFile(tt.file)

			if !tt.expectFail && err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if tt.expectFail && err == nil {
				t.Errorf("expected error")
			}
		})
	}
}
