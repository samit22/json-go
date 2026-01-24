package cmd

import (
	"encoding/json"
	"fmt"
	"go/format"
	"strings"

	"github.com/atotto/clipboard"
	"github.com/samit22/json-go/logger"
	"github.com/spf13/cobra"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
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
		if len(args) == 0 {
			logger.Error("Please provide a JSON string as argument")
			return
		}
		generateStruct(args[0])
	},
}

func generateStruct(input string) (string, error) {
	logger.Info("Generating struct...\n")
	var data interface{}
	err := json.Unmarshal([]byte(input), &data)
	if err != nil {
		err := fmt.Errorf("invalid json err: %s", err)
		logger.Error(err.Error())
		return "", err
	}
	var strct string

	switch data.(type) {
	case map[string]interface{}:
		strct += "type AutoStruct struct {\n"
		data := handleMapStringInterface(data.(map[string]interface{}), "", 0)
		strct += data + "}\n"

	case []interface{}:
		for _, d := range data.([]interface{}) {
			switch d.(type) {
			case map[string]interface{}:
				strct += "type AutoStruct struct {\n"
				data := handleMapStringInterface(d.(map[string]interface{}), "", 0)
				strct += data + "}\n"
			default:
				msg := fmt.Sprintf("array is not a map but %T", d)
				logger.Error(msg)
				return "", fmt.Errorf("array is not a map but %T", d)
			}
		}
	}

	// Format the generated struct using go/format
	formatted, err := format.Source([]byte(strct))
	if err != nil {
		logger.Error(fmt.Sprintf("Failed to format struct: %s", err))
		fmt.Printf("\n%s", strct)
	} else {
		strct = string(formatted)
		logger.Info("Generated struct.")
		fmt.Printf("\n%s\n", strct)
	}

	if err := clipboard.WriteAll(strct); err == nil {
		logger.Success("Copied to clipboard.")
	}
	return strct, nil
}

func handleMapStringInterface(data map[string]interface{}, key string, nested int) string {
	var row string
	indent := strings.Repeat("\t", nested+1)
	if nested > 0 {
		row += fmt.Sprintf("%s%s struct {\n", strings.Repeat("\t", nested), GenerateAttrName(key))
	}
	for k, v := range data {
		switch v.(type) {
		case string, int, int32, int64, float32, float64, bool:
			row += generateRow(indent, k, fmt.Sprintf("%T", v))
		case map[string]interface{}:
			mpData := handleMapStringInterface(v.(map[string]interface{}), k, nested+1)
			row += mpData
		case []interface{}:
			for _, d := range v.([]interface{}) {
				row += generateRow(indent, k, fmt.Sprintf("[]%T", d))
				break // only need to check the first element to determine type
			}
		default:
			row += generateRow(indent, k, "interface{}")
		}
	}
	if nested > 0 {
		row += fmt.Sprintf("%s} `json:\"%s\"`\n", strings.Repeat("\t", nested), key)
	}
	return row
}

func generateRow(indent, key, tp string) string {
	return fmt.Sprintf("%s%s %s `json:\"%s\"`\n", indent, GenerateAttrName(key), tp, key)
}

func GenerateAttrName(inp string) string {
	str := strings.Split(inp, "_")
	var resp string
	caser := cases.Title(language.English)
	for _, v := range str {
		resp += caser.String(v)
	}
	return resp
}
