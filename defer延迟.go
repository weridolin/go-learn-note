package main

// 1. 关键字 defer 用于注册延迟调用。
// 2. 这些调用直到 return 前才被执。因此，可以用来做资源清理。
// 3. 多个defer语句，按先进后出的方式执行。
// 4. defer语句中的变量，在defer声明时就决定了。

// import "fmt"

// func main() {
// 	var whatever [5]struct{}
// 	for i := range whatever {
// 		// 这里defer是闭包.i是指向外部变量的引用，所以最后i的值是5
// 		defer func() { fmt.Println(i) }() // 5个defer语句注册，但是i的值在return前就已经变成了5,这里defer语句中的i是引用，而不是值拷贝
// 	}
// }

import "fmt"

type Test struct {
	name string
}

func (t *Test) Close() {
	fmt.Println(t.name, " closed")
}

// func main() {
// 	ts := []Test{{"a"}, {"b"}, {"c"}}
// 	for _, t := range ts {
// 		fmt.Println(&t)
// 		defer t.Close() // t为指针，所以这里的t是引用，而不是值拷贝，defer后面的语句在执行的时候，函数调用的参数会被保存起来，但是不执行。也就是复制了一份。
// 	}
// }

// func main() {
// 	ts := []Test{{"a"}, {"b"}, {"c"}}
// 	for _, t := range ts {
// 		fmt.Println(&t)
// 		t2 := t
// 		defer t2.Close() // t2为值，所以这里的t2是值拷贝
// 	}
// }

// defer发生错误时，不影响后续的defer执行

// func errorFunc() {
// 	defer fmt.Println("defer 1")
// 	defer fmt.Println("defer 2")
// 	defer func() {
// 		panic("defer 3 error")
// 	}()
// 	defer func() {
// 		panic("defer 4 error")
// 	}()
// 	defer fmt.Println("defer 5")
// }
// 输出:
// defer 5
// defer 2
// defer 1
// panic: defer 4 error
//         panic: defer 3 error

// func main() {
// 	errorFunc()
// }

// 在多次循环里面，defer会比普通函数结尾的执行更加耗时间
