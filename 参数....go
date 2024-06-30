// package main

// import "fmt"

// 1：无论是值传递，还是引用传递，传递给函数的都是变量的副本，不过，值传递是值的拷贝。引用传递是地址的拷贝，一般来说，地址拷贝更为高效。而值拷贝取决于拷贝的对象大小，对象越大，则性能越低。

// 2：map、slice、chan、指针、interface默认以引用的方式传递。

// 不定参数传值 就是函数的参数不是固定的，后面的类型是固定的。（可变参数）相当于python的 *args

// Golang 可变参数本质上就是 slice。只能有一个，且必须是最后一个。

// 在参数赋值时可以不用用一个一个的赋值，可以直接传递一个数组或者切片，特别注意的是在参数后加上“…”即可

// slice... 语法糖，将切片打散传入

// func Test(OptionsArgs ...int) {
// 	for _, v := range OptionsArgs {
// 		println(v)
// 	}
// }

// func main() {
// 	Test(1, 2, 3, 4, 5)
// 	Test([]int{1, 2, 3, 4, 5}...)                         // ... 展开slice传入
// 	fmt.Println("Hello, World!", [...]int{1, 2, 3, 4, 5}) // ... 表示不确定的长度
// }
