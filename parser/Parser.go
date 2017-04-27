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
	Consumer TreeConsumer
}

type TreeNode interface {
	SetAttr(attrname string, values []string)
	Finish()
}

type TreeConsumer interface {
	NodeType(nodetype string, nodeid string) (TreeNode)
	Report()
}

type GccNodeTest struct {
	Filename string
	Buffer   string
	Globals  *ParserGlobal
	
	//
	//Nodeids   map[string]int
	//Nodetypes map[string]map[string]int
}
type ParserInstance struct {
	Statements int
	State      Token
	Pos        int
	LastPos        int
	Line       []byte // collected line so far for debugging
	Token      []byte
	NodeId     string
	NodeType   string
	AttrName   string
	AttrValues []string
	C          rune // current rune
	Parser     *GccNodeTest
	Current    TreeNode   // interface
}

func NewParser(Consumer TreeConsumer) *ParserGlobal {

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
		ATTRVALUE_STRG:   "ATTRVALUE_STRG",
		ATTRVALUE_LNGT:   "ATTRVALUE_LNGT",
	}

	return &ParserGlobal{
		//DebugLevel:      100,
		//DebugLevel:      11,
		DebugLevel:      0,
		StateLookup:     lookup,
		Attrnames:       make(map[string]int),
		FieldMachine:    NewTrieFields(Fields[:]),
		NotesMachine:    NewTrie(Notes[:]),
		AccessMachine:   NewTrie(Access[:]),
		OperatorMachine: NewTrie(Operators[:]),
		SpecMachine:     NewTrie(Spec[:]),
		LinkMachine:     NewTrie(Link[:]),
		Consumer:Consumer,
	}
}

func (t *GccNodeTest) Init(Filename string, Globals *ParserGlobal) {
	fmt.Printf("init %s\n", Filename)
	t.Filename = Filename
	t.Globals = Globals

	//t.Nodeids = make(map[string]int)
	//t.Nodetypes = make(map[string]map[string]int)
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
	ATTRVALUE_STRG
	ATTRVALUE_LNGT

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
	//t.Nodeids[nodeid] = 1
}

func (p *ParserInstance) ProcessNodeType(nodetype string, nodeid string) {
	p.Current = p.Parser.Globals.Consumer.NodeType(nodetype,nodeid)
	// if val, ok := t.Nodetypes[nodetype]; ok {
	// 	val[nodeid] = 1
	// } else {

	// 	t.Nodetypes[nodetype] = make(map[string]int)
	// 	t.Nodetypes[nodetype][nodeid] = 1
	// }
}

// func (t *GccNodeTest) AttrValue(field string) {
// }

// func (t *GccNodeTest) AttrName(field string) {
// }



func (p *ParserInstance) SetAttrName(name string) {
	if len(name) == 0 {
		p.Panic("nul name")
	}
	if name[0] == '\n' {
		fmt.Printf("err Set attrname:%s", name)
		p.Panic("nl name")
	}
	p.AttrName = name
}

func (p *ParserInstance) ResetAttrName() {
	if len(p.AttrName) == 0 {
		p.Panic("rest null")
	}
	p.AttrName = ""
}

func (p *ParserInstance) ResetAttrValues() {
	p.AttrValues = make([]string, 0)
}

func (p *ParserInstance) AddAttrValue() {
	if len(p.Token) == 0 {
		p.Panic("attr value token null")
	}
	if p.AttrName ==string(p.Token) {
		p.Panic("err")
	}
	//fmt.Printf("adding attr: '%s' -> '%s'\n", p.AttrName, string(p.Token))
	
	p.AttrValues = append(p.AttrValues,string(p.Token))
}


func (p *ParserInstance) FinishAttribute() {
	if (p.Current != nil) {
		if len(p.AttrName) == 0 {
			p.Panic("No attr name")
		}
		p.Current.SetAttr(p.AttrName, p.AttrValues)
	} else {
		fmt.Printf("nil\n")
	}
	
	// we are done with this attribute, now clear the flags
	p.ResetAttrName()
	p.AttrValues = make([]string, 0)
}

func (p *ParserInstance) Add() {
	if  p.LastPos !=p.Pos {
		p.LastPos =p.Pos
		p.Token = append(p.Token, byte(p.C))
		p.Line = append(p.Line, byte(p.C))
	} else {
		p.Panic(string(p.Token))
	}
}

func (p *ParserInstance) AddNotWs() {
	if p.C == ' ' ||
		p.C == '\t' ||
		p.C == '\n' ||
		p.C == '\r' {
		p.Panic("ws")
	}
	p.Add()

}

