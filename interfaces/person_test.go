package interfaces

import (
	"bytes"
	"testing"
)

func TestTheOldest(t *testing.T) {
	testData := []struct {
		name string
		args []Person
		want int
	}{
		{
			name: "Test #1",
			args: []Person{&Customer{age: 20}},
			want: 20,
		},
		{
			name: "Test #2",
			args: []Person{&Customer{age: 20}, &Employee{age: 25}, &Employee{age: 23}},
			want: 25,
		},
		{
			name: "Test #3",
			args: []Person{&Customer{age: 0}, &Employee{age: 1}, &Employee{age: 1}},
			want: 1,
		},
	}

	for _, td := range testData {
		t.Run(td.name, func(t *testing.T) {
			if got := TheOldest(td.args...); got != td.want {
				t.Errorf("TheOldest() = %v, wanted %v", got, td.want)
			}
		})
	}
}

func TestTheOldestObject(t *testing.T) {
	type args struct {
		args []interface{}
	}
	p1 := Customer{age: 20}
	p2 := Employee{age: 25}
	p3 := Employee{age: 23}

	var d1, d2 []interface{}
	d1 = append(d1, p1)
	d2 = append(d1, p1, p2, p3)

	testData := []struct {
		name string
		args args
		want interface{}
	}{
		{
			name: "Test #1",
			args: args{d1},
			want: p1,
		},
		{
			name: "Test #2",
			args: args{d2},
			want: p2,
		},
	}

	for _, td := range testData {
		t.Run(td.name, func(t *testing.T) {
			if got := TheOldestObject(td.args.args...); got != td.want {
				t.Errorf("TheOldestObject() = %v, wanted %v", got, td.want)
			}
		})
	}
}

func TestPrintStrings(t *testing.T) {
	var a1, a2, a3 []interface{}
	a1 = append(a1, 1, 1.1, "string")
	a2 = append(a2, 1, 2, 3, 4, true)
	a3 = append(a3, 1, true, false, "string1", "string2")

	type args struct {
		args []interface{}
	}

	testData := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test #1",
			args: args{a1},
			want: "string",
		},
		{
			name: "Test #2",
			args: args{a2},
			want: "",
		},
		{
			name: "Test #3",
			args: args{a3},
			want: "string1string2",
		},
	}

	for _, td := range testData {
		t.Run(td.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			PrintStrings(w, td.args.args...)
			if got := w.String(); got != td.want {
				t.Errorf("PrintStrings() = %v, want %v", got, td.want)
			}
		})
	}
}
