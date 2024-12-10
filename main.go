package main

import "fmt"

type store struct {
	data map[string]string
}

func (s *store) set(key, value string) {
	s.data[key] = value
}

func (s *store) get(key string) string {
	return s.data[key]
}

func main() {

	s := store{data: make(map[string]string)}
	s.set("foo", "boo")
	fmt.Println(s.get("foo"))
}
