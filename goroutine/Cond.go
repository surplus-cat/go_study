package main

import (
	"fmt"
	"sync"
	"time"
)

func main()  {
	//3个人赛跑，1个裁判员发号施令
	cond := sync.NewCond(&sync.Mutex{})
	var wg sync.WaitGroup;
	wg.Add(4); //3选手+1裁判
	for i := 1; i <= 3; i++ {
		go func(num int)  {
			defer wg.Done()
			fmt.Println(num, "号选手已经就位")
			cond.L.Lock()
			cond.Wait() //等待发令枪响
			fmt.Println(num, "号选手开始跑……")
			cond.L.Unlock()

		}(i)
	}

	//等待所有goroutine都进入wait状态
	time.Sleep(2 * time.Second);
	go func ()  {
		defer wg.Done();
		fmt.Println("裁判：“各就各位~~预备~~”");
		fmt.Println("啪！。。")
		cond.Broadcast(); // 发令枪响
	}();
	//防止函数提前返回退出
	wg.Wait();
}


