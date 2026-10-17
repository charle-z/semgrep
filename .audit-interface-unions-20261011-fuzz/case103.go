package interfacefuzz
type Bound103 interface{
 ~map[string]bool | Pair[string,int] | ~chan string | ~*string
}
func Scan103[T Bound103](v T, xs []int){
 auditSink(v)
 for k:=0;k<len(xs);k++ {println(xs[k])}
}
