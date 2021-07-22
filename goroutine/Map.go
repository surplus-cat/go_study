package main

import (
	"fmt"
	"sync"
)

func main()  {
	
	var syMap sync.Map;
	// sync.Map 没有获取 map 数量的方法，可以在 遍历的时候自行计算数量，
	// sync.Map 为了保证并发安全，牺牲了一些性能，如果没有并发场景，推荐使用内置的 map 类。

	// 将键值对保存到sync.Map
	syMap.Store("aaa", 111);
	syMap.Store("bbb", 222);
	syMap.Store("ccc", 333);
	fmt.Println(syMap.LoadOrStore("dddd", 444));

	// 从sync.Map中根据键取值
	fmt.Println(syMap.Load("aaa"));

	// 根据键删除对应的键值对
	syMap.Delete("aaa");

	// 遍历所有sync.Map中的键值对
	syMap.Range(func(k, v interface{}) bool {
		fmt.Println("k:", k , "=> v:", v);
		return true;
	})
}