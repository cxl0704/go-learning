package constant

import (
	"common"
	"unsafe"
	"fmt"
)

/* 常量
	1、常量表示运行时恒定不可改变的值，通常是一些字面量。
	2、常量必须是编译期可确定的字符、字符串、数字或者布尔值。可指定常量类型，或由编译器通过初始化值推断
	3、可在函数代码块定义常量，不曾使用的常量不会引发编译错误

*/
func Define001()  {
	const x,y int=123,0x22
	const s = "hello world"
	const c = "我"
	println(common.PrintFuncName(),"===>",x,y,s,c)

	const (
		i, f = 1, 0.123 //int, float64(默认)
		b = false
	)
	println(common.PrintFuncName(),"===>",i,f,b)

	const z = 1234 //未使用，不会引发编译错误

	{
		const x = "abc" //在不同的作用域定义同名常量
		println(common.PrintFuncName(),"===>",x)
	}
	println(common.PrintFuncName(),"===>",x)
	
	{
		const (
			x,y int = 99, -999
			b byte = byte(x)  //如果是显示指定类型，必须保证左右类型一致，需要时可以做显示转换
			//n = uint8(y) // 右值不能超出常量类型取值范围，否则会引发溢出错误。constant -999 overflows uint8
		)
		println(common.PrintFuncName(),"===>",b)
	}

	//常量也可以是某些编译器能计算出结果的表达式，如：unsafe.Sizeof、len、cap等
	{
		const (
			ptrSize = unsafe.Sizeof(uintptr(0))
			strSize = len("hello world!")
		)
		println(common.PrintFuncName(),"===>ptrSize,strSize",ptrSize,strSize)
	}

	//在常量组中，如不指定类型和初始化值，则与上一行非空常量右值（表达式文本）相同
	{
		const (
			x uint16 = 120
			y				//与上一行x类型、右值相同
			s = "abc"
			z				//与上一行s类型、右值相同
			t				//与s类型、右值相同
		)
		fmt.Printf("%v,===> y %T, %v\n",common.PrintFuncName(),y,y)
		fmt.Printf("%v,===> z %T, %v\n",common.PrintFuncName(),z,z)
		fmt.Printf("%v,===> t %T, %v\n",common.PrintFuncName(),t,t)
	}
}

//go并没有明确意义上的enum定义，不过可以借助iota标识符实现一组自增变量值来实现枚举类型
func Define002()  {
	const (
		x = iota 	//0
		y			//1
		z			//2
	)
	fmt.Printf("%v,===> x %T, %v\n",common.PrintFuncName(),x,x)
	fmt.Printf("%v,===> y %T, %v\n",common.PrintFuncName(),y,y)
	fmt.Printf("%v,===> z %T, %v\n",common.PrintFuncName(),z,z)


	const (
		_ = iota			//0
		KB = 1<<(10*iota)	//1<<(10*1)
		MB					//1<<(10*2)
		GB					//1<<(10*3)
	)
	fmt.Printf("%v,===> KB %T, %v\n",common.PrintFuncName(),KB,KB)
	fmt.Printf("%v,===> MB %T, %v\n",common.PrintFuncName(),MB,MB)
	fmt.Printf("%v,===> GB %T, %v\n",common.PrintFuncName(),GB,GB)

	//自增作用范围为常量组。可以在多常量定义中使用多个iota，它们各自单独计数，只需确保组中每行常量的列数相同即可。
	{
		const (
			_, _ = iota, iota*10	//0, 0*10
			a, b					//1, 1*10
			c, d					//2. 2*10
		)
		fmt.Printf("%v,===> a %T, %v\n",common.PrintFuncName(),a,a)
		fmt.Printf("%v,===> b %T, %v\n",common.PrintFuncName(),b,b)
		fmt.Printf("%v,===> c %T, %v\n",common.PrintFuncName(),c,c)
		fmt.Printf("%v,===> d %T, %v\n",common.PrintFuncName(),d,d)
		
	}

	//如果中断iota自增，则必须显示恢复。且后续自增值按行序递增
	{
		const (
			a = iota		//0
			b				//1
			c = 100			//100
			d				//100 与上一行常量右值表达式相同
			e = iota		//4 恢复iota自增，计数包括c、d
			f				//5
		)
		println(common.PrintFuncName(),"===> ",a,b,c,d,e,f)
	}

	//自增默认类型为int，可以显示指定类型
	{
		const (
			a = iota //int
			b float32 = iota //float32
			c //如果不显示指定类型，则与b类型相同
			d = iota //int
		)
		fmt.Printf("%v,===> a %T, %v\n",common.PrintFuncName(),a,a)
		fmt.Printf("%v,===> b %T, %v\n",common.PrintFuncName(),b,b)
		fmt.Printf("%v,===> c %T, %v\n",common.PrintFuncName(),c,c)
		fmt.Printf("%v,===> d %T, %v\n",common.PrintFuncName(),d,d)
	}

	//变量和常量的区别
	//不同与变量在运行期间分配内存（非优化状态），常量通常会被编译器在预处理阶段直接展开，作为指令数据使用
	var x1 = 0x100
	const y1 = 0x200
	{
		println(common.PrintFuncName(),"x1 ===> ",&x1,x1)
		//println(common.PrintFuncName(),"y1 ===> ",&y1,y1) //cannot take the address of y1
	}

	{
		const x =100		//无类型声明的常量
		const y byte = x	//直接展开下，相当于const y byte = 100
		const a int = 100	//显示指定类型，编译器会强制类型检查
		//const b byte = a	//错误，cannot use a (type int) as type byte in const initializer
	}
}