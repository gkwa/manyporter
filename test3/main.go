package test3

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/yuin/goldmark"
	meta "github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
)

func Run() {
	var mapData map[string]interface{}

	filename := "test3/example.md"
	markdownBytes, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	md := goldmark.New(goldmark.WithExtensions(meta.Meta))
	context := parser.NewContext()

	md.Parser().Parse(text.NewReader(markdownBytes), parser.WithContext(context))
	mapData = meta.Get(context)

	// Convert map to json string
	jsonStr, err := json.MarshalIndent(mapData, "", "  ")
	if err != nil {
		err1 := fmt.Errorf("marshalIndent failed: %w", err)
		fmt.Println(err1)
		return
	}

	// Output
	fmt.Println(string(jsonStr))
}
