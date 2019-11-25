package main

import (
	"2-type/variable"
	"2-type/constant"
)

func main()  {
	variableDefine() //2-1 变量定义
	variableName() //2-2 变量命名
	constantDefine() //2-3 常量
	
}

func constantDefine(){
	constant.Define001()
	constant.Define002()
}

func variableName(){
	variable.Name()
}

func variableDefine(){
	variable.Define001()
	variable.Define002()
	variable.Define003()
	variable.Define004()
	variable.Define005()
	variable.Define006()
	variable.Define007()
	variable.Define008()
	variable.Define009()
}