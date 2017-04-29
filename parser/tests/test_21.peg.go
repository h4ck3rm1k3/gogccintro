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
	ruleNoteAttr
	rulePegText
	rulews
	ruleAction0
)

var rul3s = [...]string{
	"Unknown",
	"NoteAttr",
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

			getNode().AddNote(buffer[begin:end])

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
		/* 0 NoteAttr <- <(<(('n' / 'N') ('o' / 'O') ('t' / 'T') ('e' / 'E'))> ws ':' ws <((('p' / 'P') ('t' / 'T') ('r' / 'R') ('m' / 'M') ('e' / 'E') ('m' / 'M')) / ((&('D' | 'd') (('d' / 'D') ('e' / 'E') ('s' / 'S') ('t' / 'T') ('r' / 'R') ('u' / 'U') ('c' / 'C') ('t' / 'T') ('o' / 'O') ('r' / 'R'))) | (&('P' | 'p') (('p' / 'P') ('s' / 'S') ('e' / 'E') ('u' / 'U') ('d' / 'D') ('o' / 'O') ' ' ('t' / 'T') ('m' / 'M') ('p' / 'P') ('l' / 'L'))) | (&('O' | 'o') (('o' / 'O') ('p' / 'P') ('e' / 'E') ('r' / 'R') ('a' / 'A') ('t' / 'T') ('o' / 'O') ('r' / 'R') ws ((('a' / 'A') ('n' / 'N') ('d' / 'D') ('a' / 'A') ('s' / 'S') ('s' / 'S') ('i' / 'I') ('g' / 'G') ('n' / 'N')) / (('a' / 'A') ('n' / 'N') ('d' / 'D')) / (('d' / 'D') ('e' / 'E') ('l' / 'L') ('e' / 'E') ('t' / 'T') ('e' / 'E')) / (('d' / 'D') ('e' / 'E') ('r' / 'R') ('e' / 'E') ('f' / 'F')) / (('d' / 'D') ('i' / 'I') ('v' / 'V') ('a' / 'A') ('s' / 'S') ('s' / 'S') ('i' / 'I') ('g' / 'G') ('n' / 'N')) / (('g' / 'G') ('e' / 'E')) / (('g' / 'G') ('e' / 'E')) / (('l' / 'L') ('a' / 'A') ('n' / 'N') ('d' / 'D')) / (('l' / 'L') ('e' / 'E')) / (('l' / 'L') ('n' / 'N') ('o' / 'O') ('t' / 'T')) / (('l' / 'L') ('o' / 'O') ('r' / 'R')) / (('l' / 'L') ('s' / 'S') ('h' / 'H') ('i' / 'I') ('f' / 'F') ('t' / 'T') ('a' / 'A') ('s' / 'S') ('s' / 'S') ('i' / 'I') ('g' / 'G') ('n' / 'N')) / (('l' / 'L') ('s' / 'S') ('h' / 'H') ('i' / 'I') ('f' / 'F') ('t' / 'T')) / (('m' / 'M') ('i' / 'I') ('n' / 'N') ('u' / 'U') ('s' / 'S') ('a' / 'A') ('s' / 'S') ('s' / 'S') ('i' / 'I') ('g' / 'G') ('n' / 'N')) / (('m' / 'M') ('i' / 'I') ('n' / 'N') ('u' / 'U') ('s' / 'S')) / (('m' / 'M') ('u' / 'U') ('l' / 'L') ('t' / 'T') ('a' / 'A') ('s' / 'S') ('s' / 'S') ('i' / 'I') ('g' / 'G') ('n' / 'N')) / (('n' / 'N') ('e' / 'E') ('g' / 'G')) / (('n' / 'N') ('e' / 'E') ('w' / 'W')) / (('n' / 'N') ('e' / 'E')) / (('o' / 'O') ('r' / 'R') ('a' / 'A') ('s' / 'S') ('s' / 'S') ('i' / 'I') ('g' / 'G') ('n' / 'N')) / (('p' / 'P') ('l' / 'L') ('u' / 'U') ('s' / 'S') ('a' / 'A') ('s' / 'S') ('s' / 'S') ('i' / 'I') ('g' / 'G') ('n' / 'N')) / (('p' / 'P') ('l' / 'L') ('u' / 'U') ('s' / 'S')) / (('p' / 'P') ('o' / 'O') ('s' / 'S') ('t' / 'T') ('d' / 'D') ('e' / 'E') ('c' / 'C')) / (('p' / 'P') ('o' / 'O') ('s' / 'S') ('t' / 'T') ('i' / 'I') ('n' / 'N') ('c' / 'C')) / (('p' / 'P') ('o' / 'O') ('s' / 'S')) / (('p' / 'P') ('r' / 'R') ('e' / 'E') ('d' / 'D') ('e' / 'E') ('c' / 'C')) / (('r' / 'R') ('e' / 'E') ('f' / 'F')) / (('r' / 'R') ('s' / 'S') ('h' / 'H') ('i' / 'I') ('f' / 'F') ('t' / 'T') ('a' / 'A') ('s' / 'S') ('s' / 'S') ('i' / 'I') ('g' / 'G') ('n' / 'N')) / (('v' / 'V') ('e' / 'E') ('c' / 'C') ('d' / 'D') ('e' / 'E') ('l' / 'L') ('e' / 'E') ('t' / 'T') ('e' / 'E')) / (('x' / 'X') ('o' / 'O') ('r' / 'R') ('a' / 'A') ('s' / 'S') ('s' / 'S') ('i' / 'I') ('g' / 'G') ('n' / 'N')) / ((&('X' | 'x') (('x' / 'X') ('o' / 'O') ('r' / 'R'))) | (&('V' | 'v') (('v' / 'V') ('e' / 'E') ('c' / 'C') ('n' / 'N') ('e' / 'E') ('w' / 'W'))) | (&('S' | 's') (('s' / 'S') ('u' / 'U') ('b' / 'B') ('s' / 'S'))) | (&('R' | 'r') (('r' / 'R') ('s' / 'S') ('h' / 'H') ('i' / 'I') ('f' / 'F') ('t' / 'T'))) | (&('P' | 'p') (('p' / 'P') ('r' / 'R') ('e' / 'E') ('i' / 'I') ('n' / 'N') ('c' / 'C'))) | (&('O' | 'o') (('o' / 'O') ('r' / 'R'))) | (&('N' | 'n') (('n' / 'N') ('o' / 'O') ('t' / 'T'))) | (&('M' | 'm') (('m' / 'M') ('u' / 'U') ('l' / 'L') ('t' / 'T'))) | (&('L' | 'l') (('l' / 'L') ('t' / 'T'))) | (&('G' | 'g') (('g' / 'G') ('t' / 'T'))) | (&('E' | 'e') (('e' / 'E') ('q' / 'Q'))) | (&('D' | 'd') (('d' / 'D') ('i' / 'I') ('v' / 'V'))) | (&('C' | 'c') (('c' / 'C') ('a' / 'A') ('l' / 'L') ('l' / 'L'))) | (&('A' | 'a') (('a' / 'A') ('s' / 'S') ('s' / 'S') ('i' / 'I') ('g' / 'G') ('n' / 'N'))) | (&() )))?)) | (&('M' | 'm') (('m' / 'M') ('e' / 'E') ('m' / 'M') ('b' / 'B') ('e' / 'E') ('r' / 'R'))) | (&('A' | 'a') (('a' / 'A') ('r' / 'R') ('t' / 'T') ('i' / 'I') ('f' / 'F') ('i' / 'I') ('c' / 'C') ('i' / 'I') ('a' / 'A') ('l' / 'L')))))> Action0)> */
		func() bool {
			position0, tokenIndex0 := position, tokenIndex
			{
				position1 := position
				{
					position2 := position
					{
						position3, tokenIndex3 := position, tokenIndex
						if buffer[position] != rune('n') {
							goto l4
						}
						position++
						goto l3
					l4:
						position, tokenIndex = position3, tokenIndex3
						if buffer[position] != rune('N') {
							goto l0
						}
						position++
					}
				l3:
					{
						position5, tokenIndex5 := position, tokenIndex
						if buffer[position] != rune('o') {
							goto l6
						}
						position++
						goto l5
					l6:
						position, tokenIndex = position5, tokenIndex5
						if buffer[position] != rune('O') {
							goto l0
						}
						position++
					}
				l5:
					{
						position7, tokenIndex7 := position, tokenIndex
						if buffer[position] != rune('t') {
							goto l8
						}
						position++
						goto l7
					l8:
						position, tokenIndex = position7, tokenIndex7
						if buffer[position] != rune('T') {
							goto l0
						}
						position++
					}
				l7:
					{
						position9, tokenIndex9 := position, tokenIndex
						if buffer[position] != rune('e') {
							goto l10
						}
						position++
						goto l9
					l10:
						position, tokenIndex = position9, tokenIndex9
						if buffer[position] != rune('E') {
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
							if buffer[position] != rune('p') {
								goto l15
							}
							position++
							goto l14
						l15:
							position, tokenIndex = position14, tokenIndex14
							if buffer[position] != rune('P') {
								goto l13
							}
							position++
						}
					l14:
						{
							position16, tokenIndex16 := position, tokenIndex
							if buffer[position] != rune('t') {
								goto l17
							}
							position++
							goto l16
						l17:
							position, tokenIndex = position16, tokenIndex16
							if buffer[position] != rune('T') {
								goto l13
							}
							position++
						}
					l16:
						{
							position18, tokenIndex18 := position, tokenIndex
							if buffer[position] != rune('r') {
								goto l19
							}
							position++
							goto l18
						l19:
							position, tokenIndex = position18, tokenIndex18
							if buffer[position] != rune('R') {
								goto l13
							}
							position++
						}
					l18:
						{
							position20, tokenIndex20 := position, tokenIndex
							if buffer[position] != rune('m') {
								goto l21
							}
							position++
							goto l20
						l21:
							position, tokenIndex = position20, tokenIndex20
							if buffer[position] != rune('M') {
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
							if buffer[position] != rune('m') {
								goto l25
							}
							position++
							goto l24
						l25:
							position, tokenIndex = position24, tokenIndex24
							if buffer[position] != rune('M') {
								goto l13
							}
							position++
						}
					l24:
						goto l12
					l13:
						position, tokenIndex = position12, tokenIndex12
						{
							switch buffer[position] {
							case 'D', 'd':
								{
									position27, tokenIndex27 := position, tokenIndex
									if buffer[position] != rune('d') {
										goto l28
									}
									position++
									goto l27
								l28:
									position, tokenIndex = position27, tokenIndex27
									if buffer[position] != rune('D') {
										goto l0
									}
									position++
								}
							l27:
								{
									position29, tokenIndex29 := position, tokenIndex
									if buffer[position] != rune('e') {
										goto l30
									}
									position++
									goto l29
								l30:
									position, tokenIndex = position29, tokenIndex29
									if buffer[position] != rune('E') {
										goto l0
									}
									position++
								}
							l29:
								{
									position31, tokenIndex31 := position, tokenIndex
									if buffer[position] != rune('s') {
										goto l32
									}
									position++
									goto l31
								l32:
									position, tokenIndex = position31, tokenIndex31
									if buffer[position] != rune('S') {
										goto l0
									}
									position++
								}
							l31:
								{
									position33, tokenIndex33 := position, tokenIndex
									if buffer[position] != rune('t') {
										goto l34
									}
									position++
									goto l33
								l34:
									position, tokenIndex = position33, tokenIndex33
									if buffer[position] != rune('T') {
										goto l0
									}
									position++
								}
							l33:
								{
									position35, tokenIndex35 := position, tokenIndex
									if buffer[position] != rune('r') {
										goto l36
									}
									position++
									goto l35
								l36:
									position, tokenIndex = position35, tokenIndex35
									if buffer[position] != rune('R') {
										goto l0
									}
									position++
								}
							l35:
								{
									position37, tokenIndex37 := position, tokenIndex
									if buffer[position] != rune('u') {
										goto l38
									}
									position++
									goto l37
								l38:
									position, tokenIndex = position37, tokenIndex37
									if buffer[position] != rune('U') {
										goto l0
									}
									position++
								}
							l37:
								{
									position39, tokenIndex39 := position, tokenIndex
									if buffer[position] != rune('c') {
										goto l40
									}
									position++
									goto l39
								l40:
									position, tokenIndex = position39, tokenIndex39
									if buffer[position] != rune('C') {
										goto l0
									}
									position++
								}
							l39:
								{
									position41, tokenIndex41 := position, tokenIndex
									if buffer[position] != rune('t') {
										goto l42
									}
									position++
									goto l41
								l42:
									position, tokenIndex = position41, tokenIndex41
									if buffer[position] != rune('T') {
										goto l0
									}
									position++
								}
							l41:
								{
									position43, tokenIndex43 := position, tokenIndex
									if buffer[position] != rune('o') {
										goto l44
									}
									position++
									goto l43
								l44:
									position, tokenIndex = position43, tokenIndex43
									if buffer[position] != rune('O') {
										goto l0
									}
									position++
								}
							l43:
								{
									position45, tokenIndex45 := position, tokenIndex
									if buffer[position] != rune('r') {
										goto l46
									}
									position++
									goto l45
								l46:
									position, tokenIndex = position45, tokenIndex45
									if buffer[position] != rune('R') {
										goto l0
									}
									position++
								}
							l45:
								break
							case 'P', 'p':
								{
									position47, tokenIndex47 := position, tokenIndex
									if buffer[position] != rune('p') {
										goto l48
									}
									position++
									goto l47
								l48:
									position, tokenIndex = position47, tokenIndex47
									if buffer[position] != rune('P') {
										goto l0
									}
									position++
								}
							l47:
								{
									position49, tokenIndex49 := position, tokenIndex
									if buffer[position] != rune('s') {
										goto l50
									}
									position++
									goto l49
								l50:
									position, tokenIndex = position49, tokenIndex49
									if buffer[position] != rune('S') {
										goto l0
									}
									position++
								}
							l49:
								{
									position51, tokenIndex51 := position, tokenIndex
									if buffer[position] != rune('e') {
										goto l52
									}
									position++
									goto l51
								l52:
									position, tokenIndex = position51, tokenIndex51
									if buffer[position] != rune('E') {
										goto l0
									}
									position++
								}
							l51:
								{
									position53, tokenIndex53 := position, tokenIndex
									if buffer[position] != rune('u') {
										goto l54
									}
									position++
									goto l53
								l54:
									position, tokenIndex = position53, tokenIndex53
									if buffer[position] != rune('U') {
										goto l0
									}
									position++
								}
							l53:
								{
									position55, tokenIndex55 := position, tokenIndex
									if buffer[position] != rune('d') {
										goto l56
									}
									position++
									goto l55
								l56:
									position, tokenIndex = position55, tokenIndex55
									if buffer[position] != rune('D') {
										goto l0
									}
									position++
								}
							l55:
								{
									position57, tokenIndex57 := position, tokenIndex
									if buffer[position] != rune('o') {
										goto l58
									}
									position++
									goto l57
								l58:
									position, tokenIndex = position57, tokenIndex57
									if buffer[position] != rune('O') {
										goto l0
									}
									position++
								}
							l57:
								if buffer[position] != rune(' ') {
									goto l0
								}
								position++
								{
									position59, tokenIndex59 := position, tokenIndex
									if buffer[position] != rune('t') {
										goto l60
									}
									position++
									goto l59
								l60:
									position, tokenIndex = position59, tokenIndex59
									if buffer[position] != rune('T') {
										goto l0
									}
									position++
								}
							l59:
								{
									position61, tokenIndex61 := position, tokenIndex
									if buffer[position] != rune('m') {
										goto l62
									}
									position++
									goto l61
								l62:
									position, tokenIndex = position61, tokenIndex61
									if buffer[position] != rune('M') {
										goto l0
									}
									position++
								}
							l61:
								{
									position63, tokenIndex63 := position, tokenIndex
									if buffer[position] != rune('p') {
										goto l64
									}
									position++
									goto l63
								l64:
									position, tokenIndex = position63, tokenIndex63
									if buffer[position] != rune('P') {
										goto l0
									}
									position++
								}
							l63:
								{
									position65, tokenIndex65 := position, tokenIndex
									if buffer[position] != rune('l') {
										goto l66
									}
									position++
									goto l65
								l66:
									position, tokenIndex = position65, tokenIndex65
									if buffer[position] != rune('L') {
										goto l0
									}
									position++
								}
							l65:
								break
							case 'O', 'o':
								{
									position67, tokenIndex67 := position, tokenIndex
									if buffer[position] != rune('o') {
										goto l68
									}
									position++
									goto l67
								l68:
									position, tokenIndex = position67, tokenIndex67
									if buffer[position] != rune('O') {
										goto l0
									}
									position++
								}
							l67:
								{
									position69, tokenIndex69 := position, tokenIndex
									if buffer[position] != rune('p') {
										goto l70
									}
									position++
									goto l69
								l70:
									position, tokenIndex = position69, tokenIndex69
									if buffer[position] != rune('P') {
										goto l0
									}
									position++
								}
							l69:
								{
									position71, tokenIndex71 := position, tokenIndex
									if buffer[position] != rune('e') {
										goto l72
									}
									position++
									goto l71
								l72:
									position, tokenIndex = position71, tokenIndex71
									if buffer[position] != rune('E') {
										goto l0
									}
									position++
								}
							l71:
								{
									position73, tokenIndex73 := position, tokenIndex
									if buffer[position] != rune('r') {
										goto l74
									}
									position++
									goto l73
								l74:
									position, tokenIndex = position73, tokenIndex73
									if buffer[position] != rune('R') {
										goto l0
									}
									position++
								}
							l73:
								{
									position75, tokenIndex75 := position, tokenIndex
									if buffer[position] != rune('a') {
										goto l76
									}
									position++
									goto l75
								l76:
									position, tokenIndex = position75, tokenIndex75
									if buffer[position] != rune('A') {
										goto l0
									}
									position++
								}
							l75:
								{
									position77, tokenIndex77 := position, tokenIndex
									if buffer[position] != rune('t') {
										goto l78
									}
									position++
									goto l77
								l78:
									position, tokenIndex = position77, tokenIndex77
									if buffer[position] != rune('T') {
										goto l0
									}
									position++
								}
							l77:
								{
									position79, tokenIndex79 := position, tokenIndex
									if buffer[position] != rune('o') {
										goto l80
									}
									position++
									goto l79
								l80:
									position, tokenIndex = position79, tokenIndex79
									if buffer[position] != rune('O') {
										goto l0
									}
									position++
								}
							l79:
								{
									position81, tokenIndex81 := position, tokenIndex
									if buffer[position] != rune('r') {
										goto l82
									}
									position++
									goto l81
								l82:
									position, tokenIndex = position81, tokenIndex81
									if buffer[position] != rune('R') {
										goto l0
									}
									position++
								}
							l81:
								if !_rules[rulews]() {
									goto l0
								}
								{
									position83, tokenIndex83 := position, tokenIndex
									{
										position85, tokenIndex85 := position, tokenIndex
										{
											position87, tokenIndex87 := position, tokenIndex
											if buffer[position] != rune('a') {
												goto l88
											}
											position++
											goto l87
										l88:
											position, tokenIndex = position87, tokenIndex87
											if buffer[position] != rune('A') {
												goto l86
											}
											position++
										}
									l87:
										{
											position89, tokenIndex89 := position, tokenIndex
											if buffer[position] != rune('n') {
												goto l90
											}
											position++
											goto l89
										l90:
											position, tokenIndex = position89, tokenIndex89
											if buffer[position] != rune('N') {
												goto l86
											}
											position++
										}
									l89:
										{
											position91, tokenIndex91 := position, tokenIndex
											if buffer[position] != rune('d') {
												goto l92
											}
											position++
											goto l91
										l92:
											position, tokenIndex = position91, tokenIndex91
											if buffer[position] != rune('D') {
												goto l86
											}
											position++
										}
									l91:
										{
											position93, tokenIndex93 := position, tokenIndex
											if buffer[position] != rune('a') {
												goto l94
											}
											position++
											goto l93
										l94:
											position, tokenIndex = position93, tokenIndex93
											if buffer[position] != rune('A') {
												goto l86
											}
											position++
										}
									l93:
										{
											position95, tokenIndex95 := position, tokenIndex
											if buffer[position] != rune('s') {
												goto l96
											}
											position++
											goto l95
										l96:
											position, tokenIndex = position95, tokenIndex95
											if buffer[position] != rune('S') {
												goto l86
											}
											position++
										}
									l95:
										{
											position97, tokenIndex97 := position, tokenIndex
											if buffer[position] != rune('s') {
												goto l98
											}
											position++
											goto l97
										l98:
											position, tokenIndex = position97, tokenIndex97
											if buffer[position] != rune('S') {
												goto l86
											}
											position++
										}
									l97:
										{
											position99, tokenIndex99 := position, tokenIndex
											if buffer[position] != rune('i') {
												goto l100
											}
											position++
											goto l99
										l100:
											position, tokenIndex = position99, tokenIndex99
											if buffer[position] != rune('I') {
												goto l86
											}
											position++
										}
									l99:
										{
											position101, tokenIndex101 := position, tokenIndex
											if buffer[position] != rune('g') {
												goto l102
											}
											position++
											goto l101
										l102:
											position, tokenIndex = position101, tokenIndex101
											if buffer[position] != rune('G') {
												goto l86
											}
											position++
										}
									l101:
										{
											position103, tokenIndex103 := position, tokenIndex
											if buffer[position] != rune('n') {
												goto l104
											}
											position++
											goto l103
										l104:
											position, tokenIndex = position103, tokenIndex103
											if buffer[position] != rune('N') {
												goto l86
											}
											position++
										}
									l103:
										goto l85
									l86:
										position, tokenIndex = position85, tokenIndex85
										{
											position106, tokenIndex106 := position, tokenIndex
											if buffer[position] != rune('a') {
												goto l107
											}
											position++
											goto l106
										l107:
											position, tokenIndex = position106, tokenIndex106
											if buffer[position] != rune('A') {
												goto l105
											}
											position++
										}
									l106:
										{
											position108, tokenIndex108 := position, tokenIndex
											if buffer[position] != rune('n') {
												goto l109
											}
											position++
											goto l108
										l109:
											position, tokenIndex = position108, tokenIndex108
											if buffer[position] != rune('N') {
												goto l105
											}
											position++
										}
									l108:
										{
											position110, tokenIndex110 := position, tokenIndex
											if buffer[position] != rune('d') {
												goto l111
											}
											position++
											goto l110
										l111:
											position, tokenIndex = position110, tokenIndex110
											if buffer[position] != rune('D') {
												goto l105
											}
											position++
										}
									l110:
										goto l85
									l105:
										position, tokenIndex = position85, tokenIndex85
										{
											position113, tokenIndex113 := position, tokenIndex
											if buffer[position] != rune('d') {
												goto l114
											}
											position++
											goto l113
										l114:
											position, tokenIndex = position113, tokenIndex113
											if buffer[position] != rune('D') {
												goto l112
											}
											position++
										}
									l113:
										{
											position115, tokenIndex115 := position, tokenIndex
											if buffer[position] != rune('e') {
												goto l116
											}
											position++
											goto l115
										l116:
											position, tokenIndex = position115, tokenIndex115
											if buffer[position] != rune('E') {
												goto l112
											}
											position++
										}
									l115:
										{
											position117, tokenIndex117 := position, tokenIndex
											if buffer[position] != rune('l') {
												goto l118
											}
											position++
											goto l117
										l118:
											position, tokenIndex = position117, tokenIndex117
											if buffer[position] != rune('L') {
												goto l112
											}
											position++
										}
									l117:
										{
											position119, tokenIndex119 := position, tokenIndex
											if buffer[position] != rune('e') {
												goto l120
											}
											position++
											goto l119
										l120:
											position, tokenIndex = position119, tokenIndex119
											if buffer[position] != rune('E') {
												goto l112
											}
											position++
										}
									l119:
										{
											position121, tokenIndex121 := position, tokenIndex
											if buffer[position] != rune('t') {
												goto l122
											}
											position++
											goto l121
										l122:
											position, tokenIndex = position121, tokenIndex121
											if buffer[position] != rune('T') {
												goto l112
											}
											position++
										}
									l121:
										{
											position123, tokenIndex123 := position, tokenIndex
											if buffer[position] != rune('e') {
												goto l124
											}
											position++
											goto l123
										l124:
											position, tokenIndex = position123, tokenIndex123
											if buffer[position] != rune('E') {
												goto l112
											}
											position++
										}
									l123:
										goto l85
									l112:
										position, tokenIndex = position85, tokenIndex85
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
												goto l125
											}
											position++
										}
									l126:
										{
											position128, tokenIndex128 := position, tokenIndex
											if buffer[position] != rune('e') {
												goto l129
											}
											position++
											goto l128
										l129:
											position, tokenIndex = position128, tokenIndex128
											if buffer[position] != rune('E') {
												goto l125
											}
											position++
										}
									l128:
										{
											position130, tokenIndex130 := position, tokenIndex
											if buffer[position] != rune('r') {
												goto l131
											}
											position++
											goto l130
										l131:
											position, tokenIndex = position130, tokenIndex130
											if buffer[position] != rune('R') {
												goto l125
											}
											position++
										}
									l130:
										{
											position132, tokenIndex132 := position, tokenIndex
											if buffer[position] != rune('e') {
												goto l133
											}
											position++
											goto l132
										l133:
											position, tokenIndex = position132, tokenIndex132
											if buffer[position] != rune('E') {
												goto l125
											}
											position++
										}
									l132:
										{
											position134, tokenIndex134 := position, tokenIndex
											if buffer[position] != rune('f') {
												goto l135
											}
											position++
											goto l134
										l135:
											position, tokenIndex = position134, tokenIndex134
											if buffer[position] != rune('F') {
												goto l125
											}
											position++
										}
									l134:
										goto l85
									l125:
										position, tokenIndex = position85, tokenIndex85
										{
											position137, tokenIndex137 := position, tokenIndex
											if buffer[position] != rune('d') {
												goto l138
											}
											position++
											goto l137
										l138:
											position, tokenIndex = position137, tokenIndex137
											if buffer[position] != rune('D') {
												goto l136
											}
											position++
										}
									l137:
										{
											position139, tokenIndex139 := position, tokenIndex
											if buffer[position] != rune('i') {
												goto l140
											}
											position++
											goto l139
										l140:
											position, tokenIndex = position139, tokenIndex139
											if buffer[position] != rune('I') {
												goto l136
											}
											position++
										}
									l139:
										{
											position141, tokenIndex141 := position, tokenIndex
											if buffer[position] != rune('v') {
												goto l142
											}
											position++
											goto l141
										l142:
											position, tokenIndex = position141, tokenIndex141
											if buffer[position] != rune('V') {
												goto l136
											}
											position++
										}
									l141:
										{
											position143, tokenIndex143 := position, tokenIndex
											if buffer[position] != rune('a') {
												goto l144
											}
											position++
											goto l143
										l144:
											position, tokenIndex = position143, tokenIndex143
											if buffer[position] != rune('A') {
												goto l136
											}
											position++
										}
									l143:
										{
											position145, tokenIndex145 := position, tokenIndex
											if buffer[position] != rune('s') {
												goto l146
											}
											position++
											goto l145
										l146:
											position, tokenIndex = position145, tokenIndex145
											if buffer[position] != rune('S') {
												goto l136
											}
											position++
										}
									l145:
										{
											position147, tokenIndex147 := position, tokenIndex
											if buffer[position] != rune('s') {
												goto l148
											}
											position++
											goto l147
										l148:
											position, tokenIndex = position147, tokenIndex147
											if buffer[position] != rune('S') {
												goto l136
											}
											position++
										}
									l147:
										{
											position149, tokenIndex149 := position, tokenIndex
											if buffer[position] != rune('i') {
												goto l150
											}
											position++
											goto l149
										l150:
											position, tokenIndex = position149, tokenIndex149
											if buffer[position] != rune('I') {
												goto l136
											}
											position++
										}
									l149:
										{
											position151, tokenIndex151 := position, tokenIndex
											if buffer[position] != rune('g') {
												goto l152
											}
											position++
											goto l151
										l152:
											position, tokenIndex = position151, tokenIndex151
											if buffer[position] != rune('G') {
												goto l136
											}
											position++
										}
									l151:
										{
											position153, tokenIndex153 := position, tokenIndex
											if buffer[position] != rune('n') {
												goto l154
											}
											position++
											goto l153
										l154:
											position, tokenIndex = position153, tokenIndex153
											if buffer[position] != rune('N') {
												goto l136
											}
											position++
										}
									l153:
										goto l85
									l136:
										position, tokenIndex = position85, tokenIndex85
										{
											position156, tokenIndex156 := position, tokenIndex
											if buffer[position] != rune('g') {
												goto l157
											}
											position++
											goto l156
										l157:
											position, tokenIndex = position156, tokenIndex156
											if buffer[position] != rune('G') {
												goto l155
											}
											position++
										}
									l156:
										{
											position158, tokenIndex158 := position, tokenIndex
											if buffer[position] != rune('e') {
												goto l159
											}
											position++
											goto l158
										l159:
											position, tokenIndex = position158, tokenIndex158
											if buffer[position] != rune('E') {
												goto l155
											}
											position++
										}
									l158:
										goto l85
									l155:
										position, tokenIndex = position85, tokenIndex85
										{
											position161, tokenIndex161 := position, tokenIndex
											if buffer[position] != rune('g') {
												goto l162
											}
											position++
											goto l161
										l162:
											position, tokenIndex = position161, tokenIndex161
											if buffer[position] != rune('G') {
												goto l160
											}
											position++
										}
									l161:
										{
											position163, tokenIndex163 := position, tokenIndex
											if buffer[position] != rune('e') {
												goto l164
											}
											position++
											goto l163
										l164:
											position, tokenIndex = position163, tokenIndex163
											if buffer[position] != rune('E') {
												goto l160
											}
											position++
										}
									l163:
										goto l85
									l160:
										position, tokenIndex = position85, tokenIndex85
										{
											position166, tokenIndex166 := position, tokenIndex
											if buffer[position] != rune('l') {
												goto l167
											}
											position++
											goto l166
										l167:
											position, tokenIndex = position166, tokenIndex166
											if buffer[position] != rune('L') {
												goto l165
											}
											position++
										}
									l166:
										{
											position168, tokenIndex168 := position, tokenIndex
											if buffer[position] != rune('a') {
												goto l169
											}
											position++
											goto l168
										l169:
											position, tokenIndex = position168, tokenIndex168
											if buffer[position] != rune('A') {
												goto l165
											}
											position++
										}
									l168:
										{
											position170, tokenIndex170 := position, tokenIndex
											if buffer[position] != rune('n') {
												goto l171
											}
											position++
											goto l170
										l171:
											position, tokenIndex = position170, tokenIndex170
											if buffer[position] != rune('N') {
												goto l165
											}
											position++
										}
									l170:
										{
											position172, tokenIndex172 := position, tokenIndex
											if buffer[position] != rune('d') {
												goto l173
											}
											position++
											goto l172
										l173:
											position, tokenIndex = position172, tokenIndex172
											if buffer[position] != rune('D') {
												goto l165
											}
											position++
										}
									l172:
										goto l85
									l165:
										position, tokenIndex = position85, tokenIndex85
										{
											position175, tokenIndex175 := position, tokenIndex
											if buffer[position] != rune('l') {
												goto l176
											}
											position++
											goto l175
										l176:
											position, tokenIndex = position175, tokenIndex175
											if buffer[position] != rune('L') {
												goto l174
											}
											position++
										}
									l175:
										{
											position177, tokenIndex177 := position, tokenIndex
											if buffer[position] != rune('e') {
												goto l178
											}
											position++
											goto l177
										l178:
											position, tokenIndex = position177, tokenIndex177
											if buffer[position] != rune('E') {
												goto l174
											}
											position++
										}
									l177:
										goto l85
									l174:
										position, tokenIndex = position85, tokenIndex85
										{
											position180, tokenIndex180 := position, tokenIndex
											if buffer[position] != rune('l') {
												goto l181
											}
											position++
											goto l180
										l181:
											position, tokenIndex = position180, tokenIndex180
											if buffer[position] != rune('L') {
												goto l179
											}
											position++
										}
									l180:
										{
											position182, tokenIndex182 := position, tokenIndex
											if buffer[position] != rune('n') {
												goto l183
											}
											position++
											goto l182
										l183:
											position, tokenIndex = position182, tokenIndex182
											if buffer[position] != rune('N') {
												goto l179
											}
											position++
										}
									l182:
										{
											position184, tokenIndex184 := position, tokenIndex
											if buffer[position] != rune('o') {
												goto l185
											}
											position++
											goto l184
										l185:
											position, tokenIndex = position184, tokenIndex184
											if buffer[position] != rune('O') {
												goto l179
											}
											position++
										}
									l184:
										{
											position186, tokenIndex186 := position, tokenIndex
											if buffer[position] != rune('t') {
												goto l187
											}
											position++
											goto l186
										l187:
											position, tokenIndex = position186, tokenIndex186
											if buffer[position] != rune('T') {
												goto l179
											}
											position++
										}
									l186:
										goto l85
									l179:
										position, tokenIndex = position85, tokenIndex85
										{
											position189, tokenIndex189 := position, tokenIndex
											if buffer[position] != rune('l') {
												goto l190
											}
											position++
											goto l189
										l190:
											position, tokenIndex = position189, tokenIndex189
											if buffer[position] != rune('L') {
												goto l188
											}
											position++
										}
									l189:
										{
											position191, tokenIndex191 := position, tokenIndex
											if buffer[position] != rune('o') {
												goto l192
											}
											position++
											goto l191
										l192:
											position, tokenIndex = position191, tokenIndex191
											if buffer[position] != rune('O') {
												goto l188
											}
											position++
										}
									l191:
										{
											position193, tokenIndex193 := position, tokenIndex
											if buffer[position] != rune('r') {
												goto l194
											}
											position++
											goto l193
										l194:
											position, tokenIndex = position193, tokenIndex193
											if buffer[position] != rune('R') {
												goto l188
											}
											position++
										}
									l193:
										goto l85
									l188:
										position, tokenIndex = position85, tokenIndex85
										{
											position196, tokenIndex196 := position, tokenIndex
											if buffer[position] != rune('l') {
												goto l197
											}
											position++
											goto l196
										l197:
											position, tokenIndex = position196, tokenIndex196
											if buffer[position] != rune('L') {
												goto l195
											}
											position++
										}
									l196:
										{
											position198, tokenIndex198 := position, tokenIndex
											if buffer[position] != rune('s') {
												goto l199
											}
											position++
											goto l198
										l199:
											position, tokenIndex = position198, tokenIndex198
											if buffer[position] != rune('S') {
												goto l195
											}
											position++
										}
									l198:
										{
											position200, tokenIndex200 := position, tokenIndex
											if buffer[position] != rune('h') {
												goto l201
											}
											position++
											goto l200
										l201:
											position, tokenIndex = position200, tokenIndex200
											if buffer[position] != rune('H') {
												goto l195
											}
											position++
										}
									l200:
										{
											position202, tokenIndex202 := position, tokenIndex
											if buffer[position] != rune('i') {
												goto l203
											}
											position++
											goto l202
										l203:
											position, tokenIndex = position202, tokenIndex202
											if buffer[position] != rune('I') {
												goto l195
											}
											position++
										}
									l202:
										{
											position204, tokenIndex204 := position, tokenIndex
											if buffer[position] != rune('f') {
												goto l205
											}
											position++
											goto l204
										l205:
											position, tokenIndex = position204, tokenIndex204
											if buffer[position] != rune('F') {
												goto l195
											}
											position++
										}
									l204:
										{
											position206, tokenIndex206 := position, tokenIndex
											if buffer[position] != rune('t') {
												goto l207
											}
											position++
											goto l206
										l207:
											position, tokenIndex = position206, tokenIndex206
											if buffer[position] != rune('T') {
												goto l195
											}
											position++
										}
									l206:
										{
											position208, tokenIndex208 := position, tokenIndex
											if buffer[position] != rune('a') {
												goto l209
											}
											position++
											goto l208
										l209:
											position, tokenIndex = position208, tokenIndex208
											if buffer[position] != rune('A') {
												goto l195
											}
											position++
										}
									l208:
										{
											position210, tokenIndex210 := position, tokenIndex
											if buffer[position] != rune('s') {
												goto l211
											}
											position++
											goto l210
										l211:
											position, tokenIndex = position210, tokenIndex210
											if buffer[position] != rune('S') {
												goto l195
											}
											position++
										}
									l210:
										{
											position212, tokenIndex212 := position, tokenIndex
											if buffer[position] != rune('s') {
												goto l213
											}
											position++
											goto l212
										l213:
											position, tokenIndex = position212, tokenIndex212
											if buffer[position] != rune('S') {
												goto l195
											}
											position++
										}
									l212:
										{
											position214, tokenIndex214 := position, tokenIndex
											if buffer[position] != rune('i') {
												goto l215
											}
											position++
											goto l214
										l215:
											position, tokenIndex = position214, tokenIndex214
											if buffer[position] != rune('I') {
												goto l195
											}
											position++
										}
									l214:
										{
											position216, tokenIndex216 := position, tokenIndex
											if buffer[position] != rune('g') {
												goto l217
											}
											position++
											goto l216
										l217:
											position, tokenIndex = position216, tokenIndex216
											if buffer[position] != rune('G') {
												goto l195
											}
											position++
										}
									l216:
										{
											position218, tokenIndex218 := position, tokenIndex
											if buffer[position] != rune('n') {
												goto l219
											}
											position++
											goto l218
										l219:
											position, tokenIndex = position218, tokenIndex218
											if buffer[position] != rune('N') {
												goto l195
											}
											position++
										}
									l218:
										goto l85
									l195:
										position, tokenIndex = position85, tokenIndex85
										{
											position221, tokenIndex221 := position, tokenIndex
											if buffer[position] != rune('l') {
												goto l222
											}
											position++
											goto l221
										l222:
											position, tokenIndex = position221, tokenIndex221
											if buffer[position] != rune('L') {
												goto l220
											}
											position++
										}
									l221:
										{
											position223, tokenIndex223 := position, tokenIndex
											if buffer[position] != rune('s') {
												goto l224
											}
											position++
											goto l223
										l224:
											position, tokenIndex = position223, tokenIndex223
											if buffer[position] != rune('S') {
												goto l220
											}
											position++
										}
									l223:
										{
											position225, tokenIndex225 := position, tokenIndex
											if buffer[position] != rune('h') {
												goto l226
											}
											position++
											goto l225
										l226:
											position, tokenIndex = position225, tokenIndex225
											if buffer[position] != rune('H') {
												goto l220
											}
											position++
										}
									l225:
										{
											position227, tokenIndex227 := position, tokenIndex
											if buffer[position] != rune('i') {
												goto l228
											}
											position++
											goto l227
										l228:
											position, tokenIndex = position227, tokenIndex227
											if buffer[position] != rune('I') {
												goto l220
											}
											position++
										}
									l227:
										{
											position229, tokenIndex229 := position, tokenIndex
											if buffer[position] != rune('f') {
												goto l230
											}
											position++
											goto l229
										l230:
											position, tokenIndex = position229, tokenIndex229
											if buffer[position] != rune('F') {
												goto l220
											}
											position++
										}
									l229:
										{
											position231, tokenIndex231 := position, tokenIndex
											if buffer[position] != rune('t') {
												goto l232
											}
											position++
											goto l231
										l232:
											position, tokenIndex = position231, tokenIndex231
											if buffer[position] != rune('T') {
												goto l220
											}
											position++
										}
									l231:
										goto l85
									l220:
										position, tokenIndex = position85, tokenIndex85
										{
											position234, tokenIndex234 := position, tokenIndex
											if buffer[position] != rune('m') {
												goto l235
											}
											position++
											goto l234
										l235:
											position, tokenIndex = position234, tokenIndex234
											if buffer[position] != rune('M') {
												goto l233
											}
											position++
										}
									l234:
										{
											position236, tokenIndex236 := position, tokenIndex
											if buffer[position] != rune('i') {
												goto l237
											}
											position++
											goto l236
										l237:
											position, tokenIndex = position236, tokenIndex236
											if buffer[position] != rune('I') {
												goto l233
											}
											position++
										}
									l236:
										{
											position238, tokenIndex238 := position, tokenIndex
											if buffer[position] != rune('n') {
												goto l239
											}
											position++
											goto l238
										l239:
											position, tokenIndex = position238, tokenIndex238
											if buffer[position] != rune('N') {
												goto l233
											}
											position++
										}
									l238:
										{
											position240, tokenIndex240 := position, tokenIndex
											if buffer[position] != rune('u') {
												goto l241
											}
											position++
											goto l240
										l241:
											position, tokenIndex = position240, tokenIndex240
											if buffer[position] != rune('U') {
												goto l233
											}
											position++
										}
									l240:
										{
											position242, tokenIndex242 := position, tokenIndex
											if buffer[position] != rune('s') {
												goto l243
											}
											position++
											goto l242
										l243:
											position, tokenIndex = position242, tokenIndex242
											if buffer[position] != rune('S') {
												goto l233
											}
											position++
										}
									l242:
										{
											position244, tokenIndex244 := position, tokenIndex
											if buffer[position] != rune('a') {
												goto l245
											}
											position++
											goto l244
										l245:
											position, tokenIndex = position244, tokenIndex244
											if buffer[position] != rune('A') {
												goto l233
											}
											position++
										}
									l244:
										{
											position246, tokenIndex246 := position, tokenIndex
											if buffer[position] != rune('s') {
												goto l247
											}
											position++
											goto l246
										l247:
											position, tokenIndex = position246, tokenIndex246
											if buffer[position] != rune('S') {
												goto l233
											}
											position++
										}
									l246:
										{
											position248, tokenIndex248 := position, tokenIndex
											if buffer[position] != rune('s') {
												goto l249
											}
											position++
											goto l248
										l249:
											position, tokenIndex = position248, tokenIndex248
											if buffer[position] != rune('S') {
												goto l233
											}
											position++
										}
									l248:
										{
											position250, tokenIndex250 := position, tokenIndex
											if buffer[position] != rune('i') {
												goto l251
											}
											position++
											goto l250
										l251:
											position, tokenIndex = position250, tokenIndex250
											if buffer[position] != rune('I') {
												goto l233
											}
											position++
										}
									l250:
										{
											position252, tokenIndex252 := position, tokenIndex
											if buffer[position] != rune('g') {
												goto l253
											}
											position++
											goto l252
										l253:
											position, tokenIndex = position252, tokenIndex252
											if buffer[position] != rune('G') {
												goto l233
											}
											position++
										}
									l252:
										{
											position254, tokenIndex254 := position, tokenIndex
											if buffer[position] != rune('n') {
												goto l255
											}
											position++
											goto l254
										l255:
											position, tokenIndex = position254, tokenIndex254
											if buffer[position] != rune('N') {
												goto l233
											}
											position++
										}
									l254:
										goto l85
									l233:
										position, tokenIndex = position85, tokenIndex85
										{
											position257, tokenIndex257 := position, tokenIndex
											if buffer[position] != rune('m') {
												goto l258
											}
											position++
											goto l257
										l258:
											position, tokenIndex = position257, tokenIndex257
											if buffer[position] != rune('M') {
												goto l256
											}
											position++
										}
									l257:
										{
											position259, tokenIndex259 := position, tokenIndex
											if buffer[position] != rune('i') {
												goto l260
											}
											position++
											goto l259
										l260:
											position, tokenIndex = position259, tokenIndex259
											if buffer[position] != rune('I') {
												goto l256
											}
											position++
										}
									l259:
										{
											position261, tokenIndex261 := position, tokenIndex
											if buffer[position] != rune('n') {
												goto l262
											}
											position++
											goto l261
										l262:
											position, tokenIndex = position261, tokenIndex261
											if buffer[position] != rune('N') {
												goto l256
											}
											position++
										}
									l261:
										{
											position263, tokenIndex263 := position, tokenIndex
											if buffer[position] != rune('u') {
												goto l264
											}
											position++
											goto l263
										l264:
											position, tokenIndex = position263, tokenIndex263
											if buffer[position] != rune('U') {
												goto l256
											}
											position++
										}
									l263:
										{
											position265, tokenIndex265 := position, tokenIndex
											if buffer[position] != rune('s') {
												goto l266
											}
											position++
											goto l265
										l266:
											position, tokenIndex = position265, tokenIndex265
											if buffer[position] != rune('S') {
												goto l256
											}
											position++
										}
									l265:
										goto l85
									l256:
										position, tokenIndex = position85, tokenIndex85
										{
											position268, tokenIndex268 := position, tokenIndex
											if buffer[position] != rune('m') {
												goto l269
											}
											position++
											goto l268
										l269:
											position, tokenIndex = position268, tokenIndex268
											if buffer[position] != rune('M') {
												goto l267
											}
											position++
										}
									l268:
										{
											position270, tokenIndex270 := position, tokenIndex
											if buffer[position] != rune('u') {
												goto l271
											}
											position++
											goto l270
										l271:
											position, tokenIndex = position270, tokenIndex270
											if buffer[position] != rune('U') {
												goto l267
											}
											position++
										}
									l270:
										{
											position272, tokenIndex272 := position, tokenIndex
											if buffer[position] != rune('l') {
												goto l273
											}
											position++
											goto l272
										l273:
											position, tokenIndex = position272, tokenIndex272
											if buffer[position] != rune('L') {
												goto l267
											}
											position++
										}
									l272:
										{
											position274, tokenIndex274 := position, tokenIndex
											if buffer[position] != rune('t') {
												goto l275
											}
											position++
											goto l274
										l275:
											position, tokenIndex = position274, tokenIndex274
											if buffer[position] != rune('T') {
												goto l267
											}
											position++
										}
									l274:
										{
											position276, tokenIndex276 := position, tokenIndex
											if buffer[position] != rune('a') {
												goto l277
											}
											position++
											goto l276
										l277:
											position, tokenIndex = position276, tokenIndex276
											if buffer[position] != rune('A') {
												goto l267
											}
											position++
										}
									l276:
										{
											position278, tokenIndex278 := position, tokenIndex
											if buffer[position] != rune('s') {
												goto l279
											}
											position++
											goto l278
										l279:
											position, tokenIndex = position278, tokenIndex278
											if buffer[position] != rune('S') {
												goto l267
											}
											position++
										}
									l278:
										{
											position280, tokenIndex280 := position, tokenIndex
											if buffer[position] != rune('s') {
												goto l281
											}
											position++
											goto l280
										l281:
											position, tokenIndex = position280, tokenIndex280
											if buffer[position] != rune('S') {
												goto l267
											}
											position++
										}
									l280:
										{
											position282, tokenIndex282 := position, tokenIndex
											if buffer[position] != rune('i') {
												goto l283
											}
											position++
											goto l282
										l283:
											position, tokenIndex = position282, tokenIndex282
											if buffer[position] != rune('I') {
												goto l267
											}
											position++
										}
									l282:
										{
											position284, tokenIndex284 := position, tokenIndex
											if buffer[position] != rune('g') {
												goto l285
											}
											position++
											goto l284
										l285:
											position, tokenIndex = position284, tokenIndex284
											if buffer[position] != rune('G') {
												goto l267
											}
											position++
										}
									l284:
										{
											position286, tokenIndex286 := position, tokenIndex
											if buffer[position] != rune('n') {
												goto l287
											}
											position++
											goto l286
										l287:
											position, tokenIndex = position286, tokenIndex286
											if buffer[position] != rune('N') {
												goto l267
											}
											position++
										}
									l286:
										goto l85
									l267:
										position, tokenIndex = position85, tokenIndex85
										{
											position289, tokenIndex289 := position, tokenIndex
											if buffer[position] != rune('n') {
												goto l290
											}
											position++
											goto l289
										l290:
											position, tokenIndex = position289, tokenIndex289
											if buffer[position] != rune('N') {
												goto l288
											}
											position++
										}
									l289:
										{
											position291, tokenIndex291 := position, tokenIndex
											if buffer[position] != rune('e') {
												goto l292
											}
											position++
											goto l291
										l292:
											position, tokenIndex = position291, tokenIndex291
											if buffer[position] != rune('E') {
												goto l288
											}
											position++
										}
									l291:
										{
											position293, tokenIndex293 := position, tokenIndex
											if buffer[position] != rune('g') {
												goto l294
											}
											position++
											goto l293
										l294:
											position, tokenIndex = position293, tokenIndex293
											if buffer[position] != rune('G') {
												goto l288
											}
											position++
										}
									l293:
										goto l85
									l288:
										position, tokenIndex = position85, tokenIndex85
										{
											position296, tokenIndex296 := position, tokenIndex
											if buffer[position] != rune('n') {
												goto l297
											}
											position++
											goto l296
										l297:
											position, tokenIndex = position296, tokenIndex296
											if buffer[position] != rune('N') {
												goto l295
											}
											position++
										}
									l296:
										{
											position298, tokenIndex298 := position, tokenIndex
											if buffer[position] != rune('e') {
												goto l299
											}
											position++
											goto l298
										l299:
											position, tokenIndex = position298, tokenIndex298
											if buffer[position] != rune('E') {
												goto l295
											}
											position++
										}
									l298:
										{
											position300, tokenIndex300 := position, tokenIndex
											if buffer[position] != rune('w') {
												goto l301
											}
											position++
											goto l300
										l301:
											position, tokenIndex = position300, tokenIndex300
											if buffer[position] != rune('W') {
												goto l295
											}
											position++
										}
									l300:
										goto l85
									l295:
										position, tokenIndex = position85, tokenIndex85
										{
											position303, tokenIndex303 := position, tokenIndex
											if buffer[position] != rune('n') {
												goto l304
											}
											position++
											goto l303
										l304:
											position, tokenIndex = position303, tokenIndex303
											if buffer[position] != rune('N') {
												goto l302
											}
											position++
										}
									l303:
										{
											position305, tokenIndex305 := position, tokenIndex
											if buffer[position] != rune('e') {
												goto l306
											}
											position++
											goto l305
										l306:
											position, tokenIndex = position305, tokenIndex305
											if buffer[position] != rune('E') {
												goto l302
											}
											position++
										}
									l305:
										goto l85
									l302:
										position, tokenIndex = position85, tokenIndex85
										{
											position308, tokenIndex308 := position, tokenIndex
											if buffer[position] != rune('o') {
												goto l309
											}
											position++
											goto l308
										l309:
											position, tokenIndex = position308, tokenIndex308
											if buffer[position] != rune('O') {
												goto l307
											}
											position++
										}
									l308:
										{
											position310, tokenIndex310 := position, tokenIndex
											if buffer[position] != rune('r') {
												goto l311
											}
											position++
											goto l310
										l311:
											position, tokenIndex = position310, tokenIndex310
											if buffer[position] != rune('R') {
												goto l307
											}
											position++
										}
									l310:
										{
											position312, tokenIndex312 := position, tokenIndex
											if buffer[position] != rune('a') {
												goto l313
											}
											position++
											goto l312
										l313:
											position, tokenIndex = position312, tokenIndex312
											if buffer[position] != rune('A') {
												goto l307
											}
											position++
										}
									l312:
										{
											position314, tokenIndex314 := position, tokenIndex
											if buffer[position] != rune('s') {
												goto l315
											}
											position++
											goto l314
										l315:
											position, tokenIndex = position314, tokenIndex314
											if buffer[position] != rune('S') {
												goto l307
											}
											position++
										}
									l314:
										{
											position316, tokenIndex316 := position, tokenIndex
											if buffer[position] != rune('s') {
												goto l317
											}
											position++
											goto l316
										l317:
											position, tokenIndex = position316, tokenIndex316
											if buffer[position] != rune('S') {
												goto l307
											}
											position++
										}
									l316:
										{
											position318, tokenIndex318 := position, tokenIndex
											if buffer[position] != rune('i') {
												goto l319
											}
											position++
											goto l318
										l319:
											position, tokenIndex = position318, tokenIndex318
											if buffer[position] != rune('I') {
												goto l307
											}
											position++
										}
									l318:
										{
											position320, tokenIndex320 := position, tokenIndex
											if buffer[position] != rune('g') {
												goto l321
											}
											position++
											goto l320
										l321:
											position, tokenIndex = position320, tokenIndex320
											if buffer[position] != rune('G') {
												goto l307
											}
											position++
										}
									l320:
										{
											position322, tokenIndex322 := position, tokenIndex
											if buffer[position] != rune('n') {
												goto l323
											}
											position++
											goto l322
										l323:
											position, tokenIndex = position322, tokenIndex322
											if buffer[position] != rune('N') {
												goto l307
											}
											position++
										}
									l322:
										goto l85
									l307:
										position, tokenIndex = position85, tokenIndex85
										{
											position325, tokenIndex325 := position, tokenIndex
											if buffer[position] != rune('p') {
												goto l326
											}
											position++
											goto l325
										l326:
											position, tokenIndex = position325, tokenIndex325
											if buffer[position] != rune('P') {
												goto l324
											}
											position++
										}
									l325:
										{
											position327, tokenIndex327 := position, tokenIndex
											if buffer[position] != rune('l') {
												goto l328
											}
											position++
											goto l327
										l328:
											position, tokenIndex = position327, tokenIndex327
											if buffer[position] != rune('L') {
												goto l324
											}
											position++
										}
									l327:
										{
											position329, tokenIndex329 := position, tokenIndex
											if buffer[position] != rune('u') {
												goto l330
											}
											position++
											goto l329
										l330:
											position, tokenIndex = position329, tokenIndex329
											if buffer[position] != rune('U') {
												goto l324
											}
											position++
										}
									l329:
										{
											position331, tokenIndex331 := position, tokenIndex
											if buffer[position] != rune('s') {
												goto l332
											}
											position++
											goto l331
										l332:
											position, tokenIndex = position331, tokenIndex331
											if buffer[position] != rune('S') {
												goto l324
											}
											position++
										}
									l331:
										{
											position333, tokenIndex333 := position, tokenIndex
											if buffer[position] != rune('a') {
												goto l334
											}
											position++
											goto l333
										l334:
											position, tokenIndex = position333, tokenIndex333
											if buffer[position] != rune('A') {
												goto l324
											}
											position++
										}
									l333:
										{
											position335, tokenIndex335 := position, tokenIndex
											if buffer[position] != rune('s') {
												goto l336
											}
											position++
											goto l335
										l336:
											position, tokenIndex = position335, tokenIndex335
											if buffer[position] != rune('S') {
												goto l324
											}
											position++
										}
									l335:
										{
											position337, tokenIndex337 := position, tokenIndex
											if buffer[position] != rune('s') {
												goto l338
											}
											position++
											goto l337
										l338:
											position, tokenIndex = position337, tokenIndex337
											if buffer[position] != rune('S') {
												goto l324
											}
											position++
										}
									l337:
										{
											position339, tokenIndex339 := position, tokenIndex
											if buffer[position] != rune('i') {
												goto l340
											}
											position++
											goto l339
										l340:
											position, tokenIndex = position339, tokenIndex339
											if buffer[position] != rune('I') {
												goto l324
											}
											position++
										}
									l339:
										{
											position341, tokenIndex341 := position, tokenIndex
											if buffer[position] != rune('g') {
												goto l342
											}
											position++
											goto l341
										l342:
											position, tokenIndex = position341, tokenIndex341
											if buffer[position] != rune('G') {
												goto l324
											}
											position++
										}
									l341:
										{
											position343, tokenIndex343 := position, tokenIndex
											if buffer[position] != rune('n') {
												goto l344
											}
											position++
											goto l343
										l344:
											position, tokenIndex = position343, tokenIndex343
											if buffer[position] != rune('N') {
												goto l324
											}
											position++
										}
									l343:
										goto l85
									l324:
										position, tokenIndex = position85, tokenIndex85
										{
											position346, tokenIndex346 := position, tokenIndex
											if buffer[position] != rune('p') {
												goto l347
											}
											position++
											goto l346
										l347:
											position, tokenIndex = position346, tokenIndex346
											if buffer[position] != rune('P') {
												goto l345
											}
											position++
										}
									l346:
										{
											position348, tokenIndex348 := position, tokenIndex
											if buffer[position] != rune('l') {
												goto l349
											}
											position++
											goto l348
										l349:
											position, tokenIndex = position348, tokenIndex348
											if buffer[position] != rune('L') {
												goto l345
											}
											position++
										}
									l348:
										{
											position350, tokenIndex350 := position, tokenIndex
											if buffer[position] != rune('u') {
												goto l351
											}
											position++
											goto l350
										l351:
											position, tokenIndex = position350, tokenIndex350
											if buffer[position] != rune('U') {
												goto l345
											}
											position++
										}
									l350:
										{
											position352, tokenIndex352 := position, tokenIndex
											if buffer[position] != rune('s') {
												goto l353
											}
											position++
											goto l352
										l353:
											position, tokenIndex = position352, tokenIndex352
											if buffer[position] != rune('S') {
												goto l345
											}
											position++
										}
									l352:
										goto l85
									l345:
										position, tokenIndex = position85, tokenIndex85
										{
											position355, tokenIndex355 := position, tokenIndex
											if buffer[position] != rune('p') {
												goto l356
											}
											position++
											goto l355
										l356:
											position, tokenIndex = position355, tokenIndex355
											if buffer[position] != rune('P') {
												goto l354
											}
											position++
										}
									l355:
										{
											position357, tokenIndex357 := position, tokenIndex
											if buffer[position] != rune('o') {
												goto l358
											}
											position++
											goto l357
										l358:
											position, tokenIndex = position357, tokenIndex357
											if buffer[position] != rune('O') {
												goto l354
											}
											position++
										}
									l357:
										{
											position359, tokenIndex359 := position, tokenIndex
											if buffer[position] != rune('s') {
												goto l360
											}
											position++
											goto l359
										l360:
											position, tokenIndex = position359, tokenIndex359
											if buffer[position] != rune('S') {
												goto l354
											}
											position++
										}
									l359:
										{
											position361, tokenIndex361 := position, tokenIndex
											if buffer[position] != rune('t') {
												goto l362
											}
											position++
											goto l361
										l362:
											position, tokenIndex = position361, tokenIndex361
											if buffer[position] != rune('T') {
												goto l354
											}
											position++
										}
									l361:
										{
											position363, tokenIndex363 := position, tokenIndex
											if buffer[position] != rune('d') {
												goto l364
											}
											position++
											goto l363
										l364:
											position, tokenIndex = position363, tokenIndex363
											if buffer[position] != rune('D') {
												goto l354
											}
											position++
										}
									l363:
										{
											position365, tokenIndex365 := position, tokenIndex
											if buffer[position] != rune('e') {
												goto l366
											}
											position++
											goto l365
										l366:
											position, tokenIndex = position365, tokenIndex365
											if buffer[position] != rune('E') {
												goto l354
											}
											position++
										}
									l365:
										{
											position367, tokenIndex367 := position, tokenIndex
											if buffer[position] != rune('c') {
												goto l368
											}
											position++
											goto l367
										l368:
											position, tokenIndex = position367, tokenIndex367
											if buffer[position] != rune('C') {
												goto l354
											}
											position++
										}
									l367:
										goto l85
									l354:
										position, tokenIndex = position85, tokenIndex85
										{
											position370, tokenIndex370 := position, tokenIndex
											if buffer[position] != rune('p') {
												goto l371
											}
											position++
											goto l370
										l371:
											position, tokenIndex = position370, tokenIndex370
											if buffer[position] != rune('P') {
												goto l369
											}
											position++
										}
									l370:
										{
											position372, tokenIndex372 := position, tokenIndex
											if buffer[position] != rune('o') {
												goto l373
											}
											position++
											goto l372
										l373:
											position, tokenIndex = position372, tokenIndex372
											if buffer[position] != rune('O') {
												goto l369
											}
											position++
										}
									l372:
										{
											position374, tokenIndex374 := position, tokenIndex
											if buffer[position] != rune('s') {
												goto l375
											}
											position++
											goto l374
										l375:
											position, tokenIndex = position374, tokenIndex374
											if buffer[position] != rune('S') {
												goto l369
											}
											position++
										}
									l374:
										{
											position376, tokenIndex376 := position, tokenIndex
											if buffer[position] != rune('t') {
												goto l377
											}
											position++
											goto l376
										l377:
											position, tokenIndex = position376, tokenIndex376
											if buffer[position] != rune('T') {
												goto l369
											}
											position++
										}
									l376:
										{
											position378, tokenIndex378 := position, tokenIndex
											if buffer[position] != rune('i') {
												goto l379
											}
											position++
											goto l378
										l379:
											position, tokenIndex = position378, tokenIndex378
											if buffer[position] != rune('I') {
												goto l369
											}
											position++
										}
									l378:
										{
											position380, tokenIndex380 := position, tokenIndex
											if buffer[position] != rune('n') {
												goto l381
											}
											position++
											goto l380
										l381:
											position, tokenIndex = position380, tokenIndex380
											if buffer[position] != rune('N') {
												goto l369
											}
											position++
										}
									l380:
										{
											position382, tokenIndex382 := position, tokenIndex
											if buffer[position] != rune('c') {
												goto l383
											}
											position++
											goto l382
										l383:
											position, tokenIndex = position382, tokenIndex382
											if buffer[position] != rune('C') {
												goto l369
											}
											position++
										}
									l382:
										goto l85
									l369:
										position, tokenIndex = position85, tokenIndex85
										{
											position385, tokenIndex385 := position, tokenIndex
											if buffer[position] != rune('p') {
												goto l386
											}
											position++
											goto l385
										l386:
											position, tokenIndex = position385, tokenIndex385
											if buffer[position] != rune('P') {
												goto l384
											}
											position++
										}
									l385:
										{
											position387, tokenIndex387 := position, tokenIndex
											if buffer[position] != rune('o') {
												goto l388
											}
											position++
											goto l387
										l388:
											position, tokenIndex = position387, tokenIndex387
											if buffer[position] != rune('O') {
												goto l384
											}
											position++
										}
									l387:
										{
											position389, tokenIndex389 := position, tokenIndex
											if buffer[position] != rune('s') {
												goto l390
											}
											position++
											goto l389
										l390:
											position, tokenIndex = position389, tokenIndex389
											if buffer[position] != rune('S') {
												goto l384
											}
											position++
										}
									l389:
										goto l85
									l384:
										position, tokenIndex = position85, tokenIndex85
										{
											position392, tokenIndex392 := position, tokenIndex
											if buffer[position] != rune('p') {
												goto l393
											}
											position++
											goto l392
										l393:
											position, tokenIndex = position392, tokenIndex392
											if buffer[position] != rune('P') {
												goto l391
											}
											position++
										}
									l392:
										{
											position394, tokenIndex394 := position, tokenIndex
											if buffer[position] != rune('r') {
												goto l395
											}
											position++
											goto l394
										l395:
											position, tokenIndex = position394, tokenIndex394
											if buffer[position] != rune('R') {
												goto l391
											}
											position++
										}
									l394:
										{
											position396, tokenIndex396 := position, tokenIndex
											if buffer[position] != rune('e') {
												goto l397
											}
											position++
											goto l396
										l397:
											position, tokenIndex = position396, tokenIndex396
											if buffer[position] != rune('E') {
												goto l391
											}
											position++
										}
									l396:
										{
											position398, tokenIndex398 := position, tokenIndex
											if buffer[position] != rune('d') {
												goto l399
											}
											position++
											goto l398
										l399:
											position, tokenIndex = position398, tokenIndex398
											if buffer[position] != rune('D') {
												goto l391
											}
											position++
										}
									l398:
										{
											position400, tokenIndex400 := position, tokenIndex
											if buffer[position] != rune('e') {
												goto l401
											}
											position++
											goto l400
										l401:
											position, tokenIndex = position400, tokenIndex400
											if buffer[position] != rune('E') {
												goto l391
											}
											position++
										}
									l400:
										{
											position402, tokenIndex402 := position, tokenIndex
											if buffer[position] != rune('c') {
												goto l403
											}
											position++
											goto l402
										l403:
											position, tokenIndex = position402, tokenIndex402
											if buffer[position] != rune('C') {
												goto l391
											}
											position++
										}
									l402:
										goto l85
									l391:
										position, tokenIndex = position85, tokenIndex85
										{
											position405, tokenIndex405 := position, tokenIndex
											if buffer[position] != rune('r') {
												goto l406
											}
											position++
											goto l405
										l406:
											position, tokenIndex = position405, tokenIndex405
											if buffer[position] != rune('R') {
												goto l404
											}
											position++
										}
									l405:
										{
											position407, tokenIndex407 := position, tokenIndex
											if buffer[position] != rune('e') {
												goto l408
											}
											position++
											goto l407
										l408:
											position, tokenIndex = position407, tokenIndex407
											if buffer[position] != rune('E') {
												goto l404
											}
											position++
										}
									l407:
										{
											position409, tokenIndex409 := position, tokenIndex
											if buffer[position] != rune('f') {
												goto l410
											}
											position++
											goto l409
										l410:
											position, tokenIndex = position409, tokenIndex409
											if buffer[position] != rune('F') {
												goto l404
											}
											position++
										}
									l409:
										goto l85
									l404:
										position, tokenIndex = position85, tokenIndex85
										{
											position412, tokenIndex412 := position, tokenIndex
											if buffer[position] != rune('r') {
												goto l413
											}
											position++
											goto l412
										l413:
											position, tokenIndex = position412, tokenIndex412
											if buffer[position] != rune('R') {
												goto l411
											}
											position++
										}
									l412:
										{
											position414, tokenIndex414 := position, tokenIndex
											if buffer[position] != rune('s') {
												goto l415
											}
											position++
											goto l414
										l415:
											position, tokenIndex = position414, tokenIndex414
											if buffer[position] != rune('S') {
												goto l411
											}
											position++
										}
									l414:
										{
											position416, tokenIndex416 := position, tokenIndex
											if buffer[position] != rune('h') {
												goto l417
											}
											position++
											goto l416
										l417:
											position, tokenIndex = position416, tokenIndex416
											if buffer[position] != rune('H') {
												goto l411
											}
											position++
										}
									l416:
										{
											position418, tokenIndex418 := position, tokenIndex
											if buffer[position] != rune('i') {
												goto l419
											}
											position++
											goto l418
										l419:
											position, tokenIndex = position418, tokenIndex418
											if buffer[position] != rune('I') {
												goto l411
											}
											position++
										}
									l418:
										{
											position420, tokenIndex420 := position, tokenIndex
											if buffer[position] != rune('f') {
												goto l421
											}
											position++
											goto l420
										l421:
											position, tokenIndex = position420, tokenIndex420
											if buffer[position] != rune('F') {
												goto l411
											}
											position++
										}
									l420:
										{
											position422, tokenIndex422 := position, tokenIndex
											if buffer[position] != rune('t') {
												goto l423
											}
											position++
											goto l422
										l423:
											position, tokenIndex = position422, tokenIndex422
											if buffer[position] != rune('T') {
												goto l411
											}
											position++
										}
									l422:
										{
											position424, tokenIndex424 := position, tokenIndex
											if buffer[position] != rune('a') {
												goto l425
											}
											position++
											goto l424
										l425:
											position, tokenIndex = position424, tokenIndex424
											if buffer[position] != rune('A') {
												goto l411
											}
											position++
										}
									l424:
										{
											position426, tokenIndex426 := position, tokenIndex
											if buffer[position] != rune('s') {
												goto l427
											}
											position++
											goto l426
										l427:
											position, tokenIndex = position426, tokenIndex426
											if buffer[position] != rune('S') {
												goto l411
											}
											position++
										}
									l426:
										{
											position428, tokenIndex428 := position, tokenIndex
											if buffer[position] != rune('s') {
												goto l429
											}
											position++
											goto l428
										l429:
											position, tokenIndex = position428, tokenIndex428
											if buffer[position] != rune('S') {
												goto l411
											}
											position++
										}
									l428:
										{
											position430, tokenIndex430 := position, tokenIndex
											if buffer[position] != rune('i') {
												goto l431
											}
											position++
											goto l430
										l431:
											position, tokenIndex = position430, tokenIndex430
											if buffer[position] != rune('I') {
												goto l411
											}
											position++
										}
									l430:
										{
											position432, tokenIndex432 := position, tokenIndex
											if buffer[position] != rune('g') {
												goto l433
											}
											position++
											goto l432
										l433:
											position, tokenIndex = position432, tokenIndex432
											if buffer[position] != rune('G') {
												goto l411
											}
											position++
										}
									l432:
										{
											position434, tokenIndex434 := position, tokenIndex
											if buffer[position] != rune('n') {
												goto l435
											}
											position++
											goto l434
										l435:
											position, tokenIndex = position434, tokenIndex434
											if buffer[position] != rune('N') {
												goto l411
											}
											position++
										}
									l434:
										goto l85
									l411:
										position, tokenIndex = position85, tokenIndex85
										{
											position437, tokenIndex437 := position, tokenIndex
											if buffer[position] != rune('v') {
												goto l438
											}
											position++
											goto l437
										l438:
											position, tokenIndex = position437, tokenIndex437
											if buffer[position] != rune('V') {
												goto l436
											}
											position++
										}
									l437:
										{
											position439, tokenIndex439 := position, tokenIndex
											if buffer[position] != rune('e') {
												goto l440
											}
											position++
											goto l439
										l440:
											position, tokenIndex = position439, tokenIndex439
											if buffer[position] != rune('E') {
												goto l436
											}
											position++
										}
									l439:
										{
											position441, tokenIndex441 := position, tokenIndex
											if buffer[position] != rune('c') {
												goto l442
											}
											position++
											goto l441
										l442:
											position, tokenIndex = position441, tokenIndex441
											if buffer[position] != rune('C') {
												goto l436
											}
											position++
										}
									l441:
										{
											position443, tokenIndex443 := position, tokenIndex
											if buffer[position] != rune('d') {
												goto l444
											}
											position++
											goto l443
										l444:
											position, tokenIndex = position443, tokenIndex443
											if buffer[position] != rune('D') {
												goto l436
											}
											position++
										}
									l443:
										{
											position445, tokenIndex445 := position, tokenIndex
											if buffer[position] != rune('e') {
												goto l446
											}
											position++
											goto l445
										l446:
											position, tokenIndex = position445, tokenIndex445
											if buffer[position] != rune('E') {
												goto l436
											}
											position++
										}
									l445:
										{
											position447, tokenIndex447 := position, tokenIndex
											if buffer[position] != rune('l') {
												goto l448
											}
											position++
											goto l447
										l448:
											position, tokenIndex = position447, tokenIndex447
											if buffer[position] != rune('L') {
												goto l436
											}
											position++
										}
									l447:
										{
											position449, tokenIndex449 := position, tokenIndex
											if buffer[position] != rune('e') {
												goto l450
											}
											position++
											goto l449
										l450:
											position, tokenIndex = position449, tokenIndex449
											if buffer[position] != rune('E') {
												goto l436
											}
											position++
										}
									l449:
										{
											position451, tokenIndex451 := position, tokenIndex
											if buffer[position] != rune('t') {
												goto l452
											}
											position++
											goto l451
										l452:
											position, tokenIndex = position451, tokenIndex451
											if buffer[position] != rune('T') {
												goto l436
											}
											position++
										}
									l451:
										{
											position453, tokenIndex453 := position, tokenIndex
											if buffer[position] != rune('e') {
												goto l454
											}
											position++
											goto l453
										l454:
											position, tokenIndex = position453, tokenIndex453
											if buffer[position] != rune('E') {
												goto l436
											}
											position++
										}
									l453:
										goto l85
									l436:
										position, tokenIndex = position85, tokenIndex85
										{
											position456, tokenIndex456 := position, tokenIndex
											if buffer[position] != rune('x') {
												goto l457
											}
											position++
											goto l456
										l457:
											position, tokenIndex = position456, tokenIndex456
											if buffer[position] != rune('X') {
												goto l455
											}
											position++
										}
									l456:
										{
											position458, tokenIndex458 := position, tokenIndex
											if buffer[position] != rune('o') {
												goto l459
											}
											position++
											goto l458
										l459:
											position, tokenIndex = position458, tokenIndex458
											if buffer[position] != rune('O') {
												goto l455
											}
											position++
										}
									l458:
										{
											position460, tokenIndex460 := position, tokenIndex
											if buffer[position] != rune('r') {
												goto l461
											}
											position++
											goto l460
										l461:
											position, tokenIndex = position460, tokenIndex460
											if buffer[position] != rune('R') {
												goto l455
											}
											position++
										}
									l460:
										{
											position462, tokenIndex462 := position, tokenIndex
											if buffer[position] != rune('a') {
												goto l463
											}
											position++
											goto l462
										l463:
											position, tokenIndex = position462, tokenIndex462
											if buffer[position] != rune('A') {
												goto l455
											}
											position++
										}
									l462:
										{
											position464, tokenIndex464 := position, tokenIndex
											if buffer[position] != rune('s') {
												goto l465
											}
											position++
											goto l464
										l465:
											position, tokenIndex = position464, tokenIndex464
											if buffer[position] != rune('S') {
												goto l455
											}
											position++
										}
									l464:
										{
											position466, tokenIndex466 := position, tokenIndex
											if buffer[position] != rune('s') {
												goto l467
											}
											position++
											goto l466
										l467:
											position, tokenIndex = position466, tokenIndex466
											if buffer[position] != rune('S') {
												goto l455
											}
											position++
										}
									l466:
										{
											position468, tokenIndex468 := position, tokenIndex
											if buffer[position] != rune('i') {
												goto l469
											}
											position++
											goto l468
										l469:
											position, tokenIndex = position468, tokenIndex468
											if buffer[position] != rune('I') {
												goto l455
											}
											position++
										}
									l468:
										{
											position470, tokenIndex470 := position, tokenIndex
											if buffer[position] != rune('g') {
												goto l471
											}
											position++
											goto l470
										l471:
											position, tokenIndex = position470, tokenIndex470
											if buffer[position] != rune('G') {
												goto l455
											}
											position++
										}
									l470:
										{
											position472, tokenIndex472 := position, tokenIndex
											if buffer[position] != rune('n') {
												goto l473
											}
											position++
											goto l472
										l473:
											position, tokenIndex = position472, tokenIndex472
											if buffer[position] != rune('N') {
												goto l455
											}
											position++
										}
									l472:
										goto l85
									l455:
										position, tokenIndex = position85, tokenIndex85
										{
											switch buffer[position] {
											case 'X', 'x':
												{
													position475, tokenIndex475 := position, tokenIndex
													if buffer[position] != rune('x') {
														goto l476
													}
													position++
													goto l475
												l476:
													position, tokenIndex = position475, tokenIndex475
													if buffer[position] != rune('X') {
														goto l83
													}
													position++
												}
											l475:
												{
													position477, tokenIndex477 := position, tokenIndex
													if buffer[position] != rune('o') {
														goto l478
													}
													position++
													goto l477
												l478:
													position, tokenIndex = position477, tokenIndex477
													if buffer[position] != rune('O') {
														goto l83
													}
													position++
												}
											l477:
												{
													position479, tokenIndex479 := position, tokenIndex
													if buffer[position] != rune('r') {
														goto l480
													}
													position++
													goto l479
												l480:
													position, tokenIndex = position479, tokenIndex479
													if buffer[position] != rune('R') {
														goto l83
													}
													position++
												}
											l479:
												break
											case 'V', 'v':
												{
													position481, tokenIndex481 := position, tokenIndex
													if buffer[position] != rune('v') {
														goto l482
													}
													position++
													goto l481
												l482:
													position, tokenIndex = position481, tokenIndex481
													if buffer[position] != rune('V') {
														goto l83
													}
													position++
												}
											l481:
												{
													position483, tokenIndex483 := position, tokenIndex
													if buffer[position] != rune('e') {
														goto l484
													}
													position++
													goto l483
												l484:
													position, tokenIndex = position483, tokenIndex483
													if buffer[position] != rune('E') {
														goto l83
													}
													position++
												}
											l483:
												{
													position485, tokenIndex485 := position, tokenIndex
													if buffer[position] != rune('c') {
														goto l486
													}
													position++
													goto l485
												l486:
													position, tokenIndex = position485, tokenIndex485
													if buffer[position] != rune('C') {
														goto l83
													}
													position++
												}
											l485:
												{
													position487, tokenIndex487 := position, tokenIndex
													if buffer[position] != rune('n') {
														goto l488
													}
													position++
													goto l487
												l488:
													position, tokenIndex = position487, tokenIndex487
													if buffer[position] != rune('N') {
														goto l83
													}
													position++
												}
											l487:
												{
													position489, tokenIndex489 := position, tokenIndex
													if buffer[position] != rune('e') {
														goto l490
													}
													position++
													goto l489
												l490:
													position, tokenIndex = position489, tokenIndex489
													if buffer[position] != rune('E') {
														goto l83
													}
													position++
												}
											l489:
												{
													position491, tokenIndex491 := position, tokenIndex
													if buffer[position] != rune('w') {
														goto l492
													}
													position++
													goto l491
												l492:
													position, tokenIndex = position491, tokenIndex491
													if buffer[position] != rune('W') {
														goto l83
													}
													position++
												}
											l491:
												break
											case 'S', 's':
												{
													position493, tokenIndex493 := position, tokenIndex
													if buffer[position] != rune('s') {
														goto l494
													}
													position++
													goto l493
												l494:
													position, tokenIndex = position493, tokenIndex493
													if buffer[position] != rune('S') {
														goto l83
													}
													position++
												}
											l493:
												{
													position495, tokenIndex495 := position, tokenIndex
													if buffer[position] != rune('u') {
														goto l496
													}
													position++
													goto l495
												l496:
													position, tokenIndex = position495, tokenIndex495
													if buffer[position] != rune('U') {
														goto l83
													}
													position++
												}
											l495:
												{
													position497, tokenIndex497 := position, tokenIndex
													if buffer[position] != rune('b') {
														goto l498
													}
													position++
													goto l497
												l498:
													position, tokenIndex = position497, tokenIndex497
													if buffer[position] != rune('B') {
														goto l83
													}
													position++
												}
											l497:
												{
													position499, tokenIndex499 := position, tokenIndex
													if buffer[position] != rune('s') {
														goto l500
													}
													position++
													goto l499
												l500:
													position, tokenIndex = position499, tokenIndex499
													if buffer[position] != rune('S') {
														goto l83
													}
													position++
												}
											l499:
												break
											case 'R', 'r':
												{
													position501, tokenIndex501 := position, tokenIndex
													if buffer[position] != rune('r') {
														goto l502
													}
													position++
													goto l501
												l502:
													position, tokenIndex = position501, tokenIndex501
													if buffer[position] != rune('R') {
														goto l83
													}
													position++
												}
											l501:
												{
													position503, tokenIndex503 := position, tokenIndex
													if buffer[position] != rune('s') {
														goto l504
													}
													position++
													goto l503
												l504:
													position, tokenIndex = position503, tokenIndex503
													if buffer[position] != rune('S') {
														goto l83
													}
													position++
												}
											l503:
												{
													position505, tokenIndex505 := position, tokenIndex
													if buffer[position] != rune('h') {
														goto l506
													}
													position++
													goto l505
												l506:
													position, tokenIndex = position505, tokenIndex505
													if buffer[position] != rune('H') {
														goto l83
													}
													position++
												}
											l505:
												{
													position507, tokenIndex507 := position, tokenIndex
													if buffer[position] != rune('i') {
														goto l508
													}
													position++
													goto l507
												l508:
													position, tokenIndex = position507, tokenIndex507
													if buffer[position] != rune('I') {
														goto l83
													}
													position++
												}
											l507:
												{
													position509, tokenIndex509 := position, tokenIndex
													if buffer[position] != rune('f') {
														goto l510
													}
													position++
													goto l509
												l510:
													position, tokenIndex = position509, tokenIndex509
													if buffer[position] != rune('F') {
														goto l83
													}
													position++
												}
											l509:
												{
													position511, tokenIndex511 := position, tokenIndex
													if buffer[position] != rune('t') {
														goto l512
													}
													position++
													goto l511
												l512:
													position, tokenIndex = position511, tokenIndex511
													if buffer[position] != rune('T') {
														goto l83
													}
													position++
												}
											l511:
												break
											case 'P', 'p':
												{
													position513, tokenIndex513 := position, tokenIndex
													if buffer[position] != rune('p') {
														goto l514
													}
													position++
													goto l513
												l514:
													position, tokenIndex = position513, tokenIndex513
													if buffer[position] != rune('P') {
														goto l83
													}
													position++
												}
											l513:
												{
													position515, tokenIndex515 := position, tokenIndex
													if buffer[position] != rune('r') {
														goto l516
													}
													position++
													goto l515
												l516:
													position, tokenIndex = position515, tokenIndex515
													if buffer[position] != rune('R') {
														goto l83
													}
													position++
												}
											l515:
												{
													position517, tokenIndex517 := position, tokenIndex
													if buffer[position] != rune('e') {
														goto l518
													}
													position++
													goto l517
												l518:
													position, tokenIndex = position517, tokenIndex517
													if buffer[position] != rune('E') {
														goto l83
													}
													position++
												}
											l517:
												{
													position519, tokenIndex519 := position, tokenIndex
													if buffer[position] != rune('i') {
														goto l520
													}
													position++
													goto l519
												l520:
													position, tokenIndex = position519, tokenIndex519
													if buffer[position] != rune('I') {
														goto l83
													}
													position++
												}
											l519:
												{
													position521, tokenIndex521 := position, tokenIndex
													if buffer[position] != rune('n') {
														goto l522
													}
													position++
													goto l521
												l522:
													position, tokenIndex = position521, tokenIndex521
													if buffer[position] != rune('N') {
														goto l83
													}
													position++
												}
											l521:
												{
													position523, tokenIndex523 := position, tokenIndex
													if buffer[position] != rune('c') {
														goto l524
													}
													position++
													goto l523
												l524:
													position, tokenIndex = position523, tokenIndex523
													if buffer[position] != rune('C') {
														goto l83
													}
													position++
												}
											l523:
												break
											case 'O', 'o':
												{
													position525, tokenIndex525 := position, tokenIndex
													if buffer[position] != rune('o') {
														goto l526
													}
													position++
													goto l525
												l526:
													position, tokenIndex = position525, tokenIndex525
													if buffer[position] != rune('O') {
														goto l83
													}
													position++
												}
											l525:
												{
													position527, tokenIndex527 := position, tokenIndex
													if buffer[position] != rune('r') {
														goto l528
													}
													position++
													goto l527
												l528:
													position, tokenIndex = position527, tokenIndex527
													if buffer[position] != rune('R') {
														goto l83
													}
													position++
												}
											l527:
												break
											case 'N', 'n':
												{
													position529, tokenIndex529 := position, tokenIndex
													if buffer[position] != rune('n') {
														goto l530
													}
													position++
													goto l529
												l530:
													position, tokenIndex = position529, tokenIndex529
													if buffer[position] != rune('N') {
														goto l83
													}
													position++
												}
											l529:
												{
													position531, tokenIndex531 := position, tokenIndex
													if buffer[position] != rune('o') {
														goto l532
													}
													position++
													goto l531
												l532:
													position, tokenIndex = position531, tokenIndex531
													if buffer[position] != rune('O') {
														goto l83
													}
													position++
												}
											l531:
												{
													position533, tokenIndex533 := position, tokenIndex
													if buffer[position] != rune('t') {
														goto l534
													}
													position++
													goto l533
												l534:
													position, tokenIndex = position533, tokenIndex533
													if buffer[position] != rune('T') {
														goto l83
													}
													position++
												}
											l533:
												break
											case 'M', 'm':
												{
													position535, tokenIndex535 := position, tokenIndex
													if buffer[position] != rune('m') {
														goto l536
													}
													position++
													goto l535
												l536:
													position, tokenIndex = position535, tokenIndex535
													if buffer[position] != rune('M') {
														goto l83
													}
													position++
												}
											l535:
												{
													position537, tokenIndex537 := position, tokenIndex
													if buffer[position] != rune('u') {
														goto l538
													}
													position++
													goto l537
												l538:
													position, tokenIndex = position537, tokenIndex537
													if buffer[position] != rune('U') {
														goto l83
													}
													position++
												}
											l537:
												{
													position539, tokenIndex539 := position, tokenIndex
													if buffer[position] != rune('l') {
														goto l540
													}
													position++
													goto l539
												l540:
													position, tokenIndex = position539, tokenIndex539
													if buffer[position] != rune('L') {
														goto l83
													}
													position++
												}
											l539:
												{
													position541, tokenIndex541 := position, tokenIndex
													if buffer[position] != rune('t') {
														goto l542
													}
													position++
													goto l541
												l542:
													position, tokenIndex = position541, tokenIndex541
													if buffer[position] != rune('T') {
														goto l83
													}
													position++
												}
											l541:
												break
											case 'L', 'l':
												{
													position543, tokenIndex543 := position, tokenIndex
													if buffer[position] != rune('l') {
														goto l544
													}
													position++
													goto l543
												l544:
													position, tokenIndex = position543, tokenIndex543
													if buffer[position] != rune('L') {
														goto l83
													}
													position++
												}
											l543:
												{
													position545, tokenIndex545 := position, tokenIndex
													if buffer[position] != rune('t') {
														goto l546
													}
													position++
													goto l545
												l546:
													position, tokenIndex = position545, tokenIndex545
													if buffer[position] != rune('T') {
														goto l83
													}
													position++
												}
											l545:
												break
											case 'G', 'g':
												{
													position547, tokenIndex547 := position, tokenIndex
													if buffer[position] != rune('g') {
														goto l548
													}
													position++
													goto l547
												l548:
													position, tokenIndex = position547, tokenIndex547
													if buffer[position] != rune('G') {
														goto l83
													}
													position++
												}
											l547:
												{
													position549, tokenIndex549 := position, tokenIndex
													if buffer[position] != rune('t') {
														goto l550
													}
													position++
													goto l549
												l550:
													position, tokenIndex = position549, tokenIndex549
													if buffer[position] != rune('T') {
														goto l83
													}
													position++
												}
											l549:
												break
											case 'E', 'e':
												{
													position551, tokenIndex551 := position, tokenIndex
													if buffer[position] != rune('e') {
														goto l552
													}
													position++
													goto l551
												l552:
													position, tokenIndex = position551, tokenIndex551
													if buffer[position] != rune('E') {
														goto l83
													}
													position++
												}
											l551:
												{
													position553, tokenIndex553 := position, tokenIndex
													if buffer[position] != rune('q') {
														goto l554
													}
													position++
													goto l553
												l554:
													position, tokenIndex = position553, tokenIndex553
													if buffer[position] != rune('Q') {
														goto l83
													}
													position++
												}
											l553:
												break
											case 'D', 'd':
												{
													position555, tokenIndex555 := position, tokenIndex
													if buffer[position] != rune('d') {
														goto l556
													}
													position++
													goto l555
												l556:
													position, tokenIndex = position555, tokenIndex555
													if buffer[position] != rune('D') {
														goto l83
													}
													position++
												}
											l555:
												{
													position557, tokenIndex557 := position, tokenIndex
													if buffer[position] != rune('i') {
														goto l558
													}
													position++
													goto l557
												l558:
													position, tokenIndex = position557, tokenIndex557
													if buffer[position] != rune('I') {
														goto l83
													}
													position++
												}
											l557:
												{
													position559, tokenIndex559 := position, tokenIndex
													if buffer[position] != rune('v') {
														goto l560
													}
													position++
													goto l559
												l560:
													position, tokenIndex = position559, tokenIndex559
													if buffer[position] != rune('V') {
														goto l83
													}
													position++
												}
											l559:
												break
											case 'C', 'c':
												{
													position561, tokenIndex561 := position, tokenIndex
													if buffer[position] != rune('c') {
														goto l562
													}
													position++
													goto l561
												l562:
													position, tokenIndex = position561, tokenIndex561
													if buffer[position] != rune('C') {
														goto l83
													}
													position++
												}
											l561:
												{
													position563, tokenIndex563 := position, tokenIndex
													if buffer[position] != rune('a') {
														goto l564
													}
													position++
													goto l563
												l564:
													position, tokenIndex = position563, tokenIndex563
													if buffer[position] != rune('A') {
														goto l83
													}
													position++
												}
											l563:
												{
													position565, tokenIndex565 := position, tokenIndex
													if buffer[position] != rune('l') {
														goto l566
													}
													position++
													goto l565
												l566:
													position, tokenIndex = position565, tokenIndex565
													if buffer[position] != rune('L') {
														goto l83
													}
													position++
												}
											l565:
												{
													position567, tokenIndex567 := position, tokenIndex
													if buffer[position] != rune('l') {
														goto l568
													}
													position++
													goto l567
												l568:
													position, tokenIndex = position567, tokenIndex567
													if buffer[position] != rune('L') {
														goto l83
													}
													position++
												}
											l567:
												break
											case 'A', 'a':
												{
													position569, tokenIndex569 := position, tokenIndex
													if buffer[position] != rune('a') {
														goto l570
													}
													position++
													goto l569
												l570:
													position, tokenIndex = position569, tokenIndex569
													if buffer[position] != rune('A') {
														goto l83
													}
													position++
												}
											l569:
												{
													position571, tokenIndex571 := position, tokenIndex
													if buffer[position] != rune('s') {
														goto l572
													}
													position++
													goto l571
												l572:
													position, tokenIndex = position571, tokenIndex571
													if buffer[position] != rune('S') {
														goto l83
													}
													position++
												}
											l571:
												{
													position573, tokenIndex573 := position, tokenIndex
													if buffer[position] != rune('s') {
														goto l574
													}
													position++
													goto l573
												l574:
													position, tokenIndex = position573, tokenIndex573
													if buffer[position] != rune('S') {
														goto l83
													}
													position++
												}
											l573:
												{
													position575, tokenIndex575 := position, tokenIndex
													if buffer[position] != rune('i') {
														goto l576
													}
													position++
													goto l575
												l576:
													position, tokenIndex = position575, tokenIndex575
													if buffer[position] != rune('I') {
														goto l83
													}
													position++
												}
											l575:
												{
													position577, tokenIndex577 := position, tokenIndex
													if buffer[position] != rune('g') {
														goto l578
													}
													position++
													goto l577
												l578:
													position, tokenIndex = position577, tokenIndex577
													if buffer[position] != rune('G') {
														goto l83
													}
													position++
												}
											l577:
												{
													position579, tokenIndex579 := position, tokenIndex
													if buffer[position] != rune('n') {
														goto l580
													}
													position++
													goto l579
												l580:
													position, tokenIndex = position579, tokenIndex579
													if buffer[position] != rune('N') {
														goto l83
													}
													position++
												}
											l579:
												break
											default:
												break
											}
										}

									}
								l85:
									goto l84
								l83:
									position, tokenIndex = position83, tokenIndex83
								}
							l84:
								break
							case 'M', 'm':
								{
									position581, tokenIndex581 := position, tokenIndex
									if buffer[position] != rune('m') {
										goto l582
									}
									position++
									goto l581
								l582:
									position, tokenIndex = position581, tokenIndex581
									if buffer[position] != rune('M') {
										goto l0
									}
									position++
								}
							l581:
								{
									position583, tokenIndex583 := position, tokenIndex
									if buffer[position] != rune('e') {
										goto l584
									}
									position++
									goto l583
								l584:
									position, tokenIndex = position583, tokenIndex583
									if buffer[position] != rune('E') {
										goto l0
									}
									position++
								}
							l583:
								{
									position585, tokenIndex585 := position, tokenIndex
									if buffer[position] != rune('m') {
										goto l586
									}
									position++
									goto l585
								l586:
									position, tokenIndex = position585, tokenIndex585
									if buffer[position] != rune('M') {
										goto l0
									}
									position++
								}
							l585:
								{
									position587, tokenIndex587 := position, tokenIndex
									if buffer[position] != rune('b') {
										goto l588
									}
									position++
									goto l587
								l588:
									position, tokenIndex = position587, tokenIndex587
									if buffer[position] != rune('B') {
										goto l0
									}
									position++
								}
							l587:
								{
									position589, tokenIndex589 := position, tokenIndex
									if buffer[position] != rune('e') {
										goto l590
									}
									position++
									goto l589
								l590:
									position, tokenIndex = position589, tokenIndex589
									if buffer[position] != rune('E') {
										goto l0
									}
									position++
								}
							l589:
								{
									position591, tokenIndex591 := position, tokenIndex
									if buffer[position] != rune('r') {
										goto l592
									}
									position++
									goto l591
								l592:
									position, tokenIndex = position591, tokenIndex591
									if buffer[position] != rune('R') {
										goto l0
									}
									position++
								}
							l591:
								break
							default:
								{
									position593, tokenIndex593 := position, tokenIndex
									if buffer[position] != rune('a') {
										goto l594
									}
									position++
									goto l593
								l594:
									position, tokenIndex = position593, tokenIndex593
									if buffer[position] != rune('A') {
										goto l0
									}
									position++
								}
							l593:
								{
									position595, tokenIndex595 := position, tokenIndex
									if buffer[position] != rune('r') {
										goto l596
									}
									position++
									goto l595
								l596:
									position, tokenIndex = position595, tokenIndex595
									if buffer[position] != rune('R') {
										goto l0
									}
									position++
								}
							l595:
								{
									position597, tokenIndex597 := position, tokenIndex
									if buffer[position] != rune('t') {
										goto l598
									}
									position++
									goto l597
								l598:
									position, tokenIndex = position597, tokenIndex597
									if buffer[position] != rune('T') {
										goto l0
									}
									position++
								}
							l597:
								{
									position599, tokenIndex599 := position, tokenIndex
									if buffer[position] != rune('i') {
										goto l600
									}
									position++
									goto l599
								l600:
									position, tokenIndex = position599, tokenIndex599
									if buffer[position] != rune('I') {
										goto l0
									}
									position++
								}
							l599:
								{
									position601, tokenIndex601 := position, tokenIndex
									if buffer[position] != rune('f') {
										goto l602
									}
									position++
									goto l601
								l602:
									position, tokenIndex = position601, tokenIndex601
									if buffer[position] != rune('F') {
										goto l0
									}
									position++
								}
							l601:
								{
									position603, tokenIndex603 := position, tokenIndex
									if buffer[position] != rune('i') {
										goto l604
									}
									position++
									goto l603
								l604:
									position, tokenIndex = position603, tokenIndex603
									if buffer[position] != rune('I') {
										goto l0
									}
									position++
								}
							l603:
								{
									position605, tokenIndex605 := position, tokenIndex
									if buffer[position] != rune('c') {
										goto l606
									}
									position++
									goto l605
								l606:
									position, tokenIndex = position605, tokenIndex605
									if buffer[position] != rune('C') {
										goto l0
									}
									position++
								}
							l605:
								{
									position607, tokenIndex607 := position, tokenIndex
									if buffer[position] != rune('i') {
										goto l608
									}
									position++
									goto l607
								l608:
									position, tokenIndex = position607, tokenIndex607
									if buffer[position] != rune('I') {
										goto l0
									}
									position++
								}
							l607:
								{
									position609, tokenIndex609 := position, tokenIndex
									if buffer[position] != rune('a') {
										goto l610
									}
									position++
									goto l609
								l610:
									position, tokenIndex = position609, tokenIndex609
									if buffer[position] != rune('A') {
										goto l0
									}
									position++
								}
							l609:
								{
									position611, tokenIndex611 := position, tokenIndex
									if buffer[position] != rune('l') {
										goto l612
									}
									position++
									goto l611
								l612:
									position, tokenIndex = position611, tokenIndex611
									if buffer[position] != rune('L') {
										goto l0
									}
									position++
								}
							l611:
								break
							}
						}

					}
				l12:
					add(rulePegText, position11)
				}
				{
					add(ruleAction0, position)
				}
				add(ruleNoteAttr, position1)
			}
			return true
		l0:
			position, tokenIndex = position0, tokenIndex0
			return false
		},
		nil,
		nil,
		/* 4 Action0 <- <{
			getNode().AddNote(buffer[begin:end])
		}> */
		nil,
	}
	p.rules = _rules
}
