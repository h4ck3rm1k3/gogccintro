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
	ruleSpecValue
	rulews
	rulePegText
	ruleAction0
)

var rul3s = [...]string{
	"Unknown",
	"SpecValue",
	"ws",
	"PegText",
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
				print(node.up, depth+1)
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
	tree []token32
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
		expanded := make([]token32, 2*len(tree))
		copy(expanded, tree)
		t.tree = expanded
	}
	t.tree[index] = token32{
		pegRule: rule,
		begin:   begin,
		end:     end,
	}
}

func (t *tokens32) Tokens() []token32 {
	return t.tree
}

type GccNode struct {
	Buffer string
	buffer []rune
	rules  [5]func() bool
	parse  func(rule ...int) error
	reset  func()
	Pretty bool
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

type textPositionMap map[int]textPosition

func translatePositions(buffer []rune, positions []int) textPositionMap {
	length, translations, j, line, symbol := len(positions), make(textPositionMap, len(positions)), 0, 1, 0
	sort.Ints(positions)

search:
	for i, c := range buffer {
		if c == '\n' {
			line, symbol = line+1, 0
		} else {
			symbol++
		}
		if i == positions[j] {
			translations[positions[j]] = textPosition{line, symbol}
			for j++; j < length; j++ {
				if i != positions[j] {
					continue search
				}
			}
			break search
		}
	}

	return translations
}

type parseError struct {
	p   *GccNode
	max token32
}

func (e *parseError) Error2() string {
	tokens, error := []token32{e.max}, "\n"
	positions, p := make([]int, 2*len(tokens)), 0
	for _, token := range tokens {
		positions[p], p = int(token.begin), p+1
		positions[p], p = int(token.end), p+1
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
		switch token.pegRule {

		case rulePegText:
			begin, end = int(token.begin), int(token.end)
			text = string(_buffer[begin:end])

		case ruleAction0:

			s := astproto.Spec(astproto.Spec_value[buffer[begin:end]])
			getNode().AddSpec(s)

		}
	}
	_, _, _, _, _ = buffer, _buffer, text, begin, end
}

func (p *GccNode) Init() {
	var (
		max                  token32
		position, tokenIndex uint32
		buffer               []rune
	)
	p.reset = func() {
		max = token32{}
		position, tokenIndex = 0, 0

		p.buffer = []rune(p.Buffer)
		if len(p.buffer) == 0 || p.buffer[len(p.buffer)-1] != endSymbol {
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

	_rules = [...]func() bool{
		nil,
		/* 0 SpecValue <- <(ws (('s' / 'S') ('p' / 'P') ('e' / 'E') ('c' / 'C') ':')? ws (<((&('R' | 'r') (('r' / 'R') ('e' / 'E') ('g' / 'G') ('i' / 'I') ('s' / 'S') ('t' / 'T') ('e' / 'E') ('r' / 'R'))) | (&('V' | 'v') (('v' / 'V') ('i' / 'I') ('r' / 'R') ('t' / 'T'))) | (&('P' | 'p') (('p' / 'P') ('u' / 'U') ('r' / 'R') ('e' / 'E'))) | (&('B' | 'b') (('b' / 'B') ('i' / 'I') ('t' / 'T') ('f' / 'F') ('i' / 'I') ('e' / 'E') ('l' / 'L') ('d' / 'D'))) | (&('M' | 'm') (('m' / 'M') ('u' / 'U') ('t' / 'T') ('a' / 'A') ('b' / 'B') ('l' / 'L') ('e' / 'E'))))> ws)+ Action0)> */
		func() bool {
			position0, tokenIndex0 := position, tokenIndex
			{
				position1 := position
				if !_rules[rulews]() {
					goto l0
				}
				{
					position2, tokenIndex2 := position, tokenIndex
					{
						position4, tokenIndex4 := position, tokenIndex
						if buffer[position] != rune('s') {
							goto l5
						}
						position++
						goto l4
					l5:
						position, tokenIndex = position4, tokenIndex4
						if buffer[position] != rune('S') {
							goto l2
						}
						position++
					}
				l4:
					{
						position6, tokenIndex6 := position, tokenIndex
						if buffer[position] != rune('p') {
							goto l7
						}
						position++
						goto l6
					l7:
						position, tokenIndex = position6, tokenIndex6
						if buffer[position] != rune('P') {
							goto l2
						}
						position++
					}
				l6:
					{
						position8, tokenIndex8 := position, tokenIndex
						if buffer[position] != rune('e') {
							goto l9
						}
						position++
						goto l8
					l9:
						position, tokenIndex = position8, tokenIndex8
						if buffer[position] != rune('E') {
							goto l2
						}
						position++
					}
				l8:
					{
						position10, tokenIndex10 := position, tokenIndex
						if buffer[position] != rune('c') {
							goto l11
						}
						position++
						goto l10
					l11:
						position, tokenIndex = position10, tokenIndex10
						if buffer[position] != rune('C') {
							goto l2
						}
						position++
					}
				l10:
					if buffer[position] != rune(':') {
						goto l2
					}
					position++
					goto l3
				l2:
					position, tokenIndex = position2, tokenIndex2
				}
			l3:
				if !_rules[rulews]() {
					goto l0
				}
				{
					position14 := position
					{
						switch buffer[position] {
						case 'R', 'r':
							{
								position16, tokenIndex16 := position, tokenIndex
								if buffer[position] != rune('r') {
									goto l17
								}
								position++
								goto l16
							l17:
								position, tokenIndex = position16, tokenIndex16
								if buffer[position] != rune('R') {
									goto l0
								}
								position++
							}
						l16:
							{
								position18, tokenIndex18 := position, tokenIndex
								if buffer[position] != rune('e') {
									goto l19
								}
								position++
								goto l18
							l19:
								position, tokenIndex = position18, tokenIndex18
								if buffer[position] != rune('E') {
									goto l0
								}
								position++
							}
						l18:
							{
								position20, tokenIndex20 := position, tokenIndex
								if buffer[position] != rune('g') {
									goto l21
								}
								position++
								goto l20
							l21:
								position, tokenIndex = position20, tokenIndex20
								if buffer[position] != rune('G') {
									goto l0
								}
								position++
							}
						l20:
							{
								position22, tokenIndex22 := position, tokenIndex
								if buffer[position] != rune('i') {
									goto l23
								}
								position++
								goto l22
							l23:
								position, tokenIndex = position22, tokenIndex22
								if buffer[position] != rune('I') {
									goto l0
								}
								position++
							}
						l22:
							{
								position24, tokenIndex24 := position, tokenIndex
								if buffer[position] != rune('s') {
									goto l25
								}
								position++
								goto l24
							l25:
								position, tokenIndex = position24, tokenIndex24
								if buffer[position] != rune('S') {
									goto l0
								}
								position++
							}
						l24:
							{
								position26, tokenIndex26 := position, tokenIndex
								if buffer[position] != rune('t') {
									goto l27
								}
								position++
								goto l26
							l27:
								position, tokenIndex = position26, tokenIndex26
								if buffer[position] != rune('T') {
									goto l0
								}
								position++
							}
						l26:
							{
								position28, tokenIndex28 := position, tokenIndex
								if buffer[position] != rune('e') {
									goto l29
								}
								position++
								goto l28
							l29:
								position, tokenIndex = position28, tokenIndex28
								if buffer[position] != rune('E') {
									goto l0
								}
								position++
							}
						l28:
							{
								position30, tokenIndex30 := position, tokenIndex
								if buffer[position] != rune('r') {
									goto l31
								}
								position++
								goto l30
							l31:
								position, tokenIndex = position30, tokenIndex30
								if buffer[position] != rune('R') {
									goto l0
								}
								position++
							}
						l30:
							break
						case 'V', 'v':
							{
								position32, tokenIndex32 := position, tokenIndex
								if buffer[position] != rune('v') {
									goto l33
								}
								position++
								goto l32
							l33:
								position, tokenIndex = position32, tokenIndex32
								if buffer[position] != rune('V') {
									goto l0
								}
								position++
							}
						l32:
							{
								position34, tokenIndex34 := position, tokenIndex
								if buffer[position] != rune('i') {
									goto l35
								}
								position++
								goto l34
							l35:
								position, tokenIndex = position34, tokenIndex34
								if buffer[position] != rune('I') {
									goto l0
								}
								position++
							}
						l34:
							{
								position36, tokenIndex36 := position, tokenIndex
								if buffer[position] != rune('r') {
									goto l37
								}
								position++
								goto l36
							l37:
								position, tokenIndex = position36, tokenIndex36
								if buffer[position] != rune('R') {
									goto l0
								}
								position++
							}
						l36:
							{
								position38, tokenIndex38 := position, tokenIndex
								if buffer[position] != rune('t') {
									goto l39
								}
								position++
								goto l38
							l39:
								position, tokenIndex = position38, tokenIndex38
								if buffer[position] != rune('T') {
									goto l0
								}
								position++
							}
						l38:
							break
						case 'P', 'p':
							{
								position40, tokenIndex40 := position, tokenIndex
								if buffer[position] != rune('p') {
									goto l41
								}
								position++
								goto l40
							l41:
								position, tokenIndex = position40, tokenIndex40
								if buffer[position] != rune('P') {
									goto l0
								}
								position++
							}
						l40:
							{
								position42, tokenIndex42 := position, tokenIndex
								if buffer[position] != rune('u') {
									goto l43
								}
								position++
								goto l42
							l43:
								position, tokenIndex = position42, tokenIndex42
								if buffer[position] != rune('U') {
									goto l0
								}
								position++
							}
						l42:
							{
								position44, tokenIndex44 := position, tokenIndex
								if buffer[position] != rune('r') {
									goto l45
								}
								position++
								goto l44
							l45:
								position, tokenIndex = position44, tokenIndex44
								if buffer[position] != rune('R') {
									goto l0
								}
								position++
							}
						l44:
							{
								position46, tokenIndex46 := position, tokenIndex
								if buffer[position] != rune('e') {
									goto l47
								}
								position++
								goto l46
							l47:
								position, tokenIndex = position46, tokenIndex46
								if buffer[position] != rune('E') {
									goto l0
								}
								position++
							}
						l46:
							break
						case 'B', 'b':
							{
								position48, tokenIndex48 := position, tokenIndex
								if buffer[position] != rune('b') {
									goto l49
								}
								position++
								goto l48
							l49:
								position, tokenIndex = position48, tokenIndex48
								if buffer[position] != rune('B') {
									goto l0
								}
								position++
							}
						l48:
							{
								position50, tokenIndex50 := position, tokenIndex
								if buffer[position] != rune('i') {
									goto l51
								}
								position++
								goto l50
							l51:
								position, tokenIndex = position50, tokenIndex50
								if buffer[position] != rune('I') {
									goto l0
								}
								position++
							}
						l50:
							{
								position52, tokenIndex52 := position, tokenIndex
								if buffer[position] != rune('t') {
									goto l53
								}
								position++
								goto l52
							l53:
								position, tokenIndex = position52, tokenIndex52
								if buffer[position] != rune('T') {
									goto l0
								}
								position++
							}
						l52:
							{
								position54, tokenIndex54 := position, tokenIndex
								if buffer[position] != rune('f') {
									goto l55
								}
								position++
								goto l54
							l55:
								position, tokenIndex = position54, tokenIndex54
								if buffer[position] != rune('F') {
									goto l0
								}
								position++
							}
						l54:
							{
								position56, tokenIndex56 := position, tokenIndex
								if buffer[position] != rune('i') {
									goto l57
								}
								position++
								goto l56
							l57:
								position, tokenIndex = position56, tokenIndex56
								if buffer[position] != rune('I') {
									goto l0
								}
								position++
							}
						l56:
							{
								position58, tokenIndex58 := position, tokenIndex
								if buffer[position] != rune('e') {
									goto l59
								}
								position++
								goto l58
							l59:
								position, tokenIndex = position58, tokenIndex58
								if buffer[position] != rune('E') {
									goto l0
								}
								position++
							}
						l58:
							{
								position60, tokenIndex60 := position, tokenIndex
								if buffer[position] != rune('l') {
									goto l61
								}
								position++
								goto l60
							l61:
								position, tokenIndex = position60, tokenIndex60
								if buffer[position] != rune('L') {
									goto l0
								}
								position++
							}
						l60:
							{
								position62, tokenIndex62 := position, tokenIndex
								if buffer[position] != rune('d') {
									goto l63
								}
								position++
								goto l62
							l63:
								position, tokenIndex = position62, tokenIndex62
								if buffer[position] != rune('D') {
									goto l0
								}
								position++
							}
						l62:
							break
						default:
							{
								position64, tokenIndex64 := position, tokenIndex
								if buffer[position] != rune('m') {
									goto l65
								}
								position++
								goto l64
							l65:
								position, tokenIndex = position64, tokenIndex64
								if buffer[position] != rune('M') {
									goto l0
								}
								position++
							}
						l64:
							{
								position66, tokenIndex66 := position, tokenIndex
								if buffer[position] != rune('u') {
									goto l67
								}
								position++
								goto l66
							l67:
								position, tokenIndex = position66, tokenIndex66
								if buffer[position] != rune('U') {
									goto l0
								}
								position++
							}
						l66:
							{
								position68, tokenIndex68 := position, tokenIndex
								if buffer[position] != rune('t') {
									goto l69
								}
								position++
								goto l68
							l69:
								position, tokenIndex = position68, tokenIndex68
								if buffer[position] != rune('T') {
									goto l0
								}
								position++
							}
						l68:
							{
								position70, tokenIndex70 := position, tokenIndex
								if buffer[position] != rune('a') {
									goto l71
								}
								position++
								goto l70
							l71:
								position, tokenIndex = position70, tokenIndex70
								if buffer[position] != rune('A') {
									goto l0
								}
								position++
							}
						l70:
							{
								position72, tokenIndex72 := position, tokenIndex
								if buffer[position] != rune('b') {
									goto l73
								}
								position++
								goto l72
							l73:
								position, tokenIndex = position72, tokenIndex72
								if buffer[position] != rune('B') {
									goto l0
								}
								position++
							}
						l72:
							{
								position74, tokenIndex74 := position, tokenIndex
								if buffer[position] != rune('l') {
									goto l75
								}
								position++
								goto l74
							l75:
								position, tokenIndex = position74, tokenIndex74
								if buffer[position] != rune('L') {
									goto l0
								}
								position++
							}
						l74:
							{
								position76, tokenIndex76 := position, tokenIndex
								if buffer[position] != rune('e') {
									goto l77
								}
								position++
								goto l76
							l77:
								position, tokenIndex = position76, tokenIndex76
								if buffer[position] != rune('E') {
									goto l0
								}
								position++
							}
						l76:
							break
						}
					}

					add(rulePegText, position14)
				}
				if !_rules[rulews]() {
					goto l0
				}
			l12:
				{
					position13, tokenIndex13 := position, tokenIndex
					{
						position78 := position
						{
							switch buffer[position] {
							case 'R', 'r':
								{
									position80, tokenIndex80 := position, tokenIndex
									if buffer[position] != rune('r') {
										goto l81
									}
									position++
									goto l80
								l81:
									position, tokenIndex = position80, tokenIndex80
									if buffer[position] != rune('R') {
										goto l13
									}
									position++
								}
							l80:
								{
									position82, tokenIndex82 := position, tokenIndex
									if buffer[position] != rune('e') {
										goto l83
									}
									position++
									goto l82
								l83:
									position, tokenIndex = position82, tokenIndex82
									if buffer[position] != rune('E') {
										goto l13
									}
									position++
								}
							l82:
								{
									position84, tokenIndex84 := position, tokenIndex
									if buffer[position] != rune('g') {
										goto l85
									}
									position++
									goto l84
								l85:
									position, tokenIndex = position84, tokenIndex84
									if buffer[position] != rune('G') {
										goto l13
									}
									position++
								}
							l84:
								{
									position86, tokenIndex86 := position, tokenIndex
									if buffer[position] != rune('i') {
										goto l87
									}
									position++
									goto l86
								l87:
									position, tokenIndex = position86, tokenIndex86
									if buffer[position] != rune('I') {
										goto l13
									}
									position++
								}
							l86:
								{
									position88, tokenIndex88 := position, tokenIndex
									if buffer[position] != rune('s') {
										goto l89
									}
									position++
									goto l88
								l89:
									position, tokenIndex = position88, tokenIndex88
									if buffer[position] != rune('S') {
										goto l13
									}
									position++
								}
							l88:
								{
									position90, tokenIndex90 := position, tokenIndex
									if buffer[position] != rune('t') {
										goto l91
									}
									position++
									goto l90
								l91:
									position, tokenIndex = position90, tokenIndex90
									if buffer[position] != rune('T') {
										goto l13
									}
									position++
								}
							l90:
								{
									position92, tokenIndex92 := position, tokenIndex
									if buffer[position] != rune('e') {
										goto l93
									}
									position++
									goto l92
								l93:
									position, tokenIndex = position92, tokenIndex92
									if buffer[position] != rune('E') {
										goto l13
									}
									position++
								}
							l92:
								{
									position94, tokenIndex94 := position, tokenIndex
									if buffer[position] != rune('r') {
										goto l95
									}
									position++
									goto l94
								l95:
									position, tokenIndex = position94, tokenIndex94
									if buffer[position] != rune('R') {
										goto l13
									}
									position++
								}
							l94:
								break
							case 'V', 'v':
								{
									position96, tokenIndex96 := position, tokenIndex
									if buffer[position] != rune('v') {
										goto l97
									}
									position++
									goto l96
								l97:
									position, tokenIndex = position96, tokenIndex96
									if buffer[position] != rune('V') {
										goto l13
									}
									position++
								}
							l96:
								{
									position98, tokenIndex98 := position, tokenIndex
									if buffer[position] != rune('i') {
										goto l99
									}
									position++
									goto l98
								l99:
									position, tokenIndex = position98, tokenIndex98
									if buffer[position] != rune('I') {
										goto l13
									}
									position++
								}
							l98:
								{
									position100, tokenIndex100 := position, tokenIndex
									if buffer[position] != rune('r') {
										goto l101
									}
									position++
									goto l100
								l101:
									position, tokenIndex = position100, tokenIndex100
									if buffer[position] != rune('R') {
										goto l13
									}
									position++
								}
							l100:
								{
									position102, tokenIndex102 := position, tokenIndex
									if buffer[position] != rune('t') {
										goto l103
									}
									position++
									goto l102
								l103:
									position, tokenIndex = position102, tokenIndex102
									if buffer[position] != rune('T') {
										goto l13
									}
									position++
								}
							l102:
								break
							case 'P', 'p':
								{
									position104, tokenIndex104 := position, tokenIndex
									if buffer[position] != rune('p') {
										goto l105
									}
									position++
									goto l104
								l105:
									position, tokenIndex = position104, tokenIndex104
									if buffer[position] != rune('P') {
										goto l13
									}
									position++
								}
							l104:
								{
									position106, tokenIndex106 := position, tokenIndex
									if buffer[position] != rune('u') {
										goto l107
									}
									position++
									goto l106
								l107:
									position, tokenIndex = position106, tokenIndex106
									if buffer[position] != rune('U') {
										goto l13
									}
									position++
								}
							l106:
								{
									position108, tokenIndex108 := position, tokenIndex
									if buffer[position] != rune('r') {
										goto l109
									}
									position++
									goto l108
								l109:
									position, tokenIndex = position108, tokenIndex108
									if buffer[position] != rune('R') {
										goto l13
									}
									position++
								}
							l108:
								{
									position110, tokenIndex110 := position, tokenIndex
									if buffer[position] != rune('e') {
										goto l111
									}
									position++
									goto l110
								l111:
									position, tokenIndex = position110, tokenIndex110
									if buffer[position] != rune('E') {
										goto l13
									}
									position++
								}
							l110:
								break
							case 'B', 'b':
								{
									position112, tokenIndex112 := position, tokenIndex
									if buffer[position] != rune('b') {
										goto l113
									}
									position++
									goto l112
								l113:
									position, tokenIndex = position112, tokenIndex112
									if buffer[position] != rune('B') {
										goto l13
									}
									position++
								}
							l112:
								{
									position114, tokenIndex114 := position, tokenIndex
									if buffer[position] != rune('i') {
										goto l115
									}
									position++
									goto l114
								l115:
									position, tokenIndex = position114, tokenIndex114
									if buffer[position] != rune('I') {
										goto l13
									}
									position++
								}
							l114:
								{
									position116, tokenIndex116 := position, tokenIndex
									if buffer[position] != rune('t') {
										goto l117
									}
									position++
									goto l116
								l117:
									position, tokenIndex = position116, tokenIndex116
									if buffer[position] != rune('T') {
										goto l13
									}
									position++
								}
							l116:
								{
									position118, tokenIndex118 := position, tokenIndex
									if buffer[position] != rune('f') {
										goto l119
									}
									position++
									goto l118
								l119:
									position, tokenIndex = position118, tokenIndex118
									if buffer[position] != rune('F') {
										goto l13
									}
									position++
								}
							l118:
								{
									position120, tokenIndex120 := position, tokenIndex
									if buffer[position] != rune('i') {
										goto l121
									}
									position++
									goto l120
								l121:
									position, tokenIndex = position120, tokenIndex120
									if buffer[position] != rune('I') {
										goto l13
									}
									position++
								}
							l120:
								{
									position122, tokenIndex122 := position, tokenIndex
									if buffer[position] != rune('e') {
										goto l123
									}
									position++
									goto l122
								l123:
									position, tokenIndex = position122, tokenIndex122
									if buffer[position] != rune('E') {
										goto l13
									}
									position++
								}
							l122:
								{
									position124, tokenIndex124 := position, tokenIndex
									if buffer[position] != rune('l') {
										goto l125
									}
									position++
									goto l124
								l125:
									position, tokenIndex = position124, tokenIndex124
									if buffer[position] != rune('L') {
										goto l13
									}
									position++
								}
							l124:
								{
									position126, tokenIndex126 := position, tokenIndex
									if buffer[position] != rune('d') {
										goto l127
									}
									position++
									goto l126
								l127:
									position, tokenIndex = position126, tokenIndex126
									if buffer[position] != rune('D') {
										goto l13
									}
									position++
								}
							l126:
								break
							default:
								{
									position128, tokenIndex128 := position, tokenIndex
									if buffer[position] != rune('m') {
										goto l129
									}
									position++
									goto l128
								l129:
									position, tokenIndex = position128, tokenIndex128
									if buffer[position] != rune('M') {
										goto l13
									}
									position++
								}
							l128:
								{
									position130, tokenIndex130 := position, tokenIndex
									if buffer[position] != rune('u') {
										goto l131
									}
									position++
									goto l130
								l131:
									position, tokenIndex = position130, tokenIndex130
									if buffer[position] != rune('U') {
										goto l13
									}
									position++
								}
							l130:
								{
									position132, tokenIndex132 := position, tokenIndex
									if buffer[position] != rune('t') {
										goto l133
									}
									position++
									goto l132
								l133:
									position, tokenIndex = position132, tokenIndex132
									if buffer[position] != rune('T') {
										goto l13
									}
									position++
								}
							l132:
								{
									position134, tokenIndex134 := position, tokenIndex
									if buffer[position] != rune('a') {
										goto l135
									}
									position++
									goto l134
								l135:
									position, tokenIndex = position134, tokenIndex134
									if buffer[position] != rune('A') {
										goto l13
									}
									position++
								}
							l134:
								{
									position136, tokenIndex136 := position, tokenIndex
									if buffer[position] != rune('b') {
										goto l137
									}
									position++
									goto l136
								l137:
									position, tokenIndex = position136, tokenIndex136
									if buffer[position] != rune('B') {
										goto l13
									}
									position++
								}
							l136:
								{
									position138, tokenIndex138 := position, tokenIndex
									if buffer[position] != rune('l') {
										goto l139
									}
									position++
									goto l138
								l139:
									position, tokenIndex = position138, tokenIndex138
									if buffer[position] != rune('L') {
										goto l13
									}
									position++
								}
							l138:
								{
									position140, tokenIndex140 := position, tokenIndex
									if buffer[position] != rune('e') {
										goto l141
									}
									position++
									goto l140
								l141:
									position, tokenIndex = position140, tokenIndex140
									if buffer[position] != rune('E') {
										goto l13
									}
									position++
								}
							l140:
								break
							}
						}

						add(rulePegText, position78)
					}
					if !_rules[rulews]() {
						goto l13
					}
					goto l12
				l13:
					position, tokenIndex = position13, tokenIndex13
				}
				{
					add(ruleAction0, position)
				}
				add(ruleSpecValue, position1)
			}
			return true
		l0:
			position, tokenIndex = position0, tokenIndex0
			return false
		},
		nil,
		nil,
		/* 4 Action0 <- <{
			s := astproto.Spec( astproto.Spec_value[buffer[begin:end]] )
			getNode().AddSpec(s)
		}> */
		nil,
	}
	p.rules = _rules
}
