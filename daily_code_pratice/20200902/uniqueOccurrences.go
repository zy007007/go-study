func uniqueOccurrences(arr []int) bool {
	dic1 := make(map[int]int)

	for _, v := range arr {
		dic1[v] += 1
	}

	res := make(map[int]bool)

	for _, v := range dic1 {
		if res[v] {
			return false
		} else {
			res[v] = true
		}
	}
	return true
}



type Stack struct {
	Top int // 栈顶
	Cap int // 容量
	Data *[]interface   //指向的指针，或者说是当前值
}


func (s *Stack) StackInit(cap int) {
	s.Top = 0
	s.Cap = cap
	m := make([]interface{}, cap)
	s.Data = m
}


// 入栈	
func (s *Stack) Push(data interface{}) int {
	//s.Top += 1
	//s.Data = append(s.Data, data)
	if s.Full() {
		fmt.Printf("stack is full")
		return 0
	} 

	(*s.Data)[s.Top] = data
	s.Top ++
	return s.Top
}


// 出栈
func (s *Stack) Pop() (interface{}) {
	if s.Empty() {
		fmt.Printf("stack is null")
		return nil
	}

	s.Top --
	em := (*s.Data)[s.Top]
	return em
}