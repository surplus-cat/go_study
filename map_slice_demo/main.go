package main

import (
	"fmt"
	"sort"
)

//map切片
//切片的数据类型如果是 map，则我们称为 slice of map，map 切片，这样使用则 map个数就可以动态变化了。
func main()  {
	var monsters []map[string]string
	monsters = make([]map[string]string, 2)	

	if (monsters[0] == nil) {
		monsters[0] = make(map[string]string, 2)
		monsters[0]["name"] = "牛"
		monsters[0]["age"] = "500"
	}

	if (monsters[1] == nil) {
		monsters[0] = make(map[string]string, 2)
		monsters[0]["name"] = "玉"
		monsters[0]["age"] = "400"
	}

	newMoster := map[string]string{
		"name": "猪头三",
		"age": "300",
	};

	monsters = append(monsters, newMoster);
	fmt.Println(monsters);

	//map排序
	//1.golang中没有一个专门的方法针对map的key进行排序
	//2.golang中的map默认是无序的，注意也不是按照添加的顺序存放的，你每次遍历，得到的输出 可能不一样
	//3.golang中map的排序，是先将key进行排序，然后根据key值遍历输出即可

	map1 := make(map[int]int, 10);
	map1[1] = 10;
	map1[10] = 4;
	map1[4] = 12;

	fmt.Println(map1);

	var keys []int
	for k, _ := range(map1) {
		keys = append(keys, k)
	}

	sort.Ints(keys);
	fmt.Println(keys);

	for _, k := range(keys) {
		fmt.Printf("map1[%v]=%v\n", k, map1[k])
	}

	users := make(map[string]map[string]string, 10);
	users["smith"] = make(map[string]string, 3);
	users["smith"]["pwd"] = "99999"
	users["smith"]["nick"] = "haha"
	users["smith"]["name"] = "hanpi";

	modifyUser(users, "tom")
	modifyUser(users, "mary")
	modifyUser(users, "smith")

	fmt.Println(users);

}

//细节说明
//1. map是引用类型，遵守引用类型传递的机制，在一个函数接收map，修改后，会直接修改原来的 map 
//2. map 的容量达到后，再想 map 增加元素，会自动扩容，并不会发生 panic，也就是说 map 能动 态的增长 键值对(key-value)
//3. map 的 value 也经常使用 struct 类型，更适合管理复杂的数据(比前面 value 是一个 map 更好)，比如 value 为 Student 结构体
func modifyUser(users map[string]map[string]string, name string)  {
	
	if (users[name] != nil) {
		users["name"]["pwd"] = "11111";
	} else {
		users["name"] = make(map[string]string, 2);
		users["name"]["pwd"] = "22222";
		users["name"]["nick"] = "牛呀"
	}
}