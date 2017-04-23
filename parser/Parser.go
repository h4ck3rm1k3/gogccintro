package main

import (
	//"github.com/anknown/ahocorasick"
	//"bytes"
	"fmt"
	"strings"
	//"unicode/utf8"
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
	DebugLevel      int
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
	LastPos        int
	Line       string // collected line so far for debugging
	//Token      string
	Token      []byte
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
		//DebugLevel:      100,
		DebugLevel:      0,
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

func (t *GccNodeTest) CheckName(machine *Machine, value []byte ) bool {
	//machine.PrintOutput()
	return machine.ExactSearch(value)
}

func (t *GccNodeTest) CheckNamePartial(machine *Machine, value []byte) bool {

	//machine.PrintOutput()
	
	found := machine.MultiPatternSearch(value) 
	if found{
		//fmt.Printf("Partial match %s \n",value)
		return true
	} else {
		//fmt.Printf("Partial fail %s \n",value)
	}
	return false
}

func (t *GccNodeTest) CheckFieldName(field []byte) bool {
	return t.CheckName(t.Globals.FieldMachine, field)
}

func (t *GccNodeTest) CheckNote(field []byte) bool {
	return t.CheckName(t.Globals.NotesMachine, field)
}

func (t *GccNodeTest) CheckNotePartial(field []byte) bool {
	
	return t.CheckNamePartial(t.Globals.NotesMachine, field)
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

func (p *ParserInstance) FinishAttribute() {
	// we are done with this attribute, now clear the flags
	p.AttrName = ""
	p.AttrValues = make([]string, 0)
}
func (p *ParserInstance) Add() {
	// fmt.Printf("adding %s %s\n",p.Token, string(p.C))
	// if len(p.Token) > 0 {
	// 	if (string(p.TokenEnd()) == string(p.C)) {
	// 		fmt.Printf("check %s %s",p.Token, string(p.C))
	// 		//panic("dup")
	// 	} else {
	// 		//fmt.Printf("adding %s %s",p.Token, string(p.C))
	// 	}
	// }
	if  p.LastPos !=p.Pos {
		p.LastPos =p.Pos
		p.Token = append(p.Token, byte(p.C))
		//p.Line = p.Line + string(p.C)
	} else {
		p.Panic(string(p.Token))
	}
}

// skip this in the token but gather it for the line
func (p *ParserInstance) Skip() {

	if  p.LastPos !=p.Pos {
		p.LastPos =p.Pos
		//p.Line = p.Line + string(p.C)
	} else {
		p.Panic(string(p.Token))
	}

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
			p.Skip()
			return false
		}
	} else {
		fmt.Printf("error too large '%s'\n", p.Token)
		p.Panic("")
		return false
	}

}

func (p *ParserInstance) ConsumeToken() {
	if len(p.Token) > 1 {
		p.Token = []byte{}
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
	} else {
		p.Skip()
	}
}

func (p *ParserInstance) NODEID() {
	//else if p.State == NODEID {
	if isDigit(p.C) {
		p.Add()
	} else if p.C == ' ' {
		p.NodeId = string(p.Token)
		p.Parser.NodeId(p.NodeId)
		//fmt.Printf("found node id %s\n", p.NodeId)
		p.Skip()
		p.ConsumeToken()
		p.SetState(EXPECT_NODE_TYPE)
	} else {
		p.Skip()
	}
}

func (p *ParserInstance) EXPECT_NODE_TYPE() {
	if p.C == ' ' {
		p.Skip()
		p.ConsumeToken()
	} else {
		p.Add()
		p.SetState(NODETYPE)
	}
}

func (p *ParserInstance) NODETYPE() {
	if p.C == ' ' {
		p.NodeType = string(p.Token)
		//fmt.Printf("found node type %s %s\n", p.NodeId, p.NodeType)
		p.Parser.NodeType(p.NodeType, p.NodeId)
		p.Skip()
		p.ConsumeToken()
		p.SetState(EXPECT_ATTRNAME)
	} else {
		p.Add()
	}
}

func (p *ParserInstance) EXPECT_ATTRNAME() {
	if p.C == ' ' {
		p.Skip()
		p.ConsumeToken()
	} else {
		p.Add()
		p.SetState(ATTRNAME)
	}
}
func (p *ParserInstance) TokenEnd() byte {
	t:=p.Token[len(p.Token)-1]
	p.Debug3("Token end '%s'",string(t))
	return t
}

