package main

import "fmt"

// 接口 声明一系列方法的集合
// /    空接口可以作为任何类型数据的容器。

// 抽象类声明
type People interface {
	Speak(string) string
}

// 派生类实现具体逻辑
type Student struct{}

func (stu *Student) Speak(think string) (talk string) {
	if think == "sb" {
		talk = "你是个大帅比"
	} else {
		talk = "您好"
	}
	return
}

func main() {

	var peo People = &Student{} // peo是一个人类，具体是什么人,  抽象类 ---> 派生类
	think := "bitch"
	fmt.Println(peo.Speak(think))
}
