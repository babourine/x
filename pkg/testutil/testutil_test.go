package testutil

import (
	"fmt"
	"testing"
)

func TestUnitM(t *testing.T) {
	for i := 100; i < 200; i++ {
		expected := fmt.Sprintf("i: `%d`", i)
		if v := M(&P{`i`: i}); v != expected {
			t.Errorf("fail: [%v] `%v` != `%v`", i, v, expected)
		}
	}
}

func TestUnitMNil(t *testing.T) {
	expected := fmt.Sprintf("[prefix]")
	if v := M(nil, `prefix`); v != expected {
		t.Errorf("fail: `%v` != `%v`", v, expected)
	}
}

func TestUnitMFull(t *testing.T) {
	expected := fmt.Sprintf("[prefix] abc: `xyz`, foo: `bar`")
	if v := M(&P{`foo`: `bar`, `abc`: `xyz`}, `prefix`); v != expected {
		t.Errorf("fail: `%v` != `%v`", v, expected)
	}
}

func TestUnitE(t *testing.T) {
	tt := &testing.T{}
	if E(tt, nil); tt.Failed() {
		return
	}
	t.Errorf("must fail")
}
