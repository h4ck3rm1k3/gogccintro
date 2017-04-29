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
	ruleIntAttr
	rulePegText
	rulews
	ruleInteger
	ruleAction0
)

var rul3s = [...]string{
	"Unknown",
	"IntAttr",
	"PegText",
	"ws",
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
	rules  [6]func() bool
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

			// base
			getNode().AddIntField(FieldName, Atoi(buffer[begin:end]))

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
		/* 0 IntAttr <- <(<((&('P' | 'p') (('p' / 'P') ('r' / 'R') ('e' / 'E') ('c' / 'C'))) | (&('U' | 'u') (('u' / 'U') ('s' / 'S') ('e' / 'E') ('d' / 'D'))) | (&('L' | 'l') (('l' / 'L') ('o' / 'O') ('w' / 'W'))) | (&('B' | 'b') (('b' / 'B') ('a' / 'A') ('s' / 'S') ('e' / 'E') ('s' / 'S'))))> ws ':' ws <Integer> ws Action0)> */
		func() bool {
			position0, tokenIndex0 := position, tokenIndex
			{
				position1 := position
				{
					position2 := position
					{
						switch buffer[position] {
						case 'P', 'p':
							{
								position4, tokenIndex4 := position, tokenIndex
								if buffer[position] != rune('p') {
									goto l5
								}
								position++
								goto l4
							l5:
								position, tokenIndex = position4, tokenIndex4
								if buffer[position] != rune('P') {
									goto l0
								}
								position++
							}
						l4:
							{
								position6, tokenIndex6 := position, tokenIndex
								if buffer[position] != rune('r') {
									goto l7
								}
								position++
								goto l6
							l7:
								position, tokenIndex = position6, tokenIndex6
								if buffer[position] != rune('R') {
									goto l0
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
									goto l0
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
									goto l0
								}
								position++
							}
						l10:
							break
						case 'U', 'u':
							{
								position12, tokenIndex12 := position, tokenIndex
								if buffer[position] != rune('u') {
									goto l13
								}
								position++
								goto l12
							l13:
								position, tokenIndex = position12, tokenIndex12
								if buffer[position] != rune('U') {
									goto l0
								}
								position++
							}
						l12:
							{
								position14, tokenIndex14 := position, tokenIndex
								if buffer[position] != rune('s') {
									goto l15
								}
								position++
								goto l14
							l15:
								position, tokenIndex = position14, tokenIndex14
								if buffer[position] != rune('S') {
									goto l0
								}
								position++
							}
						l14:
							{
								position16, tokenIndex16 := position, tokenIndex
								if buffer[position] != rune('e') {
									goto l17
								}
								position++
								goto l16
							l17:
								position, tokenIndex = position16, tokenIndex16
								if buffer[position] != rune('E') {
									goto l0
								}
								position++
							}
						l16:
							{
								position18, tokenIndex18 := position, tokenIndex
								if buffer[position] != rune('d') {
									goto l19
								}
								position++
								goto l18
							l19:
								position, tokenIndex = position18, tokenIndex18
								if buffer[position] != rune('D') {
									goto l0
								}
								position++
							}
						l18:
							break
						case 'L', 'l':
							{
								position20, tokenIndex20 := position, tokenIndex
								if buffer[position] != rune('l') {
									goto l21
								}
								position++
								goto l20
							l21:
								position, tokenIndex = position20, tokenIndex20
								if buffer[position] != rune('L') {
									goto l0
								}
								position++
							}
						l20:
							{
								position22, tokenIndex22 := position, tokenIndex
								if buffer[position] != rune('o') {
									goto l23
								}
								position++
								goto l22
							l23:
								position, tokenIndex = position22, tokenIndex22
								if buffer[position] != rune('O') {
									goto l0
								}
								position++
							}
						l22:
							{
								position24, tokenIndex24 := position, tokenIndex
								if buffer[position] != rune('w') {
									goto l25
								}
								position++
								goto l24
							l25:
								position, tokenIndex = position24, tokenIndex24
								if buffer[position] != rune('W') {
									goto l0
								}
								position++
							}
						l24:
							break
						default:
							{
								position26, tokenIndex26 := position, tokenIndex
								if buffer[position] != rune('b') {
									goto l27
								}
								position++
								goto l26
							l27:
								position, tokenIndex = position26, tokenIndex26
								if buffer[position] != rune('B') {
									goto l0
								}
								position++
							}
						l26:
							{
								position28, tokenIndex28 := position, tokenIndex
								if buffer[position] != rune('a') {
									goto l29
								}
								position++
								goto l28
							l29:
								position, tokenIndex = position28, tokenIndex28
								if buffer[position] != rune('A') {
									goto l0
								}
								position++
							}
						l28:
							{
								position30, tokenIndex30 := position, tokenIndex
								if buffer[position] != rune('s') {
									goto l31
								}
								position++
								goto l30
							l31:
								position, tokenIndex = position30, tokenIndex30
								if buffer[position] != rune('S') {
									goto l0
								}
								position++
							}
						l30:
							{
								position32, tokenIndex32 := position, tokenIndex
								if buffer[position] != rune('e') {
									goto l33
								}
								position++
								goto l32
							l33:
								position, tokenIndex = position32, tokenIndex32
								if buffer[position] != rune('E') {
									goto l0
								}
								position++
							}
						l32:
							{
								position34, tokenIndex34 := position, tokenIndex
								if buffer[position] != rune('s') {
									goto l35
								}
								position++
								goto l34
							l35:
								position, tokenIndex = position34, tokenIndex34
								if buffer[position] != rune('S') {
									goto l0
								}
								position++
							}
						l34:
							break
						}
					}

					add(rulePegText, position2)
				}
				if !_rules[rulews]() {
					goto l0
				}
				if buffer[position] != rune(':') {
					goto l0
				}
				position++
				if !_rules[rulews]() {
					goto l0
				}
				{
					position36 := position
					{
						position37 := position
						add(ruleInteger, position37)
					}
					add(rulePegText, position36)
				}
				if !_rules[rulews]() {
					goto l0
				}
				{
					add(ruleAction0, position)
				}
				add(ruleIntAttr, position1)
			}
			return true
		l0:
			position, tokenIndex = position0, tokenIndex0
			return false
		},
		nil,
		nil,
		nil,
		/* 5 Action0 <- <{
			// base
			getNode().AddIntField(FieldName,Atoi(buffer[begin:end]))
		}> */
		nil,
	}
	p.rules = _rules
}
