package base

import "fmt"

func xor() {
	var a = 10
	var b = 20

	fmt.Println(a, b)

	//fmt.Println(a ^ a)  // 两数异或 = 0
	//fmt.Println(0 ^ a)  // 0和任何数异或 = 本身

	// 交换两个地址不相同的值，可以用下面这种;推理如下
	a = a ^ b
	b = a ^ b // b = a ^ b^b = a ^ 0 = a;
	a = a ^ b // a = a ^ b ^ a = a^a^b = 0^b = b;   => 至此，a和b的值已经互换；
	fmt.Println(a, b)

	// 两个十进制的数做与或非，要学会先将其转为二进制，再对位数进行操作
	var x = 10                   // x = 01010
	var y = 20                   // y = 10100
	fmt.Println("x & y = ", x&y) // 00000 = 0
	fmt.Println("x | y = ", x|y) // 11110 = 30
	fmt.Println("x ^ y = ", x^y) // 11110 = 30

	var m = 6                    // m = 0110
	var n = 10                   // n = 1010
	fmt.Println("m & n = ", m&n) // 0010 = 2
	fmt.Println("m | n = ", m|n) // 1110 = 14
	fmt.Println("m ^ n = ", m^n) // 1100 = 12

	var z = 10                     // 取出10最低位的1    1010
	fmt.Println("z & -z = ", z&-z) // 1010 &
}
