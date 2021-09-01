package graphqlparser

type Reader struct {
	source    string
	sourceLen int
	readerPos int
	startPos  int
	line      int
	lineStart int
}

func NewReader(source string) *Reader {
	return &Reader{
		source:    source,
		sourceLen: len(source),
	}
}

func (r *Reader) Peekahead(_index int) byte {
	index := _index + r.Pos()
	if index < r.sourceLen {
		return r.source[index]
	}
	return 0
}

func (r *Reader) Slice(start int, end int) string {
	return r.source[start:end]
}

func (r *Reader) Pos() int {
	return r.readerPos
}

func (r *Reader) LineNum() int {
	return r.line
}

func (r *Reader) Skip(val int) {
	r.readerPos += val
}

func (r *Reader) Increment() {
	r.readerPos += 1
}

func (r *Reader) IncrementLine() {
	r.lineStart = r.Pos()
	r.line += 1
}

func (r *Reader) CanRead(readerPos int) bool {
	return readerPos < r.sourceLen
}

func (r *Reader) hasSome() bool {
	return r.readerPos < r.sourceLen
}

func (r *Reader) IsEOF() bool {
	if r.IsEmptySource() {
		return true
	}

	return !r.CanRead(r.readerPos)
}

func (r *Reader) IsEmptySource() bool {
	return r.sourceLen == 0
}

func (r *Reader) Peek() byte {
	if r.IsEOF() {
		return 0
	}
	return r.source[r.readerPos]
}

func (r *Reader) PeekNext() byte {
	index := r.Pos() + 1
	if r.CanRead(index) {
		return r.source[index]
	}
	return 0
}

func (r *Reader) Get() byte {
	if r.IsEOF() {
		return 0
	}
	c := r.source[r.readerPos]
	r.readerPos += 1
	return c
}

func (r *Reader) SkipWhiteSpace() {
	for {
		c := r.Peek()
		if c == '\t' || c == ' ' || c == '\n' || c == '\r' || c == ',' {
			r.Increment()
			continue
		}
		break
	}
}

func (r *Reader) IsBom() bool {
	b := r.PeekMany(3)
	if len(b) != 3 {
		return false
	}

	return b[0] == 0xef && b[1] == 0xbb && b[2] == 0xbf
}

func (r *Reader) PeekMany(num int) []byte {
	ret := []byte{}
	for i := 0; i < num; i++ {
		index := r.Pos() + i
		if r.CanRead(index) {
			ret = append(ret, r.source[index])
		}
	}
	return ret
}
