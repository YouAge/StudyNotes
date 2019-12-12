package main

import "fmt"

func arraySum(x [3]int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}

func main() {
	// 定义切片
	a1 := [...]int{1, 2, 3, 4, 5, 6}

	// 第二种方式
	var slicts []int = []int{1, 2, 3, 4}
	fmt.Println(slicts)

	//第三种
	c2 := make([]int, 0)

	sum := 0
	for _, v := range a1 {
		sum = sum + v
	}
	fmt.Println(sum)
	// 传入参数必须有定义
	/// 函数固定参数
	a := [3] int{1, 2, 3}
	pric := arraySum(a)
	fmt.Println(pric)

	//基于数组定义 切片, cap,容量 从切的第一个元素到最后
	// 切边不保存值，都是用底层数组保存
	a2 := [5]int{44, 55, 66, 77, 99}
	b1 := a2[1:5] // 左闭右开 4-1 =3 ，角标123
	c1 := b1[0:2]
	c1 = append(c1, 343, 324523, 32432, 22, 43634636, 2345, 235, 342, 23)
	// append追加切片, 不能是数组
	c1 = append(c1, b1...)
	fmt.Println(b1)
	fmt.Println("type of b:%T", b1)
	fmt.Println(c1)
	fmt.Println(len(b1), len(a2), cap(b1), cap(c1))

	fmt.Println(c2, cap(c2), append(c2, 23, 32, 43, 65, 7))
	c3 := make([]string, 3, 10)
	fmt.Println(cap(c3), len(c3), append(c3, "w", "f", "g", "g", "g", "g", "g", "g", "r", "ds"), cap(c3), len(c3), c3)

	// 删除切片中的元素， 切片没有删除功能， 只能利用切片本身的特性删除
	c1 = append(c1[1:])
	fmt.Println(c1)

	// TODO 切片在传递时， 属于浅拷贝，  引用使用， 都会改变
	var a3 = make([]string, 5, 10)
	for i := 0; i < 10; i++ {
		a3 = append(a3, fmt.Sprintf("%v", i))
	}
	fmt.Println(a3)

	// 切片在内存中的布局

	// string 底层是一个byte数组，因此string也可以进行切片

	// TODO 修改字符串， 需要先转byte类型，修改后再转回string
	// 转成byte后 可以处理中文数字， 但是不能处理中文，    中文是3个字节 会出现乱码
	// 解决方法 string 转成 【】rune即可  因为【】rune兼容汉字 安字符串处理的
	str3 := "aove"
	arr3 := []byte(str3)
	//arr3:=[]rune(str3)
	arr3[0] = 'l'
	//arr3[0]='我'
	str3 = string(arr3)
	fmt.Println(str3)
	//
	fmt.Println(Fbn(10))
}

// 能够将斐波那契数列到切片中
func Fbn(n int) ([]uint64) {
	fbSlice := make([]uint64, n)
	fbSlice[0] = 1
	fbSlice[1] = 1
	for i := 2; i < n; i++ {
		fbSlice[i] = fbSlice[i-1] + fbSlice[i-2]
	}
	return fbSlice

}
