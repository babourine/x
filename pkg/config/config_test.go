package config

import (
	"path"
	"path/filepath"
	"regexp"
	"runtime"
	"testing"

	. "github.com/babourine/x/pkg/testutil"
)

type MyTest struct {
	M map[string]string `yaml:"bar"`
}

type MyTest2 struct {
	I int
}

func (x *MyTest2) Init() error {
	x.I = 123
	return nil
}

func TestUnitConfig(t *testing.T) {

	testFunc := func(configFile string, fn func(string, interface{}) error) {
		var d int
		if err := fn(configFile, &d); err == nil || err.Error() != `open `+configFile+`: no such file or directory` {
			E(t, &P{`d`: d, `err`: err}, configFile)
		}
	}

	testFunc(`unknown.yaml`, LoadYAML)
	testFunc(`unknown.json`, LoadJSON)

}

func TestUnitConfigBadYAML(t *testing.T) {

	var d int

	re := regexp.MustCompile(`\s+`)

	if err := LoadYAML(path.Join(testDir(), `bad.yaml`), &d); err == nil {
		E(t, &P{`d`: d, `err`: err})
	} else {
		if e := re.ReplaceAllString(err.Error(), ` `); e != "yaml: unmarshal errors: line 1: cannot unmarshal !!str `blob` into int" {
			E(t, &P{`e`: e})
		}
	}

}

func TestUnitConfigYAMLWithoutInit(t *testing.T) {
	var x MyTest
	if err := LoadYAML(path.Join(testDir(), `foo.yaml`), &x); err != nil {
		F(t, &P{`err`: err, `x`: x})
	}

	if v, ok := x.M[`some`]; !ok || v != `text` {
		E(t, &P{`ok`: ok, `x`: x})
	}
}

func TestUnitConfigYAMLWithInit(t *testing.T) {
	var x MyTest2
	if err := LoadYAML(path.Join(testDir(), `foo.yaml`), &x); err != nil || x.I != 123 {
		E(t, &P{`err`: err, `x`: x})
	}
}

func testDir() string {
	_, filename, _, _ := runtime.Caller(0)
	return path.Join(filepath.Dir(filename), `/../../test/config`)
}
