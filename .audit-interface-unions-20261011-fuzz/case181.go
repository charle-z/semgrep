package interfacefuzz
type Bound181 interface{
 ~chan int | Pair[string,int]
}
func Scan181[T Bound181](v T, xs []int){
 auditSink(v)
 for k:=0;k<len(xs);k++ {println(xs[k])}
}
