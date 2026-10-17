package interfacefuzz
type Bound219 interface{
 Pair[int,string] | **struct{ A int } | int | ~chan string | Box[int]
}
type Alias219[T Bound219] = []T
func Scan219[T Bound219](v T, xs []int){
 auditSink(v)
 for k:=0;k<len(xs);k++ {println(xs[k])}
}
