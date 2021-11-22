package main

import "fmt"

func main() {
	var v int = 1
	fmt.Printf("# Main frame: Value of v:\t\t %v, address: %p\n", v, &v)
	PassValue(v, &v)
	fmt.Printf("# Main frame: Value of v:\t\t %v, address: %p\n", v, &v)
}

func PassValue(fv int, addV *int) {
	// fv 的地址只属于该函数, 由该函数栈分配
	fmt.Printf("# Func frame: Value of fv:\t\t %v, address: %p\n", fv, &fv)
	// 本次修改只在该函数上生效
	fv = 0
	fmt.Printf("# Func frame: Value of fv:\t\t %v, address: %p\n", fv, &fv)

	/*
	 *	根据main函数传入的全局地址, 对指针执行操作外部是可见的,
	 *  因为改指针操作的都是同一个内存块的内容
	 */
	*addV = 3
	fmt.Printf("# Func frame: Value of addV:\t %v, address: %p\n", *addV, addV)
}
