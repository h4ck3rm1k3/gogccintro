package main
import (
	"fmt"
	"strings"
	"sort"
)

type  CoOccuranceField struct {
	Fields map [string] int
}
func (t *CoOccuranceField) Report(a string) {
	for k,v := range( t.Fields) {
		fmt.Printf("\tCO '%s' -> '%s' %d\n", a,k,v)
	}	
}

func (t *CoOccuranceField) Field(b string) {
	if t.Fields == nil {
		t.Fields = make(map [string] int)
	}
	if val, ok := t.Fields[b]; ok {
		t.Fields[b] = val  + 1
	} else {
		t.Fields[b]=1
	}

}

type  CoOccurance struct {
	TheFields map [string] *CoOccuranceField
}

func (t *CoOccurance) Report() {

	n := map[int][]string{}
	
	for k,v := range( t.TheFields) {
		//fmt.Printf("\tkv '%s' -> '%s'\n", k,v)
		//v.Report(k)
		for k2,v2 := range( v.Fields) {
			//fmt.Printf("\tCO '%s' -> '%s' %d\n", k,k2,v2)
			key := fmt.Sprintf("\tCO '%s' -> '%s'", k,k2)
			n[v2] = append(n[v2], key)
		}	

	}

	var a []int
	for k := range n {
		a = append(a, k)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(a)))
	for _, k := range a {
		for _, s := range n[k] {
			fmt.Printf("%s, %d\n", s, k)
		}
	}
}

func (t *CoOccurance) Fields(a string, b string) {
	// skip the digit fields
	if (a[0]>= '0' && a[0] <= '9'){
		return
	}
	if (b[0]>= '0' && b[0] <= '9'){
		return
	}
	
	if strings.Compare(a, b) < 0 {
		c :=a
		a=b
		b=c
	}

	if t.TheFields == nil {
		t.TheFields = make(map [string] *CoOccuranceField)
	}
	if val, ok := t.TheFields[a]; ok {
		val.Field(b)
	} else {
		o := &CoOccuranceField{}
		t.TheFields[a]=o
		o.Field(b)
	}

}
