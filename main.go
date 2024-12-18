package main

import (
	"fmt"
	"time"
)

type store struct {
	data map[string]string
}

func (s *store) set(key, value string) {
	s.data[key] = value
}

func (s *store) get(key string) string {
	if value, ok := s.data[key]; ok {
		return value
	}
	return "Data not found"
}

func (s *store) del(key string) {
	delete(s.data, key)
}

func main() {

	starttime := time.Now()
	s := &store{data: make(map[string]string)}
	s.set("foo", "boo")
	s.set("foo1", "boo1")
	s.set("foo2", "boo2")
	s.set("foo3", "boo3")
	s.set("foo4", "boo4")
	s.set("foo5", "boo5")
	endtime := time.Now()

	duration := endtime.Sub(starttime)
	fmt.Print(duration)
	fmt.Println(s.get("foo5"))
	s.del("foo3")
	fmt.Println(s.get("3"))
}
