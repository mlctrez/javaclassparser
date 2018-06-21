package main

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"

	"github.com/mlctrez/javaclassparser/cpool"
	"github.com/mlctrez/javaclassparser/ioutil"
	"github.com/mlctrez/javaclassparser/parser"
)

func main() {
	config := &parser.Config{Archive: "sample.zip"}

	// key of class name to
	deps := NewDependencyManager()

	reader := ioutil.MustOpenZipReader("sample.zip")
	parser.Parse(config, reader, deps.WorkCallback)

	deps.Debug()
	deps.DumpDependencies()

}

type DependencyManager struct {
	ExcludedPrefixes []string
	includeOnly      string
	forward          map[string]map[string]bool
	reverse          map[string]map[string]bool
}

func NewDependencyManager() *DependencyManager {
	d := &DependencyManager{}
	d.forward = make(map[string]map[string]bool)
	d.reverse = make(map[string]map[string]bool)
	return d
}

func (d *DependencyManager) ExcludePrefixes(prefixes ...string) {
	d.ExcludedPrefixes = prefixes
}

func (d *DependencyManager) IncludeOnly(name string) {
	d.includeOnly = name

}

func ExtractClassName(name string) string {
	cn := strings.Split(name, " ")[0]
	cn = strings.Split(cn, "$")[0]
	cn = strings.TrimPrefix(cn, "[L")
	cn = strings.TrimSuffix(cn, ";")
	return cn
}

func (d *DependencyManager) AddDependency(from, to string) {

	// no restriction
	if d.includeOnly == "" {
		AddToMap(d.forward, from, to)
		AddToMap(d.reverse, to, from)
	}

	// add only if from is included
	if from == d.includeOnly || to == d.includeOnly{
		AddToMap(d.forward, from, to)
		AddToMap(d.reverse, to, from)
		return
	}
}

func AddToMap(m map[string]map[string]bool, from, to string) {
	var sm map[string]bool
	var ok bool

	if sm, ok = m[from]; !ok {
		m[from] = make(map[string]bool)
		sm = m[from]
	}

	sm[to] = true
}

func (d *DependencyManager) WorkCallback(work *parser.Work) {

	className := ExtractClassName(work.Class.Name)
	work.Class.Visit(func(i interface{}) {
		var ref string
		if mr, ok := i.(*cpool.ConstantMethodrefInfo); ok {
			ref = mr.String()
		}
		if mr, ok := i.(*cpool.ConstantInterfaceMethodrefInfo); ok {
			ref = mr.String()
		}
		if ref == "" {
			return
		}
		ref = ExtractClassName(ref)
		for _, r := range d.ExcludedPrefixes {
			if strings.HasPrefix(ref, r) {
				return
			}
		}
		d.AddDependency(className, ref)
	})
}

func (d *DependencyManager) Debug() {
	dumpMapValueCount("<--", d.forward)
	dumpMapValueCount("-->", d.reverse)
}

func (d *DependencyManager) DumpDependencies() {
	bytes, _ := json.MarshalIndent(d.forward, "", "  ")
	fmt.Println(string(bytes))
}

func dumpMapValueCount(mapName string, m map[string]map[string]bool) {
	var srt []string
	for k := range m {
		srt = append(srt, k)
	}
	sort.Strings(srt)
	for _, c := range srt {
		fmt.Printf("%03d %s %s\n", len(m[c]), mapName, c)
	}
}
