package graphqlparser

import (
	"errors"
	"unicode/utf8"
)

const (
	UNIT_SEPARATOR = 0x001f
	TAB            = '\t'
	UNDERSCORE     = 95
)

type Lexer struct {
	source    string
	sourceLen int
	lastToken *Token
	token     *Token
	line      int
	lineStart int
	pos       int

	// options map[string]bool
}

func InitLexer(source string) *Lexer {
	return &Lexer{
		source:    source,
		sourceLen: len(source),
		lastToken: &Token{TK_SOF, 0, 0, 0, "", 0, nil, nil},
		token:     &Token{TK_SOF, 0, 0, 0, "", 0, nil, nil},
		line:      1,
		lineStart: 0,
		pos:       0,
	}
}

func (l *Lexer) get() rune {
	token, count := utf8.DecodeRuneInString(l.source[l.pos:])
	l.pos += count
	return token
}

func (l *Lexer) peek() (rune, int) {
	return utf8.DecodeRuneInString(l.source[l.pos:])
}

func (l *Lexer) nth(val int) rune {
	index := l.pos + val
	if index < l.sourceLen {
		c, _ := utf8.DecodeRuneInString(l.source[index:])
		return c
	}
	return 0
}

func (l *Lexer) skip(num int) {
	l.pos += num
}

func (l *Lexer) createToken(tokenKind TokenKind, start int) (*Token, error) {
	return &Token{tokenKind, start, l.pos, l.line, string(l.source[start:l.pos]), 0, nil, nil}, nil
}

func (l *Lexer) createSingleToken(tk TokenKind) (*Token, error) {
	start := l.pos
	l.pos += 1 // increment
	return &Token{tk, start, start + 1, l.line, string(tk), 0, nil, nil}, nil
}

func (l *Lexer) ReadToken() (*Token, error) {

	// l.positionAfterWhitespace()

	if l.pos >= l.sourceLen {
		return l.createSingleToken(TK_EOF)
	}

	c, _ := l.peek()
	switch c {
	case '!':
		return l.createSingleToken(TK_BANG)
		// 	case '#': // #
	// 		return l.readComment($line, $col, $prev)
	case '$':
		return l.createSingleToken(TK_STRING)
	case '&':
		return l.createSingleToken(TK_AMP)
	case '(':
		return l.createSingleToken(TK_PAREN_L)
	case ')':
		return l.createSingleToken(TK_PAREN_R)
	case '.':
		if l.nth(2) == '.' && l.nth(3) == '.' {
			return l.createToken(TK_SPREAD, l.pos+3)
		}
		break // why break?

		// 		case 34:
	// 				[, $nextCode]     = $this->readChar();
	// 				[, $nextNextCode] = $this->moveStringCursor(1, 1)->readChar();

	// 				if ($nextCode === 34 && $nextNextCode === 34) {
	// 						return $this->moveStringCursor(-2, (-1 * $bytes) - 1)
	// 								->readBlockString($line, $col, $prev);
	// 				}

	// 				return $this->moveStringCursor(-2, (-1 * $bytes) - 1)
	// 						->readString($line, $col, $prev);
	// }

	case ':':
		return l.createSingleToken(TK_COLON)
	case '=':
		return l.createSingleToken(TK_EQUALS)
	case '@':
		return l.createSingleToken(TK_AT)
	case '[':
		return l.createSingleToken(TK_BRACE_L)
	case ']':
		return l.createSingleToken(TK_BRACE_R)
	case '|':
		return l.createSingleToken(TK_PIPE)
	case '{':
		return l.createSingleToken(TK_BRACKET_L)
	case '}':
		return l.createSingleToken(TK_BRACKET_R)

	default:
		if isNameStart(c) {
			return l.readName()
		}

		if isNumberStart(c) {
			return l.readNumber()
		}
	}

	return nil, errors.New("Unaexpted character")
}

func (l *Lexer) readName() (*Token, error) {
	start := l.pos
	for l.pos < l.sourceLen {
		r, w := l.peek()
		if (r >= '0' && r <= '9') || (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z') || r == '_' {
			l.pos += w
		} else {
			break
		}
	}

	return l.createToken(TK_NAME, start)
}

func (l *Lexer) readNumber() (*Token, error) {
	start := l.pos
	isFloat := false
	c, _ := l.peek()
	if c == '-' {
		l.skip(1)
	}

	c, _ = l.peek()
	if c == '0' {
		l.skip(1)
		c, _ = l.peek()
		if isDigit(c) {
			return nil, errors.New(`Invalid number, unexpected digit after 0`)
		}
	} else {
		err := l.readDigits()
		if err != nil {
			return nil, err
		}
		c, _ = l.peek()
	}

	if c == '.' {
		isFloat = true
		l.skip(1)

		l.readDigits()
		c, _ = l.peek()
	}

	if c == 'E' || c == 'e' {
		isFloat = true
		c = l.get()
		// + -
		if c == '+' || c == '-' {
			c = l.get()
		}
		l.readDigits()
		c, _ = l.peek()
	}

	if c == '.' || isNameStart(c) {
		return nil, errors.New("Invalid number, expected digit but got . or name")
	}

	if isFloat {
		return l.createToken(TK_FLOAT, start)
	}

	return l.createToken(TK_INT, start)
}

func isDigit(code rune) bool {
	return code >= '0' && code <= '9'
}

func (l *Lexer) readDigits() error {
	c, _ := l.peek()
	if !isDigit(c) {
		return errors.New(`Invalid number, unexpected digit after 0`)
	}

	l.skip(1)

	for {
		c, _ = l.peek()
		if isDigit(c) {
			l.skip(1)
		} else {
			break
		}
	}

	return nil
}

func isNameStart(c rune) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || c == '_'
}

func isNumberStart(c rune) bool {
	return (c >= '0' && c <= '9') || c == '-'
}
