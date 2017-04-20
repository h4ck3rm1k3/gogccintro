package main

import (
	//"github.com/anknown/ahocorasick"
	"bytes"
	"fmt"
	"strings"
)

type AttrNames map[string]int

type ParserGlobal struct {
	StateLookup     map[Token]string
	Attrnames       AttrNames
	FieldMachine    *Machine
	NotesMachine    *Machine
	AccessMachine   *Machine
	OperatorMachine *Machine
	SpecMachine     *Machine
	LinkMachine     *Machine
	DebugLevel int
}
type GccNodeTest struct {
	Filename string
	Buffer   string
	Globals  *ParserGlobal

	//
	Nodeids   map[string]int
	Nodetypes map[string]map[string]int
}
type ParserInstance struct {
	State      Token
	Pos        int
	Token      string
	NodeId     string
	NodeType   string
	AttrName   string
	AttrValues []string
	C          rune // current rune
	Parser     *GccNodeTest

}

func NewParser() *ParserGlobal {

	lookup := map[Token]string{
		START:            "START",
		AT:               "AT",
		NODEID:           "NODEID",
		EXPECT_NODE_TYPE: "EXPECT_NODE_TYPE",
		NODETYPE:         "NODETYPE",
		EXPECT_ATTRNAME:  "EXPECT_ATTRNAME",
		ATTRNAME:         "ATTRNAME",
		EXPECT_ATTRVALUE: "EXPECT_ATTRVALUE",
		ATTRVALUE:        "ATTRVALUE",
		AFTERATTRVALUE:   "AFTERATTRVALUE",
		AFTERATTRVALUE2:  "AFTERATTRVALUE2",
		NEWLINE:          "NEWLINE",
		NOTEVALUE:        "NOTEVALUE",
		OPERATORVALUE:    "OPERATORVALUE",
	}

	return &ParserGlobal{
		DebugLevel : 0,
		StateLookup:     lookup,
		Attrnames:       make(map[string]int),
		FieldMachine:    NewTrieFields(Fields[:]),
		NotesMachine:    NewTrie(Notes[:]),
		AccessMachine:   NewTrie(Access[:]),
		OperatorMachine: NewTrie(Operators[:]),
		SpecMachine:     NewTrie(Spec[:]),
		LinkMachine:     NewTrie(Link[:]),
	}
}


func (t *GccNodeTest) Init(Filename string, Globals *ParserGlobal) {
	fmt.Printf("init %s\n", Filename)
	t.Filename = Filename
	t.Globals = Globals

	t.Nodeids = make(map[string]int)
	t.Nodetypes = make(map[string]map[string]int)
}

type Token int

const (
	START Token = iota
	AT
	NODEID
	EXPECT_NODE_TYPE
	NODETYPE
	EXPECT_ATTRNAME
	ATTRNAME
	EXPECT_ATTRVALUE
	ATTRVALUE
	AFTERATTRVALUE
	AFTERATTRVALUE2
	NEWLINE

	// Note: ...
	NOTEVALUE
	// Note : operator ...
	OPERATORVALUE
)

func isIdent(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || ch == '_'
}

func isDigit(ch rune) bool {
	return (ch >= '0' && ch <= '9')
}

func (t *GccNodeTest) CheckName(machine *Machine, value string) bool {
	return machine.Machine.ExactSearch(bytes.Runes([]byte(value))) != nil
}

func (t *GccNodeTest) CheckFieldName(field string) bool {
	return t.CheckName(t.Globals.FieldMachine, field)
}

func (t *GccNodeTest) NodeId(nodeid string) {
	t.Nodeids[nodeid] = 1
}

func (t *GccNodeTest) NodeType(nodetype string, nodeid string) {
	if val, ok := t.Nodetypes[nodetype]; ok {
		val[nodeid] = 1
	} else {

		t.Nodetypes[nodetype] = make(map[string]int)
		t.Nodetypes[nodetype][nodeid] = 1
	}

}

func (t *GccNodeTest) AttrValue(field string) {
}

func (t *GccNodeTest) AttrName(field string) {
}


func (p *ParserInstance) Add() {
	p.Token = p.Token + string(p.C)
}

func (p *ParserInstance) NextState(val bool, nextState Token) {
	if val {
		p.State = nextState
	}
}

func (p *ParserInstance) SetState(nextState Token) {
	p.State = nextState
}

func (p *ParserInstance) Check(m *Machine) bool {
	if len(p.Token) <= m.MaxLen {
		if p.Parser.CheckName(m, p.Token) {
			p.Add()
			return true
		} else {
			return false
		}
	} else {
		panic(fmt.Sprintf("error operator %s", p.Token))
	}

}

func (p *ParserInstance) ConsumeToken() {
	if len(p.Token) > 1 {
		p.Token = ""
	}
}


