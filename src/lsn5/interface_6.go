package main

import "fmt"

func classifier(items ...interface{}) {
	for _, v := range items {
		switch v.(type) {
		case bool:
			fmt.Println("bool type")
		case int, int64, int32:
			fmt.Println("int")
		case float32, float64:
			fmt.Println("float32")
		case string:
			fmt.Println("string")
		}
	}
}

func main() {
	classifier(2, 8.2, "ok string", false)
}
