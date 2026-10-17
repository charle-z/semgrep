package interfacefuzz
type Bound145 interface{
 Box[string]
}
func Scan145[T Bound145](v T, xs []int){
 auditSink(v)
 for k:=0;k<len(xs);k++ {println(xs[k])}
}
