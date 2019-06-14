package main

import (
	"fmt"
	"time"
)

func main1() {

	//var intVal int
	//
	//intVal :=3 // 这时候会产生编译错误
	//
	//intVal,intVal1 := 1,2 // 此时不会产生编译错误，因为有声明新的变量，因为 := 是一个声明语句
	//
	//fmt.Println(&intVal, intVal1)
	//
	//fmt.Println("Google" + "Runoob")

	//const LENGTH int = 10
	//const WIDTH int = 5
	//var area int
	//const a, b, c = 1, false, "str" //多重赋值
	//
	//area = LENGTH * WIDTH
	//fmt.Printf("面积为 : %d", area)
	//println()
	//println(a, b, c)

	//
	//const (
	//	a = 5   //0
	//	b = iota         //1
	//	c          //2
	//	d = "ha"   //独立值，iota += 1
	//	e          //"ha"   iota += 1
	//	f = 100    //iota +=1
	//	g          //100  iota +=1
	//	h = iota   //7,恢复计数
	//	i          //8
	//)
	//fmt.Println(a,b,c,d,e,f,g,h,i)
	//
	//var a int = 21
	//var b int = 10
	//var c int
	//
	//c = a + b
	//fmt.Printf("第一行 - c 的值为 %d\n", c )
	//c = a - b
	//fmt.Printf("第二行 - c 的值为 %d\n", c )
	//c = a * b
	//fmt.Printf("第三行 - c 的值为 %d\n", c )
	//c = a / b
	//fmt.Printf("第四行 - c 的值为 %d\n", c )
	//c = a % b
	//fmt.Printf("第五行 - c 的值为 %d\n", c )
	//a++
	//fmt.Printf("第六行 - a 的值为 %d\n", a )
	//a=21   // 为了方便测试，a 这里重新赋值为 21
	//a--
	//fmt.Printf("第七行 - a 的值为 %d\n", a )

	//var a int = 21
	//var b int = 10
	//
	//if( a == b ) {
	//	fmt.Printf("第一行 - a 等于 b\n" )
	//} else {
	//	fmt.Printf("第一行 - a 不等于 b\n" )
	//}
	//if ( a < b ) {
	//	fmt.Printf("第二行 - a 小于 b\n" )
	//} else {
	//	fmt.Printf("第二行 - a 不小于 b\n" )
	//}
	//
	//if ( a > b ) {
	//	fmt.Printf("第三行 - a 大于 b\n" )
	//} else {
	//	fmt.Printf("第三行 - a 不大于 b\n" )
	//}
	///* Lets change value of a and b */
	//a = 5
	//b = 20
	//if ( a <= b ) {
	//	fmt.Printf("第四行 - a 小于等于 b\n" )
	//}
	//if ( b >= a ) {
	//	fmt.Printf("第五行 - b 大于等于 a\n" )
	//}

	//var a bool = true
	//var b bool = false
	//if ( a && b ) {
	//	fmt.Printf("第一行 - 条件为 true\n" )
	//}
	//if ( a || b ) {
	//	fmt.Printf("第二行 - 条件为 true\n" )
	//}
	///* 修改 a 和 b 的值 */
	//a = false
	//b = true
	//if ( a && b ) {
	//	fmt.Printf("第三行 - 条件为 true\n" )
	//} else {
	//	fmt.Printf("第三行 - 条件为 false\n" )
	//}
	//if ( !(a && b) ) {
	//	fmt.Printf("第四行 - 条件为 true\n" )
	//}

	//var a uint = 60      /* 60 = 0011 1100 */
	//var b uint = 13      /* 13 = 0000 1101 */
	//var c uint = 0
	//
	//c = a & b       /* 12 = 0000 1100 */
	//fmt.Printf("第一行 - c 的值为 %d\n", c )
	//
	//c = a | b       /* 61 = 0011 1101 */
	//fmt.Printf("第二行 - c 的值为 %d\n", c )
	//
	//c = a ^ b       /* 49 = 0011 0001 */
	//fmt.Printf("第三行 - c 的值为 %d\n", c )
	//
	//c = a << 2     /* 240 = 1111 0000 */
	//fmt.Printf("第四行 - c 的值为 %d\n", c )
	//
	//c = a >> 2     /* 15 = 0000 1111 */
	//fmt.Printf("第五行 - c 的值为 %d\n", c )

	//var a int = 4
	//var b int32
	//var c float32
	//var ptr *int
	//
	///* 运算符实例 */
	//fmt.Printf("第 1 行 - a 变量类型为 = %T\n", a );
	//fmt.Printf("第 2 行 - b 变量类型为 = %T\n", b );
	//fmt.Printf("第 3 行 - c 变量类型为 = %T\n", c );
	//
	///*  & 和 * 运算符实例 */
	//ptr = &a     /* 'ptr' 包含了 'a' 变量的地址 */
	//fmt.Printf("a 的值为  %d\n", a);
	//fmt.Printf("ptr 为 %d\n", ptr);
	//fmt.Printf("*ptr 为 %d\n", *ptr);

	//print9x()
	//gotoTag()

	///* 定义局部变量 */
	//var a int = 100
	//var b int = 200
	//var ret int
	//
	///* 调用函数并返回最大值 */
	//ret = max(a, b)
	//
	//fmt.Printf( "最大值是 : %d\n", ret )

	//var balance = [][]int{
	//	{1,2,3,4,5},
	//	{1,2,3,4,5},
	//}
	//
	//fmt.Println(balance[1][3])

	//var a int= 20   /* 声明实际变量 */
	//var ip *int        /* 声明指针变量 */
	//
	//ip = &a  /* 指针变量的存储地址 */
	//
	//fmt.Printf("a 变量的地址是: %x\n", &a  )
	//
	///* 指针变量的存储地址 */
	//fmt.Printf("ip 变量储存的指针地址: %x\n", ip )
	//
	///* 使用指针访问值 */
	//fmt.Printf("*ip 变量的值: %d\n", *ip )

	go say("world")
	say("hello")
}

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

/* 函数返回两个数的最大值 */
//func max(num1, num2 int) int {
//	/* 声明局部变量 */
//	var result int
//
//	if (num1 > num2) {
//		result = num1
//	} else {
//		result = num2
//	}
//	return result
//}

////嵌套for循环打印九九乘法表
//func print9x() {
//	for m := 1; m < 10; m++ {
//		for n := 1; n <= m; n++ {
//			fmt.Printf("%dx%d=%d ",n,m,m*n)
//		}
//		fmt.Println("")
//	}
//}
//
////for循环配合goto打印九九乘法表
//func gotoTag() {
//	for m := 1; m < 10; m++ {
//		n := 1
//	LOOP: if n <= m {
//		fmt.Printf("%dx%d=%d ",n,m,m*n)
//		n++
//		goto LOOP
//	} else {
//		fmt.Println("")
//	}
//		n++
//	}
//}
