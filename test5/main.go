package test5

import (
	"fmt"
	"os"

	jsoniter "github.com/json-iterator/go"
	"github.com/yuin/goldmark"
	meta "github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
)

func Run() {
	filename := "test5/example.md"
	markdownBytes, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	md := goldmark.New(goldmark.WithExtensions(meta.Meta))
	context := parser.NewContext()

	md.Parser().Parse(text.NewReader(markdownBytes), parser.WithContext(context))
	metaData := meta.Get(context)

	jsonBytes, err := jsoniter.MarshalIndent(metaData, "", " ")
	if err != nil {
		fmt.Printf("error marshaling metadata: %v", err)
		return
	}

	fmt.Println(string(jsonBytes))
}
