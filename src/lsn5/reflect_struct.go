package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type Student struct {
	Name  string  `json:"name"`
	Age   int     `json:"age"`
	Score float32 `json:"score"`
	Sex   string  `json:"sex"`
}

func (s Student) Print() {
	fmt.Println(s)
}

func OpStruct(inter interface{}) {
	val := reflect.ValueOf(inter)
	kind := val.Kind()

	if kind != reflect.Ptr && val.Elem().Kind() == reflect.Struct {
		fmt.Println("expect struct")
		return
	}

	num := val.Elem().NumField()

	for i := 0; i < num; i++ {
		fmt.Printf("%d %s ", i, val.Elem().Field(i).Kind())
	}
	numMethod := val.Elem().NumMethod()
	fmt.Println("Num ", num, numMethod)

	//var params []reflect.Value
	val.Elem().Method(0).Call(nil)

	tp := reflect.TypeOf(inter)
	for i := 0; i < num; i++ {
		tag := tp.Elem().Field(i).Tag.Get("json")
		fmt.Println("tag", tag)
	}
}

func main() {
	stu := Student{
		Name:  "ming",
		Age:   23,
		Score: 34.5,
	}
	fmt.Println(stu)
	OpStruct(&stu)

	result, _ := json.Marshal(stu)
	fmt.Println("json", string(result))
}
