package graphqlparser

type Token struct {
	Kind   TokenKind
	Start  int
	End    int
	Line   int
	Value  string
	Column int
	Prev   *Token
	Next   *Token
}
