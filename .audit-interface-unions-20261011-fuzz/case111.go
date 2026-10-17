package interfacefuzz
type Bound111 interface{
 ~chan int | ~*string
}
type Alias111[T Bound111] = []T
func Scan111[T Bound111](v T, xs []int){
 auditSink(v)
 for k:=0;k<len(xs);k++ {println(xs[k])}
}
