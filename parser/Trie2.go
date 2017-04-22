package main;

import (
	//"github.com/anknown/ahocorasick"
	"github.com/cloudflare/ahocorasick"
	//"bytes"
	"fmt"
)

type Machine3 struct {
	Machine *ahocorasick.Matcher
	Dict[] string
	MaxLen int
}

//type [][]rune DictType

func GenericTrieFields3(
	dict []string,
	maxlen int,
) (*Machine3){
	fmt.Printf("Build %s", dict)
	m :=  ahocorasick.NewStringMatcher(dict)
	return &Machine3{
		Machine: m,
		Dict: dict,
		MaxLen: maxlen,
	}
}

func NewTrieFields3(in []string) (*Machine3){
	dict := []string{}
	maxlen := 0
	// process the fields
	for _,l := range(in) {
		ln := len(l)
		if ln > maxlen {
			maxlen = ln
		}

		if len(l) == 3 {
			fmt.Printf("adding2 %s\n", l + " :")
			dict = append(dict, l + " :")
		} else if len(l) == 2 {
			fmt.Printf("adding3 %s\n", l + "  :")
			dict = append(dict, l + "  :")
		} else if len(l) >= 4 {
			fmt.Printf("adding2 %s\n", l + ":")
			dict = append(dict, l + ":")
		} else {
			panic(l)
		}
	}
	return GenericTrieFields3(dict, maxlen)
}

func NewTrie3(in []string) (*Machine3){
	maxlen := 0
	dict := []string{}

	for _,l := range(in) {
		ln := len(l)
		if ln > maxlen {
			maxlen = ln
		}
		fmt.Printf("adding %s\n", l)
		dict = append(dict, l)
	}
	return GenericTrieFields3(dict,maxlen)
}

func (t *Machine3) MultiPatternSearch(s string) ([]int){
	fmt.Printf("looking for %s in %s", s, t.Dict)
	return t.Machine.Match([]byte(s))
}
func (t *Machine3) ExactSearch(s string) (bool){
	r := t.Machine.Match([]byte(s))

	if (len(r) == 1 ){
		fmt.Printf("Exact found: %s = %s\n", r, s)
		return true
	}
	fmt.Printf("Not Exact %s\n", r)
	return false
}
func (t*Machine3) PrintOutput(){
	fmt.Printf("Machine %s", t)
}