func (p *ParserInstance) START() {
	if p.C == '@' {
		p.State = AT
		p.ConsumeToken()
	}
}

// AT
func (p *ParserInstance) AT() {
	if isDigit(p.C) {
		p.SetState(NODEID)
		p.Add()
	}
}

func (p *ParserInstance) NODEID() {
	//else if p.State == NODEID {
	if isDigit(p.C) {
		p.Add()
	} else if p.C == ' ' {
		p.NodeId = p.Token
		p.Parser.NodeId(p.NodeId)
		//fmt.Printf("found node id %s\n", p.NodeId)
		p.ConsumeToken()
		p.SetState(EXPECT_NODE_TYPE)
	}
}

func (p *ParserInstance) EXPECT_NODE_TYPE() {
	if p.C == ' ' {
		p.ConsumeToken()
	} else {
		p.Add()
		p.SetState(NODETYPE)
	}
}

func (p *ParserInstance) NODETYPE() {
	if p.C == ' ' {
		p.NodeType = p.Token
		//fmt.Printf("found node type %s %s\n", p.NodeId, p.NodeType)
		p.Parser.NodeType(p.NodeType, p.NodeId)
		p.ConsumeToken()
		p.SetState(EXPECT_ATTRNAME)
	} else {
		p.Add()
	}
}

func (p *ParserInstance) EXPECT_ATTRNAME() {
	if p.C == ' ' {
		p.ConsumeToken()
	} else {
		p.Add()
		p.SetState(ATTRNAME)
	}
}

func (p *ParserInstance) ATTRNAME() {
	if p.C == ' ' {
		if p.Token == "op" {
			p.Add()
		} else {
			//size := utf8.RuneCountInString(p.Token)
			if strings.HasSuffix(p.Token, ":") {
				if p.Parser.CheckFieldName(p.Token) {
					p.AttrName = p.Token
					p.ConsumeToken()

					//fmt.Printf("found p.AttrName %s %s %s\n", p.NodeId, p.NodeType, attrname)
					p.Parser.Globals.Attrnames[p.AttrName] = p.Parser.Globals.Attrnames[p.AttrName] + 1
					p.AttrValues = make([]string, 0)
					p.SetState(EXPECT_ATTRVALUE)
				} else {
					p.Add()
					// it is not a attr name, but a value, so switch
					p.SetState(ATTRVALUE)
				}
			} else {
				p.Add()
			}
		}
	} else {
		p.Add()
	}
}

func (p *ParserInstance) EXPECT_ATTRVALUE() {

	if p.C == ' ' {
		p.ConsumeToken()
	} else {
		p.SetState(ATTRVALUE)
		p.Add()
	}
}



func (p *ParserInstance) ATTRVALUE() {
	//p.Debug()
	l := len(p.AttrValues) // size of array
	var v string

	if l > 0 {
		v = p.AttrValues[l-1] // current p.Token
	} else {
		v = "<null>"
	}
	p.Debug2(fmt.Sprintf("check len %d\n", l))
	p.Debug2(fmt.Sprintf("check last %s\n", v))
	p.Debug2(fmt.Sprintf("check values %s\n", p.AttrValues))

	if (p.C == ' ')||(p.C == '\n') { // ws separator

		if strings.HasPrefix(p.Token, "@") {
			// prefix
			p.AttrValues = append(p.AttrValues, p.Token)
			//fmt.Printf("found attrvalue @ %s %s %s %s\n", p.NodeId, p.NodeType,p.AttrName,						p.AttrValues)
			p.ConsumeToken()
			p.SetState(AFTERATTRVALUE2)
		} else if strings.HasSuffix(p.Token, ":") {
			if p.Parser.CheckFieldName(p.Token) {
				// it is an attribute name
				p.AttrName = p.Token
				//fmt.Printf("found p.AttrName %s %s %s\n", p.NodeId, p.NodeType, p.AttrName)
				p.Parser.Globals.Attrnames[p.AttrName] = p.Parser.Globals.Attrnames[p.AttrName] + 1
				p.AttrValues = make([]string, 0)
				p.ConsumeToken()
				p.SetState(EXPECT_ATTRVALUE)
			} else {
				p.Add()
			}
		} else {
			if p.AttrName == "note:" {
				//p.Debug()
				fmt.Printf("in note operator \n")
				if l > 0 {
					if v == "operator" {
						p.Debug2(fmt.Sprintf("check note operator %s \n", p.Token))
						p.AttrValues = append(p.AttrValues, p.Token)
						p.ConsumeToken()
						// p.State on p.State
						p.SetState(OPERATORVALUE)
					} else {
						p.Debug2(fmt.Sprintf("check note operator2 %s\n", p.Token))
						p.Add()
						p.SetState(NOTEVALUE)
					}

				} else { // first work
					p.AttrValues = append(p.AttrValues, p.Token)
					p.Debug2(fmt.Sprintf(fmt.Sprintf("found first note %s %s %s %s\n", p.NodeId, p.NodeType,p.AttrName,								p.AttrValues)))
					p.ConsumeToken()
					if p.C != ' ' {
						p.Add()
					}
					p.SetState(AFTERATTRVALUE)
				}
				
			} else { // not note:
				//attrvalue = p.Token
				p.AttrValues = append(p.AttrValues, p.Token)
				//fmt.Printf("found attrvalue %s %s %s %s\n", p.NodeId, p.NodeType,p.AttrName,							p.AttrValues)
				p.ConsumeToken()
				p.SetState(AFTERATTRVALUE)
			}
		}
	} else {
		p.Add()
	}
}

