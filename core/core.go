package core

import (
	"fmt"
	"os"

	"github.com/yuin/goldmark"
	meta "github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
)

func Run() {
	filename := "example.md"
	markdownBytes, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	md := goldmark.New(goldmark.WithExtensions(meta.Meta))
	context := parser.NewContext()

	md.Parser().Parse(text.NewReader(markdownBytes), parser.WithContext(context))
	metaData := meta.Get(context)
	fmt.Printf("%#v\n", metaData)
}
