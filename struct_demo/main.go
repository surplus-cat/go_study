package main

import (
	"fmt"
	"encoding/json"
)

//结构体

//Golang也支持面向对象编程(OOP)，
//但是和传统的面向对象编程有区别，并不是纯粹的面向对 象语言。所以我们说 Golang 支持面向对象编程特性是比较准确的。
//Golang没有类(class)，Go语言的结构体(struct)和其它编程语言的类(class)有同等的地位，
//你可以理解 Golang 是基于 struct 来实现 OOP 特性的。
//Golang面向对象编程非常简洁，去掉了传统OOP语言的继承、方法重载、构造函数和析构函数、隐藏的this指针等等
//Golang仍然有面向对象编程的继承，封装和多态的特性，只是实现的方式和其它OOP语言不
//一样，比如继承 :Golang 没有 extends 关键字，继承是通过匿名字段来实现。
//Golang面向对象(OOP)很优雅，OOP本身就是语言类型系统(typesystem)的一部分，通过接口 (interface)关联，耦合性低，也非常灵活。
//后面同学们会充分体会到这个特点。也就是说在 Golang 中面向接口编程是非常重要的特性。

type Cat struct {
	Name string;
	Age int;
	Color string;
	Hobby string;
}

type Person struct {
	Name string;
	Age int
}

type Point struct {
	x int;
	y int;
}

type Reat struct {
	leftUp, rightDown Point
}

type Reat2 struct {
	leftUp, rightDown *Point
}

type A struct {
	Num int;
}

type B struct {
	Num int;
}

type Student struct {
	Name string;
	Age int;
}

type integer int;

type Stu Student;

type Monster struct {
	Name string `json: "name"` //`json:"name"`就是struct tag
	Age int `json: "age"`
	Skill string `json: "skill"`
}


func main() {
	var cat1 Cat;
	cat1.Name = "xiaobai";
	cat1.Age = 10;
	cat1.Color = "xiaobai";
	cat1.Hobby = "eat";

	fmt.Println(cat1)

	//创建结构体变量和访问结构体字段
	//1.方式 1-直接声明
	//案例演示: var person Person 前面我们已经说了。
	//方式 2-{}
	p2 := Cat{"hhh", 20, "黑", "吃"};
	fmt.Println(p2);

	// 方法3-&
	var c4 *Cat = &Cat{};
	(*c4).Name = "dfsf";
	c4.Age = 111;
	c4.Color = "蓝";
	c4.Hobby = "飞";

	fmt.Println(c4)

	//第 3 种和第 4 种方式返回的是 结构体指针。
	//结构体指针访问字段的标准方式应该是:(*结构体指针).字段名 ，比如 (*person).Name = "tom"
	//但 go 做了一个简化，也支持 结构体指针.字段名, 比如 person.Name = "tom"。更加符合程序员
	//使用的习惯，go 编译器底层 对 person.Name 做了转化 (*person).Name。

	var p1 Person;
	p1.Age = 111;
	p1.Name = "Tim";

	var p3 *Person = &p1;
	fmt.Println((*p3).Age);
	fmt.Println(p3.Age);

	p3.Name = "Tony";
	fmt.Printf("p3.name=%v,p1.name=%v\n", p3.Name, p1.Name)//Tony Tony
	fmt.Printf("p3.name=%v,p3.name=%v\n", (*p3).Name, p1.Name) //Tony Tony

	fmt.Printf("p1的地址%p\n", &p1)
	fmt.Printf("p3的地址%p, p3的值%p\n", &p3, p3)

	//结构体的所有字段在内存中是连续的
	r1 := Reat{Point{1, 2}, Point{3, 4}};
	fmt.Printf(
		"r1.leftUp.x 地址=%p r1.leftUp.y 地址=%p, r1.rightDown.x 地址=%p r1.rightDown.y 地址=%p\n",
		&r1.leftUp.x, &r1.leftUp.y, &r1.rightDown.x, &r1.rightDown.y)

	//r2有两个 *Point类型，这两个*Point类型本身的地址是连续的。但是他们指向的地址不一定是连续的。
	r2 := Reat2{&Point{10,20}, &Point{30,49}};
	fmt.Printf("r2.leftUp 本身的地址是=%p r2.rightDown 本身的地址是%p\n", &r2.leftUp, &r2.rightDown)

	fmt.Printf("r2 .leftUp指向的地址是=%p, r2.rightDown指向的地址是=%p\n", r2.leftUp, r2.rightDown)

	//结构体是用户单独定义的类型，和其它类型进行转换时需要有完全相同的字段(名字、个数和类 型)
	var a A;
	var b B;

	a = A(b); //可以转换，但是有要求，就是结构体的字段要求完全一样，名字、个数和类型
	fmt.Println(a, b);

	//结构体进行 type 重新定义(相当于取别名)，Golang 认为是新的数据类型，但是相互间可以强转
	var stu2 Student;
	var stu3 Stu;

	stu3 = Stu(stu2); //不可以 stu3 = stu2
	fmt.Println(stu2, stu3)

	var i integer = 10;
	var j int = 20;
	j = int(i); // 不可以使用 j = i
	
	fmt.Println(i, j);

	//struct 的每个字段上，可以写上一个 tag, 该 tag 可以通过反射机制获取，常见的使用场景就是序列化和反序列化。
	monster := Monster{"牛", 2000, "芭蕉扇"};
	//将monster序列化为json字符串
	jsonStr, err := json.Marshal(monster)
	

	if err != nil {
		fmt.Println("处理错误");
	}

	fmt.Println("jsonStr", string(jsonStr))
}

//细节说明
//结构体和结构体变量(实例)的区别和联系
//1. 结构体是自定义的数据类型，代表一类事物.
//2. 结构体变量(实例)是具体的，实际的，代表一个具体变量

//如何声明结构体
//type 结构体名称 struct {
//field1 type
//field2 type }

//从概念或叫法上看: 结构体字段 = 属性 = field
//字段是结构体的一个组成部分，一般是基本数据类型、数组,也可是引用类型。
//比如我们前面定义猫结构体 的 Name string 就是属性

//注意事项
//1.字段声明语法同变量，示例:字段名 字段类型
//2.字段的类型可以为:基本类型、数组或引用类型
//3.在创建一个结构体变量后，如果没有给字段赋值，都对应一个零值(默认值)，规则同前面一样:
//布尔类型是 false ，数值是 0 ，字符串是 ""。
//数组类型的默认值和它的元素类型相关，比如 score [3]int 则为[0, 0, 0]
//指针，slice，和 map 的零值都是 nil ，即还没有分配空间。
//4.不同结构体变量的字段是独立，互不影响，一个结构体变量字段的更改，不影响另外一个, 结构体是值类型。