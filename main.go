package main

import (
	"fmt"

	"github.com/mcwz/gopretemplate/parseTemplate"
)

func main() {
	parse, err := parseTemplate.New("test.html")
	if err != nil {
		fmt.Println("err founded:", err)
	} else {
		parse.Parse()
	}
}
