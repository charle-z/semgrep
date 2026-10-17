package interfacefuzz
type Bound197 interface{
 ~chan string | *int | Box[string]
}
func Scan197[T Bound197](v T, xs []int){
 auditSink(v)
 for k:=0;k<len(xs);k++ {println(xs[k])}
}
