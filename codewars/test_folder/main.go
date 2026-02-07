package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "go go go"
	fmt.Println(strings.ReplaceAll(str, "go", "no"))

	str2 := "gogogo"
	fmt.Println(strings.ReplaceAll(str2, "go", "no"))
}
