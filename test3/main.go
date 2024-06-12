package test3

import (
	"encoding/json"
	"fmt"
)

func Run() {
	// map data
	mapData := map[string]interface{}{
		"Name":    "noknow",
		"Age":     2,
		"Admin":   true,
		"Hobbies": []string{"IT", "Travel"},
		"Address": map[string]interface{}{
			"PostalCode": 1111,
			"Country":    "Japan",
		},
		"Null": nil,
	}

	// Convert map to json string
	jsonStr, err := json.MarshalIndent(mapData, "", "  ")
	if err != nil {
		fmt.Println(err)
	}

	// Output
	fmt.Println(string(jsonStr))
}
