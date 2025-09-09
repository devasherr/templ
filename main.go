package main

import (
	"fmt"
	"github.com/devasherr/templ/parser"
	"io"
	"os"
)

type Templ struct {
	writer io.Writer
}

func NewTempl() *Templ {
	return &Templ{
		writer: os.Stdout,
	}
}

func (t *Templ) Print(a ...any) (int, error) {
	str := fmt.Sprint(a...)
	return t.writer.Write([]byte(str))
}

func (t *Templ) Printf(format string, a ...any) (int, error) {
	p_format, p_any := parser.Parse(format, a)
	str := fmt.Sprintf(p_format, p_any...)
	return t.writer.Write([]byte(str))
}

func (t *Templ) Sprintf(format string, a ...any) string {
	p_format, p_any := parser.Parse(format, a)
	return fmt.Sprintf(p_format, p_any...)
}

func main() {
	templ := NewTempl()
	name := "bob"
	templ.Printf("this is {bob}\n", name)

	age := 6
	msg := templ.Sprintf("this is {bob} and he is {age} years old\n", name, age)
	templ.Print(msg)
}
