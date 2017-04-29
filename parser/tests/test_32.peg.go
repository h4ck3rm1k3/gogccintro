package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	
)

const endSymbol rune = 1114112

/* The rule types inferred from the grammar are below. */
type pegRule uint8

const (
	ruleUnknown pegRule = iota
	ruleOneAttr
	ruleStringAttr
	ruleAddrAttr
	ruleSpecValue
	ruleNodeAttr
	ruleSourceAttr
	ruleIntAttr
	ruleSignAttr
	ruleIntAttr3
	ruleTagAttr
	ruleRandomSpec
	ruleBodyAttr
	ruleAccsAttr
	ruleNoteAttr
	ruleLinkAttr
	ruleQualAttr
	ruleIntAttr2
	ruleSignedIntAttr
	ruleLngtAttr
	
)

var rul3s = [...]string {
	"Unknown",
	"OneAttr",
	"StringAttr",
	"AddrAttr",
	"SpecValue",
	"NodeAttr",
	"SourceAttr",
	"IntAttr",
	"SignAttr",
	"IntAttr3",
	"TagAttr",
	"RandomSpec",
	"BodyAttr",
	"AccsAttr",
	"NoteAttr",
	"LinkAttr",
	"QualAttr",
	"IntAttr2",
	"SignedIntAttr",
	"LngtAttr",
	
}

type token32 struct {
	pegRule
	begin, end uint32
}

func (t *token32) String() string {
	return fmt.Sprintf("\x1B[34m%v\x1B[m %v %v", rul3s[t.pegRule], t.begin, t.end)
}


type node32 struct {
	token32
	up, next *node32
}

func (node *node32) print(pretty bool, buffer string) {
	var print func(node *node32, depth int)
	print = func(node *node32, depth int) {
		for node != nil {
			for c := 0; c < depth; c++ {
				fmt.Printf(" ")
			}
			rule := rul3s[node.pegRule]
			quote := strconv.Quote(string(([]rune(buffer)[node.begin:node.end])))
			if !pretty {
				fmt.Printf("%v %v\n", rule, quote)
			} else {
				fmt.Printf("\x1B[34m%v\x1B[m %v\n", rule, quote)
			}
			if node.up != nil {
				print(node.up, depth + 1)
			}
			node = node.next
		}
	}
	print(node, 0)
}

func (node *node32) Print(buffer string) {
	node.print(false, buffer)
}

func (node *node32) PrettyPrint(buffer string) {
	node.print(true, buffer)
}

type tokens32 struct {
	tree		[]token32
}

func (t *tokens32) Trim(length uint32) {
	t.tree = t.tree[:length]
}

func (t *tokens32) Print() {
	for _, token := range t.tree {
		fmt.Println(token.String())
	}
}

func (t *tokens32) AST() *node32 {
	type element struct {
		node *node32
		down *element
	}
	tokens := t.Tokens()
	var stack *element
	for _, token := range tokens {
		if token.begin == token.end {
			continue
		}
		node := &node32{token32: token}
		for stack != nil && stack.node.begin >= token.begin && stack.node.end <= token.end {
			stack.node.next = node.up
			node.up = stack.node
			stack = stack.down
		}
		stack = &element{node: node, down: stack}
	}
	if stack != nil {
		return stack.node
	}
	return nil
}

func (t *tokens32) PrintSyntaxTree(buffer string) {
	t.AST().Print(buffer)
}

func (t *tokens32) PrettyPrintSyntaxTree(buffer string) {
	t.AST().PrettyPrint(buffer)
}

func (t *tokens32) Add(rule pegRule, begin, end, index uint32) {
	if tree := t.tree; int(index) >= len(tree) {
		expanded := make([]token32, 2 * len(tree))
		copy(expanded, tree)
		t.tree = expanded
	}
	t.tree[index] = token32{
		pegRule: rule,
		begin: begin,
		end: end,
	}
}

func (t *tokens32) Tokens() []token32 {
	return t.tree
}


type GccNode struct {
	
	Buffer		string
	buffer		[]rune
	rules		[20]func() bool
	parse		func(rule ...int) error
	reset		func()
	Pretty 	bool
tokens32
}

func (p *GccNode) Parse(rule ...int) error {
	return p.parse(rule...)
}

func (p *GccNode) Reset() {
	p.reset()
}

type textPosition struct {
	line, symbol int
}

type textPositionMap map[int] textPosition

func translatePositions(buffer []rune, positions []int) textPositionMap {
	length, translations, j, line, symbol := len(positions), make(textPositionMap, len(positions)), 0, 1, 0
	sort.Ints(positions)

	search: for i, c := range buffer {
		if c == '\n' {line, symbol = line + 1, 0} else {symbol++}
		if i == positions[j] {
			translations[positions[j]] = textPosition{line, symbol}
			for j++; j < length; j++ {if i != positions[j] {continue search}}
			break search
		}
 	}

	return translations
}

type parseError struct {
	p *GccNode
	max token32
}

