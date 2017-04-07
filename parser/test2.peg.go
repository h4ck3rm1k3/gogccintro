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
	ruleTUFILE
	ruleOpAttr
	ruleFileRef
	ruleSourceAttr
	ruleIntAttr
	ruleIntAttr3
	ruleIntAttr2
	ruleSignedIntAttr
	ruleAddr
	ruleHex
	ruleAddrAttr
	ruleTagAttr
	ruleBodyAttr
	ruleLinkAttr
	ruleNoteAttr
	ruleAccsAttr
	ruleQualAttr
	ruleSignAttr
	ruleNodeName
	ruleNodeAttr
	ruleSpecValue
	ruleLngtAttr
	ruleStringAttr
	ruleRandomSpec
	ruleOneAttr
	ruleAttrs
	ruleAttr
	ruleStatement
	ruleNode
	ruleInteger
	ruleNodeType
	ruleDecimalDigit
	ruleNonZeroDecimalDigit
	rulews
	ruleEOF
)

var rul3s = [...]string{
	"Unknown",
	"TUFILE",
	"OpAttr",
	"FileRef",
	"SourceAttr",
	"IntAttr",
	"IntAttr3",
	"IntAttr2",
	"SignedIntAttr",
	"Addr",
	"Hex",
	"AddrAttr",
	"TagAttr",
	"BodyAttr",
	"LinkAttr",
	"NoteAttr",
	"AccsAttr",
	"QualAttr",
	"SignAttr",
	"NodeName",
	"NodeAttr",
	"SpecValue",
	"LngtAttr",
	"StringAttr",
	"RandomSpec",
	"OneAttr",
	"Attrs",
	"Attr",
	"Statement",
	"Node",
	"Integer",
	"NodeType",
	"DecimalDigit",
	"NonZeroDecimalDigit",
	"ws",
	"EOF",
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
	rules  [36]func() bool
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

	matchDot := func() bool {
		if buffer[position] != endSymbol {
			position++
			return true
		}
		return false
	}

	/*matchChar := func(c byte) bool {
		if buffer[position] == c {
			position++
			return true
		}
		return false
	}*/

	/*matchRange := func(lower byte, upper byte) bool {
		if c := buffer[position]; c >= lower && c <= upper {
			position++
			return true
		}
		return false
	}*/

	_rules = [...]func() bool{
		nil,
		/* 0 TUFILE <- <(ws Statement+ EOF)> */
		func() bool {
			position0, tokenIndex0 := position, tokenIndex
			{
				position1 := position
				if !_rules[rulews]() {
					goto l0
				}
				{
					position4 := position
					if !_rules[ruleNode]() {
						goto l0
					}
					if !_rules[rulews]() {
						goto l0
					}
					{
						position5 := position
						{
							position8, tokenIndex8 := position, tokenIndex
							if c := buffer[position]; c < rune('a') || c > rune('z') {
								goto l9
							}
							position++
							goto l8
						l9:
							position, tokenIndex = position8, tokenIndex8
							if buffer[position] != rune('_') {
								goto l0
							}
							position++
						}
					l8:
					l6:
						{
							position7, tokenIndex7 := position, tokenIndex
							{
								position10, tokenIndex10 := position, tokenIndex
								if c := buffer[position]; c < rune('a') || c > rune('z') {
									goto l11
								}
								position++
								goto l10
							l11:
								position, tokenIndex = position10, tokenIndex10
								if buffer[position] != rune('_') {
									goto l7
								}
								position++
							}
						l10:
							goto l6
						l7:
							position, tokenIndex = position7, tokenIndex7
						}
						add(ruleNodeType, position5)
					}
					if !_rules[rulews]() {
						goto l0
					}
					{
						position12 := position
						{
							position13, tokenIndex13 := position, tokenIndex
							if !_rules[ruleOneAttr]() {
								goto l13
							}
							if !_rules[rulews]() {
								goto l13
							}
							{
								position15 := position
							l16:
								{
									position17, tokenIndex17 := position, tokenIndex
									if !_rules[rulews]() {
										goto l17
									}
									if !_rules[ruleOneAttr]() {
										goto l17
									}
									goto l16
								l17:
									position, tokenIndex = position17, tokenIndex17
								}
								add(ruleAttrs, position15)
							}
							if !_rules[rulews]() {
								goto l13
							}
							goto l14
						l13:
							position, tokenIndex = position13, tokenIndex13
						}
					l14:
						add(ruleAttr, position12)
					}
					if !_rules[rulews]() {
						goto l0
					}
					add(ruleStatement, position4)
				}
			l2:
				{
					position3, tokenIndex3 := position, tokenIndex
					{
						position18 := position
						if !_rules[ruleNode]() {
							goto l3
						}
						if !_rules[rulews]() {
							goto l3
						}
						{
							position19 := position
							{
								position22, tokenIndex22 := position, tokenIndex
								if c := buffer[position]; c < rune('a') || c > rune('z') {
									goto l23
								}
								position++
								goto l22
							l23:
								position, tokenIndex = position22, tokenIndex22
								if buffer[position] != rune('_') {
									goto l3
								}
								position++
							}
						l22:
						l20:
							{
								position21, tokenIndex21 := position, tokenIndex
								{
									position24, tokenIndex24 := position, tokenIndex
									if c := buffer[position]; c < rune('a') || c > rune('z') {
										goto l25
									}
									position++
									goto l24
								l25:
									position, tokenIndex = position24, tokenIndex24
									if buffer[position] != rune('_') {
										goto l21
									}
									position++
								}
							l24:
								goto l20
							l21:
								position, tokenIndex = position21, tokenIndex21
							}
							add(ruleNodeType, position19)
						}
						if !_rules[rulews]() {
							goto l3
						}
						{
							position26 := position
							{
								position27, tokenIndex27 := position, tokenIndex
								if !_rules[ruleOneAttr]() {
									goto l27
								}
								if !_rules[rulews]() {
									goto l27
								}
								{
									position29 := position
								l30:
									{
										position31, tokenIndex31 := position, tokenIndex
										if !_rules[rulews]() {
											goto l31
										}
										if !_rules[ruleOneAttr]() {
											goto l31
										}
										goto l30
									l31:
										position, tokenIndex = position31, tokenIndex31
									}
									add(ruleAttrs, position29)
								}
								if !_rules[rulews]() {
									goto l27
								}
								goto l28
							l27:
								position, tokenIndex = position27, tokenIndex27
							}
						l28:
							add(ruleAttr, position26)
						}
						if !_rules[rulews]() {
							goto l3
						}
						add(ruleStatement, position18)
					}
					goto l2
				l3:
					position, tokenIndex = position3, tokenIndex3
				}
				{
					position32 := position
					{
						position33, tokenIndex33 := position, tokenIndex
						if !matchDot() {
							goto l33
						}
						goto l0
					l33:
						position, tokenIndex = position33, tokenIndex33
					}
					add(ruleEOF, position32)
				}
				add(ruleTUFILE, position1)
			}
			return true
		l0:
			position, tokenIndex = position0, tokenIndex0
			return false
		},
		/* 1 OpAttr <- <(('o' / 'O') ('p' / 'P') ' ' [0-9])> */
		nil,
		/* 2 FileRef <- <((('_'? ((&('_') '_') | (&('0' | '1' | '2' | '3' | '4' | '5' | '6' | '7' | '8' | '9') [0-9]) | (&('a' | 'b' | 'c' | 'd' | 'e' | 'f' | 'g' | 'h' | 'i' | 'j' | 'k' | 'l' | 'm' | 'n' | 'o' | 'p' | 'q' | 'r' | 's' | 't' | 'u' | 'v' | 'w' | 'x' | 'y' | 'z') [a-z]) | (&('-') '-') | (&('A' | 'B' | 'C' | 'D' | 'E' | 'F' | 'G' | 'H' | 'I' | 'J' | 'K' | 'L' | 'M' | 'N' | 'O' | 'P' | 'Q' | 'R' | 'S' | 'T' | 'U' | 'V' | 'W' | 'X' | 'Y' | 'Z') [A-Z]))+ ('.' ((&('_') '_') | (&('a' | 'b' | 'c' | 'd' | 'e' | 'f' | 'g' | 'h' | 'i' | 'j' | 'k' | 'l' | 'm' | 'n' | 'o' | 'p' | 'q' | 'r' | 's' | 't' | 'u' | 'v' | 'w' | 'x' | 'y' | 'z') [a-z]) | (&('A' | 'B' | 'C' | 'D' | 'E' | 'F' | 'G' | 'H' | 'I' | 'J' | 'K' | 'L' | 'M' | 'N' | 'O' | 'P' | 'Q' | 'R' | 'S' | 'T' | 'U' | 'V' | 'W' | 'X' | 'Y' | 'Z') [A-Z]))+)*) / ('<' ('b' / 'B') ('u' / 'U') ('i' / 'I') ('l' / 'L') ('t' / 'T') '-' ('i' / 'I') ('n' / 'N') '>')) ':' Integer ws)> */
		nil,
		/* 3 SourceAttr <- <('s' 'r' 'c' 'p' ':' ws FileRef)> */
		nil,
		/* 4 IntAttr <- <(((&('P' | 'p') (('p' / 'P') ('r' / 'R') ('e' / 'E') ('c' / 'C'))) | (&('U' | 'u') (('u' / 'U') ('s' / 'S') ('e' / 'E') ('d' / 'D'))) | (&('L' | 'l') (('l' / 'L') ('o' / 'O') ('w' / 'W'))) | (&('B' | 'b') (('b' / 'B') ('a' / 'A') ('s' / 'S') ('e' / 'E') ('s' / 'S')))) ws ':' ws Integer ws)> */
		nil,
		/* 5 IntAttr3 <- <(((('l' / 'L') ('o' / 'O') ('w' / 'W')) / (('h' / 'H') ('i' / 'I') ('g' / 'G') ('h' / 'H'))) ws ':' ws '-'? Integer ws)> */
		nil,
		/* 6 IntAttr2 <- <(('a' / 'A') ('l' / 'L') ('g' / 'G') ('n' / 'N') ':' ws [0-9]+ ws)> */
		nil,
		/* 7 SignedIntAttr <- <(('i' / 'I') ('n' / 'N') ('t' / 'T') ws ':' ws ('-' / ('0' ('x' / 'X')))? (Integer / Hex)+ ws)> */
		nil,
		/* 8 Addr <- <(DecimalDigit / Hex)+> */
		nil,
		/* 9 Hex <- <[a-h]> */
		func() bool {
			position42, tokenIndex42 := position, tokenIndex
			{
				position43 := position
				if c := buffer[position]; c < rune('a') || c > rune('h') {
					goto l42
				}
				position++
				add(ruleHex, position43)
			}
			return true
		l42:
			position, tokenIndex = position42, tokenIndex42
			return false
		},
		/* 10 AddrAttr <- <(('a' / 'A') ('d' / 'D') ('d' / 'D') ('r' / 'R') ws ':' ws Addr)> */
		nil,
		/* 11 TagAttr <- <(('t' / 'T') ('a' / 'A') ('g' / 'G') ws ':' ws ((('s' / 'S') ('t' / 'T') ('r' / 'R') ('u' / 'U') ('c' / 'C') ('t' / 'T')) / (('u' / 'U') ('n' / 'N') ('i' / 'I') ('o' / 'O') ('n' / 'N'))))> */
		nil,
		/* 12 BodyAttr <- <(('b' / 'B') ('o' / 'O') ('d' / 'D') ('y' / 'Y') ws ':' ws ((('u' / 'U') ('n' / 'N') ('d' / 'D') ('e' / 'E') ('f' / 'F') ('i' / 'I') ('n' / 'N') ('e' / 'E') ('d' / 'D')) / Node))> */
		nil,
		/* 13 LinkAttr <- <(('l' / 'L') ('i' / 'I') ('n' / 'N') ('k' / 'K') ws ':' ws ((('e' / 'E') ('x' / 'X') ('t' / 'T') ('e' / 'E') ('r' / 'R') ('n' / 'N')) / (('s' / 'S') ('t' / 'T') ('a' / 'A') ('t' / 'T') ('i' / 'I') ('c' / 'C'))))> */
		nil,
		/* 14 NoteAttr <- <(('n' / 'N') ('o' / 'O') ('t' / 'T') ('e' / 'E') ws ':' ws ((('p' / 'P') ('t' / 'T') ('r' / 'R') ('m' / 'M') ('e' / 'E') ('m' / 'M')) / ((&('D' | 'd') (('d' / 'D') ('e' / 'E') ('s' / 'S') ('t' / 'T') ('r' / 'R') ('u' / 'U') ('c' / 'C') ('t' / 'T') ('o' / 'O') ('r' / 'R'))) | (&('P' | 'p') (('p' / 'P') ('s' / 'S') ('e' / 'E') ('u' / 'U') ('d' / 'D') ('o' / 'O') ' ' ('t' / 'T') ('m' / 'M') ('p' / 'P') ('l' / 'L'))) | (&('O' | 'o') (('o' / 'O') ('p' / 'P') ('e' / 'E') ('r' / 'R') ('a' / 'A') ('t' / 'T') ('o' / 'O') ('r' / 'R') ws ((('a' / 'A') ('n' / 'N') ('d' / 'D') ('a' / 'A') ('s' / 'S') ('s' / 'S') ('i' / 'I') ('g' / 'G') ('n' / 'N')) / (('a' / 'A') ('n' / 'N') ('d' / 'D')) / (('d' / 'D') ('e' / 'E') ('l' / 'L') ('e' / 'E') ('t' / 'T') ('e' / 'E')) / (('d' / 'D') ('e' / 'E') ('r' / 'R') ('e' / 'E') ('f' / 'F')) / (('d' / 'D') ('i' / 'I') ('v' / 'V') ('a' / 'A') ('s' / 'S') ('s' / 'S') ('i' / 'I') ('g' / 'G') ('n' / 'N')) / (('g' / 'G') ('e' / 'E')) / (('g' / 'G') ('e' / 'E')) / (('l' / 'L') ('a' / 'A') ('n' / 'N') ('d' / 'D')) / (('l' / 'L') ('e' / 'E')) / (('l' / 'L') ('n' / 'N') ('o' / 'O') ('t' / 'T')) / (('l' / 'L') ('o' / 'O') ('r' / 'R')) / (('l' / 'L') ('s' / 'S') ('h' / 'H') ('i' / 'I') ('f' / 'F') ('t' / 'T') ('a' / 'A') ('s' / 'S') ('s' / 'S') ('i' / 'I') ('g' / 'G') ('n' / 'N')) / (('l' / 'L') ('s' / 'S') ('h' / 'H') ('i' / 'I') ('f' / 'F') ('t' / 'T')) / (('m' / 'M') ('i' / 'I') ('n' / 'N') ('u' / 'U') ('s' / 'S') ('a' / 'A') ('s' / 'S') ('s' / 'S') ('i' / 'I') ('g' / 'G') ('n' / 'N')) / (('m' / 'M') ('i' / 'I') ('n' / 'N') ('u' / 'U') ('s' / 'S')) / (('m' / 'M') ('u' / 'U') ('l' / 'L') ('t' / 'T') ('a' / 'A') ('s' / 'S') ('s' / 'S') ('i' / 'I') ('g' / 'G') ('n' / 'N')) / (('n' / 'N') ('e' / 'E') ('g' / 'G')) / (('n' / 'N') ('e' / 'E') ('w' / 'W')) / (('n' / 'N') ('e' / 'E')) / (('o' / 'O') ('r' / 'R') ('a' / 'A') ('s' / 'S') ('s' / 'S') ('i' / 'I') ('g' / 'G') ('n' / 'N')) / (('p' / 'P') ('l' / 'L') ('u' / 'U') ('s' / 'S') ('a' / 'A') ('s' / 'S') ('s' / 'S') ('i' / 'I') ('g' / 'G') ('n' / 'N')) / (('p' / 'P') ('l' / 'L') ('u' / 'U') ('s' / 'S')) / (('p' / 'P') ('o' / 'O') ('s' / 'S') ('t' / 'T') ('d' / 'D') ('e' / 'E') ('c' / 'C')) / (('p' / 'P') ('o' / 'O') ('s' / 'S') ('t' / 'T') ('i' / 'I') ('n' / 'N') ('c' / 'C')) / (('p' / 'P') ('o' / 'O') ('s' / 'S')) / (('p' / 'P') ('r' / 'R') ('e' / 'E') ('d' / 'D') ('e' / 'E') ('c' / 'C')) / (('r' / 'R') ('e' / 'E') ('f' / 'F')) / (('r' / 'R') ('s' / 'S') ('h' / 'H') ('i' / 'I') ('f' / 'F') ('t' / 'T') ('a' / 'A') ('s' / 'S') ('s' / 'S') ('i' / 'I') ('g' / 'G') ('n' / 'N')) / (('v' / 'V') ('e' / 'E') ('c' / 'C') ('d' / 'D') ('e' / 'E') ('l' / 'L') ('e' / 'E') ('t' / 'T') ('e' / 'E')) / (('x' / 'X') ('o' / 'O') ('r' / 'R') ('a' / 'A') ('s' / 'S') ('s' / 'S') ('i' / 'I') ('g' / 'G') ('n' / 'N')) / ((&('X' | 'x') (('x' / 'X') ('o' / 'O') ('r' / 'R'))) | (&('V' | 'v') (('v' / 'V') ('e' / 'E') ('c' / 'C') ('n' / 'N') ('e' / 'E') ('w' / 'W'))) | (&('S' | 's') (('s' / 'S') ('u' / 'U') ('b' / 'B') ('s' / 'S'))) | (&('R' | 'r') (('r' / 'R') ('s' / 'S') ('h' / 'H') ('i' / 'I') ('f' / 'F') ('t' / 'T'))) | (&('P' | 'p') (('p' / 'P') ('r' / 'R') ('e' / 'E') ('i' / 'I') ('n' / 'N') ('c' / 'C'))) | (&('O' | 'o') (('o' / 'O') ('r' / 'R'))) | (&('N' | 'n') (('n' / 'N') ('o' / 'O') ('t' / 'T'))) | (&('M' | 'm') (('m' / 'M') ('u' / 'U') ('l' / 'L') ('t' / 'T'))) | (&('L' | 'l') (('l' / 'L') ('t' / 'T'))) | (&('G' | 'g') (('g' / 'G') ('t' / 'T'))) | (&('E' | 'e') (('e' / 'E') ('q' / 'Q'))) | (&('D' | 'd') (('d' / 'D') ('i' / 'I') ('v' / 'V'))) | (&('C' | 'c') (('c' / 'C') ('a' / 'A') ('l' / 'L') ('l' / 'L'))) | (&('A' | 'a') (('a' / 'A') ('s' / 'S') ('s' / 'S') ('i' / 'I') ('g' / 'G') ('n' / 'N'))) | (&() )))?)) | (&('M' | 'm') (('m' / 'M') ('e' / 'E') ('m' / 'M') ('b' / 'B') ('e' / 'E') ('r' / 'R'))) | (&('A' | 'a') (('a' / 'A') ('r' / 'R') ('t' / 'T') ('i' / 'I') ('f' / 'F') ('i' / 'I') ('c' / 'C') ('i' / 'I') ('a' / 'A') ('l' / 'L'))))))> */
		nil,
		/* 15 AccsAttr <- <(('a' / 'A') ('c' / 'C') ('c' / 'C') ('s' / 'S') ws ':' ws ((('p' / 'P') ('r' / 'R') ('i' / 'I') ('v' / 'V')) / (('p' / 'P') ('u' / 'U') ('b' / 'B')) / (('p' / 'P') ('r' / 'R') ('o' / 'O') ('t' / 'T'))))> */
		nil,
		/* 16 QualAttr <- <(('q' / 'Q') ('u' / 'U') ('a' / 'A') ('l' / 'L') ws ':' ws ((&('R' | 'r') ('r' / 'R')) | (&('C') 'C') | (&('c') 'c') | (&('V' | 'v') ('v' / 'V')))+ ws)> */
		nil,
		/* 17 SignAttr <- <(('s' / 'S') ('i' / 'I') ('g' / 'G') ('n' / 'N') ws ':' ws ((('s' / 'S') ('i' / 'I') ('g' / 'G') ('n' / 'N') ('e' / 'E') ('d' / 'D')) / (('u' / 'U') ('n' / 'N') ('s' / 'S') ('i' / 'I') ('g' / 'G') ('n' / 'N') ('e' / 'E') ('d' / 'D'))))> */
		nil,
		/* 18 NodeName <- <((('r' / 'R') ('e' / 'E') ('t' / 'T') ('n' / 'N')) / (('p' / 'P') ('r' / 'R') ('m' / 'M') ('s' / 'S')) / (('a' / 'A') ('r' / 'R') ('g' / 'G') ('s' / 'S')) / (('b' / 'B') ('a' / 'A') ('s' / 'S') ('e' / 'E')) / (('o' / 'O') ('r' / 'R') ('i' / 'I') ('g' / 'G')) / (('o' / 'O') ('p' / 'P') '0') / (('d' / 'D') ('e' / 'E') ('c' / 'C') ('l' / 'L')) / (('d' / 'D') ('c' / 'C') ('l' / 'L') ('s' / 'S')) / (('e' / 'E') ('x' / 'X') ('p' / 'P') ('r' / 'R')) / (('f' / 'F') ('n' / 'N') ('c' / 'C') ('s' / 'S')) / (('f' / 'F') ('n' / 'N')) / (('c' / 'C') ('n' / 'N') ('s' / 'S') ('t' / 'T')) / (('v' / 'V') ('a' / 'A') ('r' / 'R') ('s' / 'S')) / (('v' / 'V') ('f' / 'F') ('l' / 'L') ('d' / 'D')) / (('v' / 'V') ('a' / 'A') ('l' / 'L') ('u' / 'U')) / (('c' / 'C') ('h' / 'H') ('a' / 'A') ('n' / 'N')) / (('p' / 'P') ('u' / 'U') ('r' / 'R') ('p' / 'P')) / (('c' / 'C') ('l' / 'L') ('s' / 'S')) / (('c' / 'C') ('s' / 'S') ('t' / 'T') ('s' / 'S')) / (('t' / 'T') ('y' / 'Y') ('p' / 'P') ('e' / 'E')) / (('m' / 'M') ('i' / 'I') ('n' / 'N')) / (('m' / 'M') ('a' / 'A') ('x' / 'X')) / (('b' / 'B') ('i' / 'I') ('n' / 'N') ('f' / 'F')) / (('s' / 'S') ('c' / 'C') ('p' / 'P') ('e' / 'E')) / (('i' / 'I') ('n' / 'N') ('i' / 'I') ('t' / 'T')) / ((&('O' | 'o') OpAttr) | (&('V' | 'v') (('v' / 'V') ('a' / 'A') ('l' / 'L'))) | (&('I' | 'i') (('i' / 'I') ('d' / 'D') ('x' / 'X'))) | (&('S' | 's') (('s' / 'S') ('i' / 'I') ('z' / 'Z') ('e' / 'E'))) | (&('B' | 'b') (('b' / 'B') ('p' / 'P') ('o' / 'O') ('s' / 'S'))) | (&('C' | 'c') (('c' / 'C') ('h' / 'H') ('a' / 'A') ('i' / 'I') ('n' / 'N'))) | (&('E' | 'e') (('e' / 'E') ('l' / 'L') ('t' / 'T') ('s' / 'S'))) | (&('D' | 'd') (('d' / 'D') ('o' / 'O') ('m' / 'M') ('n' / 'N'))) | (&('T' | 't') (('t' / 'T') ('y' / 'Y') ('p' / 'P') ('e' / 'E'))) | (&('M' | 'm') (('m' / 'M') ('n' / 'N') ('g' / 'G') ('l' / 'L'))) | (&('N' | 'n') (('n' / 'N') ('a' / 'A') ('m' / 'M') ('e' / 'E'))) | (&('U' | 'u') (('u' / 'U') ('n' / 'N') ('q' / 'Q') ('l' / 'L'))) | (&('L' | 'l') (('l' / 'L') ('a' / 'A') ('b' / 'B') ('l' / 'L'))) | (&('P' | 'p') (('p' / 'P') ('t' / 'T') ('d' / 'D'))) | (&('F' | 'f') (('f' / 'F') ('l' / 'L') ('d' / 'D') ('s' / 'S'))) | (&('R' | 'r') (('r' / 'R') ('e' / 'E') ('f' / 'F') ('d' / 'D'))) | (&('A' | 'a') (('a' / 'A') ('r' / 'R') ('g' / 'G') ('t' / 'T'))) | (&('0' | '1' | '2' | '3' | '4' | '5' | '6' | '7' | '8' | '9') Integer)))> */
		nil,
		/* 19 NodeAttr <- <(ws NodeName ws ':' ws Node ws)> */
		nil,
		/* 20 SpecValue <- <(ws (('s' / 'S') ('p' / 'P') ('e' / 'E') ('c' / 'C') ':')? ws (((&('R' | 'r') (('r' / 'R') ('e' / 'E') ('g' / 'G') ('i' / 'I') ('s' / 'S') ('t' / 'T') ('e' / 'E') ('r' / 'R'))) | (&('V' | 'v') (('v' / 'V') ('i' / 'I') ('r' / 'R') ('t' / 'T'))) | (&('P' | 'p') (('p' / 'P') ('u' / 'U') ('r' / 'R') ('e' / 'E'))) | (&('B' | 'b') (('b' / 'B') ('i' / 'I') ('t' / 'T') ('f' / 'F') ('i' / 'I') ('e' / 'E') ('l' / 'L') ('d' / 'D'))) | (&('M' | 'm') (('m' / 'M') ('u' / 'U') ('t' / 'T') ('a' / 'A') ('b' / 'B') ('l' / 'L') ('e' / 'E')))) ws)+)> */
		nil,
		/* 21 LngtAttr <- <(('l' / 'L') ('n' / 'N') ('g' / 'G') ('t' / 'T') ws ':' ws Integer ws)> */
		nil,
		/* 22 StringAttr <- <(('s' / 'S') ('t' / 'T') ('r' / 'R') ('g' / 'G') ':' .+)> */
		nil,
		/* 23 RandomSpec <- <(ws (('s' / 'S') ('p' / 'P') ('e' / 'E') ('c' / 'C')) ws)> */
		nil,
		/* 24 OneAttr <- <(AddrAttr / SpecValue / NodeAttr / SourceAttr / IntAttr / StringAttr / SignAttr / IntAttr3 / LngtAttr / AccsAttr / ((&('A' | 'a') IntAttr2) | (&('Q' | 'q') QualAttr) | (&('L' | 'l') LinkAttr) | (&('N' | 'n') NoteAttr) | (&('B' | 'b') BodyAttr) | (&('T' | 't') TagAttr) | (&('I' | 'i') SignedIntAttr) | (&('\t' | '\n' | '\r' | ' ' | 'S' | 's') RandomSpec)))> */
		func() bool {
			position58, tokenIndex58 := position, tokenIndex
			{
				position59 := position
				{
					position60, tokenIndex60 := position, tokenIndex
					{
						position62 := position
						{
							position63, tokenIndex63 := position, tokenIndex
							if buffer[position] != rune('a') {
								goto l64
							}
							position++
							goto l63
						l64:
							position, tokenIndex = position63, tokenIndex63
							if buffer[position] != rune('A') {
								goto l61
							}
							position++
						}
					l63:
						{
							position65, tokenIndex65 := position, tokenIndex
							if buffer[position] != rune('d') {
								goto l66
							}
							position++
							goto l65
						l66:
							position, tokenIndex = position65, tokenIndex65
							if buffer[position] != rune('D') {
								goto l61
							}
							position++
						}
					l65:
						{
							position67, tokenIndex67 := position, tokenIndex
							if buffer[position] != rune('d') {
								goto l68
							}
							position++
							goto l67
						l68:
							position, tokenIndex = position67, tokenIndex67
							if buffer[position] != rune('D') {
								goto l61
							}
							position++
						}
					l67:
						{
							position69, tokenIndex69 := position, tokenIndex
							if buffer[position] != rune('r') {
								goto l70
							}
							position++
							goto l69
						l70:
							position, tokenIndex = position69, tokenIndex69
							if buffer[position] != rune('R') {
								goto l61
							}
							position++
						}
					l69:
						if !_rules[rulews]() {
							goto l61
						}
						if buffer[position] != rune(':') {
							goto l61
						}
						position++
						if !_rules[rulews]() {
							goto l61
						}
						{
							position71 := position
							{
								position74, tokenIndex74 := position, tokenIndex
								if !_rules[ruleDecimalDigit]() {
									goto l75
								}
								goto l74
							l75:
								position, tokenIndex = position74, tokenIndex74
								if !_rules[ruleHex]() {
									goto l61
								}
							}
						l74:
						l72:
							{
								position73, tokenIndex73 := position, tokenIndex
								{
									position76, tokenIndex76 := position, tokenIndex
									if !_rules[ruleDecimalDigit]() {
										goto l77
									}
									goto l76
								l77:
									position, tokenIndex = position76, tokenIndex76
									if !_rules[ruleHex]() {
										goto l73
									}
								}
							l76:
								goto l72
							l73:
								position, tokenIndex = position73, tokenIndex73
							}
							add(ruleAddr, position71)
						}
						add(ruleAddrAttr, position62)
					}
					goto l60
				l61:
					position, tokenIndex = position60, tokenIndex60
					{
						position79 := position
						if !_rules[rulews]() {
							goto l78
						}
						{
							position80, tokenIndex80 := position, tokenIndex
							{
								position82, tokenIndex82 := position, tokenIndex
								if buffer[position] != rune('s') {
									goto l83
								}
								position++
								goto l82
							l83:
								position, tokenIndex = position82, tokenIndex82
								if buffer[position] != rune('S') {
									goto l80
								}
								position++
							}
						l82:
							{
								position84, tokenIndex84 := position, tokenIndex
								if buffer[position] != rune('p') {
									goto l85
								}
								position++
								goto l84
							l85:
								position, tokenIndex = position84, tokenIndex84
								if buffer[position] != rune('P') {
									goto l80
								}
								position++
							}
						l84:
							{
								position86, tokenIndex86 := position, tokenIndex
								if buffer[position] != rune('e') {
									goto l87
								}
								position++
								goto l86
							l87:
								position, tokenIndex = position86, tokenIndex86
								if buffer[position] != rune('E') {
									goto l80
								}
								position++
							}
						l86:
							{
								position88, tokenIndex88 := position, tokenIndex
								if buffer[position] != rune('c') {
									goto l89
								}
								position++
								goto l88
							l89:
								position, tokenIndex = position88, tokenIndex88
								if buffer[position] != rune('C') {
									goto l80
								}
								position++
							}
						l88:
							if buffer[position] != rune(':') {
								goto l80
							}
							position++
							goto l81
						l80:
							position, tokenIndex = position80, tokenIndex80
						}
					l81:
						if !_rules[rulews]() {
							goto l78
						}
						{
							switch buffer[position] {
							case 'R', 'r':
								{
									position93, tokenIndex93 := position, tokenIndex
									if buffer[position] != rune('r') {
										goto l94
									}
									position++
									goto l93
								l94:
									position, tokenIndex = position93, tokenIndex93
									if buffer[position] != rune('R') {
										goto l78
									}
									position++
								}
							l93:
								{
									position95, tokenIndex95 := position, tokenIndex
									if buffer[position] != rune('e') {
										goto l96
									}
									position++
									goto l95
								l96:
									position, tokenIndex = position95, tokenIndex95
									if buffer[position] != rune('E') {
										goto l78
									}
									position++
								}
							l95:
								{
									position97, tokenIndex97 := position, tokenIndex
									if buffer[position] != rune('g') {
										goto l98
									}
									position++
									goto l97
								l98:
									position, tokenIndex = position97, tokenIndex97
									if buffer[position] != rune('G') {
										goto l78
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
										goto l78
									}
									position++
								}
							l99:
								{
									position101, tokenIndex101 := position, tokenIndex
									if buffer[position] != rune('s') {
										goto l102
									}
									position++
									goto l101
								l102:
									position, tokenIndex = position101, tokenIndex101
									if buffer[position] != rune('S') {
										goto l78
									}
									position++
								}
							l101:
								{
									position103, tokenIndex103 := position, tokenIndex
									if buffer[position] != rune('t') {
										goto l104
									}
									position++
									goto l103
								l104:
									position, tokenIndex = position103, tokenIndex103
									if buffer[position] != rune('T') {
										goto l78
									}
									position++
								}
							l103:
								{
									position105, tokenIndex105 := position, tokenIndex
									if buffer[position] != rune('e') {
										goto l106
									}
									position++
									goto l105
								l106:
									position, tokenIndex = position105, tokenIndex105
									if buffer[position] != rune('E') {
										goto l78
									}
									position++
								}
							l105:
								{
									position107, tokenIndex107 := position, tokenIndex
									if buffer[position] != rune('r') {
										goto l108
									}
									position++
									goto l107
								l108:
									position, tokenIndex = position107, tokenIndex107
									if buffer[position] != rune('R') {
										goto l78
									}
									position++
								}
							l107:
								break
							case 'V', 'v':
								{
									position109, tokenIndex109 := position, tokenIndex
									if buffer[position] != rune('v') {
										goto l110
									}
									position++
									goto l109
								l110:
									position, tokenIndex = position109, tokenIndex109
									if buffer[position] != rune('V') {
										goto l78
									}
									position++
								}
							l109:
								{
									position111, tokenIndex111 := position, tokenIndex
									if buffer[position] != rune('i') {
										goto l112
									}
									position++
									goto l111
								l112:
									position, tokenIndex = position111, tokenIndex111
									if buffer[position] != rune('I') {
										goto l78
									}
									position++
								}
							l111:
								{
									position113, tokenIndex113 := position, tokenIndex
									if buffer[position] != rune('r') {
										goto l114
									}
									position++
									goto l113
								l114:
									position, tokenIndex = position113, tokenIndex113
									if buffer[position] != rune('R') {
										goto l78
									}
									position++
								}
							l113:
								{
									position115, tokenIndex115 := position, tokenIndex
									if buffer[position] != rune('t') {
										goto l116
									}
									position++
									goto l115
								l116:
									position, tokenIndex = position115, tokenIndex115
									if buffer[position] != rune('T') {
										goto l78
									}
									position++
								}
							l115:
								break
							case 'P', 'p':
								{
									position117, tokenIndex117 := position, tokenIndex
									if buffer[position] != rune('p') {
										goto l118
									}
									position++
									goto l117
								l118:
									position, tokenIndex = position117, tokenIndex117
									if buffer[position] != rune('P') {
										goto l78
									}
									position++
								}
							l117:
								{
									position119, tokenIndex119 := position, tokenIndex
									if buffer[position] != rune('u') {
										goto l120
									}
									position++
									goto l119
								l120:
									position, tokenIndex = position119, tokenIndex119
									if buffer[position] != rune('U') {
										goto l78
									}
									position++
								}
							l119:
								{
									position121, tokenIndex121 := position, tokenIndex
									if buffer[position] != rune('r') {
										goto l122
									}
									position++
									goto l121
								l122:
									position, tokenIndex = position121, tokenIndex121
									if buffer[position] != rune('R') {
										goto l78
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
										goto l78
									}
									position++
								}
							l123:
								break
							case 'B', 'b':
								{
									position125, tokenIndex125 := position, tokenIndex
									if buffer[position] != rune('b') {
										goto l126
									}
									position++
									goto l125
								l126:
									position, tokenIndex = position125, tokenIndex125
									if buffer[position] != rune('B') {
										goto l78
									}
									position++
								}
							l125:
								{
									position127, tokenIndex127 := position, tokenIndex
									if buffer[position] != rune('i') {
										goto l128
									}
									position++
									goto l127
								l128:
									position, tokenIndex = position127, tokenIndex127
									if buffer[position] != rune('I') {
										goto l78
									}
									position++
								}
							l127:
								{
									position129, tokenIndex129 := position, tokenIndex
									if buffer[position] != rune('t') {
										goto l130
									}
									position++
									goto l129
								l130:
									position, tokenIndex = position129, tokenIndex129
									if buffer[position] != rune('T') {
										goto l78
									}
									position++
								}
							l129:
								{
									position131, tokenIndex131 := position, tokenIndex
									if buffer[position] != rune('f') {
										goto l132
									}
									position++
									goto l131
								l132:
									position, tokenIndex = position131, tokenIndex131
									if buffer[position] != rune('F') {
										goto l78
									}
									position++
								}
							l131:
								{
									position133, tokenIndex133 := position, tokenIndex
									if buffer[position] != rune('i') {
										goto l134
									}
									position++
									goto l133
								l134:
									position, tokenIndex = position133, tokenIndex133
									if buffer[position] != rune('I') {
										goto l78
									}
									position++
								}
							l133:
								{
									position135, tokenIndex135 := position, tokenIndex
									if buffer[position] != rune('e') {
										goto l136
									}
									position++
									goto l135
								l136:
									position, tokenIndex = position135, tokenIndex135
									if buffer[position] != rune('E') {
										goto l78
									}
									position++
								}
							l135:
								{
									position137, tokenIndex137 := position, tokenIndex
									if buffer[position] != rune('l') {
										goto l138
									}
									position++
									goto l137
								l138:
									position, tokenIndex = position137, tokenIndex137
									if buffer[position] != rune('L') {
										goto l78
									}
									position++
								}
							l137:
								{
									position139, tokenIndex139 := position, tokenIndex
									if buffer[position] != rune('d') {
										goto l140
									}
									position++
									goto l139
								l140:
									position, tokenIndex = position139, tokenIndex139
									if buffer[position] != rune('D') {
										goto l78
									}
									position++
								}
							l139:
								break
							default:
								{
									position141, tokenIndex141 := position, tokenIndex
									if buffer[position] != rune('m') {
										goto l142
									}
									position++
									goto l141
								l142:
									position, tokenIndex = position141, tokenIndex141
									if buffer[position] != rune('M') {
										goto l78
									}
									position++
								}
							l141:
								{
									position143, tokenIndex143 := position, tokenIndex
									if buffer[position] != rune('u') {
										goto l144
									}
									position++
									goto l143
								l144:
									position, tokenIndex = position143, tokenIndex143
									if buffer[position] != rune('U') {
										goto l78
									}
									position++
								}
							l143:
								{
									position145, tokenIndex145 := position, tokenIndex
									if buffer[position] != rune('t') {
										goto l146
									}
									position++
									goto l145
								l146:
									position, tokenIndex = position145, tokenIndex145
									if buffer[position] != rune('T') {
										goto l78
									}
									position++
								}
							l145:
								{
									position147, tokenIndex147 := position, tokenIndex
									if buffer[position] != rune('a') {
										goto l148
									}
									position++
									goto l147
								l148:
									position, tokenIndex = position147, tokenIndex147
									if buffer[position] != rune('A') {
										goto l78
									}
									position++
								}
							l147:
								{
									position149, tokenIndex149 := position, tokenIndex
									if buffer[position] != rune('b') {
										goto l150
									}
									position++
									goto l149
								l150:
									position, tokenIndex = position149, tokenIndex149
									if buffer[position] != rune('B') {
										goto l78
									}
									position++
								}
							l149:
								{
									position151, tokenIndex151 := position, tokenIndex
									if buffer[position] != rune('l') {
										goto l152
									}
									position++
									goto l151
								l152:
									position, tokenIndex = position151, tokenIndex151
									if buffer[position] != rune('L') {
										goto l78
									}
									position++
								}
							l151:
								{
									position153, tokenIndex153 := position, tokenIndex
									if buffer[position] != rune('e') {
										goto l154
									}
									position++
									goto l153
								l154:
									position, tokenIndex = position153, tokenIndex153
									if buffer[position] != rune('E') {
										goto l78
									}
									position++
								}
							l153:
								break
							}
						}

						if !_rules[rulews]() {
							goto l78
						}
					l90:
						{
							position91, tokenIndex91 := position, tokenIndex
							{
								switch buffer[position] {
								case 'R', 'r':
									{
										position156, tokenIndex156 := position, tokenIndex
										if buffer[position] != rune('r') {
											goto l157
										}
										position++
										goto l156
									l157:
										position, tokenIndex = position156, tokenIndex156
										if buffer[position] != rune('R') {
											goto l91
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
											goto l91
										}
										position++
									}
								l158:
									{
										position160, tokenIndex160 := position, tokenIndex
										if buffer[position] != rune('g') {
											goto l161
										}
										position++
										goto l160
									l161:
										position, tokenIndex = position160, tokenIndex160
										if buffer[position] != rune('G') {
											goto l91
										}
										position++
									}
								l160:
									{
										position162, tokenIndex162 := position, tokenIndex
										if buffer[position] != rune('i') {
											goto l163
										}
										position++
										goto l162
									l163:
										position, tokenIndex = position162, tokenIndex162
										if buffer[position] != rune('I') {
											goto l91
										}
										position++
									}
								l162:
									{
										position164, tokenIndex164 := position, tokenIndex
										if buffer[position] != rune('s') {
											goto l165
										}
										position++
										goto l164
									l165:
										position, tokenIndex = position164, tokenIndex164
										if buffer[position] != rune('S') {
											goto l91
										}
										position++
									}
								l164:
									{
										position166, tokenIndex166 := position, tokenIndex
										if buffer[position] != rune('t') {
											goto l167
										}
										position++
										goto l166
									l167:
										position, tokenIndex = position166, tokenIndex166
										if buffer[position] != rune('T') {
											goto l91
										}
										position++
									}
								l166:
									{
										position168, tokenIndex168 := position, tokenIndex
										if buffer[position] != rune('e') {
											goto l169
										}
										position++
										goto l168
									l169:
										position, tokenIndex = position168, tokenIndex168
										if buffer[position] != rune('E') {
											goto l91
										}
										position++
									}
								l168:
									{
										position170, tokenIndex170 := position, tokenIndex
										if buffer[position] != rune('r') {
											goto l171
										}
										position++
										goto l170
									l171:
										position, tokenIndex = position170, tokenIndex170
										if buffer[position] != rune('R') {
											goto l91
										}
										position++
									}
								l170:
									break
								case 'V', 'v':
									{
										position172, tokenIndex172 := position, tokenIndex
										if buffer[position] != rune('v') {
											goto l173
										}
										position++
										goto l172
									l173:
										position, tokenIndex = position172, tokenIndex172
										if buffer[position] != rune('V') {
											goto l91
										}
										position++
									}
								l172:
									{
										position174, tokenIndex174 := position, tokenIndex
										if buffer[position] != rune('i') {
											goto l175
										}
										position++
										goto l174
									l175:
										position, tokenIndex = position174, tokenIndex174
										if buffer[position] != rune('I') {
											goto l91
										}
										position++
									}
								l174:
									{
										position176, tokenIndex176 := position, tokenIndex
										if buffer[position] != rune('r') {
											goto l177
										}
										position++
										goto l176
									l177:
										position, tokenIndex = position176, tokenIndex176
										if buffer[position] != rune('R') {
											goto l91
										}
										position++
									}
								l176:
									{
										position178, tokenIndex178 := position, tokenIndex
										if buffer[position] != rune('t') {
											goto l179
										}
										position++
										goto l178
									l179:
										position, tokenIndex = position178, tokenIndex178
										if buffer[position] != rune('T') {
											goto l91
										}
										position++
									}
								l178:
									break
								case 'P', 'p':
									{
										position180, tokenIndex180 := position, tokenIndex
										if buffer[position] != rune('p') {
											goto l181
										}
										position++
										goto l180
									l181:
										position, tokenIndex = position180, tokenIndex180
										if buffer[position] != rune('P') {
											goto l91
										}
										position++
									}
								l180:
									{
										position182, tokenIndex182 := position, tokenIndex
										if buffer[position] != rune('u') {
											goto l183
										}
										position++
										goto l182
									l183:
										position, tokenIndex = position182, tokenIndex182
										if buffer[position] != rune('U') {
											goto l91
										}
										position++
									}
								l182:
									{
										position184, tokenIndex184 := position, tokenIndex
										if buffer[position] != rune('r') {
											goto l185
										}
										position++
										goto l184
									l185:
										position, tokenIndex = position184, tokenIndex184
										if buffer[position] != rune('R') {
											goto l91
										}
										position++
									}
								l184:
									{
										position186, tokenIndex186 := position, tokenIndex
										if buffer[position] != rune('e') {
											goto l187
										}
										position++
										goto l186
									l187:
										position, tokenIndex = position186, tokenIndex186
										if buffer[position] != rune('E') {
											goto l91
										}
										position++
									}
								l186:
									break
								case 'B', 'b':
									{
										position188, tokenIndex188 := position, tokenIndex
										if buffer[position] != rune('b') {
											goto l189
										}
										position++
										goto l188
									l189:
										position, tokenIndex = position188, tokenIndex188
										if buffer[position] != rune('B') {
											goto l91
										}
										position++
									}
								l188:
									{
										position190, tokenIndex190 := position, tokenIndex
										if buffer[position] != rune('i') {
											goto l191
										}
										position++
										goto l190
									l191:
										position, tokenIndex = position190, tokenIndex190
										if buffer[position] != rune('I') {
											goto l91
										}
										position++
									}
								l190:
									{
										position192, tokenIndex192 := position, tokenIndex
										if buffer[position] != rune('t') {
											goto l193
										}
										position++
										goto l192
									l193:
										position, tokenIndex = position192, tokenIndex192
										if buffer[position] != rune('T') {
											goto l91
										}
										position++
									}
								l192:
									{
										position194, tokenIndex194 := position, tokenIndex
										if buffer[position] != rune('f') {
											goto l195
										}
										position++
										goto l194
									l195:
										position, tokenIndex = position194, tokenIndex194
										if buffer[position] != rune('F') {
											goto l91
										}
										position++
									}
								l194:
									{
										position196, tokenIndex196 := position, tokenIndex
										if buffer[position] != rune('i') {
											goto l197
										}
										position++
										goto l196
									l197:
										position, tokenIndex = position196, tokenIndex196
										if buffer[position] != rune('I') {
											goto l91
										}
										position++
									}
								l196:
									{
										position198, tokenIndex198 := position, tokenIndex
										if buffer[position] != rune('e') {
											goto l199
										}
										position++
										goto l198
									l199:
										position, tokenIndex = position198, tokenIndex198
										if buffer[position] != rune('E') {
											goto l91
										}
										position++
									}
								l198:
									{
										position200, tokenIndex200 := position, tokenIndex
										if buffer[position] != rune('l') {
											goto l201
										}
										position++
										goto l200
									l201:
										position, tokenIndex = position200, tokenIndex200
										if buffer[position] != rune('L') {
											goto l91
										}
										position++
									}
								l200:
									{
										position202, tokenIndex202 := position, tokenIndex
										if buffer[position] != rune('d') {
											goto l203
										}
										position++
										goto l202
									l203:
										position, tokenIndex = position202, tokenIndex202
										if buffer[position] != rune('D') {
											goto l91
										}
										position++
									}
								l202:
									break
								default:
									{
										position204, tokenIndex204 := position, tokenIndex
										if buffer[position] != rune('m') {
											goto l205
										}
										position++
										goto l204
									l205:
										position, tokenIndex = position204, tokenIndex204
										if buffer[position] != rune('M') {
											goto l91
										}
										position++
									}
								l204:
									{
										position206, tokenIndex206 := position, tokenIndex
										if buffer[position] != rune('u') {
											goto l207
										}
										position++
										goto l206
									l207:
										position, tokenIndex = position206, tokenIndex206
										if buffer[position] != rune('U') {
											goto l91
										}
										position++
									}
								l206:
									{
										position208, tokenIndex208 := position, tokenIndex
										if buffer[position] != rune('t') {
											goto l209
										}
										position++
										goto l208
									l209:
										position, tokenIndex = position208, tokenIndex208
										if buffer[position] != rune('T') {
											goto l91
										}
										position++
									}
								l208:
									{
										position210, tokenIndex210 := position, tokenIndex
										if buffer[position] != rune('a') {
											goto l211
										}
										position++
										goto l210
									l211:
										position, tokenIndex = position210, tokenIndex210
										if buffer[position] != rune('A') {
											goto l91
										}
										position++
									}
								l210:
									{
										position212, tokenIndex212 := position, tokenIndex
										if buffer[position] != rune('b') {
											goto l213
										}
										position++
										goto l212
									l213:
										position, tokenIndex = position212, tokenIndex212
										if buffer[position] != rune('B') {
											goto l91
										}
										position++
									}
								l212:
									{
										position214, tokenIndex214 := position, tokenIndex
										if buffer[position] != rune('l') {
											goto l215
										}
										position++
										goto l214
									l215:
										position, tokenIndex = position214, tokenIndex214
										if buffer[position] != rune('L') {
											goto l91
										}
										position++
									}
								l214:
									{
										position216, tokenIndex216 := position, tokenIndex
										if buffer[position] != rune('e') {
											goto l217
										}
										position++
										goto l216
									l217:
										position, tokenIndex = position216, tokenIndex216
										if buffer[position] != rune('E') {
											goto l91
										}
										position++
									}
								l216:
									break
								}
							}

							if !_rules[rulews]() {
								goto l91
							}
							goto l90
						l91:
							position, tokenIndex = position91, tokenIndex91
						}
						add(ruleSpecValue, position79)
					}
					goto l60
				l78:
					position, tokenIndex = position60, tokenIndex60
					{
						position219 := position
						if !_rules[rulews]() {
							goto l218
						}
						{
							position220 := position
							{
								position221, tokenIndex221 := position, tokenIndex
								{
									position223, tokenIndex223 := position, tokenIndex
									if buffer[position] != rune('r') {
										goto l224
									}
									position++
									goto l223
								l224:
									position, tokenIndex = position223, tokenIndex223
									if buffer[position] != rune('R') {
										goto l222
									}
									position++
								}
							l223:
								{
									position225, tokenIndex225 := position, tokenIndex
									if buffer[position] != rune('e') {
										goto l226
									}
									position++
									goto l225
								l226:
									position, tokenIndex = position225, tokenIndex225
									if buffer[position] != rune('E') {
										goto l222
									}
									position++
								}
							l225:
								{
									position227, tokenIndex227 := position, tokenIndex
									if buffer[position] != rune('t') {
										goto l228
									}
									position++
									goto l227
								l228:
									position, tokenIndex = position227, tokenIndex227
									if buffer[position] != rune('T') {
										goto l222
									}
									position++
								}
							l227:
								{
									position229, tokenIndex229 := position, tokenIndex
									if buffer[position] != rune('n') {
										goto l230
									}
									position++
									goto l229
								l230:
									position, tokenIndex = position229, tokenIndex229
									if buffer[position] != rune('N') {
										goto l222
									}
									position++
								}
							l229:
								goto l221
							l222:
								position, tokenIndex = position221, tokenIndex221
								{
									position232, tokenIndex232 := position, tokenIndex
									if buffer[position] != rune('p') {
										goto l233
									}
									position++
									goto l232
								l233:
									position, tokenIndex = position232, tokenIndex232
									if buffer[position] != rune('P') {
										goto l231
									}
									position++
								}
							l232:
								{
									position234, tokenIndex234 := position, tokenIndex
									if buffer[position] != rune('r') {
										goto l235
									}
									position++
									goto l234
								l235:
									position, tokenIndex = position234, tokenIndex234
									if buffer[position] != rune('R') {
										goto l231
									}
									position++
								}
							l234:
								{
									position236, tokenIndex236 := position, tokenIndex
									if buffer[position] != rune('m') {
										goto l237
									}
									position++
									goto l236
								l237:
									position, tokenIndex = position236, tokenIndex236
									if buffer[position] != rune('M') {
										goto l231
									}
									position++
								}
							l236:
								{
									position238, tokenIndex238 := position, tokenIndex
									if buffer[position] != rune('s') {
										goto l239
									}
									position++
									goto l238
								l239:
									position, tokenIndex = position238, tokenIndex238
									if buffer[position] != rune('S') {
										goto l231
									}
									position++
								}
							l238:
								goto l221
							l231:
								position, tokenIndex = position221, tokenIndex221
								{
									position241, tokenIndex241 := position, tokenIndex
									if buffer[position] != rune('a') {
										goto l242
									}
									position++
									goto l241
								l242:
									position, tokenIndex = position241, tokenIndex241
									if buffer[position] != rune('A') {
										goto l240
									}
									position++
								}
							l241:
								{
									position243, tokenIndex243 := position, tokenIndex
									if buffer[position] != rune('r') {
										goto l244
									}
									position++
									goto l243
								l244:
									position, tokenIndex = position243, tokenIndex243
									if buffer[position] != rune('R') {
										goto l240
									}
									position++
								}
							l243:
								{
									position245, tokenIndex245 := position, tokenIndex
									if buffer[position] != rune('g') {
										goto l246
									}
									position++
									goto l245
								l246:
									position, tokenIndex = position245, tokenIndex245
									if buffer[position] != rune('G') {
										goto l240
									}
									position++
								}
							l245:
								{
									position247, tokenIndex247 := position, tokenIndex
									if buffer[position] != rune('s') {
										goto l248
									}
									position++
									goto l247
								l248:
									position, tokenIndex = position247, tokenIndex247
									if buffer[position] != rune('S') {
										goto l240
									}
									position++
								}
							l247:
								goto l221
							l240:
								position, tokenIndex = position221, tokenIndex221
								{
									position250, tokenIndex250 := position, tokenIndex
									if buffer[position] != rune('b') {
										goto l251
									}
									position++
									goto l250
								l251:
									position, tokenIndex = position250, tokenIndex250
									if buffer[position] != rune('B') {
										goto l249
									}
									position++
								}
							l250:
								{
									position252, tokenIndex252 := position, tokenIndex
									if buffer[position] != rune('a') {
										goto l253
									}
									position++
									goto l252
								l253:
									position, tokenIndex = position252, tokenIndex252
									if buffer[position] != rune('A') {
										goto l249
									}
									position++
								}
							l252:
								{
									position254, tokenIndex254 := position, tokenIndex
									if buffer[position] != rune('s') {
										goto l255
									}
									position++
									goto l254
								l255:
									position, tokenIndex = position254, tokenIndex254
									if buffer[position] != rune('S') {
										goto l249
									}
									position++
								}
							l254:
								{
									position256, tokenIndex256 := position, tokenIndex
									if buffer[position] != rune('e') {
										goto l257
									}
									position++
									goto l256
								l257:
									position, tokenIndex = position256, tokenIndex256
									if buffer[position] != rune('E') {
										goto l249
									}
									position++
								}
							l256:
								goto l221
							l249:
								position, tokenIndex = position221, tokenIndex221
								{
									position259, tokenIndex259 := position, tokenIndex
									if buffer[position] != rune('o') {
										goto l260
									}
									position++
									goto l259
								l260:
									position, tokenIndex = position259, tokenIndex259
									if buffer[position] != rune('O') {
										goto l258
									}
									position++
								}
							l259:
								{
									position261, tokenIndex261 := position, tokenIndex
									if buffer[position] != rune('r') {
										goto l262
									}
									position++
									goto l261
								l262:
									position, tokenIndex = position261, tokenIndex261
									if buffer[position] != rune('R') {
										goto l258
									}
									position++
								}
							l261:
								{
									position263, tokenIndex263 := position, tokenIndex
									if buffer[position] != rune('i') {
										goto l264
									}
									position++
									goto l263
								l264:
									position, tokenIndex = position263, tokenIndex263
									if buffer[position] != rune('I') {
										goto l258
									}
									position++
								}
							l263:
								{
									position265, tokenIndex265 := position, tokenIndex
									if buffer[position] != rune('g') {
										goto l266
									}
									position++
									goto l265
								l266:
									position, tokenIndex = position265, tokenIndex265
									if buffer[position] != rune('G') {
										goto l258
									}
									position++
								}
							l265:
								goto l221
							l258:
								position, tokenIndex = position221, tokenIndex221
								{
									position268, tokenIndex268 := position, tokenIndex
									if buffer[position] != rune('o') {
										goto l269
									}
									position++
									goto l268
								l269:
									position, tokenIndex = position268, tokenIndex268
									if buffer[position] != rune('O') {
										goto l267
									}
									position++
								}
							l268:
								{
									position270, tokenIndex270 := position, tokenIndex
									if buffer[position] != rune('p') {
										goto l271
									}
									position++
									goto l270
								l271:
									position, tokenIndex = position270, tokenIndex270
									if buffer[position] != rune('P') {
										goto l267
									}
									position++
								}
							l270:
								if buffer[position] != rune('0') {
									goto l267
								}
								position++
								goto l221
							l267:
								position, tokenIndex = position221, tokenIndex221
								{
									position273, tokenIndex273 := position, tokenIndex
									if buffer[position] != rune('d') {
										goto l274
									}
									position++
									goto l273
								l274:
									position, tokenIndex = position273, tokenIndex273
									if buffer[position] != rune('D') {
										goto l272
									}
									position++
								}
							l273:
								{
									position275, tokenIndex275 := position, tokenIndex
									if buffer[position] != rune('e') {
										goto l276
									}
									position++
									goto l275
								l276:
									position, tokenIndex = position275, tokenIndex275
									if buffer[position] != rune('E') {
										goto l272
									}
									position++
								}
							l275:
								{
									position277, tokenIndex277 := position, tokenIndex
									if buffer[position] != rune('c') {
										goto l278
									}
									position++
									goto l277
								l278:
									position, tokenIndex = position277, tokenIndex277
									if buffer[position] != rune('C') {
										goto l272
									}
									position++
								}
							l277:
								{
									position279, tokenIndex279 := position, tokenIndex
									if buffer[position] != rune('l') {
										goto l280
									}
									position++
									goto l279
								l280:
									position, tokenIndex = position279, tokenIndex279
									if buffer[position] != rune('L') {
										goto l272
									}
									position++
								}
							l279:
								goto l221
							l272:
								position, tokenIndex = position221, tokenIndex221
								{
									position282, tokenIndex282 := position, tokenIndex
									if buffer[position] != rune('d') {
										goto l283
									}
									position++
									goto l282
								l283:
									position, tokenIndex = position282, tokenIndex282
									if buffer[position] != rune('D') {
										goto l281
									}
									position++
								}
							l282:
								{
									position284, tokenIndex284 := position, tokenIndex
									if buffer[position] != rune('c') {
										goto l285
									}
									position++
									goto l284
								l285:
									position, tokenIndex = position284, tokenIndex284
									if buffer[position] != rune('C') {
										goto l281
									}
									position++
								}
							l284:
								{
									position286, tokenIndex286 := position, tokenIndex
									if buffer[position] != rune('l') {
										goto l287
									}
									position++
									goto l286
								l287:
									position, tokenIndex = position286, tokenIndex286
									if buffer[position] != rune('L') {
										goto l281
									}
									position++
								}
							l286:
								{
									position288, tokenIndex288 := position, tokenIndex
									if buffer[position] != rune('s') {
										goto l289
									}
									position++
									goto l288
								l289:
									position, tokenIndex = position288, tokenIndex288
									if buffer[position] != rune('S') {
										goto l281
									}
									position++
								}
							l288:
								goto l221
							l281:
								position, tokenIndex = position221, tokenIndex221
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
										goto l290
									}
									position++
								}
							l291:
								{
									position293, tokenIndex293 := position, tokenIndex
									if buffer[position] != rune('x') {
										goto l294
									}
									position++
									goto l293
								l294:
									position, tokenIndex = position293, tokenIndex293
									if buffer[position] != rune('X') {
										goto l290
									}
									position++
								}
							l293:
								{
									position295, tokenIndex295 := position, tokenIndex
									if buffer[position] != rune('p') {
										goto l296
									}
									position++
									goto l295
								l296:
									position, tokenIndex = position295, tokenIndex295
									if buffer[position] != rune('P') {
										goto l290
									}
									position++
								}
							l295:
								{
									position297, tokenIndex297 := position, tokenIndex
									if buffer[position] != rune('r') {
										goto l298
									}
									position++
									goto l297
								l298:
									position, tokenIndex = position297, tokenIndex297
									if buffer[position] != rune('R') {
										goto l290
									}
									position++
								}
							l297:
								goto l221
							l290:
								position, tokenIndex = position221, tokenIndex221
								{
									position300, tokenIndex300 := position, tokenIndex
									if buffer[position] != rune('f') {
										goto l301
									}
									position++
									goto l300
								l301:
									position, tokenIndex = position300, tokenIndex300
									if buffer[position] != rune('F') {
										goto l299
									}
									position++
								}
							l300:
								{
									position302, tokenIndex302 := position, tokenIndex
									if buffer[position] != rune('n') {
										goto l303
									}
									position++
									goto l302
								l303:
									position, tokenIndex = position302, tokenIndex302
									if buffer[position] != rune('N') {
										goto l299
									}
									position++
								}
							l302:
								{
									position304, tokenIndex304 := position, tokenIndex
									if buffer[position] != rune('c') {
										goto l305
									}
									position++
									goto l304
								l305:
									position, tokenIndex = position304, tokenIndex304
									if buffer[position] != rune('C') {
										goto l299
									}
									position++
								}
							l304:
								{
									position306, tokenIndex306 := position, tokenIndex
									if buffer[position] != rune('s') {
										goto l307
									}
									position++
									goto l306
								l307:
									position, tokenIndex = position306, tokenIndex306
									if buffer[position] != rune('S') {
										goto l299
									}
									position++
								}
							l306:
								goto l221
							l299:
								position, tokenIndex = position221, tokenIndex221
								{
									position309, tokenIndex309 := position, tokenIndex
									if buffer[position] != rune('f') {
										goto l310
									}
									position++
									goto l309
								l310:
									position, tokenIndex = position309, tokenIndex309
									if buffer[position] != rune('F') {
										goto l308
									}
									position++
								}
							l309:
								{
									position311, tokenIndex311 := position, tokenIndex
									if buffer[position] != rune('n') {
										goto l312
									}
									position++
									goto l311
								l312:
									position, tokenIndex = position311, tokenIndex311
									if buffer[position] != rune('N') {
										goto l308
									}
									position++
								}
							l311:
								goto l221
							l308:
								position, tokenIndex = position221, tokenIndex221
								{
									position314, tokenIndex314 := position, tokenIndex
									if buffer[position] != rune('c') {
										goto l315
									}
									position++
									goto l314
								l315:
									position, tokenIndex = position314, tokenIndex314
									if buffer[position] != rune('C') {
										goto l313
									}
									position++
								}
							l314:
								{
									position316, tokenIndex316 := position, tokenIndex
									if buffer[position] != rune('n') {
										goto l317
									}
									position++
									goto l316
								l317:
									position, tokenIndex = position316, tokenIndex316
									if buffer[position] != rune('N') {
										goto l313
									}
									position++
								}
							l316:
								{
									position318, tokenIndex318 := position, tokenIndex
									if buffer[position] != rune('s') {
										goto l319
									}
									position++
									goto l318
								l319:
									position, tokenIndex = position318, tokenIndex318
									if buffer[position] != rune('S') {
										goto l313
									}
									position++
								}
							l318:
								{
									position320, tokenIndex320 := position, tokenIndex
									if buffer[position] != rune('t') {
										goto l321
									}
									position++
									goto l320
								l321:
									position, tokenIndex = position320, tokenIndex320
									if buffer[position] != rune('T') {
										goto l313
									}
									position++
								}
							l320:
								goto l221
							l313:
								position, tokenIndex = position221, tokenIndex221
								{
									position323, tokenIndex323 := position, tokenIndex
									if buffer[position] != rune('v') {
										goto l324
									}
									position++
									goto l323
								l324:
									position, tokenIndex = position323, tokenIndex323
									if buffer[position] != rune('V') {
										goto l322
									}
									position++
								}
							l323:
								{
									position325, tokenIndex325 := position, tokenIndex
									if buffer[position] != rune('a') {
										goto l326
									}
									position++
									goto l325
								l326:
									position, tokenIndex = position325, tokenIndex325
									if buffer[position] != rune('A') {
										goto l322
									}
									position++
								}
							l325:
								{
									position327, tokenIndex327 := position, tokenIndex
									if buffer[position] != rune('r') {
										goto l328
									}
									position++
									goto l327
								l328:
									position, tokenIndex = position327, tokenIndex327
									if buffer[position] != rune('R') {
										goto l322
									}
									position++
								}
							l327:
								{
									position329, tokenIndex329 := position, tokenIndex
									if buffer[position] != rune('s') {
										goto l330
									}
									position++
									goto l329
								l330:
									position, tokenIndex = position329, tokenIndex329
									if buffer[position] != rune('S') {
										goto l322
									}
									position++
								}
							l329:
								goto l221
							l322:
								position, tokenIndex = position221, tokenIndex221
								{
									position332, tokenIndex332 := position, tokenIndex
									if buffer[position] != rune('v') {
										goto l333
									}
									position++
									goto l332
								l333:
									position, tokenIndex = position332, tokenIndex332
									if buffer[position] != rune('V') {
										goto l331
									}
									position++
								}
							l332:
								{
									position334, tokenIndex334 := position, tokenIndex
									if buffer[position] != rune('f') {
										goto l335
									}
									position++
									goto l334
								l335:
									position, tokenIndex = position334, tokenIndex334
									if buffer[position] != rune('F') {
										goto l331
									}
									position++
								}
							l334:
								{
									position336, tokenIndex336 := position, tokenIndex
									if buffer[position] != rune('l') {
										goto l337
									}
									position++
									goto l336
								l337:
									position, tokenIndex = position336, tokenIndex336
									if buffer[position] != rune('L') {
										goto l331
									}
									position++
								}
							l336:
								{
									position338, tokenIndex338 := position, tokenIndex
									if buffer[position] != rune('d') {
										goto l339
									}
									position++
									goto l338
								l339:
									position, tokenIndex = position338, tokenIndex338
									if buffer[position] != rune('D') {
										goto l331
									}
									position++
								}
							l338:
								goto l221
							l331:
								position, tokenIndex = position221, tokenIndex221
								{
									position341, tokenIndex341 := position, tokenIndex
									if buffer[position] != rune('v') {
										goto l342
									}
									position++
									goto l341
								l342:
									position, tokenIndex = position341, tokenIndex341
									if buffer[position] != rune('V') {
										goto l340
									}
									position++
								}
							l341:
								{
									position343, tokenIndex343 := position, tokenIndex
									if buffer[position] != rune('a') {
										goto l344
									}
									position++
									goto l343
								l344:
									position, tokenIndex = position343, tokenIndex343
									if buffer[position] != rune('A') {
										goto l340
									}
									position++
								}
							l343:
								{
									position345, tokenIndex345 := position, tokenIndex
									if buffer[position] != rune('l') {
										goto l346
									}
									position++
									goto l345
								l346:
									position, tokenIndex = position345, tokenIndex345
									if buffer[position] != rune('L') {
										goto l340
									}
									position++
								}
							l345:
								{
									position347, tokenIndex347 := position, tokenIndex
									if buffer[position] != rune('u') {
										goto l348
									}
									position++
									goto l347
								l348:
									position, tokenIndex = position347, tokenIndex347
									if buffer[position] != rune('U') {
										goto l340
									}
									position++
								}
							l347:
								goto l221
							l340:
								position, tokenIndex = position221, tokenIndex221
								{
									position350, tokenIndex350 := position, tokenIndex
									if buffer[position] != rune('c') {
										goto l351
									}
									position++
									goto l350
								l351:
									position, tokenIndex = position350, tokenIndex350
									if buffer[position] != rune('C') {
										goto l349
									}
									position++
								}
							l350:
								{
									position352, tokenIndex352 := position, tokenIndex
									if buffer[position] != rune('h') {
										goto l353
									}
									position++
									goto l352
								l353:
									position, tokenIndex = position352, tokenIndex352
									if buffer[position] != rune('H') {
										goto l349
									}
									position++
								}
							l352:
								{
									position354, tokenIndex354 := position, tokenIndex
									if buffer[position] != rune('a') {
										goto l355
									}
									position++
									goto l354
								l355:
									position, tokenIndex = position354, tokenIndex354
									if buffer[position] != rune('A') {
										goto l349
									}
									position++
								}
							l354:
								{
									position356, tokenIndex356 := position, tokenIndex
									if buffer[position] != rune('n') {
										goto l357
									}
									position++
									goto l356
								l357:
									position, tokenIndex = position356, tokenIndex356
									if buffer[position] != rune('N') {
										goto l349
									}
									position++
								}
							l356:
								goto l221
							l349:
								position, tokenIndex = position221, tokenIndex221
								{
									position359, tokenIndex359 := position, tokenIndex
									if buffer[position] != rune('p') {
										goto l360
									}
									position++
									goto l359
								l360:
									position, tokenIndex = position359, tokenIndex359
									if buffer[position] != rune('P') {
										goto l358
									}
									position++
								}
							l359:
								{
									position361, tokenIndex361 := position, tokenIndex
									if buffer[position] != rune('u') {
										goto l362
									}
									position++
									goto l361
								l362:
									position, tokenIndex = position361, tokenIndex361
									if buffer[position] != rune('U') {
										goto l358
									}
									position++
								}
							l361:
								{
									position363, tokenIndex363 := position, tokenIndex
									if buffer[position] != rune('r') {
										goto l364
									}
									position++
									goto l363
								l364:
									position, tokenIndex = position363, tokenIndex363
									if buffer[position] != rune('R') {
										goto l358
									}
									position++
								}
							l363:
								{
									position365, tokenIndex365 := position, tokenIndex
									if buffer[position] != rune('p') {
										goto l366
									}
									position++
									goto l365
								l366:
									position, tokenIndex = position365, tokenIndex365
									if buffer[position] != rune('P') {
										goto l358
									}
									position++
								}
							l365:
								goto l221
							l358:
								position, tokenIndex = position221, tokenIndex221
								{
									position368, tokenIndex368 := position, tokenIndex
									if buffer[position] != rune('c') {
										goto l369
									}
									position++
									goto l368
								l369:
									position, tokenIndex = position368, tokenIndex368
									if buffer[position] != rune('C') {
										goto l367
									}
									position++
								}
							l368:
								{
									position370, tokenIndex370 := position, tokenIndex
									if buffer[position] != rune('l') {
										goto l371
									}
									position++
									goto l370
								l371:
									position, tokenIndex = position370, tokenIndex370
									if buffer[position] != rune('L') {
										goto l367
									}
									position++
								}
							l370:
								{
									position372, tokenIndex372 := position, tokenIndex
									if buffer[position] != rune('s') {
										goto l373
									}
									position++
									goto l372
								l373:
									position, tokenIndex = position372, tokenIndex372
									if buffer[position] != rune('S') {
										goto l367
									}
									position++
								}
							l372:
								goto l221
							l367:
								position, tokenIndex = position221, tokenIndex221
								{
									position375, tokenIndex375 := position, tokenIndex
									if buffer[position] != rune('c') {
										goto l376
									}
									position++
									goto l375
								l376:
									position, tokenIndex = position375, tokenIndex375
									if buffer[position] != rune('C') {
										goto l374
									}
									position++
								}
							l375:
								{
									position377, tokenIndex377 := position, tokenIndex
									if buffer[position] != rune('s') {
										goto l378
									}
									position++
									goto l377
								l378:
									position, tokenIndex = position377, tokenIndex377
									if buffer[position] != rune('S') {
										goto l374
									}
									position++
								}
							l377:
								{
									position379, tokenIndex379 := position, tokenIndex
									if buffer[position] != rune('t') {
										goto l380
									}
									position++
									goto l379
								l380:
									position, tokenIndex = position379, tokenIndex379
									if buffer[position] != rune('T') {
										goto l374
									}
									position++
								}
							l379:
								{
									position381, tokenIndex381 := position, tokenIndex
									if buffer[position] != rune('s') {
										goto l382
									}
									position++
									goto l381
								l382:
									position, tokenIndex = position381, tokenIndex381
									if buffer[position] != rune('S') {
										goto l374
									}
									position++
								}
							l381:
								goto l221
							l374:
								position, tokenIndex = position221, tokenIndex221
								{
									position384, tokenIndex384 := position, tokenIndex
									if buffer[position] != rune('t') {
										goto l385
									}
									position++
									goto l384
								l385:
									position, tokenIndex = position384, tokenIndex384
									if buffer[position] != rune('T') {
										goto l383
									}
									position++
								}
							l384:
								{
									position386, tokenIndex386 := position, tokenIndex
									if buffer[position] != rune('y') {
										goto l387
									}
									position++
									goto l386
								l387:
									position, tokenIndex = position386, tokenIndex386
									if buffer[position] != rune('Y') {
										goto l383
									}
									position++
								}
							l386:
								{
									position388, tokenIndex388 := position, tokenIndex
									if buffer[position] != rune('p') {
										goto l389
									}
									position++
									goto l388
								l389:
									position, tokenIndex = position388, tokenIndex388
									if buffer[position] != rune('P') {
										goto l383
									}
									position++
								}
							l388:
								{
									position390, tokenIndex390 := position, tokenIndex
									if buffer[position] != rune('e') {
										goto l391
									}
									position++
									goto l390
								l391:
									position, tokenIndex = position390, tokenIndex390
									if buffer[position] != rune('E') {
										goto l383
									}
									position++
								}
							l390:
								goto l221
							l383:
								position, tokenIndex = position221, tokenIndex221
								{
									position393, tokenIndex393 := position, tokenIndex
									if buffer[position] != rune('m') {
										goto l394
									}
									position++
									goto l393
								l394:
									position, tokenIndex = position393, tokenIndex393
									if buffer[position] != rune('M') {
										goto l392
									}
									position++
								}
							l393:
								{
									position395, tokenIndex395 := position, tokenIndex
									if buffer[position] != rune('i') {
										goto l396
									}
									position++
									goto l395
								l396:
									position, tokenIndex = position395, tokenIndex395
									if buffer[position] != rune('I') {
										goto l392
									}
									position++
								}
							l395:
								{
									position397, tokenIndex397 := position, tokenIndex
									if buffer[position] != rune('n') {
										goto l398
									}
									position++
									goto l397
								l398:
									position, tokenIndex = position397, tokenIndex397
									if buffer[position] != rune('N') {
										goto l392
									}
									position++
								}
							l397:
								goto l221
							l392:
								position, tokenIndex = position221, tokenIndex221
								{
									position400, tokenIndex400 := position, tokenIndex
									if buffer[position] != rune('m') {
										goto l401
									}
									position++
									goto l400
								l401:
									position, tokenIndex = position400, tokenIndex400
									if buffer[position] != rune('M') {
										goto l399
									}
									position++
								}
							l400:
								{
									position402, tokenIndex402 := position, tokenIndex
									if buffer[position] != rune('a') {
										goto l403
									}
									position++
									goto l402
								l403:
									position, tokenIndex = position402, tokenIndex402
									if buffer[position] != rune('A') {
										goto l399
									}
									position++
								}
							l402:
								{
									position404, tokenIndex404 := position, tokenIndex
									if buffer[position] != rune('x') {
										goto l405
									}
									position++
									goto l404
								l405:
									position, tokenIndex = position404, tokenIndex404
									if buffer[position] != rune('X') {
										goto l399
									}
									position++
								}
							l404:
								goto l221
							l399:
								position, tokenIndex = position221, tokenIndex221
								{
									position407, tokenIndex407 := position, tokenIndex
									if buffer[position] != rune('b') {
										goto l408
									}
									position++
									goto l407
								l408:
									position, tokenIndex = position407, tokenIndex407
									if buffer[position] != rune('B') {
										goto l406
									}
									position++
								}
							l407:
								{
									position409, tokenIndex409 := position, tokenIndex
									if buffer[position] != rune('i') {
										goto l410
									}
									position++
									goto l409
								l410:
									position, tokenIndex = position409, tokenIndex409
									if buffer[position] != rune('I') {
										goto l406
									}
									position++
								}
							l409:
								{
									position411, tokenIndex411 := position, tokenIndex
									if buffer[position] != rune('n') {
										goto l412
									}
									position++
									goto l411
								l412:
									position, tokenIndex = position411, tokenIndex411
									if buffer[position] != rune('N') {
										goto l406
									}
									position++
								}
							l411:
								{
									position413, tokenIndex413 := position, tokenIndex
									if buffer[position] != rune('f') {
										goto l414
									}
									position++
									goto l413
								l414:
									position, tokenIndex = position413, tokenIndex413
									if buffer[position] != rune('F') {
										goto l406
									}
									position++
								}
							l413:
								goto l221
							l406:
								position, tokenIndex = position221, tokenIndex221
								{
									position416, tokenIndex416 := position, tokenIndex
									if buffer[position] != rune('s') {
										goto l417
									}
									position++
									goto l416
								l417:
									position, tokenIndex = position416, tokenIndex416
									if buffer[position] != rune('S') {
										goto l415
									}
									position++
								}
							l416:
								{
									position418, tokenIndex418 := position, tokenIndex
									if buffer[position] != rune('c') {
										goto l419
									}
									position++
									goto l418
								l419:
									position, tokenIndex = position418, tokenIndex418
									if buffer[position] != rune('C') {
										goto l415
									}
									position++
								}
							l418:
								{
									position420, tokenIndex420 := position, tokenIndex
									if buffer[position] != rune('p') {
										goto l421
									}
									position++
									goto l420
								l421:
									position, tokenIndex = position420, tokenIndex420
									if buffer[position] != rune('P') {
										goto l415
									}
									position++
								}
							l420:
								{
									position422, tokenIndex422 := position, tokenIndex
									if buffer[position] != rune('e') {
										goto l423
									}
									position++
									goto l422
								l423:
									position, tokenIndex = position422, tokenIndex422
									if buffer[position] != rune('E') {
										goto l415
									}
									position++
								}
							l422:
								goto l221
							l415:
								position, tokenIndex = position221, tokenIndex221
								{
									position425, tokenIndex425 := position, tokenIndex
									if buffer[position] != rune('i') {
										goto l426
									}
									position++
									goto l425
								l426:
									position, tokenIndex = position425, tokenIndex425
									if buffer[position] != rune('I') {
										goto l424
									}
									position++
								}
							l425:
								{
									position427, tokenIndex427 := position, tokenIndex
									if buffer[position] != rune('n') {
										goto l428
									}
									position++
									goto l427
								l428:
									position, tokenIndex = position427, tokenIndex427
									if buffer[position] != rune('N') {
										goto l424
									}
									position++
								}
							l427:
								{
									position429, tokenIndex429 := position, tokenIndex
									if buffer[position] != rune('i') {
										goto l430
									}
									position++
									goto l429
								l430:
									position, tokenIndex = position429, tokenIndex429
									if buffer[position] != rune('I') {
										goto l424
									}
									position++
								}
							l429:
								{
									position431, tokenIndex431 := position, tokenIndex
									if buffer[position] != rune('t') {
										goto l432
									}
									position++
									goto l431
								l432:
									position, tokenIndex = position431, tokenIndex431
									if buffer[position] != rune('T') {
										goto l424
									}
									position++
								}
							l431:
								goto l221
							l424:
								position, tokenIndex = position221, tokenIndex221
								{
									switch buffer[position] {
									case 'O', 'o':
										{
											position434 := position
											{
												position435, tokenIndex435 := position, tokenIndex
												if buffer[position] != rune('o') {
													goto l436
												}
												position++
												goto l435
											l436:
												position, tokenIndex = position435, tokenIndex435
												if buffer[position] != rune('O') {
													goto l218
												}
												position++
											}
										l435:
											{
												position437, tokenIndex437 := position, tokenIndex
												if buffer[position] != rune('p') {
													goto l438
												}
												position++
												goto l437
											l438:
												position, tokenIndex = position437, tokenIndex437
												if buffer[position] != rune('P') {
													goto l218
												}
												position++
											}
										l437:
											if buffer[position] != rune(' ') {
												goto l218
											}
											position++
											if c := buffer[position]; c < rune('0') || c > rune('9') {
												goto l218
											}
											position++
											add(ruleOpAttr, position434)
										}
										break
									case 'V', 'v':
										{
											position439, tokenIndex439 := position, tokenIndex
											if buffer[position] != rune('v') {
												goto l440
											}
											position++
											goto l439
										l440:
											position, tokenIndex = position439, tokenIndex439
											if buffer[position] != rune('V') {
												goto l218
											}
											position++
										}
									l439:
										{
											position441, tokenIndex441 := position, tokenIndex
											if buffer[position] != rune('a') {
												goto l442
											}
											position++
											goto l441
										l442:
											position, tokenIndex = position441, tokenIndex441
											if buffer[position] != rune('A') {
												goto l218
											}
											position++
										}
									l441:
										{
											position443, tokenIndex443 := position, tokenIndex
											if buffer[position] != rune('l') {
												goto l444
											}
											position++
											goto l443
										l444:
											position, tokenIndex = position443, tokenIndex443
											if buffer[position] != rune('L') {
												goto l218
											}
											position++
										}
									l443:
										break
									case 'I', 'i':
										{
											position445, tokenIndex445 := position, tokenIndex
											if buffer[position] != rune('i') {
												goto l446
											}
											position++
											goto l445
										l446:
											position, tokenIndex = position445, tokenIndex445
											if buffer[position] != rune('I') {
												goto l218
											}
											position++
										}
									l445:
										{
											position447, tokenIndex447 := position, tokenIndex
											if buffer[position] != rune('d') {
												goto l448
											}
											position++
											goto l447
										l448:
											position, tokenIndex = position447, tokenIndex447
											if buffer[position] != rune('D') {
												goto l218
											}
											position++
										}
									l447:
										{
											position449, tokenIndex449 := position, tokenIndex
											if buffer[position] != rune('x') {
												goto l450
											}
											position++
											goto l449
										l450:
											position, tokenIndex = position449, tokenIndex449
											if buffer[position] != rune('X') {
												goto l218
											}
											position++
										}
									l449:
										break
									case 'S', 's':
										{
											position451, tokenIndex451 := position, tokenIndex
											if buffer[position] != rune('s') {
												goto l452
											}
											position++
											goto l451
										l452:
											position, tokenIndex = position451, tokenIndex451
											if buffer[position] != rune('S') {
												goto l218
											}
											position++
										}
									l451:
										{
											position453, tokenIndex453 := position, tokenIndex
											if buffer[position] != rune('i') {
												goto l454
											}
											position++
											goto l453
										l454:
											position, tokenIndex = position453, tokenIndex453
											if buffer[position] != rune('I') {
												goto l218
											}
											position++
										}
									l453:
										{
											position455, tokenIndex455 := position, tokenIndex
											if buffer[position] != rune('z') {
												goto l456
											}
											position++
											goto l455
										l456:
											position, tokenIndex = position455, tokenIndex455
											if buffer[position] != rune('Z') {
												goto l218
											}
											position++
										}
									l455:
										{
											position457, tokenIndex457 := position, tokenIndex
											if buffer[position] != rune('e') {
												goto l458
											}
											position++
											goto l457
										l458:
											position, tokenIndex = position457, tokenIndex457
											if buffer[position] != rune('E') {
												goto l218
											}
											position++
										}
									l457:
										break
									case 'B', 'b':
										{
											position459, tokenIndex459 := position, tokenIndex
											if buffer[position] != rune('b') {
												goto l460
											}
											position++
											goto l459
										l460:
											position, tokenIndex = position459, tokenIndex459
											if buffer[position] != rune('B') {
												goto l218
											}
											position++
										}
									l459:
										{
											position461, tokenIndex461 := position, tokenIndex
											if buffer[position] != rune('p') {
												goto l462
											}
											position++
											goto l461
										l462:
											position, tokenIndex = position461, tokenIndex461
											if buffer[position] != rune('P') {
												goto l218
											}
											position++
										}
									l461:
										{
											position463, tokenIndex463 := position, tokenIndex
											if buffer[position] != rune('o') {
												goto l464
											}
											position++
											goto l463
										l464:
											position, tokenIndex = position463, tokenIndex463
											if buffer[position] != rune('O') {
												goto l218
											}
											position++
										}
									l463:
										{
											position465, tokenIndex465 := position, tokenIndex
											if buffer[position] != rune('s') {
												goto l466
											}
											position++
											goto l465
										l466:
											position, tokenIndex = position465, tokenIndex465
											if buffer[position] != rune('S') {
												goto l218
											}
											position++
										}
									l465:
										break
									case 'C', 'c':
										{
											position467, tokenIndex467 := position, tokenIndex
											if buffer[position] != rune('c') {
												goto l468
											}
											position++
											goto l467
										l468:
											position, tokenIndex = position467, tokenIndex467
											if buffer[position] != rune('C') {
												goto l218
											}
											position++
										}
									l467:
										{
											position469, tokenIndex469 := position, tokenIndex
											if buffer[position] != rune('h') {
												goto l470
											}
											position++
											goto l469
										l470:
											position, tokenIndex = position469, tokenIndex469
											if buffer[position] != rune('H') {
												goto l218
											}
											position++
										}
									l469:
										{
											position471, tokenIndex471 := position, tokenIndex
											if buffer[position] != rune('a') {
												goto l472
											}
											position++
											goto l471
										l472:
											position, tokenIndex = position471, tokenIndex471
											if buffer[position] != rune('A') {
												goto l218
											}
											position++
										}
									l471:
										{
											position473, tokenIndex473 := position, tokenIndex
											if buffer[position] != rune('i') {
												goto l474
											}
											position++
											goto l473
										l474:
											position, tokenIndex = position473, tokenIndex473
											if buffer[position] != rune('I') {
												goto l218
											}
											position++
										}
									l473:
										{
											position475, tokenIndex475 := position, tokenIndex
											if buffer[position] != rune('n') {
												goto l476
											}
											position++
											goto l475
										l476:
											position, tokenIndex = position475, tokenIndex475
											if buffer[position] != rune('N') {
												goto l218
											}
											position++
										}
									l475:
										break
									case 'E', 'e':
										{
											position477, tokenIndex477 := position, tokenIndex
											if buffer[position] != rune('e') {
												goto l478
											}
											position++
											goto l477
										l478:
											position, tokenIndex = position477, tokenIndex477
											if buffer[position] != rune('E') {
												goto l218
											}
											position++
										}
									l477:
										{
											position479, tokenIndex479 := position, tokenIndex
											if buffer[position] != rune('l') {
												goto l480
											}
											position++
											goto l479
										l480:
											position, tokenIndex = position479, tokenIndex479
											if buffer[position] != rune('L') {
												goto l218
											}
											position++
										}
									l479:
										{
											position481, tokenIndex481 := position, tokenIndex
											if buffer[position] != rune('t') {
												goto l482
											}
											position++
											goto l481
										l482:
											position, tokenIndex = position481, tokenIndex481
											if buffer[position] != rune('T') {
												goto l218
											}
											position++
										}
									l481:
										{
											position483, tokenIndex483 := position, tokenIndex
											if buffer[position] != rune('s') {
												goto l484
											}
											position++
											goto l483
										l484:
											position, tokenIndex = position483, tokenIndex483
											if buffer[position] != rune('S') {
												goto l218
											}
											position++
										}
									l483:
										break
									case 'D', 'd':
										{
											position485, tokenIndex485 := position, tokenIndex
											if buffer[position] != rune('d') {
												goto l486
											}
											position++
											goto l485
										l486:
											position, tokenIndex = position485, tokenIndex485
											if buffer[position] != rune('D') {
												goto l218
											}
											position++
										}
									l485:
										{
											position487, tokenIndex487 := position, tokenIndex
											if buffer[position] != rune('o') {
												goto l488
											}
											position++
											goto l487
										l488:
											position, tokenIndex = position487, tokenIndex487
											if buffer[position] != rune('O') {
												goto l218
											}
											position++
										}
									l487:
										{
											position489, tokenIndex489 := position, tokenIndex
											if buffer[position] != rune('m') {
												goto l490
											}
											position++
											goto l489
										l490:
											position, tokenIndex = position489, tokenIndex489
											if buffer[position] != rune('M') {
												goto l218
											}
											position++
										}
									l489:
										{
											position491, tokenIndex491 := position, tokenIndex
											if buffer[position] != rune('n') {
												goto l492
											}
											position++
											goto l491
										l492:
											position, tokenIndex = position491, tokenIndex491
											if buffer[position] != rune('N') {
												goto l218
											}
											position++
										}
									l491:
										break
									case 'T', 't':
										{
											position493, tokenIndex493 := position, tokenIndex
											if buffer[position] != rune('t') {
												goto l494
											}
											position++
											goto l493
										l494:
											position, tokenIndex = position493, tokenIndex493
											if buffer[position] != rune('T') {
												goto l218
											}
											position++
										}
									l493:
										{
											position495, tokenIndex495 := position, tokenIndex
											if buffer[position] != rune('y') {
												goto l496
											}
											position++
											goto l495
										l496:
											position, tokenIndex = position495, tokenIndex495
											if buffer[position] != rune('Y') {
												goto l218
											}
											position++
										}
									l495:
										{
											position497, tokenIndex497 := position, tokenIndex
											if buffer[position] != rune('p') {
												goto l498
											}
											position++
											goto l497
										l498:
											position, tokenIndex = position497, tokenIndex497
											if buffer[position] != rune('P') {
												goto l218
											}
											position++
										}
									l497:
										{
											position499, tokenIndex499 := position, tokenIndex
											if buffer[position] != rune('e') {
												goto l500
											}
											position++
											goto l499
										l500:
											position, tokenIndex = position499, tokenIndex499
											if buffer[position] != rune('E') {
												goto l218
											}
											position++
										}
									l499:
										break
									case 'M', 'm':
										{
											position501, tokenIndex501 := position, tokenIndex
											if buffer[position] != rune('m') {
												goto l502
											}
											position++
											goto l501
										l502:
											position, tokenIndex = position501, tokenIndex501
											if buffer[position] != rune('M') {
												goto l218
											}
											position++
										}
									l501:
										{
											position503, tokenIndex503 := position, tokenIndex
											if buffer[position] != rune('n') {
												goto l504
											}
											position++
											goto l503
										l504:
											position, tokenIndex = position503, tokenIndex503
											if buffer[position] != rune('N') {
												goto l218
											}
											position++
										}
									l503:
										{
											position505, tokenIndex505 := position, tokenIndex
											if buffer[position] != rune('g') {
												goto l506
											}
											position++
											goto l505
										l506:
											position, tokenIndex = position505, tokenIndex505
											if buffer[position] != rune('G') {
												goto l218
											}
											position++
										}
									l505:
										{
											position507, tokenIndex507 := position, tokenIndex
											if buffer[position] != rune('l') {
												goto l508
											}
											position++
											goto l507
										l508:
											position, tokenIndex = position507, tokenIndex507
											if buffer[position] != rune('L') {
												goto l218
											}
											position++
										}
									l507:
										break
									case 'N', 'n':
										{
											position509, tokenIndex509 := position, tokenIndex
											if buffer[position] != rune('n') {
												goto l510
											}
											position++
											goto l509
										l510:
											position, tokenIndex = position509, tokenIndex509
											if buffer[position] != rune('N') {
												goto l218
											}
											position++
										}
									l509:
										{
											position511, tokenIndex511 := position, tokenIndex
											if buffer[position] != rune('a') {
												goto l512
											}
											position++
											goto l511
										l512:
											position, tokenIndex = position511, tokenIndex511
											if buffer[position] != rune('A') {
												goto l218
											}
											position++
										}
									l511:
										{
											position513, tokenIndex513 := position, tokenIndex
											if buffer[position] != rune('m') {
												goto l514
											}
											position++
											goto l513
										l514:
											position, tokenIndex = position513, tokenIndex513
											if buffer[position] != rune('M') {
												goto l218
											}
											position++
										}
									l513:
										{
											position515, tokenIndex515 := position, tokenIndex
											if buffer[position] != rune('e') {
												goto l516
											}
											position++
											goto l515
										l516:
											position, tokenIndex = position515, tokenIndex515
											if buffer[position] != rune('E') {
												goto l218
											}
											position++
										}
									l515:
										break
									case 'U', 'u':
										{
											position517, tokenIndex517 := position, tokenIndex
											if buffer[position] != rune('u') {
												goto l518
											}
											position++
											goto l517
										l518:
											position, tokenIndex = position517, tokenIndex517
											if buffer[position] != rune('U') {
												goto l218
											}
											position++
										}
									l517:
										{
											position519, tokenIndex519 := position, tokenIndex
											if buffer[position] != rune('n') {
												goto l520
											}
											position++
											goto l519
										l520:
											position, tokenIndex = position519, tokenIndex519
											if buffer[position] != rune('N') {
												goto l218
											}
											position++
										}
									l519:
										{
											position521, tokenIndex521 := position, tokenIndex
											if buffer[position] != rune('q') {
												goto l522
											}
											position++
											goto l521
										l522:
											position, tokenIndex = position521, tokenIndex521
											if buffer[position] != rune('Q') {
												goto l218
											}
											position++
										}
									l521:
										{
											position523, tokenIndex523 := position, tokenIndex
											if buffer[position] != rune('l') {
												goto l524
											}
											position++
											goto l523
										l524:
											position, tokenIndex = position523, tokenIndex523
											if buffer[position] != rune('L') {
												goto l218
											}
											position++
										}
									l523:
										break
									case 'L', 'l':
										{
											position525, tokenIndex525 := position, tokenIndex
											if buffer[position] != rune('l') {
												goto l526
											}
											position++
											goto l525
										l526:
											position, tokenIndex = position525, tokenIndex525
											if buffer[position] != rune('L') {
												goto l218
											}
											position++
										}
									l525:
										{
											position527, tokenIndex527 := position, tokenIndex
											if buffer[position] != rune('a') {
												goto l528
											}
											position++
											goto l527
										l528:
											position, tokenIndex = position527, tokenIndex527
											if buffer[position] != rune('A') {
												goto l218
											}
											position++
										}
									l527:
										{
											position529, tokenIndex529 := position, tokenIndex
											if buffer[position] != rune('b') {
												goto l530
											}
											position++
											goto l529
										l530:
											position, tokenIndex = position529, tokenIndex529
											if buffer[position] != rune('B') {
												goto l218
											}
											position++
										}
									l529:
										{
											position531, tokenIndex531 := position, tokenIndex
											if buffer[position] != rune('l') {
												goto l532
											}
											position++
											goto l531
										l532:
											position, tokenIndex = position531, tokenIndex531
											if buffer[position] != rune('L') {
												goto l218
											}
											position++
										}
									l531:
										break
									case 'P', 'p':
										{
											position533, tokenIndex533 := position, tokenIndex
											if buffer[position] != rune('p') {
												goto l534
											}
											position++
											goto l533
										l534:
											position, tokenIndex = position533, tokenIndex533
											if buffer[position] != rune('P') {
												goto l218
											}
											position++
										}
									l533:
										{
											position535, tokenIndex535 := position, tokenIndex
											if buffer[position] != rune('t') {
												goto l536
											}
											position++
											goto l535
										l536:
											position, tokenIndex = position535, tokenIndex535
											if buffer[position] != rune('T') {
												goto l218
											}
											position++
										}
									l535:
										{
											position537, tokenIndex537 := position, tokenIndex
											if buffer[position] != rune('d') {
												goto l538
											}
											position++
											goto l537
										l538:
											position, tokenIndex = position537, tokenIndex537
											if buffer[position] != rune('D') {
												goto l218
											}
											position++
										}
									l537:
										break
									case 'F', 'f':
										{
											position539, tokenIndex539 := position, tokenIndex
											if buffer[position] != rune('f') {
												goto l540
											}
											position++
											goto l539
										l540:
											position, tokenIndex = position539, tokenIndex539
											if buffer[position] != rune('F') {
												goto l218
											}
											position++
										}
									l539:
										{
											position541, tokenIndex541 := position, tokenIndex
											if buffer[position] != rune('l') {
												goto l542
											}
											position++
											goto l541
										l542:
											position, tokenIndex = position541, tokenIndex541
											if buffer[position] != rune('L') {
												goto l218
											}
											position++
										}
									l541:
										{
											position543, tokenIndex543 := position, tokenIndex
											if buffer[position] != rune('d') {
												goto l544
											}
											position++
											goto l543
										l544:
											position, tokenIndex = position543, tokenIndex543
											if buffer[position] != rune('D') {
												goto l218
											}
											position++
										}
									l543:
										{
											position545, tokenIndex545 := position, tokenIndex
											if buffer[position] != rune('s') {
												goto l546
											}
											position++
											goto l545
										l546:
											position, tokenIndex = position545, tokenIndex545
											if buffer[position] != rune('S') {
												goto l218
											}
											position++
										}
									l545:
										break
									case 'R', 'r':
										{
											position547, tokenIndex547 := position, tokenIndex
											if buffer[position] != rune('r') {
												goto l548
											}
											position++
											goto l547
										l548:
											position, tokenIndex = position547, tokenIndex547
											if buffer[position] != rune('R') {
												goto l218
											}
											position++
										}
									l547:
										{
											position549, tokenIndex549 := position, tokenIndex
											if buffer[position] != rune('e') {
												goto l550
											}
											position++
											goto l549
										l550:
											position, tokenIndex = position549, tokenIndex549
											if buffer[position] != rune('E') {
												goto l218
											}
											position++
										}
									l549:
										{
											position551, tokenIndex551 := position, tokenIndex
											if buffer[position] != rune('f') {
												goto l552
											}
											position++
											goto l551
										l552:
											position, tokenIndex = position551, tokenIndex551
											if buffer[position] != rune('F') {
												goto l218
											}
											position++
										}
									l551:
										{
											position553, tokenIndex553 := position, tokenIndex
											if buffer[position] != rune('d') {
												goto l554
											}
											position++
											goto l553
										l554:
											position, tokenIndex = position553, tokenIndex553
											if buffer[position] != rune('D') {
												goto l218
											}
											position++
										}
									l553:
										break
									case 'A', 'a':
										{
											position555, tokenIndex555 := position, tokenIndex
											if buffer[position] != rune('a') {
												goto l556
											}
											position++
											goto l555
										l556:
											position, tokenIndex = position555, tokenIndex555
											if buffer[position] != rune('A') {
												goto l218
											}
											position++
										}
									l555:
										{
											position557, tokenIndex557 := position, tokenIndex
											if buffer[position] != rune('r') {
												goto l558
											}
											position++
											goto l557
										l558:
											position, tokenIndex = position557, tokenIndex557
											if buffer[position] != rune('R') {
												goto l218
											}
											position++
										}
									l557:
										{
											position559, tokenIndex559 := position, tokenIndex
											if buffer[position] != rune('g') {
												goto l560
											}
											position++
											goto l559
										l560:
											position, tokenIndex = position559, tokenIndex559
											if buffer[position] != rune('G') {
												goto l218
											}
											position++
										}
									l559:
										{
											position561, tokenIndex561 := position, tokenIndex
											if buffer[position] != rune('t') {
												goto l562
											}
											position++
											goto l561
										l562:
											position, tokenIndex = position561, tokenIndex561
											if buffer[position] != rune('T') {
												goto l218
											}
											position++
										}
									l561:
										break
									default:
										if !_rules[ruleInteger]() {
											goto l218
										}
										break
									}
								}

							}
						l221:
							add(ruleNodeName, position220)
						}
						if !_rules[rulews]() {
							goto l218
						}
						if buffer[position] != rune(':') {
							goto l218
						}
						position++
						if !_rules[rulews]() {
							goto l218
						}
						if !_rules[ruleNode]() {
							goto l218
						}
						if !_rules[rulews]() {
							goto l218
						}
						add(ruleNodeAttr, position219)
					}
					goto l60
				l218:
					position, tokenIndex = position60, tokenIndex60
					{
						position564 := position
						if buffer[position] != rune('s') {
							goto l563
						}
						position++
						if buffer[position] != rune('r') {
							goto l563
						}
						position++
						if buffer[position] != rune('c') {
							goto l563
						}
						position++
						if buffer[position] != rune('p') {
							goto l563
						}
						position++
						if buffer[position] != rune(':') {
							goto l563
						}
						position++
						if !_rules[rulews]() {
							goto l563
						}
						{
							position565 := position
							{
								position566, tokenIndex566 := position, tokenIndex
								{
									position568, tokenIndex568 := position, tokenIndex
									if buffer[position] != rune('_') {
										goto l568
									}
									position++
									goto l569
								l568:
									position, tokenIndex = position568, tokenIndex568
								}
							l569:
								{
									switch buffer[position] {
									case '_':
										if buffer[position] != rune('_') {
											goto l567
										}
										position++
										break
									case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
										if c := buffer[position]; c < rune('0') || c > rune('9') {
											goto l567
										}
										position++
										break
									case 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z':
										if c := buffer[position]; c < rune('a') || c > rune('z') {
											goto l567
										}
										position++
										break
									case '-':
										if buffer[position] != rune('-') {
											goto l567
										}
										position++
										break
									default:
										if c := buffer[position]; c < rune('A') || c > rune('Z') {
											goto l567
										}
										position++
										break
									}
								}

							l570:
								{
									position571, tokenIndex571 := position, tokenIndex
									{
										switch buffer[position] {
										case '_':
											if buffer[position] != rune('_') {
												goto l571
											}
											position++
											break
										case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
											if c := buffer[position]; c < rune('0') || c > rune('9') {
												goto l571
											}
											position++
											break
										case 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z':
											if c := buffer[position]; c < rune('a') || c > rune('z') {
												goto l571
											}
											position++
											break
										case '-':
											if buffer[position] != rune('-') {
												goto l571
											}
											position++
											break
										default:
											if c := buffer[position]; c < rune('A') || c > rune('Z') {
												goto l571
											}
											position++
											break
										}
									}

									goto l570
								l571:
									position, tokenIndex = position571, tokenIndex571
								}
							l574:
								{
									position575, tokenIndex575 := position, tokenIndex
									if buffer[position] != rune('.') {
										goto l575
									}
									position++
									{
										switch buffer[position] {
										case '_':
											if buffer[position] != rune('_') {
												goto l575
											}
											position++
											break
										case 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z':
											if c := buffer[position]; c < rune('a') || c > rune('z') {
												goto l575
											}
											position++
											break
										default:
											if c := buffer[position]; c < rune('A') || c > rune('Z') {
												goto l575
											}
											position++
											break
										}
									}

								l576:
									{
										position577, tokenIndex577 := position, tokenIndex
										{
											switch buffer[position] {
											case '_':
												if buffer[position] != rune('_') {
													goto l577
												}
												position++
												break
											case 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z':
												if c := buffer[position]; c < rune('a') || c > rune('z') {
													goto l577
												}
												position++
												break
											default:
												if c := buffer[position]; c < rune('A') || c > rune('Z') {
													goto l577
												}
												position++
												break
											}
										}

										goto l576
									l577:
										position, tokenIndex = position577, tokenIndex577
									}
									goto l574
								l575:
									position, tokenIndex = position575, tokenIndex575
								}
								goto l566
							l567:
								position, tokenIndex = position566, tokenIndex566
								if buffer[position] != rune('<') {
									goto l563
								}
								position++
								{
									position580, tokenIndex580 := position, tokenIndex
									if buffer[position] != rune('b') {
										goto l581
									}
									position++
									goto l580
								l581:
									position, tokenIndex = position580, tokenIndex580
									if buffer[position] != rune('B') {
										goto l563
									}
									position++
								}
							l580:
								{
									position582, tokenIndex582 := position, tokenIndex
									if buffer[position] != rune('u') {
										goto l583
									}
									position++
									goto l582
								l583:
									position, tokenIndex = position582, tokenIndex582
									if buffer[position] != rune('U') {
										goto l563
									}
									position++
								}
							l582:
								{
									position584, tokenIndex584 := position, tokenIndex
									if buffer[position] != rune('i') {
										goto l585
									}
									position++
									goto l584
								l585:
									position, tokenIndex = position584, tokenIndex584
									if buffer[position] != rune('I') {
										goto l563
									}
									position++
								}
							l584:
								{
									position586, tokenIndex586 := position, tokenIndex
									if buffer[position] != rune('l') {
										goto l587
									}
									position++
									goto l586
								l587:
									position, tokenIndex = position586, tokenIndex586
									if buffer[position] != rune('L') {
										goto l563
									}
									position++
								}
							l586:
								{
									position588, tokenIndex588 := position, tokenIndex
									if buffer[position] != rune('t') {
										goto l589
									}
									position++
									goto l588
								l589:
									position, tokenIndex = position588, tokenIndex588
									if buffer[position] != rune('T') {
										goto l563
									}
									position++
								}
							l588:
								if buffer[position] != rune('-') {
									goto l563
								}
								position++
								{
									position590, tokenIndex590 := position, tokenIndex
									if buffer[position] != rune('i') {
										goto l591
									}
									position++
									goto l590
								l591:
									position, tokenIndex = position590, tokenIndex590
									if buffer[position] != rune('I') {
										goto l563
									}
									position++
								}
							l590:
								{
									position592, tokenIndex592 := position, tokenIndex
									if buffer[position] != rune('n') {
										goto l593
									}
									position++
									goto l592
								l593:
									position, tokenIndex = position592, tokenIndex592
									if buffer[position] != rune('N') {
										goto l563
									}
									position++
								}
							l592:
								if buffer[position] != rune('>') {
									goto l563
								}
								position++
							}
						l566:
							if buffer[position] != rune(':') {
								goto l563
							}
							position++
							if !_rules[ruleInteger]() {
								goto l563
							}
							if !_rules[rulews]() {
								goto l563
							}
							add(ruleFileRef, position565)
						}
						add(ruleSourceAttr, position564)
					}
					goto l60
				l563:
					position, tokenIndex = position60, tokenIndex60
					{
						position595 := position
						{
							switch buffer[position] {
							case 'P', 'p':
								{
									position597, tokenIndex597 := position, tokenIndex
									if buffer[position] != rune('p') {
										goto l598
									}
									position++
									goto l597
								l598:
									position, tokenIndex = position597, tokenIndex597
									if buffer[position] != rune('P') {
										goto l594
									}
									position++
								}
							l597:
								{
									position599, tokenIndex599 := position, tokenIndex
									if buffer[position] != rune('r') {
										goto l600
									}
									position++
									goto l599
								l600:
									position, tokenIndex = position599, tokenIndex599
									if buffer[position] != rune('R') {
										goto l594
									}
									position++
								}
							l599:
								{
									position601, tokenIndex601 := position, tokenIndex
									if buffer[position] != rune('e') {
										goto l602
									}
									position++
									goto l601
								l602:
									position, tokenIndex = position601, tokenIndex601
									if buffer[position] != rune('E') {
										goto l594
									}
									position++
								}
							l601:
								{
									position603, tokenIndex603 := position, tokenIndex
									if buffer[position] != rune('c') {
										goto l604
									}
									position++
									goto l603
								l604:
									position, tokenIndex = position603, tokenIndex603
									if buffer[position] != rune('C') {
										goto l594
									}
									position++
								}
							l603:
								break
							case 'U', 'u':
								{
									position605, tokenIndex605 := position, tokenIndex
									if buffer[position] != rune('u') {
										goto l606
									}
									position++
									goto l605
								l606:
									position, tokenIndex = position605, tokenIndex605
									if buffer[position] != rune('U') {
										goto l594
									}
									position++
								}
							l605:
								{
									position607, tokenIndex607 := position, tokenIndex
									if buffer[position] != rune('s') {
										goto l608
									}
									position++
									goto l607
								l608:
									position, tokenIndex = position607, tokenIndex607
									if buffer[position] != rune('S') {
										goto l594
									}
									position++
								}
							l607:
								{
									position609, tokenIndex609 := position, tokenIndex
									if buffer[position] != rune('e') {
										goto l610
									}
									position++
									goto l609
								l610:
									position, tokenIndex = position609, tokenIndex609
									if buffer[position] != rune('E') {
										goto l594
									}
									position++
								}
							l609:
								{
									position611, tokenIndex611 := position, tokenIndex
									if buffer[position] != rune('d') {
										goto l612
									}
									position++
									goto l611
								l612:
									position, tokenIndex = position611, tokenIndex611
									if buffer[position] != rune('D') {
										goto l594
									}
									position++
								}
							l611:
								break
							case 'L', 'l':
								{
									position613, tokenIndex613 := position, tokenIndex
									if buffer[position] != rune('l') {
										goto l614
									}
									position++
									goto l613
								l614:
									position, tokenIndex = position613, tokenIndex613
									if buffer[position] != rune('L') {
										goto l594
									}
									position++
								}
							l613:
								{
									position615, tokenIndex615 := position, tokenIndex
									if buffer[position] != rune('o') {
										goto l616
									}
									position++
									goto l615
								l616:
									position, tokenIndex = position615, tokenIndex615
									if buffer[position] != rune('O') {
										goto l594
									}
									position++
								}
							l615:
								{
									position617, tokenIndex617 := position, tokenIndex
									if buffer[position] != rune('w') {
										goto l618
									}
									position++
									goto l617
								l618:
									position, tokenIndex = position617, tokenIndex617
									if buffer[position] != rune('W') {
										goto l594
									}
									position++
								}
							l617:
								break
							default:
								{
									position619, tokenIndex619 := position, tokenIndex
									if buffer[position] != rune('b') {
										goto l620
									}
									position++
									goto l619
								l620:
									position, tokenIndex = position619, tokenIndex619
									if buffer[position] != rune('B') {
										goto l594
									}
									position++
								}
							l619:
								{
									position621, tokenIndex621 := position, tokenIndex
									if buffer[position] != rune('a') {
										goto l622
									}
									position++
									goto l621
								l622:
									position, tokenIndex = position621, tokenIndex621
									if buffer[position] != rune('A') {
										goto l594
									}
									position++
								}
							l621:
								{
									position623, tokenIndex623 := position, tokenIndex
									if buffer[position] != rune('s') {
										goto l624
									}
									position++
									goto l623
								l624:
									position, tokenIndex = position623, tokenIndex623
									if buffer[position] != rune('S') {
										goto l594
									}
									position++
								}
							l623:
								{
									position625, tokenIndex625 := position, tokenIndex
									if buffer[position] != rune('e') {
										goto l626
									}
									position++
									goto l625
								l626:
									position, tokenIndex = position625, tokenIndex625
									if buffer[position] != rune('E') {
										goto l594
									}
									position++
								}
							l625:
								{
									position627, tokenIndex627 := position, tokenIndex
									if buffer[position] != rune('s') {
										goto l628
									}
									position++
									goto l627
								l628:
									position, tokenIndex = position627, tokenIndex627
									if buffer[position] != rune('S') {
										goto l594
									}
									position++
								}
							l627:
								break
							}
						}

						if !_rules[rulews]() {
							goto l594
						}
						if buffer[position] != rune(':') {
							goto l594
						}
						position++
						if !_rules[rulews]() {
							goto l594
						}
						if !_rules[ruleInteger]() {
							goto l594
						}
						if !_rules[rulews]() {
							goto l594
						}
						add(ruleIntAttr, position595)
					}
					goto l60
				l594:
					position, tokenIndex = position60, tokenIndex60
					{
						position630 := position
						{
							position631, tokenIndex631 := position, tokenIndex
							if buffer[position] != rune('s') {
								goto l632
							}
							position++
							goto l631
						l632:
							position, tokenIndex = position631, tokenIndex631
							if buffer[position] != rune('S') {
								goto l629
							}
							position++
						}
					l631:
						{
							position633, tokenIndex633 := position, tokenIndex
							if buffer[position] != rune('t') {
								goto l634
							}
							position++
							goto l633
						l634:
							position, tokenIndex = position633, tokenIndex633
							if buffer[position] != rune('T') {
								goto l629
							}
							position++
						}
					l633:
						{
							position635, tokenIndex635 := position, tokenIndex
							if buffer[position] != rune('r') {
								goto l636
							}
							position++
							goto l635
						l636:
							position, tokenIndex = position635, tokenIndex635
							if buffer[position] != rune('R') {
								goto l629
							}
							position++
						}
					l635:
						{
							position637, tokenIndex637 := position, tokenIndex
							if buffer[position] != rune('g') {
								goto l638
							}
							position++
							goto l637
						l638:
							position, tokenIndex = position637, tokenIndex637
							if buffer[position] != rune('G') {
								goto l629
							}
							position++
						}
					l637:
						if buffer[position] != rune(':') {
							goto l629
						}
						position++
						if !matchDot() {
							goto l629
						}
					l639:
						{
							position640, tokenIndex640 := position, tokenIndex
							if !matchDot() {
								goto l640
							}
							goto l639
						l640:
							position, tokenIndex = position640, tokenIndex640
						}
						add(ruleStringAttr, position630)
					}
					goto l60
				l629:
					position, tokenIndex = position60, tokenIndex60
					{
						position642 := position
						{
							position643, tokenIndex643 := position, tokenIndex
							if buffer[position] != rune('s') {
								goto l644
							}
							position++
							goto l643
						l644:
							position, tokenIndex = position643, tokenIndex643
							if buffer[position] != rune('S') {
								goto l641
							}
							position++
						}
					l643:
						{
							position645, tokenIndex645 := position, tokenIndex
							if buffer[position] != rune('i') {
								goto l646
							}
							position++
							goto l645
						l646:
							position, tokenIndex = position645, tokenIndex645
							if buffer[position] != rune('I') {
								goto l641
							}
							position++
						}
					l645:
						{
							position647, tokenIndex647 := position, tokenIndex
							if buffer[position] != rune('g') {
								goto l648
							}
							position++
							goto l647
						l648:
							position, tokenIndex = position647, tokenIndex647
							if buffer[position] != rune('G') {
								goto l641
							}
							position++
						}
					l647:
						{
							position649, tokenIndex649 := position, tokenIndex
							if buffer[position] != rune('n') {
								goto l650
							}
							position++
							goto l649
						l650:
							position, tokenIndex = position649, tokenIndex649
							if buffer[position] != rune('N') {
								goto l641
							}
							position++
						}
					l649:
						if !_rules[rulews]() {
							goto l641
						}
						if buffer[position] != rune(':') {
							goto l641
						}
						position++
						if !_rules[rulews]() {
							goto l641
						}
						{
							position651, tokenIndex651 := position, tokenIndex
							{
								position653, tokenIndex653 := position, tokenIndex
								if buffer[position] != rune('s') {
									goto l654
								}
								position++
								goto l653
							l654:
								position, tokenIndex = position653, tokenIndex653
								if buffer[position] != rune('S') {
									goto l652
								}
								position++
							}
						l653:
							{
								position655, tokenIndex655 := position, tokenIndex
								if buffer[position] != rune('i') {
									goto l656
								}
								position++
								goto l655
							l656:
								position, tokenIndex = position655, tokenIndex655
								if buffer[position] != rune('I') {
									goto l652
								}
								position++
							}
						l655:
							{
								position657, tokenIndex657 := position, tokenIndex
								if buffer[position] != rune('g') {
									goto l658
								}
								position++
								goto l657
							l658:
								position, tokenIndex = position657, tokenIndex657
								if buffer[position] != rune('G') {
									goto l652
								}
								position++
							}
						l657:
							{
								position659, tokenIndex659 := position, tokenIndex
								if buffer[position] != rune('n') {
									goto l660
								}
								position++
								goto l659
							l660:
								position, tokenIndex = position659, tokenIndex659
								if buffer[position] != rune('N') {
									goto l652
								}
								position++
							}
						l659:
							{
								position661, tokenIndex661 := position, tokenIndex
								if buffer[position] != rune('e') {
									goto l662
								}
								position++
								goto l661
							l662:
								position, tokenIndex = position661, tokenIndex661
								if buffer[position] != rune('E') {
									goto l652
								}
								position++
							}
						l661:
							{
								position663, tokenIndex663 := position, tokenIndex
								if buffer[position] != rune('d') {
									goto l664
								}
								position++
								goto l663
							l664:
								position, tokenIndex = position663, tokenIndex663
								if buffer[position] != rune('D') {
									goto l652
								}
								position++
							}
						l663:
							goto l651
						l652:
							position, tokenIndex = position651, tokenIndex651
							{
								position665, tokenIndex665 := position, tokenIndex
								if buffer[position] != rune('u') {
									goto l666
								}
								position++
								goto l665
							l666:
								position, tokenIndex = position665, tokenIndex665
								if buffer[position] != rune('U') {
									goto l641
								}
								position++
							}
						l665:
							{
								position667, tokenIndex667 := position, tokenIndex
								if buffer[position] != rune('n') {
									goto l668
								}
								position++
								goto l667
							l668:
								position, tokenIndex = position667, tokenIndex667
								if buffer[position] != rune('N') {
									goto l641
								}
								position++
							}
						l667:
							{
								position669, tokenIndex669 := position, tokenIndex
								if buffer[position] != rune('s') {
									goto l670
								}
								position++
								goto l669
							l670:
								position, tokenIndex = position669, tokenIndex669
								if buffer[position] != rune('S') {
									goto l641
								}
								position++
							}
						l669:
							{
								position671, tokenIndex671 := position, tokenIndex
								if buffer[position] != rune('i') {
									goto l672
								}
								position++
								goto l671
							l672:
								position, tokenIndex = position671, tokenIndex671
								if buffer[position] != rune('I') {
									goto l641
								}
								position++
							}
						l671:
							{
								position673, tokenIndex673 := position, tokenIndex
								if buffer[position] != rune('g') {
									goto l674
								}
								position++
								goto l673
							l674:
								position, tokenIndex = position673, tokenIndex673
								if buffer[position] != rune('G') {
									goto l641
								}
								position++
							}
						l673:
							{
								position675, tokenIndex675 := position, tokenIndex
								if buffer[position] != rune('n') {
									goto l676
								}
								position++
								goto l675
							l676:
								position, tokenIndex = position675, tokenIndex675
								if buffer[position] != rune('N') {
									goto l641
								}
								position++
							}
						l675:
							{
								position677, tokenIndex677 := position, tokenIndex
								if buffer[position] != rune('e') {
									goto l678
								}
								position++
								goto l677
							l678:
								position, tokenIndex = position677, tokenIndex677
								if buffer[position] != rune('E') {
									goto l641
								}
								position++
							}
						l677:
							{
								position679, tokenIndex679 := position, tokenIndex
								if buffer[position] != rune('d') {
									goto l680
								}
								position++
								goto l679
							l680:
								position, tokenIndex = position679, tokenIndex679
								if buffer[position] != rune('D') {
									goto l641
								}
								position++
							}
						l679:
						}
					l651:
						add(ruleSignAttr, position642)
					}
					goto l60
				l641:
					position, tokenIndex = position60, tokenIndex60
					{
						position682 := position
						{
							position683, tokenIndex683 := position, tokenIndex
							{
								position685, tokenIndex685 := position, tokenIndex
								if buffer[position] != rune('l') {
									goto l686
								}
								position++
								goto l685
							l686:
								position, tokenIndex = position685, tokenIndex685
								if buffer[position] != rune('L') {
									goto l684
								}
								position++
							}
						l685:
							{
								position687, tokenIndex687 := position, tokenIndex
								if buffer[position] != rune('o') {
									goto l688
								}
								position++
								goto l687
							l688:
								position, tokenIndex = position687, tokenIndex687
								if buffer[position] != rune('O') {
									goto l684
								}
								position++
							}
						l687:
							{
								position689, tokenIndex689 := position, tokenIndex
								if buffer[position] != rune('w') {
									goto l690
								}
								position++
								goto l689
							l690:
								position, tokenIndex = position689, tokenIndex689
								if buffer[position] != rune('W') {
									goto l684
								}
								position++
							}
						l689:
							goto l683
						l684:
							position, tokenIndex = position683, tokenIndex683
							{
								position691, tokenIndex691 := position, tokenIndex
								if buffer[position] != rune('h') {
									goto l692
								}
								position++
								goto l691
							l692:
								position, tokenIndex = position691, tokenIndex691
								if buffer[position] != rune('H') {
									goto l681
								}
								position++
							}
						l691:
							{
								position693, tokenIndex693 := position, tokenIndex
								if buffer[position] != rune('i') {
									goto l694
								}
								position++
								goto l693
							l694:
								position, tokenIndex = position693, tokenIndex693
								if buffer[position] != rune('I') {
									goto l681
								}
								position++
							}
						l693:
							{
								position695, tokenIndex695 := position, tokenIndex
								if buffer[position] != rune('g') {
									goto l696
								}
								position++
								goto l695
							l696:
								position, tokenIndex = position695, tokenIndex695
								if buffer[position] != rune('G') {
									goto l681
								}
								position++
							}
						l695:
							{
								position697, tokenIndex697 := position, tokenIndex
								if buffer[position] != rune('h') {
									goto l698
								}
								position++
								goto l697
							l698:
								position, tokenIndex = position697, tokenIndex697
								if buffer[position] != rune('H') {
									goto l681
								}
								position++
							}
						l697:
						}
					l683:
						if !_rules[rulews]() {
							goto l681
						}
						if buffer[position] != rune(':') {
							goto l681
						}
						position++
						if !_rules[rulews]() {
							goto l681
						}
						{
							position699, tokenIndex699 := position, tokenIndex
							if buffer[position] != rune('-') {
								goto l699
							}
							position++
							goto l700
						l699:
							position, tokenIndex = position699, tokenIndex699
						}
					l700:
						if !_rules[ruleInteger]() {
							goto l681
						}
						if !_rules[rulews]() {
							goto l681
						}
						add(ruleIntAttr3, position682)
					}
					goto l60
				l681:
					position, tokenIndex = position60, tokenIndex60
					{
						position702 := position
						{
							position703, tokenIndex703 := position, tokenIndex
							if buffer[position] != rune('l') {
								goto l704
							}
							position++
							goto l703
						l704:
							position, tokenIndex = position703, tokenIndex703
							if buffer[position] != rune('L') {
								goto l701
							}
							position++
						}
					l703:
						{
							position705, tokenIndex705 := position, tokenIndex
							if buffer[position] != rune('n') {
								goto l706
							}
							position++
							goto l705
						l706:
							position, tokenIndex = position705, tokenIndex705
							if buffer[position] != rune('N') {
								goto l701
							}
							position++
						}
					l705:
						{
							position707, tokenIndex707 := position, tokenIndex
							if buffer[position] != rune('g') {
								goto l708
							}
							position++
							goto l707
						l708:
							position, tokenIndex = position707, tokenIndex707
							if buffer[position] != rune('G') {
								goto l701
							}
							position++
						}
					l707:
						{
							position709, tokenIndex709 := position, tokenIndex
							if buffer[position] != rune('t') {
								goto l710
							}
							position++
							goto l709
						l710:
							position, tokenIndex = position709, tokenIndex709
							if buffer[position] != rune('T') {
								goto l701
							}
							position++
						}
					l709:
						if !_rules[rulews]() {
							goto l701
						}
						if buffer[position] != rune(':') {
							goto l701
						}
						position++
						if !_rules[rulews]() {
							goto l701
						}
						if !_rules[ruleInteger]() {
							goto l701
						}
						if !_rules[rulews]() {
							goto l701
						}
						add(ruleLngtAttr, position702)
					}
					goto l60
				l701:
					position, tokenIndex = position60, tokenIndex60
					{
						position712 := position
						{
							position713, tokenIndex713 := position, tokenIndex
							if buffer[position] != rune('a') {
								goto l714
							}
							position++
							goto l713
						l714:
							position, tokenIndex = position713, tokenIndex713
							if buffer[position] != rune('A') {
								goto l711
							}
							position++
						}
					l713:
						{
							position715, tokenIndex715 := position, tokenIndex
							if buffer[position] != rune('c') {
								goto l716
							}
							position++
							goto l715
						l716:
							position, tokenIndex = position715, tokenIndex715
							if buffer[position] != rune('C') {
								goto l711
							}
							position++
						}
					l715:
						{
							position717, tokenIndex717 := position, tokenIndex
							if buffer[position] != rune('c') {
								goto l718
							}
							position++
							goto l717
						l718:
							position, tokenIndex = position717, tokenIndex717
							if buffer[position] != rune('C') {
								goto l711
							}
							position++
						}
					l717:
						{
							position719, tokenIndex719 := position, tokenIndex
							if buffer[position] != rune('s') {
								goto l720
							}
							position++
							goto l719
						l720:
							position, tokenIndex = position719, tokenIndex719
							if buffer[position] != rune('S') {
								goto l711
							}
							position++
						}
					l719:
						if !_rules[rulews]() {
							goto l711
						}
						if buffer[position] != rune(':') {
							goto l711
						}
						position++
						if !_rules[rulews]() {
							goto l711
						}
						{
							position721, tokenIndex721 := position, tokenIndex
							{
								position723, tokenIndex723 := position, tokenIndex
								if buffer[position] != rune('p') {
									goto l724
								}
								position++
								goto l723
							l724:
								position, tokenIndex = position723, tokenIndex723
								if buffer[position] != rune('P') {
									goto l722
								}
								position++
							}
						l723:
							{
								position725, tokenIndex725 := position, tokenIndex
								if buffer[position] != rune('r') {
									goto l726
								}
								position++
								goto l725
							l726:
								position, tokenIndex = position725, tokenIndex725
								if buffer[position] != rune('R') {
									goto l722
								}
								position++
							}
						l725:
							{
								position727, tokenIndex727 := position, tokenIndex
								if buffer[position] != rune('i') {
									goto l728
								}
								position++
								goto l727
							l728:
								position, tokenIndex = position727, tokenIndex727
								if buffer[position] != rune('I') {
									goto l722
								}
								position++
							}
						l727:
							{
								position729, tokenIndex729 := position, tokenIndex
								if buffer[position] != rune('v') {
									goto l730
								}
								position++
								goto l729
							l730:
								position, tokenIndex = position729, tokenIndex729
								if buffer[position] != rune('V') {
									goto l722
								}
								position++
							}
						l729:
							goto l721
						l722:
							position, tokenIndex = position721, tokenIndex721
							{
								position732, tokenIndex732 := position, tokenIndex
								if buffer[position] != rune('p') {
									goto l733
								}
								position++
								goto l732
							l733:
								position, tokenIndex = position732, tokenIndex732
								if buffer[position] != rune('P') {
									goto l731
								}
								position++
							}
						l732:
							{
								position734, tokenIndex734 := position, tokenIndex
								if buffer[position] != rune('u') {
									goto l735
								}
								position++
								goto l734
							l735:
								position, tokenIndex = position734, tokenIndex734
								if buffer[position] != rune('U') {
									goto l731
								}
								position++
							}
						l734:
							{
								position736, tokenIndex736 := position, tokenIndex
								if buffer[position] != rune('b') {
									goto l737
								}
								position++
								goto l736
							l737:
								position, tokenIndex = position736, tokenIndex736
								if buffer[position] != rune('B') {
									goto l731
								}
								position++
							}
						l736:
							goto l721
						l731:
							position, tokenIndex = position721, tokenIndex721
							{
								position738, tokenIndex738 := position, tokenIndex
								if buffer[position] != rune('p') {
									goto l739
								}
								position++
								goto l738
							l739:
								position, tokenIndex = position738, tokenIndex738
								if buffer[position] != rune('P') {
									goto l711
								}
								position++
							}
						l738:
							{
								position740, tokenIndex740 := position, tokenIndex
								if buffer[position] != rune('r') {
									goto l741
								}
								position++
								goto l740
							l741:
								position, tokenIndex = position740, tokenIndex740
								if buffer[position] != rune('R') {
									goto l711
								}
								position++
							}
						l740:
							{
								position742, tokenIndex742 := position, tokenIndex
								if buffer[position] != rune('o') {
									goto l743
								}
								position++
								goto l742
							l743:
								position, tokenIndex = position742, tokenIndex742
								if buffer[position] != rune('O') {
									goto l711
								}
								position++
							}
						l742:
							{
								position744, tokenIndex744 := position, tokenIndex
								if buffer[position] != rune('t') {
									goto l745
								}
								position++
								goto l744
							l745:
								position, tokenIndex = position744, tokenIndex744
								if buffer[position] != rune('T') {
									goto l711
								}
								position++
							}
						l744:
						}
					l721:
						add(ruleAccsAttr, position712)
					}
					goto l60
				l711:
					position, tokenIndex = position60, tokenIndex60
					{
						switch buffer[position] {
						case 'A', 'a':
							{
								position747 := position
								{
									position748, tokenIndex748 := position, tokenIndex
									if buffer[position] != rune('a') {
										goto l749
									}
									position++
									goto l748
								l749:
									position, tokenIndex = position748, tokenIndex748
									if buffer[position] != rune('A') {
										goto l58
									}
									position++
								}
							l748:
								{
									position750, tokenIndex750 := position, tokenIndex
									if buffer[position] != rune('l') {
										goto l751
									}
									position++
									goto l750
								l751:
									position, tokenIndex = position750, tokenIndex750
									if buffer[position] != rune('L') {
										goto l58
									}
									position++
								}
							l750:
								{
									position752, tokenIndex752 := position, tokenIndex
									if buffer[position] != rune('g') {
										goto l753
									}
									position++
									goto l752
								l753:
									position, tokenIndex = position752, tokenIndex752
									if buffer[position] != rune('G') {
										goto l58
									}
									position++
								}
							l752:
								{
									position754, tokenIndex754 := position, tokenIndex
									if buffer[position] != rune('n') {
										goto l755
									}
									position++
									goto l754
								l755:
									position, tokenIndex = position754, tokenIndex754
									if buffer[position] != rune('N') {
										goto l58
									}
									position++
								}
							l754:
								if buffer[position] != rune(':') {
									goto l58
								}
								position++
								if !_rules[rulews]() {
									goto l58
								}
								if c := buffer[position]; c < rune('0') || c > rune('9') {
									goto l58
								}
								position++
							l756:
								{
									position757, tokenIndex757 := position, tokenIndex
									if c := buffer[position]; c < rune('0') || c > rune('9') {
										goto l757
									}
									position++
									goto l756
								l757:
									position, tokenIndex = position757, tokenIndex757
								}
								if !_rules[rulews]() {
									goto l58
								}
								add(ruleIntAttr2, position747)
							}
							break
						case 'Q', 'q':
							{
								position758 := position
								{
									position759, tokenIndex759 := position, tokenIndex
									if buffer[position] != rune('q') {
										goto l760
									}
									position++
									goto l759
								l760:
									position, tokenIndex = position759, tokenIndex759
									if buffer[position] != rune('Q') {
										goto l58
									}
									position++
								}
							l759:
								{
									position761, tokenIndex761 := position, tokenIndex
									if buffer[position] != rune('u') {
										goto l762
									}
									position++
									goto l761
								l762:
									position, tokenIndex = position761, tokenIndex761
									if buffer[position] != rune('U') {
										goto l58
									}
									position++
								}
							l761:
								{
									position763, tokenIndex763 := position, tokenIndex
									if buffer[position] != rune('a') {
										goto l764
									}
									position++
									goto l763
								l764:
									position, tokenIndex = position763, tokenIndex763
									if buffer[position] != rune('A') {
										goto l58
									}
									position++
								}
							l763:
								{
									position765, tokenIndex765 := position, tokenIndex
									if buffer[position] != rune('l') {
										goto l766
									}
									position++
									goto l765
								l766:
									position, tokenIndex = position765, tokenIndex765
									if buffer[position] != rune('L') {
										goto l58
									}
									position++
								}
							l765:
								if !_rules[rulews]() {
									goto l58
								}
								if buffer[position] != rune(':') {
									goto l58
								}
								position++
								if !_rules[rulews]() {
									goto l58
								}
								{
									switch buffer[position] {
									case 'R', 'r':
										{
											position770, tokenIndex770 := position, tokenIndex
											if buffer[position] != rune('r') {
												goto l771
											}
											position++
											goto l770
										l771:
											position, tokenIndex = position770, tokenIndex770
											if buffer[position] != rune('R') {
												goto l58
											}
											position++
										}
									l770:
										break
									case 'C':
										if buffer[position] != rune('C') {
											goto l58
										}
										position++
										break
									case 'c':
										if buffer[position] != rune('c') {
											goto l58
										}
										position++
										break
									default:
										{
											position772, tokenIndex772 := position, tokenIndex
											if buffer[position] != rune('v') {
												goto l773
											}
											position++
											goto l772
										l773:
											position, tokenIndex = position772, tokenIndex772
											if buffer[position] != rune('V') {
												goto l58
											}
											position++
										}
									l772:
										break
									}
								}

							l767:
								{
									position768, tokenIndex768 := position, tokenIndex
									{
										switch buffer[position] {
										case 'R', 'r':
											{
												position775, tokenIndex775 := position, tokenIndex
												if buffer[position] != rune('r') {
													goto l776
												}
												position++
												goto l775
											l776:
												position, tokenIndex = position775, tokenIndex775
												if buffer[position] != rune('R') {
													goto l768
												}
												position++
											}
										l775:
											break
										case 'C':
											if buffer[position] != rune('C') {
												goto l768
											}
											position++
											break
										case 'c':
											if buffer[position] != rune('c') {
												goto l768
											}
											position++
											break
										default:
											{
												position777, tokenIndex777 := position, tokenIndex
												if buffer[position] != rune('v') {
													goto l778
												}
												position++
												goto l777
											l778:
												position, tokenIndex = position777, tokenIndex777
												if buffer[position] != rune('V') {
													goto l768
												}
												position++
											}
										l777:
											break
										}
									}

									goto l767
								l768:
									position, tokenIndex = position768, tokenIndex768
								}
								if !_rules[rulews]() {
									goto l58
								}
								add(ruleQualAttr, position758)
							}
							break
						case 'L', 'l':
							{
								position779 := position
								{
									position780, tokenIndex780 := position, tokenIndex
									if buffer[position] != rune('l') {
										goto l781
									}
									position++
									goto l780
								l781:
									position, tokenIndex = position780, tokenIndex780
									if buffer[position] != rune('L') {
										goto l58
									}
									position++
								}
							l780:
								{
									position782, tokenIndex782 := position, tokenIndex
									if buffer[position] != rune('i') {
										goto l783
									}
									position++
									goto l782
								l783:
									position, tokenIndex = position782, tokenIndex782
									if buffer[position] != rune('I') {
										goto l58
									}
									position++
								}
							l782:
								{
									position784, tokenIndex784 := position, tokenIndex
									if buffer[position] != rune('n') {
										goto l785
									}
									position++
									goto l784
								l785:
									position, tokenIndex = position784, tokenIndex784
									if buffer[position] != rune('N') {
										goto l58
									}
									position++
								}
							l784:
								{
									position786, tokenIndex786 := position, tokenIndex
									if buffer[position] != rune('k') {
										goto l787
									}
									position++
									goto l786
								l787:
									position, tokenIndex = position786, tokenIndex786
									if buffer[position] != rune('K') {
										goto l58
									}
									position++
								}
							l786:
								if !_rules[rulews]() {
									goto l58
								}
								if buffer[position] != rune(':') {
									goto l58
								}
								position++
								if !_rules[rulews]() {
									goto l58
								}
								{
									position788, tokenIndex788 := position, tokenIndex
									{
										position790, tokenIndex790 := position, tokenIndex
										if buffer[position] != rune('e') {
											goto l791
										}
										position++
										goto l790
									l791:
										position, tokenIndex = position790, tokenIndex790
										if buffer[position] != rune('E') {
											goto l789
										}
										position++
									}
								l790:
									{
										position792, tokenIndex792 := position, tokenIndex
										if buffer[position] != rune('x') {
											goto l793
										}
										position++
										goto l792
									l793:
										position, tokenIndex = position792, tokenIndex792
										if buffer[position] != rune('X') {
											goto l789
										}
										position++
									}
								l792:
									{
										position794, tokenIndex794 := position, tokenIndex
										if buffer[position] != rune('t') {
											goto l795
										}
										position++
										goto l794
									l795:
										position, tokenIndex = position794, tokenIndex794
										if buffer[position] != rune('T') {
											goto l789
										}
										position++
									}
								l794:
									{
										position796, tokenIndex796 := position, tokenIndex
										if buffer[position] != rune('e') {
											goto l797
										}
										position++
										goto l796
									l797:
										position, tokenIndex = position796, tokenIndex796
										if buffer[position] != rune('E') {
											goto l789
										}
										position++
									}
								l796:
									{
										position798, tokenIndex798 := position, tokenIndex
										if buffer[position] != rune('r') {
											goto l799
										}
										position++
										goto l798
									l799:
										position, tokenIndex = position798, tokenIndex798
										if buffer[position] != rune('R') {
											goto l789
										}
										position++
									}
								l798:
									{
										position800, tokenIndex800 := position, tokenIndex
										if buffer[position] != rune('n') {
											goto l801
										}
										position++
										goto l800
									l801:
										position, tokenIndex = position800, tokenIndex800
										if buffer[position] != rune('N') {
											goto l789
										}
										position++
									}
								l800:
									goto l788
								l789:
									position, tokenIndex = position788, tokenIndex788
									{
										position802, tokenIndex802 := position, tokenIndex
										if buffer[position] != rune('s') {
											goto l803
										}
										position++
										goto l802
									l803:
										position, tokenIndex = position802, tokenIndex802
										if buffer[position] != rune('S') {
											goto l58
										}
										position++
									}
								l802:
									{
										position804, tokenIndex804 := position, tokenIndex
										if buffer[position] != rune('t') {
											goto l805
										}
										position++
										goto l804
									l805:
										position, tokenIndex = position804, tokenIndex804
										if buffer[position] != rune('T') {
											goto l58
										}
										position++
									}
								l804:
									{
										position806, tokenIndex806 := position, tokenIndex
										if buffer[position] != rune('a') {
											goto l807
										}
										position++
										goto l806
									l807:
										position, tokenIndex = position806, tokenIndex806
										if buffer[position] != rune('A') {
											goto l58
										}
										position++
									}
								l806:
									{
										position808, tokenIndex808 := position, tokenIndex
										if buffer[position] != rune('t') {
											goto l809
										}
										position++
										goto l808
									l809:
										position, tokenIndex = position808, tokenIndex808
										if buffer[position] != rune('T') {
											goto l58
										}
										position++
									}
								l808:
									{
										position810, tokenIndex810 := position, tokenIndex
										if buffer[position] != rune('i') {
											goto l811
										}
										position++
										goto l810
									l811:
										position, tokenIndex = position810, tokenIndex810
										if buffer[position] != rune('I') {
											goto l58
										}
										position++
									}
								l810:
									{
										position812, tokenIndex812 := position, tokenIndex
										if buffer[position] != rune('c') {
											goto l813
										}
										position++
										goto l812
									l813:
										position, tokenIndex = position812, tokenIndex812
										if buffer[position] != rune('C') {
											goto l58
										}
										position++
									}
								l812:
								}
							l788:
								add(ruleLinkAttr, position779)
							}
							break
						case 'N', 'n':
							{
								position814 := position
								{
									position815, tokenIndex815 := position, tokenIndex
									if buffer[position] != rune('n') {
										goto l816
									}
									position++
									goto l815
								l816:
									position, tokenIndex = position815, tokenIndex815
									if buffer[position] != rune('N') {
										goto l58
									}
									position++
								}
							l815:
								{
									position817, tokenIndex817 := position, tokenIndex
									if buffer[position] != rune('o') {
										goto l818
									}
									position++
									goto l817
								l818:
									position, tokenIndex = position817, tokenIndex817
									if buffer[position] != rune('O') {
										goto l58
									}
									position++
								}
							l817:
								{
									position819, tokenIndex819 := position, tokenIndex
									if buffer[position] != rune('t') {
										goto l820
									}
									position++
									goto l819
								l820:
									position, tokenIndex = position819, tokenIndex819
									if buffer[position] != rune('T') {
										goto l58
									}
									position++
								}
							l819:
								{
									position821, tokenIndex821 := position, tokenIndex
									if buffer[position] != rune('e') {
										goto l822
									}
									position++
									goto l821
								l822:
									position, tokenIndex = position821, tokenIndex821
									if buffer[position] != rune('E') {
										goto l58
									}
									position++
								}
							l821:
								if !_rules[rulews]() {
									goto l58
								}
								if buffer[position] != rune(':') {
									goto l58
								}
								position++
								if !_rules[rulews]() {
									goto l58
								}
								{
									position823, tokenIndex823 := position, tokenIndex
									{
										position825, tokenIndex825 := position, tokenIndex
										if buffer[position] != rune('p') {
											goto l826
										}
										position++
										goto l825
									l826:
										position, tokenIndex = position825, tokenIndex825
										if buffer[position] != rune('P') {
											goto l824
										}
										position++
									}
								l825:
									{
										position827, tokenIndex827 := position, tokenIndex
										if buffer[position] != rune('t') {
											goto l828
										}
										position++
										goto l827
									l828:
										position, tokenIndex = position827, tokenIndex827
										if buffer[position] != rune('T') {
											goto l824
										}
										position++
									}
								l827:
									{
										position829, tokenIndex829 := position, tokenIndex
										if buffer[position] != rune('r') {
											goto l830
										}
										position++
										goto l829
									l830:
										position, tokenIndex = position829, tokenIndex829
										if buffer[position] != rune('R') {
											goto l824
										}
										position++
									}
								l829:
									{
										position831, tokenIndex831 := position, tokenIndex
										if buffer[position] != rune('m') {
											goto l832
										}
										position++
										goto l831
									l832:
										position, tokenIndex = position831, tokenIndex831
										if buffer[position] != rune('M') {
											goto l824
										}
										position++
									}
								l831:
									{
										position833, tokenIndex833 := position, tokenIndex
										if buffer[position] != rune('e') {
											goto l834
										}
										position++
										goto l833
									l834:
										position, tokenIndex = position833, tokenIndex833
										if buffer[position] != rune('E') {
											goto l824
										}
										position++
									}
								l833:
									{
										position835, tokenIndex835 := position, tokenIndex
										if buffer[position] != rune('m') {
											goto l836
										}
										position++
										goto l835
									l836:
										position, tokenIndex = position835, tokenIndex835
										if buffer[position] != rune('M') {
											goto l824
										}
										position++
									}
								l835:
									goto l823
								l824:
									position, tokenIndex = position823, tokenIndex823
									{
										switch buffer[position] {
										case 'D', 'd':
											{
												position838, tokenIndex838 := position, tokenIndex
												if buffer[position] != rune('d') {
													goto l839
												}
												position++
												goto l838
											l839:
												position, tokenIndex = position838, tokenIndex838
												if buffer[position] != rune('D') {
													goto l58
												}
												position++
											}
										l838:
											{
												position840, tokenIndex840 := position, tokenIndex
												if buffer[position] != rune('e') {
													goto l841
												}
												position++
												goto l840
											l841:
												position, tokenIndex = position840, tokenIndex840
												if buffer[position] != rune('E') {
													goto l58
												}
												position++
											}
										l840:
											{
												position842, tokenIndex842 := position, tokenIndex
												if buffer[position] != rune('s') {
													goto l843
												}
												position++
												goto l842
											l843:
												position, tokenIndex = position842, tokenIndex842
												if buffer[position] != rune('S') {
													goto l58
												}
												position++
											}
										l842:
											{
												position844, tokenIndex844 := position, tokenIndex
												if buffer[position] != rune('t') {
													goto l845
												}
												position++
												goto l844
											l845:
												position, tokenIndex = position844, tokenIndex844
												if buffer[position] != rune('T') {
													goto l58
												}
												position++
											}
										l844:
											{
												position846, tokenIndex846 := position, tokenIndex
												if buffer[position] != rune('r') {
													goto l847
												}
												position++
												goto l846
											l847:
												position, tokenIndex = position846, tokenIndex846
												if buffer[position] != rune('R') {
													goto l58
												}
												position++
											}
										l846:
											{
												position848, tokenIndex848 := position, tokenIndex
												if buffer[position] != rune('u') {
													goto l849
												}
												position++
												goto l848
											l849:
												position, tokenIndex = position848, tokenIndex848
												if buffer[position] != rune('U') {
													goto l58
												}
												position++
											}
										l848:
											{
												position850, tokenIndex850 := position, tokenIndex
												if buffer[position] != rune('c') {
													goto l851
												}
												position++
												goto l850
											l851:
												position, tokenIndex = position850, tokenIndex850
												if buffer[position] != rune('C') {
													goto l58
												}
												position++
											}
										l850:
											{
												position852, tokenIndex852 := position, tokenIndex
												if buffer[position] != rune('t') {
													goto l853
												}
												position++
												goto l852
											l853:
												position, tokenIndex = position852, tokenIndex852
												if buffer[position] != rune('T') {
													goto l58
												}
												position++
											}
										l852:
											{
												position854, tokenIndex854 := position, tokenIndex
												if buffer[position] != rune('o') {
													goto l855
												}
												position++
												goto l854
											l855:
												position, tokenIndex = position854, tokenIndex854
												if buffer[position] != rune('O') {
													goto l58
												}
												position++
											}
										l854:
											{
												position856, tokenIndex856 := position, tokenIndex
												if buffer[position] != rune('r') {
													goto l857
												}
												position++
												goto l856
											l857:
												position, tokenIndex = position856, tokenIndex856
												if buffer[position] != rune('R') {
													goto l58
												}
												position++
											}
										l856:
											break
										case 'P', 'p':
											{
												position858, tokenIndex858 := position, tokenIndex
												if buffer[position] != rune('p') {
													goto l859
												}
												position++
												goto l858
											l859:
												position, tokenIndex = position858, tokenIndex858
												if buffer[position] != rune('P') {
													goto l58
												}
												position++
											}
										l858:
											{
												position860, tokenIndex860 := position, tokenIndex
												if buffer[position] != rune('s') {
													goto l861
												}
												position++
												goto l860
											l861:
												position, tokenIndex = position860, tokenIndex860
												if buffer[position] != rune('S') {
													goto l58
												}
												position++
											}
										l860:
											{
												position862, tokenIndex862 := position, tokenIndex
												if buffer[position] != rune('e') {
													goto l863
												}
												position++
												goto l862
											l863:
												position, tokenIndex = position862, tokenIndex862
												if buffer[position] != rune('E') {
													goto l58
												}
												position++
											}
										l862:
											{
												position864, tokenIndex864 := position, tokenIndex
												if buffer[position] != rune('u') {
													goto l865
												}
												position++
												goto l864
											l865:
												position, tokenIndex = position864, tokenIndex864
												if buffer[position] != rune('U') {
													goto l58
												}
												position++
											}
										l864:
											{
												position866, tokenIndex866 := position, tokenIndex
												if buffer[position] != rune('d') {
													goto l867
												}
												position++
												goto l866
											l867:
												position, tokenIndex = position866, tokenIndex866
												if buffer[position] != rune('D') {
													goto l58
												}
												position++
											}
										l866:
											{
												position868, tokenIndex868 := position, tokenIndex
												if buffer[position] != rune('o') {
													goto l869
												}
												position++
												goto l868
											l869:
												position, tokenIndex = position868, tokenIndex868
												if buffer[position] != rune('O') {
													goto l58
												}
												position++
											}
										l868:
											if buffer[position] != rune(' ') {
												goto l58
											}
											position++
											{
												position870, tokenIndex870 := position, tokenIndex
												if buffer[position] != rune('t') {
													goto l871
												}
												position++
												goto l870
											l871:
												position, tokenIndex = position870, tokenIndex870
												if buffer[position] != rune('T') {
													goto l58
												}
												position++
											}
										l870:
											{
												position872, tokenIndex872 := position, tokenIndex
												if buffer[position] != rune('m') {
													goto l873
												}
												position++
												goto l872
											l873:
												position, tokenIndex = position872, tokenIndex872
												if buffer[position] != rune('M') {
													goto l58
												}
												position++
											}
										l872:
											{
												position874, tokenIndex874 := position, tokenIndex
												if buffer[position] != rune('p') {
													goto l875
												}
												position++
												goto l874
											l875:
												position, tokenIndex = position874, tokenIndex874
												if buffer[position] != rune('P') {
													goto l58
												}
												position++
											}
										l874:
											{
												position876, tokenIndex876 := position, tokenIndex
												if buffer[position] != rune('l') {
													goto l877
												}
												position++
												goto l876
											l877:
												position, tokenIndex = position876, tokenIndex876
												if buffer[position] != rune('L') {
													goto l58
												}
												position++
											}
										l876:
											break
										case 'O', 'o':
											{
												position878, tokenIndex878 := position, tokenIndex
												if buffer[position] != rune('o') {
													goto l879
												}
												position++
												goto l878
											l879:
												position, tokenIndex = position878, tokenIndex878
												if buffer[position] != rune('O') {
													goto l58
												}
												position++
											}
										l878:
											{
												position880, tokenIndex880 := position, tokenIndex
												if buffer[position] != rune('p') {
													goto l881
												}
												position++
												goto l880
											l881:
												position, tokenIndex = position880, tokenIndex880
												if buffer[position] != rune('P') {
													goto l58
												}
												position++
											}
										l880:
											{
												position882, tokenIndex882 := position, tokenIndex
												if buffer[position] != rune('e') {
													goto l883
												}
												position++
												goto l882
											l883:
												position, tokenIndex = position882, tokenIndex882
												if buffer[position] != rune('E') {
													goto l58
												}
												position++
											}
										l882:
											{
												position884, tokenIndex884 := position, tokenIndex
												if buffer[position] != rune('r') {
													goto l885
												}
												position++
												goto l884
											l885:
												position, tokenIndex = position884, tokenIndex884
												if buffer[position] != rune('R') {
													goto l58
												}
												position++
											}
										l884:
											{
												position886, tokenIndex886 := position, tokenIndex
												if buffer[position] != rune('a') {
													goto l887
												}
												position++
												goto l886
											l887:
												position, tokenIndex = position886, tokenIndex886
												if buffer[position] != rune('A') {
													goto l58
												}
												position++
											}
										l886:
											{
												position888, tokenIndex888 := position, tokenIndex
												if buffer[position] != rune('t') {
													goto l889
												}
												position++
												goto l888
											l889:
												position, tokenIndex = position888, tokenIndex888
												if buffer[position] != rune('T') {
													goto l58
												}
												position++
											}
										l888:
											{
												position890, tokenIndex890 := position, tokenIndex
												if buffer[position] != rune('o') {
													goto l891
												}
												position++
												goto l890
											l891:
												position, tokenIndex = position890, tokenIndex890
												if buffer[position] != rune('O') {
													goto l58
												}
												position++
											}
										l890:
											{
												position892, tokenIndex892 := position, tokenIndex
												if buffer[position] != rune('r') {
													goto l893
												}
												position++
												goto l892
											l893:
												position, tokenIndex = position892, tokenIndex892
												if buffer[position] != rune('R') {
													goto l58
												}
												position++
											}
										l892:
											if !_rules[rulews]() {
												goto l58
											}
											{
												position894, tokenIndex894 := position, tokenIndex
												{
													position896, tokenIndex896 := position, tokenIndex
													{
														position898, tokenIndex898 := position, tokenIndex
														if buffer[position] != rune('a') {
															goto l899
														}
														position++
														goto l898
													l899:
														position, tokenIndex = position898, tokenIndex898
														if buffer[position] != rune('A') {
															goto l897
														}
														position++
													}
												l898:
													{
														position900, tokenIndex900 := position, tokenIndex
														if buffer[position] != rune('n') {
															goto l901
														}
														position++
														goto l900
													l901:
														position, tokenIndex = position900, tokenIndex900
														if buffer[position] != rune('N') {
															goto l897
														}
														position++
													}
												l900:
													{
														position902, tokenIndex902 := position, tokenIndex
														if buffer[position] != rune('d') {
															goto l903
														}
														position++
														goto l902
													l903:
														position, tokenIndex = position902, tokenIndex902
														if buffer[position] != rune('D') {
															goto l897
														}
														position++
													}
												l902:
													{
														position904, tokenIndex904 := position, tokenIndex
														if buffer[position] != rune('a') {
															goto l905
														}
														position++
														goto l904
													l905:
														position, tokenIndex = position904, tokenIndex904
														if buffer[position] != rune('A') {
															goto l897
														}
														position++
													}
												l904:
													{
														position906, tokenIndex906 := position, tokenIndex
														if buffer[position] != rune('s') {
															goto l907
														}
														position++
														goto l906
													l907:
														position, tokenIndex = position906, tokenIndex906
														if buffer[position] != rune('S') {
															goto l897
														}
														position++
													}
												l906:
													{
														position908, tokenIndex908 := position, tokenIndex
														if buffer[position] != rune('s') {
															goto l909
														}
														position++
														goto l908
													l909:
														position, tokenIndex = position908, tokenIndex908
														if buffer[position] != rune('S') {
															goto l897
														}
														position++
													}
												l908:
													{
														position910, tokenIndex910 := position, tokenIndex
														if buffer[position] != rune('i') {
															goto l911
														}
														position++
														goto l910
													l911:
														position, tokenIndex = position910, tokenIndex910
														if buffer[position] != rune('I') {
															goto l897
														}
														position++
													}
												l910:
													{
														position912, tokenIndex912 := position, tokenIndex
														if buffer[position] != rune('g') {
															goto l913
														}
														position++
														goto l912
													l913:
														position, tokenIndex = position912, tokenIndex912
														if buffer[position] != rune('G') {
															goto l897
														}
														position++
													}
												l912:
													{
														position914, tokenIndex914 := position, tokenIndex
														if buffer[position] != rune('n') {
															goto l915
														}
														position++
														goto l914
													l915:
														position, tokenIndex = position914, tokenIndex914
														if buffer[position] != rune('N') {
															goto l897
														}
														position++
													}
												l914:
													goto l896
												l897:
													position, tokenIndex = position896, tokenIndex896
													{
														position917, tokenIndex917 := position, tokenIndex
														if buffer[position] != rune('a') {
															goto l918
														}
														position++
														goto l917
													l918:
														position, tokenIndex = position917, tokenIndex917
														if buffer[position] != rune('A') {
															goto l916
														}
														position++
													}
												l917:
													{
														position919, tokenIndex919 := position, tokenIndex
														if buffer[position] != rune('n') {
															goto l920
														}
														position++
														goto l919
													l920:
														position, tokenIndex = position919, tokenIndex919
														if buffer[position] != rune('N') {
															goto l916
														}
														position++
													}
												l919:
													{
														position921, tokenIndex921 := position, tokenIndex
														if buffer[position] != rune('d') {
															goto l922
														}
														position++
														goto l921
													l922:
														position, tokenIndex = position921, tokenIndex921
														if buffer[position] != rune('D') {
															goto l916
														}
														position++
													}
												l921:
													goto l896
												l916:
													position, tokenIndex = position896, tokenIndex896
													{
														position924, tokenIndex924 := position, tokenIndex
														if buffer[position] != rune('d') {
															goto l925
														}
														position++
														goto l924
													l925:
														position, tokenIndex = position924, tokenIndex924
														if buffer[position] != rune('D') {
															goto l923
														}
														position++
													}
												l924:
													{
														position926, tokenIndex926 := position, tokenIndex
														if buffer[position] != rune('e') {
															goto l927
														}
														position++
														goto l926
													l927:
														position, tokenIndex = position926, tokenIndex926
														if buffer[position] != rune('E') {
															goto l923
														}
														position++
													}
												l926:
													{
														position928, tokenIndex928 := position, tokenIndex
														if buffer[position] != rune('l') {
															goto l929
														}
														position++
														goto l928
													l929:
														position, tokenIndex = position928, tokenIndex928
														if buffer[position] != rune('L') {
															goto l923
														}
														position++
													}
												l928:
													{
														position930, tokenIndex930 := position, tokenIndex
														if buffer[position] != rune('e') {
															goto l931
														}
														position++
														goto l930
													l931:
														position, tokenIndex = position930, tokenIndex930
														if buffer[position] != rune('E') {
															goto l923
														}
														position++
													}
												l930:
													{
														position932, tokenIndex932 := position, tokenIndex
														if buffer[position] != rune('t') {
															goto l933
														}
														position++
														goto l932
													l933:
														position, tokenIndex = position932, tokenIndex932
														if buffer[position] != rune('T') {
															goto l923
														}
														position++
													}
												l932:
													{
														position934, tokenIndex934 := position, tokenIndex
														if buffer[position] != rune('e') {
															goto l935
														}
														position++
														goto l934
													l935:
														position, tokenIndex = position934, tokenIndex934
														if buffer[position] != rune('E') {
															goto l923
														}
														position++
													}
												l934:
													goto l896
												l923:
													position, tokenIndex = position896, tokenIndex896
													{
														position937, tokenIndex937 := position, tokenIndex
														if buffer[position] != rune('d') {
															goto l938
														}
														position++
														goto l937
													l938:
														position, tokenIndex = position937, tokenIndex937
														if buffer[position] != rune('D') {
															goto l936
														}
														position++
													}
												l937:
													{
														position939, tokenIndex939 := position, tokenIndex
														if buffer[position] != rune('e') {
															goto l940
														}
														position++
														goto l939
													l940:
														position, tokenIndex = position939, tokenIndex939
														if buffer[position] != rune('E') {
															goto l936
														}
														position++
													}
												l939:
													{
														position941, tokenIndex941 := position, tokenIndex
														if buffer[position] != rune('r') {
															goto l942
														}
														position++
														goto l941
													l942:
														position, tokenIndex = position941, tokenIndex941
														if buffer[position] != rune('R') {
															goto l936
														}
														position++
													}
												l941:
													{
														position943, tokenIndex943 := position, tokenIndex
														if buffer[position] != rune('e') {
															goto l944
														}
														position++
														goto l943
													l944:
														position, tokenIndex = position943, tokenIndex943
														if buffer[position] != rune('E') {
															goto l936
														}
														position++
													}
												l943:
													{
														position945, tokenIndex945 := position, tokenIndex
														if buffer[position] != rune('f') {
															goto l946
														}
														position++
														goto l945
													l946:
														position, tokenIndex = position945, tokenIndex945
														if buffer[position] != rune('F') {
															goto l936
														}
														position++
													}
												l945:
													goto l896
												l936:
													position, tokenIndex = position896, tokenIndex896
													{
														position948, tokenIndex948 := position, tokenIndex
														if buffer[position] != rune('d') {
															goto l949
														}
														position++
														goto l948
													l949:
														position, tokenIndex = position948, tokenIndex948
														if buffer[position] != rune('D') {
															goto l947
														}
														position++
													}
												l948:
													{
														position950, tokenIndex950 := position, tokenIndex
														if buffer[position] != rune('i') {
															goto l951
														}
														position++
														goto l950
													l951:
														position, tokenIndex = position950, tokenIndex950
														if buffer[position] != rune('I') {
															goto l947
														}
														position++
													}
												l950:
													{
														position952, tokenIndex952 := position, tokenIndex
														if buffer[position] != rune('v') {
															goto l953
														}
														position++
														goto l952
													l953:
														position, tokenIndex = position952, tokenIndex952
														if buffer[position] != rune('V') {
															goto l947
														}
														position++
													}
												l952:
													{
														position954, tokenIndex954 := position, tokenIndex
														if buffer[position] != rune('a') {
															goto l955
														}
														position++
														goto l954
													l955:
														position, tokenIndex = position954, tokenIndex954
														if buffer[position] != rune('A') {
															goto l947
														}
														position++
													}
												l954:
													{
														position956, tokenIndex956 := position, tokenIndex
														if buffer[position] != rune('s') {
															goto l957
														}
														position++
														goto l956
													l957:
														position, tokenIndex = position956, tokenIndex956
														if buffer[position] != rune('S') {
															goto l947
														}
														position++
													}
												l956:
													{
														position958, tokenIndex958 := position, tokenIndex
														if buffer[position] != rune('s') {
															goto l959
														}
														position++
														goto l958
													l959:
														position, tokenIndex = position958, tokenIndex958
														if buffer[position] != rune('S') {
															goto l947
														}
														position++
													}
												l958:
													{
														position960, tokenIndex960 := position, tokenIndex
														if buffer[position] != rune('i') {
															goto l961
														}
														position++
														goto l960
													l961:
														position, tokenIndex = position960, tokenIndex960
														if buffer[position] != rune('I') {
															goto l947
														}
														position++
													}
												l960:
													{
														position962, tokenIndex962 := position, tokenIndex
														if buffer[position] != rune('g') {
															goto l963
														}
														position++
														goto l962
													l963:
														position, tokenIndex = position962, tokenIndex962
														if buffer[position] != rune('G') {
															goto l947
														}
														position++
													}
												l962:
													{
														position964, tokenIndex964 := position, tokenIndex
														if buffer[position] != rune('n') {
															goto l965
														}
														position++
														goto l964
													l965:
														position, tokenIndex = position964, tokenIndex964
														if buffer[position] != rune('N') {
															goto l947
														}
														position++
													}
												l964:
													goto l896
												l947:
													position, tokenIndex = position896, tokenIndex896
													{
														position967, tokenIndex967 := position, tokenIndex
														if buffer[position] != rune('g') {
															goto l968
														}
														position++
														goto l967
													l968:
														position, tokenIndex = position967, tokenIndex967
														if buffer[position] != rune('G') {
															goto l966
														}
														position++
													}
												l967:
													{
														position969, tokenIndex969 := position, tokenIndex
														if buffer[position] != rune('e') {
															goto l970
														}
														position++
														goto l969
													l970:
														position, tokenIndex = position969, tokenIndex969
														if buffer[position] != rune('E') {
															goto l966
														}
														position++
													}
												l969:
													goto l896
												l966:
													position, tokenIndex = position896, tokenIndex896
													{
														position972, tokenIndex972 := position, tokenIndex
														if buffer[position] != rune('g') {
															goto l973
														}
														position++
														goto l972
													l973:
														position, tokenIndex = position972, tokenIndex972
														if buffer[position] != rune('G') {
															goto l971
														}
														position++
													}
												l972:
													{
														position974, tokenIndex974 := position, tokenIndex
														if buffer[position] != rune('e') {
															goto l975
														}
														position++
														goto l974
													l975:
														position, tokenIndex = position974, tokenIndex974
														if buffer[position] != rune('E') {
															goto l971
														}
														position++
													}
												l974:
													goto l896
												l971:
													position, tokenIndex = position896, tokenIndex896
													{
														position977, tokenIndex977 := position, tokenIndex
														if buffer[position] != rune('l') {
															goto l978
														}
														position++
														goto l977
													l978:
														position, tokenIndex = position977, tokenIndex977
														if buffer[position] != rune('L') {
															goto l976
														}
														position++
													}
												l977:
													{
														position979, tokenIndex979 := position, tokenIndex
														if buffer[position] != rune('a') {
															goto l980
														}
														position++
														goto l979
													l980:
														position, tokenIndex = position979, tokenIndex979
														if buffer[position] != rune('A') {
															goto l976
														}
														position++
													}
												l979:
													{
														position981, tokenIndex981 := position, tokenIndex
														if buffer[position] != rune('n') {
															goto l982
														}
														position++
														goto l981
													l982:
														position, tokenIndex = position981, tokenIndex981
														if buffer[position] != rune('N') {
															goto l976
														}
														position++
													}
												l981:
													{
														position983, tokenIndex983 := position, tokenIndex
														if buffer[position] != rune('d') {
															goto l984
														}
														position++
														goto l983
													l984:
														position, tokenIndex = position983, tokenIndex983
														if buffer[position] != rune('D') {
															goto l976
														}
														position++
													}
												l983:
													goto l896
												l976:
													position, tokenIndex = position896, tokenIndex896
													{
														position986, tokenIndex986 := position, tokenIndex
														if buffer[position] != rune('l') {
															goto l987
														}
														position++
														goto l986
													l987:
														position, tokenIndex = position986, tokenIndex986
														if buffer[position] != rune('L') {
															goto l985
														}
														position++
													}
												l986:
													{
														position988, tokenIndex988 := position, tokenIndex
														if buffer[position] != rune('e') {
															goto l989
														}
														position++
														goto l988
													l989:
														position, tokenIndex = position988, tokenIndex988
														if buffer[position] != rune('E') {
															goto l985
														}
														position++
													}
												l988:
													goto l896
												l985:
													position, tokenIndex = position896, tokenIndex896
													{
														position991, tokenIndex991 := position, tokenIndex
														if buffer[position] != rune('l') {
															goto l992
														}
														position++
														goto l991
													l992:
														position, tokenIndex = position991, tokenIndex991
														if buffer[position] != rune('L') {
															goto l990
														}
														position++
													}
												l991:
													{
														position993, tokenIndex993 := position, tokenIndex
														if buffer[position] != rune('n') {
															goto l994
														}
														position++
														goto l993
													l994:
														position, tokenIndex = position993, tokenIndex993
														if buffer[position] != rune('N') {
															goto l990
														}
														position++
													}
												l993:
													{
														position995, tokenIndex995 := position, tokenIndex
														if buffer[position] != rune('o') {
															goto l996
														}
														position++
														goto l995
													l996:
														position, tokenIndex = position995, tokenIndex995
														if buffer[position] != rune('O') {
															goto l990
														}
														position++
													}
												l995:
													{
														position997, tokenIndex997 := position, tokenIndex
														if buffer[position] != rune('t') {
															goto l998
														}
														position++
														goto l997
													l998:
														position, tokenIndex = position997, tokenIndex997
														if buffer[position] != rune('T') {
															goto l990
														}
														position++
													}
												l997:
													goto l896
												l990:
													position, tokenIndex = position896, tokenIndex896
													{
														position1000, tokenIndex1000 := position, tokenIndex
														if buffer[position] != rune('l') {
															goto l1001
														}
														position++
														goto l1000
													l1001:
														position, tokenIndex = position1000, tokenIndex1000
														if buffer[position] != rune('L') {
															goto l999
														}
														position++
													}
												l1000:
													{
														position1002, tokenIndex1002 := position, tokenIndex
														if buffer[position] != rune('o') {
															goto l1003
														}
														position++
														goto l1002
													l1003:
														position, tokenIndex = position1002, tokenIndex1002
														if buffer[position] != rune('O') {
															goto l999
														}
														position++
													}
												l1002:
													{
														position1004, tokenIndex1004 := position, tokenIndex
														if buffer[position] != rune('r') {
															goto l1005
														}
														position++
														goto l1004
													l1005:
														position, tokenIndex = position1004, tokenIndex1004
														if buffer[position] != rune('R') {
															goto l999
														}
														position++
													}
												l1004:
													goto l896
												l999:
													position, tokenIndex = position896, tokenIndex896
													{
														position1007, tokenIndex1007 := position, tokenIndex
														if buffer[position] != rune('l') {
															goto l1008
														}
														position++
														goto l1007
													l1008:
														position, tokenIndex = position1007, tokenIndex1007
														if buffer[position] != rune('L') {
															goto l1006
														}
														position++
													}
												l1007:
													{
														position1009, tokenIndex1009 := position, tokenIndex
														if buffer[position] != rune('s') {
															goto l1010
														}
														position++
														goto l1009
													l1010:
														position, tokenIndex = position1009, tokenIndex1009
														if buffer[position] != rune('S') {
															goto l1006
														}
														position++
													}
												l1009:
													{
														position1011, tokenIndex1011 := position, tokenIndex
														if buffer[position] != rune('h') {
															goto l1012
														}
														position++
														goto l1011
													l1012:
														position, tokenIndex = position1011, tokenIndex1011
														if buffer[position] != rune('H') {
															goto l1006
														}
														position++
													}
												l1011:
													{
														position1013, tokenIndex1013 := position, tokenIndex
														if buffer[position] != rune('i') {
															goto l1014
														}
														position++
														goto l1013
													l1014:
														position, tokenIndex = position1013, tokenIndex1013
														if buffer[position] != rune('I') {
															goto l1006
														}
														position++
													}
												l1013:
													{
														position1015, tokenIndex1015 := position, tokenIndex
														if buffer[position] != rune('f') {
															goto l1016
														}
														position++
														goto l1015
													l1016:
														position, tokenIndex = position1015, tokenIndex1015
														if buffer[position] != rune('F') {
															goto l1006
														}
														position++
													}
												l1015:
													{
														position1017, tokenIndex1017 := position, tokenIndex
														if buffer[position] != rune('t') {
															goto l1018
														}
														position++
														goto l1017
													l1018:
														position, tokenIndex = position1017, tokenIndex1017
														if buffer[position] != rune('T') {
															goto l1006
														}
														position++
													}
												l1017:
													{
														position1019, tokenIndex1019 := position, tokenIndex
														if buffer[position] != rune('a') {
															goto l1020
														}
														position++
														goto l1019
													l1020:
														position, tokenIndex = position1019, tokenIndex1019
														if buffer[position] != rune('A') {
															goto l1006
														}
														position++
													}
												l1019:
													{
														position1021, tokenIndex1021 := position, tokenIndex
														if buffer[position] != rune('s') {
															goto l1022
														}
														position++
														goto l1021
													l1022:
														position, tokenIndex = position1021, tokenIndex1021
														if buffer[position] != rune('S') {
															goto l1006
														}
														position++
													}
												l1021:
													{
														position1023, tokenIndex1023 := position, tokenIndex
														if buffer[position] != rune('s') {
															goto l1024
														}
														position++
														goto l1023
													l1024:
														position, tokenIndex = position1023, tokenIndex1023
														if buffer[position] != rune('S') {
															goto l1006
														}
														position++
													}
												l1023:
													{
														position1025, tokenIndex1025 := position, tokenIndex
														if buffer[position] != rune('i') {
															goto l1026
														}
														position++
														goto l1025
													l1026:
														position, tokenIndex = position1025, tokenIndex1025
														if buffer[position] != rune('I') {
															goto l1006
														}
														position++
													}
												l1025:
													{
														position1027, tokenIndex1027 := position, tokenIndex
														if buffer[position] != rune('g') {
															goto l1028
														}
														position++
														goto l1027
													l1028:
														position, tokenIndex = position1027, tokenIndex1027
														if buffer[position] != rune('G') {
															goto l1006
														}
														position++
													}
												l1027:
													{
														position1029, tokenIndex1029 := position, tokenIndex
														if buffer[position] != rune('n') {
															goto l1030
														}
														position++
														goto l1029
													l1030:
														position, tokenIndex = position1029, tokenIndex1029
														if buffer[position] != rune('N') {
															goto l1006
														}
														position++
													}
												l1029:
													goto l896
												l1006:
													position, tokenIndex = position896, tokenIndex896
													{
														position1032, tokenIndex1032 := position, tokenIndex
														if buffer[position] != rune('l') {
															goto l1033
														}
														position++
														goto l1032
													l1033:
														position, tokenIndex = position1032, tokenIndex1032
														if buffer[position] != rune('L') {
															goto l1031
														}
														position++
													}
												l1032:
													{
														position1034, tokenIndex1034 := position, tokenIndex
														if buffer[position] != rune('s') {
															goto l1035
														}
														position++
														goto l1034
													l1035:
														position, tokenIndex = position1034, tokenIndex1034
														if buffer[position] != rune('S') {
															goto l1031
														}
														position++
													}
												l1034:
													{
														position1036, tokenIndex1036 := position, tokenIndex
														if buffer[position] != rune('h') {
															goto l1037
														}
														position++
														goto l1036
													l1037:
														position, tokenIndex = position1036, tokenIndex1036
														if buffer[position] != rune('H') {
															goto l1031
														}
														position++
													}
												l1036:
													{
														position1038, tokenIndex1038 := position, tokenIndex
														if buffer[position] != rune('i') {
															goto l1039
														}
														position++
														goto l1038
													l1039:
														position, tokenIndex = position1038, tokenIndex1038
														if buffer[position] != rune('I') {
															goto l1031
														}
														position++
													}
												l1038:
													{
														position1040, tokenIndex1040 := position, tokenIndex
														if buffer[position] != rune('f') {
															goto l1041
														}
														position++
														goto l1040
													l1041:
														position, tokenIndex = position1040, tokenIndex1040
														if buffer[position] != rune('F') {
															goto l1031
														}
														position++
													}
												l1040:
													{
														position1042, tokenIndex1042 := position, tokenIndex
														if buffer[position] != rune('t') {
															goto l1043
														}
														position++
														goto l1042
													l1043:
														position, tokenIndex = position1042, tokenIndex1042
														if buffer[position] != rune('T') {
															goto l1031
														}
														position++
													}
												l1042:
													goto l896
												l1031:
													position, tokenIndex = position896, tokenIndex896
													{
														position1045, tokenIndex1045 := position, tokenIndex
														if buffer[position] != rune('m') {
															goto l1046
														}
														position++
														goto l1045
													l1046:
														position, tokenIndex = position1045, tokenIndex1045
														if buffer[position] != rune('M') {
															goto l1044
														}
														position++
													}
												l1045:
													{
														position1047, tokenIndex1047 := position, tokenIndex
														if buffer[position] != rune('i') {
															goto l1048
														}
														position++
														goto l1047
													l1048:
														position, tokenIndex = position1047, tokenIndex1047
														if buffer[position] != rune('I') {
															goto l1044
														}
														position++
													}
												l1047:
													{
														position1049, tokenIndex1049 := position, tokenIndex
														if buffer[position] != rune('n') {
															goto l1050
														}
														position++
														goto l1049
													l1050:
														position, tokenIndex = position1049, tokenIndex1049
														if buffer[position] != rune('N') {
															goto l1044
														}
														position++
													}
												l1049:
													{
														position1051, tokenIndex1051 := position, tokenIndex
														if buffer[position] != rune('u') {
															goto l1052
														}
														position++
														goto l1051
													l1052:
														position, tokenIndex = position1051, tokenIndex1051
														if buffer[position] != rune('U') {
															goto l1044
														}
														position++
													}
												l1051:
													{
														position1053, tokenIndex1053 := position, tokenIndex
														if buffer[position] != rune('s') {
															goto l1054
														}
														position++
														goto l1053
													l1054:
														position, tokenIndex = position1053, tokenIndex1053
														if buffer[position] != rune('S') {
															goto l1044
														}
														position++
													}
												l1053:
													{
														position1055, tokenIndex1055 := position, tokenIndex
														if buffer[position] != rune('a') {
															goto l1056
														}
														position++
														goto l1055
													l1056:
														position, tokenIndex = position1055, tokenIndex1055
														if buffer[position] != rune('A') {
															goto l1044
														}
														position++
													}
												l1055:
													{
														position1057, tokenIndex1057 := position, tokenIndex
														if buffer[position] != rune('s') {
															goto l1058
														}
														position++
														goto l1057
													l1058:
														position, tokenIndex = position1057, tokenIndex1057
														if buffer[position] != rune('S') {
															goto l1044
														}
														position++
													}
												l1057:
													{
														position1059, tokenIndex1059 := position, tokenIndex
														if buffer[position] != rune('s') {
															goto l1060
														}
														position++
														goto l1059
													l1060:
														position, tokenIndex = position1059, tokenIndex1059
														if buffer[position] != rune('S') {
															goto l1044
														}
														position++
													}
												l1059:
													{
														position1061, tokenIndex1061 := position, tokenIndex
														if buffer[position] != rune('i') {
															goto l1062
														}
														position++
														goto l1061
													l1062:
														position, tokenIndex = position1061, tokenIndex1061
														if buffer[position] != rune('I') {
															goto l1044
														}
														position++
													}
												l1061:
													{
														position1063, tokenIndex1063 := position, tokenIndex
														if buffer[position] != rune('g') {
															goto l1064
														}
														position++
														goto l1063
													l1064:
														position, tokenIndex = position1063, tokenIndex1063
														if buffer[position] != rune('G') {
															goto l1044
														}
														position++
													}
												l1063:
													{
														position1065, tokenIndex1065 := position, tokenIndex
														if buffer[position] != rune('n') {
															goto l1066
														}
														position++
														goto l1065
													l1066:
														position, tokenIndex = position1065, tokenIndex1065
														if buffer[position] != rune('N') {
															goto l1044
														}
														position++
													}
												l1065:
													goto l896
												l1044:
													position, tokenIndex = position896, tokenIndex896
													{
														position1068, tokenIndex1068 := position, tokenIndex
														if buffer[position] != rune('m') {
															goto l1069
														}
														position++
														goto l1068
													l1069:
														position, tokenIndex = position1068, tokenIndex1068
														if buffer[position] != rune('M') {
															goto l1067
														}
														position++
													}
												l1068:
													{
														position1070, tokenIndex1070 := position, tokenIndex
														if buffer[position] != rune('i') {
															goto l1071
														}
														position++
														goto l1070
													l1071:
														position, tokenIndex = position1070, tokenIndex1070
														if buffer[position] != rune('I') {
															goto l1067
														}
														position++
													}
												l1070:
													{
														position1072, tokenIndex1072 := position, tokenIndex
														if buffer[position] != rune('n') {
															goto l1073
														}
														position++
														goto l1072
													l1073:
														position, tokenIndex = position1072, tokenIndex1072
														if buffer[position] != rune('N') {
															goto l1067
														}
														position++
													}
												l1072:
													{
														position1074, tokenIndex1074 := position, tokenIndex
														if buffer[position] != rune('u') {
															goto l1075
														}
														position++
														goto l1074
													l1075:
														position, tokenIndex = position1074, tokenIndex1074
														if buffer[position] != rune('U') {
															goto l1067
														}
														position++
													}
												l1074:
													{
														position1076, tokenIndex1076 := position, tokenIndex
														if buffer[position] != rune('s') {
															goto l1077
														}
														position++
														goto l1076
													l1077:
														position, tokenIndex = position1076, tokenIndex1076
														if buffer[position] != rune('S') {
															goto l1067
														}
														position++
													}
												l1076:
													goto l896
												l1067:
													position, tokenIndex = position896, tokenIndex896
													{
														position1079, tokenIndex1079 := position, tokenIndex
														if buffer[position] != rune('m') {
															goto l1080
														}
														position++
														goto l1079
													l1080:
														position, tokenIndex = position1079, tokenIndex1079
														if buffer[position] != rune('M') {
															goto l1078
														}
														position++
													}
												l1079:
													{
														position1081, tokenIndex1081 := position, tokenIndex
														if buffer[position] != rune('u') {
															goto l1082
														}
														position++
														goto l1081
													l1082:
														position, tokenIndex = position1081, tokenIndex1081
														if buffer[position] != rune('U') {
															goto l1078
														}
														position++
													}
												l1081:
													{
														position1083, tokenIndex1083 := position, tokenIndex
														if buffer[position] != rune('l') {
															goto l1084
														}
														position++
														goto l1083
													l1084:
														position, tokenIndex = position1083, tokenIndex1083
														if buffer[position] != rune('L') {
															goto l1078
														}
														position++
													}
												l1083:
													{
														position1085, tokenIndex1085 := position, tokenIndex
														if buffer[position] != rune('t') {
															goto l1086
														}
														position++
														goto l1085
													l1086:
														position, tokenIndex = position1085, tokenIndex1085
														if buffer[position] != rune('T') {
															goto l1078
														}
														position++
													}
												l1085:
													{
														position1087, tokenIndex1087 := position, tokenIndex
														if buffer[position] != rune('a') {
															goto l1088
														}
														position++
														goto l1087
													l1088:
														position, tokenIndex = position1087, tokenIndex1087
														if buffer[position] != rune('A') {
															goto l1078
														}
														position++
													}
												l1087:
													{
														position1089, tokenIndex1089 := position, tokenIndex
														if buffer[position] != rune('s') {
															goto l1090
														}
														position++
														goto l1089
													l1090:
														position, tokenIndex = position1089, tokenIndex1089
														if buffer[position] != rune('S') {
															goto l1078
														}
														position++
													}
												l1089:
													{
														position1091, tokenIndex1091 := position, tokenIndex
														if buffer[position] != rune('s') {
															goto l1092
														}
														position++
														goto l1091
													l1092:
														position, tokenIndex = position1091, tokenIndex1091
														if buffer[position] != rune('S') {
															goto l1078
														}
														position++
													}
												l1091:
													{
														position1093, tokenIndex1093 := position, tokenIndex
														if buffer[position] != rune('i') {
															goto l1094
														}
														position++
														goto l1093
													l1094:
														position, tokenIndex = position1093, tokenIndex1093
														if buffer[position] != rune('I') {
															goto l1078
														}
														position++
													}
												l1093:
													{
														position1095, tokenIndex1095 := position, tokenIndex
														if buffer[position] != rune('g') {
															goto l1096
														}
														position++
														goto l1095
													l1096:
														position, tokenIndex = position1095, tokenIndex1095
														if buffer[position] != rune('G') {
															goto l1078
														}
														position++
													}
												l1095:
													{
														position1097, tokenIndex1097 := position, tokenIndex
														if buffer[position] != rune('n') {
															goto l1098
														}
														position++
														goto l1097
													l1098:
														position, tokenIndex = position1097, tokenIndex1097
														if buffer[position] != rune('N') {
															goto l1078
														}
														position++
													}
												l1097:
													goto l896
												l1078:
													position, tokenIndex = position896, tokenIndex896
													{
														position1100, tokenIndex1100 := position, tokenIndex
														if buffer[position] != rune('n') {
															goto l1101
														}
														position++
														goto l1100
													l1101:
														position, tokenIndex = position1100, tokenIndex1100
														if buffer[position] != rune('N') {
															goto l1099
														}
														position++
													}
												l1100:
													{
														position1102, tokenIndex1102 := position, tokenIndex
														if buffer[position] != rune('e') {
															goto l1103
														}
														position++
														goto l1102
													l1103:
														position, tokenIndex = position1102, tokenIndex1102
														if buffer[position] != rune('E') {
															goto l1099
														}
														position++
													}
												l1102:
													{
														position1104, tokenIndex1104 := position, tokenIndex
														if buffer[position] != rune('g') {
															goto l1105
														}
														position++
														goto l1104
													l1105:
														position, tokenIndex = position1104, tokenIndex1104
														if buffer[position] != rune('G') {
															goto l1099
														}
														position++
													}
												l1104:
													goto l896
												l1099:
													position, tokenIndex = position896, tokenIndex896
													{
														position1107, tokenIndex1107 := position, tokenIndex
														if buffer[position] != rune('n') {
															goto l1108
														}
														position++
														goto l1107
													l1108:
														position, tokenIndex = position1107, tokenIndex1107
														if buffer[position] != rune('N') {
															goto l1106
														}
														position++
													}
												l1107:
													{
														position1109, tokenIndex1109 := position, tokenIndex
														if buffer[position] != rune('e') {
															goto l1110
														}
														position++
														goto l1109
													l1110:
														position, tokenIndex = position1109, tokenIndex1109
														if buffer[position] != rune('E') {
															goto l1106
														}
														position++
													}
												l1109:
													{
														position1111, tokenIndex1111 := position, tokenIndex
														if buffer[position] != rune('w') {
															goto l1112
														}
														position++
														goto l1111
													l1112:
														position, tokenIndex = position1111, tokenIndex1111
														if buffer[position] != rune('W') {
															goto l1106
														}
														position++
													}
												l1111:
													goto l896
												l1106:
													position, tokenIndex = position896, tokenIndex896
													{
														position1114, tokenIndex1114 := position, tokenIndex
														if buffer[position] != rune('n') {
															goto l1115
														}
														position++
														goto l1114
													l1115:
														position, tokenIndex = position1114, tokenIndex1114
														if buffer[position] != rune('N') {
															goto l1113
														}
														position++
													}
												l1114:
													{
														position1116, tokenIndex1116 := position, tokenIndex
														if buffer[position] != rune('e') {
															goto l1117
														}
														position++
														goto l1116
													l1117:
														position, tokenIndex = position1116, tokenIndex1116
														if buffer[position] != rune('E') {
															goto l1113
														}
														position++
													}
												l1116:
													goto l896
												l1113:
													position, tokenIndex = position896, tokenIndex896
													{
														position1119, tokenIndex1119 := position, tokenIndex
														if buffer[position] != rune('o') {
															goto l1120
														}
														position++
														goto l1119
													l1120:
														position, tokenIndex = position1119, tokenIndex1119
														if buffer[position] != rune('O') {
															goto l1118
														}
														position++
													}
												l1119:
													{
														position1121, tokenIndex1121 := position, tokenIndex
														if buffer[position] != rune('r') {
															goto l1122
														}
														position++
														goto l1121
													l1122:
														position, tokenIndex = position1121, tokenIndex1121
														if buffer[position] != rune('R') {
															goto l1118
														}
														position++
													}
												l1121:
													{
														position1123, tokenIndex1123 := position, tokenIndex
														if buffer[position] != rune('a') {
															goto l1124
														}
														position++
														goto l1123
													l1124:
														position, tokenIndex = position1123, tokenIndex1123
														if buffer[position] != rune('A') {
															goto l1118
														}
														position++
													}
												l1123:
													{
														position1125, tokenIndex1125 := position, tokenIndex
														if buffer[position] != rune('s') {
															goto l1126
														}
														position++
														goto l1125
													l1126:
														position, tokenIndex = position1125, tokenIndex1125
														if buffer[position] != rune('S') {
															goto l1118
														}
														position++
													}
												l1125:
													{
														position1127, tokenIndex1127 := position, tokenIndex
														if buffer[position] != rune('s') {
															goto l1128
														}
														position++
														goto l1127
													l1128:
														position, tokenIndex = position1127, tokenIndex1127
														if buffer[position] != rune('S') {
															goto l1118
														}
														position++
													}
												l1127:
													{
														position1129, tokenIndex1129 := position, tokenIndex
														if buffer[position] != rune('i') {
															goto l1130
														}
														position++
														goto l1129
													l1130:
														position, tokenIndex = position1129, tokenIndex1129
														if buffer[position] != rune('I') {
															goto l1118
														}
														position++
													}
												l1129:
													{
														position1131, tokenIndex1131 := position, tokenIndex
														if buffer[position] != rune('g') {
															goto l1132
														}
														position++
														goto l1131
													l1132:
														position, tokenIndex = position1131, tokenIndex1131
														if buffer[position] != rune('G') {
															goto l1118
														}
														position++
													}
												l1131:
													{
														position1133, tokenIndex1133 := position, tokenIndex
														if buffer[position] != rune('n') {
															goto l1134
														}
														position++
														goto l1133
													l1134:
														position, tokenIndex = position1133, tokenIndex1133
														if buffer[position] != rune('N') {
															goto l1118
														}
														position++
													}
												l1133:
													goto l896
												l1118:
													position, tokenIndex = position896, tokenIndex896
													{
														position1136, tokenIndex1136 := position, tokenIndex
														if buffer[position] != rune('p') {
															goto l1137
														}
														position++
														goto l1136
													l1137:
														position, tokenIndex = position1136, tokenIndex1136
														if buffer[position] != rune('P') {
															goto l1135
														}
														position++
													}
												l1136:
													{
														position1138, tokenIndex1138 := position, tokenIndex
														if buffer[position] != rune('l') {
															goto l1139
														}
														position++
														goto l1138
													l1139:
														position, tokenIndex = position1138, tokenIndex1138
														if buffer[position] != rune('L') {
															goto l1135
														}
														position++
													}
												l1138:
													{
														position1140, tokenIndex1140 := position, tokenIndex
														if buffer[position] != rune('u') {
															goto l1141
														}
														position++
														goto l1140
													l1141:
														position, tokenIndex = position1140, tokenIndex1140
														if buffer[position] != rune('U') {
															goto l1135
														}
														position++
													}
												l1140:
													{
														position1142, tokenIndex1142 := position, tokenIndex
														if buffer[position] != rune('s') {
															goto l1143
														}
														position++
														goto l1142
													l1143:
														position, tokenIndex = position1142, tokenIndex1142
														if buffer[position] != rune('S') {
															goto l1135
														}
														position++
													}
												l1142:
													{
														position1144, tokenIndex1144 := position, tokenIndex
														if buffer[position] != rune('a') {
															goto l1145
														}
														position++
														goto l1144
													l1145:
														position, tokenIndex = position1144, tokenIndex1144
														if buffer[position] != rune('A') {
															goto l1135
														}
														position++
													}
												l1144:
													{
														position1146, tokenIndex1146 := position, tokenIndex
														if buffer[position] != rune('s') {
															goto l1147
														}
														position++
														goto l1146
													l1147:
														position, tokenIndex = position1146, tokenIndex1146
														if buffer[position] != rune('S') {
															goto l1135
														}
														position++
													}
												l1146:
													{
														position1148, tokenIndex1148 := position, tokenIndex
														if buffer[position] != rune('s') {
															goto l1149
														}
														position++
														goto l1148
													l1149:
														position, tokenIndex = position1148, tokenIndex1148
														if buffer[position] != rune('S') {
															goto l1135
														}
														position++
													}
												l1148:
													{
														position1150, tokenIndex1150 := position, tokenIndex
														if buffer[position] != rune('i') {
															goto l1151
														}
														position++
														goto l1150
													l1151:
														position, tokenIndex = position1150, tokenIndex1150
														if buffer[position] != rune('I') {
															goto l1135
														}
														position++
													}
												l1150:
													{
														position1152, tokenIndex1152 := position, tokenIndex
														if buffer[position] != rune('g') {
															goto l1153
														}
														position++
														goto l1152
													l1153:
														position, tokenIndex = position1152, tokenIndex1152
														if buffer[position] != rune('G') {
															goto l1135
														}
														position++
													}
												l1152:
													{
														position1154, tokenIndex1154 := position, tokenIndex
														if buffer[position] != rune('n') {
															goto l1155
														}
														position++
														goto l1154
													l1155:
														position, tokenIndex = position1154, tokenIndex1154
														if buffer[position] != rune('N') {
															goto l1135
														}
														position++
													}
												l1154:
													goto l896
												l1135:
													position, tokenIndex = position896, tokenIndex896
													{
														position1157, tokenIndex1157 := position, tokenIndex
														if buffer[position] != rune('p') {
															goto l1158
														}
														position++
														goto l1157
													l1158:
														position, tokenIndex = position1157, tokenIndex1157
														if buffer[position] != rune('P') {
															goto l1156
														}
														position++
													}
												l1157:
													{
														position1159, tokenIndex1159 := position, tokenIndex
														if buffer[position] != rune('l') {
															goto l1160
														}
														position++
														goto l1159
													l1160:
														position, tokenIndex = position1159, tokenIndex1159
														if buffer[position] != rune('L') {
															goto l1156
														}
														position++
													}
												l1159:
													{
														position1161, tokenIndex1161 := position, tokenIndex
														if buffer[position] != rune('u') {
															goto l1162
														}
														position++
														goto l1161
													l1162:
														position, tokenIndex = position1161, tokenIndex1161
														if buffer[position] != rune('U') {
															goto l1156
														}
														position++
													}
												l1161:
													{
														position1163, tokenIndex1163 := position, tokenIndex
														if buffer[position] != rune('s') {
															goto l1164
														}
														position++
														goto l1163
													l1164:
														position, tokenIndex = position1163, tokenIndex1163
														if buffer[position] != rune('S') {
															goto l1156
														}
														position++
													}
												l1163:
													goto l896
												l1156:
													position, tokenIndex = position896, tokenIndex896
													{
														position1166, tokenIndex1166 := position, tokenIndex
														if buffer[position] != rune('p') {
															goto l1167
														}
														position++
														goto l1166
													l1167:
														position, tokenIndex = position1166, tokenIndex1166
														if buffer[position] != rune('P') {
															goto l1165
														}
														position++
													}
												l1166:
													{
														position1168, tokenIndex1168 := position, tokenIndex
														if buffer[position] != rune('o') {
															goto l1169
														}
														position++
														goto l1168
													l1169:
														position, tokenIndex = position1168, tokenIndex1168
														if buffer[position] != rune('O') {
															goto l1165
														}
														position++
													}
												l1168:
													{
														position1170, tokenIndex1170 := position, tokenIndex
														if buffer[position] != rune('s') {
															goto l1171
														}
														position++
														goto l1170
													l1171:
														position, tokenIndex = position1170, tokenIndex1170
														if buffer[position] != rune('S') {
															goto l1165
														}
														position++
													}
												l1170:
													{
														position1172, tokenIndex1172 := position, tokenIndex
														if buffer[position] != rune('t') {
															goto l1173
														}
														position++
														goto l1172
													l1173:
														position, tokenIndex = position1172, tokenIndex1172
														if buffer[position] != rune('T') {
															goto l1165
														}
														position++
													}
												l1172:
													{
														position1174, tokenIndex1174 := position, tokenIndex
														if buffer[position] != rune('d') {
															goto l1175
														}
														position++
														goto l1174
													l1175:
														position, tokenIndex = position1174, tokenIndex1174
														if buffer[position] != rune('D') {
															goto l1165
														}
														position++
													}
												l1174:
													{
														position1176, tokenIndex1176 := position, tokenIndex
														if buffer[position] != rune('e') {
															goto l1177
														}
														position++
														goto l1176
													l1177:
														position, tokenIndex = position1176, tokenIndex1176
														if buffer[position] != rune('E') {
															goto l1165
														}
														position++
													}
												l1176:
													{
														position1178, tokenIndex1178 := position, tokenIndex
														if buffer[position] != rune('c') {
															goto l1179
														}
														position++
														goto l1178
													l1179:
														position, tokenIndex = position1178, tokenIndex1178
														if buffer[position] != rune('C') {
															goto l1165
														}
														position++
													}
												l1178:
													goto l896
												l1165:
													position, tokenIndex = position896, tokenIndex896
													{
														position1181, tokenIndex1181 := position, tokenIndex
														if buffer[position] != rune('p') {
															goto l1182
														}
														position++
														goto l1181
													l1182:
														position, tokenIndex = position1181, tokenIndex1181
														if buffer[position] != rune('P') {
															goto l1180
														}
														position++
													}
												l1181:
													{
														position1183, tokenIndex1183 := position, tokenIndex
														if buffer[position] != rune('o') {
															goto l1184
														}
														position++
														goto l1183
													l1184:
														position, tokenIndex = position1183, tokenIndex1183
														if buffer[position] != rune('O') {
															goto l1180
														}
														position++
													}
												l1183:
													{
														position1185, tokenIndex1185 := position, tokenIndex
														if buffer[position] != rune('s') {
															goto l1186
														}
														position++
														goto l1185
													l1186:
														position, tokenIndex = position1185, tokenIndex1185
														if buffer[position] != rune('S') {
															goto l1180
														}
														position++
													}
												l1185:
													{
														position1187, tokenIndex1187 := position, tokenIndex
														if buffer[position] != rune('t') {
															goto l1188
														}
														position++
														goto l1187
													l1188:
														position, tokenIndex = position1187, tokenIndex1187
														if buffer[position] != rune('T') {
															goto l1180
														}
														position++
													}
												l1187:
													{
														position1189, tokenIndex1189 := position, tokenIndex
														if buffer[position] != rune('i') {
															goto l1190
														}
														position++
														goto l1189
													l1190:
														position, tokenIndex = position1189, tokenIndex1189
														if buffer[position] != rune('I') {
															goto l1180
														}
														position++
													}
												l1189:
													{
														position1191, tokenIndex1191 := position, tokenIndex
														if buffer[position] != rune('n') {
															goto l1192
														}
														position++
														goto l1191
													l1192:
														position, tokenIndex = position1191, tokenIndex1191
														if buffer[position] != rune('N') {
															goto l1180
														}
														position++
													}
												l1191:
													{
														position1193, tokenIndex1193 := position, tokenIndex
														if buffer[position] != rune('c') {
															goto l1194
														}
														position++
														goto l1193
													l1194:
														position, tokenIndex = position1193, tokenIndex1193
														if buffer[position] != rune('C') {
															goto l1180
														}
														position++
													}
												l1193:
													goto l896
												l1180:
													position, tokenIndex = position896, tokenIndex896
													{
														position1196, tokenIndex1196 := position, tokenIndex
														if buffer[position] != rune('p') {
															goto l1197
														}
														position++
														goto l1196
													l1197:
														position, tokenIndex = position1196, tokenIndex1196
														if buffer[position] != rune('P') {
															goto l1195
														}
														position++
													}
												l1196:
													{
														position1198, tokenIndex1198 := position, tokenIndex
														if buffer[position] != rune('o') {
															goto l1199
														}
														position++
														goto l1198
													l1199:
														position, tokenIndex = position1198, tokenIndex1198
														if buffer[position] != rune('O') {
															goto l1195
														}
														position++
													}
												l1198:
													{
														position1200, tokenIndex1200 := position, tokenIndex
														if buffer[position] != rune('s') {
															goto l1201
														}
														position++
														goto l1200
													l1201:
														position, tokenIndex = position1200, tokenIndex1200
														if buffer[position] != rune('S') {
															goto l1195
														}
														position++
													}
												l1200:
													goto l896
												l1195:
													position, tokenIndex = position896, tokenIndex896
													{
														position1203, tokenIndex1203 := position, tokenIndex
														if buffer[position] != rune('p') {
															goto l1204
														}
														position++
														goto l1203
													l1204:
														position, tokenIndex = position1203, tokenIndex1203
														if buffer[position] != rune('P') {
															goto l1202
														}
														position++
													}
												l1203:
													{
														position1205, tokenIndex1205 := position, tokenIndex
														if buffer[position] != rune('r') {
															goto l1206
														}
														position++
														goto l1205
													l1206:
														position, tokenIndex = position1205, tokenIndex1205
														if buffer[position] != rune('R') {
															goto l1202
														}
														position++
													}
												l1205:
													{
														position1207, tokenIndex1207 := position, tokenIndex
														if buffer[position] != rune('e') {
															goto l1208
														}
														position++
														goto l1207
													l1208:
														position, tokenIndex = position1207, tokenIndex1207
														if buffer[position] != rune('E') {
															goto l1202
														}
														position++
													}
												l1207:
													{
														position1209, tokenIndex1209 := position, tokenIndex
														if buffer[position] != rune('d') {
															goto l1210
														}
														position++
														goto l1209
													l1210:
														position, tokenIndex = position1209, tokenIndex1209
														if buffer[position] != rune('D') {
															goto l1202
														}
														position++
													}
												l1209:
													{
														position1211, tokenIndex1211 := position, tokenIndex
														if buffer[position] != rune('e') {
															goto l1212
														}
														position++
														goto l1211
													l1212:
														position, tokenIndex = position1211, tokenIndex1211
														if buffer[position] != rune('E') {
															goto l1202
														}
														position++
													}
												l1211:
													{
														position1213, tokenIndex1213 := position, tokenIndex
														if buffer[position] != rune('c') {
															goto l1214
														}
														position++
														goto l1213
													l1214:
														position, tokenIndex = position1213, tokenIndex1213
														if buffer[position] != rune('C') {
															goto l1202
														}
														position++
													}
												l1213:
													goto l896
												l1202:
													position, tokenIndex = position896, tokenIndex896
													{
														position1216, tokenIndex1216 := position, tokenIndex
														if buffer[position] != rune('r') {
															goto l1217
														}
														position++
														goto l1216
													l1217:
														position, tokenIndex = position1216, tokenIndex1216
														if buffer[position] != rune('R') {
															goto l1215
														}
														position++
													}
												l1216:
													{
														position1218, tokenIndex1218 := position, tokenIndex
														if buffer[position] != rune('e') {
															goto l1219
														}
														position++
														goto l1218
													l1219:
														position, tokenIndex = position1218, tokenIndex1218
														if buffer[position] != rune('E') {
															goto l1215
														}
														position++
													}
												l1218:
													{
														position1220, tokenIndex1220 := position, tokenIndex
														if buffer[position] != rune('f') {
															goto l1221
														}
														position++
														goto l1220
													l1221:
														position, tokenIndex = position1220, tokenIndex1220
														if buffer[position] != rune('F') {
															goto l1215
														}
														position++
													}
												l1220:
													goto l896
												l1215:
													position, tokenIndex = position896, tokenIndex896
													{
														position1223, tokenIndex1223 := position, tokenIndex
														if buffer[position] != rune('r') {
															goto l1224
														}
														position++
														goto l1223
													l1224:
														position, tokenIndex = position1223, tokenIndex1223
														if buffer[position] != rune('R') {
															goto l1222
														}
														position++
													}
												l1223:
													{
														position1225, tokenIndex1225 := position, tokenIndex
														if buffer[position] != rune('s') {
															goto l1226
														}
														position++
														goto l1225
													l1226:
														position, tokenIndex = position1225, tokenIndex1225
														if buffer[position] != rune('S') {
															goto l1222
														}
														position++
													}
												l1225:
													{
														position1227, tokenIndex1227 := position, tokenIndex
														if buffer[position] != rune('h') {
															goto l1228
														}
														position++
														goto l1227
													l1228:
														position, tokenIndex = position1227, tokenIndex1227
														if buffer[position] != rune('H') {
															goto l1222
														}
														position++
													}
												l1227:
													{
														position1229, tokenIndex1229 := position, tokenIndex
														if buffer[position] != rune('i') {
															goto l1230
														}
														position++
														goto l1229
													l1230:
														position, tokenIndex = position1229, tokenIndex1229
														if buffer[position] != rune('I') {
															goto l1222
														}
														position++
													}
												l1229:
													{
														position1231, tokenIndex1231 := position, tokenIndex
														if buffer[position] != rune('f') {
															goto l1232
														}
														position++
														goto l1231
													l1232:
														position, tokenIndex = position1231, tokenIndex1231
														if buffer[position] != rune('F') {
															goto l1222
														}
														position++
													}
												l1231:
													{
														position1233, tokenIndex1233 := position, tokenIndex
														if buffer[position] != rune('t') {
															goto l1234
														}
														position++
														goto l1233
													l1234:
														position, tokenIndex = position1233, tokenIndex1233
														if buffer[position] != rune('T') {
															goto l1222
														}
														position++
													}
												l1233:
													{
														position1235, tokenIndex1235 := position, tokenIndex
														if buffer[position] != rune('a') {
															goto l1236
														}
														position++
														goto l1235
													l1236:
														position, tokenIndex = position1235, tokenIndex1235
														if buffer[position] != rune('A') {
															goto l1222
														}
														position++
													}
												l1235:
													{
														position1237, tokenIndex1237 := position, tokenIndex
														if buffer[position] != rune('s') {
															goto l1238
														}
														position++
														goto l1237
													l1238:
														position, tokenIndex = position1237, tokenIndex1237
														if buffer[position] != rune('S') {
															goto l1222
														}
														position++
													}
												l1237:
													{
														position1239, tokenIndex1239 := position, tokenIndex
														if buffer[position] != rune('s') {
															goto l1240
														}
														position++
														goto l1239
													l1240:
														position, tokenIndex = position1239, tokenIndex1239
														if buffer[position] != rune('S') {
															goto l1222
														}
														position++
													}
												l1239:
													{
														position1241, tokenIndex1241 := position, tokenIndex
														if buffer[position] != rune('i') {
															goto l1242
														}
														position++
														goto l1241
													l1242:
														position, tokenIndex = position1241, tokenIndex1241
														if buffer[position] != rune('I') {
															goto l1222
														}
														position++
													}
												l1241:
													{
														position1243, tokenIndex1243 := position, tokenIndex
														if buffer[position] != rune('g') {
															goto l1244
														}
														position++
														goto l1243
													l1244:
														position, tokenIndex = position1243, tokenIndex1243
														if buffer[position] != rune('G') {
															goto l1222
														}
														position++
													}
												l1243:
													{
														position1245, tokenIndex1245 := position, tokenIndex
														if buffer[position] != rune('n') {
															goto l1246
														}
														position++
														goto l1245
													l1246:
														position, tokenIndex = position1245, tokenIndex1245
														if buffer[position] != rune('N') {
															goto l1222
														}
														position++
													}
												l1245:
													goto l896
												l1222:
													position, tokenIndex = position896, tokenIndex896
													{
														position1248, tokenIndex1248 := position, tokenIndex
														if buffer[position] != rune('v') {
															goto l1249
														}
														position++
														goto l1248
													l1249:
														position, tokenIndex = position1248, tokenIndex1248
														if buffer[position] != rune('V') {
															goto l1247
														}
														position++
													}
												l1248:
													{
														position1250, tokenIndex1250 := position, tokenIndex
														if buffer[position] != rune('e') {
															goto l1251
														}
														position++
														goto l1250
													l1251:
														position, tokenIndex = position1250, tokenIndex1250
														if buffer[position] != rune('E') {
															goto l1247
														}
														position++
													}
												l1250:
													{
														position1252, tokenIndex1252 := position, tokenIndex
														if buffer[position] != rune('c') {
															goto l1253
														}
														position++
														goto l1252
													l1253:
														position, tokenIndex = position1252, tokenIndex1252
														if buffer[position] != rune('C') {
															goto l1247
														}
														position++
													}
												l1252:
													{
														position1254, tokenIndex1254 := position, tokenIndex
														if buffer[position] != rune('d') {
															goto l1255
														}
														position++
														goto l1254
													l1255:
														position, tokenIndex = position1254, tokenIndex1254
														if buffer[position] != rune('D') {
															goto l1247
														}
														position++
													}
												l1254:
													{
														position1256, tokenIndex1256 := position, tokenIndex
														if buffer[position] != rune('e') {
															goto l1257
														}
														position++
														goto l1256
													l1257:
														position, tokenIndex = position1256, tokenIndex1256
														if buffer[position] != rune('E') {
															goto l1247
														}
														position++
													}
												l1256:
													{
														position1258, tokenIndex1258 := position, tokenIndex
														if buffer[position] != rune('l') {
															goto l1259
														}
														position++
														goto l1258
													l1259:
														position, tokenIndex = position1258, tokenIndex1258
														if buffer[position] != rune('L') {
															goto l1247
														}
														position++
													}
												l1258:
													{
														position1260, tokenIndex1260 := position, tokenIndex
														if buffer[position] != rune('e') {
															goto l1261
														}
														position++
														goto l1260
													l1261:
														position, tokenIndex = position1260, tokenIndex1260
														if buffer[position] != rune('E') {
															goto l1247
														}
														position++
													}
												l1260:
													{
														position1262, tokenIndex1262 := position, tokenIndex
														if buffer[position] != rune('t') {
															goto l1263
														}
														position++
														goto l1262
													l1263:
														position, tokenIndex = position1262, tokenIndex1262
														if buffer[position] != rune('T') {
															goto l1247
														}
														position++
													}
												l1262:
													{
														position1264, tokenIndex1264 := position, tokenIndex
														if buffer[position] != rune('e') {
															goto l1265
														}
														position++
														goto l1264
													l1265:
														position, tokenIndex = position1264, tokenIndex1264
														if buffer[position] != rune('E') {
															goto l1247
														}
														position++
													}
												l1264:
													goto l896
												l1247:
													position, tokenIndex = position896, tokenIndex896
													{
														position1267, tokenIndex1267 := position, tokenIndex
														if buffer[position] != rune('x') {
															goto l1268
														}
														position++
														goto l1267
													l1268:
														position, tokenIndex = position1267, tokenIndex1267
														if buffer[position] != rune('X') {
															goto l1266
														}
														position++
													}
												l1267:
													{
														position1269, tokenIndex1269 := position, tokenIndex
														if buffer[position] != rune('o') {
															goto l1270
														}
														position++
														goto l1269
													l1270:
														position, tokenIndex = position1269, tokenIndex1269
														if buffer[position] != rune('O') {
															goto l1266
														}
														position++
													}
												l1269:
													{
														position1271, tokenIndex1271 := position, tokenIndex
														if buffer[position] != rune('r') {
															goto l1272
														}
														position++
														goto l1271
													l1272:
														position, tokenIndex = position1271, tokenIndex1271
														if buffer[position] != rune('R') {
															goto l1266
														}
														position++
													}
												l1271:
													{
														position1273, tokenIndex1273 := position, tokenIndex
														if buffer[position] != rune('a') {
															goto l1274
														}
														position++
														goto l1273
													l1274:
														position, tokenIndex = position1273, tokenIndex1273
														if buffer[position] != rune('A') {
															goto l1266
														}
														position++
													}
												l1273:
													{
														position1275, tokenIndex1275 := position, tokenIndex
														if buffer[position] != rune('s') {
															goto l1276
														}
														position++
														goto l1275
													l1276:
														position, tokenIndex = position1275, tokenIndex1275
														if buffer[position] != rune('S') {
															goto l1266
														}
														position++
													}
												l1275:
													{
														position1277, tokenIndex1277 := position, tokenIndex
														if buffer[position] != rune('s') {
															goto l1278
														}
														position++
														goto l1277
													l1278:
														position, tokenIndex = position1277, tokenIndex1277
														if buffer[position] != rune('S') {
															goto l1266
														}
														position++
													}
												l1277:
													{
														position1279, tokenIndex1279 := position, tokenIndex
														if buffer[position] != rune('i') {
															goto l1280
														}
														position++
														goto l1279
													l1280:
														position, tokenIndex = position1279, tokenIndex1279
														if buffer[position] != rune('I') {
															goto l1266
														}
														position++
													}
												l1279:
													{
														position1281, tokenIndex1281 := position, tokenIndex
														if buffer[position] != rune('g') {
															goto l1282
														}
														position++
														goto l1281
													l1282:
														position, tokenIndex = position1281, tokenIndex1281
														if buffer[position] != rune('G') {
															goto l1266
														}
														position++
													}
												l1281:
													{
														position1283, tokenIndex1283 := position, tokenIndex
														if buffer[position] != rune('n') {
															goto l1284
														}
														position++
														goto l1283
													l1284:
														position, tokenIndex = position1283, tokenIndex1283
														if buffer[position] != rune('N') {
															goto l1266
														}
														position++
													}
												l1283:
													goto l896
												l1266:
													position, tokenIndex = position896, tokenIndex896
													{
														switch buffer[position] {
														case 'X', 'x':
															{
																position1286, tokenIndex1286 := position, tokenIndex
																if buffer[position] != rune('x') {
																	goto l1287
																}
																position++
																goto l1286
															l1287:
																position, tokenIndex = position1286, tokenIndex1286
																if buffer[position] != rune('X') {
																	goto l894
																}
																position++
															}
														l1286:
															{
																position1288, tokenIndex1288 := position, tokenIndex
																if buffer[position] != rune('o') {
																	goto l1289
																}
																position++
																goto l1288
															l1289:
																position, tokenIndex = position1288, tokenIndex1288
																if buffer[position] != rune('O') {
																	goto l894
																}
																position++
															}
														l1288:
															{
																position1290, tokenIndex1290 := position, tokenIndex
																if buffer[position] != rune('r') {
																	goto l1291
																}
																position++
																goto l1290
															l1291:
																position, tokenIndex = position1290, tokenIndex1290
																if buffer[position] != rune('R') {
																	goto l894
																}
																position++
															}
														l1290:
															break
														case 'V', 'v':
															{
																position1292, tokenIndex1292 := position, tokenIndex
																if buffer[position] != rune('v') {
																	goto l1293
																}
																position++
																goto l1292
															l1293:
																position, tokenIndex = position1292, tokenIndex1292
																if buffer[position] != rune('V') {
																	goto l894
																}
																position++
															}
														l1292:
															{
																position1294, tokenIndex1294 := position, tokenIndex
																if buffer[position] != rune('e') {
																	goto l1295
																}
																position++
																goto l1294
															l1295:
																position, tokenIndex = position1294, tokenIndex1294
																if buffer[position] != rune('E') {
																	goto l894
																}
																position++
															}
														l1294:
															{
																position1296, tokenIndex1296 := position, tokenIndex
																if buffer[position] != rune('c') {
																	goto l1297
																}
																position++
																goto l1296
															l1297:
																position, tokenIndex = position1296, tokenIndex1296
																if buffer[position] != rune('C') {
																	goto l894
																}
																position++
															}
														l1296:
															{
																position1298, tokenIndex1298 := position, tokenIndex
																if buffer[position] != rune('n') {
																	goto l1299
																}
																position++
																goto l1298
															l1299:
																position, tokenIndex = position1298, tokenIndex1298
																if buffer[position] != rune('N') {
																	goto l894
																}
																position++
															}
														l1298:
															{
																position1300, tokenIndex1300 := position, tokenIndex
																if buffer[position] != rune('e') {
																	goto l1301
																}
																position++
																goto l1300
															l1301:
																position, tokenIndex = position1300, tokenIndex1300
																if buffer[position] != rune('E') {
																	goto l894
																}
																position++
															}
														l1300:
															{
																position1302, tokenIndex1302 := position, tokenIndex
																if buffer[position] != rune('w') {
																	goto l1303
																}
																position++
																goto l1302
															l1303:
																position, tokenIndex = position1302, tokenIndex1302
																if buffer[position] != rune('W') {
																	goto l894
																}
																position++
															}
														l1302:
															break
														case 'S', 's':
															{
																position1304, tokenIndex1304 := position, tokenIndex
																if buffer[position] != rune('s') {
																	goto l1305
																}
																position++
																goto l1304
															l1305:
																position, tokenIndex = position1304, tokenIndex1304
																if buffer[position] != rune('S') {
																	goto l894
																}
																position++
															}
														l1304:
															{
																position1306, tokenIndex1306 := position, tokenIndex
																if buffer[position] != rune('u') {
																	goto l1307
																}
																position++
																goto l1306
															l1307:
																position, tokenIndex = position1306, tokenIndex1306
																if buffer[position] != rune('U') {
																	goto l894
																}
																position++
															}
														l1306:
															{
																position1308, tokenIndex1308 := position, tokenIndex
																if buffer[position] != rune('b') {
																	goto l1309
																}
																position++
																goto l1308
															l1309:
																position, tokenIndex = position1308, tokenIndex1308
																if buffer[position] != rune('B') {
																	goto l894
																}
																position++
															}
														l1308:
															{
																position1310, tokenIndex1310 := position, tokenIndex
																if buffer[position] != rune('s') {
																	goto l1311
																}
																position++
																goto l1310
															l1311:
																position, tokenIndex = position1310, tokenIndex1310
																if buffer[position] != rune('S') {
																	goto l894
																}
																position++
															}
														l1310:
															break
														case 'R', 'r':
															{
																position1312, tokenIndex1312 := position, tokenIndex
																if buffer[position] != rune('r') {
																	goto l1313
																}
																position++
																goto l1312
															l1313:
																position, tokenIndex = position1312, tokenIndex1312
																if buffer[position] != rune('R') {
																	goto l894
																}
																position++
															}
														l1312:
															{
																position1314, tokenIndex1314 := position, tokenIndex
																if buffer[position] != rune('s') {
																	goto l1315
																}
																position++
																goto l1314
															l1315:
																position, tokenIndex = position1314, tokenIndex1314
																if buffer[position] != rune('S') {
																	goto l894
																}
																position++
															}
														l1314:
															{
																position1316, tokenIndex1316 := position, tokenIndex
																if buffer[position] != rune('h') {
																	goto l1317
																}
																position++
																goto l1316
															l1317:
																position, tokenIndex = position1316, tokenIndex1316
																if buffer[position] != rune('H') {
																	goto l894
																}
																position++
															}
														l1316:
															{
																position1318, tokenIndex1318 := position, tokenIndex
																if buffer[position] != rune('i') {
																	goto l1319
																}
																position++
																goto l1318
															l1319:
																position, tokenIndex = position1318, tokenIndex1318
																if buffer[position] != rune('I') {
																	goto l894
																}
																position++
															}
														l1318:
															{
																position1320, tokenIndex1320 := position, tokenIndex
																if buffer[position] != rune('f') {
																	goto l1321
																}
																position++
																goto l1320
															l1321:
																position, tokenIndex = position1320, tokenIndex1320
																if buffer[position] != rune('F') {
																	goto l894
																}
																position++
															}
														l1320:
															{
																position1322, tokenIndex1322 := position, tokenIndex
																if buffer[position] != rune('t') {
																	goto l1323
																}
																position++
																goto l1322
															l1323:
																position, tokenIndex = position1322, tokenIndex1322
																if buffer[position] != rune('T') {
																	goto l894
																}
																position++
															}
														l1322:
															break
														case 'P', 'p':
															{
																position1324, tokenIndex1324 := position, tokenIndex
																if buffer[position] != rune('p') {
																	goto l1325
																}
																position++
																goto l1324
															l1325:
																position, tokenIndex = position1324, tokenIndex1324
																if buffer[position] != rune('P') {
																	goto l894
																}
																position++
															}
														l1324:
															{
																position1326, tokenIndex1326 := position, tokenIndex
																if buffer[position] != rune('r') {
																	goto l1327
																}
																position++
																goto l1326
															l1327:
																position, tokenIndex = position1326, tokenIndex1326
																if buffer[position] != rune('R') {
																	goto l894
																}
																position++
															}
														l1326:
															{
																position1328, tokenIndex1328 := position, tokenIndex
																if buffer[position] != rune('e') {
																	goto l1329
																}
																position++
																goto l1328
															l1329:
																position, tokenIndex = position1328, tokenIndex1328
																if buffer[position] != rune('E') {
																	goto l894
																}
																position++
															}
														l1328:
															{
																position1330, tokenIndex1330 := position, tokenIndex
																if buffer[position] != rune('i') {
																	goto l1331
																}
																position++
																goto l1330
															l1331:
																position, tokenIndex = position1330, tokenIndex1330
																if buffer[position] != rune('I') {
																	goto l894
																}
																position++
															}
														l1330:
															{
																position1332, tokenIndex1332 := position, tokenIndex
																if buffer[position] != rune('n') {
																	goto l1333
																}
																position++
																goto l1332
															l1333:
																position, tokenIndex = position1332, tokenIndex1332
																if buffer[position] != rune('N') {
																	goto l894
																}
																position++
															}
														l1332:
															{
																position1334, tokenIndex1334 := position, tokenIndex
																if buffer[position] != rune('c') {
																	goto l1335
																}
																position++
																goto l1334
															l1335:
																position, tokenIndex = position1334, tokenIndex1334
																if buffer[position] != rune('C') {
																	goto l894
																}
																position++
															}
														l1334:
															break
														case 'O', 'o':
															{
																position1336, tokenIndex1336 := position, tokenIndex
																if buffer[position] != rune('o') {
																	goto l1337
																}
																position++
																goto l1336
															l1337:
																position, tokenIndex = position1336, tokenIndex1336
																if buffer[position] != rune('O') {
																	goto l894
																}
																position++
															}
														l1336:
															{
																position1338, tokenIndex1338 := position, tokenIndex
																if buffer[position] != rune('r') {
																	goto l1339
																}
																position++
																goto l1338
															l1339:
																position, tokenIndex = position1338, tokenIndex1338
																if buffer[position] != rune('R') {
																	goto l894
																}
																position++
															}
														l1338:
															break
														case 'N', 'n':
															{
																position1340, tokenIndex1340 := position, tokenIndex
																if buffer[position] != rune('n') {
																	goto l1341
																}
																position++
																goto l1340
															l1341:
																position, tokenIndex = position1340, tokenIndex1340
																if buffer[position] != rune('N') {
																	goto l894
																}
																position++
															}
														l1340:
															{
																position1342, tokenIndex1342 := position, tokenIndex
																if buffer[position] != rune('o') {
																	goto l1343
																}
																position++
																goto l1342
															l1343:
																position, tokenIndex = position1342, tokenIndex1342
																if buffer[position] != rune('O') {
																	goto l894
																}
																position++
															}
														l1342:
															{
																position1344, tokenIndex1344 := position, tokenIndex
																if buffer[position] != rune('t') {
																	goto l1345
																}
																position++
																goto l1344
															l1345:
																position, tokenIndex = position1344, tokenIndex1344
																if buffer[position] != rune('T') {
																	goto l894
																}
																position++
															}
														l1344:
															break
														case 'M', 'm':
															{
																position1346, tokenIndex1346 := position, tokenIndex
																if buffer[position] != rune('m') {
																	goto l1347
																}
																position++
																goto l1346
															l1347:
																position, tokenIndex = position1346, tokenIndex1346
																if buffer[position] != rune('M') {
																	goto l894
																}
																position++
															}
														l1346:
															{
																position1348, tokenIndex1348 := position, tokenIndex
																if buffer[position] != rune('u') {
																	goto l1349
																}
																position++
																goto l1348
															l1349:
																position, tokenIndex = position1348, tokenIndex1348
																if buffer[position] != rune('U') {
																	goto l894
																}
																position++
															}
														l1348:
															{
																position1350, tokenIndex1350 := position, tokenIndex
																if buffer[position] != rune('l') {
																	goto l1351
																}
																position++
																goto l1350
															l1351:
																position, tokenIndex = position1350, tokenIndex1350
																if buffer[position] != rune('L') {
																	goto l894
																}
																position++
															}
														l1350:
															{
																position1352, tokenIndex1352 := position, tokenIndex
																if buffer[position] != rune('t') {
																	goto l1353
																}
																position++
																goto l1352
															l1353:
																position, tokenIndex = position1352, tokenIndex1352
																if buffer[position] != rune('T') {
																	goto l894
																}
																position++
															}
														l1352:
															break
														case 'L', 'l':
															{
																position1354, tokenIndex1354 := position, tokenIndex
																if buffer[position] != rune('l') {
																	goto l1355
																}
																position++
																goto l1354
															l1355:
																position, tokenIndex = position1354, tokenIndex1354
																if buffer[position] != rune('L') {
																	goto l894
																}
																position++
															}
														l1354:
															{
																position1356, tokenIndex1356 := position, tokenIndex
																if buffer[position] != rune('t') {
																	goto l1357
																}
																position++
																goto l1356
															l1357:
																position, tokenIndex = position1356, tokenIndex1356
																if buffer[position] != rune('T') {
																	goto l894
																}
																position++
															}
														l1356:
															break
														case 'G', 'g':
															{
																position1358, tokenIndex1358 := position, tokenIndex
																if buffer[position] != rune('g') {
																	goto l1359
																}
																position++
																goto l1358
															l1359:
																position, tokenIndex = position1358, tokenIndex1358
																if buffer[position] != rune('G') {
																	goto l894
																}
																position++
															}
														l1358:
															{
																position1360, tokenIndex1360 := position, tokenIndex
																if buffer[position] != rune('t') {
																	goto l1361
																}
																position++
																goto l1360
															l1361:
																position, tokenIndex = position1360, tokenIndex1360
																if buffer[position] != rune('T') {
																	goto l894
																}
																position++
															}
														l1360:
															break
														case 'E', 'e':
															{
																position1362, tokenIndex1362 := position, tokenIndex
																if buffer[position] != rune('e') {
																	goto l1363
																}
																position++
																goto l1362
															l1363:
																position, tokenIndex = position1362, tokenIndex1362
																if buffer[position] != rune('E') {
																	goto l894
																}
																position++
															}
														l1362:
															{
																position1364, tokenIndex1364 := position, tokenIndex
																if buffer[position] != rune('q') {
																	goto l1365
																}
																position++
																goto l1364
															l1365:
																position, tokenIndex = position1364, tokenIndex1364
																if buffer[position] != rune('Q') {
																	goto l894
																}
																position++
															}
														l1364:
															break
														case 'D', 'd':
															{
																position1366, tokenIndex1366 := position, tokenIndex
																if buffer[position] != rune('d') {
																	goto l1367
																}
																position++
																goto l1366
															l1367:
																position, tokenIndex = position1366, tokenIndex1366
																if buffer[position] != rune('D') {
																	goto l894
																}
																position++
															}
														l1366:
															{
																position1368, tokenIndex1368 := position, tokenIndex
																if buffer[position] != rune('i') {
																	goto l1369
																}
																position++
																goto l1368
															l1369:
																position, tokenIndex = position1368, tokenIndex1368
																if buffer[position] != rune('I') {
																	goto l894
																}
																position++
															}
														l1368:
															{
																position1370, tokenIndex1370 := position, tokenIndex
																if buffer[position] != rune('v') {
																	goto l1371
																}
																position++
																goto l1370
															l1371:
																position, tokenIndex = position1370, tokenIndex1370
																if buffer[position] != rune('V') {
																	goto l894
																}
																position++
															}
														l1370:
															break
														case 'C', 'c':
															{
																position1372, tokenIndex1372 := position, tokenIndex
																if buffer[position] != rune('c') {
																	goto l1373
																}
																position++
																goto l1372
															l1373:
																position, tokenIndex = position1372, tokenIndex1372
																if buffer[position] != rune('C') {
																	goto l894
																}
																position++
															}
														l1372:
															{
																position1374, tokenIndex1374 := position, tokenIndex
																if buffer[position] != rune('a') {
																	goto l1375
																}
																position++
																goto l1374
															l1375:
																position, tokenIndex = position1374, tokenIndex1374
																if buffer[position] != rune('A') {
																	goto l894
																}
																position++
															}
														l1374:
															{
																position1376, tokenIndex1376 := position, tokenIndex
																if buffer[position] != rune('l') {
																	goto l1377
																}
																position++
																goto l1376
															l1377:
																position, tokenIndex = position1376, tokenIndex1376
																if buffer[position] != rune('L') {
																	goto l894
																}
																position++
															}
														l1376:
															{
																position1378, tokenIndex1378 := position, tokenIndex
																if buffer[position] != rune('l') {
																	goto l1379
																}
																position++
																goto l1378
															l1379:
																position, tokenIndex = position1378, tokenIndex1378
																if buffer[position] != rune('L') {
																	goto l894
																}
																position++
															}
														l1378:
															break
														case 'A', 'a':
															{
																position1380, tokenIndex1380 := position, tokenIndex
																if buffer[position] != rune('a') {
																	goto l1381
																}
																position++
																goto l1380
															l1381:
																position, tokenIndex = position1380, tokenIndex1380
																if buffer[position] != rune('A') {
																	goto l894
																}
																position++
															}
														l1380:
															{
																position1382, tokenIndex1382 := position, tokenIndex
																if buffer[position] != rune('s') {
																	goto l1383
																}
																position++
																goto l1382
															l1383:
																position, tokenIndex = position1382, tokenIndex1382
																if buffer[position] != rune('S') {
																	goto l894
																}
																position++
															}
														l1382:
															{
																position1384, tokenIndex1384 := position, tokenIndex
																if buffer[position] != rune('s') {
																	goto l1385
																}
																position++
																goto l1384
															l1385:
																position, tokenIndex = position1384, tokenIndex1384
																if buffer[position] != rune('S') {
																	goto l894
																}
																position++
															}
														l1384:
															{
																position1386, tokenIndex1386 := position, tokenIndex
																if buffer[position] != rune('i') {
																	goto l1387
																}
																position++
																goto l1386
															l1387:
																position, tokenIndex = position1386, tokenIndex1386
																if buffer[position] != rune('I') {
																	goto l894
																}
																position++
															}
														l1386:
															{
																position1388, tokenIndex1388 := position, tokenIndex
																if buffer[position] != rune('g') {
																	goto l1389
																}
																position++
																goto l1388
															l1389:
																position, tokenIndex = position1388, tokenIndex1388
																if buffer[position] != rune('G') {
																	goto l894
																}
																position++
															}
														l1388:
															{
																position1390, tokenIndex1390 := position, tokenIndex
																if buffer[position] != rune('n') {
																	goto l1391
																}
																position++
																goto l1390
															l1391:
																position, tokenIndex = position1390, tokenIndex1390
																if buffer[position] != rune('N') {
																	goto l894
																}
																position++
															}
														l1390:
															break
														default:
															break
														}
													}

												}
											l896:
												goto l895
											l894:
												position, tokenIndex = position894, tokenIndex894
											}
										l895:
											break
										case 'M', 'm':
											{
												position1392, tokenIndex1392 := position, tokenIndex
												if buffer[position] != rune('m') {
													goto l1393
												}
												position++
												goto l1392
											l1393:
												position, tokenIndex = position1392, tokenIndex1392
												if buffer[position] != rune('M') {
													goto l58
												}
												position++
											}
										l1392:
											{
												position1394, tokenIndex1394 := position, tokenIndex
												if buffer[position] != rune('e') {
													goto l1395
												}
												position++
												goto l1394
											l1395:
												position, tokenIndex = position1394, tokenIndex1394
												if buffer[position] != rune('E') {
													goto l58
												}
												position++
											}
										l1394:
											{
												position1396, tokenIndex1396 := position, tokenIndex
												if buffer[position] != rune('m') {
													goto l1397
												}
												position++
												goto l1396
											l1397:
												position, tokenIndex = position1396, tokenIndex1396
												if buffer[position] != rune('M') {
													goto l58
												}
												position++
											}
										l1396:
											{
												position1398, tokenIndex1398 := position, tokenIndex
												if buffer[position] != rune('b') {
													goto l1399
												}
												position++
												goto l1398
											l1399:
												position, tokenIndex = position1398, tokenIndex1398
												if buffer[position] != rune('B') {
													goto l58
												}
												position++
											}
										l1398:
											{
												position1400, tokenIndex1400 := position, tokenIndex
												if buffer[position] != rune('e') {
													goto l1401
												}
												position++
												goto l1400
											l1401:
												position, tokenIndex = position1400, tokenIndex1400
												if buffer[position] != rune('E') {
													goto l58
												}
												position++
											}
										l1400:
											{
												position1402, tokenIndex1402 := position, tokenIndex
												if buffer[position] != rune('r') {
													goto l1403
												}
												position++
												goto l1402
											l1403:
												position, tokenIndex = position1402, tokenIndex1402
												if buffer[position] != rune('R') {
													goto l58
												}
												position++
											}
										l1402:
											break
										default:
											{
												position1404, tokenIndex1404 := position, tokenIndex
												if buffer[position] != rune('a') {
													goto l1405
												}
												position++
												goto l1404
											l1405:
												position, tokenIndex = position1404, tokenIndex1404
												if buffer[position] != rune('A') {
													goto l58
												}
												position++
											}
										l1404:
											{
												position1406, tokenIndex1406 := position, tokenIndex
												if buffer[position] != rune('r') {
													goto l1407
												}
												position++
												goto l1406
											l1407:
												position, tokenIndex = position1406, tokenIndex1406
												if buffer[position] != rune('R') {
													goto l58
												}
												position++
											}
										l1406:
											{
												position1408, tokenIndex1408 := position, tokenIndex
												if buffer[position] != rune('t') {
													goto l1409
												}
												position++
												goto l1408
											l1409:
												position, tokenIndex = position1408, tokenIndex1408
												if buffer[position] != rune('T') {
													goto l58
												}
												position++
											}
										l1408:
											{
												position1410, tokenIndex1410 := position, tokenIndex
												if buffer[position] != rune('i') {
													goto l1411
												}
												position++
												goto l1410
											l1411:
												position, tokenIndex = position1410, tokenIndex1410
												if buffer[position] != rune('I') {
													goto l58
												}
												position++
											}
										l1410:
											{
												position1412, tokenIndex1412 := position, tokenIndex
												if buffer[position] != rune('f') {
													goto l1413
												}
												position++
												goto l1412
											l1413:
												position, tokenIndex = position1412, tokenIndex1412
												if buffer[position] != rune('F') {
													goto l58
												}
												position++
											}
										l1412:
											{
												position1414, tokenIndex1414 := position, tokenIndex
												if buffer[position] != rune('i') {
													goto l1415
												}
												position++
												goto l1414
											l1415:
												position, tokenIndex = position1414, tokenIndex1414
												if buffer[position] != rune('I') {
													goto l58
												}
												position++
											}
										l1414:
											{
												position1416, tokenIndex1416 := position, tokenIndex
												if buffer[position] != rune('c') {
													goto l1417
												}
												position++
												goto l1416
											l1417:
												position, tokenIndex = position1416, tokenIndex1416
												if buffer[position] != rune('C') {
													goto l58
												}
												position++
											}
										l1416:
											{
												position1418, tokenIndex1418 := position, tokenIndex
												if buffer[position] != rune('i') {
													goto l1419
												}
												position++
												goto l1418
											l1419:
												position, tokenIndex = position1418, tokenIndex1418
												if buffer[position] != rune('I') {
													goto l58
												}
												position++
											}
										l1418:
											{
												position1420, tokenIndex1420 := position, tokenIndex
												if buffer[position] != rune('a') {
													goto l1421
												}
												position++
												goto l1420
											l1421:
												position, tokenIndex = position1420, tokenIndex1420
												if buffer[position] != rune('A') {
													goto l58
												}
												position++
											}
										l1420:
											{
												position1422, tokenIndex1422 := position, tokenIndex
												if buffer[position] != rune('l') {
													goto l1423
												}
												position++
												goto l1422
											l1423:
												position, tokenIndex = position1422, tokenIndex1422
												if buffer[position] != rune('L') {
													goto l58
												}
												position++
											}
										l1422:
											break
										}
									}

								}
							l823:
								add(ruleNoteAttr, position814)
							}
							break
						case 'B', 'b':
							{
								position1424 := position
								{
									position1425, tokenIndex1425 := position, tokenIndex
									if buffer[position] != rune('b') {
										goto l1426
									}
									position++
									goto l1425
								l1426:
									position, tokenIndex = position1425, tokenIndex1425
									if buffer[position] != rune('B') {
										goto l58
									}
									position++
								}
							l1425:
								{
									position1427, tokenIndex1427 := position, tokenIndex
									if buffer[position] != rune('o') {
										goto l1428
									}
									position++
									goto l1427
								l1428:
									position, tokenIndex = position1427, tokenIndex1427
									if buffer[position] != rune('O') {
										goto l58
									}
									position++
								}
							l1427:
								{
									position1429, tokenIndex1429 := position, tokenIndex
									if buffer[position] != rune('d') {
										goto l1430
									}
									position++
									goto l1429
								l1430:
									position, tokenIndex = position1429, tokenIndex1429
									if buffer[position] != rune('D') {
										goto l58
									}
									position++
								}
							l1429:
								{
									position1431, tokenIndex1431 := position, tokenIndex
									if buffer[position] != rune('y') {
										goto l1432
									}
									position++
									goto l1431
								l1432:
									position, tokenIndex = position1431, tokenIndex1431
									if buffer[position] != rune('Y') {
										goto l58
									}
									position++
								}
							l1431:
								if !_rules[rulews]() {
									goto l58
								}
								if buffer[position] != rune(':') {
									goto l58
								}
								position++
								if !_rules[rulews]() {
									goto l58
								}
								{
									position1433, tokenIndex1433 := position, tokenIndex
									{
										position1435, tokenIndex1435 := position, tokenIndex
										if buffer[position] != rune('u') {
											goto l1436
										}
										position++
										goto l1435
									l1436:
										position, tokenIndex = position1435, tokenIndex1435
										if buffer[position] != rune('U') {
											goto l1434
										}
										position++
									}
								l1435:
									{
										position1437, tokenIndex1437 := position, tokenIndex
										if buffer[position] != rune('n') {
											goto l1438
										}
										position++
										goto l1437
									l1438:
										position, tokenIndex = position1437, tokenIndex1437
										if buffer[position] != rune('N') {
											goto l1434
										}
										position++
									}
								l1437:
									{
										position1439, tokenIndex1439 := position, tokenIndex
										if buffer[position] != rune('d') {
											goto l1440
										}
										position++
										goto l1439
									l1440:
										position, tokenIndex = position1439, tokenIndex1439
										if buffer[position] != rune('D') {
											goto l1434
										}
										position++
									}
								l1439:
									{
										position1441, tokenIndex1441 := position, tokenIndex
										if buffer[position] != rune('e') {
											goto l1442
										}
										position++
										goto l1441
									l1442:
										position, tokenIndex = position1441, tokenIndex1441
										if buffer[position] != rune('E') {
											goto l1434
										}
										position++
									}
								l1441:
									{
										position1443, tokenIndex1443 := position, tokenIndex
										if buffer[position] != rune('f') {
											goto l1444
										}
										position++
										goto l1443
									l1444:
										position, tokenIndex = position1443, tokenIndex1443
										if buffer[position] != rune('F') {
											goto l1434
										}
										position++
									}
								l1443:
									{
										position1445, tokenIndex1445 := position, tokenIndex
										if buffer[position] != rune('i') {
											goto l1446
										}
										position++
										goto l1445
									l1446:
										position, tokenIndex = position1445, tokenIndex1445
										if buffer[position] != rune('I') {
											goto l1434
										}
										position++
									}
								l1445:
									{
										position1447, tokenIndex1447 := position, tokenIndex
										if buffer[position] != rune('n') {
											goto l1448
										}
										position++
										goto l1447
									l1448:
										position, tokenIndex = position1447, tokenIndex1447
										if buffer[position] != rune('N') {
											goto l1434
										}
										position++
									}
								l1447:
									{
										position1449, tokenIndex1449 := position, tokenIndex
										if buffer[position] != rune('e') {
											goto l1450
										}
										position++
										goto l1449
									l1450:
										position, tokenIndex = position1449, tokenIndex1449
										if buffer[position] != rune('E') {
											goto l1434
										}
										position++
									}
								l1449:
									{
										position1451, tokenIndex1451 := position, tokenIndex
										if buffer[position] != rune('d') {
											goto l1452
										}
										position++
										goto l1451
									l1452:
										position, tokenIndex = position1451, tokenIndex1451
										if buffer[position] != rune('D') {
											goto l1434
										}
										position++
									}
								l1451:
									goto l1433
								l1434:
									position, tokenIndex = position1433, tokenIndex1433
									if !_rules[ruleNode]() {
										goto l58
									}
								}
							l1433:
								add(ruleBodyAttr, position1424)
							}
							break
						case 'T', 't':
							{
								position1453 := position
								{
									position1454, tokenIndex1454 := position, tokenIndex
									if buffer[position] != rune('t') {
										goto l1455
									}
									position++
									goto l1454
								l1455:
									position, tokenIndex = position1454, tokenIndex1454
									if buffer[position] != rune('T') {
										goto l58
									}
									position++
								}
							l1454:
								{
									position1456, tokenIndex1456 := position, tokenIndex
									if buffer[position] != rune('a') {
										goto l1457
									}
									position++
									goto l1456
								l1457:
									position, tokenIndex = position1456, tokenIndex1456
									if buffer[position] != rune('A') {
										goto l58
									}
									position++
								}
							l1456:
								{
									position1458, tokenIndex1458 := position, tokenIndex
									if buffer[position] != rune('g') {
										goto l1459
									}
									position++
									goto l1458
								l1459:
									position, tokenIndex = position1458, tokenIndex1458
									if buffer[position] != rune('G') {
										goto l58
									}
									position++
								}
							l1458:
								if !_rules[rulews]() {
									goto l58
								}
								if buffer[position] != rune(':') {
									goto l58
								}
								position++
								if !_rules[rulews]() {
									goto l58
								}
								{
									position1460, tokenIndex1460 := position, tokenIndex
									{
										position1462, tokenIndex1462 := position, tokenIndex
										if buffer[position] != rune('s') {
											goto l1463
										}
										position++
										goto l1462
									l1463:
										position, tokenIndex = position1462, tokenIndex1462
										if buffer[position] != rune('S') {
											goto l1461
										}
										position++
									}
								l1462:
									{
										position1464, tokenIndex1464 := position, tokenIndex
										if buffer[position] != rune('t') {
											goto l1465
										}
										position++
										goto l1464
									l1465:
										position, tokenIndex = position1464, tokenIndex1464
										if buffer[position] != rune('T') {
											goto l1461
										}
										position++
									}
								l1464:
									{
										position1466, tokenIndex1466 := position, tokenIndex
										if buffer[position] != rune('r') {
											goto l1467
										}
										position++
										goto l1466
									l1467:
										position, tokenIndex = position1466, tokenIndex1466
										if buffer[position] != rune('R') {
											goto l1461
										}
										position++
									}
								l1466:
									{
										position1468, tokenIndex1468 := position, tokenIndex
										if buffer[position] != rune('u') {
											goto l1469
										}
										position++
										goto l1468
									l1469:
										position, tokenIndex = position1468, tokenIndex1468
										if buffer[position] != rune('U') {
											goto l1461
										}
										position++
									}
								l1468:
									{
										position1470, tokenIndex1470 := position, tokenIndex
										if buffer[position] != rune('c') {
											goto l1471
										}
										position++
										goto l1470
									l1471:
										position, tokenIndex = position1470, tokenIndex1470
										if buffer[position] != rune('C') {
											goto l1461
										}
										position++
									}
								l1470:
									{
										position1472, tokenIndex1472 := position, tokenIndex
										if buffer[position] != rune('t') {
											goto l1473
										}
										position++
										goto l1472
									l1473:
										position, tokenIndex = position1472, tokenIndex1472
										if buffer[position] != rune('T') {
											goto l1461
										}
										position++
									}
								l1472:
									goto l1460
								l1461:
									position, tokenIndex = position1460, tokenIndex1460
									{
										position1474, tokenIndex1474 := position, tokenIndex
										if buffer[position] != rune('u') {
											goto l1475
										}
										position++
										goto l1474
									l1475:
										position, tokenIndex = position1474, tokenIndex1474
										if buffer[position] != rune('U') {
											goto l58
										}
										position++
									}
								l1474:
									{
										position1476, tokenIndex1476 := position, tokenIndex
										if buffer[position] != rune('n') {
											goto l1477
										}
										position++
										goto l1476
									l1477:
										position, tokenIndex = position1476, tokenIndex1476
										if buffer[position] != rune('N') {
											goto l58
										}
										position++
									}
								l1476:
									{
										position1478, tokenIndex1478 := position, tokenIndex
										if buffer[position] != rune('i') {
											goto l1479
										}
										position++
										goto l1478
									l1479:
										position, tokenIndex = position1478, tokenIndex1478
										if buffer[position] != rune('I') {
											goto l58
										}
										position++
									}
								l1478:
									{
										position1480, tokenIndex1480 := position, tokenIndex
										if buffer[position] != rune('o') {
											goto l1481
										}
										position++
										goto l1480
									l1481:
										position, tokenIndex = position1480, tokenIndex1480
										if buffer[position] != rune('O') {
											goto l58
										}
										position++
									}
								l1480:
									{
										position1482, tokenIndex1482 := position, tokenIndex
										if buffer[position] != rune('n') {
											goto l1483
										}
										position++
										goto l1482
									l1483:
										position, tokenIndex = position1482, tokenIndex1482
										if buffer[position] != rune('N') {
											goto l58
										}
										position++
									}
								l1482:
								}
							l1460:
								add(ruleTagAttr, position1453)
							}
							break
						case 'I', 'i':
							{
								position1484 := position
								{
									position1485, tokenIndex1485 := position, tokenIndex
									if buffer[position] != rune('i') {
										goto l1486
									}
									position++
									goto l1485
								l1486:
									position, tokenIndex = position1485, tokenIndex1485
									if buffer[position] != rune('I') {
										goto l58
									}
									position++
								}
							l1485:
								{
									position1487, tokenIndex1487 := position, tokenIndex
									if buffer[position] != rune('n') {
										goto l1488
									}
									position++
									goto l1487
								l1488:
									position, tokenIndex = position1487, tokenIndex1487
									if buffer[position] != rune('N') {
										goto l58
									}
									position++
								}
							l1487:
								{
									position1489, tokenIndex1489 := position, tokenIndex
									if buffer[position] != rune('t') {
										goto l1490
									}
									position++
									goto l1489
								l1490:
									position, tokenIndex = position1489, tokenIndex1489
									if buffer[position] != rune('T') {
										goto l58
									}
									position++
								}
							l1489:
								if !_rules[rulews]() {
									goto l58
								}
								if buffer[position] != rune(':') {
									goto l58
								}
								position++
								if !_rules[rulews]() {
									goto l58
								}
								{
									position1491, tokenIndex1491 := position, tokenIndex
									{
										position1493, tokenIndex1493 := position, tokenIndex
										if buffer[position] != rune('-') {
											goto l1494
										}
										position++
										goto l1493
									l1494:
										position, tokenIndex = position1493, tokenIndex1493
										if buffer[position] != rune('0') {
											goto l1491
										}
										position++
										{
											position1495, tokenIndex1495 := position, tokenIndex
											if buffer[position] != rune('x') {
												goto l1496
											}
											position++
											goto l1495
										l1496:
											position, tokenIndex = position1495, tokenIndex1495
											if buffer[position] != rune('X') {
												goto l1491
											}
											position++
										}
									l1495:
									}
								l1493:
									goto l1492
								l1491:
									position, tokenIndex = position1491, tokenIndex1491
								}
							l1492:
								{
									position1499, tokenIndex1499 := position, tokenIndex
									if !_rules[ruleInteger]() {
										goto l1500
									}
									goto l1499
								l1500:
									position, tokenIndex = position1499, tokenIndex1499
									if !_rules[ruleHex]() {
										goto l58
									}
								}
							l1499:
							l1497:
								{
									position1498, tokenIndex1498 := position, tokenIndex
									{
										position1501, tokenIndex1501 := position, tokenIndex
										if !_rules[ruleInteger]() {
											goto l1502
										}
										goto l1501
									l1502:
										position, tokenIndex = position1501, tokenIndex1501
										if !_rules[ruleHex]() {
											goto l1498
										}
									}
								l1501:
									goto l1497
								l1498:
									position, tokenIndex = position1498, tokenIndex1498
								}
								if !_rules[rulews]() {
									goto l58
								}
								add(ruleSignedIntAttr, position1484)
							}
							break
						default:
							{
								position1503 := position
								if !_rules[rulews]() {
									goto l58
								}
								{
									position1504, tokenIndex1504 := position, tokenIndex
									if buffer[position] != rune('s') {
										goto l1505
									}
									position++
									goto l1504
								l1505:
									position, tokenIndex = position1504, tokenIndex1504
									if buffer[position] != rune('S') {
										goto l58
									}
									position++
								}
							l1504:
								{
									position1506, tokenIndex1506 := position, tokenIndex
									if buffer[position] != rune('p') {
										goto l1507
									}
									position++
									goto l1506
								l1507:
									position, tokenIndex = position1506, tokenIndex1506
									if buffer[position] != rune('P') {
										goto l58
									}
									position++
								}
							l1506:
								{
									position1508, tokenIndex1508 := position, tokenIndex
									if buffer[position] != rune('e') {
										goto l1509
									}
									position++
									goto l1508
								l1509:
									position, tokenIndex = position1508, tokenIndex1508
									if buffer[position] != rune('E') {
										goto l58
									}
									position++
								}
							l1508:
								{
									position1510, tokenIndex1510 := position, tokenIndex
									if buffer[position] != rune('c') {
										goto l1511
									}
									position++
									goto l1510
								l1511:
									position, tokenIndex = position1510, tokenIndex1510
									if buffer[position] != rune('C') {
										goto l58
									}
									position++
								}
							l1510:
								if !_rules[rulews]() {
									goto l58
								}
								add(ruleRandomSpec, position1503)
							}
							break
						}
					}

				}
			l60:
				add(ruleOneAttr, position59)
			}
			return true
		l58:
			position, tokenIndex = position58, tokenIndex58
			return false
		},
		/* 25 Attrs <- <(ws OneAttr)*> */
		nil,
		/* 26 Attr <- <(OneAttr ws Attrs ws)?> */
		nil,
		/* 27 Statement <- <(Node ws NodeType ws Attr ws)> */
		nil,
		/* 28 Node <- <('@' NonZeroDecimalDigit DecimalDigit*)> */
		func() bool {
			position1515, tokenIndex1515 := position, tokenIndex
			{
				position1516 := position
				if buffer[position] != rune('@') {
					goto l1515
				}
				position++
				if !_rules[ruleNonZeroDecimalDigit]() {
					goto l1515
				}
			l1517:
				{
					position1518, tokenIndex1518 := position, tokenIndex
					if !_rules[ruleDecimalDigit]() {
						goto l1518
					}
					goto l1517
				l1518:
					position, tokenIndex = position1518, tokenIndex1518
				}
				add(ruleNode, position1516)
			}
			return true
		l1515:
			position, tokenIndex = position1515, tokenIndex1515
			return false
		},
		/* 29 Integer <- <('0' / (NonZeroDecimalDigit DecimalDigit*))> */
		func() bool {
			position1519, tokenIndex1519 := position, tokenIndex
			{
				position1520 := position
				{
					position1521, tokenIndex1521 := position, tokenIndex
					if buffer[position] != rune('0') {
						goto l1522
					}
					position++
					goto l1521
				l1522:
					position, tokenIndex = position1521, tokenIndex1521
					if !_rules[ruleNonZeroDecimalDigit]() {
						goto l1519
					}
				l1523:
					{
						position1524, tokenIndex1524 := position, tokenIndex
						if !_rules[ruleDecimalDigit]() {
							goto l1524
						}
						goto l1523
					l1524:
						position, tokenIndex = position1524, tokenIndex1524
					}
				}
			l1521:
				add(ruleInteger, position1520)
			}
			return true
		l1519:
			position, tokenIndex = position1519, tokenIndex1519
			return false
		},
		/* 30 NodeType <- <([a-z] / '_')+> */
		nil,
		/* 31 DecimalDigit <- <[0-9]> */
		func() bool {
			position1526, tokenIndex1526 := position, tokenIndex
			{
				position1527 := position
				if c := buffer[position]; c < rune('0') || c > rune('9') {
					goto l1526
				}
				position++
				add(ruleDecimalDigit, position1527)
			}
			return true
		l1526:
			position, tokenIndex = position1526, tokenIndex1526
			return false
		},
		/* 32 NonZeroDecimalDigit <- <[1-9]> */
		func() bool {
			position1528, tokenIndex1528 := position, tokenIndex
			{
				position1529 := position
				if c := buffer[position]; c < rune('1') || c > rune('9') {
					goto l1528
				}
				position++
				add(ruleNonZeroDecimalDigit, position1529)
			}
			return true
		l1528:
			position, tokenIndex = position1528, tokenIndex1528
			return false
		},
		/* 33 ws <- <((&('\n') '\n') | (&('\r') '\r') | (&('\t') '\t') | (&(' ') ' '))*> */
		func() bool {
			{
				position1531 := position
			l1532:
				{
					position1533, tokenIndex1533 := position, tokenIndex
					{
						switch buffer[position] {
						case '\n':
							if buffer[position] != rune('\n') {
								goto l1533
							}
							position++
							break
						case '\r':
							if buffer[position] != rune('\r') {
								goto l1533
							}
							position++
							break
						case '\t':
							if buffer[position] != rune('\t') {
								goto l1533
							}
							position++
							break
						default:
							if buffer[position] != rune(' ') {
								goto l1533
							}
							position++
							break
						}
					}

					goto l1532
				l1533:
					position, tokenIndex = position1533, tokenIndex1533
				}
				add(rulews, position1531)
			}
			return true
		},
		/* 34 EOF <- <!.> */
		nil,
	}
	p.rules = _rules
}
