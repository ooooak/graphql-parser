package graphqlparser

type ParserOptions struct {
	NoLocation                   bool
	AllowLegacyFragmentVariables bool
}

type Parser struct {
	Options *ParserOptions
	Reader  *Reader
}

func NewParser(source string, p *ParserOptions) *Parser {
	return &Parser{
		Options: p,
		Reader:  NewReader(source),
	}
}
