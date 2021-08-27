package graphqlparser

type Reader struct {
	source    string
	sourceLen int
	readerPos int
	startPos  int
}

func NewReader(source string) *Reader {
	return &Reader{
		source:    source,
		sourceLen: len(source),
		readerPos: 0,
		startPos:  0,
	}
}

func (r *Reader) Pos() int {
	return r.readerPos
}

func (r *Reader) Increment() {
	r.readerPos += 1
}

func (r *Reader) CanRead(readerPos int) bool {
	return readerPos < r.sourceLen
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

func (r *Reader) PeekMany(num int) string {
	ret := []byte{}
	for i := 0; i < num; i++ {
		index := r.Pos() + i
		if r.CanRead(index) {
			ret = append(ret, r.source[index])
		}
	}
	return string(ret)
}
