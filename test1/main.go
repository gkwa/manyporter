package test1

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
	filename := "test1/example.md"
	markdownBytes, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	md := goldmark.New(goldmark.WithExtensions(meta.Meta))
	context := parser.NewContext()

	md.Parser().Parse(text.NewReader(markdownBytes), parser.WithContext(context))
	metaData := meta.Get(context)

	convertedMetaData := convertMap(metaData)

	jsonBytes, err := json.MarshalIndent(convertedMetaData, "", "  ")
	if err != nil {
		fmt.Printf("error marshalling metadata: %v", err)
		return
	}

	fmt.Println(string(jsonBytes))
}

func convertMap(data interface{}) interface{} {
	switch value := data.(type) {
	case map[string]interface{}:
		converted := make(map[string]interface{})
		for k, v := range value {
			converted[k] = convertMap(v)
		}
		return converted
	case map[interface{}]interface{}:
		converted := make(map[string]interface{})
		for k, v := range value {
			strKey := fmt.Sprintf("%v", k)
			converted[strKey] = convertMap(v)
		}
		return converted
	default:
		return data
	}
}
