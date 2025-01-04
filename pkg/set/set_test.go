package set

import (
	"encoding/json"
	"testing"

	"gopkg.in/yaml.v3"
)

func TestUnitNew(t *testing.T) {
	s := New([]int{1, 2, 3})
	if l := s.Len(); l != 3 {
		t.Errorf("unexpected length: expect 3, got %v", l)
	}
}

func TestUnitNewEmpty(t *testing.T) {
	s := New([]int{})
	if l := s.Len(); l != 0 {
		t.Errorf("unexpected length: expect 0, got %v", l)
	}
}

func TestUnitAdd(t *testing.T) {
	s := New([]string{`1`, `2`, `3`})
	s.Add(`4`)
	if l := s.Len(); l != 4 {
		t.Errorf("unexpected length: expect 4, got %v", l)
	}
}

func TestUnitAddExisting(t *testing.T) {
	s := New([]byte{1, 2, 3})
	s.Add(3)
	if l := s.Len(); l != 3 {
		t.Errorf("unexpected length: expect 3, got %v", l)
	}
}

func TestUnitDelete(t *testing.T) {
	s := New([]uint32{1, 2, 3})
	s.Delete(4)
	if l := s.Len(); l != 3 {
		t.Errorf("unexpected length: expect 2, got %v", l)
	}
}

func TestUnitDeleteExisting(t *testing.T) {
	s := New([]int64{1, 2, 3})
	s.Delete(2)
	if l := s.Len(); l != 2 {
		t.Errorf("unexpected length: expect 2, got %v", l)
	}
}

func TestUnitClear(t *testing.T) {
	s := New([]int64{1, 2, 3})
	if l := s.Len(); l != 3 {
		t.Fatalf("unexpected length: expect 3, got %v", l)
	}
	s.Clear()
	if l := s.Len(); l != 0 {
		t.Errorf("unexpected length: expect 0, got %v", l)
	}
}

func TestUnitHas(t *testing.T) {
	s := New([]int64{1, 2, 3})
	if !s.Has(2) {
		t.Errorf("missing item")
	}
}

func TestUnitHasNot(t *testing.T) {
	s := New([]int64{1, 2, 3})
	if s.Has(4) {
		t.Fatalf("what? :)")
	}
}

func TestUnitUnmarshalJSON(t *testing.T) {
	s := &Set[int32]{}
	if err := json.Unmarshal([]byte(`[1, 2, 3]`), s); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}
	if l := s.Len(); l != 3 || !s.Has(2) {
		t.Errorf("unexpected set: %v", s)
	}
}

func TestUnitUnmarshalJSONError(t *testing.T) {
	s := &Set[int64]{}
	if err := json.Unmarshal([]byte(`[`), s); err == nil {
		t.Fatalf("missing expected error")
	}
}

func TestUnitUnmarshalYAML(t *testing.T) {
	s := &Set[int32]{}
	if err := yaml.Unmarshal([]byte(`[1, 2, 3]`), s); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}
	if l := s.Len(); l != 3 || !s.Has(2) {
		t.Errorf("unexpected set: %v", s)
	}
}

func TestUnitUnmarshalYAMLError(t *testing.T) {
	s := &Set[int64]{}
	if err := yaml.Unmarshal([]byte(`[`), s); err == nil {
		t.Fatalf("missing expected error")
	}
}

func TestUnitContainsSlice(t *testing.T) {
	s := New([]int64{1, 2, 3})
	n := New(s.Slice())

	if !s.Contains(n) {
		t.Errorf("%v != %v", s, n)
	}
}

func TestUnitNotContains(t *testing.T) {
	s := New([]int64{1, 2, 3})
	n := New([]int64{4, 5})

	if s.Contains(n) {
		t.Errorf("%v == %v", s, n)
	}
}

func TestUnitContainsAnyError(t *testing.T) {
	s := New([]int64{1, 2, 3})
	n := New([]int64{4, 5})

	if s.ContainsAny(n) {
		t.Errorf("%v == %v", s, n)
	}
}

func TestUnitMarshalJSON(t *testing.T) {
	s := New([]string{`foo`, `bar`, `xyz`, `abc`})

	if data, err := s.MarshalJSON(); err != nil || string(data) != `["abc","bar","foo","xyz"]` {
		t.Errorf("%v %v", string(data), err)
	}

}

func TestUnitMarshalYAML(t *testing.T) {
	s := New([]string{`foo`, `bar`, `xyz`, `abc`})

	if data, err := s.MarshalYAML(); err != nil || string(data) != "- abc\n- bar\n- foo\n- xyz\n" {
		t.Errorf("%v %v", string(data), err)
	}

}
