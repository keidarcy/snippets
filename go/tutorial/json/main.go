package main

import (
	"encoding/json"
	"fmt"
	// "strings"
)

type GoStruct struct {
	A string `json:"first"`
	B string `json:"second"`
}

// https://go.dev/blog/json
func main() {
	// encode(GoStruct{A: "aa", B: "bb"})
	// decode(`{"first":"1", "second": "2", "second": "1"}`)

	// decode any json
	// interface
	// var data interface{}
	// jsonString := `{"a": "aaa", "b": "bbb"}`
	// if err := json.Unmarshal([]byte(jsonString), &data); err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Printf("data: %v, type: %T\n", data, data)
	// m := data.(map[string]interface{})
	// for k, v := range m {
	// 	switch vv := v.(type) {
	// 	case string:
	// 		fmt.Println(k, "is string", vv)
	// 	case float64:
	// 		fmt.Println(k, "is float64", vv)
	// 	case []interface{}:
	// 		fmt.Println(k, "is an array:")
	// 		for i, u := range vv {
	// 			fmt.Println(i, u)
	// 		}
	// 	default:
	// 		fmt.Println(k, "is of a type I don't know how to handle")
	// 	}
	// }
	// any
	var data any
	jsonString := `{"a": "aaa", "b": "bbb"}`
	if err := json.Unmarshal([]byte(jsonString), &data); err != nil {
		fmt.Println(err)
	}
	fmt.Printf("data: %v, type: %T\n", data, data)

	m := data.(map[string]interface{})

	if a, ok := m["a"]; ok {
		fmt.Printf("a: %v\n", a)
		m["a"] = "dynamic"
	}
	fmt.Printf("m: %v\n", m)

}

func encode(s GoStruct) {
	// s := GoStruct{"a", "b"}

	// json.Marshal
	// fmt.Printf("s: %q\n", s)
	// jsonData, err := json.Marshal(s)
	// if err != nil {
	// 	panic("bad json")
	// }

	// fmt.Printf("jsonData: %s\n", jsonData)
	// fmt.Printf("jsonData: %v\n", string(jsonData))

	// json.NewEncoder
	// file, err := os.Create("./1.json")
	// if err != nil {
	// 	panic("bad json")
	// }
	// // jsonData, _ := json.Marshal(s)
	// // file.Write(jsonData)

	// json.NewEncoder(file).Encode(s)
	// err := json.NewEncoder(os.Stdout).Encode(s)
	// if err != nil {
	// 	panic("bad json")
	// }
	fmt.Println("====== struct -> json ======")
	jsonData, err := json.Marshal(s)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("jsonData: %q\n", string(jsonData))
}

func decode(jsonString string) {
	// var s GoStruct
	// jsonString := `{"A":1, "B": 2}`

	// // json.Unmarshal
	// if err := json.Unmarshal([]byte(jsonString), &s); err != nil {
	// 	fmt.Printf("s: %v\n", s)
	// 	fmt.Println("1")
	// 	return
	// }

	// fmt.Printf("s: %+v\n", s)

	// json.NewDecoder
	// var s GoStruct
	// jsonString := `{"A":"1", "B": "bbb"}`

	// if err := json.NewDecoder(strings.NewReader(jsonString)).Decode(&s); err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// fmt.Printf("s: %+v\n", s)
	fmt.Println("====== json -> struct ======")
	var s GoStruct
	if err := json.Unmarshal([]byte(jsonString), &s); err != nil {
		fmt.Println(err)

	}
	fmt.Printf("s: %+v\n", s)
}
