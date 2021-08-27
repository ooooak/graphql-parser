package graphqlparser_test

import (
	"fmt"
	"log"
	"testing"

	"github.com/ooooak/apj/pkg/graphqlparser"
)

func TestInit(t *testing.T) {
	reader := graphqlparser.NewReader("type")
	if reader.Pos() != 0 {
		t.Fatal("should be 0")
	}
}

func TestEmpty(t *testing.T) {
	r := graphqlparser.NewReader("")
	if !r.IsEmptySource() {
		log.Fatal("source should be empty")
	}

	r2 := graphqlparser.NewReader("type")
	if r2.IsEmptySource() {
		log.Fatal("should not be empty")
	}
}

func TestGet(t *testing.T) {
	r := graphqlparser.NewReader("")
	if r.Get() != 0 {
		log.Fatal("get should be zero")
	}

	r2 := graphqlparser.NewReader("type")
	if r2.Pos() != 0 {
		log.Fatal("pos should be zero")
	}

	if r2.Get() != 't' {
		log.Fatal("Should be t")
	}
	if r2.Pos() != 1 {
		log.Fatal("pos should be 1")
	}
	if r2.Get() != 'y' {
		log.Fatal("Should be y")
	}
	if r2.Pos() != 2 {
		log.Fatal("pos should be 2")
	}

	if r2.Get() != 'p' {
		log.Fatal("Should be p")
	}
	if r2.Pos() != 3 {
		log.Fatal("pos should be 3")
	}

	if r2.Get() != 'e' {
		log.Fatal("Should be e")
	}
	if r2.Pos() != 4 {
		log.Fatal("pos should be 4")
	}

	if r2.Get() != 0 {
		log.Fatal("should be an empty byte")
	}

	if r2.Pos() != 4 {
		log.Fatal("pos should be 4")
	}
}

func TestPeek(t *testing.T) {
	r := graphqlparser.NewReader("")
	if r.Peek() != 0 {
		log.Fatal("should be zero")
	}
	if r.Pos() != 0 {
		log.Fatal("should be zero")
	}

	r2 := graphqlparser.NewReader("type")
	if r2.Pos() != 0 {
		log.Fatal("should be zero")
	}
	if r2.Peek() != 't' {
		log.Fatal("should be t")
	}

	if r2.Pos() != 0 {
		log.Fatal("should be zero")
	}
	if r2.Peek() != 't' {
		log.Fatal("should be t")
	}

	r2.Get() // skip t

	if r2.Pos() != 1 {
		log.Fatal("should be 1")
	}

	if r2.Peek() != 'y' {
		log.Fatal("should be y")
	}
}

func TestSkipWhiteSpace(t *testing.T) {
	r := graphqlparser.NewReader("   , ,, ,, , t")
	if r.Get() != ' ' {
		log.Fatal("should be white space")
	}

	r.SkipWhiteSpace()

	if r.Get() != 't' {
		log.Fatal("should be t")
	}
}

func TestPeekMany(t *testing.T) {
	r := graphqlparser.NewReader("12345")
	if r.Get() != '1' {
		log.Fatal("should be 1")
	}

	if r.PeekMany(3) != "234" {
		log.Fatal("should be 234 " + r.PeekMany(3))
	}

	if r.Get() != '2' {
		log.Fatal("should be 2")
	}
	if r.Get() != '3' {
		log.Fatal("should be 3")
	}

	if r.PeekMany(3) != "45" {
		log.Fatal("should be 45 ")
	}
	if r.PeekMany(1) != "4" {
		log.Fatal("should be 4")
	}
	if r.PeekMany(1) != "4" {
		log.Fatal("should be 4")
	}
	if r.PeekMany(10) != "45" {
		log.Fatal("should be 45")
	}
}

func BenchmarkPeekMany(b *testing.B) {
	r := graphqlparser.NewReader("12345")
	fmt.Println(b.N)
	for n := 0; n < b.N; n++ {
		r.PeekMany(3)
	}
}
