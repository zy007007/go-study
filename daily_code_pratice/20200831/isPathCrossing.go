func isPathCrossing(path string) bool {
	var res int

	for _, v := range path {
		switch w {
		case "N":
			res += 1
		case "S":
			res += -1
		case "E":
			res += 2
		case "W":
			res += -2
		}
	}

	return res == 0
}	




type Point struct {
	x int
	y int
}

func (p *Point) Walk(x, y int) {
	p.x += x
	p.y += y
}

func isPathCrossing(path string) bool {

	mp := map[rune][2]int{
		'E':[2]int{1, 0},
		'W':[2]int{-1, 0},
		'N':[2]int{0, 1},
		'S':[2]int{0, -1}
	}

	visited := map[Point]bool{}
	point := Point{0, 0}
	visited[point] = true
	for _, v := range path {
		point.Walk(mp[v][0], mp[v][1])
		if visited[point] {
			return true
		}
		visited[point] = true
	}

	return false
}




select a,b,c,d from table1 left join table2 on table1.AimId == table2.AimId


inner join
left join
right join


select Email from Person group by Email having count(Email) > 1;

select ( select sal from Sal order by sal DESC limit 1 offset 1 ) as NewSal;


头尾一起找


func searchRange(nums []int, target int) []int {
	var res []int

	for k, v := range nums {
		if v==target {
			res = append(res, k)
		}
		if len(res) > 1 {
			start = res[0]
		}
	}

	if len(res) > 1 {
		end := res[-1]
	} else if len(res) == 1{
		start, end := res[0], res[0]
	} else {
		start, end := -1, -1
	}

}



sort.Ints()



select id, jan, ..., dec from Department 



select class from courses group by class having count(class) > 5;


select class from courses group by class having count(DISTINCT student) >= 5;