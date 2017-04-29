package main
import (
	"fmt"
	"strings"
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
	for k,v := range( t.TheFields) {
		//fmt.Printf("\tkv '%s' -> '%s'\n", k,v)
		v.Report(k)
	}
}

func (t *CoOccurance) Fields(a string, b string) {

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
