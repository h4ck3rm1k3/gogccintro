package main

import (
	//"github.com/anknown/ahocorasick"
	//"bytes"
	"fmt"
	"sync"
	"strings"
	//"unicode/utf8"
)

type AttrNames map[string]int

type StateNames map[Token]string
	
// global
var StateLookup    StateNames

func CreateStateLookup() (StateNames ){
	return map[Token]string{
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
	OPERATORVALUE_MATCH:    "OPERATORVALUE_MATCH",
	OPERATORVALUE_NL:    "OPERATORVALUE_NL",
	ATTRVALUE_STRG:   "ATTRVALUE_STRG",
	ATTRVALUE_LNGT:   "ATTRVALUE_LNGT",
	}
}

type ParserGlobal struct {
	Lock        sync.RWMutex
	Attrnames       AttrNames
	FieldMachine    *Machine
	FlagsMachine     *Machine
	NotesMachine    *Machine
	AccessMachine   *Machine
	OperatorMachine *Machine
	SpecMachine     *Machine
	LinkMachine     *Machine
	DebugLevel      int
	Consumer TreeConsumer
}

func (p *ParserGlobal) AddAttr(attrname string ) {
	p.Lock.Lock()
	defer p.Lock.Unlock()
	p.Attrnames[attrname] = p.Attrnames[attrname] + 1
}

type TreeNode interface {
	SetAttr(attrname string, values []string)
	//Finish(TreeConsumer)
}

