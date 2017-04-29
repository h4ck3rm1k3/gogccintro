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
	ruleNodeFieldName
	rulePegText
	ruleInteger
	ruleAction0
	
)

var rul3s = [...]string {
	"Unknown",
	"NodeFieldName",
	"PegText",
	"Integer",
	"Action0",
	
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
	rules		[5]func() bool
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


func (p *GccNode) Execute() {
	buffer, _buffer, text, begin, end := p.Buffer, p.buffer, "", 0, 0
	for _, token := range p.Tokens() {
		switch (token.pegRule) {
		
		case rulePegText:
			begin, end = int(token.begin), int(token.end)
			text = string(_buffer[begin:end])
		
		case ruleAction0:
			

	FieldName=buffer[begin:end]
   
		
		}
	}
	_, _, _, _, _ = buffer, _buffer, text, begin, end
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

	

	
	/*matchChar := func(c byte) bool {
		if buffer[position] == c {
			position++
			return true
		}
		return false
	}*/
	

	

	

	_rules = [...]func() bool {
		nil,
  /* 0 NodeFieldName <- <(<((('p' / 'P') ('r' / 'R') ('m' / 'M') ('s' / 'S')) / (('b' / 'B') ('a' / 'A') ('s' / 'S') ('e' / 'E')) / (('d' / 'D') ('e' / 'E') ('c' / 'C') ('l' / 'L')) / (('f' / 'F') ('n' / 'N') ('c' / 'C') ('s' / 'S')) / (('c' / 'C') ('n' / 'N') ('s' / 'S') ('t' / 'T')) / (('v' / 'V') ('a' / 'A') ('r' / 'R') ('s' / 'S')) / (('v' / 'V') ('f' / 'F') ('l' / 'L') ('d' / 'D')) / (('c' / 'C') ('h' / 'H') ('a' / 'A') ('n' / 'N')) / (('c' / 'C') ('l' / 'L') ('s' / 'S')) / (('m' / 'M') ('i' / 'I') ('n' / 'N')) / ((&() ((&('V' | 'v') (('v' / 'V') ('a' / 'A') ('l' / 'L'))) | (&('I' | 'i') (('i' / 'I') ('d' / 'D') ('x' / 'X'))) | (&('S' | 's') (('s' / 'S') ('i' / 'I') ('z' / 'Z') ('e' / 'E'))) | (&('B' | 'b') (('b' / 'B') ('p' / 'P') ('o' / 'O') ('s' / 'S'))) | (&('C' | 'c') (('c' / 'C') ('h' / 'H') ('a' / 'A') ('i' / 'I') ('n' / 'N'))) | (&('E' | 'e') (('e' / 'E') ('l' / 'L') ('t' / 'T') ('s' / 'S'))) | (&('D' | 'd') (('d' / 'D') ('o' / 'O') ('m' / 'M') ('n' / 'N'))) | (&('T' | 't') (('t' / 'T') ('y' / 'Y') ('p' / 'P') ('e' / 'E'))) | (&('M' | 'm') (('m' / 'M') ('n' / 'N') ('g' / 'G') ('l' / 'L'))) | (&('N' | 'n') (('n' / 'N') ('a' / 'A') ('m' / 'M') ('e' / 'E'))) | (&('U' | 'u') (('u' / 'U') ('n' / 'N') ('q' / 'Q') ('l' / 'L'))) | (&('L' | 'l') (('l' / 'L') ('a' / 'A') ('b' / 'B') ('l' / 'L'))) | (&('P' | 'p') (('p' / 'P') ('t' / 'T') ('d' / 'D'))) | (&('F' | 'f') (('f' / 'F') ('l' / 'L') ('d' / 'D') ('s' / 'S'))) | (&('R' | 'r') (('r' / 'R') ('e' / 'E') ('f' / 'F') ('d' / 'D'))) | (&() Integer) | (&('A' | 'a') (('a' / 'A') ('r' / 'R') ('g' / 'G') ('t' / 'T'))) | (&('O' | 'o') (('o' / 'O') ('p' / 'P') '0')))) | (&('I' | 'i') (('i' / 'I') ('n' / 'N') ('i' / 'I') ('t' / 'T'))) | (&('S' | 's') (('s' / 'S') ('c' / 'C') ('p' / 'P') ('e' / 'E'))) | (&('B' | 'b') (('b' / 'B') ('i' / 'I') ('n' / 'N') ('f' / 'F'))) | (&('M' | 'm') (('m' / 'M') ('a' / 'A') ('x' / 'X'))) | (&('T' | 't') (('t' / 'T') ('y' / 'Y') ('p' / 'P') ('e' / 'E'))) | (&('C' | 'c') (('c' / 'C') ('s' / 'S') ('t' / 'T') ('s' / 'S'))) | (&('P' | 'p') (('p' / 'P') ('u' / 'U') ('r' / 'R') ('p' / 'P'))) | (&('V' | 'v') (('v' / 'V') ('a' / 'A') ('l' / 'L') ('u' / 'U'))) | (&('F' | 'f') (('f' / 'F') ('n' / 'N'))) | (&('E' | 'e') (('e' / 'E') ('x' / 'X') ('p' / 'P') ('r' / 'R'))) | (&('D' | 'd') (('d' / 'D') ('c' / 'C') ('l' / 'L') ('s' / 'S'))) | (&('O' | 'o') (('o' / 'O') ('r' / 'R') ('i' / 'I') ('g' / 'G'))) | (&('A' | 'a') (('a' / 'A') ('r' / 'R') ('g' / 'G') ('s' / 'S'))) | (&('R' | 'r') (('r' / 'R') ('e' / 'E') ('t' / 'T') ('n' / 'N')))))> Action0)> */
  func() bool {
   position0, tokenIndex0 := position, tokenIndex
   {
position1 := position
   {
position2 := position
   {
   position3, tokenIndex3 := position, tokenIndex
   {
   position5, tokenIndex5 := position, tokenIndex
   if buffer[position] != rune('p') {
   goto l6}
position++
   goto l5
   l6:	
   position, tokenIndex = position5, tokenIndex5
   if buffer[position] != rune('P') {
   goto l4}
position++
   }
   l5:	
   {
   position7, tokenIndex7 := position, tokenIndex
   if buffer[position] != rune('r') {
   goto l8}
position++
   goto l7
   l8:	
   position, tokenIndex = position7, tokenIndex7
   if buffer[position] != rune('R') {
   goto l4}
position++
   }
   l7:	
   {
   position9, tokenIndex9 := position, tokenIndex
   if buffer[position] != rune('m') {
   goto l10}
position++
   goto l9
   l10:	
   position, tokenIndex = position9, tokenIndex9
   if buffer[position] != rune('M') {
   goto l4}
position++
   }
   l9:	
   {
   position11, tokenIndex11 := position, tokenIndex
   if buffer[position] != rune('s') {
   goto l12}
position++
   goto l11
   l12:	
   position, tokenIndex = position11, tokenIndex11
   if buffer[position] != rune('S') {
   goto l4}
position++
   }
   l11:	
   goto l3
   l4:	
   position, tokenIndex = position3, tokenIndex3
   {
   position14, tokenIndex14 := position, tokenIndex
   if buffer[position] != rune('b') {
   goto l15}
position++
   goto l14
   l15:	
   position, tokenIndex = position14, tokenIndex14
   if buffer[position] != rune('B') {
   goto l13}
position++
   }
   l14:	
   {
   position16, tokenIndex16 := position, tokenIndex
   if buffer[position] != rune('a') {
   goto l17}
position++
   goto l16
   l17:	
   position, tokenIndex = position16, tokenIndex16
   if buffer[position] != rune('A') {
   goto l13}
position++
   }
   l16:	
   {
   position18, tokenIndex18 := position, tokenIndex
   if buffer[position] != rune('s') {
   goto l19}
position++
   goto l18
   l19:	
   position, tokenIndex = position18, tokenIndex18
   if buffer[position] != rune('S') {
   goto l13}
position++
   }
   l18:	
   {
   position20, tokenIndex20 := position, tokenIndex
   if buffer[position] != rune('e') {
   goto l21}
position++
   goto l20
   l21:	
   position, tokenIndex = position20, tokenIndex20
   if buffer[position] != rune('E') {
   goto l13}
position++
   }
   l20:	
   goto l3
   l13:	
   position, tokenIndex = position3, tokenIndex3
   {
   position23, tokenIndex23 := position, tokenIndex
   if buffer[position] != rune('d') {
   goto l24}
position++
   goto l23
   l24:	
   position, tokenIndex = position23, tokenIndex23
   if buffer[position] != rune('D') {
   goto l22}
position++
   }
   l23:	
   {
   position25, tokenIndex25 := position, tokenIndex
   if buffer[position] != rune('e') {
   goto l26}
position++
   goto l25
   l26:	
   position, tokenIndex = position25, tokenIndex25
   if buffer[position] != rune('E') {
   goto l22}
position++
   }
   l25:	
   {
   position27, tokenIndex27 := position, tokenIndex
   if buffer[position] != rune('c') {
   goto l28}
position++
   goto l27
   l28:	
   position, tokenIndex = position27, tokenIndex27
   if buffer[position] != rune('C') {
   goto l22}
position++
   }
   l27:	
   {
   position29, tokenIndex29 := position, tokenIndex
   if buffer[position] != rune('l') {
   goto l30}
position++
   goto l29
   l30:	
   position, tokenIndex = position29, tokenIndex29
   if buffer[position] != rune('L') {
   goto l22}
position++
   }
   l29:	
   goto l3
   l22:	
   position, tokenIndex = position3, tokenIndex3
   {
   position32, tokenIndex32 := position, tokenIndex
   if buffer[position] != rune('f') {
   goto l33}
position++
   goto l32
   l33:	
   position, tokenIndex = position32, tokenIndex32
   if buffer[position] != rune('F') {
   goto l31}
position++
   }
   l32:	
   {
   position34, tokenIndex34 := position, tokenIndex
   if buffer[position] != rune('n') {
   goto l35}
position++
   goto l34
   l35:	
   position, tokenIndex = position34, tokenIndex34
   if buffer[position] != rune('N') {
   goto l31}
position++
   }
   l34:	
   {
   position36, tokenIndex36 := position, tokenIndex
   if buffer[position] != rune('c') {
   goto l37}
position++
   goto l36
   l37:	
   position, tokenIndex = position36, tokenIndex36
   if buffer[position] != rune('C') {
   goto l31}
position++
   }
   l36:	
   {
   position38, tokenIndex38 := position, tokenIndex
   if buffer[position] != rune('s') {
   goto l39}
position++
   goto l38
   l39:	
   position, tokenIndex = position38, tokenIndex38
   if buffer[position] != rune('S') {
   goto l31}
position++
   }
   l38:	
   goto l3
   l31:	
   position, tokenIndex = position3, tokenIndex3
   {
   position41, tokenIndex41 := position, tokenIndex
   if buffer[position] != rune('c') {
   goto l42}
position++
   goto l41
   l42:	
   position, tokenIndex = position41, tokenIndex41
   if buffer[position] != rune('C') {
   goto l40}
position++
   }
   l41:	
   {
   position43, tokenIndex43 := position, tokenIndex
   if buffer[position] != rune('n') {
   goto l44}
position++
   goto l43
   l44:	
   position, tokenIndex = position43, tokenIndex43
   if buffer[position] != rune('N') {
   goto l40}
position++
   }
   l43:	
   {
   position45, tokenIndex45 := position, tokenIndex
   if buffer[position] != rune('s') {
   goto l46}
position++
   goto l45
   l46:	
   position, tokenIndex = position45, tokenIndex45
   if buffer[position] != rune('S') {
   goto l40}
position++
   }
   l45:	
   {
   position47, tokenIndex47 := position, tokenIndex
   if buffer[position] != rune('t') {
   goto l48}
position++
   goto l47
   l48:	
   position, tokenIndex = position47, tokenIndex47
   if buffer[position] != rune('T') {
   goto l40}
position++
   }
   l47:	
   goto l3
   l40:	
   position, tokenIndex = position3, tokenIndex3
   {
   position50, tokenIndex50 := position, tokenIndex
   if buffer[position] != rune('v') {
   goto l51}
position++
   goto l50
   l51:	
   position, tokenIndex = position50, tokenIndex50
   if buffer[position] != rune('V') {
   goto l49}
position++
   }
   l50:	
   {
   position52, tokenIndex52 := position, tokenIndex
   if buffer[position] != rune('a') {
   goto l53}
position++
   goto l52
   l53:	
   position, tokenIndex = position52, tokenIndex52
   if buffer[position] != rune('A') {
   goto l49}
position++
   }
   l52:	
   {
   position54, tokenIndex54 := position, tokenIndex
   if buffer[position] != rune('r') {
   goto l55}
position++
   goto l54
   l55:	
   position, tokenIndex = position54, tokenIndex54
   if buffer[position] != rune('R') {
   goto l49}
position++
   }
   l54:	
   {
   position56, tokenIndex56 := position, tokenIndex
   if buffer[position] != rune('s') {
   goto l57}
position++
   goto l56
   l57:	
   position, tokenIndex = position56, tokenIndex56
   if buffer[position] != rune('S') {
   goto l49}
position++
   }
   l56:	
   goto l3
   l49:	
   position, tokenIndex = position3, tokenIndex3
   {
   position59, tokenIndex59 := position, tokenIndex
   if buffer[position] != rune('v') {
   goto l60}
position++
   goto l59
   l60:	
   position, tokenIndex = position59, tokenIndex59
   if buffer[position] != rune('V') {
   goto l58}
position++
   }
   l59:	
   {
   position61, tokenIndex61 := position, tokenIndex
   if buffer[position] != rune('f') {
   goto l62}
position++
   goto l61
   l62:	
   position, tokenIndex = position61, tokenIndex61
   if buffer[position] != rune('F') {
   goto l58}
position++
   }
   l61:	
   {
   position63, tokenIndex63 := position, tokenIndex
   if buffer[position] != rune('l') {
   goto l64}
position++
   goto l63
   l64:	
   position, tokenIndex = position63, tokenIndex63
   if buffer[position] != rune('L') {
   goto l58}
position++
   }
   l63:	
   {
   position65, tokenIndex65 := position, tokenIndex
   if buffer[position] != rune('d') {
   goto l66}
position++
   goto l65
   l66:	
   position, tokenIndex = position65, tokenIndex65
   if buffer[position] != rune('D') {
   goto l58}
position++
   }
   l65:	
   goto l3
   l58:	
   position, tokenIndex = position3, tokenIndex3
   {
   position68, tokenIndex68 := position, tokenIndex
   if buffer[position] != rune('c') {
   goto l69}
position++
   goto l68
   l69:	
   position, tokenIndex = position68, tokenIndex68
   if buffer[position] != rune('C') {
   goto l67}
position++
   }
   l68:	
   {
   position70, tokenIndex70 := position, tokenIndex
   if buffer[position] != rune('h') {
   goto l71}
position++
   goto l70
   l71:	
   position, tokenIndex = position70, tokenIndex70
   if buffer[position] != rune('H') {
   goto l67}
position++
   }
   l70:	
   {
   position72, tokenIndex72 := position, tokenIndex
   if buffer[position] != rune('a') {
   goto l73}
position++
   goto l72
   l73:	
   position, tokenIndex = position72, tokenIndex72
   if buffer[position] != rune('A') {
   goto l67}
position++
   }
   l72:	
   {
   position74, tokenIndex74 := position, tokenIndex
   if buffer[position] != rune('n') {
   goto l75}
position++
   goto l74
   l75:	
   position, tokenIndex = position74, tokenIndex74
   if buffer[position] != rune('N') {
   goto l67}
position++
   }
   l74:	
   goto l3
   l67:	
   position, tokenIndex = position3, tokenIndex3
   {
   position77, tokenIndex77 := position, tokenIndex
   if buffer[position] != rune('c') {
   goto l78}
position++
   goto l77
   l78:	
   position, tokenIndex = position77, tokenIndex77
   if buffer[position] != rune('C') {
   goto l76}
position++
   }
   l77:	
   {
   position79, tokenIndex79 := position, tokenIndex
   if buffer[position] != rune('l') {
   goto l80}
position++
   goto l79
   l80:	
   position, tokenIndex = position79, tokenIndex79
   if buffer[position] != rune('L') {
   goto l76}
position++
   }
   l79:	
   {
   position81, tokenIndex81 := position, tokenIndex
   if buffer[position] != rune('s') {
   goto l82}
position++
   goto l81
   l82:	
   position, tokenIndex = position81, tokenIndex81
   if buffer[position] != rune('S') {
   goto l76}
position++
   }
   l81:	
   goto l3
   l76:	
   position, tokenIndex = position3, tokenIndex3
   {
   position84, tokenIndex84 := position, tokenIndex
   if buffer[position] != rune('m') {
   goto l85}
position++
   goto l84
   l85:	
   position, tokenIndex = position84, tokenIndex84
   if buffer[position] != rune('M') {
   goto l83}
position++
   }
   l84:	
   {
   position86, tokenIndex86 := position, tokenIndex
   if buffer[position] != rune('i') {
   goto l87}
position++
   goto l86
   l87:	
   position, tokenIndex = position86, tokenIndex86
   if buffer[position] != rune('I') {
   goto l83}
position++
   }
   l86:	
   {
   position88, tokenIndex88 := position, tokenIndex
   if buffer[position] != rune('n') {
   goto l89}
position++
   goto l88
   l89:	
   position, tokenIndex = position88, tokenIndex88
   if buffer[position] != rune('N') {
   goto l83}
position++
   }
   l88:	
   goto l3
   l83:	
   position, tokenIndex = position3, tokenIndex3
   {
   switch buffer[position] {
   case '<nil>':
   {
   switch buffer[position] {
   case 'V', 'v':
   {
   position92, tokenIndex92 := position, tokenIndex
   if buffer[position] != rune('v') {
   goto l93}
position++
   goto l92
   l93:	
   position, tokenIndex = position92, tokenIndex92
   if buffer[position] != rune('V') {
   goto l0}
position++
   }
   l92:	
   {
   position94, tokenIndex94 := position, tokenIndex
   if buffer[position] != rune('a') {
   goto l95}
position++
   goto l94
   l95:	
   position, tokenIndex = position94, tokenIndex94
   if buffer[position] != rune('A') {
   goto l0}
position++
   }
   l94:	
   {
   position96, tokenIndex96 := position, tokenIndex
   if buffer[position] != rune('l') {
   goto l97}
position++
   goto l96
   l97:	
   position, tokenIndex = position96, tokenIndex96
   if buffer[position] != rune('L') {
   goto l0}
position++
   }
   l96:	
break
   case 'I', 'i':
   {
   position98, tokenIndex98 := position, tokenIndex
   if buffer[position] != rune('i') {
   goto l99}
position++
   goto l98
   l99:	
   position, tokenIndex = position98, tokenIndex98
   if buffer[position] != rune('I') {
   goto l0}
position++
   }
   l98:	
   {
   position100, tokenIndex100 := position, tokenIndex
   if buffer[position] != rune('d') {
   goto l101}
position++
   goto l100
   l101:	
   position, tokenIndex = position100, tokenIndex100
   if buffer[position] != rune('D') {
   goto l0}
position++
   }
   l100:	
   {
   position102, tokenIndex102 := position, tokenIndex
   if buffer[position] != rune('x') {
   goto l103}
position++
   goto l102
   l103:	
   position, tokenIndex = position102, tokenIndex102
   if buffer[position] != rune('X') {
   goto l0}
position++
   }
   l102:	
break
   case 'S', 's':
   {
   position104, tokenIndex104 := position, tokenIndex
   if buffer[position] != rune('s') {
   goto l105}
position++
   goto l104
   l105:	
   position, tokenIndex = position104, tokenIndex104
   if buffer[position] != rune('S') {
   goto l0}
position++
   }
   l104:	
   {
   position106, tokenIndex106 := position, tokenIndex
   if buffer[position] != rune('i') {
   goto l107}
position++
   goto l106
   l107:	
   position, tokenIndex = position106, tokenIndex106
   if buffer[position] != rune('I') {
   goto l0}
position++
   }
   l106:	
   {
   position108, tokenIndex108 := position, tokenIndex
   if buffer[position] != rune('z') {
   goto l109}
position++
   goto l108
   l109:	
   position, tokenIndex = position108, tokenIndex108
   if buffer[position] != rune('Z') {
   goto l0}
position++
   }
   l108:	
   {
   position110, tokenIndex110 := position, tokenIndex
   if buffer[position] != rune('e') {
   goto l111}
position++
   goto l110
   l111:	
   position, tokenIndex = position110, tokenIndex110
   if buffer[position] != rune('E') {
   goto l0}
position++
   }
   l110:	
break
   case 'B', 'b':
   {
   position112, tokenIndex112 := position, tokenIndex
   if buffer[position] != rune('b') {
   goto l113}
position++
   goto l112
   l113:	
   position, tokenIndex = position112, tokenIndex112
   if buffer[position] != rune('B') {
   goto l0}
position++
   }
   l112:	
   {
   position114, tokenIndex114 := position, tokenIndex
   if buffer[position] != rune('p') {
   goto l115}
position++
   goto l114
   l115:	
   position, tokenIndex = position114, tokenIndex114
   if buffer[position] != rune('P') {
   goto l0}
position++
   }
   l114:	
   {
   position116, tokenIndex116 := position, tokenIndex
   if buffer[position] != rune('o') {
   goto l117}
position++
   goto l116
   l117:	
   position, tokenIndex = position116, tokenIndex116
   if buffer[position] != rune('O') {
   goto l0}
position++
   }
   l116:	
   {
   position118, tokenIndex118 := position, tokenIndex
   if buffer[position] != rune('s') {
   goto l119}
position++
   goto l118
   l119:	
   position, tokenIndex = position118, tokenIndex118
   if buffer[position] != rune('S') {
   goto l0}
position++
   }
   l118:	
break
   case 'C', 'c':
   {
   position120, tokenIndex120 := position, tokenIndex
   if buffer[position] != rune('c') {
   goto l121}
position++
   goto l120
   l121:	
   position, tokenIndex = position120, tokenIndex120
   if buffer[position] != rune('C') {
   goto l0}
position++
   }
   l120:	
   {
   position122, tokenIndex122 := position, tokenIndex
   if buffer[position] != rune('h') {
   goto l123}
position++
   goto l122
   l123:	
   position, tokenIndex = position122, tokenIndex122
   if buffer[position] != rune('H') {
   goto l0}
position++
   }
   l122:	
   {
   position124, tokenIndex124 := position, tokenIndex
   if buffer[position] != rune('a') {
   goto l125}
position++
   goto l124
   l125:	
   position, tokenIndex = position124, tokenIndex124
   if buffer[position] != rune('A') {
   goto l0}
position++
   }
   l124:	
   {
   position126, tokenIndex126 := position, tokenIndex
   if buffer[position] != rune('i') {
   goto l127}
position++
   goto l126
   l127:	
   position, tokenIndex = position126, tokenIndex126
   if buffer[position] != rune('I') {
   goto l0}
position++
   }
   l126:	
   {
   position128, tokenIndex128 := position, tokenIndex
   if buffer[position] != rune('n') {
   goto l129}
position++
   goto l128
   l129:	
   position, tokenIndex = position128, tokenIndex128
   if buffer[position] != rune('N') {
   goto l0}
position++
   }
   l128:	
break
   case 'E', 'e':
   {
   position130, tokenIndex130 := position, tokenIndex
   if buffer[position] != rune('e') {
   goto l131}
position++
   goto l130
   l131:	
   position, tokenIndex = position130, tokenIndex130
   if buffer[position] != rune('E') {
   goto l0}
position++
   }
   l130:	
   {
   position132, tokenIndex132 := position, tokenIndex
   if buffer[position] != rune('l') {
   goto l133}
position++
   goto l132
   l133:	
   position, tokenIndex = position132, tokenIndex132
   if buffer[position] != rune('L') {
   goto l0}
position++
   }
   l132:	
   {
   position134, tokenIndex134 := position, tokenIndex
   if buffer[position] != rune('t') {
   goto l135}
position++
   goto l134
   l135:	
   position, tokenIndex = position134, tokenIndex134
   if buffer[position] != rune('T') {
   goto l0}
position++
   }
   l134:	
   {
   position136, tokenIndex136 := position, tokenIndex
   if buffer[position] != rune('s') {
   goto l137}
position++
   goto l136
   l137:	
   position, tokenIndex = position136, tokenIndex136
   if buffer[position] != rune('S') {
   goto l0}
position++
   }
   l136:	
break
   case 'D', 'd':
   {
   position138, tokenIndex138 := position, tokenIndex
   if buffer[position] != rune('d') {
   goto l139}
position++
   goto l138
   l139:	
   position, tokenIndex = position138, tokenIndex138
   if buffer[position] != rune('D') {
   goto l0}
position++
   }
   l138:	
   {
   position140, tokenIndex140 := position, tokenIndex
   if buffer[position] != rune('o') {
   goto l141}
position++
   goto l140
   l141:	
   position, tokenIndex = position140, tokenIndex140
   if buffer[position] != rune('O') {
   goto l0}
position++
   }
   l140:	
   {
   position142, tokenIndex142 := position, tokenIndex
   if buffer[position] != rune('m') {
   goto l143}
position++
   goto l142
   l143:	
   position, tokenIndex = position142, tokenIndex142
   if buffer[position] != rune('M') {
   goto l0}
position++
   }
   l142:	
   {
   position144, tokenIndex144 := position, tokenIndex
   if buffer[position] != rune('n') {
   goto l145}
position++
   goto l144
   l145:	
   position, tokenIndex = position144, tokenIndex144
   if buffer[position] != rune('N') {
   goto l0}
position++
   }
   l144:	
break
   case 'T', 't':
   {
   position146, tokenIndex146 := position, tokenIndex
   if buffer[position] != rune('t') {
   goto l147}
position++
   goto l146
   l147:	
   position, tokenIndex = position146, tokenIndex146
   if buffer[position] != rune('T') {
   goto l0}
position++
   }
   l146:	
   {
   position148, tokenIndex148 := position, tokenIndex
   if buffer[position] != rune('y') {
   goto l149}
position++
   goto l148
   l149:	
   position, tokenIndex = position148, tokenIndex148
   if buffer[position] != rune('Y') {
   goto l0}
position++
   }
   l148:	
   {
   position150, tokenIndex150 := position, tokenIndex
   if buffer[position] != rune('p') {
   goto l151}
position++
   goto l150
   l151:	
   position, tokenIndex = position150, tokenIndex150
   if buffer[position] != rune('P') {
   goto l0}
position++
   }
   l150:	
   {
   position152, tokenIndex152 := position, tokenIndex
   if buffer[position] != rune('e') {
   goto l153}
position++
   goto l152
   l153:	
   position, tokenIndex = position152, tokenIndex152
   if buffer[position] != rune('E') {
   goto l0}
position++
   }
   l152:	
break
   case 'M', 'm':
   {
   position154, tokenIndex154 := position, tokenIndex
   if buffer[position] != rune('m') {
   goto l155}
position++
   goto l154
   l155:	
   position, tokenIndex = position154, tokenIndex154
   if buffer[position] != rune('M') {
   goto l0}
position++
   }
   l154:	
   {
   position156, tokenIndex156 := position, tokenIndex
   if buffer[position] != rune('n') {
   goto l157}
position++
   goto l156
   l157:	
   position, tokenIndex = position156, tokenIndex156
   if buffer[position] != rune('N') {
   goto l0}
position++
   }
   l156:	
   {
   position158, tokenIndex158 := position, tokenIndex
   if buffer[position] != rune('g') {
   goto l159}
position++
   goto l158
   l159:	
   position, tokenIndex = position158, tokenIndex158
   if buffer[position] != rune('G') {
   goto l0}
position++
   }
   l158:	
   {
   position160, tokenIndex160 := position, tokenIndex
   if buffer[position] != rune('l') {
   goto l161}
position++
   goto l160
   l161:	
   position, tokenIndex = position160, tokenIndex160
   if buffer[position] != rune('L') {
   goto l0}
position++
   }
   l160:	
break
   case 'N', 'n':
   {
   position162, tokenIndex162 := position, tokenIndex
   if buffer[position] != rune('n') {
   goto l163}
position++
   goto l162
   l163:	
   position, tokenIndex = position162, tokenIndex162
   if buffer[position] != rune('N') {
   goto l0}
position++
   }
   l162:	
   {
   position164, tokenIndex164 := position, tokenIndex
   if buffer[position] != rune('a') {
   goto l165}
position++
   goto l164
   l165:	
   position, tokenIndex = position164, tokenIndex164
   if buffer[position] != rune('A') {
   goto l0}
position++
   }
   l164:	
   {
   position166, tokenIndex166 := position, tokenIndex
   if buffer[position] != rune('m') {
   goto l167}
position++
   goto l166
   l167:	
   position, tokenIndex = position166, tokenIndex166
   if buffer[position] != rune('M') {
   goto l0}
position++
   }
   l166:	
   {
   position168, tokenIndex168 := position, tokenIndex
   if buffer[position] != rune('e') {
   goto l169}
position++
   goto l168
   l169:	
   position, tokenIndex = position168, tokenIndex168
   if buffer[position] != rune('E') {
   goto l0}
position++
   }
   l168:	
break
   case 'U', 'u':
   {
   position170, tokenIndex170 := position, tokenIndex
   if buffer[position] != rune('u') {
   goto l171}
position++
   goto l170
   l171:	
   position, tokenIndex = position170, tokenIndex170
   if buffer[position] != rune('U') {
   goto l0}
position++
   }
   l170:	
   {
   position172, tokenIndex172 := position, tokenIndex
   if buffer[position] != rune('n') {
   goto l173}
position++
   goto l172
   l173:	
   position, tokenIndex = position172, tokenIndex172
   if buffer[position] != rune('N') {
   goto l0}
position++
   }
   l172:	
   {
   position174, tokenIndex174 := position, tokenIndex
   if buffer[position] != rune('q') {
   goto l175}
position++
   goto l174
   l175:	
   position, tokenIndex = position174, tokenIndex174
   if buffer[position] != rune('Q') {
   goto l0}
position++
   }
   l174:	
   {
   position176, tokenIndex176 := position, tokenIndex
   if buffer[position] != rune('l') {
   goto l177}
position++
   goto l176
   l177:	
   position, tokenIndex = position176, tokenIndex176
   if buffer[position] != rune('L') {
   goto l0}
position++
   }
   l176:	
break
   case 'L', 'l':
   {
   position178, tokenIndex178 := position, tokenIndex
   if buffer[position] != rune('l') {
   goto l179}
position++
   goto l178
   l179:	
   position, tokenIndex = position178, tokenIndex178
   if buffer[position] != rune('L') {
   goto l0}
position++
   }
   l178:	
   {
   position180, tokenIndex180 := position, tokenIndex
   if buffer[position] != rune('a') {
   goto l181}
position++
   goto l180
   l181:	
   position, tokenIndex = position180, tokenIndex180
   if buffer[position] != rune('A') {
   goto l0}
position++
   }
   l180:	
   {
   position182, tokenIndex182 := position, tokenIndex
   if buffer[position] != rune('b') {
   goto l183}
position++
   goto l182
   l183:	
   position, tokenIndex = position182, tokenIndex182
   if buffer[position] != rune('B') {
   goto l0}
position++
   }
   l182:	
   {
   position184, tokenIndex184 := position, tokenIndex
   if buffer[position] != rune('l') {
   goto l185}
position++
   goto l184
   l185:	
   position, tokenIndex = position184, tokenIndex184
   if buffer[position] != rune('L') {
   goto l0}
position++
   }
   l184:	
break
   case 'P', 'p':
   {
   position186, tokenIndex186 := position, tokenIndex
   if buffer[position] != rune('p') {
   goto l187}
position++
   goto l186
   l187:	
   position, tokenIndex = position186, tokenIndex186
   if buffer[position] != rune('P') {
   goto l0}
position++
   }
   l186:	
   {
   position188, tokenIndex188 := position, tokenIndex
   if buffer[position] != rune('t') {
   goto l189}
position++
   goto l188
   l189:	
   position, tokenIndex = position188, tokenIndex188
   if buffer[position] != rune('T') {
   goto l0}
position++
   }
   l188:	
   {
   position190, tokenIndex190 := position, tokenIndex
   if buffer[position] != rune('d') {
   goto l191}
position++
   goto l190
   l191:	
   position, tokenIndex = position190, tokenIndex190
   if buffer[position] != rune('D') {
   goto l0}
position++
   }
   l190:	
break
   case 'F', 'f':
   {
   position192, tokenIndex192 := position, tokenIndex
   if buffer[position] != rune('f') {
   goto l193}
position++
   goto l192
   l193:	
   position, tokenIndex = position192, tokenIndex192
   if buffer[position] != rune('F') {
   goto l0}
position++
   }
   l192:	
   {
   position194, tokenIndex194 := position, tokenIndex
   if buffer[position] != rune('l') {
   goto l195}
position++
   goto l194
   l195:	
   position, tokenIndex = position194, tokenIndex194
   if buffer[position] != rune('L') {
   goto l0}
position++
   }
   l194:	
   {
   position196, tokenIndex196 := position, tokenIndex
   if buffer[position] != rune('d') {
   goto l197}
position++
   goto l196
   l197:	
   position, tokenIndex = position196, tokenIndex196
   if buffer[position] != rune('D') {
   goto l0}
position++
   }
   l196:	
   {
   position198, tokenIndex198 := position, tokenIndex
   if buffer[position] != rune('s') {
   goto l199}
position++
   goto l198
   l199:	
   position, tokenIndex = position198, tokenIndex198
   if buffer[position] != rune('S') {
   goto l0}
position++
   }
   l198:	
break
   case 'R', 'r':
   {
   position200, tokenIndex200 := position, tokenIndex
   if buffer[position] != rune('r') {
   goto l201}
position++
   goto l200
   l201:	
   position, tokenIndex = position200, tokenIndex200
   if buffer[position] != rune('R') {
   goto l0}
position++
   }
   l200:	
   {
   position202, tokenIndex202 := position, tokenIndex
   if buffer[position] != rune('e') {
   goto l203}
position++
   goto l202
   l203:	
   position, tokenIndex = position202, tokenIndex202
   if buffer[position] != rune('E') {
   goto l0}
position++
   }
   l202:	
   {
   position204, tokenIndex204 := position, tokenIndex
   if buffer[position] != rune('f') {
   goto l205}
position++
   goto l204
   l205:	
   position, tokenIndex = position204, tokenIndex204
   if buffer[position] != rune('F') {
   goto l0}
position++
   }
   l204:	
   {
   position206, tokenIndex206 := position, tokenIndex
   if buffer[position] != rune('d') {
   goto l207}
position++
   goto l206
   l207:	
   position, tokenIndex = position206, tokenIndex206
   if buffer[position] != rune('D') {
   goto l0}
position++
   }
   l206:	
break
   case '<nil>':
   {
position208 := position
add(ruleInteger, position208)
   }
break
   case 'A', 'a':
   {
   position209, tokenIndex209 := position, tokenIndex
   if buffer[position] != rune('a') {
   goto l210}
position++
   goto l209
   l210:	
   position, tokenIndex = position209, tokenIndex209
   if buffer[position] != rune('A') {
   goto l0}
position++
   }
   l209:	
   {
   position211, tokenIndex211 := position, tokenIndex
   if buffer[position] != rune('r') {
   goto l212}
position++
   goto l211
   l212:	
   position, tokenIndex = position211, tokenIndex211
   if buffer[position] != rune('R') {
   goto l0}
position++
   }
   l211:	
   {
   position213, tokenIndex213 := position, tokenIndex
   if buffer[position] != rune('g') {
   goto l214}
position++
   goto l213
   l214:	
   position, tokenIndex = position213, tokenIndex213
   if buffer[position] != rune('G') {
   goto l0}
position++
   }
   l213:	
   {
   position215, tokenIndex215 := position, tokenIndex
   if buffer[position] != rune('t') {
   goto l216}
position++
   goto l215
   l216:	
   position, tokenIndex = position215, tokenIndex215
   if buffer[position] != rune('T') {
   goto l0}
position++
   }
   l215:	
break
   default:
   {
   position217, tokenIndex217 := position, tokenIndex
   if buffer[position] != rune('o') {
   goto l218}
position++
   goto l217
   l218:	
   position, tokenIndex = position217, tokenIndex217
   if buffer[position] != rune('O') {
   goto l0}
position++
   }
   l217:	
   {
   position219, tokenIndex219 := position, tokenIndex
   if buffer[position] != rune('p') {
   goto l220}
position++
   goto l219
   l220:	
   position, tokenIndex = position219, tokenIndex219
   if buffer[position] != rune('P') {
   goto l0}
position++
   }
   l219:	
   if buffer[position] != rune('0') {
   goto l0}
position++
break
   }
   }

break
   case 'I', 'i':
   {
   position221, tokenIndex221 := position, tokenIndex
   if buffer[position] != rune('i') {
   goto l222}
position++
   goto l221
   l222:	
   position, tokenIndex = position221, tokenIndex221
   if buffer[position] != rune('I') {
   goto l0}
position++
   }
   l221:	
   {
   position223, tokenIndex223 := position, tokenIndex
   if buffer[position] != rune('n') {
   goto l224}
position++
   goto l223
   l224:	
   position, tokenIndex = position223, tokenIndex223
   if buffer[position] != rune('N') {
   goto l0}
position++
   }
   l223:	
   {
   position225, tokenIndex225 := position, tokenIndex
   if buffer[position] != rune('i') {
   goto l226}
position++
   goto l225
   l226:	
   position, tokenIndex = position225, tokenIndex225
   if buffer[position] != rune('I') {
   goto l0}
position++
   }
   l225:	
   {
   position227, tokenIndex227 := position, tokenIndex
   if buffer[position] != rune('t') {
   goto l228}
position++
   goto l227
   l228:	
   position, tokenIndex = position227, tokenIndex227
   if buffer[position] != rune('T') {
   goto l0}
position++
   }
   l227:	
break
   case 'S', 's':
   {
   position229, tokenIndex229 := position, tokenIndex
   if buffer[position] != rune('s') {
   goto l230}
position++
   goto l229
   l230:	
   position, tokenIndex = position229, tokenIndex229
   if buffer[position] != rune('S') {
   goto l0}
position++
   }
   l229:	
   {
   position231, tokenIndex231 := position, tokenIndex
   if buffer[position] != rune('c') {
   goto l232}
position++
   goto l231
   l232:	
   position, tokenIndex = position231, tokenIndex231
   if buffer[position] != rune('C') {
   goto l0}
position++
   }
   l231:	
   {
   position233, tokenIndex233 := position, tokenIndex
   if buffer[position] != rune('p') {
   goto l234}
position++
   goto l233
   l234:	
   position, tokenIndex = position233, tokenIndex233
   if buffer[position] != rune('P') {
   goto l0}
position++
   }
   l233:	
   {
   position235, tokenIndex235 := position, tokenIndex
   if buffer[position] != rune('e') {
   goto l236}
position++
   goto l235
   l236:	
   position, tokenIndex = position235, tokenIndex235
   if buffer[position] != rune('E') {
   goto l0}
position++
   }
   l235:	
break
   case 'B', 'b':
   {
   position237, tokenIndex237 := position, tokenIndex
   if buffer[position] != rune('b') {
   goto l238}
position++
   goto l237
   l238:	
   position, tokenIndex = position237, tokenIndex237
   if buffer[position] != rune('B') {
   goto l0}
position++
   }
   l237:	
   {
   position239, tokenIndex239 := position, tokenIndex
   if buffer[position] != rune('i') {
   goto l240}
position++
   goto l239
   l240:	
   position, tokenIndex = position239, tokenIndex239
   if buffer[position] != rune('I') {
   goto l0}
position++
   }
   l239:	
   {
   position241, tokenIndex241 := position, tokenIndex
   if buffer[position] != rune('n') {
   goto l242}
position++
   goto l241
   l242:	
   position, tokenIndex = position241, tokenIndex241
   if buffer[position] != rune('N') {
   goto l0}
position++
   }
   l241:	
   {
   position243, tokenIndex243 := position, tokenIndex
   if buffer[position] != rune('f') {
   goto l244}
position++
   goto l243
   l244:	
   position, tokenIndex = position243, tokenIndex243
   if buffer[position] != rune('F') {
   goto l0}
position++
   }
   l243:	
break
   case 'M', 'm':
   {
   position245, tokenIndex245 := position, tokenIndex
   if buffer[position] != rune('m') {
   goto l246}
position++
   goto l245
   l246:	
   position, tokenIndex = position245, tokenIndex245
   if buffer[position] != rune('M') {
   goto l0}
position++
   }
   l245:	
   {
   position247, tokenIndex247 := position, tokenIndex
   if buffer[position] != rune('a') {
   goto l248}
position++
   goto l247
   l248:	
   position, tokenIndex = position247, tokenIndex247
   if buffer[position] != rune('A') {
   goto l0}
position++
   }
   l247:	
   {
   position249, tokenIndex249 := position, tokenIndex
   if buffer[position] != rune('x') {
   goto l250}
position++
   goto l249
   l250:	
   position, tokenIndex = position249, tokenIndex249
   if buffer[position] != rune('X') {
   goto l0}
position++
   }
   l249:	
break
   case 'T', 't':
   {
   position251, tokenIndex251 := position, tokenIndex
   if buffer[position] != rune('t') {
   goto l252}
position++
   goto l251
   l252:	
   position, tokenIndex = position251, tokenIndex251
   if buffer[position] != rune('T') {
   goto l0}
position++
   }
   l251:	
   {
   position253, tokenIndex253 := position, tokenIndex
   if buffer[position] != rune('y') {
   goto l254}
position++
   goto l253
   l254:	
   position, tokenIndex = position253, tokenIndex253
   if buffer[position] != rune('Y') {
   goto l0}
position++
   }
   l253:	
   {
   position255, tokenIndex255 := position, tokenIndex
   if buffer[position] != rune('p') {
   goto l256}
position++
   goto l255
   l256:	
   position, tokenIndex = position255, tokenIndex255
   if buffer[position] != rune('P') {
   goto l0}
position++
   }
   l255:	
   {
   position257, tokenIndex257 := position, tokenIndex
   if buffer[position] != rune('e') {
   goto l258}
position++
   goto l257
   l258:	
   position, tokenIndex = position257, tokenIndex257
   if buffer[position] != rune('E') {
   goto l0}
position++
   }
   l257:	
break
   case 'C', 'c':
   {
   position259, tokenIndex259 := position, tokenIndex
   if buffer[position] != rune('c') {
   goto l260}
position++
   goto l259
   l260:	
   position, tokenIndex = position259, tokenIndex259
   if buffer[position] != rune('C') {
   goto l0}
position++
   }
   l259:	
   {
   position261, tokenIndex261 := position, tokenIndex
   if buffer[position] != rune('s') {
   goto l262}
position++
   goto l261
   l262:	
   position, tokenIndex = position261, tokenIndex261
   if buffer[position] != rune('S') {
   goto l0}
position++
   }
   l261:	
   {
   position263, tokenIndex263 := position, tokenIndex
   if buffer[position] != rune('t') {
   goto l264}
position++
   goto l263
   l264:	
   position, tokenIndex = position263, tokenIndex263
   if buffer[position] != rune('T') {
   goto l0}
position++
   }
   l263:	
   {
   position265, tokenIndex265 := position, tokenIndex
   if buffer[position] != rune('s') {
   goto l266}
position++
   goto l265
   l266:	
   position, tokenIndex = position265, tokenIndex265
   if buffer[position] != rune('S') {
   goto l0}
position++
   }
   l265:	
break
   case 'P', 'p':
   {
   position267, tokenIndex267 := position, tokenIndex
   if buffer[position] != rune('p') {
   goto l268}
position++
   goto l267
   l268:	
   position, tokenIndex = position267, tokenIndex267
   if buffer[position] != rune('P') {
   goto l0}
position++
   }
   l267:	
   {
   position269, tokenIndex269 := position, tokenIndex
   if buffer[position] != rune('u') {
   goto l270}
position++
   goto l269
   l270:	
   position, tokenIndex = position269, tokenIndex269
   if buffer[position] != rune('U') {
   goto l0}
position++
   }
   l269:	
   {
   position271, tokenIndex271 := position, tokenIndex
   if buffer[position] != rune('r') {
   goto l272}
position++
   goto l271
   l272:	
   position, tokenIndex = position271, tokenIndex271
   if buffer[position] != rune('R') {
   goto l0}
position++
   }
   l271:	
   {
   position273, tokenIndex273 := position, tokenIndex
   if buffer[position] != rune('p') {
   goto l274}
position++
   goto l273
   l274:	
   position, tokenIndex = position273, tokenIndex273
   if buffer[position] != rune('P') {
   goto l0}
position++
   }
   l273:	
break
   case 'V', 'v':
   {
   position275, tokenIndex275 := position, tokenIndex
   if buffer[position] != rune('v') {
   goto l276}
position++
   goto l275
   l276:	
   position, tokenIndex = position275, tokenIndex275
   if buffer[position] != rune('V') {
   goto l0}
position++
   }
   l275:	
   {
   position277, tokenIndex277 := position, tokenIndex
   if buffer[position] != rune('a') {
   goto l278}
position++
   goto l277
   l278:	
   position, tokenIndex = position277, tokenIndex277
   if buffer[position] != rune('A') {
   goto l0}
position++
   }
   l277:	
   {
   position279, tokenIndex279 := position, tokenIndex
   if buffer[position] != rune('l') {
   goto l280}
position++
   goto l279
   l280:	
   position, tokenIndex = position279, tokenIndex279
   if buffer[position] != rune('L') {
   goto l0}
position++
   }
   l279:	
   {
   position281, tokenIndex281 := position, tokenIndex
   if buffer[position] != rune('u') {
   goto l282}
position++
   goto l281
   l282:	
   position, tokenIndex = position281, tokenIndex281
   if buffer[position] != rune('U') {
   goto l0}
position++
   }
   l281:	
break
   case 'F', 'f':
   {
   position283, tokenIndex283 := position, tokenIndex
   if buffer[position] != rune('f') {
   goto l284}
position++
   goto l283
   l284:	
   position, tokenIndex = position283, tokenIndex283
   if buffer[position] != rune('F') {
   goto l0}
position++
   }
   l283:	
   {
   position285, tokenIndex285 := position, tokenIndex
   if buffer[position] != rune('n') {
   goto l286}
position++
   goto l285
   l286:	
   position, tokenIndex = position285, tokenIndex285
   if buffer[position] != rune('N') {
   goto l0}
position++
   }
   l285:	
break
   case 'E', 'e':
   {
   position287, tokenIndex287 := position, tokenIndex
   if buffer[position] != rune('e') {
   goto l288}
position++
   goto l287
   l288:	
   position, tokenIndex = position287, tokenIndex287
   if buffer[position] != rune('E') {
   goto l0}
position++
   }
   l287:	
   {
   position289, tokenIndex289 := position, tokenIndex
   if buffer[position] != rune('x') {
   goto l290}
position++
   goto l289
   l290:	
   position, tokenIndex = position289, tokenIndex289
   if buffer[position] != rune('X') {
   goto l0}
position++
   }
   l289:	
   {
   position291, tokenIndex291 := position, tokenIndex
   if buffer[position] != rune('p') {
   goto l292}
position++
   goto l291
   l292:	
   position, tokenIndex = position291, tokenIndex291
   if buffer[position] != rune('P') {
   goto l0}
position++
   }
   l291:	
   {
   position293, tokenIndex293 := position, tokenIndex
   if buffer[position] != rune('r') {
   goto l294}
position++
   goto l293
   l294:	
   position, tokenIndex = position293, tokenIndex293
   if buffer[position] != rune('R') {
   goto l0}
position++
   }
   l293:	
break
   case 'D', 'd':
   {
   position295, tokenIndex295 := position, tokenIndex
   if buffer[position] != rune('d') {
   goto l296}
position++
   goto l295
   l296:	
   position, tokenIndex = position295, tokenIndex295
   if buffer[position] != rune('D') {
   goto l0}
position++
   }
   l295:	
   {
   position297, tokenIndex297 := position, tokenIndex
   if buffer[position] != rune('c') {
   goto l298}
position++
   goto l297
   l298:	
   position, tokenIndex = position297, tokenIndex297
   if buffer[position] != rune('C') {
   goto l0}
position++
   }
   l297:	
   {
   position299, tokenIndex299 := position, tokenIndex
   if buffer[position] != rune('l') {
   goto l300}
position++
   goto l299
   l300:	
   position, tokenIndex = position299, tokenIndex299
   if buffer[position] != rune('L') {
   goto l0}
position++
   }
   l299:	
   {
   position301, tokenIndex301 := position, tokenIndex
   if buffer[position] != rune('s') {
   goto l302}
position++
   goto l301
   l302:	
   position, tokenIndex = position301, tokenIndex301
   if buffer[position] != rune('S') {
   goto l0}
position++
   }
   l301:	
break
   case 'O', 'o':
   {
   position303, tokenIndex303 := position, tokenIndex
   if buffer[position] != rune('o') {
   goto l304}
position++
   goto l303
   l304:	
   position, tokenIndex = position303, tokenIndex303
   if buffer[position] != rune('O') {
   goto l0}
position++
   }
   l303:	
   {
   position305, tokenIndex305 := position, tokenIndex
   if buffer[position] != rune('r') {
   goto l306}
position++
   goto l305
   l306:	
   position, tokenIndex = position305, tokenIndex305
   if buffer[position] != rune('R') {
   goto l0}
position++
   }
   l305:	
   {
   position307, tokenIndex307 := position, tokenIndex
   if buffer[position] != rune('i') {
   goto l308}
position++
   goto l307
   l308:	
   position, tokenIndex = position307, tokenIndex307
   if buffer[position] != rune('I') {
   goto l0}
position++
   }
   l307:	
   {
   position309, tokenIndex309 := position, tokenIndex
   if buffer[position] != rune('g') {
   goto l310}
position++
   goto l309
   l310:	
   position, tokenIndex = position309, tokenIndex309
   if buffer[position] != rune('G') {
   goto l0}
position++
   }
   l309:	
break
   case 'A', 'a':
   {
   position311, tokenIndex311 := position, tokenIndex
   if buffer[position] != rune('a') {
   goto l312}
position++
   goto l311
   l312:	
   position, tokenIndex = position311, tokenIndex311
   if buffer[position] != rune('A') {
   goto l0}
position++
   }
   l311:	
   {
   position313, tokenIndex313 := position, tokenIndex
   if buffer[position] != rune('r') {
   goto l314}
position++
   goto l313
   l314:	
   position, tokenIndex = position313, tokenIndex313
   if buffer[position] != rune('R') {
   goto l0}
position++
   }
   l313:	
   {
   position315, tokenIndex315 := position, tokenIndex
   if buffer[position] != rune('g') {
   goto l316}
position++
   goto l315
   l316:	
   position, tokenIndex = position315, tokenIndex315
   if buffer[position] != rune('G') {
   goto l0}
position++
   }
   l315:	
   {
   position317, tokenIndex317 := position, tokenIndex
   if buffer[position] != rune('s') {
   goto l318}
position++
   goto l317
   l318:	
   position, tokenIndex = position317, tokenIndex317
   if buffer[position] != rune('S') {
   goto l0}
position++
   }
   l317:	
break
   default:
   {
   position319, tokenIndex319 := position, tokenIndex
   if buffer[position] != rune('r') {
   goto l320}
position++
   goto l319
   l320:	
   position, tokenIndex = position319, tokenIndex319
   if buffer[position] != rune('R') {
   goto l0}
position++
   }
   l319:	
   {
   position321, tokenIndex321 := position, tokenIndex
   if buffer[position] != rune('e') {
   goto l322}
position++
   goto l321
   l322:	
   position, tokenIndex = position321, tokenIndex321
   if buffer[position] != rune('E') {
   goto l0}
position++
   }
   l321:	
   {
   position323, tokenIndex323 := position, tokenIndex
   if buffer[position] != rune('t') {
   goto l324}
position++
   goto l323
   l324:	
   position, tokenIndex = position323, tokenIndex323
   if buffer[position] != rune('T') {
   goto l0}
position++
   }
   l323:	
   {
   position325, tokenIndex325 := position, tokenIndex
   if buffer[position] != rune('n') {
   goto l326}
position++
   goto l325
   l326:	
   position, tokenIndex = position325, tokenIndex325
   if buffer[position] != rune('N') {
   goto l0}
position++
   }
   l325:	
break
   }
   }

   }
   l3:	
add(rulePegText, position2)
   }
   {
add(ruleAction0, position)
   }
add(ruleNodeFieldName, position1)
   }
   return true
   l0:	
   position, tokenIndex = position0, tokenIndex0
   return false
  },
  nil,
  nil,
  /* 4 Action0 <- <{

	FieldName=buffer[begin:end]
   }> */
  nil,
 }
 p.rules = _rules
}
