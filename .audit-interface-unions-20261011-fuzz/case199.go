package interfacefuzz
type Bound199 interface{
 Box[int] | **struct{ A int } | Pair[string,int] | ~struct{ A int } | ~[2]int
}
func Scan199[T Bound199](v T, xs []int){
 auditSink(v)
 for k:=0;k<len(xs);k++ {println(xs[k])}
}
