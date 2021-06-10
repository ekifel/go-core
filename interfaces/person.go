package interfaces

import (
	"io"
	"sort"
)

type Person interface {
	Age() int
}

type Employee struct {
	age int
}

type Customer struct {
	age int
}

func (c *Customer) Age() int {
	return c.age
}

func (e *Employee) Age() int {
	return e.age
}

func TheOldest(p ...Person) int {
	sort.Slice(p, func(i, j int) bool { return p[i].Age() >= p[j].Age() })
	return p[0].Age()
}

func TheOldestObject(args ...interface{}) interface{} {
	max := 0
	var r interface{}
	for _, p := range args {
		switch t := p.(type) {
		case Employee:
			if t.age > max {
				max = t.age
				r = t
			}
		case Customer:
			if t.age > max {
				max = t.age
				r = t
			}
		}
	}

	return r
}

func PrintStrings(w io.Writer, args ...interface{}) {
	for _, v := range args {
		if p, ok := v.(string); ok {
			w.Write([]byte(p))
		}
	}
}
