package main
 
import(
    "fmt"
    //"sort"
)
 
func RemoveDuplicatesAndEmpty(a []int64) (ret []int64){
    a_len := len(a)
    for i:=0; i < a_len; i++{
        if (i > 0 && a[i-1] == a[i]){
            continue;
        }
        ret = append(ret, a[i])
    }
    return
}
 
func main(){
	ids := []int64{11,22,33,22,11,333,444,22,11,33}
	tmp := make(map[int64]int64)
	for _, v := range ids {
		if tmp[v] != 0 {
			continue
		} 
		tmp[v] = v
	}

	var res []int64
	for _, v := range tmp {
		res = append(res, v)
	}
	fmt.Println(res)
}
