package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
)

/*

算法与操作对象的结构分离的一种方法
https://en.wikipedia.org/wiki/Visitor_pattern
*/

type Visitor func(shape Shape)

type Shape interface {
	accept(visitor Visitor)
}

type Circle struct {
	Radius int
}

func (c Circle) accept(v Visitor) {
	v(c)
}

type Rectangle struct {
	Width, Height int
}

func (r Rectangle) accept(v Visitor) {
	v(r)
}

// JsonVisitor /*
func JsonVisitor(shape Shape) {
	bytes, err := json.Marshal(shape)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}

// XmlVisitor /*
func XmlVisitor(shape Shape) {
	bytes, err := xml.Marshal(shape)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}

func main() {
	c := Circle{10}
	rectangle := Rectangle{100, 200}
	shapes := []Shape{c, rectangle}
	for _, shape := range shapes {
		shape.accept(JsonVisitor)
		shape.accept(XmlVisitor)
	}

}
