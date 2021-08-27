package graphqlparser

type Token struct {
	Kind TokenKind

	// The character offset at which this Node begins.
	Start int
	// The character offset at which this Node ends.
	End int

	/**
	 * The 1-indexed line number on which this Token appears.
	 */
	Line int

	/**
	 * The 1-indexed column number at which this Token begins.
	 */
	Column int

	/**
	 * For non-punctuation tokens, represents the interpreted value of the token.
	 *
	 * Note: is undefined for punctuation tokens, but typed as string for
	 * convenience in the parser.
	 */
	Value string

	/**
	 * Tokens exist as nodes in a double-linked-list amongst all tokens
	 * including ignored tokens. <SOF> is always the first node and <EOF>
	 * the last.
	 */
	Prev *Token
	Next *Token
}

func InitToken(kind TokenKind, start int, end int, line int, column int) *Token {
	return &Token{
		Kind:   kind,
		Start:  start,
		End:    end,
		Line:   line,
		Column: column,
		Value:  "",
		Prev:   &Token{},
		Next:   &Token{},
	}
}
