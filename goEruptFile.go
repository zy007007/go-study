最近工作中需要完成的一个任务点：从数据库中获取某一列的数据(用户名)存入文件后，通过接口得到用户名和对应部门。
对于数据库导出数据，总结看到的各种文章后主要有三个方法：1.mysqldump 2.into oufile '/pathfile'(注意权限)3.mysql -e > /pathfile

在成功获得所需的初始数据之后，开始编写功能代码，从开始就决定使用go并发来完成，而不是顺序处理一条条数据得到结果。
之前有大概看到过go并发的一些代码，但也没有具体的实操过，不知道合格的go并发代码是什么样。
以下是今天通过查漏补缺，最终到达了预期目标：使用go并发处理一个文件内容

首先go并发在网上云搜索学习看到了主要是
1.直接使用go yourfunc()即可开启一个协程
2.在主进程结束前完成所有子任务

在抓住了这两大关键点之后，结合相关教程，目前个人实操后有两种方法。
1.sync.WaitGroup{}
sync的这个功能，就可以形象的理解为“等待组”，而多个子任务就可以看成一个组。waitgroup很好的提供了主进程等待所有子任务完成后，再结束整个程序。
func doThisMessage(user string, wg *sync.WaitGroup) {
	resp, err := doSomeThing(user)
	if err != nil {
		fmt.Println("error")
		wg.Done()           // 注意结束此次任务
		return
	}
	fmt.Println("得到期望的结果", resp)
	wg.Done()             // 同err上，结束此次任务
	return
}

func main() {
	wg := sync.WaitGroup{}     // 定义功能
	for _,v := range user_list {
		wg.Add(1)                // 任务新增1
		go doThisMessage(v, &wg)
	}
	wg.Wait()                  // 等待wg中没有任务，在结束主进程
}
可以看到，使用WaitGroup的关键在于处理一个子任务后，及时的使用Done()结束掉此任务。

2.channel
可以使用chan通道来传递子任务的数据，从而达到并发的效果。在使用通道时，最需要注意的就是通道的缓冲区大小。在实际使用的过程中，出现了如下情况。
func departmentName(user string, c chan string) {
	resp, err := doSomeThing(user)
	if err != nil {
		c <- fmt.Sprintf("error")     // 给这个通道传值
		return
	}
	c <- fmt.Sprintf("得到期望的结果", resp)   // 同err,给这个通道传值
	return
}


func main() {
	c := make(chan string, usernums)  // 文件有多少行，就创建多少缓冲区
  
	for i := 0; i < usernums; i++ {   
		go departmentName(content[i], c)  
	}

	for i := 0; i < cap(c); i++ {   // 从缓冲区里面，获取预期数据值
		result := <-c
		fmt.Println(result)
	}

	defer close(c)
}
使用chan的关键就在于对管道数据的传入和输出。

对于以上代码，存在的一个问题是缓冲区的大小过大，容易使程序运行时占用过多的内存。
那么如果不指定缓冲区大小呢，则就叫无缓冲区channel，说是每次管道中只能存有一个数据，那可不可以说无缓冲区就是大小为1的缓冲区呢？
如果使用无缓冲channel，那么上面的代码需要稍微进行修改
func main() {
	c := make(chan string)  // 创建无缓冲的管道
  
	for i := 0; i < usernums; i++ {   
		go departmentName(content[i], c)  
	}

	for i := 0; i < usernums; i++ {   // 这里循环需要变成文件行数
		result := <-c
		fmt.Println(result)
	}

	defer close(c)
}
在使用无缓冲的chan时，最后读取通道的数据时，保证所有的数据能被读到，因此需要指定文件行数。

那么，如果使用指定其他的缓冲区大小，结果又会是怎么样呢
func main() {
	c := make(chan string, 10)  // 指定10个缓冲区大小
  
	for i := 0; i < usernums; i++ {   
		go departmentName(content[i], c)  
	}

	for i := 0; i < usernums; i++ {   
		result := <-c
		fmt.Println(result)
	}

	defer close(c)
}
指定管道的缓冲区大小为10后，表示一个chan最多可以存放10个数值。

通过以上的学习和测试，我们可以看到实现go并发的两种方法。
那么接下来在回顾以上学习过程中，还需研究的地方有
1."即可开启一个协程" 到底什么是协程
2.此次处理的文件只有270行，并发处理符合预期。准备1k 10k 100k...的数据，进一步进行学习和测试