type TreeConsumer interface {
	NodeType(nodetype string, nodeid string, filename string) (TreeNode)
	Report()
	StateTransition(from int, to int) // state
	StateUsage(from int) // state usage
	Node(TreeNode)
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

	
	return &ParserGlobal{
		Lock :sync.RWMutex{},
		//DebugLevel:      100,
		//DebugLevel:      11,
		DebugLevel:      0,
		//StateLookup:     lookup,
		Attrnames:       make(map[string]int),
		FieldMachine:    NewTrieFields(Fields[:]),
		FlagsMachine:    NewTrie(Flags[:]),
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
	OPERATORVALUE_MATCH
	OPERATORVALUE_NL
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

func (t *GccNodeTest) CheckFlagName(field []byte) bool {
	return t.CheckName(t.Globals.FlagsMachine, field)
}

func (t *GccNodeTest) AddFlag(flag string) {
	//return t.CheckName(t.Globals.FlagsMachine, field)
	if t.Globals.DebugLevel > 1000 {
		fmt.Printf("added flag: %s\n",flag)
	}
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
	p.Current = p.Parser.Globals.Consumer.NodeType(nodetype,nodeid,p.Parser.Filename)
	// if val, ok := t.Nodetypes[nodetype]; ok {
	// 	val[nodeid] = 1
	// } else {

	// 	t.Nodetypes[nodetype] = make(map[string]int)
	// 	t.Nodetypes[nodetype][nodeid] = 1
	// }
}

func (p *ParserInstance) SetAttrName(name string) {
	
	if len(name) == 0 {
		p.Panic("nul name")
	}
	sc := 0 
	for _,v := range(name) {

		if v == '\n' || v == '@' {
			fmt.Printf("err Set attrname:%s", name)
			p.Panic("nl name")
		}

		if v == ' ' {
			sc = sc +1
			// 0   : has 3 spaces
			if sc > 4 {
				fmt.Printf("err Set attrname2:%s", name)
				p.Panic("2 spaces name")
			}
		}
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
	if  p.LastPos ==  p.Pos  - 1 {  // only one char
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

	p.Parser.Globals.Consumer.StateTransition(int(p.State),int(nextState))

	p.State = nextState
}

func (p *ParserInstance) Check(m *Machine) bool {
	if len(p.Token) <= m.MaxLen {
		if p.Parser.CheckName(m, p.Token) {

			return true
		} else {

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
		p.ConsumeToken()
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
	p.Debug3("Debug attrname token:'%s'\n",string(p.Token))
	p.Debug3("Debug attrname nextchar:'%s'\n",string(p.C))
	if p.C == '@' {
		p.Panic("not allowd in name")
	}
	if p.C == '\n' {
		// no attribute name, just end the statement
		p.SetState(NEWLINE)
		p.Skip()
		return
	}
	token := string(p.Token)
	if p.C == ' ' {		
		//if ((len(p.Token) > 1)(p.Token[0] == 'o' && p.Token[1] =='p')) {
		if p.Parser.CheckFlagName(p.Token) {
			p.Parser.AddFlag(token)
			p.Skip()
			p.ConsumeToken()
			p.SetState(EXPECT_ATTRVALUE)
		} else if strings.HasPrefix(token,"op") {
			p.Debug2("Starting op")
			if p.TokenEnd()== ':' {
				s := string(p.Token)
				s = strings.Replace(s," ","_",-1)
				p.SetAttrName(s)
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
					p.Parser.Globals.AddAttr(p.AttrName)
					//p.Parser.Globals.Attrnames[p.AttrName] = p.Parser.Globals.Attrnames[p.AttrName] + 1
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
	p.DebugAttrs("check values:")

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
				p.Parser.Globals.AddAttr(p.AttrName)
				//p.Parser.Globals.Attrnames[p.AttrName] = p.Parser.Globals.Attrnames[p.AttrName] + 1
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
				if p.AttrName != "" {
					p.Debug3("Attrname : %s\n", p.AttrName)
					p.FinishAttribute()
				}
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
			end := len(p.Token) - len(" lngt:")
			if end > 0 && end < len(p.Token)  {
				p.Token= p.Token[0:end] // strip off end
			}
			p.AddAttrValue()
			p.FinishAttribute()
			p.Skip()
			p.ConsumeToken()

			p.SetAttrName("lngt:")

			p.Parser.Globals.AddAttr(p.AttrName)
			//p.Parser.Globals.Attrnames[p.AttrName] = p.Parser.Globals.Attrnames[p.AttrName] + 1
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
	p.Debug3("enter AFTERATTRVALUE: '%s'\t", string(p.C))
	if p.C == '\n' {
		p.SetState(NEWLINE)
		p.Skip()
		p.ConsumeToken() // reset the token to get rid of ws
	} else if p.C == ' ' || p.C == '\t' {
		p.Skip()
		p.ConsumeToken() // reset the token to get rid of ws
	} else {

		//p.Debug3("Attrname: '%s'\t", p.AttrName)
		p.DebugAttrs("AttrValues:\n")
		// the next could be a name of an attr or a value

		if p.AttrName == "note:" {
			p.SetState(ATTRVALUE)
			p.Add()
		} else {			
			p.SetState(ATTRNAME)
			p.Add()
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

func (p *ParserInstance) OPERATORVALUE_NL() {
	// after operator, if we get a newline, this is a special case.
	p.Debug2(fmt.Sprintf("OPERATORVALUE_NL: Token('%s') Char('%s')\n", p.Token, string(p.C)))
	if p.C == '@' {
		p.FinishStatement()
		p.Skip()
		p.SetState(NODEID)
	}
}

func (p *ParserInstance) OPERATORVALUE_MATCH (){
	// could have matched not, but we get note:
	
	// after we match an operator, look for the space because the names are ambigious
	if (p.C == ' ' ) {
		if p.Check(p.Parser.Globals.OperatorMachine) { 
			p.AddAttrValue()
			p.FinishAttribute()
			p.Debug2(fmt.Sprintf("found operator value: '%s' '%s'\n", p.Token, string(p.C)))
			p.ConsumeToken()
			p.Add()
			p.NextState(true, AFTERATTRVALUE)
		} else {
			t:=string(p.Token)
			if (t == "note:") {
				// this is a special case, with note:operator <note:> method
				// where the note: could have matched not
				p.SetAttrName(t)
				p.ConsumeToken()
				p.SetState(EXPECT_ATTRVALUE)
				p.Add()
			} else {
				p.Debug2(fmt.Sprintf("other operator value: '%s' '%s'\n", p.Token, string(p.C)))
				p.Add()
			}
		}
	} else {
		p.Add()
	}
	//else if p.Check(p.Parser.Globals.OperatorMachine) { 
	//	}
}

func (p *ParserInstance) OPERATORVALUE() {
	// this could be an operator or just the next token or end of line
	if len(p.Token) > 20 {
		p.Panic(fmt.Sprintf("Too big %s", p.Token))
	}
	if (p.C == '\n' ) {
		if p.Check(p.Parser.Globals.OperatorMachine) { 
			p.AddAttrValue()
			p.FinishAttribute()
			p.Debug2(fmt.Sprintf("found operator value: Token('%s') Char('%s')\n", p.Token, string(p.C)))
			p.ConsumeToken()
			p.Add()
			p.NextState(true, NEWLINE)
		} else {
			p.NextState(true, OPERATORVALUE_NL)
			p.Skip() 
		}
	} else if p.C == ' ' || p.C == '\t'  {
		if p.Check(p.Parser.Globals.OperatorMachine) { 
			p.AddAttrValue()
			p.FinishAttribute()
			p.Debug2(fmt.Sprintf("found operator value: Token('%s') Char('%s')\n", p.Token, string(p.C)))
			p.ConsumeToken()
			p.Add()
			p.NextState(true, NEWLINE)
		} else {
			p.Debug2(fmt.Sprintf("check operator value: Token('%s') Char('%s')\n", p.Token, string(p.C)))
			p.Skip() 
		}

	} else {
		
		if p.Check(p.Parser.Globals.OperatorMachine) { 
			p.Add()
			p.NextState(true, OPERATORVALUE_MATCH)
		} else {
			if p.C == '@' {
				p.NextState(true, AFTERATTRVALUE)
				p.Add()
				p.Debug2(fmt.Sprintf("empty operator, next value: '%s' '%s'\n", p.Token, string(p.C)))
			} else {
				// could be another 
				p.Debug2(fmt.Sprintf("skip operator value: '%s' '%s'\n", p.Token, string(p.C)))
				p.Add()
			}
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
//		fmt.Printf("Node: %#v\n", p.Current)
		p.Parser.Globals.Consumer.Node(p.Current)
		p.Current = nil
	} else {
		p.Panic("no current\n")
	}
}

func (p *ParserInstance) Panic(l string ) {
	p.Debug3("Panic: %s\n", l)
	
	p.Parser.Globals.DebugLevel=1000
	p.DebugAttrs("Panic")
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
	if len(p.Token) > 10250 {
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
		OPERATORVALUE_NL:    p.OPERATORVALUE_NL,
		OPERATORVALUE_MATCH:    p.OPERATORVALUE_MATCH,
		START:            p.START,
		ATTRVALUE_STRG:   p.ATTRVALUE_STRG,
		ATTRVALUE_LNGT:   p.ATTRVALUE_LNGT,
	}

	for _, p.C = range t.Buffer {
		p.DebugIn()
		Funcs[p.State]()
		p.LenCheck()
		p.Debug()

		// usage of a state
		p.Parser.Globals.Consumer.StateUsage(int(p.State))
		
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

func (p *ParserInstance) Debug3(fmts string, val string) {
	if p.Parser.Globals.DebugLevel > 12 {
		fmt.Printf(fmts, val)
	}
}

func (p *ParserInstance) DebugBytes(fmts string,  s []byte ) {
	if p.Parser.Globals.DebugLevel > 13 {
		fmt.Printf(fmts, string(s))
	}
}

func (p *ParserInstance) DebugInt(fmts string,  s int ) {
	if p.Parser.Globals.DebugLevel > 14 {
		fmt.Printf(fmts, s)
	}
}

func (p *ParserInstance) DebugAttrs(fmts string) {
	if p.Parser.Globals.DebugLevel > 200 {
		fmt.Printf("State: '%s'\t", StateLookup[p.State])
		fmt.Printf("Attrs: '%s'\t", fmts)
		fmt.Printf("Statements: '%d'\t", p.Statements)
		fmt.Printf("Pos: '%d'\t", p.Pos)
		fmt.Printf("LastPos: '%d'\t", p.LastPos)
		fmt.Printf("Token: '%s'\t", p.Token)
		fmt.Printf("Line: '%s'\t", p.Line)
		fmt.Printf("NodeId: '%s'\t", p.NodeId)
		fmt.Printf("NodeType: '%s'\t", p.NodeType)
		fmt.Printf("AttrName: '%s'\t", p.AttrName)
		fmt.Printf("AttrValues: '%s'\t", p.AttrValues)
		//fmt.Printf("C: '%s'\n", string(p.C))
		if p.C == '\n' {
			fmt.Printf("In Char: '<newline>'\t")
		}else if p.C == '\t' {
			fmt.Printf("In Char: '<tab>'\t")
		}else if p.C == ' ' {
			fmt.Printf("In Char: '<space>'\t")
		} else {
			fmt.Printf("In Char: '%s'\t", string(p.C))
		}
	}
	
	
}

func (p *ParserInstance) Debug() {
	if p.Parser.Globals.DebugLevel > 10 {
		//fmt.Printf("Char: '%s'\t", string(p.C))
		if p.C == '\n' {
			fmt.Printf("In Char: '<newline>'\t")
		}else if p.C == '\t' {
			fmt.Printf("In Char: '<tab>'\t")
		}else if p.C == ' ' {
			fmt.Printf("In Char: '<space>'\t")
		} else {
			fmt.Printf("In Char: '%s'\t", string(p.C))
		}
		fmt.Printf("State: '%s'\t", StateLookup[p.State])
		fmt.Printf("Token: '%s'\n", p.Token)
	}
}

func (p *ParserInstance) DebugIn() {
	if p.Parser.Globals.DebugLevel > 10 {
		if p.C == '\n' {
			fmt.Printf("In Char: '<newline>'\t")
		}else if p.C == '\t' {
			fmt.Printf("In Char: '<tab>'\t")
		}else if p.C == ' ' {
			fmt.Printf("In Char: '<space>'\t")
		} else {
			fmt.Printf("In Char: '%s'\t", string(p.C))
		}
		fmt.Printf("In State: '%s'\t", StateLookup[p.State])
		t := string(p.Token)
		t = strings.Replace(t,"\n","<newline>",-1)
		t = strings.Replace(t," ","<space>",-1)
		t = strings.Replace(t,"\t","<tab>",-1)
		fmt.Printf("In Token: '%s'\n", p.Token)
	}
}

func (p *ParserInstance) DebugNow() {

	//fmt.Printf("Char: '%s'\t", string(p.C))
	if p.C == '\n' {
		fmt.Printf("Char: '<newline>'\t")
	}else if p.C == '\t' {
		fmt.Printf("Char: '<tab>'\t")
	}else if p.C == ' ' {
		fmt.Printf("Char: '<space>'\t")
	} else {
		fmt.Printf("Char: '%s'\t", string(p.C))
	}
	
	fmt.Printf("State: '%s'\t", StateLookup[p.State])
	fmt.Printf("Token: '%s'\n", p.Token)
	
}
