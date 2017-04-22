package main;

import (
	"github.com/anknown/ahocorasick"
	"bytes"
	"fmt"
)

type Machine2 struct {
	Machine *goahocorasick.Machine
	MaxLen int
}

//type [][]rune DictType

func GenericTrieFields2(
	dict [][]rune,
	maxlen int,
) (*Machine2){
	fmt.Printf("Build %s", dict)
	m := new(goahocorasick.Machine)
	if err := m.Build(dict); err != nil {
		fmt.Println(err)
		return nil
	}
	return &Machine2{
		Machine: m,
		MaxLen: maxlen,
	}

}

func NewTrieFields2(in []string) (*Machine2){
	dict := [][]rune{}
	maxlen := 0
	// process the fields
	for _,l := range(in) {
		ln := len(l)
		if ln > maxlen {
			maxlen = ln
		}
		fmt.Printf("adding2 %s\n", l)
		if len(l) == 3 {

			dict = append(dict, bytes.Runes( []byte(l + " :")))
		} else if len(l) == 4 {
			dict = append(dict, bytes.Runes( []byte(l + ":")))
		}
	}
	return GenericTrieFields2(dict, maxlen)
}

func NewTrie2(in []string) (*Machine2){
	maxlen := 0
	dict := [][]rune{}
	for _,l := range(in) {
		ln := len(l)
		if ln > maxlen {
			maxlen = ln
		}
		fmt.Printf("adding %s\n", l)
		dict = append(dict, bytes.Runes( []byte(l)))
	}
	return GenericTrieFields2(dict,maxlen)
}
