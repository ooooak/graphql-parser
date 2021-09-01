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

func (l *Lexer) peek() (rune, int) {
	return utf8.DecodeRuneInString(l.source[l.pos:])
}

func (l *Lexer) get() rune {
	token, count := utf8.DecodeRuneInString(l.source[l.pos:])
	l.pos += count
	return token
}

func (l *Lexer) createToken(tokenKind TokenKind, start int) (*Token, error) {
	return &Token{tokenKind, start, l.pos, l.line, string(l.source[start:l.pos]), 0, nil, nil}, nil
}

func (l *Lexer) createSingleToken(tokenKind TokenKind) (*Token, error) {
	start := l.pos
	l.pos += 1 // increment
	return &Token{tokenKind, start, start + 1, l.line, string(TK_EOF), 0, nil, nil}, nil
}

func (l *Lexer) ReadToken() (*Token, error) {

	// l.positionAfterWhitespace()

	if l.pos >= l.sourceLen {
		return l.createSingleToken(TK_EOF)
	}

	c, _ := l.peek()
	switch c {
	case '!':
		return l.createSingleToken(TK_EOF)
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
		// 		case self::TOKEN_DOT: // .
		// 				[, $charCode1] = $this->readChar(true);
		// 				[, $charCode2] = $this->readChar(true);

		// 				if ($charCode1 === self::TOKEN_DOT && $charCode2 === self::TOKEN_DOT) {
		// 						return new Token(Token::SPREAD, $position, $position + 3, $line, $col, $prev);
		// 				}

		// 				break;

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

	case '_', 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i',
		'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't',
		'u', 'v', 'w', 'x', 'y', 'z', 'A', 'B', 'C', 'D', 'E',
		'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P',
		'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z':
		return l.readName()

	case '-', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		return l.readNumber()

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
