package main

//设计一个支持 push，pop，top 操作，并能在常数时间内检索到最小元素的栈。
//
//push(x) -- 将元素 x 推入栈中。
//pop() -- 删除栈顶的元素。
//top() -- 获取栈顶元素。
//getMin() -- 检索栈中的最小元素。
//示例:
//
//MinStack minStack = new MinStack();
//minStack.push(-2);
//minStack.push(0);
//minStack.push(-3);
//minStack.getMin();   --> 返回 -3.
//minStack.pop();
//minStack.top();      --> 返回 0.
//minStack.getMin();   --> 返回 -2.

// 个人思路
// 栈,让我现在能想到的是，就是对切片的操作了

// 参考练习了解题的一个，过了一遍，感觉是为了获取最小的值而增大了切片的长度

type MinStack struct {
	min   int
	datas []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		datas: make([]int, 0),
		min:   math.MaxInt64,
	}
}

func (this *MinStack) Push(x int) {
	if x <= this.min {
		this.datas = append(this.datas, this.min)
		this.min = x
	}
	this.datas = append(this.datas, x)
}

func (this *MinStack) Pop() {
	element := this.datas[len(this.datas)-1:][0]
	this.datas = this.datas[:len(this.datas)-1]
	if element == this.min {
		this.min = this.datas[len(this.datas)-1:][0]
		this.datas = this.datas[:len(this.datas)-1]
	}
}

func (this *MinStack) Top() int {
	return this.datas[len(this.datas)-1:][0]
}

func (this *MinStack) GetMin() int {
	return this.min
}

// 收获:1.可以适当增加容量，让目标前后有标记，方便定位；2.切片中[:]的操作，又新增了个完成项：整理一篇[:]的相关博客吧
