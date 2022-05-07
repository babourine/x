package set

import (
	"encoding/json"
)

var empty struct{}

type Set[T comparable] map[T]struct{}

func New[T comparable](items []T) *Set[T] {
	l := Set[T](make(map[T]struct{}))
	for _, item := range items {
		l[item] = empty
	}
	return &l
}

func (s *Set[T]) Add(item T) {
	(*s)[item] = empty
}

func (s *Set[T]) Delete(item T) {
	delete(*s, item)
}

func (s *Set[T]) Has(item T) bool {
	_, found := (*s)[item]
	return found
}

func (s *Set[T]) Len() int {
	return len(*s)
}

func (s *Set[T]) Clear() {
	*s = make(map[T]struct{})
}

func (s *Set[T]) UnmarshalJSON(data []byte) error {

	temp := make([]T, 0, 10)

	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	items := New(temp)
	*s = *items

	return nil

}

func (s *Set[T]) UnmarshalYAML(unmarshal func(interface{}) error) error {

	temp := make([]T, 0, 10)

	if err := unmarshal(&temp); err != nil {
		return err
	}

	items := New(temp)
	*s = *items

	return nil

}
