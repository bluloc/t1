package p1

import (
	"encoding/json"
	"fmt"
)

func F1(input string) string {
	fmt.Println(input)
	x := struct {
		Input string `json:"inpt"`
	}{
		input,
	}
	b, _ := json.Marshal(x)
	return string(b)
}
