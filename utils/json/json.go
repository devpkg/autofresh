package json

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func ToJSON(obj interface{}) string {
	str, err := json.MarshalIndent(obj, "", "    ")
	if err != nil {
		fmt.Println(err.Error())
	}
	return string(str)
}

func FromJSON(str string) interface{} {
	d := json.NewDecoder(bytes.NewReader([]byte(str)))
	var obj interface{}
	d.Decode(&obj)
	return obj
}
