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
	ruleSignAttr
	rulePegText
	rulews
	ruleAction0
)

var rul3s = [...]string{
	"Unknown",
	"SignAttr",
	"PegText",
	"ws",
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

			getNode().AddSign(astproto.SignType(astproto.SignType_value[buffer[begin:end]]))

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
		/* 0 SignAttr <- <(<(('s' / 'S') ('i' / 'I') ('g' / 'G') ('n' / 'N'))> ws ':' ws <((('s' / 'S') ('i' / 'I') ('g' / 'G') ('n' / 'N') ('e' / 'E') ('d' / 'D')) / (('u' / 'U') ('n' / 'N') ('s' / 'S') ('i' / 'I') ('g' / 'G') ('n' / 'N') ('e' / 'E') ('d' / 'D')))> Action0)> */
		func() bool {
			position0, tokenIndex0 := position, tokenIndex
			{
				position1 := position
				{
					position2 := position
					{
						position3, tokenIndex3 := position, tokenIndex
						if buffer[position] != rune('s') {
							goto l4
						}
						position++
						goto l3
					l4:
						position, tokenIndex = position3, tokenIndex3
						if buffer[position] != rune('S') {
							goto l0
						}
						position++
					}
				l3:
					{
						position5, tokenIndex5 := position, tokenIndex
						if buffer[position] != rune('i') {
							goto l6
						}
						position++
						goto l5
					l6:
						position, tokenIndex = position5, tokenIndex5
						if buffer[position] != rune('I') {
							goto l0
						}
						position++
					}
				l5:
					{
						position7, tokenIndex7 := position, tokenIndex
						if buffer[position] != rune('g') {
							goto l8
						}
						position++
						goto l7
					l8:
						position, tokenIndex = position7, tokenIndex7
						if buffer[position] != rune('G') {
							goto l0
						}
						position++
					}
				l7:
					{
						position9, tokenIndex9 := position, tokenIndex
						if buffer[position] != rune('n') {
							goto l10
						}
						position++
						goto l9
					l10:
						position, tokenIndex = position9, tokenIndex9
						if buffer[position] != rune('N') {
							goto l0
						}
						position++
					}
				l9:
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
					position11 := position
					{
						position12, tokenIndex12 := position, tokenIndex
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
								goto l13
							}
							position++
						}
					l14:
						{
							position16, tokenIndex16 := position, tokenIndex
							if buffer[position] != rune('i') {
								goto l17
							}
							position++
							goto l16
						l17:
							position, tokenIndex = position16, tokenIndex16
							if buffer[position] != rune('I') {
								goto l13
							}
							position++
						}
					l16:
						{
							position18, tokenIndex18 := position, tokenIndex
							if buffer[position] != rune('g') {
								goto l19
							}
							position++
							goto l18
						l19:
							position, tokenIndex = position18, tokenIndex18
							if buffer[position] != rune('G') {
								goto l13
							}
							position++
						}
					l18:
						{
							position20, tokenIndex20 := position, tokenIndex
							if buffer[position] != rune('n') {
								goto l21
							}
							position++
							goto l20
						l21:
							position, tokenIndex = position20, tokenIndex20
							if buffer[position] != rune('N') {
								goto l13
							}
							position++
						}
					l20:
						{
							position22, tokenIndex22 := position, tokenIndex
							if buffer[position] != rune('e') {
								goto l23
							}
							position++
							goto l22
						l23:
							position, tokenIndex = position22, tokenIndex22
							if buffer[position] != rune('E') {
								goto l13
							}
							position++
						}
					l22:
						{
							position24, tokenIndex24 := position, tokenIndex
							if buffer[position] != rune('d') {
								goto l25
							}
							position++
							goto l24
						l25:
							position, tokenIndex = position24, tokenIndex24
							if buffer[position] != rune('D') {
								goto l13
							}
							position++
						}
					l24:
						goto l12
					l13:
						position, tokenIndex = position12, tokenIndex12
						{
							position26, tokenIndex26 := position, tokenIndex
							if buffer[position] != rune('u') {
								goto l27
							}
							position++
							goto l26
						l27:
							position, tokenIndex = position26, tokenIndex26
							if buffer[position] != rune('U') {
								goto l0
							}
							position++
						}
					l26:
						{
							position28, tokenIndex28 := position, tokenIndex
							if buffer[position] != rune('n') {
								goto l29
							}
							position++
							goto l28
						l29:
							position, tokenIndex = position28, tokenIndex28
							if buffer[position] != rune('N') {
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
							if buffer[position] != rune('i') {
								goto l33
							}
							position++
							goto l32
						l33:
							position, tokenIndex = position32, tokenIndex32
							if buffer[position] != rune('I') {
								goto l0
							}
							position++
						}
					l32:
						{
							position34, tokenIndex34 := position, tokenIndex
							if buffer[position] != rune('g') {
								goto l35
							}
							position++
							goto l34
						l35:
							position, tokenIndex = position34, tokenIndex34
							if buffer[position] != rune('G') {
								goto l0
							}
							position++
						}
					l34:
						{
							position36, tokenIndex36 := position, tokenIndex
							if buffer[position] != rune('n') {
								goto l37
							}
							position++
							goto l36
						l37:
							position, tokenIndex = position36, tokenIndex36
							if buffer[position] != rune('N') {
								goto l0
							}
							position++
						}
					l36:
						{
							position38, tokenIndex38 := position, tokenIndex
							if buffer[position] != rune('e') {
								goto l39
							}
							position++
							goto l38
						l39:
							position, tokenIndex = position38, tokenIndex38
							if buffer[position] != rune('E') {
								goto l0
							}
							position++
						}
					l38:
						{
							position40, tokenIndex40 := position, tokenIndex
							if buffer[position] != rune('d') {
								goto l41
							}
							position++
							goto l40
						l41:
							position, tokenIndex = position40, tokenIndex40
							if buffer[position] != rune('D') {
								goto l0
							}
							position++
						}
					l40:
					}
				l12:
					add(rulePegText, position11)
				}
				{
					add(ruleAction0, position)
				}
				add(ruleSignAttr, position1)
			}
			return true
		l0:
			position, tokenIndex = position0, tokenIndex0
			return false
		},
		nil,
		nil,
		/* 4 Action0 <- <{
			getNode().AddSign(astproto.SignType(astproto.SignType_value[buffer[begin:end]]))
		}> */
		nil,
	}
	p.rules = _rules
}