// skip this in the token but gather it for the line
func (p *ParserInstance) Skip() {

	if  p.LastPos !=p.Pos {
		p.LastPos =p.Pos
		p.Line = append(p.Line, byte(p.C))
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
	if len(p.Token) > 0 {
		p.DebugBytes("Token '%s'\n", p.Token)
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
		p.Debug3("found node id %s\n", p.NodeId)
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
		if p.NodeType[0] == '1' || p.NodeType[0] == '2' || p.NodeType[0] == '6'|| p.NodeType[0] == '7' {
			panic(p.NodeType)
		}
		p.Debug3("found node type %s %s\n", p.NodeType)
		p.ProcessNodeType(p.NodeType, p.NodeId)
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
	} else	if p.C == '\t' {
		p.Skip()
		p.ConsumeToken()
	} else

	if p.C == '\n' {
		p.SetState(NEWLINE)
		p.Skip()
		return
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
	p.Debug3("Debug attrname '%s'\n",string(p.C))
	if p.C == '@' {
		p.Panic("not allowd in name")
	}
	if p.C == '\n' {
		// no attribute name, just end the statement
		p.SetState(NEWLINE)
		p.Skip()
		return
	}
	if p.C == ' ' {		
		//if ((len(p.Token) > 1)(p.Token[0] == 'o' && p.Token[1] =='p')) {
		if strings.HasPrefix(string(p.Token),"op") {
			p.Debug2("Starting op")
			if p.TokenEnd()== ':' {
				p.SetAttrName(string(p.Token))
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
					p.SetAttrName( string( p.Token))
					p.Skip()
					p.ConsumeToken()

					//p.Debug3("found p.AttrName %s %s %s\n", p.NodeId, p.NodeType, attrname)
					p.Parser.Globals.Attrnames[p.AttrName] = p.Parser.Globals.Attrnames[p.AttrName] + 1
					p.ResetAttrValues()
					if p.AttrName == "note:" {
						p.SetState(NOTEVALUE)
					} else {
						p.SetState(EXPECT_ATTRVALUE)
					}
				} else {
					p.Debug3("debug name2 %s\n", string(p.C))
					p.SetAttrName(string(p.Token))
					p.Add()
					// it is not a attr name, but a value, so switch				
					p.SetState(ATTRVALUE)
				}
			} else {
				p.Debug3("debug name4 %s\n", string(p.C))
				p.Debug()
				p.Add()
			}
		}
	} else {
		p.Debug3("debug name5 %s\n", string(p.C))
		p.Add()
	}
}

func (p *ParserInstance) EXPECT_ATTRVALUE() {

	if p.C == ' ' {
		p.Skip()
		p.ConsumeToken()
	} else {
		if p.AttrName == "strg:" {
			p.SetState(ATTRVALUE_STRG)
			p.Add()
		} else if p.AttrName == "lngt:" {
			p.SetState(ATTRVALUE_LNGT)
			p.Add()
		} else {
			p.SetState(ATTRVALUE)
			p.Add()
		}
		
	}
}


func (p *ParserInstance) NOTE_OPERATOR_VALUE() {
	p.DebugBytes("check note operator %s \n", p.Token)
	p.AddAttrValue()
	p.Skip()
	p.ConsumeToken()
	// p.State on p.State
	p.SetState(OPERATORVALUE)
}

func (p *ParserInstance) NOTE_OPERATOR() {
	p.AddAttrValue()
	p.ConsumeToken()
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

	if (p.C == '\n') {
		p.SetState(NEWLINE)
		p.AddAttrValue()
		p.FinishAttribute()
		p.Skip()
		p.ConsumeToken()
		return

	}
	
	if (p.C == ' ')  { // ws separator

		if p.Token[0] == '@' {
			// prefix
			p.AddAttrValue()
			// p.Debug3("found attrvalue @ %s %s %s %s\n", p.NodeId, p.NodeType,
			// 	p.AttrName,p.AttrValues)
			p.FinishAttribute()
			p.Skip()
			p.ConsumeToken()
			if (p.C== ' ') {
				p.SetState(AFTERATTRVALUE2)
				return
			}
			
			
		} else if p.TokenEnd()== ':' {

			//p.Debug3("debug name2 %s\n", string(p.C))
			if p.Parser.CheckFieldName(p.Token) {
				// it is an attribute name
				p.SetAttrName( string(p.Token))
				//p.Debug3("found p.AttrName %s %s %s\n", p.NodeId, p.NodeType, p.AttrName)
				p.Parser.Globals.Attrnames[p.AttrName] = p.Parser.Globals.Attrnames[p.AttrName] + 1
				p.ResetAttrValues()
				p.Skip()
				p.ConsumeToken()
				p.SetState(EXPECT_ATTRVALUE)
			} else {
				// there was a space so it is the end with a : but the field name does not match known fields so ignore it.
				// could be a : in a string, but we should be able to match that better
				p.Add()

				//p.Debug3("debug name3 %s\n", string(p.C))
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
					// now look up the note and end
					if p.Parser.CheckNote(p.Token) {
						// end of the token.
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
				p.AddAttrValue()
				p.Skip()
				p.ConsumeToken()
				//p.Debug3("debug name4 %s\n", string(p.C))
				p.Debug3("Attrname : %s\n", p.AttrName)
				p.FinishAttribute()
				p.SetState(AFTERATTRVALUE)
			}
		}
	} else {
		//p.Debug3("debug name1 %s\n", string(p.C))
		p.Add()
	}
}

func (p *ParserInstance) ATTRVALUE_STRG() {
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
	if p.TokenEnd()== ':' {
		if strings.HasSuffix(string(p.Token),"lngt:") {
			p.Token= p.Token[0:len(p.Token) - len(" lngt:")] // strip off end
			p.AddAttrValue()
			p.FinishAttribute()
			p.Skip()
			p.ConsumeToken()

			p.SetAttrName("lngt:")
			
			p.Parser.Globals.Attrnames[p.AttrName] = p.Parser.Globals.Attrnames[p.AttrName] + 1
			p.ResetAttrValues()
			p.SetState(EXPECT_ATTRVALUE)
			return
		}
	}
	
	p.Add()
}

func (p *ParserInstance) TrimStringLength(){
	//
	
}

func (p *ParserInstance) ATTRVALUE_LNGT() {
	if p.C != ' ' {
		p.Add()
	} else {
		p.TrimStringLength()
		p.AddAttrValue()
		p.FinishAttribute()
		p.Skip()
		p.ConsumeToken()
		p.SetState(AFTERATTRVALUE)
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
	p.Debug3("AFTERATTRVALUE2 '%s'\n",string(p.C))
	if p.C == '\n' {
		p.Skip()
		p.SetState(NEWLINE)
	} else if p.C == ' ' {
		p.Skip()
		//p.SetState(AFTERATTRVALUE2)
	} else {
		p.SetState(ATTRNAME)
		p.AddNotWs()
	}
}

func (p *ParserInstance) OPERATORVALUE() {
	if p.C == ' ' || p.C == '\t' || p.C == '\n' {
		p.Skip() //
	} else {
		if p.Check(p.Parser.Globals.OperatorMachine) { // will call add or skip
			p.AddAttrValue()
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
				p.AddAttrValue()
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
	
	p.Line = []byte{}
	if (p.Current != nil) {
		p.Statements++
		p.Current.Finish()
		p.Current = nil
	} else {
		p.Panic("no current\n")
	}
}

func (p *ParserInstance) Panic(l string ) {
	p.Debug3("Panic: %s\n", l)
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
		Line:   []byte{},
		Statements : 0,
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
		ATTRVALUE_STRG:   p.ATTRVALUE_STRG,
		ATTRVALUE_LNGT:   p.ATTRVALUE_LNGT,
	}

	for _, p.C = range t.Buffer {
		p.DebugIn()
		Funcs[p.State]()
		p.LenCheck()
		p.Debug()
	} // while
	// final finish statement
	p.FinishStatement()
	if (p.Statements==0){
		p.Panic("no statements found")
	}
	return nil
}

func (t *GccNodeTest) Execute() *int {
	r := 0
	//fmt.Printf("exec %#v\n",t)
	return &r
}

func (t *GccNodeTest) Report() {
	t.Globals.Consumer.Report()
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

func (p *ParserInstance) DebugIn() {
	if p.Parser.Globals.DebugLevel > 10 {
		fmt.Printf("In Char: '%s'\t", string(p.C))
		fmt.Printf("In State: '%s'\t", p.Parser.Globals.StateLookup[p.State])
		fmt.Printf("In Token: '%s'\n", p.Token)
	}
}

func (p *ParserInstance) DebugNow() {

	fmt.Printf("Char: '%s'\t", string(p.C))
	fmt.Printf("State: '%s'\t", p.Parser.Globals.StateLookup[p.State])
	fmt.Printf("Token: '%s'\n", p.Token)
	
}
