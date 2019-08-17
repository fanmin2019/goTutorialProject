//每个 Go 程序都是由包构成的
//程序从 main 包开始运行
package main

//此代码用圆括号组合了导入，这是“分组”形式的导入语句。
import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func basic1() {
	//*注意：* 此程序的运行环境是固定的，因此 rand.Intn 总是会返回相同的数字。
	// （要得到不同的数字，需为生成器提供不同的种子数，参见 rand.Seed。 练习场中的时间为常量，因此你需要用其它的值作为种子数。）
	rand.Seed(time.Now().UnixNano())
	fmt.Println("My favorite number2 is", rand.Intn(10))
}

func basic2() {
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
}

//注意类型在变量名 之后
func add(x int, y int) int {
	return x + y
}

//当连续两个或多个函数的已命名形参类型相同时，除最后一个类型以外，其它都可以省略
//x int, y int 被缩写为x, y int
func add2(x, y int) int {
	return x + y
}

func basic3() {
	//在 Go 中，如果一个名字以大写字母开头，那么它就是已导出的
	//在导入一个包时，你只能引用其中已导出的名字。任何“未导出”的名字在该包外均无法访问。
	//NG : math.pi
	fmt.Println(math.Pi)
}

func basic4() {
	fmt.Println(add2(42, 13))
}

//函数可以返回任意数量的返回值。
//function_name(variable1 type1, variable2 type2)(return type1, return type2)
func swap(x, y string) (string, string) {
	return y, x
}

//Go 的返回值可被命名，它们会被视作定义在函数顶部的变量。
//返回值的名称应当具有一定的意义，它可以作为文档使用
//没有参数的 return 语句返回已命名的返回值。也就是 直接 返回
func split(sum int) (y, x int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func basic6() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}

func basic7() {
	fmt.Println(split(17))
}

//var 语句用于声明一个变量列表，跟函数的参数列表一样，类型在最后
//var 语句可以出现在包或函数级别
var c, python, java string

//var c, python, java bool

func basic8() {
	var i int
	fmt.Println(i, c, python, java)
}

func main() {
	//basic1()
	//basic2()
	//basic3()
	//fmt.Println(add2(42, 13))
	//basic6()
	basic8()
}
