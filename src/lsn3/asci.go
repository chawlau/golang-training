package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "http:://www.baidu.com"
	fmt.Println("%v", strings.HasPrefix(str, "https"))
	fmt.Println("%v", strings.HasSuffix(str, "com"))
}
