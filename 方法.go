// package main

// 方法总是绑定对象实例，并隐式将实例作为第一实参 (receiver)
// 方法只能在同一个包内定义

// func (receiver type) methodName(参数列表)(返回值列表){}

// type InstanceType struct {
// 	Width, Height int
// }

// func (i InstanceType) Area() int {
// 	return i.Width * i.Height
// }

// func (i *InstanceType) Perimeter() int {
// 	return 2 * (i.Width + i.Height)
// }

// 方法的接收者可以是值类型或指针类型

// 1. 如果接收者是值类型，那么它只能接收值类型
// 2. 如果接收者是指针类型，那么它既能接收值类型，也能接收指针类型

// func main() {
// 	i := InstanceType{10, 20} // 其实也可以调用指针类型的方法？
// 	println(i.Area())
// 	println(i.Perimeter())

// 	p := &InstanceType{10, 20}
// 	println(p.Area())
// 	println(p.Perimeter())

// 	fmt.Println(i, p)
// }

// 相当于集成调用
// • 类型 T 方法集包含全部 receiver T 方法。
// • 类型 *T 方法集包含全部 receiver T + *T 方法。
// • 如类型 S 包含匿名字段 T，则 S 和 *S 方法集包含 T 方法。
// • 如类型 S 包含匿名字段 *T，则 S 和 *S 方法集包含 T + *T 方法。
// • 不管嵌入 T 或 *T，*S 方法集总是包含 T + *T 方法。

// type Type1 struct {
// }

// type Type2 struct {
// 	Type1
// }

// func (t *Type1) Method1() {
// 	fmt.Println("Method1")
// }

// func (t Type1) Method2() {
// 	fmt.Println("Method2")
// }

// // 相当于重新Method1
// func (t *Type2) Method1() {
// 	fmt.Println("Method1 override")
// }

// func main() {
// 	i := Type2{}
// 	i.Method1() // 重写Type1的Method1
// 	i.Method2() // 集成Type1的Method2
// }

// 方法调用两种方式
// instance.method(args...) ---> <type>.func(instance, args...)

// type Data struct{}

// func (Data) TestValue() {}

// func (*Data) TestPointer() {}

// func main() {
// 	var p *Data = nil
// 	p.TestPointer()

// 	(*Data)(nil).TestPointer() // method value ## 先类型转换，在调用转换的类型所属方法
// 	(*Data).TestPointer(nil)   // method expression  ## 先调用方法，再传入实例参数

// p.TestValue()            // invalid memory address or nil pointer dereference

// (Data)(nil).TestValue()  // cannot convert nil to type Data
// Data.TestValue(nil)      // cannot use nil as type Data in function argument
// }
