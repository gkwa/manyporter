package core

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
	filename := "example.md"
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

func convertMap(data map[string]interface{}) map[string]interface{} {
	converted := make(map[string]interface{})
	for k, v := range data {
		switch value := v.(type) {
		case map[interface{}]interface{}:
			converted[k] = convertInterfaceMap(value)
		default:
			converted[k] = v
		}
	}
	return converted
}

func convertInterfaceMap(data map[interface{}]interface{}) map[string]interface{} {
	converted := make(map[string]interface{})
	for k, v := range data {
		converted[fmt.Sprintf("%v", k)] = v
	}
	return converted
}