func (p *ParserInstance) AFTERATTRVALUE() {

	if p.C == '\n' {
		p.SetState(NEWLINE)
	} else if p.C == ' ' {
		p.SetState(AFTERATTRVALUE)
	} else {

		p.Debug2(fmt.Sprintf("Attrname: '%s'\t",p.AttrName))
		p.Debug2(fmt.Sprintf("AttrValues: '%v'\n",p.AttrValues))
		// the next could be a name of an attr or a value
		
		if (p.AttrName == "note:") {
			p.SetState(ATTRVALUE)
			p.Add()
		} else {
			p.SetState(ATTRNAME)
		}
	}
}

func (p *ParserInstance) AFTERATTRVALUE2() {

	if p.C == '\n' {
		p.SetState(NEWLINE)
	} else if p.C == ' ' {
		p.SetState(AFTERATTRVALUE2)
	} else {
		p.SetState(ATTRNAME)
		p.Add()
	}
}

func (p *ParserInstance) OPERATORVALUE() {
	// one of Operators
	// max length 12
	//p.Debug()
	p.NextState(p.Check(p.Parser.Globals.OperatorMachine), AFTERATTRVALUE)
}

func (p *ParserInstance) NOTEVALUE() {

	//one of
	//Notes
	// max length 16
	p.Add()

	//p.Debug()

	if len(p.Token) <= p.Parser.Globals.OperatorMachine.MaxLen {
		p.Parser.CheckName(p.Parser.Globals.OperatorMachine, p.Token)
		p.Add()
	} else {
		panic(fmt.Sprintf("error note '%s'", p.Token))
	}
}

func (p *ParserInstance) NEWLINE() {
	// new p.Statement
	if p.C == '@' {
		p.SetState(NODEID)
	} else {
		p.SetState(EXPECT_ATTRNAME)
	}
}

type ParseFunc func()

func (p *ParserInstance) LenCheck() {
	if len(p.Token) > 250 {
		fmt.Printf("Long P.Token %d: %s\n", p.State, p.Token)
		panic("toolong")
	}
	if p.Pos > 1 {
		if p.Pos%100000 == 0 {
			fmt.Printf("Proc %d %s %d\n", p.Pos, p.Token, p.State)
		}
	}
	p.Pos = p.Pos + 1
}

func (t *GccNodeTest) Parse() *int {
	//r := ni
	//fmt.Printf("machine %s\n",t.Machine)

	p := ParserInstance{
		Parser: t,
	}

	var Funcs = map[Token]ParseFunc{
		AFTERATTRVALUE2:  p.AFTERATTRVALUE2,
		AFTERATTRVALUE:   p.AFTERATTRVALUE,
		AT:               p.AT,
		ATTRNAME:         p.ATTRNAME,
		ATTRVALUE:        p.ATTRVALUE,
		EXPECT_ATTRNAME:  p.EXPECT_ATTRNAME,
		EXPECT_ATTRVALUE: p.EXPECT_ATTRVALUE,
		EXPECT_NODE_TYPE: p.EXPECT_NODE_TYPE,
		NEWLINE:          p.NEWLINE,
		NODEID:           p.NODEID,
		NODETYPE:         p.NODETYPE,
		NOTEVALUE:        p.NOTEVALUE,
		OPERATORVALUE:    p.OPERATORVALUE,
		START:            p.START,
	}

	for _, p.C = range t.Buffer {
		Funcs[p.State]()
		p.LenCheck()
		p.Debug()
	} // while

	return nil
}

func (t *GccNodeTest) Execute() *int {
	r := 0
	//fmt.Printf("exec %#v\n",t)
	return &r
}

func (t *GccNodeTest) Report() {

}


func (p *ParserInstance) Debug2(v string ) {}
func (p *ParserInstance) Debug() {
	if p.Parser.Globals.DebugLevel > 10 {
	fmt.Printf("Char: '%s'\t", string(p.C))
	fmt.Printf("State: '%s'\t", p.Parser.Globals.StateLookup[p.State])
		fmt.Printf("Token: '%s'\n", p.Token)
	}
}
