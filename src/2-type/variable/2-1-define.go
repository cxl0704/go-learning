package variable

import (
	"common"
)

var x=1000 //全局变量x
func Define001() {
	var x int //自动初始化为0
	var y=false //自动推断为bool类型
	println(common.PrintFuncName(),"===>",x,y)
}

func Define002() {
	var x,y int //相同类型的多个变量
	var a,s=200,"abc" //不同类型的初始化值
	println(common.PrintFuncName(),"===>",x,y,a,s)
}

func Define003() {
	//建议以组方式整理多行变量定义
	var (
		x,y int
		a,s=100,"abc"
	)
	println(common.PrintFuncName(),"===>",x,y,a,s)
}

//如下例子是简短定义的例子
func Define004() {
	//除了var关键字外，还可以使用更简短的变量定义和初始化语法
	x:=100
	a,s:=1,"abc"
	println(common.PrintFuncName(),"===>",x,a,s)
}

func Define005() {
	/*
	  需要注意，简短模式有些限制
	  1、定义变量同时显示初始化 2、不能提供数据类型 3、只能用在函数内部
	  对于新手这可能造成错误。比如原本打算修改全局变量，结果变成重新定义同名的局部变量
	  从这个例子可以看出，两个打印的内存地址是不同的，说明是连个不同的变量
	*/
	println(common.PrintFuncName(),"===>",&x,x) //全局变量
	x:=10 //重新定义和初始化同名的局部变量
	println(common.PrintFuncName(),"===>",&x,x)
}

func Define006() {
	//简短定义在函数多返回值，以及if/for/switch等语句中定义局部变量非常的方便
	//简短模式并不总是重新定义变量，也可能是部分退化的赋值操作
	x:=100
	println(common.PrintFuncName(),"===>",&x,x)

	x,y:=200,"abc" //注意：x退化为赋值操作，仅有y是变量定义
	println(common.PrintFuncName(),"===>",&x,x)
	println(common.PrintFuncName(),"===>",y)

	//对比变量的内存地址，可以确认x属于同一变量

}

//退化赋值的前提条件是：至少有一个新变量被定义，且必须同属一个作用域
func Define007() {
	x:=100
	println(common.PrintFuncName(),"===>",&x,x)

	//x:=200 //错误，no new Defines on left side of :=
	println(common.PrintFuncName(),"===>",&x,x)
}

func Define008() {
	x:=100
	println(common.PrintFuncName(),"===>",&x,x)

	{
		x,y:=200,300 //不同作用域，全部是新变量定义
		println(common.PrintFuncName(),"===>",&x,x,y)
	}
	
}

func Define009() {
	x,y:=1,2
	println(common.PrintFuncName(),"===>",x,y)
	x,y=y+3,x+2 //多变量赋值时，先计算出所有右值，然后再依次完成赋值操作
	println(common.PrintFuncName(),"===>",x,y)
}

//编译器将未使用的局部变量当作错误，全局变量没有问题，不会报错误

