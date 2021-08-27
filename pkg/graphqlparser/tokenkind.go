package graphqlparser

type TokenKind string

const (
	TK_SOF          TokenKind = "<SOF>"
	TK_EOF          TokenKind = "<EOF>"
	TK_BANG         TokenKind = "!"
	TK_DOLLAR       TokenKind = "$"
	TK_AMP          TokenKind = "&"
	TK_PAREN_L      TokenKind = "("
	TK_PAREN_R      TokenKind = ")"
	TK_SPREAD       TokenKind = "..."
	TK_COLON        TokenKind = ":"
	TK_EQUALS       TokenKind = "="
	TK_AT           TokenKind = "@"
	TK_BRACKET_L    TokenKind = "["
	TK_BRACKET_R    TokenKind = "]"
	TK_BRACE_L      TokenKind = "{"
	TK_PIPE         TokenKind = "|"
	TK_BRACE_R      TokenKind = "}"
	TK_NAME         TokenKind = "Name"
	TK_INT          TokenKind = "Int"
	TK_FLOAT        TokenKind = "Float"
	TK_STRING       TokenKind = "String"
	TK_BLOCK_STRING TokenKind = "BlockString"
	TK_COMMENT      TokenKind = "Comment"
)
