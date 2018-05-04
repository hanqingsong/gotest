package main

import "fmt"

func main()  {
	arrayCopyTest()
}

func arrayCopyTest(){

	var arr1 =[...]int{1,3,5,4}
	fmt.Println(arr1)

	// 数组赋值，直接复制数组，修改互补影响
	arr2:=arr1
	fmt.Println(arr2)

	fmt.Printf("%p\n",&arr1)
	fmt.Printf("%p\n",&arr2)
	arr2[2]=88
	fmt.Println(arr1)
	fmt.Println(arr2)
	// 数组地址不同
	fmt.Printf("%p\n",&arr1)
	fmt.Printf("%p\n",&arr2)
	// 元素地址不同
	fmt.Printf("%p\n",&arr1[2])
	fmt.Printf("%p\n",&arr2[2])

	fmt.Println("=========指针赋值==========")
	arr3 := &arr1
	fmt.Println(*arr3)

	fmt.Printf("%p\n",&arr1)
	fmt.Printf("%p\n",&arr3)
	arr3[2]=88
	fmt.Println(arr1)
	fmt.Println(*arr3)
	//数组地址不同
	fmt.Printf("%p\n",&arr1)
	fmt.Printf("%p\n",&arr3)
	// 元素地址相同
	fmt.Printf("%p\n",&arr1[2])
	fmt.Printf("%p\n",&arr3[2])
}