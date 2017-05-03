package ast

type ProgramData struct {
	Count         Stat
	NodeTypes     map[string]*SNodeType
	// parser stats
	Transitions   map[int]map[int]int
	States        map[int]int	
}

type  CoOccuranceField struct {
	Fields map [string] int
}
type  CoOccurance struct {
	TheFields map [string] *CoOccuranceField
}

// simple ast
type Stat struct {
	TheCount int
}

type SValue struct {
	Count Stat

}

type SAttributeValString struct {
	Count   Stat
	TheVals map[string]*SValue
}

type SAttributeVals struct {
	Count   Stat
	TheVals map[int]*SAttributeValString
}

type SAttributeNames struct {
	Count   Stat
	TheVals map[string]*SAttributeVals
}

type SAttributeVector struct {
	Count   Stat
	TheVals map[string]*SAttributeNames
}

type SNodeType struct {
	Count          Stat
	AttributeCount map[int]*SAttributeVector // Attribute County

	AttrNames map[string]int // simple count
	CoOccurance *CoOccurance
}

type TestNode struct {
	NodeType string
	NodeId   string
	Vals     map[string][]string
	//StrVals     map[string]string
	//IntVals     map[string]int
}
