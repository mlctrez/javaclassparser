package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sort"
	"strings"

	"github.com/mlctrez/javaclassparser/parser"
)

func main() {

	// key of class name to

	deps := NewDependencyManager()
	deps.IncludePrefixes("com/foo")

	//deps.packageOnly = true
	//deps.depth = 3

	err := parser.Scan(parser.NewConfigFromArgs(), deps.WorkCallback)
	if err != nil {
		panic(err)
	}

	//deps.Debug()
	//deps.DumpDependencies()

}

type DependencyManager struct {
	prefixes    []string
	forward     map[string]map[string]uint16
	reverse     map[string]map[string]uint16
	packageOnly bool
	depth       int
}

func NewDependencyManager() *DependencyManager {
	d := &DependencyManager{}
	d.forward = make(map[string]map[string]uint16)
	d.reverse = make(map[string]map[string]uint16)
	return d
}

func (d *DependencyManager) IncludePrefixes(prefixes ...string) {
	d.prefixes = prefixes
}

func (d *DependencyManager) extractName(name string) string {

	if d.packageOnly&&d.depth>0 {
		panic("cannot set packageOnly and depth")
	}

	cn := parser.ExtractName(name)

	if d.packageOnly && strings.Contains(cn, "/") {
		cn = cn[0:strings.LastIndex(cn, "/")]
	}
	if d.depth != 0 {
		cnp := strings.Split(cn, "/")
		desired := d.depth
		if len(cnp) < desired {
			desired = len(cnp)
		}
		cn = strings.Join(cnp[0:desired], "/")
	}

	return cn
}

func (d *DependencyManager) AddDependency(from, to string) {
	AddToMap(d.forward, from, to)
	AddToMap(d.reverse, to, from)
}

func AddToMap(m map[string]map[string]uint16, from, to string) {
	var sm map[string]uint16
	var ok bool

	if sm, ok = m[from]; !ok {
		m[from] = make(map[string]uint16)
		sm = m[from]
	}

	sm[to]++
}

func (d *DependencyManager) included(name string) bool {
	for _, r := range d.prefixes {
		if strings.HasPrefix(name, r) {
			return true
		}
	}
	return false
}

func (d *DependencyManager) WorkCallback(work *parser.Work) {

	if d.packageOnly && d.depth>0 {
		panic("cannot set packageOnly and depth together")
	}

	fromName := d.extractName(work.Class.Name)
	work.Class.RefVisit(func(toName string) {
		if toName == "" {
			return
		}
		toName = d.extractName(toName)

		if fromName == toName {
			return
		}

		if d.included(fromName) && d.included(toName) {
			d.AddDependency(fromName, toName)
		}

	})
}

func (d *DependencyManager) Debug() {
	dumpMapValueCount("<--", d.forward)
	dumpMapValueCount("-->", d.reverse)
}

func (d *DependencyManager) Unique() []string {
	u := make(map[string]bool)
	for k := range d.forward {
		u[k] = true
	}
	for k := range d.reverse {
		u[k] = true
	}
	results := make([]string, 0)
	for k := range u {
		results = append(results, k)
	}
	sort.Strings(results)
	return results
}

func (d *DependencyManager) DumpDependencies() {
	dm := make(map[string]interface{})
	dm["_count_forward"] = len(d.forward)
	dm["_count_reverse"] = len(d.reverse)
	unique := d.Unique()
	dm["_unique"] = unique
	dm["_count_unique"] = len(unique)
	dm["forward"] = d.forward
	dm["reverse"] = d.reverse
	bytes, _ := json.MarshalIndent(dm, "", "  ")
	ioutil.WriteFile("deps.json", bytes, 0755)
}

func dumpMapValueCount(mapName string, m map[string]map[string]uint16) {
	var srt []string
	for k := range m {
		srt = append(srt, k)
	}
	sort.Strings(srt)
	for _, c := range srt {
		fmt.Printf("%03d %s %s\n", len(m[c]), mapName, c)
	}
}
