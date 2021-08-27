package graphqlparser

type Lexer struct {
	Reader *Reader

	lastToken: Source;

  /**
   * The previously focused non-ignored token.
   */
  lastToken  Token;

  /**
   * The currently focused non-ignored token.
   */
  token: Token;

  /**
   * The (1-indexed) line containing the current token.
   */
  line: number;

  /**
   * The character offset at which the current line begins.
   */
  lineStart: number;
}