func (e *parseError) Error2() string {
	tokens, error := []token32{e.max}, "\n"
	positions, p := make([]int, 2 * len(tokens)), 0
	for _, token := range tokens {
		positions[p], p = int(token.begin), p + 1
		positions[p], p = int(token.end), p + 1
	}
	translations := translatePositions(e.p.buffer, positions)
	format := "parse error near %v (line %v symbol %v - line %v symbol %v):\n%v\n"
	if e.p.Pretty {
		format = "parse error near \x1B[34m%v\x1B[m (line %v symbol %v - line %v symbol %v):\n%v\n"
	}
	for _, token := range tokens {
		begin, end := int(token.begin), int(token.end)
		error += fmt.Sprintf(format,
                         rul3s[token.pegRule],
                         translations[begin].line, translations[begin].symbol,
                         translations[end].line, translations[end].symbol,
                         strconv.Quote(string(e.p.buffer[begin:end])))
	}

	return error
}


func (p *GccNode) PrintSyntaxTree() {
	if p.Pretty {
		p.tokens32.PrettyPrintSyntaxTree(p.Buffer)
	} else {
		p.tokens32.PrintSyntaxTree(p.Buffer)
	}
}




func (p *GccNode) Init() {
	var (
		max token32
		position, tokenIndex uint32
		buffer []rune
)
	p.reset = func() {
		max = token32{}
		position, tokenIndex = 0, 0

		p.buffer = []rune(p.Buffer)
		if len(p.buffer) == 0 || p.buffer[len(p.buffer) - 1] != endSymbol {
			p.buffer = append(p.buffer, endSymbol)
		}
		buffer = p.buffer
	}
	p.reset()

	_rules := p.rules
tree := tokens32{tree: make([]token32, math.MaxInt16)}
p.parse = func(rule ...int) error {
		r := 1
		if len(rule) > 0 {
			r = rule[0]
		}
		matches := p.rules[r]()
p.tokens32 = tree
if matches {
p.Trim(tokenIndex)
return nil
		}
		return &parseError{p, max}
	}

	add := func(rule pegRule, begin uint32) {
tree.Add(rule, begin, position, tokenIndex)
tokenIndex++
		if begin != position && position > max.end {
			max = token32{rule, begin, position}
		}
	}

	

	

	

	

	_rules = [...]func() bool {
		nil,
  /* 0 OneAttr <- <((&() LngtAttr) | (&() SignedIntAttr) | (&() IntAttr2) | (&() QualAttr) | (&() LinkAttr) | (&() NoteAttr) | (&() AccsAttr) | (&() BodyAttr) | (&() RandomSpec) | (&() TagAttr) | (&() IntAttr3) | (&() SignAttr) | (&() IntAttr) | (&() SourceAttr) | (&() NodeAttr) | (&() SpecValue) | (&() AddrAttr) | (&() StringAttr))> */
  func() bool {
   {
position1 := position
   {
   switch buffer[position] {
   case '<nil>':
   {
position3 := position
add(ruleLngtAttr, position3)
   }
break
   case '<nil>':
   {
position4 := position
add(ruleSignedIntAttr, position4)
   }
break
   case '<nil>':
   {
position5 := position
add(ruleIntAttr2, position5)
   }
break
   case '<nil>':
   {
position6 := position
add(ruleQualAttr, position6)
   }
break
   case '<nil>':
   {
position7 := position
add(ruleLinkAttr, position7)
   }
break
   case '<nil>':
   {
position8 := position
add(ruleNoteAttr, position8)
   }
break
   case '<nil>':
   {
position9 := position
add(ruleAccsAttr, position9)
   }
break
   case '<nil>':
   {
position10 := position
add(ruleBodyAttr, position10)
   }
break
   case '<nil>':
   {
position11 := position
add(ruleRandomSpec, position11)
   }
break
   case '<nil>':
   {
position12 := position
add(ruleTagAttr, position12)
   }
break
   case '<nil>':
   {
position13 := position
add(ruleIntAttr3, position13)
   }
break
   case '<nil>':
   {
position14 := position
add(ruleSignAttr, position14)
   }
break
   case '<nil>':
   {
position15 := position
add(ruleIntAttr, position15)
   }
break
   case '<nil>':
   {
position16 := position
add(ruleSourceAttr, position16)
   }
break
   case '<nil>':
   {
position17 := position
add(ruleNodeAttr, position17)
   }
break
   case '<nil>':
   {
position18 := position
add(ruleSpecValue, position18)
   }
break
   case '<nil>':
   {
position19 := position
add(ruleAddrAttr, position19)
   }
break
   default:
   {
position20 := position
add(ruleStringAttr, position20)
   }
break
   }
   }

add(ruleOneAttr, position1)
   }
   return true
  },
  nil,
  nil,
  nil,
  nil,
  nil,
  nil,
  nil,
  nil,
  nil,
  nil,
  nil,
  nil,
  nil,
  nil,
  nil,
  nil,
  nil,
  nil,
 }
 p.rules = _rules
}