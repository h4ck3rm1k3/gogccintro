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
	fmt.Printf("CoOccurance report\n" ) // total per field
	
	// fields with the same count for cooccurance
	coccurance   := map[int][]string{}

	// for each field, get the coccurance
	for k,v := range( t.TheFields) {

		//fmt.Printf("\tCHECK FIELD '%s' -> '%#v'\n", k,v)
		
		var total int // total per field
		total = 0
		
		//v.Report(k)
		for k2,v2 := range( v.Fields) { // count with other fields

			//fmt.Printf("\tCHECK FIELD2 '%s' -> '%s'\n", k2,v2)
			
			key := fmt.Sprintf("\t\tCO '%s' -> '%s'", k,k2)
			coccurance[v2] = append(coccurance[v2], key)
			total = total + v2
		}

		fmt.Printf("\tFIELD: '%s' : %d\n", k, total ) // total per field

		// now compare with others
		for k2,_ := range( t.TheFields) {
			//if k2 not in v.Fields
			if k2 != k {

				// now swap them if k is less than k2
				a := k
				b := k2
				if strings.Compare(a, b) < 0 {
					c :=a
					a=b
					b=c
				}

				// now look up in the matrix from the beginning
				
				if count, ok := t.TheFields[a].Fields[b]; ! ok {
					fmt.Printf("\t\tNCO '%s'  '%s'\n", k,k2)
				} else {
					fmt.Printf("\t\tCOO '%s'  '%s' -> %d\n", k,k2,count)
				}
				
			}
		}
			
	}

	var a []int
	for k := range coccurance {
		a = append(a, k)
	}

	// 
	sort.Sort(sort.Reverse(sort.IntSlice(a)))
	for _, k := range a {
		for _, s := range coccurance[k] {
			fmt.Printf("COOCCURANCE : %s, %d\n", s, k)
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

func (t *CoOccuranceField) AppendTo(t2 *CoOccuranceField) {
	if t2.Fields == nil {
		t2.Fields = make(map [string] int)
	}	

	for k,v := range( t.Fields) {
		if val, ok := t2.Fields[k]; ok {
			t2.Fields[k] = val + v
		} else {
			t2.Fields[k]=v
		}		
	}	
}

func (t *CoOccurance) AppendTo(t2 *CoOccurance) {
	if t2.TheFields == nil {
		t2.TheFields = make(map [string] *CoOccuranceField)
	}

	//
	for k,v := range( t.TheFields) {
		var o * CoOccuranceField
		
		if val, ok := t2.TheFields[k]; ok {
			o = val

		} else {
			o = &CoOccuranceField{}
			t2.TheFields[k]=o
		}
		
		v.AppendTo(o)
	}
}
