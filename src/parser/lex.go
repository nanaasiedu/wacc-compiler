package parser

import (
	. "ast"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

const eof = -1

// stateFn represents the state of the scanner as a function that returns the next state.
type stateFn func(*Lexer) stateFn

// Token is a token
type Token struct {
	Typ    int    // The type of this item.
	Lexeme string // The value of this item.
	Pos    int    // The starting position, in bytes, of this item in the input string.
}

func (token Token) String() string {
	switch token.Typ {
	case eof:
		return "EOF"
	case ERROR:
		return token.Lexeme
	}
	if _, ok := TokenKeywordStrings[token.Lexeme]; ok {
		return fmt.Sprintf("%q", token.Lexeme)
	}
	if _, ok := TokenStrings[token.Lexeme]; ok {
		return fmt.Sprintf("%q", token.Lexeme)
	}
	return fmt.Sprint("Const: ", token.Lexeme)
}

// Lexer is  a struct
type Lexer struct {
	name       string
	input      string
	state      stateFn
	start      int
	pos        int
	width      int
	lastPos    int // position of most recent item returned by nextItem
	Items      chan Token
	prog       *Program // The parsed program
	lastItem   Token    // The last item emitted
	parseError bool
}

var escapeChars = []rune{'0', 'b', 't', 'n', 'f', 'r', '\\', '\'', '"'}

// TokenKeywordStrings key is a map of keywords: string keyword to integer type.
var TokenKeywordStrings = map[string]int{
	"class":   CLASS,
	"open":    OPEN,
	"close":   CLOSE,
	"begin":   BEGIN,
	"end":     END,
	"is":      IS,
	"skip":    SKIP,
	"read":    READ,
	"free":    FREE,
	"return":  RETURN,
	"exit":    EXIT,
	"println": PRINTLN,
	"print":   PRINT,
	"if":      IF,
	"then":    THEN,
	"else":    ELSE,
	"fi":      FI,
	"while":   WHILE,
	"done":    DONE,
	"do":      DO,
	"newpair": NEWPAIR,
	"call":    CALL,
	"fst":     FST,
	"snd":     SND,
	"null":    NULL,
	"int":     INT,
	"bool":    BOOL,
	"char":    CHAR,
	"string":  STRING,
	"pair":    PAIR,
	"len":     LEN,
	"ord":     ORD,
	"chr":     CHR,
	"true":    TRUE,
	"false":   FALSE,
	"for":     FOR,
	"this":    THIS,
	"new":     NEW,
}

// TokenStrings map
var TokenStrings = map[string]int{
	".":  DOT,
	",":  COMMA,
	";":  SEMICOLON,
	"%":  MOD,
	"*":  MUL,
	"/":  DIV,
	"+":  PLUS,
	"-":  SUB,
	">=": GTE,
	"<=": LTE,
	">":  GT,
	"<":  LT,
	"==": EQ,
	"!=": NEQ,
	"!":  NOT,
	"&&": AND,
	"||": OR,
	"[":  OPENSQUARE,
	"(":  OPENROUND,
	"]":  CLOSESQUARE,
	")":  CLOSEROUND,
	"=":  ASSIGNMENT,
}

// TokenLocation returns line number and column number of token in .wacc file
func (l *Lexer) TokenLocation(t Token) (line int, col int) {
	line = 1 + strings.Count(l.input[:t.Pos], "\n")
	col = t.Pos - strings.LastIndex(l.input[:t.Pos], "\n")
	return
}

//Returns line number and column number of current lexing item in .wacc file
func (l *Lexer) currLocation() (line int, col int) {
	line = 1 + strings.Count(l.input[:l.pos], "\n")
	col = l.pos - strings.LastIndex(l.input[:l.pos], "\n")
	return
}

// run runs the state machine for the lexer.
func (l *Lexer) run() {
	for state := lexText; state != nil; {
		state = state(l)
	}
	close(l.Items)
}

// NextItem returns the next item from the input.
// Called by the parser, not in the lexing goroutine.
func (l *Lexer) NextItem() Token {
	item := <-l.Items
	l.lastPos = item.Pos
	//l.lastItem = item
	return item
}

// lexText scans until an opening action "BEGIN"
func lexText(l *Lexer) stateFn {
	for {
		_ = "breakpoint"
		if strings.HasPrefix(l.input[l.pos:], "begin") {
			/*	if l.pos > l.start {
				l.emit(PLAINTEXT)
			}     */
			l.ignore()
			return lexInsideProgram
		}
		if l.next() == eof {
			break
		}
	}
	// Correctly reached eof
	return nil
}

// Lex creates a new scanner for the input string.
func newLex(name string, input string) *Lexer {
	l := &Lexer{
		name:  name,
		input: input,
		Items: make(chan Token),
	}
	go l.run()
	return l
}

// lexInsideAction scans the elements inside action delimiters.
func lexInsideProgram(l *Lexer) stateFn {
	var keysForKeywords sort.StringSlice
	for key := range TokenKeywordStrings {
		keysForKeywords = append(keysForKeywords, key)
	}
	keysForKeywords.Sort()
	sort.Sort(sort.Reverse(keysForKeywords))
	for _, str := range keysForKeywords {
		if strings.HasPrefix(l.input[l.pos:], str) {
			l.width = len(str)
			l.pos += l.width
			if isAlphaNumeric(l.peek()) || l.peek() == '_' {
				return lexIdentifier
			}
			l.emit(TokenKeywordStrings[str])
			return lexInsideProgram
		}
	}
	var keysForTokens sort.StringSlice
	for key := range TokenStrings {
		keysForTokens = append(keysForTokens, key)
	}
	keysForTokens.Sort()
	sort.Sort(sort.Reverse(keysForTokens))
	for _, str := range keysForTokens {
		if strings.HasPrefix(l.input[l.pos:], str) {
			if (l.input[l.pos] == '+' || l.input[l.pos] == '-') && '0' <= l.input[l.pos+1] && l.input[l.pos+1] <= '9' {
				switch l.lastItem.Typ {
				case ASSIGNMENT, OPENROUND, OPENSQUARE, COMMA, SEMICOLON:
					return lexNumber
				}
			}
			l.width = len(str)
			l.pos += l.width
			l.emit(TokenStrings[str])
			return lexInsideProgram
		}
	}
	switch r := l.next(); {
	case r == '"':
		return lexString
	case r == '\'':
		return lexChar
	case r == '\b', r == '\t', r == '\n', r == '\f', r == '\r':
		l.ignore()
		return lexInsideProgram
	case r == '#':
		for l.next() != '\n' {
			l.ignore()
		}
		l.backup()
		return lexInsideProgram
	case unicode.IsSpace(r):
		l.ignore()
		return lexInsideProgram
	case '0' <= r && r <= '9':
		l.backup()
		return lexNumber
	case isAlphaNumeric(r) || r == '_':
		l.backup()
		return lexIdentifier
	case r == eof:
		return nil
	}
	return l.errorf("Item Not in WACC lanuage: %s", l.input[l.start:l.start+l.width])
}

// next returns the next rune in the input.
func (l *Lexer) next() (char rune) {
	if l.pos >= len(l.input) {
		l.width = 0
		return eof
	}
	char, l.width = utf8.DecodeRuneInString(l.input[l.pos:])
	l.pos += l.width
	return char
}

// ignore skips over the pending input before this point.
func (l *Lexer) ignore() {
	l.start = l.pos
}

// backup steps back one rune. Can only be called once per call of next.
func (l *Lexer) backup() {
	l.pos -= l.width
}

// peek returns but does not consume the next rune in the input.
func (l *Lexer) peek() rune {
	char := l.next()
	l.backup()
	return char
}

// accept consumes the next rune if it's from the valid set.
func (l *Lexer) accept(vaild string) bool {
	if strings.IndexRune(vaild, l.next()) >= 0 {
		return true
	}
	l.backup()
	return false
}

// acceptRun consumes a run of runes from the valid set.
func (l *Lexer) acceptRun(valid string) {
	for strings.IndexRune(valid, l.next()) >= 0 {
	}
	l.backup()
}

// emit passes an item back to the client.
func (l *Lexer) emit(t int) {
	item := Token{Typ: t, Lexeme: l.input[l.start:l.pos], Pos: l.start}
	l.Items <- item
	l.lastItem = item
	l.start = l.pos
}

func (l *Lexer) errorf(format string, args ...interface{}) stateFn {
	line, col := l.currLocation()
	fmt.Printf("At %d:%d :: ", line, col)
	fmt.Printf(format, args)
	l.Items <- Token{Typ: ERROR, Lexeme: fmt.Sprintf(format, args...), Pos: l.start}
	return nil
}

// isAlphaNumeric reports whether r is an alphabetic, digit, or underscore.
func isAlphaNumeric(r rune) bool {
	return r == '_' || unicode.IsLetter(r) || unicode.IsDigit(r)
}

func lexIdentifier(l *Lexer) stateFn {
	if !(unicode.IsLetter(l.peek()) || l.peek() == '_') {
		l.next()
		return l.errorf("bad identifier syntax: %q", l.input[l.start:l.pos])
	}
	for isAlphaNumeric(l.peek()) || l.peek() == '_' {
		l.next()
	}
	l.emit(IDENTIFIER)
	return lexInsideProgram
}

func lexChar(l *Lexer) stateFn {
Loop:
	for {
		switch l.next() {
		case '"':
			fmt.Println(string(l.input[l.pos-2]))
			if l.input[l.pos-2] != '\\' {
				return l.errorf("unescaped char %s", string(l.input[l.pos-1]))
			}
		case '\\':
			if ok := runeIsEscape(l.peek()); !ok {
				return l.errorf("Not an escape character %s", strconv.QuoteRuneToASCII(l.next()))
			} else {
				l.next()
				break
			}
			fallthrough
		case eof, '\n':
			return l.errorf("unterminated character constant")
		case '\'':
			break Loop
		}
	}
	l.emit(CHARACTER)
	return lexInsideProgram
}

func lexString(l *Lexer) stateFn {
Loop:
	for {
		switch l.next() {
		case '\\':
			if r := l.next(); r != eof && r != '\n' {
				break
			}
			fallthrough
		case eof, '\n':
			return l.errorf("unterminated quoted string")
		case '"':
			break Loop
		}
	}
	l.emit(STRINGCONST)
	return lexInsideProgram
}

func lexNumber(l *Lexer) stateFn {
	l.accept("+-")
	digits := "0123456789"
	l.acceptRun(digits)
	l.emit(INTEGER)
	return lexInsideProgram
}

// Error is used by the yacc-generated parser to signal errors.
func (l *Lexer) Error(e string) {
	l.parseError = true
	line, col := l.currLocation()
	fmt.Printf("%q, %d:%d %s at or near %s\n", l.name, line, col, e, fmt.Sprintf(l.lastItem.Lexeme))
}

// Lex is used by the yacc-generated parser to fetch the next Lexeme.
func (l *Lexer) Lex(lval *parserSymType) int {
	token := l.NextItem()
	position := token.Pos
	switch token.Typ {
	case STRINGCONST:
		*lval = parserSymType{stringconst: Str(token.Lexeme), pos: position}
	case CHARACTER:
		*lval = parserSymType{character: Character(token.Lexeme), pos: position}
	case IDENTIFIER:
		*lval = parserSymType{ident: Ident(token.Lexeme), pos: position}
	case TRUE:
		*lval = parserSymType{boolean: Boolean(true), pos: position}
	case FALSE:
		*lval = parserSymType{boolean: Boolean(false), pos: position}
	case INTEGER:
		num, err := strconv.Atoi(token.Lexeme)
		if err != nil {
			fmt.Println(err)
			os.Exit(100)
		}
		if !checkInt(num) {
			fmt.Println("Int too big or small")
			os.Exit(100)
		}
		*lval = parserSymType{integer: Integer(num), pos: position}
	default:
		*lval = parserSymType{pos: position}
	}
	return token.Typ
}

func isInt(value interface{}) bool {
	switch value.(type) {
	case int:
		return true
	default:
		return false
	}
}

func checkInt(num int) bool {
	if num > math.MaxInt32 || num < math.MinInt32 {
		return false
	}
	return true
}

func runeIsEscape(a rune) bool {
	for _, b := range escapeChars {
		if b == a {
			return true
		}
	}
	return false
}

func checkClassIdent(classtype Ident) bool {
	a, _ := utf8.DecodeRuneInString(string(classtype))
	return unicode.IsUpper(a)

}

func checkStats(stats []Statement) bool {
	switch stats[len(stats)-1].(type) {
	case Return, Exit:
		return true
	case If:
		ifstat := stats[len(stats)-1].(If)
		return checkStats(ifstat.ThenStat) && checkStats(ifstat.ElseStat)
	default:
		return false
	}

}
