package interfacefuzz
type Bound159 interface{
 ~map[string]bool | ~[3]string | string | Box[int] | Pair[string,int]
}
type Alias159[T Bound159] = []T
func Scan159[T Bound159](v T, xs []int){
 auditSink(v)
 for k:=0;k<len(xs);k++ {println(xs[k])}
}
