/*
Copyright © 2022 Samit <samitghimire@gmail.com>
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// v4Cmd represents the v4 command
var jsonToGo = &cobra.Command{
	Use:   "jtg",
	Short: "Generates json to go structs",
	Long: `Generates json to go structs.:
It takes json object as input and generates go structs.
Example: json-to-go jtg '{"name":"Samit","age":22}'
`,
	Run: func(cmd *cobra.Command, args []string) {
		generateStruct(args[0])
	},
}

func generateStruct(input string) (string, error) {
	var data interface{}
	err := json.Unmarshal([]byte(input), &data)
	if err != nil {
		fmt.Printf("Invalid json input: error: %v\n", err)
		return "", err
	}
	var strct string

	switch data.(type) {
	case map[string]interface{}:
		strct += "type AutoStruct struct {\n"
		data := handleMapStringInterface(data.(map[string]interface{}), "", 0)
		strct += data + " }\n"
		fmt.Printf("Genererated struct \n\n %s", strct)
	case []interface{}:
	}
	return "", nil
}

func handleMapStringInterface(data map[string]interface{}, key string, nested int) string {
	var row string
	if nested > 0 {
		row += fmt.Sprintf("\t %s struct {\n", GenreateAttrName(key))
	}
	idxRpt := generateIdx(nested)
	for k, v := range data {
		switch v.(type) {
		case string, int, int32, int64, float32, float64:
			row += generateRow(idxRpt, k, fmt.Sprintf("%T", v))
		case map[string]interface{}:
			mpData := handleMapStringInterface(v.(map[string]interface{}), k, nested+1)
			row += mpData

		default:
			row += generateRow(idxRpt, k, "interface{}")
		}
	}
	if nested > 0 {
		row += fmt.Sprintf("%s } `json:\"%s\"`\n", idxRpt, key)
	}
	return row
}

func generateRow(idx, key, tp string) string {
	return fmt.Sprintf("\t%s %s %s `json:\"%s\"`\n", idx, GenreateAttrName(key), tp, key)
}

func GenreateAttrName(inp string) string {
	str := strings.Split(inp, "_")
	var resp string
	for _, v := range str {
		resp += strings.Title(v)
	}
	return resp
}

func generateIdx(inp int) string {
	return strings.Repeat("\t", inp)
}
