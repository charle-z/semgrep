package interfacefuzz
type Bound131 interface{
 ~chan string | ~chan int
}
func Scan131[T Bound131](v T, xs []int){
 auditSink(v)
 for k:=0;k<len(xs);k++ {println(xs[k])}
}
