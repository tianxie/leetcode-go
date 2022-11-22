package main

import (
	"encoding/json"
	"log"
	"os"
	"path"
	"text/template"
)

func main() {
	data, err := os.ReadFile("problem.json")
	check(err)

	var p Problem
	err = json.Unmarshal(data, &p)
	check(err)

	if _, err := os.Stat(p.Package); os.IsNotExist(err) {
		err := os.Mkdir(p.Package, 0750)
		check(err)
	}

	genSolution(p)
	genTest(p)
}

type Problem struct {
	Package string
	Func    string
	URL     string
}

func genSolution(p Problem) {
	const tpl = `// {{ .URL }}
package {{ .Package }}

type solution interface {
	{{ .Func }}
}

type solution1 struct{}

func (s1 solution1) {{ .Func }} {
	return nil
}

type solution2 struct{}

func (s2 solution2) {{ .Func }} {
	return nil
}`

	t, err := template.New("solution.go").Parse(tpl)
	check(err)

	name := path.Join(p.Package, t.Name())
	if _, err := os.Stat(name); os.IsNotExist(err) {
		f, err := os.Create(name)
		check(err)

		err = t.Execute(f, p)
		check(err)
	} else {
		log.Println(name, "existed")
	}

}

func genTest(p Problem) {
	const tpl = `// {{ .URL }}
package {{ .Package }}

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	var tests = []struct {
		
	}{
		{},
		{},
	}

	s := solution1{}
	for _, test := range tests {
		if got := s; !reflect.DeepEqual(got, test.want) {
			t.Errorf("got %q, want %q", got, test.want)
		}
	}
}`

	t, err := template.New("solution_test.go").Parse(tpl)
	check(err)

	name := path.Join(p.Package, t.Name())
	if _, err := os.Stat(name); os.IsNotExist(err) {
		f, err := os.Create(name)
		check(err)

		err = t.Execute(f, p)
		check(err)
	} else {
		log.Println(name, "existed")
	}
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
