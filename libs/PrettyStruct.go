package libs

import (
	"encoding/json"
	"fmt"
)

func PrintPrettyStruct(value interface{}) {
	valueJSON, err := json.MarshalIndent(value, "", "  ")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("%s\n", string(valueJSON))
}