func (p *ParserInstance) ATTRNAME() {
	if p.C == ' ' {		
		//if ((len(p.Token) > 1)(p.Token[0] == 'o' && p.Token[1] =='p')) {
		if strings.HasPrefix(string(p.Token),"op") {
			p.Debug2("Starting op")
			if p.TokenEnd()== ':' {
				p.AttrName = string( p.Token)
				p.Skip()
				p.ConsumeToken()
				p.SetState(EXPECT_ATTRVALUE)
			} else{
				p.Add()
			}
		} else {
			//size := utf8.RuneCountInString(p.Token)
			if p.TokenEnd()== ':' {
				if p.Parser.CheckFieldName(p.Token) {
					p.AttrName = string( p.Token)
					p.Skip()
					p.ConsumeToken()

					//fmt.Printf("found p.AttrName %s %s %s\n", p.NodeId, p.NodeType, attrname)
					p.Parser.Globals.Attrnames[p.AttrName] = p.Parser.Globals.Attrnames[p.AttrName] + 1
					p.AttrValues = make([]string, 0)
					if p.AttrName == "note:" {
						p.SetState(NOTEVALUE)
					} else {
						p.SetState(EXPECT_ATTRVALUE)
					}
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
		p.Skip()
		p.ConsumeToken()
	} else {
		p.SetState(ATTRVALUE)
		p.Add()
	}
}

func (p *ParserInstance) NOTE_OPERATOR_VALUE() {
	p.DebugBytes("check note operator %s \n", p.Token)
	p.AttrValues = append(p.AttrValues, string(p.Token))
	p.Skip()
	p.ConsumeToken()
	// p.State on p.State
	p.SetState(OPERATORVALUE)
}

func (p *ParserInstance) NOTE_OPERATOR() {
	p.AttrValues = append(p.AttrValues, string(p.Token))
	p.ConsumeToken()
	//p.FinishAttribute()
	// the operator has started
	//p.Skip()
	p.SetState(OPERATORVALUE)
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
	p.DebugInt("check len %d\n", l)
	p.Debug3("check last %s\n", v)
	p.DebugAttrs("check values %s\n")

	if (p.C == ' ') || (p.C == '\n') { // ws separator

		if p.Token[0] == '@' {
			// prefix
			p.AttrValues = append(p.AttrValues, string(p.Token))
			//fmt.Printf("found attrvalue @ %s %s %s %s\n", p.NodeId, p.NodeType,p.AttrName,						p.AttrValues)
			p.Skip()
			p.ConsumeToken()
			p.SetState(AFTERATTRVALUE2)
		} else if p.TokenEnd()== ':' {
			if p.Parser.CheckFieldName(p.Token) {
				// it is an attribute name
				p.AttrName = string(p.Token)
				//fmt.Printf("found p.AttrName %s %s %s\n", p.NodeId, p.NodeType, p.AttrName)
				p.Parser.Globals.Attrnames[p.AttrName] = p.Parser.Globals.Attrnames[p.AttrName] + 1
				p.AttrValues = make([]string, 0)
				p.Skip()
				p.ConsumeToken()
				p.SetState(EXPECT_ATTRVALUE)
			} else {
				// there was a space so it is the end with a : but the field name does not match known fields so ignore it.
				// could be a : in a string, but we should be able to match that better
				p.Add()
				p.DebugBytes("Check : in token:%s\n", p.Token)
				//panic(p.Token)
			}
		} else {
			if p.AttrName == "note:" {
				//p.Debug()
				p.DebugInt("in note operator %d\n", l)
				p.DebugBytes("in note operator token:%s\n", p.Token)
				if l > 0 {
					if v == "operator" {
						p.NOTE_OPERATOR()
					} else {
						p.DebugBytes("check note operator2 %s\n", p.Token)
						p.Add()
						p.SetState(NOTEVALUE)
					}

				} else { // first work

					//
					//

					// now look up the note and end
					if p.Parser.CheckNote(p.Token) {
						// end of the token.
						//						
						if string(p.Token) == "operator" {
							p.NOTE_OPERATOR()
							
						} else {
							//p.Panic(fmt.Sprintf(fmt.Sprintf("matched %s and ended note %s %s %s %s\n", p.Token, p.NodeId, p.NodeType, p.AttrName, p.AttrValues)))
							p.Add()
							p.SetState(AFTERATTRVALUE)
						}
													
					} else {
						p.Add()
						//p.Debug3("unmatched first note nodeid:%s nodetype:%s attrname:%s attrvals:%s : token '%s'o\n", p.NodeId, p.NodeType, p.AttrName, p.AttrValues, p.Token))
						//p.Skip()
						//p.ConsumeToken()
						p.SetState(NOTEVALUE)					
					}
					
					
										
				}

			} else { // not note:
				//attrvalue = p.Token
				p.AttrValues = append(p.AttrValues, string(p.Token))
				//fmt.Printf("found attrvalue %s %s %s %s\n", p.NodeId, p.NodeType,p.AttrName,							p.AttrValues)
				p.Skip()
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
		p.Skip()
	} else if p.C == ' ' {
		p.SetState(AFTERATTRVALUE)
		p.Skip()
	} else {

		p.Debug3("Attrname: '%s'\t", p.AttrName)
		p.DebugAttrs("AttrValues: '%v'\n")
		// the next could be a name of an attr or a value

		if p.AttrName == "note:" {
			p.SetState(ATTRVALUE)
			p.Add()
		} else {
			p.Add()
			p.SetState(ATTRNAME)
		}
	}
}

func (p *ParserInstance) AFTERATTRVALUE2() {

	if p.C == '\n' {
		p.Skip()
		p.SetState(NEWLINE)
	} else if p.C == ' ' {
		p.Skip()
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
	//p.Skip()
	if p.C == ' ' || p.C == '\t' || p.C == '\n' {
		p.Skip() //
	} else {
		//p.Add()
		if p.Check(p.Parser.Globals.OperatorMachine) { // will call add or skip
			p.AttrValues = append(p.AttrValues, string(p.Token))
			p.FinishAttribute()
			p.ConsumeToken()
			p.NextState(true, AFTERATTRVALUE)
		}
	}
}

func (p *ParserInstance) NOTEVALUE() {

	//one of
	//Notes
	// max length 16

	//p.Parser.CheckName(p.Parser.Globals.OperatorMachine, p.Token)
	//p.Add()
	if ((len(p.Token) == 0) && (p.C == ' ')) {
		p.Skip()
	} else {
		p.Add()
		if p.Parser.CheckNote(p.Token) {
			// end of the token.	
			//p.Debug3(fmt.Sprintf("matched and ended note %s %s %s %s\n", p.NodeId, p.NodeType, p.AttrName, p.AttrValues))
			if string(p.Token) == "operator" {
				p.NOTE_OPERATOR()
			} else {
				p.AttrValues = append(p.AttrValues, string(p.Token))
				p.ConsumeToken()
				p.FinishAttribute()
				p.SetState(AFTERATTRVALUE)
			}
			//if len(p.Token) <= p.Parser.Globals.NoteMachine.MaxLen {
		} else if p.Parser.CheckNotePartial(p.Token) {
			
		} else {
			p.Skip()
			fmt.Printf("error note '%s' and c :'%c'\n", p.Token, p.C)
			p.Panic("")	
		}
	}
}

func (p *ParserInstance) FinishStatement() {
	p.Line = ""
}

func (p *ParserInstance) Panic(l string ) {
	fmt.Printf("Panic: %s\n", l)
	panic(p.Line)
}

func (p *ParserInstance) NEWLINE() {

	p.Skip()
	if p.C == '@' {
		// clear the prev lin
		p.FinishStatement()
		p.SetState(NODEID)
	} else {
		p.SetState(EXPECT_ATTRNAME)
	}
}

type ParseFunc func()

func (p *ParserInstance) LenCheck() {
	if len(p.Token) > 1250 {
		fmt.Printf("Too Long P.Token %d: %s\n", p.State, p.Token)
		p.Panic("")
	}
	if p.Pos > 1 {
		if p.Pos%10000000 == 0 {
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
		Line:   "",
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

func (p *ParserInstance) Debug2(v string) {
	if p.Parser.Globals.DebugLevel > 11 {
		fmt.Printf("Debug: '%s'\n", v)
	}
}

func (p *ParserInstance) Debug3(fmt string, val string) {
}

func (p *ParserInstance) DebugBytes(fmt string,  s []byte ) {
}

func (p *ParserInstance) DebugInt(fmt string,  s int ) {
}

func (p *ParserInstance) DebugAttrs(fmt string) {
}

func (p *ParserInstance) Debug() {
	if p.Parser.Globals.DebugLevel > 10 {
		fmt.Printf("Char: '%s'\t", string(p.C))
		fmt.Printf("State: '%s'\t", p.Parser.Globals.StateLookup[p.State])
		fmt.Printf("Token: '%s'\n", p.Token)
	}
}
