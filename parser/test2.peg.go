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
	rules  [35]func() bool
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
		/* 2 FileRef <- <((('_'? ((&('_') '_') | (&('0' | '1' | '2' | '3' | '4' | '5' | '6' | '7' | '8' | '9') [0-9]) | (&('a' | 'b' | 'c' | 'd' | 'e' | 'f' | 'g' | 'h' | 'i' | 'j' | 'k' | 'l' | 'm' | 'n' | 'o' | 'p' | 'q' | 'r' | 's' | 't' | 'u' | 'v' | 'w' | 'x' | 'y' | 'z') [a-z]) | (&('-') '-') | (&('A' | 'B' | 'C' | 'D' | 'E' | 'F' | 'G' | 'H' | 'I' | 'J' | 'K' | 'L' | 'M' | 'N' | 'O' | 'P' | 'Q' | 'R' | 'S' | 'T' | 'U' | 'V' | 'W' | 'X' | 'Y' | 'Z') [A-Z]))+ '.' ((&('_') '_') | (&('a' | 'b' | 'c' | 'd' | 'e' | 'f' | 'g' | 'h' | 'i' | 'j' | 'k' | 'l' | 'm' | 'n' | 'o' | 'p' | 'q' | 'r' | 's' | 't' | 'u' | 'v' | 'w' | 'x' | 'y' | 'z') [a-z]) | (&('A' | 'B' | 'C' | 'D' | 'E' | 'F' | 'G' | 'H' | 'I' | 'J' | 'K' | 'L' | 'M' | 'N' | 'O' | 'P' | 'Q' | 'R' | 'S' | 'T' | 'U' | 'V' | 'W' | 'X' | 'Y' | 'Z') [A-Z]))+) / ('<' ('b' / 'B') ('u' / 'U') ('i' / 'I') ('l' / 'L') ('t' / 'T') '-' ('i' / 'I') ('n' / 'N') '>')) ':' Integer ws)> */
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
		/* 14 NoteAttr <- <(('n' / 'N') ('o' / 'O') ('t' / 'T') ('e' / 'E') ws ':' ws ((('p' / 'P') ('t' / 'T') ('r' / 'R') ('m' / 'M') ('e' / 'E') ('m' / 'M')) / ((&('P' | 'p') (('p' / 'P') ('s' / 'S') ('e' / 'E') ('u' / 'U') ('d' / 'D') ('o' / 'O') ' ' ('t' / 'T') ('m' / 'M') ('p' / 'P') ('l' / 'L'))) | (&('O' | 'o') (('o' / 'O') ('p' / 'P') ('e' / 'E') ('r' / 'R') ('a' / 'A') ('t' / 'T') ('o' / 'O') ('r' / 'R') ' ' ((('g' / 'G') ('e' / 'E')) / (('l' / 'L') ('e' / 'E')) / (('g' / 'G') ('t' / 'T')) / (('l' / 'L') ('t' / 'T'))))) | (&('M' | 'm') (('m' / 'M') ('e' / 'E') ('m' / 'M') ('b' / 'B') ('e' / 'E') ('r' / 'R'))) | (&('A' | 'a') (('a' / 'A') ('r' / 'R') ('t' / 'T') ('i' / 'I') ('f' / 'F') ('i' / 'I') ('c' / 'C') ('i' / 'I') ('a' / 'A') ('l' / 'L'))))))> */
		nil,
		/* 15 AccsAttr <- <(('a' / 'A') ('c' / 'C') ('c' / 'C') ('s' / 'S') ws ':' ws ((('p' / 'P') ('r' / 'R') ('i' / 'I') ('v' / 'V')) / (('p' / 'P') ('u' / 'U') ('b' / 'B'))))> */
		nil,
		/* 16 QualAttr <- <(('q' / 'Q') ('u' / 'U') ('a' / 'A') ('l' / 'L') ws ':' ws ((&('R' | 'r') ('r' / 'R')) | (&('C') 'C') | (&('c') 'c') | (&('V' | 'v') ('v' / 'V')))+ ws)> */
		nil,
		/* 17 SignAttr <- <(('s' / 'S') ('i' / 'I') ('g' / 'G') ('n' / 'N') ws ':' ws ((('s' / 'S') ('i' / 'I') ('g' / 'G') ('n' / 'N') ('e' / 'E') ('d' / 'D')) / (('u' / 'U') ('n' / 'N') ('s' / 'S') ('i' / 'I') ('g' / 'G') ('n' / 'N') ('e' / 'E') ('d' / 'D'))))> */
		nil,
		/* 18 NodeName <- <((('r' / 'R') ('e' / 'E') ('t' / 'T') ('n' / 'N')) / (('p' / 'P') ('r' / 'R') ('m' / 'M') ('s' / 'S')) / (('a' / 'A') ('r' / 'R') ('g' / 'G') ('s' / 'S')) / (('d' / 'D') ('e' / 'E') ('c' / 'C') ('l' / 'L')) / (('e' / 'E') ('x' / 'X') ('p' / 'P') ('r' / 'R')) / (('f' / 'F') ('n' / 'N')) / (('c' / 'C') ('n' / 'N') ('s' / 'S') ('t' / 'T')) / (('v' / 'V') ('a' / 'A') ('r' / 'R') ('s' / 'S')) / (('v' / 'V') ('a' / 'A') ('l' / 'L') ('u' / 'U')) / (('c' / 'C') ('h' / 'H') ('a' / 'A') ('n' / 'N')) / (('p' / 'P') ('u' / 'U') ('r' / 'R') ('p' / 'P')) / (('f' / 'F') ('l' / 'L') ('d' / 'D') ('s' / 'S')) / (('c' / 'C') ('l' / 'L') ('s' / 'S')) / (('c' / 'C') ('s' / 'S') ('t' / 'T') ('s' / 'S')) / (('t' / 'T') ('y' / 'Y') ('p' / 'P') ('e' / 'E')) / (('m' / 'M') ('i' / 'I') ('n' / 'N')) / (('m' / 'M') ('a' / 'A') ('x' / 'X')) / (('b' / 'B') ('i' / 'I') ('n' / 'N') ('f' / 'F')) / (('s' / 'S') ('c' / 'C') ('p' / 'P') ('e' / 'E')) / (('i' / 'I') ('n' / 'N') ('i' / 'I') ('t' / 'T')) / ((&('O' | 'o') OpAttr) | (&('V' | 'v') (('v' / 'V') ('a' / 'A') ('l' / 'L'))) | (&('I' | 'i') (('i' / 'I') ('d' / 'D') ('x' / 'X'))) | (&('S' | 's') (('s' / 'S') ('i' / 'I') ('z' / 'Z') ('e' / 'E'))) | (&('B' | 'b') (('b' / 'B') ('p' / 'P') ('o' / 'O') ('s' / 'S'))) | (&('C' | 'c') (('c' / 'C') ('h' / 'H') ('a' / 'A') ('i' / 'I') ('n' / 'N'))) | (&('E' | 'e') (('e' / 'E') ('l' / 'L') ('t' / 'T') ('s' / 'S'))) | (&('D' | 'd') (('d' / 'D') ('o' / 'O') ('m' / 'M') ('n' / 'N'))) | (&('T' | 't') (('t' / 'T') ('y' / 'Y') ('p' / 'P') ('e' / 'E'))) | (&('M' | 'm') (('m' / 'M') ('n' / 'N') ('g' / 'G') ('l' / 'L'))) | (&('N' | 'n') (('n' / 'N') ('a' / 'A') ('m' / 'M') ('e' / 'E'))) | (&('U' | 'u') (('u' / 'U') ('n' / 'N') ('q' / 'Q') ('l' / 'L'))) | (&('L' | 'l') (('l' / 'L') ('a' / 'A') ('b' / 'B') ('l' / 'L'))) | (&('P' | 'p') (('p' / 'P') ('t' / 'T') ('d' / 'D'))) | (&('F' | 'f') (('f' / 'F') ('n' / 'N') ('c' / 'C') ('s' / 'S'))) | (&('R' | 'r') (('r' / 'R') ('e' / 'E') ('f' / 'F') ('d' / 'D'))) | (&('A' | 'a') (('a' / 'A') ('r' / 'R') ('g' / 'G') ('t' / 'T'))) | (&('0' | '1' | '2' | '3' | '4' | '5' | '6' | '7' | '8' | '9') Integer)))> */
		nil,
		/* 19 NodeAttr <- <(ws NodeName ws ':' ws Node ws)> */
		nil,
		/* 20 SpecValue <- <(ws (('s' / 'S') ('p' / 'P') ('e' / 'E') ('c' / 'C') ':')? ws (((&('R' | 'r') (('r' / 'R') ('e' / 'E') ('g' / 'G') ('i' / 'I') ('s' / 'S') ('t' / 'T') ('e' / 'E') ('r' / 'R'))) | (&('V' | 'v') (('v' / 'V') ('i' / 'I') ('r' / 'R') ('t' / 'T'))) | (&('P' | 'p') (('p' / 'P') ('u' / 'U') ('r' / 'R') ('e' / 'E'))) | (&('B' | 'b') (('b' / 'B') ('i' / 'I') ('t' / 'T') ('f' / 'F') ('i' / 'I') ('e' / 'E') ('l' / 'L') ('d' / 'D'))) | (&('M' | 'm') (('m' / 'M') ('u' / 'U') ('t' / 'T') ('a' / 'A') ('b' / 'B') ('l' / 'L') ('e' / 'E')))) ws)*)> */
		nil,
		/* 21 LngtAttr <- <(('l' / 'L') ('n' / 'N') ('g' / 'G') ('t' / 'T') ws ':' ws Integer ws)> */
		nil,
		/* 22 StringAttr <- <(('s' / 'S') ('t' / 'T') ('r' / 'R') ('g' / 'G') ':' .+)> */
		nil,
		/* 23 OneAttr <- <(AddrAttr / SpecValue / NodeAttr / SourceAttr / IntAttr / StringAttr / IntAttr3 / LngtAttr / AccsAttr / ((&('A' | 'a') IntAttr2) | (&('Q' | 'q') QualAttr) | (&('L' | 'l') LinkAttr) | (&('N' | 'n') NoteAttr) | (&('B' | 'b') BodyAttr) | (&('T' | 't') TagAttr) | (&('I' | 'i') SignedIntAttr) | (&('S' | 's') SignAttr)))> */
		func() bool {
			position57, tokenIndex57 := position, tokenIndex
			{
				position58 := position
				{
					position59, tokenIndex59 := position, tokenIndex
					{
						position61 := position
						{
							position62, tokenIndex62 := position, tokenIndex
							if buffer[position] != rune('a') {
								goto l63
							}
							position++
							goto l62
						l63:
							position, tokenIndex = position62, tokenIndex62
							if buffer[position] != rune('A') {
								goto l60
							}
							position++
						}
					l62:
						{
							position64, tokenIndex64 := position, tokenIndex
							if buffer[position] != rune('d') {
								goto l65
							}
							position++
							goto l64
						l65:
							position, tokenIndex = position64, tokenIndex64
							if buffer[position] != rune('D') {
								goto l60
							}
							position++
						}
					l64:
						{
							position66, tokenIndex66 := position, tokenIndex
							if buffer[position] != rune('d') {
								goto l67
							}
							position++
							goto l66
						l67:
							position, tokenIndex = position66, tokenIndex66
							if buffer[position] != rune('D') {
								goto l60
							}
							position++
						}
					l66:
						{
							position68, tokenIndex68 := position, tokenIndex
							if buffer[position] != rune('r') {
								goto l69
							}
							position++
							goto l68
						l69:
							position, tokenIndex = position68, tokenIndex68
							if buffer[position] != rune('R') {
								goto l60
							}
							position++
						}
					l68:
						if !_rules[rulews]() {
							goto l60
						}
						if buffer[position] != rune(':') {
							goto l60
						}
						position++
						if !_rules[rulews]() {
							goto l60
						}
						{
							position70 := position
							{
								position73, tokenIndex73 := position, tokenIndex
								if !_rules[ruleDecimalDigit]() {
									goto l74
								}
								goto l73
							l74:
								position, tokenIndex = position73, tokenIndex73
								if !_rules[ruleHex]() {
									goto l60
								}
							}
						l73:
						l71:
							{
								position72, tokenIndex72 := position, tokenIndex
								{
									position75, tokenIndex75 := position, tokenIndex
									if !_rules[ruleDecimalDigit]() {
										goto l76
									}
									goto l75
								l76:
									position, tokenIndex = position75, tokenIndex75
									if !_rules[ruleHex]() {
										goto l72
									}
								}
							l75:
								goto l71
							l72:
								position, tokenIndex = position72, tokenIndex72
							}
							add(ruleAddr, position70)
						}
						add(ruleAddrAttr, position61)
					}
					goto l59
				l60:
					position, tokenIndex = position59, tokenIndex59
					{
						position78 := position
						if !_rules[rulews]() {
							goto l77
						}
						{
							position79, tokenIndex79 := position, tokenIndex
							{
								position81, tokenIndex81 := position, tokenIndex
								if buffer[position] != rune('s') {
									goto l82
								}
								position++
								goto l81
							l82:
								position, tokenIndex = position81, tokenIndex81
								if buffer[position] != rune('S') {
									goto l79
								}
								position++
							}
						l81:
							{
								position83, tokenIndex83 := position, tokenIndex
								if buffer[position] != rune('p') {
									goto l84
								}
								position++
								goto l83
							l84:
								position, tokenIndex = position83, tokenIndex83
								if buffer[position] != rune('P') {
									goto l79
								}
								position++
							}
						l83:
							{
								position85, tokenIndex85 := position, tokenIndex
								if buffer[position] != rune('e') {
									goto l86
								}
								position++
								goto l85
							l86:
								position, tokenIndex = position85, tokenIndex85
								if buffer[position] != rune('E') {
									goto l79
								}
								position++
							}
						l85:
							{
								position87, tokenIndex87 := position, tokenIndex
								if buffer[position] != rune('c') {
									goto l88
								}
								position++
								goto l87
							l88:
								position, tokenIndex = position87, tokenIndex87
								if buffer[position] != rune('C') {
									goto l79
								}
								position++
							}
						l87:
							if buffer[position] != rune(':') {
								goto l79
							}
							position++
							goto l80
						l79:
							position, tokenIndex = position79, tokenIndex79
						}
					l80:
						if !_rules[rulews]() {
							goto l77
						}
					l89:
						{
							position90, tokenIndex90 := position, tokenIndex
							{
								switch buffer[position] {
								case 'R', 'r':
									{
										position92, tokenIndex92 := position, tokenIndex
										if buffer[position] != rune('r') {
											goto l93
										}
										position++
										goto l92
									l93:
										position, tokenIndex = position92, tokenIndex92
										if buffer[position] != rune('R') {
											goto l90
										}
										position++
									}
								l92:
									{
										position94, tokenIndex94 := position, tokenIndex
										if buffer[position] != rune('e') {
											goto l95
										}
										position++
										goto l94
									l95:
										position, tokenIndex = position94, tokenIndex94
										if buffer[position] != rune('E') {
											goto l90
										}
										position++
									}
								l94:
									{
										position96, tokenIndex96 := position, tokenIndex
										if buffer[position] != rune('g') {
											goto l97
										}
										position++
										goto l96
									l97:
										position, tokenIndex = position96, tokenIndex96
										if buffer[position] != rune('G') {
											goto l90
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
											goto l90
										}
										position++
									}
								l98:
									{
										position100, tokenIndex100 := position, tokenIndex
										if buffer[position] != rune('s') {
											goto l101
										}
										position++
										goto l100
									l101:
										position, tokenIndex = position100, tokenIndex100
										if buffer[position] != rune('S') {
											goto l90
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
											goto l90
										}
										position++
									}
								l102:
									{
										position104, tokenIndex104 := position, tokenIndex
										if buffer[position] != rune('e') {
											goto l105
										}
										position++
										goto l104
									l105:
										position, tokenIndex = position104, tokenIndex104
										if buffer[position] != rune('E') {
											goto l90
										}
										position++
									}
								l104:
									{
										position106, tokenIndex106 := position, tokenIndex
										if buffer[position] != rune('r') {
											goto l107
										}
										position++
										goto l106
									l107:
										position, tokenIndex = position106, tokenIndex106
										if buffer[position] != rune('R') {
											goto l90
										}
										position++
									}
								l106:
									break
								case 'V', 'v':
									{
										position108, tokenIndex108 := position, tokenIndex
										if buffer[position] != rune('v') {
											goto l109
										}
										position++
										goto l108
									l109:
										position, tokenIndex = position108, tokenIndex108
										if buffer[position] != rune('V') {
											goto l90
										}
										position++
									}
								l108:
									{
										position110, tokenIndex110 := position, tokenIndex
										if buffer[position] != rune('i') {
											goto l111
										}
										position++
										goto l110
									l111:
										position, tokenIndex = position110, tokenIndex110
										if buffer[position] != rune('I') {
											goto l90
										}
										position++
									}
								l110:
									{
										position112, tokenIndex112 := position, tokenIndex
										if buffer[position] != rune('r') {
											goto l113
										}
										position++
										goto l112
									l113:
										position, tokenIndex = position112, tokenIndex112
										if buffer[position] != rune('R') {
											goto l90
										}
										position++
									}
								l112:
									{
										position114, tokenIndex114 := position, tokenIndex
										if buffer[position] != rune('t') {
											goto l115
										}
										position++
										goto l114
									l115:
										position, tokenIndex = position114, tokenIndex114
										if buffer[position] != rune('T') {
											goto l90
										}
										position++
									}
								l114:
									break
								case 'P', 'p':
									{
										position116, tokenIndex116 := position, tokenIndex
										if buffer[position] != rune('p') {
											goto l117
										}
										position++
										goto l116
									l117:
										position, tokenIndex = position116, tokenIndex116
										if buffer[position] != rune('P') {
											goto l90
										}
										position++
									}
								l116:
									{
										position118, tokenIndex118 := position, tokenIndex
										if buffer[position] != rune('u') {
											goto l119
										}
										position++
										goto l118
									l119:
										position, tokenIndex = position118, tokenIndex118
										if buffer[position] != rune('U') {
											goto l90
										}
										position++
									}
								l118:
									{
										position120, tokenIndex120 := position, tokenIndex
										if buffer[position] != rune('r') {
											goto l121
										}
										position++
										goto l120
									l121:
										position, tokenIndex = position120, tokenIndex120
										if buffer[position] != rune('R') {
											goto l90
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
											goto l90
										}
										position++
									}
								l122:
									break
								case 'B', 'b':
									{
										position124, tokenIndex124 := position, tokenIndex
										if buffer[position] != rune('b') {
											goto l125
										}
										position++
										goto l124
									l125:
										position, tokenIndex = position124, tokenIndex124
										if buffer[position] != rune('B') {
											goto l90
										}
										position++
									}
								l124:
									{
										position126, tokenIndex126 := position, tokenIndex
										if buffer[position] != rune('i') {
											goto l127
										}
										position++
										goto l126
									l127:
										position, tokenIndex = position126, tokenIndex126
										if buffer[position] != rune('I') {
											goto l90
										}
										position++
									}
								l126:
									{
										position128, tokenIndex128 := position, tokenIndex
										if buffer[position] != rune('t') {
											goto l129
										}
										position++
										goto l128
									l129:
										position, tokenIndex = position128, tokenIndex128
										if buffer[position] != rune('T') {
											goto l90
										}
										position++
									}
								l128:
									{
										position130, tokenIndex130 := position, tokenIndex
										if buffer[position] != rune('f') {
											goto l131
										}
										position++
										goto l130
									l131:
										position, tokenIndex = position130, tokenIndex130
										if buffer[position] != rune('F') {
											goto l90
										}
										position++
									}
								l130:
									{
										position132, tokenIndex132 := position, tokenIndex
										if buffer[position] != rune('i') {
											goto l133
										}
										position++
										goto l132
									l133:
										position, tokenIndex = position132, tokenIndex132
										if buffer[position] != rune('I') {
											goto l90
										}
										position++
									}
								l132:
									{
										position134, tokenIndex134 := position, tokenIndex
										if buffer[position] != rune('e') {
											goto l135
										}
										position++
										goto l134
									l135:
										position, tokenIndex = position134, tokenIndex134
										if buffer[position] != rune('E') {
											goto l90
										}
										position++
									}
								l134:
									{
										position136, tokenIndex136 := position, tokenIndex
										if buffer[position] != rune('l') {
											goto l137
										}
										position++
										goto l136
									l137:
										position, tokenIndex = position136, tokenIndex136
										if buffer[position] != rune('L') {
											goto l90
										}
										position++
									}
								l136:
									{
										position138, tokenIndex138 := position, tokenIndex
										if buffer[position] != rune('d') {
											goto l139
										}
										position++
										goto l138
									l139:
										position, tokenIndex = position138, tokenIndex138
										if buffer[position] != rune('D') {
											goto l90
										}
										position++
									}
								l138:
									break
								default:
									{
										position140, tokenIndex140 := position, tokenIndex
										if buffer[position] != rune('m') {
											goto l141
										}
										position++
										goto l140
									l141:
										position, tokenIndex = position140, tokenIndex140
										if buffer[position] != rune('M') {
											goto l90
										}
										position++
									}
								l140:
									{
										position142, tokenIndex142 := position, tokenIndex
										if buffer[position] != rune('u') {
											goto l143
										}
										position++
										goto l142
									l143:
										position, tokenIndex = position142, tokenIndex142
										if buffer[position] != rune('U') {
											goto l90
										}
										position++
									}
								l142:
									{
										position144, tokenIndex144 := position, tokenIndex
										if buffer[position] != rune('t') {
											goto l145
										}
										position++
										goto l144
									l145:
										position, tokenIndex = position144, tokenIndex144
										if buffer[position] != rune('T') {
											goto l90
										}
										position++
									}
								l144:
									{
										position146, tokenIndex146 := position, tokenIndex
										if buffer[position] != rune('a') {
											goto l147
										}
										position++
										goto l146
									l147:
										position, tokenIndex = position146, tokenIndex146
										if buffer[position] != rune('A') {
											goto l90
										}
										position++
									}
								l146:
									{
										position148, tokenIndex148 := position, tokenIndex
										if buffer[position] != rune('b') {
											goto l149
										}
										position++
										goto l148
									l149:
										position, tokenIndex = position148, tokenIndex148
										if buffer[position] != rune('B') {
											goto l90
										}
										position++
									}
								l148:
									{
										position150, tokenIndex150 := position, tokenIndex
										if buffer[position] != rune('l') {
											goto l151
										}
										position++
										goto l150
									l151:
										position, tokenIndex = position150, tokenIndex150
										if buffer[position] != rune('L') {
											goto l90
										}
										position++
									}
								l150:
									{
										position152, tokenIndex152 := position, tokenIndex
										if buffer[position] != rune('e') {
											goto l153
										}
										position++
										goto l152
									l153:
										position, tokenIndex = position152, tokenIndex152
										if buffer[position] != rune('E') {
											goto l90
										}
										position++
									}
								l152:
									break
								}
							}

							if !_rules[rulews]() {
								goto l90
							}
							goto l89
						l90:
							position, tokenIndex = position90, tokenIndex90
						}
						add(ruleSpecValue, position78)
					}
					goto l59
				l77:
					position, tokenIndex = position59, tokenIndex59
					{
						position155 := position
						if !_rules[rulews]() {
							goto l154
						}
						{
							position156 := position
							{
								position157, tokenIndex157 := position, tokenIndex
								{
									position159, tokenIndex159 := position, tokenIndex
									if buffer[position] != rune('r') {
										goto l160
									}
									position++
									goto l159
								l160:
									position, tokenIndex = position159, tokenIndex159
									if buffer[position] != rune('R') {
										goto l158
									}
									position++
								}
							l159:
								{
									position161, tokenIndex161 := position, tokenIndex
									if buffer[position] != rune('e') {
										goto l162
									}
									position++
									goto l161
								l162:
									position, tokenIndex = position161, tokenIndex161
									if buffer[position] != rune('E') {
										goto l158
									}
									position++
								}
							l161:
								{
									position163, tokenIndex163 := position, tokenIndex
									if buffer[position] != rune('t') {
										goto l164
									}
									position++
									goto l163
								l164:
									position, tokenIndex = position163, tokenIndex163
									if buffer[position] != rune('T') {
										goto l158
									}
									position++
								}
							l163:
								{
									position165, tokenIndex165 := position, tokenIndex
									if buffer[position] != rune('n') {
										goto l166
									}
									position++
									goto l165
								l166:
									position, tokenIndex = position165, tokenIndex165
									if buffer[position] != rune('N') {
										goto l158
									}
									position++
								}
							l165:
								goto l157
							l158:
								position, tokenIndex = position157, tokenIndex157
								{
									position168, tokenIndex168 := position, tokenIndex
									if buffer[position] != rune('p') {
										goto l169
									}
									position++
									goto l168
								l169:
									position, tokenIndex = position168, tokenIndex168
									if buffer[position] != rune('P') {
										goto l167
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
										goto l167
									}
									position++
								}
							l170:
								{
									position172, tokenIndex172 := position, tokenIndex
									if buffer[position] != rune('m') {
										goto l173
									}
									position++
									goto l172
								l173:
									position, tokenIndex = position172, tokenIndex172
									if buffer[position] != rune('M') {
										goto l167
									}
									position++
								}
							l172:
								{
									position174, tokenIndex174 := position, tokenIndex
									if buffer[position] != rune('s') {
										goto l175
									}
									position++
									goto l174
								l175:
									position, tokenIndex = position174, tokenIndex174
									if buffer[position] != rune('S') {
										goto l167
									}
									position++
								}
							l174:
								goto l157
							l167:
								position, tokenIndex = position157, tokenIndex157
								{
									position177, tokenIndex177 := position, tokenIndex
									if buffer[position] != rune('a') {
										goto l178
									}
									position++
									goto l177
								l178:
									position, tokenIndex = position177, tokenIndex177
									if buffer[position] != rune('A') {
										goto l176
									}
									position++
								}
							l177:
								{
									position179, tokenIndex179 := position, tokenIndex
									if buffer[position] != rune('r') {
										goto l180
									}
									position++
									goto l179
								l180:
									position, tokenIndex = position179, tokenIndex179
									if buffer[position] != rune('R') {
										goto l176
									}
									position++
								}
							l179:
								{
									position181, tokenIndex181 := position, tokenIndex
									if buffer[position] != rune('g') {
										goto l182
									}
									position++
									goto l181
								l182:
									position, tokenIndex = position181, tokenIndex181
									if buffer[position] != rune('G') {
										goto l176
									}
									position++
								}
							l181:
								{
									position183, tokenIndex183 := position, tokenIndex
									if buffer[position] != rune('s') {
										goto l184
									}
									position++
									goto l183
								l184:
									position, tokenIndex = position183, tokenIndex183
									if buffer[position] != rune('S') {
										goto l176
									}
									position++
								}
							l183:
								goto l157
							l176:
								position, tokenIndex = position157, tokenIndex157
								{
									position186, tokenIndex186 := position, tokenIndex
									if buffer[position] != rune('d') {
										goto l187
									}
									position++
									goto l186
								l187:
									position, tokenIndex = position186, tokenIndex186
									if buffer[position] != rune('D') {
										goto l185
									}
									position++
								}
							l186:
								{
									position188, tokenIndex188 := position, tokenIndex
									if buffer[position] != rune('e') {
										goto l189
									}
									position++
									goto l188
								l189:
									position, tokenIndex = position188, tokenIndex188
									if buffer[position] != rune('E') {
										goto l185
									}
									position++
								}
							l188:
								{
									position190, tokenIndex190 := position, tokenIndex
									if buffer[position] != rune('c') {
										goto l191
									}
									position++
									goto l190
								l191:
									position, tokenIndex = position190, tokenIndex190
									if buffer[position] != rune('C') {
										goto l185
									}
									position++
								}
							l190:
								{
									position192, tokenIndex192 := position, tokenIndex
									if buffer[position] != rune('l') {
										goto l193
									}
									position++
									goto l192
								l193:
									position, tokenIndex = position192, tokenIndex192
									if buffer[position] != rune('L') {
										goto l185
									}
									position++
								}
							l192:
								goto l157
							l185:
								position, tokenIndex = position157, tokenIndex157
								{
									position195, tokenIndex195 := position, tokenIndex
									if buffer[position] != rune('e') {
										goto l196
									}
									position++
									goto l195
								l196:
									position, tokenIndex = position195, tokenIndex195
									if buffer[position] != rune('E') {
										goto l194
									}
									position++
								}
							l195:
								{
									position197, tokenIndex197 := position, tokenIndex
									if buffer[position] != rune('x') {
										goto l198
									}
									position++
									goto l197
								l198:
									position, tokenIndex = position197, tokenIndex197
									if buffer[position] != rune('X') {
										goto l194
									}
									position++
								}
							l197:
								{
									position199, tokenIndex199 := position, tokenIndex
									if buffer[position] != rune('p') {
										goto l200
									}
									position++
									goto l199
								l200:
									position, tokenIndex = position199, tokenIndex199
									if buffer[position] != rune('P') {
										goto l194
									}
									position++
								}
							l199:
								{
									position201, tokenIndex201 := position, tokenIndex
									if buffer[position] != rune('r') {
										goto l202
									}
									position++
									goto l201
								l202:
									position, tokenIndex = position201, tokenIndex201
									if buffer[position] != rune('R') {
										goto l194
									}
									position++
								}
							l201:
								goto l157
							l194:
								position, tokenIndex = position157, tokenIndex157
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
										goto l203
									}
									position++
								}
							l204:
								{
									position206, tokenIndex206 := position, tokenIndex
									if buffer[position] != rune('n') {
										goto l207
									}
									position++
									goto l206
								l207:
									position, tokenIndex = position206, tokenIndex206
									if buffer[position] != rune('N') {
										goto l203
									}
									position++
								}
							l206:
								goto l157
							l203:
								position, tokenIndex = position157, tokenIndex157
								{
									position209, tokenIndex209 := position, tokenIndex
									if buffer[position] != rune('c') {
										goto l210
									}
									position++
									goto l209
								l210:
									position, tokenIndex = position209, tokenIndex209
									if buffer[position] != rune('C') {
										goto l208
									}
									position++
								}
							l209:
								{
									position211, tokenIndex211 := position, tokenIndex
									if buffer[position] != rune('n') {
										goto l212
									}
									position++
									goto l211
								l212:
									position, tokenIndex = position211, tokenIndex211
									if buffer[position] != rune('N') {
										goto l208
									}
									position++
								}
							l211:
								{
									position213, tokenIndex213 := position, tokenIndex
									if buffer[position] != rune('s') {
										goto l214
									}
									position++
									goto l213
								l214:
									position, tokenIndex = position213, tokenIndex213
									if buffer[position] != rune('S') {
										goto l208
									}
									position++
								}
							l213:
								{
									position215, tokenIndex215 := position, tokenIndex
									if buffer[position] != rune('t') {
										goto l216
									}
									position++
									goto l215
								l216:
									position, tokenIndex = position215, tokenIndex215
									if buffer[position] != rune('T') {
										goto l208
									}
									position++
								}
							l215:
								goto l157
							l208:
								position, tokenIndex = position157, tokenIndex157
								{
									position218, tokenIndex218 := position, tokenIndex
									if buffer[position] != rune('v') {
										goto l219
									}
									position++
									goto l218
								l219:
									position, tokenIndex = position218, tokenIndex218
									if buffer[position] != rune('V') {
										goto l217
									}
									position++
								}
							l218:
								{
									position220, tokenIndex220 := position, tokenIndex
									if buffer[position] != rune('a') {
										goto l221
									}
									position++
									goto l220
								l221:
									position, tokenIndex = position220, tokenIndex220
									if buffer[position] != rune('A') {
										goto l217
									}
									position++
								}
							l220:
								{
									position222, tokenIndex222 := position, tokenIndex
									if buffer[position] != rune('r') {
										goto l223
									}
									position++
									goto l222
								l223:
									position, tokenIndex = position222, tokenIndex222
									if buffer[position] != rune('R') {
										goto l217
									}
									position++
								}
							l222:
								{
									position224, tokenIndex224 := position, tokenIndex
									if buffer[position] != rune('s') {
										goto l225
									}
									position++
									goto l224
								l225:
									position, tokenIndex = position224, tokenIndex224
									if buffer[position] != rune('S') {
										goto l217
									}
									position++
								}
							l224:
								goto l157
							l217:
								position, tokenIndex = position157, tokenIndex157
								{
									position227, tokenIndex227 := position, tokenIndex
									if buffer[position] != rune('v') {
										goto l228
									}
									position++
									goto l227
								l228:
									position, tokenIndex = position227, tokenIndex227
									if buffer[position] != rune('V') {
										goto l226
									}
									position++
								}
							l227:
								{
									position229, tokenIndex229 := position, tokenIndex
									if buffer[position] != rune('a') {
										goto l230
									}
									position++
									goto l229
								l230:
									position, tokenIndex = position229, tokenIndex229
									if buffer[position] != rune('A') {
										goto l226
									}
									position++
								}
							l229:
								{
									position231, tokenIndex231 := position, tokenIndex
									if buffer[position] != rune('l') {
										goto l232
									}
									position++
									goto l231
								l232:
									position, tokenIndex = position231, tokenIndex231
									if buffer[position] != rune('L') {
										goto l226
									}
									position++
								}
							l231:
								{
									position233, tokenIndex233 := position, tokenIndex
									if buffer[position] != rune('u') {
										goto l234
									}
									position++
									goto l233
								l234:
									position, tokenIndex = position233, tokenIndex233
									if buffer[position] != rune('U') {
										goto l226
									}
									position++
								}
							l233:
								goto l157
							l226:
								position, tokenIndex = position157, tokenIndex157
								{
									position236, tokenIndex236 := position, tokenIndex
									if buffer[position] != rune('c') {
										goto l237
									}
									position++
									goto l236
								l237:
									position, tokenIndex = position236, tokenIndex236
									if buffer[position] != rune('C') {
										goto l235
									}
									position++
								}
							l236:
								{
									position238, tokenIndex238 := position, tokenIndex
									if buffer[position] != rune('h') {
										goto l239
									}
									position++
									goto l238
								l239:
									position, tokenIndex = position238, tokenIndex238
									if buffer[position] != rune('H') {
										goto l235
									}
									position++
								}
							l238:
								{
									position240, tokenIndex240 := position, tokenIndex
									if buffer[position] != rune('a') {
										goto l241
									}
									position++
									goto l240
								l241:
									position, tokenIndex = position240, tokenIndex240
									if buffer[position] != rune('A') {
										goto l235
									}
									position++
								}
							l240:
								{
									position242, tokenIndex242 := position, tokenIndex
									if buffer[position] != rune('n') {
										goto l243
									}
									position++
									goto l242
								l243:
									position, tokenIndex = position242, tokenIndex242
									if buffer[position] != rune('N') {
										goto l235
									}
									position++
								}
							l242:
								goto l157
							l235:
								position, tokenIndex = position157, tokenIndex157
								{
									position245, tokenIndex245 := position, tokenIndex
									if buffer[position] != rune('p') {
										goto l246
									}
									position++
									goto l245
								l246:
									position, tokenIndex = position245, tokenIndex245
									if buffer[position] != rune('P') {
										goto l244
									}
									position++
								}
							l245:
								{
									position247, tokenIndex247 := position, tokenIndex
									if buffer[position] != rune('u') {
										goto l248
									}
									position++
									goto l247
								l248:
									position, tokenIndex = position247, tokenIndex247
									if buffer[position] != rune('U') {
										goto l244
									}
									position++
								}
							l247:
								{
									position249, tokenIndex249 := position, tokenIndex
									if buffer[position] != rune('r') {
										goto l250
									}
									position++
									goto l249
								l250:
									position, tokenIndex = position249, tokenIndex249
									if buffer[position] != rune('R') {
										goto l244
									}
									position++
								}
							l249:
								{
									position251, tokenIndex251 := position, tokenIndex
									if buffer[position] != rune('p') {
										goto l252
									}
									position++
									goto l251
								l252:
									position, tokenIndex = position251, tokenIndex251
									if buffer[position] != rune('P') {
										goto l244
									}
									position++
								}
							l251:
								goto l157
							l244:
								position, tokenIndex = position157, tokenIndex157
								{
									position254, tokenIndex254 := position, tokenIndex
									if buffer[position] != rune('f') {
										goto l255
									}
									position++
									goto l254
								l255:
									position, tokenIndex = position254, tokenIndex254
									if buffer[position] != rune('F') {
										goto l253
									}
									position++
								}
							l254:
								{
									position256, tokenIndex256 := position, tokenIndex
									if buffer[position] != rune('l') {
										goto l257
									}
									position++
									goto l256
								l257:
									position, tokenIndex = position256, tokenIndex256
									if buffer[position] != rune('L') {
										goto l253
									}
									position++
								}
							l256:
								{
									position258, tokenIndex258 := position, tokenIndex
									if buffer[position] != rune('d') {
										goto l259
									}
									position++
									goto l258
								l259:
									position, tokenIndex = position258, tokenIndex258
									if buffer[position] != rune('D') {
										goto l253
									}
									position++
								}
							l258:
								{
									position260, tokenIndex260 := position, tokenIndex
									if buffer[position] != rune('s') {
										goto l261
									}
									position++
									goto l260
								l261:
									position, tokenIndex = position260, tokenIndex260
									if buffer[position] != rune('S') {
										goto l253
									}
									position++
								}
							l260:
								goto l157
							l253:
								position, tokenIndex = position157, tokenIndex157
								{
									position263, tokenIndex263 := position, tokenIndex
									if buffer[position] != rune('c') {
										goto l264
									}
									position++
									goto l263
								l264:
									position, tokenIndex = position263, tokenIndex263
									if buffer[position] != rune('C') {
										goto l262
									}
									position++
								}
							l263:
								{
									position265, tokenIndex265 := position, tokenIndex
									if buffer[position] != rune('l') {
										goto l266
									}
									position++
									goto l265
								l266:
									position, tokenIndex = position265, tokenIndex265
									if buffer[position] != rune('L') {
										goto l262
									}
									position++
								}
							l265:
								{
									position267, tokenIndex267 := position, tokenIndex
									if buffer[position] != rune('s') {
										goto l268
									}
									position++
									goto l267
								l268:
									position, tokenIndex = position267, tokenIndex267
									if buffer[position] != rune('S') {
										goto l262
									}
									position++
								}
							l267:
								goto l157
							l262:
								position, tokenIndex = position157, tokenIndex157
								{
									position270, tokenIndex270 := position, tokenIndex
									if buffer[position] != rune('c') {
										goto l271
									}
									position++
									goto l270
								l271:
									position, tokenIndex = position270, tokenIndex270
									if buffer[position] != rune('C') {
										goto l269
									}
									position++
								}
							l270:
								{
									position272, tokenIndex272 := position, tokenIndex
									if buffer[position] != rune('s') {
										goto l273
									}
									position++
									goto l272
								l273:
									position, tokenIndex = position272, tokenIndex272
									if buffer[position] != rune('S') {
										goto l269
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
										goto l269
									}
									position++
								}
							l274:
								{
									position276, tokenIndex276 := position, tokenIndex
									if buffer[position] != rune('s') {
										goto l277
									}
									position++
									goto l276
								l277:
									position, tokenIndex = position276, tokenIndex276
									if buffer[position] != rune('S') {
										goto l269
									}
									position++
								}
							l276:
								goto l157
							l269:
								position, tokenIndex = position157, tokenIndex157
								{
									position279, tokenIndex279 := position, tokenIndex
									if buffer[position] != rune('t') {
										goto l280
									}
									position++
									goto l279
								l280:
									position, tokenIndex = position279, tokenIndex279
									if buffer[position] != rune('T') {
										goto l278
									}
									position++
								}
							l279:
								{
									position281, tokenIndex281 := position, tokenIndex
									if buffer[position] != rune('y') {
										goto l282
									}
									position++
									goto l281
								l282:
									position, tokenIndex = position281, tokenIndex281
									if buffer[position] != rune('Y') {
										goto l278
									}
									position++
								}
							l281:
								{
									position283, tokenIndex283 := position, tokenIndex
									if buffer[position] != rune('p') {
										goto l284
									}
									position++
									goto l283
								l284:
									position, tokenIndex = position283, tokenIndex283
									if buffer[position] != rune('P') {
										goto l278
									}
									position++
								}
							l283:
								{
									position285, tokenIndex285 := position, tokenIndex
									if buffer[position] != rune('e') {
										goto l286
									}
									position++
									goto l285
								l286:
									position, tokenIndex = position285, tokenIndex285
									if buffer[position] != rune('E') {
										goto l278
									}
									position++
								}
							l285:
								goto l157
							l278:
								position, tokenIndex = position157, tokenIndex157
								{
									position288, tokenIndex288 := position, tokenIndex
									if buffer[position] != rune('m') {
										goto l289
									}
									position++
									goto l288
								l289:
									position, tokenIndex = position288, tokenIndex288
									if buffer[position] != rune('M') {
										goto l287
									}
									position++
								}
							l288:
								{
									position290, tokenIndex290 := position, tokenIndex
									if buffer[position] != rune('i') {
										goto l291
									}
									position++
									goto l290
								l291:
									position, tokenIndex = position290, tokenIndex290
									if buffer[position] != rune('I') {
										goto l287
									}
									position++
								}
							l290:
								{
									position292, tokenIndex292 := position, tokenIndex
									if buffer[position] != rune('n') {
										goto l293
									}
									position++
									goto l292
								l293:
									position, tokenIndex = position292, tokenIndex292
									if buffer[position] != rune('N') {
										goto l287
									}
									position++
								}
							l292:
								goto l157
							l287:
								position, tokenIndex = position157, tokenIndex157
								{
									position295, tokenIndex295 := position, tokenIndex
									if buffer[position] != rune('m') {
										goto l296
									}
									position++
									goto l295
								l296:
									position, tokenIndex = position295, tokenIndex295
									if buffer[position] != rune('M') {
										goto l294
									}
									position++
								}
							l295:
								{
									position297, tokenIndex297 := position, tokenIndex
									if buffer[position] != rune('a') {
										goto l298
									}
									position++
									goto l297
								l298:
									position, tokenIndex = position297, tokenIndex297
									if buffer[position] != rune('A') {
										goto l294
									}
									position++
								}
							l297:
								{
									position299, tokenIndex299 := position, tokenIndex
									if buffer[position] != rune('x') {
										goto l300
									}
									position++
									goto l299
								l300:
									position, tokenIndex = position299, tokenIndex299
									if buffer[position] != rune('X') {
										goto l294
									}
									position++
								}
							l299:
								goto l157
							l294:
								position, tokenIndex = position157, tokenIndex157
								{
									position302, tokenIndex302 := position, tokenIndex
									if buffer[position] != rune('b') {
										goto l303
									}
									position++
									goto l302
								l303:
									position, tokenIndex = position302, tokenIndex302
									if buffer[position] != rune('B') {
										goto l301
									}
									position++
								}
							l302:
								{
									position304, tokenIndex304 := position, tokenIndex
									if buffer[position] != rune('i') {
										goto l305
									}
									position++
									goto l304
								l305:
									position, tokenIndex = position304, tokenIndex304
									if buffer[position] != rune('I') {
										goto l301
									}
									position++
								}
							l304:
								{
									position306, tokenIndex306 := position, tokenIndex
									if buffer[position] != rune('n') {
										goto l307
									}
									position++
									goto l306
								l307:
									position, tokenIndex = position306, tokenIndex306
									if buffer[position] != rune('N') {
										goto l301
									}
									position++
								}
							l306:
								{
									position308, tokenIndex308 := position, tokenIndex
									if buffer[position] != rune('f') {
										goto l309
									}
									position++
									goto l308
								l309:
									position, tokenIndex = position308, tokenIndex308
									if buffer[position] != rune('F') {
										goto l301
									}
									position++
								}
							l308:
								goto l157
							l301:
								position, tokenIndex = position157, tokenIndex157
								{
									position311, tokenIndex311 := position, tokenIndex
									if buffer[position] != rune('s') {
										goto l312
									}
									position++
									goto l311
								l312:
									position, tokenIndex = position311, tokenIndex311
									if buffer[position] != rune('S') {
										goto l310
									}
									position++
								}
							l311:
								{
									position313, tokenIndex313 := position, tokenIndex
									if buffer[position] != rune('c') {
										goto l314
									}
									position++
									goto l313
								l314:
									position, tokenIndex = position313, tokenIndex313
									if buffer[position] != rune('C') {
										goto l310
									}
									position++
								}
							l313:
								{
									position315, tokenIndex315 := position, tokenIndex
									if buffer[position] != rune('p') {
										goto l316
									}
									position++
									goto l315
								l316:
									position, tokenIndex = position315, tokenIndex315
									if buffer[position] != rune('P') {
										goto l310
									}
									position++
								}
							l315:
								{
									position317, tokenIndex317 := position, tokenIndex
									if buffer[position] != rune('e') {
										goto l318
									}
									position++
									goto l317
								l318:
									position, tokenIndex = position317, tokenIndex317
									if buffer[position] != rune('E') {
										goto l310
									}
									position++
								}
							l317:
								goto l157
							l310:
								position, tokenIndex = position157, tokenIndex157
								{
									position320, tokenIndex320 := position, tokenIndex
									if buffer[position] != rune('i') {
										goto l321
									}
									position++
									goto l320
								l321:
									position, tokenIndex = position320, tokenIndex320
									if buffer[position] != rune('I') {
										goto l319
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
										goto l319
									}
									position++
								}
							l322:
								{
									position324, tokenIndex324 := position, tokenIndex
									if buffer[position] != rune('i') {
										goto l325
									}
									position++
									goto l324
								l325:
									position, tokenIndex = position324, tokenIndex324
									if buffer[position] != rune('I') {
										goto l319
									}
									position++
								}
							l324:
								{
									position326, tokenIndex326 := position, tokenIndex
									if buffer[position] != rune('t') {
										goto l327
									}
									position++
									goto l326
								l327:
									position, tokenIndex = position326, tokenIndex326
									if buffer[position] != rune('T') {
										goto l319
									}
									position++
								}
							l326:
								goto l157
							l319:
								position, tokenIndex = position157, tokenIndex157
								{
									switch buffer[position] {
									case 'O', 'o':
										{
											position329 := position
											{
												position330, tokenIndex330 := position, tokenIndex
												if buffer[position] != rune('o') {
													goto l331
												}
												position++
												goto l330
											l331:
												position, tokenIndex = position330, tokenIndex330
												if buffer[position] != rune('O') {
													goto l154
												}
												position++
											}
										l330:
											{
												position332, tokenIndex332 := position, tokenIndex
												if buffer[position] != rune('p') {
													goto l333
												}
												position++
												goto l332
											l333:
												position, tokenIndex = position332, tokenIndex332
												if buffer[position] != rune('P') {
													goto l154
												}
												position++
											}
										l332:
											if buffer[position] != rune(' ') {
												goto l154
											}
											position++
											if c := buffer[position]; c < rune('0') || c > rune('9') {
												goto l154
											}
											position++
											add(ruleOpAttr, position329)
										}
										break
									case 'V', 'v':
										{
											position334, tokenIndex334 := position, tokenIndex
											if buffer[position] != rune('v') {
												goto l335
											}
											position++
											goto l334
										l335:
											position, tokenIndex = position334, tokenIndex334
											if buffer[position] != rune('V') {
												goto l154
											}
											position++
										}
									l334:
										{
											position336, tokenIndex336 := position, tokenIndex
											if buffer[position] != rune('a') {
												goto l337
											}
											position++
											goto l336
										l337:
											position, tokenIndex = position336, tokenIndex336
											if buffer[position] != rune('A') {
												goto l154
											}
											position++
										}
									l336:
										{
											position338, tokenIndex338 := position, tokenIndex
											if buffer[position] != rune('l') {
												goto l339
											}
											position++
											goto l338
										l339:
											position, tokenIndex = position338, tokenIndex338
											if buffer[position] != rune('L') {
												goto l154
											}
											position++
										}
									l338:
										break
									case 'I', 'i':
										{
											position340, tokenIndex340 := position, tokenIndex
											if buffer[position] != rune('i') {
												goto l341
											}
											position++
											goto l340
										l341:
											position, tokenIndex = position340, tokenIndex340
											if buffer[position] != rune('I') {
												goto l154
											}
											position++
										}
									l340:
										{
											position342, tokenIndex342 := position, tokenIndex
											if buffer[position] != rune('d') {
												goto l343
											}
											position++
											goto l342
										l343:
											position, tokenIndex = position342, tokenIndex342
											if buffer[position] != rune('D') {
												goto l154
											}
											position++
										}
									l342:
										{
											position344, tokenIndex344 := position, tokenIndex
											if buffer[position] != rune('x') {
												goto l345
											}
											position++
											goto l344
										l345:
											position, tokenIndex = position344, tokenIndex344
											if buffer[position] != rune('X') {
												goto l154
											}
											position++
										}
									l344:
										break
									case 'S', 's':
										{
											position346, tokenIndex346 := position, tokenIndex
											if buffer[position] != rune('s') {
												goto l347
											}
											position++
											goto l346
										l347:
											position, tokenIndex = position346, tokenIndex346
											if buffer[position] != rune('S') {
												goto l154
											}
											position++
										}
									l346:
										{
											position348, tokenIndex348 := position, tokenIndex
											if buffer[position] != rune('i') {
												goto l349
											}
											position++
											goto l348
										l349:
											position, tokenIndex = position348, tokenIndex348
											if buffer[position] != rune('I') {
												goto l154
											}
											position++
										}
									l348:
										{
											position350, tokenIndex350 := position, tokenIndex
											if buffer[position] != rune('z') {
												goto l351
											}
											position++
											goto l350
										l351:
											position, tokenIndex = position350, tokenIndex350
											if buffer[position] != rune('Z') {
												goto l154
											}
											position++
										}
									l350:
										{
											position352, tokenIndex352 := position, tokenIndex
											if buffer[position] != rune('e') {
												goto l353
											}
											position++
											goto l352
										l353:
											position, tokenIndex = position352, tokenIndex352
											if buffer[position] != rune('E') {
												goto l154
											}
											position++
										}
									l352:
										break
									case 'B', 'b':
										{
											position354, tokenIndex354 := position, tokenIndex
											if buffer[position] != rune('b') {
												goto l355
											}
											position++
											goto l354
										l355:
											position, tokenIndex = position354, tokenIndex354
											if buffer[position] != rune('B') {
												goto l154
											}
											position++
										}
									l354:
										{
											position356, tokenIndex356 := position, tokenIndex
											if buffer[position] != rune('p') {
												goto l357
											}
											position++
											goto l356
										l357:
											position, tokenIndex = position356, tokenIndex356
											if buffer[position] != rune('P') {
												goto l154
											}
											position++
										}
									l356:
										{
											position358, tokenIndex358 := position, tokenIndex
											if buffer[position] != rune('o') {
												goto l359
											}
											position++
											goto l358
										l359:
											position, tokenIndex = position358, tokenIndex358
											if buffer[position] != rune('O') {
												goto l154
											}
											position++
										}
									l358:
										{
											position360, tokenIndex360 := position, tokenIndex
											if buffer[position] != rune('s') {
												goto l361
											}
											position++
											goto l360
										l361:
											position, tokenIndex = position360, tokenIndex360
											if buffer[position] != rune('S') {
												goto l154
											}
											position++
										}
									l360:
										break
									case 'C', 'c':
										{
											position362, tokenIndex362 := position, tokenIndex
											if buffer[position] != rune('c') {
												goto l363
											}
											position++
											goto l362
										l363:
											position, tokenIndex = position362, tokenIndex362
											if buffer[position] != rune('C') {
												goto l154
											}
											position++
										}
									l362:
										{
											position364, tokenIndex364 := position, tokenIndex
											if buffer[position] != rune('h') {
												goto l365
											}
											position++
											goto l364
										l365:
											position, tokenIndex = position364, tokenIndex364
											if buffer[position] != rune('H') {
												goto l154
											}
											position++
										}
									l364:
										{
											position366, tokenIndex366 := position, tokenIndex
											if buffer[position] != rune('a') {
												goto l367
											}
											position++
											goto l366
										l367:
											position, tokenIndex = position366, tokenIndex366
											if buffer[position] != rune('A') {
												goto l154
											}
											position++
										}
									l366:
										{
											position368, tokenIndex368 := position, tokenIndex
											if buffer[position] != rune('i') {
												goto l369
											}
											position++
											goto l368
										l369:
											position, tokenIndex = position368, tokenIndex368
											if buffer[position] != rune('I') {
												goto l154
											}
											position++
										}
									l368:
										{
											position370, tokenIndex370 := position, tokenIndex
											if buffer[position] != rune('n') {
												goto l371
											}
											position++
											goto l370
										l371:
											position, tokenIndex = position370, tokenIndex370
											if buffer[position] != rune('N') {
												goto l154
											}
											position++
										}
									l370:
										break
									case 'E', 'e':
										{
											position372, tokenIndex372 := position, tokenIndex
											if buffer[position] != rune('e') {
												goto l373
											}
											position++
											goto l372
										l373:
											position, tokenIndex = position372, tokenIndex372
											if buffer[position] != rune('E') {
												goto l154
											}
											position++
										}
									l372:
										{
											position374, tokenIndex374 := position, tokenIndex
											if buffer[position] != rune('l') {
												goto l375
											}
											position++
											goto l374
										l375:
											position, tokenIndex = position374, tokenIndex374
											if buffer[position] != rune('L') {
												goto l154
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
												goto l154
											}
											position++
										}
									l376:
										{
											position378, tokenIndex378 := position, tokenIndex
											if buffer[position] != rune('s') {
												goto l379
											}
											position++
											goto l378
										l379:
											position, tokenIndex = position378, tokenIndex378
											if buffer[position] != rune('S') {
												goto l154
											}
											position++
										}
									l378:
										break
									case 'D', 'd':
										{
											position380, tokenIndex380 := position, tokenIndex
											if buffer[position] != rune('d') {
												goto l381
											}
											position++
											goto l380
										l381:
											position, tokenIndex = position380, tokenIndex380
											if buffer[position] != rune('D') {
												goto l154
											}
											position++
										}
									l380:
										{
											position382, tokenIndex382 := position, tokenIndex
											if buffer[position] != rune('o') {
												goto l383
											}
											position++
											goto l382
										l383:
											position, tokenIndex = position382, tokenIndex382
											if buffer[position] != rune('O') {
												goto l154
											}
											position++
										}
									l382:
										{
											position384, tokenIndex384 := position, tokenIndex
											if buffer[position] != rune('m') {
												goto l385
											}
											position++
											goto l384
										l385:
											position, tokenIndex = position384, tokenIndex384
											if buffer[position] != rune('M') {
												goto l154
											}
											position++
										}
									l384:
										{
											position386, tokenIndex386 := position, tokenIndex
											if buffer[position] != rune('n') {
												goto l387
											}
											position++
											goto l386
										l387:
											position, tokenIndex = position386, tokenIndex386
											if buffer[position] != rune('N') {
												goto l154
											}
											position++
										}
									l386:
										break
									case 'T', 't':
										{
											position388, tokenIndex388 := position, tokenIndex
											if buffer[position] != rune('t') {
												goto l389
											}
											position++
											goto l388
										l389:
											position, tokenIndex = position388, tokenIndex388
											if buffer[position] != rune('T') {
												goto l154
											}
											position++
										}
									l388:
										{
											position390, tokenIndex390 := position, tokenIndex
											if buffer[position] != rune('y') {
												goto l391
											}
											position++
											goto l390
										l391:
											position, tokenIndex = position390, tokenIndex390
											if buffer[position] != rune('Y') {
												goto l154
											}
											position++
										}
									l390:
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
												goto l154
											}
											position++
										}
									l392:
										{
											position394, tokenIndex394 := position, tokenIndex
											if buffer[position] != rune('e') {
												goto l395
											}
											position++
											goto l394
										l395:
											position, tokenIndex = position394, tokenIndex394
											if buffer[position] != rune('E') {
												goto l154
											}
											position++
										}
									l394:
										break
									case 'M', 'm':
										{
											position396, tokenIndex396 := position, tokenIndex
											if buffer[position] != rune('m') {
												goto l397
											}
											position++
											goto l396
										l397:
											position, tokenIndex = position396, tokenIndex396
											if buffer[position] != rune('M') {
												goto l154
											}
											position++
										}
									l396:
										{
											position398, tokenIndex398 := position, tokenIndex
											if buffer[position] != rune('n') {
												goto l399
											}
											position++
											goto l398
										l399:
											position, tokenIndex = position398, tokenIndex398
											if buffer[position] != rune('N') {
												goto l154
											}
											position++
										}
									l398:
										{
											position400, tokenIndex400 := position, tokenIndex
											if buffer[position] != rune('g') {
												goto l401
											}
											position++
											goto l400
										l401:
											position, tokenIndex = position400, tokenIndex400
											if buffer[position] != rune('G') {
												goto l154
											}
											position++
										}
									l400:
										{
											position402, tokenIndex402 := position, tokenIndex
											if buffer[position] != rune('l') {
												goto l403
											}
											position++
											goto l402
										l403:
											position, tokenIndex = position402, tokenIndex402
											if buffer[position] != rune('L') {
												goto l154
											}
											position++
										}
									l402:
										break
									case 'N', 'n':
										{
											position404, tokenIndex404 := position, tokenIndex
											if buffer[position] != rune('n') {
												goto l405
											}
											position++
											goto l404
										l405:
											position, tokenIndex = position404, tokenIndex404
											if buffer[position] != rune('N') {
												goto l154
											}
											position++
										}
									l404:
										{
											position406, tokenIndex406 := position, tokenIndex
											if buffer[position] != rune('a') {
												goto l407
											}
											position++
											goto l406
										l407:
											position, tokenIndex = position406, tokenIndex406
											if buffer[position] != rune('A') {
												goto l154
											}
											position++
										}
									l406:
										{
											position408, tokenIndex408 := position, tokenIndex
											if buffer[position] != rune('m') {
												goto l409
											}
											position++
											goto l408
										l409:
											position, tokenIndex = position408, tokenIndex408
											if buffer[position] != rune('M') {
												goto l154
											}
											position++
										}
									l408:
										{
											position410, tokenIndex410 := position, tokenIndex
											if buffer[position] != rune('e') {
												goto l411
											}
											position++
											goto l410
										l411:
											position, tokenIndex = position410, tokenIndex410
											if buffer[position] != rune('E') {
												goto l154
											}
											position++
										}
									l410:
										break
									case 'U', 'u':
										{
											position412, tokenIndex412 := position, tokenIndex
											if buffer[position] != rune('u') {
												goto l413
											}
											position++
											goto l412
										l413:
											position, tokenIndex = position412, tokenIndex412
											if buffer[position] != rune('U') {
												goto l154
											}
											position++
										}
									l412:
										{
											position414, tokenIndex414 := position, tokenIndex
											if buffer[position] != rune('n') {
												goto l415
											}
											position++
											goto l414
										l415:
											position, tokenIndex = position414, tokenIndex414
											if buffer[position] != rune('N') {
												goto l154
											}
											position++
										}
									l414:
										{
											position416, tokenIndex416 := position, tokenIndex
											if buffer[position] != rune('q') {
												goto l417
											}
											position++
											goto l416
										l417:
											position, tokenIndex = position416, tokenIndex416
											if buffer[position] != rune('Q') {
												goto l154
											}
											position++
										}
									l416:
										{
											position418, tokenIndex418 := position, tokenIndex
											if buffer[position] != rune('l') {
												goto l419
											}
											position++
											goto l418
										l419:
											position, tokenIndex = position418, tokenIndex418
											if buffer[position] != rune('L') {
												goto l154
											}
											position++
										}
									l418:
										break
									case 'L', 'l':
										{
											position420, tokenIndex420 := position, tokenIndex
											if buffer[position] != rune('l') {
												goto l421
											}
											position++
											goto l420
										l421:
											position, tokenIndex = position420, tokenIndex420
											if buffer[position] != rune('L') {
												goto l154
											}
											position++
										}
									l420:
										{
											position422, tokenIndex422 := position, tokenIndex
											if buffer[position] != rune('a') {
												goto l423
											}
											position++
											goto l422
										l423:
											position, tokenIndex = position422, tokenIndex422
											if buffer[position] != rune('A') {
												goto l154
											}
											position++
										}
									l422:
										{
											position424, tokenIndex424 := position, tokenIndex
											if buffer[position] != rune('b') {
												goto l425
											}
											position++
											goto l424
										l425:
											position, tokenIndex = position424, tokenIndex424
											if buffer[position] != rune('B') {
												goto l154
											}
											position++
										}
									l424:
										{
											position426, tokenIndex426 := position, tokenIndex
											if buffer[position] != rune('l') {
												goto l427
											}
											position++
											goto l426
										l427:
											position, tokenIndex = position426, tokenIndex426
											if buffer[position] != rune('L') {
												goto l154
											}
											position++
										}
									l426:
										break
									case 'P', 'p':
										{
											position428, tokenIndex428 := position, tokenIndex
											if buffer[position] != rune('p') {
												goto l429
											}
											position++
											goto l428
										l429:
											position, tokenIndex = position428, tokenIndex428
											if buffer[position] != rune('P') {
												goto l154
											}
											position++
										}
									l428:
										{
											position430, tokenIndex430 := position, tokenIndex
											if buffer[position] != rune('t') {
												goto l431
											}
											position++
											goto l430
										l431:
											position, tokenIndex = position430, tokenIndex430
											if buffer[position] != rune('T') {
												goto l154
											}
											position++
										}
									l430:
										{
											position432, tokenIndex432 := position, tokenIndex
											if buffer[position] != rune('d') {
												goto l433
											}
											position++
											goto l432
										l433:
											position, tokenIndex = position432, tokenIndex432
											if buffer[position] != rune('D') {
												goto l154
											}
											position++
										}
									l432:
										break
									case 'F', 'f':
										{
											position434, tokenIndex434 := position, tokenIndex
											if buffer[position] != rune('f') {
												goto l435
											}
											position++
											goto l434
										l435:
											position, tokenIndex = position434, tokenIndex434
											if buffer[position] != rune('F') {
												goto l154
											}
											position++
										}
									l434:
										{
											position436, tokenIndex436 := position, tokenIndex
											if buffer[position] != rune('n') {
												goto l437
											}
											position++
											goto l436
										l437:
											position, tokenIndex = position436, tokenIndex436
											if buffer[position] != rune('N') {
												goto l154
											}
											position++
										}
									l436:
										{
											position438, tokenIndex438 := position, tokenIndex
											if buffer[position] != rune('c') {
												goto l439
											}
											position++
											goto l438
										l439:
											position, tokenIndex = position438, tokenIndex438
											if buffer[position] != rune('C') {
												goto l154
											}
											position++
										}
									l438:
										{
											position440, tokenIndex440 := position, tokenIndex
											if buffer[position] != rune('s') {
												goto l441
											}
											position++
											goto l440
										l441:
											position, tokenIndex = position440, tokenIndex440
											if buffer[position] != rune('S') {
												goto l154
											}
											position++
										}
									l440:
										break
									case 'R', 'r':
										{
											position442, tokenIndex442 := position, tokenIndex
											if buffer[position] != rune('r') {
												goto l443
											}
											position++
											goto l442
										l443:
											position, tokenIndex = position442, tokenIndex442
											if buffer[position] != rune('R') {
												goto l154
											}
											position++
										}
									l442:
										{
											position444, tokenIndex444 := position, tokenIndex
											if buffer[position] != rune('e') {
												goto l445
											}
											position++
											goto l444
										l445:
											position, tokenIndex = position444, tokenIndex444
											if buffer[position] != rune('E') {
												goto l154
											}
											position++
										}
									l444:
										{
											position446, tokenIndex446 := position, tokenIndex
											if buffer[position] != rune('f') {
												goto l447
											}
											position++
											goto l446
										l447:
											position, tokenIndex = position446, tokenIndex446
											if buffer[position] != rune('F') {
												goto l154
											}
											position++
										}
									l446:
										{
											position448, tokenIndex448 := position, tokenIndex
											if buffer[position] != rune('d') {
												goto l449
											}
											position++
											goto l448
										l449:
											position, tokenIndex = position448, tokenIndex448
											if buffer[position] != rune('D') {
												goto l154
											}
											position++
										}
									l448:
										break
									case 'A', 'a':
										{
											position450, tokenIndex450 := position, tokenIndex
											if buffer[position] != rune('a') {
												goto l451
											}
											position++
											goto l450
										l451:
											position, tokenIndex = position450, tokenIndex450
											if buffer[position] != rune('A') {
												goto l154
											}
											position++
										}
									l450:
										{
											position452, tokenIndex452 := position, tokenIndex
											if buffer[position] != rune('r') {
												goto l453
											}
											position++
											goto l452
										l453:
											position, tokenIndex = position452, tokenIndex452
											if buffer[position] != rune('R') {
												goto l154
											}
											position++
										}
									l452:
										{
											position454, tokenIndex454 := position, tokenIndex
											if buffer[position] != rune('g') {
												goto l455
											}
											position++
											goto l454
										l455:
											position, tokenIndex = position454, tokenIndex454
											if buffer[position] != rune('G') {
												goto l154
											}
											position++
										}
									l454:
										{
											position456, tokenIndex456 := position, tokenIndex
											if buffer[position] != rune('t') {
												goto l457
											}
											position++
											goto l456
										l457:
											position, tokenIndex = position456, tokenIndex456
											if buffer[position] != rune('T') {
												goto l154
											}
											position++
										}
									l456:
										break
									default:
										if !_rules[ruleInteger]() {
											goto l154
										}
										break
									}
								}

							}
						l157:
							add(ruleNodeName, position156)
						}
						if !_rules[rulews]() {
							goto l154
						}
						if buffer[position] != rune(':') {
							goto l154
						}
						position++
						if !_rules[rulews]() {
							goto l154
						}
						if !_rules[ruleNode]() {
							goto l154
						}
						if !_rules[rulews]() {
							goto l154
						}
						add(ruleNodeAttr, position155)
					}
					goto l59
				l154:
					position, tokenIndex = position59, tokenIndex59
					{
						position459 := position
						if buffer[position] != rune('s') {
							goto l458
						}
						position++
						if buffer[position] != rune('r') {
							goto l458
						}
						position++
						if buffer[position] != rune('c') {
							goto l458
						}
						position++
						if buffer[position] != rune('p') {
							goto l458
						}
						position++
						if buffer[position] != rune(':') {
							goto l458
						}
						position++
						if !_rules[rulews]() {
							goto l458
						}
						{
							position460 := position
							{
								position461, tokenIndex461 := position, tokenIndex
								{
									position463, tokenIndex463 := position, tokenIndex
									if buffer[position] != rune('_') {
										goto l463
									}
									position++
									goto l464
								l463:
									position, tokenIndex = position463, tokenIndex463
								}
							l464:
								{
									switch buffer[position] {
									case '_':
										if buffer[position] != rune('_') {
											goto l462
										}
										position++
										break
									case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
										if c := buffer[position]; c < rune('0') || c > rune('9') {
											goto l462
										}
										position++
										break
									case 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z':
										if c := buffer[position]; c < rune('a') || c > rune('z') {
											goto l462
										}
										position++
										break
									case '-':
										if buffer[position] != rune('-') {
											goto l462
										}
										position++
										break
									default:
										if c := buffer[position]; c < rune('A') || c > rune('Z') {
											goto l462
										}
										position++
										break
									}
								}

							l465:
								{
									position466, tokenIndex466 := position, tokenIndex
									{
										switch buffer[position] {
										case '_':
											if buffer[position] != rune('_') {
												goto l466
											}
											position++
											break
										case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
											if c := buffer[position]; c < rune('0') || c > rune('9') {
												goto l466
											}
											position++
											break
										case 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z':
											if c := buffer[position]; c < rune('a') || c > rune('z') {
												goto l466
											}
											position++
											break
										case '-':
											if buffer[position] != rune('-') {
												goto l466
											}
											position++
											break
										default:
											if c := buffer[position]; c < rune('A') || c > rune('Z') {
												goto l466
											}
											position++
											break
										}
									}

									goto l465
								l466:
									position, tokenIndex = position466, tokenIndex466
								}
								if buffer[position] != rune('.') {
									goto l462
								}
								position++
								{
									switch buffer[position] {
									case '_':
										if buffer[position] != rune('_') {
											goto l462
										}
										position++
										break
									case 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z':
										if c := buffer[position]; c < rune('a') || c > rune('z') {
											goto l462
										}
										position++
										break
									default:
										if c := buffer[position]; c < rune('A') || c > rune('Z') {
											goto l462
										}
										position++
										break
									}
								}

							l469:
								{
									position470, tokenIndex470 := position, tokenIndex
									{
										switch buffer[position] {
										case '_':
											if buffer[position] != rune('_') {
												goto l470
											}
											position++
											break
										case 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z':
											if c := buffer[position]; c < rune('a') || c > rune('z') {
												goto l470
											}
											position++
											break
										default:
											if c := buffer[position]; c < rune('A') || c > rune('Z') {
												goto l470
											}
											position++
											break
										}
									}

									goto l469
								l470:
									position, tokenIndex = position470, tokenIndex470
								}
								goto l461
							l462:
								position, tokenIndex = position461, tokenIndex461
								if buffer[position] != rune('<') {
									goto l458
								}
								position++
								{
									position473, tokenIndex473 := position, tokenIndex
									if buffer[position] != rune('b') {
										goto l474
									}
									position++
									goto l473
								l474:
									position, tokenIndex = position473, tokenIndex473
									if buffer[position] != rune('B') {
										goto l458
									}
									position++
								}
							l473:
								{
									position475, tokenIndex475 := position, tokenIndex
									if buffer[position] != rune('u') {
										goto l476
									}
									position++
									goto l475
								l476:
									position, tokenIndex = position475, tokenIndex475
									if buffer[position] != rune('U') {
										goto l458
									}
									position++
								}
							l475:
								{
									position477, tokenIndex477 := position, tokenIndex
									if buffer[position] != rune('i') {
										goto l478
									}
									position++
									goto l477
								l478:
									position, tokenIndex = position477, tokenIndex477
									if buffer[position] != rune('I') {
										goto l458
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
										goto l458
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
										goto l458
									}
									position++
								}
							l481:
								if buffer[position] != rune('-') {
									goto l458
								}
								position++
								{
									position483, tokenIndex483 := position, tokenIndex
									if buffer[position] != rune('i') {
										goto l484
									}
									position++
									goto l483
								l484:
									position, tokenIndex = position483, tokenIndex483
									if buffer[position] != rune('I') {
										goto l458
									}
									position++
								}
							l483:
								{
									position485, tokenIndex485 := position, tokenIndex
									if buffer[position] != rune('n') {
										goto l486
									}
									position++
									goto l485
								l486:
									position, tokenIndex = position485, tokenIndex485
									if buffer[position] != rune('N') {
										goto l458
									}
									position++
								}
							l485:
								if buffer[position] != rune('>') {
									goto l458
								}
								position++
							}
						l461:
							if buffer[position] != rune(':') {
								goto l458
							}
							position++
							if !_rules[ruleInteger]() {
								goto l458
							}
							if !_rules[rulews]() {
								goto l458
							}
							add(ruleFileRef, position460)
						}
						add(ruleSourceAttr, position459)
					}
					goto l59
				l458:
					position, tokenIndex = position59, tokenIndex59
					{
						position488 := position
						{
							switch buffer[position] {
							case 'P', 'p':
								{
									position490, tokenIndex490 := position, tokenIndex
									if buffer[position] != rune('p') {
										goto l491
									}
									position++
									goto l490
								l491:
									position, tokenIndex = position490, tokenIndex490
									if buffer[position] != rune('P') {
										goto l487
									}
									position++
								}
							l490:
								{
									position492, tokenIndex492 := position, tokenIndex
									if buffer[position] != rune('r') {
										goto l493
									}
									position++
									goto l492
								l493:
									position, tokenIndex = position492, tokenIndex492
									if buffer[position] != rune('R') {
										goto l487
									}
									position++
								}
							l492:
								{
									position494, tokenIndex494 := position, tokenIndex
									if buffer[position] != rune('e') {
										goto l495
									}
									position++
									goto l494
								l495:
									position, tokenIndex = position494, tokenIndex494
									if buffer[position] != rune('E') {
										goto l487
									}
									position++
								}
							l494:
								{
									position496, tokenIndex496 := position, tokenIndex
									if buffer[position] != rune('c') {
										goto l497
									}
									position++
									goto l496
								l497:
									position, tokenIndex = position496, tokenIndex496
									if buffer[position] != rune('C') {
										goto l487
									}
									position++
								}
							l496:
								break
							case 'U', 'u':
								{
									position498, tokenIndex498 := position, tokenIndex
									if buffer[position] != rune('u') {
										goto l499
									}
									position++
									goto l498
								l499:
									position, tokenIndex = position498, tokenIndex498
									if buffer[position] != rune('U') {
										goto l487
									}
									position++
								}
							l498:
								{
									position500, tokenIndex500 := position, tokenIndex
									if buffer[position] != rune('s') {
										goto l501
									}
									position++
									goto l500
								l501:
									position, tokenIndex = position500, tokenIndex500
									if buffer[position] != rune('S') {
										goto l487
									}
									position++
								}
							l500:
								{
									position502, tokenIndex502 := position, tokenIndex
									if buffer[position] != rune('e') {
										goto l503
									}
									position++
									goto l502
								l503:
									position, tokenIndex = position502, tokenIndex502
									if buffer[position] != rune('E') {
										goto l487
									}
									position++
								}
							l502:
								{
									position504, tokenIndex504 := position, tokenIndex
									if buffer[position] != rune('d') {
										goto l505
									}
									position++
									goto l504
								l505:
									position, tokenIndex = position504, tokenIndex504
									if buffer[position] != rune('D') {
										goto l487
									}
									position++
								}
							l504:
								break
							case 'L', 'l':
								{
									position506, tokenIndex506 := position, tokenIndex
									if buffer[position] != rune('l') {
										goto l507
									}
									position++
									goto l506
								l507:
									position, tokenIndex = position506, tokenIndex506
									if buffer[position] != rune('L') {
										goto l487
									}
									position++
								}
							l506:
								{
									position508, tokenIndex508 := position, tokenIndex
									if buffer[position] != rune('o') {
										goto l509
									}
									position++
									goto l508
								l509:
									position, tokenIndex = position508, tokenIndex508
									if buffer[position] != rune('O') {
										goto l487
									}
									position++
								}
							l508:
								{
									position510, tokenIndex510 := position, tokenIndex
									if buffer[position] != rune('w') {
										goto l511
									}
									position++
									goto l510
								l511:
									position, tokenIndex = position510, tokenIndex510
									if buffer[position] != rune('W') {
										goto l487
									}
									position++
								}
							l510:
								break
							default:
								{
									position512, tokenIndex512 := position, tokenIndex
									if buffer[position] != rune('b') {
										goto l513
									}
									position++
									goto l512
								l513:
									position, tokenIndex = position512, tokenIndex512
									if buffer[position] != rune('B') {
										goto l487
									}
									position++
								}
							l512:
								{
									position514, tokenIndex514 := position, tokenIndex
									if buffer[position] != rune('a') {
										goto l515
									}
									position++
									goto l514
								l515:
									position, tokenIndex = position514, tokenIndex514
									if buffer[position] != rune('A') {
										goto l487
									}
									position++
								}
							l514:
								{
									position516, tokenIndex516 := position, tokenIndex
									if buffer[position] != rune('s') {
										goto l517
									}
									position++
									goto l516
								l517:
									position, tokenIndex = position516, tokenIndex516
									if buffer[position] != rune('S') {
										goto l487
									}
									position++
								}
							l516:
								{
									position518, tokenIndex518 := position, tokenIndex
									if buffer[position] != rune('e') {
										goto l519
									}
									position++
									goto l518
								l519:
									position, tokenIndex = position518, tokenIndex518
									if buffer[position] != rune('E') {
										goto l487
									}
									position++
								}
							l518:
								{
									position520, tokenIndex520 := position, tokenIndex
									if buffer[position] != rune('s') {
										goto l521
									}
									position++
									goto l520
								l521:
									position, tokenIndex = position520, tokenIndex520
									if buffer[position] != rune('S') {
										goto l487
									}
									position++
								}
							l520:
								break
							}
						}

						if !_rules[rulews]() {
							goto l487
						}
						if buffer[position] != rune(':') {
							goto l487
						}
						position++
						if !_rules[rulews]() {
							goto l487
						}
						if !_rules[ruleInteger]() {
							goto l487
						}
						if !_rules[rulews]() {
							goto l487
						}
						add(ruleIntAttr, position488)
					}
					goto l59
				l487:
					position, tokenIndex = position59, tokenIndex59
					{
						position523 := position
						{
							position524, tokenIndex524 := position, tokenIndex
							if buffer[position] != rune('s') {
								goto l525
							}
							position++
							goto l524
						l525:
							position, tokenIndex = position524, tokenIndex524
							if buffer[position] != rune('S') {
								goto l522
							}
							position++
						}
					l524:
						{
							position526, tokenIndex526 := position, tokenIndex
							if buffer[position] != rune('t') {
								goto l527
							}
							position++
							goto l526
						l527:
							position, tokenIndex = position526, tokenIndex526
							if buffer[position] != rune('T') {
								goto l522
							}
							position++
						}
					l526:
						{
							position528, tokenIndex528 := position, tokenIndex
							if buffer[position] != rune('r') {
								goto l529
							}
							position++
							goto l528
						l529:
							position, tokenIndex = position528, tokenIndex528
							if buffer[position] != rune('R') {
								goto l522
							}
							position++
						}
					l528:
						{
							position530, tokenIndex530 := position, tokenIndex
							if buffer[position] != rune('g') {
								goto l531
							}
							position++
							goto l530
						l531:
							position, tokenIndex = position530, tokenIndex530
							if buffer[position] != rune('G') {
								goto l522
							}
							position++
						}
					l530:
						if buffer[position] != rune(':') {
							goto l522
						}
						position++
						if !matchDot() {
							goto l522
						}
					l532:
						{
							position533, tokenIndex533 := position, tokenIndex
							if !matchDot() {
								goto l533
							}
							goto l532
						l533:
							position, tokenIndex = position533, tokenIndex533
						}
						add(ruleStringAttr, position523)
					}
					goto l59
				l522:
					position, tokenIndex = position59, tokenIndex59
					{
						position535 := position
						{
							position536, tokenIndex536 := position, tokenIndex
							{
								position538, tokenIndex538 := position, tokenIndex
								if buffer[position] != rune('l') {
									goto l539
								}
								position++
								goto l538
							l539:
								position, tokenIndex = position538, tokenIndex538
								if buffer[position] != rune('L') {
									goto l537
								}
								position++
							}
						l538:
							{
								position540, tokenIndex540 := position, tokenIndex
								if buffer[position] != rune('o') {
									goto l541
								}
								position++
								goto l540
							l541:
								position, tokenIndex = position540, tokenIndex540
								if buffer[position] != rune('O') {
									goto l537
								}
								position++
							}
						l540:
							{
								position542, tokenIndex542 := position, tokenIndex
								if buffer[position] != rune('w') {
									goto l543
								}
								position++
								goto l542
							l543:
								position, tokenIndex = position542, tokenIndex542
								if buffer[position] != rune('W') {
									goto l537
								}
								position++
							}
						l542:
							goto l536
						l537:
							position, tokenIndex = position536, tokenIndex536
							{
								position544, tokenIndex544 := position, tokenIndex
								if buffer[position] != rune('h') {
									goto l545
								}
								position++
								goto l544
							l545:
								position, tokenIndex = position544, tokenIndex544
								if buffer[position] != rune('H') {
									goto l534
								}
								position++
							}
						l544:
							{
								position546, tokenIndex546 := position, tokenIndex
								if buffer[position] != rune('i') {
									goto l547
								}
								position++
								goto l546
							l547:
								position, tokenIndex = position546, tokenIndex546
								if buffer[position] != rune('I') {
									goto l534
								}
								position++
							}
						l546:
							{
								position548, tokenIndex548 := position, tokenIndex
								if buffer[position] != rune('g') {
									goto l549
								}
								position++
								goto l548
							l549:
								position, tokenIndex = position548, tokenIndex548
								if buffer[position] != rune('G') {
									goto l534
								}
								position++
							}
						l548:
							{
								position550, tokenIndex550 := position, tokenIndex
								if buffer[position] != rune('h') {
									goto l551
								}
								position++
								goto l550
							l551:
								position, tokenIndex = position550, tokenIndex550
								if buffer[position] != rune('H') {
									goto l534
								}
								position++
							}
						l550:
						}
					l536:
						if !_rules[rulews]() {
							goto l534
						}
						if buffer[position] != rune(':') {
							goto l534
						}
						position++
						if !_rules[rulews]() {
							goto l534
						}
						{
							position552, tokenIndex552 := position, tokenIndex
							if buffer[position] != rune('-') {
								goto l552
							}
							position++
							goto l553
						l552:
							position, tokenIndex = position552, tokenIndex552
						}
					l553:
						if !_rules[ruleInteger]() {
							goto l534
						}
						if !_rules[rulews]() {
							goto l534
						}
						add(ruleIntAttr3, position535)
					}
					goto l59
				l534:
					position, tokenIndex = position59, tokenIndex59
					{
						position555 := position
						{
							position556, tokenIndex556 := position, tokenIndex
							if buffer[position] != rune('l') {
								goto l557
							}
							position++
							goto l556
						l557:
							position, tokenIndex = position556, tokenIndex556
							if buffer[position] != rune('L') {
								goto l554
							}
							position++
						}
					l556:
						{
							position558, tokenIndex558 := position, tokenIndex
							if buffer[position] != rune('n') {
								goto l559
							}
							position++
							goto l558
						l559:
							position, tokenIndex = position558, tokenIndex558
							if buffer[position] != rune('N') {
								goto l554
							}
							position++
						}
					l558:
						{
							position560, tokenIndex560 := position, tokenIndex
							if buffer[position] != rune('g') {
								goto l561
							}
							position++
							goto l560
						l561:
							position, tokenIndex = position560, tokenIndex560
							if buffer[position] != rune('G') {
								goto l554
							}
							position++
						}
					l560:
						{
							position562, tokenIndex562 := position, tokenIndex
							if buffer[position] != rune('t') {
								goto l563
							}
							position++
							goto l562
						l563:
							position, tokenIndex = position562, tokenIndex562
							if buffer[position] != rune('T') {
								goto l554
							}
							position++
						}
					l562:
						if !_rules[rulews]() {
							goto l554
						}
						if buffer[position] != rune(':') {
							goto l554
						}
						position++
						if !_rules[rulews]() {
							goto l554
						}
						if !_rules[ruleInteger]() {
							goto l554
						}
						if !_rules[rulews]() {
							goto l554
						}
						add(ruleLngtAttr, position555)
					}
					goto l59
				l554:
					position, tokenIndex = position59, tokenIndex59
					{
						position565 := position
						{
							position566, tokenIndex566 := position, tokenIndex
							if buffer[position] != rune('a') {
								goto l567
							}
							position++
							goto l566
						l567:
							position, tokenIndex = position566, tokenIndex566
							if buffer[position] != rune('A') {
								goto l564
							}
							position++
						}
					l566:
						{
							position568, tokenIndex568 := position, tokenIndex
							if buffer[position] != rune('c') {
								goto l569
							}
							position++
							goto l568
						l569:
							position, tokenIndex = position568, tokenIndex568
							if buffer[position] != rune('C') {
								goto l564
							}
							position++
						}
					l568:
						{
							position570, tokenIndex570 := position, tokenIndex
							if buffer[position] != rune('c') {
								goto l571
							}
							position++
							goto l570
						l571:
							position, tokenIndex = position570, tokenIndex570
							if buffer[position] != rune('C') {
								goto l564
							}
							position++
						}
					l570:
						{
							position572, tokenIndex572 := position, tokenIndex
							if buffer[position] != rune('s') {
								goto l573
							}
							position++
							goto l572
						l573:
							position, tokenIndex = position572, tokenIndex572
							if buffer[position] != rune('S') {
								goto l564
							}
							position++
						}
					l572:
						if !_rules[rulews]() {
							goto l564
						}
						if buffer[position] != rune(':') {
							goto l564
						}
						position++
						if !_rules[rulews]() {
							goto l564
						}
						{
							position574, tokenIndex574 := position, tokenIndex
							{
								position576, tokenIndex576 := position, tokenIndex
								if buffer[position] != rune('p') {
									goto l577
								}
								position++
								goto l576
							l577:
								position, tokenIndex = position576, tokenIndex576
								if buffer[position] != rune('P') {
									goto l575
								}
								position++
							}
						l576:
							{
								position578, tokenIndex578 := position, tokenIndex
								if buffer[position] != rune('r') {
									goto l579
								}
								position++
								goto l578
							l579:
								position, tokenIndex = position578, tokenIndex578
								if buffer[position] != rune('R') {
									goto l575
								}
								position++
							}
						l578:
							{
								position580, tokenIndex580 := position, tokenIndex
								if buffer[position] != rune('i') {
									goto l581
								}
								position++
								goto l580
							l581:
								position, tokenIndex = position580, tokenIndex580
								if buffer[position] != rune('I') {
									goto l575
								}
								position++
							}
						l580:
							{
								position582, tokenIndex582 := position, tokenIndex
								if buffer[position] != rune('v') {
									goto l583
								}
								position++
								goto l582
							l583:
								position, tokenIndex = position582, tokenIndex582
								if buffer[position] != rune('V') {
									goto l575
								}
								position++
							}
						l582:
							goto l574
						l575:
							position, tokenIndex = position574, tokenIndex574
							{
								position584, tokenIndex584 := position, tokenIndex
								if buffer[position] != rune('p') {
									goto l585
								}
								position++
								goto l584
							l585:
								position, tokenIndex = position584, tokenIndex584
								if buffer[position] != rune('P') {
									goto l564
								}
								position++
							}
						l584:
							{
								position586, tokenIndex586 := position, tokenIndex
								if buffer[position] != rune('u') {
									goto l587
								}
								position++
								goto l586
							l587:
								position, tokenIndex = position586, tokenIndex586
								if buffer[position] != rune('U') {
									goto l564
								}
								position++
							}
						l586:
							{
								position588, tokenIndex588 := position, tokenIndex
								if buffer[position] != rune('b') {
									goto l589
								}
								position++
								goto l588
							l589:
								position, tokenIndex = position588, tokenIndex588
								if buffer[position] != rune('B') {
									goto l564
								}
								position++
							}
						l588:
						}
					l574:
						add(ruleAccsAttr, position565)
					}
					goto l59
				l564:
					position, tokenIndex = position59, tokenIndex59
					{
						switch buffer[position] {
						case 'A', 'a':
							{
								position591 := position
								{
									position592, tokenIndex592 := position, tokenIndex
									if buffer[position] != rune('a') {
										goto l593
									}
									position++
									goto l592
								l593:
									position, tokenIndex = position592, tokenIndex592
									if buffer[position] != rune('A') {
										goto l57
									}
									position++
								}
							l592:
								{
									position594, tokenIndex594 := position, tokenIndex
									if buffer[position] != rune('l') {
										goto l595
									}
									position++
									goto l594
								l595:
									position, tokenIndex = position594, tokenIndex594
									if buffer[position] != rune('L') {
										goto l57
									}
									position++
								}
							l594:
								{
									position596, tokenIndex596 := position, tokenIndex
									if buffer[position] != rune('g') {
										goto l597
									}
									position++
									goto l596
								l597:
									position, tokenIndex = position596, tokenIndex596
									if buffer[position] != rune('G') {
										goto l57
									}
									position++
								}
							l596:
								{
									position598, tokenIndex598 := position, tokenIndex
									if buffer[position] != rune('n') {
										goto l599
									}
									position++
									goto l598
								l599:
									position, tokenIndex = position598, tokenIndex598
									if buffer[position] != rune('N') {
										goto l57
									}
									position++
								}
							l598:
								if buffer[position] != rune(':') {
									goto l57
								}
								position++
								if !_rules[rulews]() {
									goto l57
								}
								if c := buffer[position]; c < rune('0') || c > rune('9') {
									goto l57
								}
								position++
							l600:
								{
									position601, tokenIndex601 := position, tokenIndex
									if c := buffer[position]; c < rune('0') || c > rune('9') {
										goto l601
									}
									position++
									goto l600
								l601:
									position, tokenIndex = position601, tokenIndex601
								}
								if !_rules[rulews]() {
									goto l57
								}
								add(ruleIntAttr2, position591)
							}
							break
						case 'Q', 'q':
							{
								position602 := position
								{
									position603, tokenIndex603 := position, tokenIndex
									if buffer[position] != rune('q') {
										goto l604
									}
									position++
									goto l603
								l604:
									position, tokenIndex = position603, tokenIndex603
									if buffer[position] != rune('Q') {
										goto l57
									}
									position++
								}
							l603:
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
										goto l57
									}
									position++
								}
							l605:
								{
									position607, tokenIndex607 := position, tokenIndex
									if buffer[position] != rune('a') {
										goto l608
									}
									position++
									goto l607
								l608:
									position, tokenIndex = position607, tokenIndex607
									if buffer[position] != rune('A') {
										goto l57
									}
									position++
								}
							l607:
								{
									position609, tokenIndex609 := position, tokenIndex
									if buffer[position] != rune('l') {
										goto l610
									}
									position++
									goto l609
								l610:
									position, tokenIndex = position609, tokenIndex609
									if buffer[position] != rune('L') {
										goto l57
									}
									position++
								}
							l609:
								if !_rules[rulews]() {
									goto l57
								}
								if buffer[position] != rune(':') {
									goto l57
								}
								position++
								if !_rules[rulews]() {
									goto l57
								}
								{
									switch buffer[position] {
									case 'R', 'r':
										{
											position614, tokenIndex614 := position, tokenIndex
											if buffer[position] != rune('r') {
												goto l615
											}
											position++
											goto l614
										l615:
											position, tokenIndex = position614, tokenIndex614
											if buffer[position] != rune('R') {
												goto l57
											}
											position++
										}
									l614:
										break
									case 'C':
										if buffer[position] != rune('C') {
											goto l57
										}
										position++
										break
									case 'c':
										if buffer[position] != rune('c') {
											goto l57
										}
										position++
										break
									default:
										{
											position616, tokenIndex616 := position, tokenIndex
											if buffer[position] != rune('v') {
												goto l617
											}
											position++
											goto l616
										l617:
											position, tokenIndex = position616, tokenIndex616
											if buffer[position] != rune('V') {
												goto l57
											}
											position++
										}
									l616:
										break
									}
								}

							l611:
								{
									position612, tokenIndex612 := position, tokenIndex
									{
										switch buffer[position] {
										case 'R', 'r':
											{
												position619, tokenIndex619 := position, tokenIndex
												if buffer[position] != rune('r') {
													goto l620
												}
												position++
												goto l619
											l620:
												position, tokenIndex = position619, tokenIndex619
												if buffer[position] != rune('R') {
													goto l612
												}
												position++
											}
										l619:
											break
										case 'C':
											if buffer[position] != rune('C') {
												goto l612
											}
											position++
											break
										case 'c':
											if buffer[position] != rune('c') {
												goto l612
											}
											position++
											break
										default:
											{
												position621, tokenIndex621 := position, tokenIndex
												if buffer[position] != rune('v') {
													goto l622
												}
												position++
												goto l621
											l622:
												position, tokenIndex = position621, tokenIndex621
												if buffer[position] != rune('V') {
													goto l612
												}
												position++
											}
										l621:
											break
										}
									}

									goto l611
								l612:
									position, tokenIndex = position612, tokenIndex612
								}
								if !_rules[rulews]() {
									goto l57
								}
								add(ruleQualAttr, position602)
							}
							break
						case 'L', 'l':
							{
								position623 := position
								{
									position624, tokenIndex624 := position, tokenIndex
									if buffer[position] != rune('l') {
										goto l625
									}
									position++
									goto l624
								l625:
									position, tokenIndex = position624, tokenIndex624
									if buffer[position] != rune('L') {
										goto l57
									}
									position++
								}
							l624:
								{
									position626, tokenIndex626 := position, tokenIndex
									if buffer[position] != rune('i') {
										goto l627
									}
									position++
									goto l626
								l627:
									position, tokenIndex = position626, tokenIndex626
									if buffer[position] != rune('I') {
										goto l57
									}
									position++
								}
							l626:
								{
									position628, tokenIndex628 := position, tokenIndex
									if buffer[position] != rune('n') {
										goto l629
									}
									position++
									goto l628
								l629:
									position, tokenIndex = position628, tokenIndex628
									if buffer[position] != rune('N') {
										goto l57
									}
									position++
								}
							l628:
								{
									position630, tokenIndex630 := position, tokenIndex
									if buffer[position] != rune('k') {
										goto l631
									}
									position++
									goto l630
								l631:
									position, tokenIndex = position630, tokenIndex630
									if buffer[position] != rune('K') {
										goto l57
									}
									position++
								}
							l630:
								if !_rules[rulews]() {
									goto l57
								}
								if buffer[position] != rune(':') {
									goto l57
								}
								position++
								if !_rules[rulews]() {
									goto l57
								}
								{
									position632, tokenIndex632 := position, tokenIndex
									{
										position634, tokenIndex634 := position, tokenIndex
										if buffer[position] != rune('e') {
											goto l635
										}
										position++
										goto l634
									l635:
										position, tokenIndex = position634, tokenIndex634
										if buffer[position] != rune('E') {
											goto l633
										}
										position++
									}
								l634:
									{
										position636, tokenIndex636 := position, tokenIndex
										if buffer[position] != rune('x') {
											goto l637
										}
										position++
										goto l636
									l637:
										position, tokenIndex = position636, tokenIndex636
										if buffer[position] != rune('X') {
											goto l633
										}
										position++
									}
								l636:
									{
										position638, tokenIndex638 := position, tokenIndex
										if buffer[position] != rune('t') {
											goto l639
										}
										position++
										goto l638
									l639:
										position, tokenIndex = position638, tokenIndex638
										if buffer[position] != rune('T') {
											goto l633
										}
										position++
									}
								l638:
									{
										position640, tokenIndex640 := position, tokenIndex
										if buffer[position] != rune('e') {
											goto l641
										}
										position++
										goto l640
									l641:
										position, tokenIndex = position640, tokenIndex640
										if buffer[position] != rune('E') {
											goto l633
										}
										position++
									}
								l640:
									{
										position642, tokenIndex642 := position, tokenIndex
										if buffer[position] != rune('r') {
											goto l643
										}
										position++
										goto l642
									l643:
										position, tokenIndex = position642, tokenIndex642
										if buffer[position] != rune('R') {
											goto l633
										}
										position++
									}
								l642:
									{
										position644, tokenIndex644 := position, tokenIndex
										if buffer[position] != rune('n') {
											goto l645
										}
										position++
										goto l644
									l645:
										position, tokenIndex = position644, tokenIndex644
										if buffer[position] != rune('N') {
											goto l633
										}
										position++
									}
								l644:
									goto l632
								l633:
									position, tokenIndex = position632, tokenIndex632
									{
										position646, tokenIndex646 := position, tokenIndex
										if buffer[position] != rune('s') {
											goto l647
										}
										position++
										goto l646
									l647:
										position, tokenIndex = position646, tokenIndex646
										if buffer[position] != rune('S') {
											goto l57
										}
										position++
									}
								l646:
									{
										position648, tokenIndex648 := position, tokenIndex
										if buffer[position] != rune('t') {
											goto l649
										}
										position++
										goto l648
									l649:
										position, tokenIndex = position648, tokenIndex648
										if buffer[position] != rune('T') {
											goto l57
										}
										position++
									}
								l648:
									{
										position650, tokenIndex650 := position, tokenIndex
										if buffer[position] != rune('a') {
											goto l651
										}
										position++
										goto l650
									l651:
										position, tokenIndex = position650, tokenIndex650
										if buffer[position] != rune('A') {
											goto l57
										}
										position++
									}
								l650:
									{
										position652, tokenIndex652 := position, tokenIndex
										if buffer[position] != rune('t') {
											goto l653
										}
										position++
										goto l652
									l653:
										position, tokenIndex = position652, tokenIndex652
										if buffer[position] != rune('T') {
											goto l57
										}
										position++
									}
								l652:
									{
										position654, tokenIndex654 := position, tokenIndex
										if buffer[position] != rune('i') {
											goto l655
										}
										position++
										goto l654
									l655:
										position, tokenIndex = position654, tokenIndex654
										if buffer[position] != rune('I') {
											goto l57
										}
										position++
									}
								l654:
									{
										position656, tokenIndex656 := position, tokenIndex
										if buffer[position] != rune('c') {
											goto l657
										}
										position++
										goto l656
									l657:
										position, tokenIndex = position656, tokenIndex656
										if buffer[position] != rune('C') {
											goto l57
										}
										position++
									}
								l656:
								}
							l632:
								add(ruleLinkAttr, position623)
							}
							break
						case 'N', 'n':
							{
								position658 := position
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
										goto l57
									}
									position++
								}
							l659:
								{
									position661, tokenIndex661 := position, tokenIndex
									if buffer[position] != rune('o') {
										goto l662
									}
									position++
									goto l661
								l662:
									position, tokenIndex = position661, tokenIndex661
									if buffer[position] != rune('O') {
										goto l57
									}
									position++
								}
							l661:
								{
									position663, tokenIndex663 := position, tokenIndex
									if buffer[position] != rune('t') {
										goto l664
									}
									position++
									goto l663
								l664:
									position, tokenIndex = position663, tokenIndex663
									if buffer[position] != rune('T') {
										goto l57
									}
									position++
								}
							l663:
								{
									position665, tokenIndex665 := position, tokenIndex
									if buffer[position] != rune('e') {
										goto l666
									}
									position++
									goto l665
								l666:
									position, tokenIndex = position665, tokenIndex665
									if buffer[position] != rune('E') {
										goto l57
									}
									position++
								}
							l665:
								if !_rules[rulews]() {
									goto l57
								}
								if buffer[position] != rune(':') {
									goto l57
								}
								position++
								if !_rules[rulews]() {
									goto l57
								}
								{
									position667, tokenIndex667 := position, tokenIndex
									{
										position669, tokenIndex669 := position, tokenIndex
										if buffer[position] != rune('p') {
											goto l670
										}
										position++
										goto l669
									l670:
										position, tokenIndex = position669, tokenIndex669
										if buffer[position] != rune('P') {
											goto l668
										}
										position++
									}
								l669:
									{
										position671, tokenIndex671 := position, tokenIndex
										if buffer[position] != rune('t') {
											goto l672
										}
										position++
										goto l671
									l672:
										position, tokenIndex = position671, tokenIndex671
										if buffer[position] != rune('T') {
											goto l668
										}
										position++
									}
								l671:
									{
										position673, tokenIndex673 := position, tokenIndex
										if buffer[position] != rune('r') {
											goto l674
										}
										position++
										goto l673
									l674:
										position, tokenIndex = position673, tokenIndex673
										if buffer[position] != rune('R') {
											goto l668
										}
										position++
									}
								l673:
									{
										position675, tokenIndex675 := position, tokenIndex
										if buffer[position] != rune('m') {
											goto l676
										}
										position++
										goto l675
									l676:
										position, tokenIndex = position675, tokenIndex675
										if buffer[position] != rune('M') {
											goto l668
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
											goto l668
										}
										position++
									}
								l677:
									{
										position679, tokenIndex679 := position, tokenIndex
										if buffer[position] != rune('m') {
											goto l680
										}
										position++
										goto l679
									l680:
										position, tokenIndex = position679, tokenIndex679
										if buffer[position] != rune('M') {
											goto l668
										}
										position++
									}
								l679:
									goto l667
								l668:
									position, tokenIndex = position667, tokenIndex667
									{
										switch buffer[position] {
										case 'P', 'p':
											{
												position682, tokenIndex682 := position, tokenIndex
												if buffer[position] != rune('p') {
													goto l683
												}
												position++
												goto l682
											l683:
												position, tokenIndex = position682, tokenIndex682
												if buffer[position] != rune('P') {
													goto l57
												}
												position++
											}
										l682:
											{
												position684, tokenIndex684 := position, tokenIndex
												if buffer[position] != rune('s') {
													goto l685
												}
												position++
												goto l684
											l685:
												position, tokenIndex = position684, tokenIndex684
												if buffer[position] != rune('S') {
													goto l57
												}
												position++
											}
										l684:
											{
												position686, tokenIndex686 := position, tokenIndex
												if buffer[position] != rune('e') {
													goto l687
												}
												position++
												goto l686
											l687:
												position, tokenIndex = position686, tokenIndex686
												if buffer[position] != rune('E') {
													goto l57
												}
												position++
											}
										l686:
											{
												position688, tokenIndex688 := position, tokenIndex
												if buffer[position] != rune('u') {
													goto l689
												}
												position++
												goto l688
											l689:
												position, tokenIndex = position688, tokenIndex688
												if buffer[position] != rune('U') {
													goto l57
												}
												position++
											}
										l688:
											{
												position690, tokenIndex690 := position, tokenIndex
												if buffer[position] != rune('d') {
													goto l691
												}
												position++
												goto l690
											l691:
												position, tokenIndex = position690, tokenIndex690
												if buffer[position] != rune('D') {
													goto l57
												}
												position++
											}
										l690:
											{
												position692, tokenIndex692 := position, tokenIndex
												if buffer[position] != rune('o') {
													goto l693
												}
												position++
												goto l692
											l693:
												position, tokenIndex = position692, tokenIndex692
												if buffer[position] != rune('O') {
													goto l57
												}
												position++
											}
										l692:
											if buffer[position] != rune(' ') {
												goto l57
											}
											position++
											{
												position694, tokenIndex694 := position, tokenIndex
												if buffer[position] != rune('t') {
													goto l695
												}
												position++
												goto l694
											l695:
												position, tokenIndex = position694, tokenIndex694
												if buffer[position] != rune('T') {
													goto l57
												}
												position++
											}
										l694:
											{
												position696, tokenIndex696 := position, tokenIndex
												if buffer[position] != rune('m') {
													goto l697
												}
												position++
												goto l696
											l697:
												position, tokenIndex = position696, tokenIndex696
												if buffer[position] != rune('M') {
													goto l57
												}
												position++
											}
										l696:
											{
												position698, tokenIndex698 := position, tokenIndex
												if buffer[position] != rune('p') {
													goto l699
												}
												position++
												goto l698
											l699:
												position, tokenIndex = position698, tokenIndex698
												if buffer[position] != rune('P') {
													goto l57
												}
												position++
											}
										l698:
											{
												position700, tokenIndex700 := position, tokenIndex
												if buffer[position] != rune('l') {
													goto l701
												}
												position++
												goto l700
											l701:
												position, tokenIndex = position700, tokenIndex700
												if buffer[position] != rune('L') {
													goto l57
												}
												position++
											}
										l700:
											break
										case 'O', 'o':
											{
												position702, tokenIndex702 := position, tokenIndex
												if buffer[position] != rune('o') {
													goto l703
												}
												position++
												goto l702
											l703:
												position, tokenIndex = position702, tokenIndex702
												if buffer[position] != rune('O') {
													goto l57
												}
												position++
											}
										l702:
											{
												position704, tokenIndex704 := position, tokenIndex
												if buffer[position] != rune('p') {
													goto l705
												}
												position++
												goto l704
											l705:
												position, tokenIndex = position704, tokenIndex704
												if buffer[position] != rune('P') {
													goto l57
												}
												position++
											}
										l704:
											{
												position706, tokenIndex706 := position, tokenIndex
												if buffer[position] != rune('e') {
													goto l707
												}
												position++
												goto l706
											l707:
												position, tokenIndex = position706, tokenIndex706
												if buffer[position] != rune('E') {
													goto l57
												}
												position++
											}
										l706:
											{
												position708, tokenIndex708 := position, tokenIndex
												if buffer[position] != rune('r') {
													goto l709
												}
												position++
												goto l708
											l709:
												position, tokenIndex = position708, tokenIndex708
												if buffer[position] != rune('R') {
													goto l57
												}
												position++
											}
										l708:
											{
												position710, tokenIndex710 := position, tokenIndex
												if buffer[position] != rune('a') {
													goto l711
												}
												position++
												goto l710
											l711:
												position, tokenIndex = position710, tokenIndex710
												if buffer[position] != rune('A') {
													goto l57
												}
												position++
											}
										l710:
											{
												position712, tokenIndex712 := position, tokenIndex
												if buffer[position] != rune('t') {
													goto l713
												}
												position++
												goto l712
											l713:
												position, tokenIndex = position712, tokenIndex712
												if buffer[position] != rune('T') {
													goto l57
												}
												position++
											}
										l712:
											{
												position714, tokenIndex714 := position, tokenIndex
												if buffer[position] != rune('o') {
													goto l715
												}
												position++
												goto l714
											l715:
												position, tokenIndex = position714, tokenIndex714
												if buffer[position] != rune('O') {
													goto l57
												}
												position++
											}
										l714:
											{
												position716, tokenIndex716 := position, tokenIndex
												if buffer[position] != rune('r') {
													goto l717
												}
												position++
												goto l716
											l717:
												position, tokenIndex = position716, tokenIndex716
												if buffer[position] != rune('R') {
													goto l57
												}
												position++
											}
										l716:
											if buffer[position] != rune(' ') {
												goto l57
											}
											position++
											{
												position718, tokenIndex718 := position, tokenIndex
												{
													position720, tokenIndex720 := position, tokenIndex
													if buffer[position] != rune('g') {
														goto l721
													}
													position++
													goto l720
												l721:
													position, tokenIndex = position720, tokenIndex720
													if buffer[position] != rune('G') {
														goto l719
													}
													position++
												}
											l720:
												{
													position722, tokenIndex722 := position, tokenIndex
													if buffer[position] != rune('e') {
														goto l723
													}
													position++
													goto l722
												l723:
													position, tokenIndex = position722, tokenIndex722
													if buffer[position] != rune('E') {
														goto l719
													}
													position++
												}
											l722:
												goto l718
											l719:
												position, tokenIndex = position718, tokenIndex718
												{
													position725, tokenIndex725 := position, tokenIndex
													if buffer[position] != rune('l') {
														goto l726
													}
													position++
													goto l725
												l726:
													position, tokenIndex = position725, tokenIndex725
													if buffer[position] != rune('L') {
														goto l724
													}
													position++
												}
											l725:
												{
													position727, tokenIndex727 := position, tokenIndex
													if buffer[position] != rune('e') {
														goto l728
													}
													position++
													goto l727
												l728:
													position, tokenIndex = position727, tokenIndex727
													if buffer[position] != rune('E') {
														goto l724
													}
													position++
												}
											l727:
												goto l718
											l724:
												position, tokenIndex = position718, tokenIndex718
												{
													position730, tokenIndex730 := position, tokenIndex
													if buffer[position] != rune('g') {
														goto l731
													}
													position++
													goto l730
												l731:
													position, tokenIndex = position730, tokenIndex730
													if buffer[position] != rune('G') {
														goto l729
													}
													position++
												}
											l730:
												{
													position732, tokenIndex732 := position, tokenIndex
													if buffer[position] != rune('t') {
														goto l733
													}
													position++
													goto l732
												l733:
													position, tokenIndex = position732, tokenIndex732
													if buffer[position] != rune('T') {
														goto l729
													}
													position++
												}
											l732:
												goto l718
											l729:
												position, tokenIndex = position718, tokenIndex718
												{
													position734, tokenIndex734 := position, tokenIndex
													if buffer[position] != rune('l') {
														goto l735
													}
													position++
													goto l734
												l735:
													position, tokenIndex = position734, tokenIndex734
													if buffer[position] != rune('L') {
														goto l57
													}
													position++
												}
											l734:
												{
													position736, tokenIndex736 := position, tokenIndex
													if buffer[position] != rune('t') {
														goto l737
													}
													position++
													goto l736
												l737:
													position, tokenIndex = position736, tokenIndex736
													if buffer[position] != rune('T') {
														goto l57
													}
													position++
												}
											l736:
											}
										l718:
											break
										case 'M', 'm':
											{
												position738, tokenIndex738 := position, tokenIndex
												if buffer[position] != rune('m') {
													goto l739
												}
												position++
												goto l738
											l739:
												position, tokenIndex = position738, tokenIndex738
												if buffer[position] != rune('M') {
													goto l57
												}
												position++
											}
										l738:
											{
												position740, tokenIndex740 := position, tokenIndex
												if buffer[position] != rune('e') {
													goto l741
												}
												position++
												goto l740
											l741:
												position, tokenIndex = position740, tokenIndex740
												if buffer[position] != rune('E') {
													goto l57
												}
												position++
											}
										l740:
											{
												position742, tokenIndex742 := position, tokenIndex
												if buffer[position] != rune('m') {
													goto l743
												}
												position++
												goto l742
											l743:
												position, tokenIndex = position742, tokenIndex742
												if buffer[position] != rune('M') {
													goto l57
												}
												position++
											}
										l742:
											{
												position744, tokenIndex744 := position, tokenIndex
												if buffer[position] != rune('b') {
													goto l745
												}
												position++
												goto l744
											l745:
												position, tokenIndex = position744, tokenIndex744
												if buffer[position] != rune('B') {
													goto l57
												}
												position++
											}
										l744:
											{
												position746, tokenIndex746 := position, tokenIndex
												if buffer[position] != rune('e') {
													goto l747
												}
												position++
												goto l746
											l747:
												position, tokenIndex = position746, tokenIndex746
												if buffer[position] != rune('E') {
													goto l57
												}
												position++
											}
										l746:
											{
												position748, tokenIndex748 := position, tokenIndex
												if buffer[position] != rune('r') {
													goto l749
												}
												position++
												goto l748
											l749:
												position, tokenIndex = position748, tokenIndex748
												if buffer[position] != rune('R') {
													goto l57
												}
												position++
											}
										l748:
											break
										default:
											{
												position750, tokenIndex750 := position, tokenIndex
												if buffer[position] != rune('a') {
													goto l751
												}
												position++
												goto l750
											l751:
												position, tokenIndex = position750, tokenIndex750
												if buffer[position] != rune('A') {
													goto l57
												}
												position++
											}
										l750:
											{
												position752, tokenIndex752 := position, tokenIndex
												if buffer[position] != rune('r') {
													goto l753
												}
												position++
												goto l752
											l753:
												position, tokenIndex = position752, tokenIndex752
												if buffer[position] != rune('R') {
													goto l57
												}
												position++
											}
										l752:
											{
												position754, tokenIndex754 := position, tokenIndex
												if buffer[position] != rune('t') {
													goto l755
												}
												position++
												goto l754
											l755:
												position, tokenIndex = position754, tokenIndex754
												if buffer[position] != rune('T') {
													goto l57
												}
												position++
											}
										l754:
											{
												position756, tokenIndex756 := position, tokenIndex
												if buffer[position] != rune('i') {
													goto l757
												}
												position++
												goto l756
											l757:
												position, tokenIndex = position756, tokenIndex756
												if buffer[position] != rune('I') {
													goto l57
												}
												position++
											}
										l756:
											{
												position758, tokenIndex758 := position, tokenIndex
												if buffer[position] != rune('f') {
													goto l759
												}
												position++
												goto l758
											l759:
												position, tokenIndex = position758, tokenIndex758
												if buffer[position] != rune('F') {
													goto l57
												}
												position++
											}
										l758:
											{
												position760, tokenIndex760 := position, tokenIndex
												if buffer[position] != rune('i') {
													goto l761
												}
												position++
												goto l760
											l761:
												position, tokenIndex = position760, tokenIndex760
												if buffer[position] != rune('I') {
													goto l57
												}
												position++
											}
										l760:
											{
												position762, tokenIndex762 := position, tokenIndex
												if buffer[position] != rune('c') {
													goto l763
												}
												position++
												goto l762
											l763:
												position, tokenIndex = position762, tokenIndex762
												if buffer[position] != rune('C') {
													goto l57
												}
												position++
											}
										l762:
											{
												position764, tokenIndex764 := position, tokenIndex
												if buffer[position] != rune('i') {
													goto l765
												}
												position++
												goto l764
											l765:
												position, tokenIndex = position764, tokenIndex764
												if buffer[position] != rune('I') {
													goto l57
												}
												position++
											}
										l764:
											{
												position766, tokenIndex766 := position, tokenIndex
												if buffer[position] != rune('a') {
													goto l767
												}
												position++
												goto l766
											l767:
												position, tokenIndex = position766, tokenIndex766
												if buffer[position] != rune('A') {
													goto l57
												}
												position++
											}
										l766:
											{
												position768, tokenIndex768 := position, tokenIndex
												if buffer[position] != rune('l') {
													goto l769
												}
												position++
												goto l768
											l769:
												position, tokenIndex = position768, tokenIndex768
												if buffer[position] != rune('L') {
													goto l57
												}
												position++
											}
										l768:
											break
										}
									}

								}
							l667:
								add(ruleNoteAttr, position658)
							}
							break
						case 'B', 'b':
							{
								position770 := position
								{
									position771, tokenIndex771 := position, tokenIndex
									if buffer[position] != rune('b') {
										goto l772
									}
									position++
									goto l771
								l772:
									position, tokenIndex = position771, tokenIndex771
									if buffer[position] != rune('B') {
										goto l57
									}
									position++
								}
							l771:
								{
									position773, tokenIndex773 := position, tokenIndex
									if buffer[position] != rune('o') {
										goto l774
									}
									position++
									goto l773
								l774:
									position, tokenIndex = position773, tokenIndex773
									if buffer[position] != rune('O') {
										goto l57
									}
									position++
								}
							l773:
								{
									position775, tokenIndex775 := position, tokenIndex
									if buffer[position] != rune('d') {
										goto l776
									}
									position++
									goto l775
								l776:
									position, tokenIndex = position775, tokenIndex775
									if buffer[position] != rune('D') {
										goto l57
									}
									position++
								}
							l775:
								{
									position777, tokenIndex777 := position, tokenIndex
									if buffer[position] != rune('y') {
										goto l778
									}
									position++
									goto l777
								l778:
									position, tokenIndex = position777, tokenIndex777
									if buffer[position] != rune('Y') {
										goto l57
									}
									position++
								}
							l777:
								if !_rules[rulews]() {
									goto l57
								}
								if buffer[position] != rune(':') {
									goto l57
								}
								position++
								if !_rules[rulews]() {
									goto l57
								}
								{
									position779, tokenIndex779 := position, tokenIndex
									{
										position781, tokenIndex781 := position, tokenIndex
										if buffer[position] != rune('u') {
											goto l782
										}
										position++
										goto l781
									l782:
										position, tokenIndex = position781, tokenIndex781
										if buffer[position] != rune('U') {
											goto l780
										}
										position++
									}
								l781:
									{
										position783, tokenIndex783 := position, tokenIndex
										if buffer[position] != rune('n') {
											goto l784
										}
										position++
										goto l783
									l784:
										position, tokenIndex = position783, tokenIndex783
										if buffer[position] != rune('N') {
											goto l780
										}
										position++
									}
								l783:
									{
										position785, tokenIndex785 := position, tokenIndex
										if buffer[position] != rune('d') {
											goto l786
										}
										position++
										goto l785
									l786:
										position, tokenIndex = position785, tokenIndex785
										if buffer[position] != rune('D') {
											goto l780
										}
										position++
									}
								l785:
									{
										position787, tokenIndex787 := position, tokenIndex
										if buffer[position] != rune('e') {
											goto l788
										}
										position++
										goto l787
									l788:
										position, tokenIndex = position787, tokenIndex787
										if buffer[position] != rune('E') {
											goto l780
										}
										position++
									}
								l787:
									{
										position789, tokenIndex789 := position, tokenIndex
										if buffer[position] != rune('f') {
											goto l790
										}
										position++
										goto l789
									l790:
										position, tokenIndex = position789, tokenIndex789
										if buffer[position] != rune('F') {
											goto l780
										}
										position++
									}
								l789:
									{
										position791, tokenIndex791 := position, tokenIndex
										if buffer[position] != rune('i') {
											goto l792
										}
										position++
										goto l791
									l792:
										position, tokenIndex = position791, tokenIndex791
										if buffer[position] != rune('I') {
											goto l780
										}
										position++
									}
								l791:
									{
										position793, tokenIndex793 := position, tokenIndex
										if buffer[position] != rune('n') {
											goto l794
										}
										position++
										goto l793
									l794:
										position, tokenIndex = position793, tokenIndex793
										if buffer[position] != rune('N') {
											goto l780
										}
										position++
									}
								l793:
									{
										position795, tokenIndex795 := position, tokenIndex
										if buffer[position] != rune('e') {
											goto l796
										}
										position++
										goto l795
									l796:
										position, tokenIndex = position795, tokenIndex795
										if buffer[position] != rune('E') {
											goto l780
										}
										position++
									}
								l795:
									{
										position797, tokenIndex797 := position, tokenIndex
										if buffer[position] != rune('d') {
											goto l798
										}
										position++
										goto l797
									l798:
										position, tokenIndex = position797, tokenIndex797
										if buffer[position] != rune('D') {
											goto l780
										}
										position++
									}
								l797:
									goto l779
								l780:
									position, tokenIndex = position779, tokenIndex779
									if !_rules[ruleNode]() {
										goto l57
									}
								}
							l779:
								add(ruleBodyAttr, position770)
							}
							break
						case 'T', 't':
							{
								position799 := position
								{
									position800, tokenIndex800 := position, tokenIndex
									if buffer[position] != rune('t') {
										goto l801
									}
									position++
									goto l800
								l801:
									position, tokenIndex = position800, tokenIndex800
									if buffer[position] != rune('T') {
										goto l57
									}
									position++
								}
							l800:
								{
									position802, tokenIndex802 := position, tokenIndex
									if buffer[position] != rune('a') {
										goto l803
									}
									position++
									goto l802
								l803:
									position, tokenIndex = position802, tokenIndex802
									if buffer[position] != rune('A') {
										goto l57
									}
									position++
								}
							l802:
								{
									position804, tokenIndex804 := position, tokenIndex
									if buffer[position] != rune('g') {
										goto l805
									}
									position++
									goto l804
								l805:
									position, tokenIndex = position804, tokenIndex804
									if buffer[position] != rune('G') {
										goto l57
									}
									position++
								}
							l804:
								if !_rules[rulews]() {
									goto l57
								}
								if buffer[position] != rune(':') {
									goto l57
								}
								position++
								if !_rules[rulews]() {
									goto l57
								}
								{
									position806, tokenIndex806 := position, tokenIndex
									{
										position808, tokenIndex808 := position, tokenIndex
										if buffer[position] != rune('s') {
											goto l809
										}
										position++
										goto l808
									l809:
										position, tokenIndex = position808, tokenIndex808
										if buffer[position] != rune('S') {
											goto l807
										}
										position++
									}
								l808:
									{
										position810, tokenIndex810 := position, tokenIndex
										if buffer[position] != rune('t') {
											goto l811
										}
										position++
										goto l810
									l811:
										position, tokenIndex = position810, tokenIndex810
										if buffer[position] != rune('T') {
											goto l807
										}
										position++
									}
								l810:
									{
										position812, tokenIndex812 := position, tokenIndex
										if buffer[position] != rune('r') {
											goto l813
										}
										position++
										goto l812
									l813:
										position, tokenIndex = position812, tokenIndex812
										if buffer[position] != rune('R') {
											goto l807
										}
										position++
									}
								l812:
									{
										position814, tokenIndex814 := position, tokenIndex
										if buffer[position] != rune('u') {
											goto l815
										}
										position++
										goto l814
									l815:
										position, tokenIndex = position814, tokenIndex814
										if buffer[position] != rune('U') {
											goto l807
										}
										position++
									}
								l814:
									{
										position816, tokenIndex816 := position, tokenIndex
										if buffer[position] != rune('c') {
											goto l817
										}
										position++
										goto l816
									l817:
										position, tokenIndex = position816, tokenIndex816
										if buffer[position] != rune('C') {
											goto l807
										}
										position++
									}
								l816:
									{
										position818, tokenIndex818 := position, tokenIndex
										if buffer[position] != rune('t') {
											goto l819
										}
										position++
										goto l818
									l819:
										position, tokenIndex = position818, tokenIndex818
										if buffer[position] != rune('T') {
											goto l807
										}
										position++
									}
								l818:
									goto l806
								l807:
									position, tokenIndex = position806, tokenIndex806
									{
										position820, tokenIndex820 := position, tokenIndex
										if buffer[position] != rune('u') {
											goto l821
										}
										position++
										goto l820
									l821:
										position, tokenIndex = position820, tokenIndex820
										if buffer[position] != rune('U') {
											goto l57
										}
										position++
									}
								l820:
									{
										position822, tokenIndex822 := position, tokenIndex
										if buffer[position] != rune('n') {
											goto l823
										}
										position++
										goto l822
									l823:
										position, tokenIndex = position822, tokenIndex822
										if buffer[position] != rune('N') {
											goto l57
										}
										position++
									}
								l822:
									{
										position824, tokenIndex824 := position, tokenIndex
										if buffer[position] != rune('i') {
											goto l825
										}
										position++
										goto l824
									l825:
										position, tokenIndex = position824, tokenIndex824
										if buffer[position] != rune('I') {
											goto l57
										}
										position++
									}
								l824:
									{
										position826, tokenIndex826 := position, tokenIndex
										if buffer[position] != rune('o') {
											goto l827
										}
										position++
										goto l826
									l827:
										position, tokenIndex = position826, tokenIndex826
										if buffer[position] != rune('O') {
											goto l57
										}
										position++
									}
								l826:
									{
										position828, tokenIndex828 := position, tokenIndex
										if buffer[position] != rune('n') {
											goto l829
										}
										position++
										goto l828
									l829:
										position, tokenIndex = position828, tokenIndex828
										if buffer[position] != rune('N') {
											goto l57
										}
										position++
									}
								l828:
								}
							l806:
								add(ruleTagAttr, position799)
							}
							break
						case 'I', 'i':
							{
								position830 := position
								{
									position831, tokenIndex831 := position, tokenIndex
									if buffer[position] != rune('i') {
										goto l832
									}
									position++
									goto l831
								l832:
									position, tokenIndex = position831, tokenIndex831
									if buffer[position] != rune('I') {
										goto l57
									}
									position++
								}
							l831:
								{
									position833, tokenIndex833 := position, tokenIndex
									if buffer[position] != rune('n') {
										goto l834
									}
									position++
									goto l833
								l834:
									position, tokenIndex = position833, tokenIndex833
									if buffer[position] != rune('N') {
										goto l57
									}
									position++
								}
							l833:
								{
									position835, tokenIndex835 := position, tokenIndex
									if buffer[position] != rune('t') {
										goto l836
									}
									position++
									goto l835
								l836:
									position, tokenIndex = position835, tokenIndex835
									if buffer[position] != rune('T') {
										goto l57
									}
									position++
								}
							l835:
								if !_rules[rulews]() {
									goto l57
								}
								if buffer[position] != rune(':') {
									goto l57
								}
								position++
								if !_rules[rulews]() {
									goto l57
								}
								{
									position837, tokenIndex837 := position, tokenIndex
									{
										position839, tokenIndex839 := position, tokenIndex
										if buffer[position] != rune('-') {
											goto l840
										}
										position++
										goto l839
									l840:
										position, tokenIndex = position839, tokenIndex839
										if buffer[position] != rune('0') {
											goto l837
										}
										position++
										{
											position841, tokenIndex841 := position, tokenIndex
											if buffer[position] != rune('x') {
												goto l842
											}
											position++
											goto l841
										l842:
											position, tokenIndex = position841, tokenIndex841
											if buffer[position] != rune('X') {
												goto l837
											}
											position++
										}
									l841:
									}
								l839:
									goto l838
								l837:
									position, tokenIndex = position837, tokenIndex837
								}
							l838:
								{
									position845, tokenIndex845 := position, tokenIndex
									if !_rules[ruleInteger]() {
										goto l846
									}
									goto l845
								l846:
									position, tokenIndex = position845, tokenIndex845
									if !_rules[ruleHex]() {
										goto l57
									}
								}
							l845:
							l843:
								{
									position844, tokenIndex844 := position, tokenIndex
									{
										position847, tokenIndex847 := position, tokenIndex
										if !_rules[ruleInteger]() {
											goto l848
										}
										goto l847
									l848:
										position, tokenIndex = position847, tokenIndex847
										if !_rules[ruleHex]() {
											goto l844
										}
									}
								l847:
									goto l843
								l844:
									position, tokenIndex = position844, tokenIndex844
								}
								if !_rules[rulews]() {
									goto l57
								}
								add(ruleSignedIntAttr, position830)
							}
							break
						default:
							{
								position849 := position
								{
									position850, tokenIndex850 := position, tokenIndex
									if buffer[position] != rune('s') {
										goto l851
									}
									position++
									goto l850
								l851:
									position, tokenIndex = position850, tokenIndex850
									if buffer[position] != rune('S') {
										goto l57
									}
									position++
								}
							l850:
								{
									position852, tokenIndex852 := position, tokenIndex
									if buffer[position] != rune('i') {
										goto l853
									}
									position++
									goto l852
								l853:
									position, tokenIndex = position852, tokenIndex852
									if buffer[position] != rune('I') {
										goto l57
									}
									position++
								}
							l852:
								{
									position854, tokenIndex854 := position, tokenIndex
									if buffer[position] != rune('g') {
										goto l855
									}
									position++
									goto l854
								l855:
									position, tokenIndex = position854, tokenIndex854
									if buffer[position] != rune('G') {
										goto l57
									}
									position++
								}
							l854:
								{
									position856, tokenIndex856 := position, tokenIndex
									if buffer[position] != rune('n') {
										goto l857
									}
									position++
									goto l856
								l857:
									position, tokenIndex = position856, tokenIndex856
									if buffer[position] != rune('N') {
										goto l57
									}
									position++
								}
							l856:
								if !_rules[rulews]() {
									goto l57
								}
								if buffer[position] != rune(':') {
									goto l57
								}
								position++
								if !_rules[rulews]() {
									goto l57
								}
								{
									position858, tokenIndex858 := position, tokenIndex
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
											goto l859
										}
										position++
									}
								l860:
									{
										position862, tokenIndex862 := position, tokenIndex
										if buffer[position] != rune('i') {
											goto l863
										}
										position++
										goto l862
									l863:
										position, tokenIndex = position862, tokenIndex862
										if buffer[position] != rune('I') {
											goto l859
										}
										position++
									}
								l862:
									{
										position864, tokenIndex864 := position, tokenIndex
										if buffer[position] != rune('g') {
											goto l865
										}
										position++
										goto l864
									l865:
										position, tokenIndex = position864, tokenIndex864
										if buffer[position] != rune('G') {
											goto l859
										}
										position++
									}
								l864:
									{
										position866, tokenIndex866 := position, tokenIndex
										if buffer[position] != rune('n') {
											goto l867
										}
										position++
										goto l866
									l867:
										position, tokenIndex = position866, tokenIndex866
										if buffer[position] != rune('N') {
											goto l859
										}
										position++
									}
								l866:
									{
										position868, tokenIndex868 := position, tokenIndex
										if buffer[position] != rune('e') {
											goto l869
										}
										position++
										goto l868
									l869:
										position, tokenIndex = position868, tokenIndex868
										if buffer[position] != rune('E') {
											goto l859
										}
										position++
									}
								l868:
									{
										position870, tokenIndex870 := position, tokenIndex
										if buffer[position] != rune('d') {
											goto l871
										}
										position++
										goto l870
									l871:
										position, tokenIndex = position870, tokenIndex870
										if buffer[position] != rune('D') {
											goto l859
										}
										position++
									}
								l870:
									goto l858
								l859:
									position, tokenIndex = position858, tokenIndex858
									{
										position872, tokenIndex872 := position, tokenIndex
										if buffer[position] != rune('u') {
											goto l873
										}
										position++
										goto l872
									l873:
										position, tokenIndex = position872, tokenIndex872
										if buffer[position] != rune('U') {
											goto l57
										}
										position++
									}
								l872:
									{
										position874, tokenIndex874 := position, tokenIndex
										if buffer[position] != rune('n') {
											goto l875
										}
										position++
										goto l874
									l875:
										position, tokenIndex = position874, tokenIndex874
										if buffer[position] != rune('N') {
											goto l57
										}
										position++
									}
								l874:
									{
										position876, tokenIndex876 := position, tokenIndex
										if buffer[position] != rune('s') {
											goto l877
										}
										position++
										goto l876
									l877:
										position, tokenIndex = position876, tokenIndex876
										if buffer[position] != rune('S') {
											goto l57
										}
										position++
									}
								l876:
									{
										position878, tokenIndex878 := position, tokenIndex
										if buffer[position] != rune('i') {
											goto l879
										}
										position++
										goto l878
									l879:
										position, tokenIndex = position878, tokenIndex878
										if buffer[position] != rune('I') {
											goto l57
										}
										position++
									}
								l878:
									{
										position880, tokenIndex880 := position, tokenIndex
										if buffer[position] != rune('g') {
											goto l881
										}
										position++
										goto l880
									l881:
										position, tokenIndex = position880, tokenIndex880
										if buffer[position] != rune('G') {
											goto l57
										}
										position++
									}
								l880:
									{
										position882, tokenIndex882 := position, tokenIndex
										if buffer[position] != rune('n') {
											goto l883
										}
										position++
										goto l882
									l883:
										position, tokenIndex = position882, tokenIndex882
										if buffer[position] != rune('N') {
											goto l57
										}
										position++
									}
								l882:
									{
										position884, tokenIndex884 := position, tokenIndex
										if buffer[position] != rune('e') {
											goto l885
										}
										position++
										goto l884
									l885:
										position, tokenIndex = position884, tokenIndex884
										if buffer[position] != rune('E') {
											goto l57
										}
										position++
									}
								l884:
									{
										position886, tokenIndex886 := position, tokenIndex
										if buffer[position] != rune('d') {
											goto l887
										}
										position++
										goto l886
									l887:
										position, tokenIndex = position886, tokenIndex886
										if buffer[position] != rune('D') {
											goto l57
										}
										position++
									}
								l886:
								}
							l858:
								add(ruleSignAttr, position849)
							}
							break
						}
					}

				}
			l59:
				add(ruleOneAttr, position58)
			}
			return true
		l57:
			position, tokenIndex = position57, tokenIndex57
			return false
		},
		/* 24 Attrs <- <(ws OneAttr)*> */
		nil,
		/* 25 Attr <- <(OneAttr ws Attrs ws)?> */
		nil,
		/* 26 Statement <- <(Node ws NodeType ws Attr ws)> */
		nil,
		/* 27 Node <- <('@' NonZeroDecimalDigit DecimalDigit*)> */
		func() bool {
			position891, tokenIndex891 := position, tokenIndex
			{
				position892 := position
				if buffer[position] != rune('@') {
					goto l891
				}
				position++
				if !_rules[ruleNonZeroDecimalDigit]() {
					goto l891
				}
			l893:
				{
					position894, tokenIndex894 := position, tokenIndex
					if !_rules[ruleDecimalDigit]() {
						goto l894
					}
					goto l893
				l894:
					position, tokenIndex = position894, tokenIndex894
				}
				add(ruleNode, position892)
			}
			return true
		l891:
			position, tokenIndex = position891, tokenIndex891
			return false
		},
		/* 28 Integer <- <('0' / (NonZeroDecimalDigit DecimalDigit*))> */
		func() bool {
			position895, tokenIndex895 := position, tokenIndex
			{
				position896 := position
				{
					position897, tokenIndex897 := position, tokenIndex
					if buffer[position] != rune('0') {
						goto l898
					}
					position++
					goto l897
				l898:
					position, tokenIndex = position897, tokenIndex897
					if !_rules[ruleNonZeroDecimalDigit]() {
						goto l895
					}
				l899:
					{
						position900, tokenIndex900 := position, tokenIndex
						if !_rules[ruleDecimalDigit]() {
							goto l900
						}
						goto l899
					l900:
						position, tokenIndex = position900, tokenIndex900
					}
				}
			l897:
				add(ruleInteger, position896)
			}
			return true
		l895:
			position, tokenIndex = position895, tokenIndex895
			return false
		},
		/* 29 NodeType <- <([a-z] / '_')+> */
		nil,
		/* 30 DecimalDigit <- <[0-9]> */
		func() bool {
			position902, tokenIndex902 := position, tokenIndex
			{
				position903 := position
				if c := buffer[position]; c < rune('0') || c > rune('9') {
					goto l902
				}
				position++
				add(ruleDecimalDigit, position903)
			}
			return true
		l902:
			position, tokenIndex = position902, tokenIndex902
			return false
		},
		/* 31 NonZeroDecimalDigit <- <[1-9]> */
		func() bool {
			position904, tokenIndex904 := position, tokenIndex
			{
				position905 := position
				if c := buffer[position]; c < rune('1') || c > rune('9') {
					goto l904
				}
				position++
				add(ruleNonZeroDecimalDigit, position905)
			}
			return true
		l904:
			position, tokenIndex = position904, tokenIndex904
			return false
		},
		/* 32 ws <- <((&('\n') '\n') | (&('\r') '\r') | (&('\t') '\t') | (&(' ') ' '))*> */
		func() bool {
			{
				position907 := position
			l908:
				{
					position909, tokenIndex909 := position, tokenIndex
					{
						switch buffer[position] {
						case '\n':
							if buffer[position] != rune('\n') {
								goto l909
							}
							position++
							break
						case '\r':
							if buffer[position] != rune('\r') {
								goto l909
							}
							position++
							break
						case '\t':
							if buffer[position] != rune('\t') {
								goto l909
							}
							position++
							break
						default:
							if buffer[position] != rune(' ') {
								goto l909
							}
							position++
							break
						}
					}

					goto l908
				l909:
					position, tokenIndex = position909, tokenIndex909
				}
				add(rulews, position907)
			}
			return true
		},
		/* 33 EOF <- <!.> */
		nil,
	}
	p.rules = _rules
}
