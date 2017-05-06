package main


import (
	"fmt"
	"sort"
//	"strconv"
	//"strings"
//	"io/ioutil"
	//"bytes"

)
/*
a contains b if all the fields of b are inside of a.

outer shell has the most fields, inner has the least.

b is a subset of a

http://stackoverflow.com/questions/10368218/venn-diagram-drawing-algorithms

    the dual graph, which is the Boolean lattice of the intersections of the sets.

idea:
create binary tree with a level for each attribute if the attribute exists or not?

group attributes with same counts on one level.

find disjoint attributes and groups, ones that never occur together.

1. scan from top to bottom in terms of count

2. if the count occurs the same create a group.

3. if the field has less count
    3.1.  look at the other fields it occurs with 
       compare the total fields a and b to the co occurance counts, so total(a) - cooccurance(a,b), total(b) - cooccurance(a,b), total(a) - total(b).

   look at what fields never occur together. See what fields they always occur with. Find the largest groups of disjoint groups. 


type : field(a), field(b), cooccurance(a,b)
total : field(a), field(b), cooccurance(a,b)

if cooccurance(a,c)=0 : for each cooccurance(c,d) and cooccurance(a,b)
       find where cooccurance(d,b)=0.

*/

func (t *SNodeType) ReportDisjoint(nodetype string) {

	fmt.Printf("ReportDisjoint:\t\tlen attrcount %d:\n", len(t.AttributeCount))

	// values to string
	n := map[int][]string{}
	var a []int

	for k, v := range t.AttrNames {
		n[v] = append(n[v], k)
	}

	for k := range n {
		a = append(a, k)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(a)))

	var last int
	last = 0 
	for _, k := range a {
		
		if last > 0 {
			fmt.Printf("Group %d Diff :-%d\n", k, last-k)
		} else {
			fmt.Printf("Group :%d\n", k)
		}
		
		for _, s := range n[k] {
			fmt.Printf("\t %s, %d\n", s, k)

			// make sure each of these are cooccuring with each other.
		}
		
		last = k
	}

}
